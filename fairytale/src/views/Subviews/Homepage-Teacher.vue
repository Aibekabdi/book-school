<template>
  <Notification class="notification" v-if="onError" :message="onError" />
  <div class="grid teacher">
    <h3 class="heading-3">Teacher View</h3>
    <div class="grid">
      <nav class="row justify">
        <div class="row">
          <Button
            type="Button"
            @click="toggleModal('CreateStudent'), getTeacherClasses()"
            label="Создать ученика"
            icon="add_circle"
            color="primary"
          />
          <Button
            type="Button"
            @click="toggleModal('EditClass'), getClassForTeacher()"
            label="Настройка класса"
            icon="settings"
            color="second"
          />
          <Button
            type="Button"
            @click="toggleModal('CheckStats')"
            label="Открыть статистику"
            icon="insights"
            color="success"
          />
        </div>
        <div class="row">
          <Button
            type="Router"
            link="/profile"
            label="Личный кабинет"
            icon="account_circle"
            color="warning"
          />
        </div>
      </nav>
      <LeaderDashboard />
      <Feed class="feed" />
    </div>
    <div class="modal" v-if="open">
      <div>
        <div class="close">
          <Button type="Button" @click="toggleModal()" icon="close" :rounded="true" color="black" />
        </div>
        <div class="modal-content" v-if="modalType == 'CreateStudent'">
          <div>
            <h3 class="heading-2">Создать ученика</h3>
            <p class="text-4">Введите ФИО, логин, год рождения, язык обучения и пароль ученика</p>
            <span v-if="isSuc" class="success-message">Ученик был успешно создан</span>
            <span v-else-if="isErr" class="error-message">Ошибка. Попробуйте снова</span>
          </div>
          <section>
            <component
              v-for="item in teacherForm.student"
              :key="item"
              :is="Input"
              :label="item.label"
              :name="item.name"
              :type="item.type"
              v-model:value="item.ref"
              :placeholder="item.placeholder"
              :error="item.error"
            />
            <div class="grid">
              <span class="heading-3">Выберите класс:</span>
              <v-select :options="teacherClasses" v-model="selectedOptionField"></v-select>
            </div>
            <Button
              type="Button"
              @click="createStudent"
              label="Создать"
              icon="add"
              color="success"
              width="100%"
            />
          </section>
          <p class="text-5">
            Здесь вы можете создать ученика, <a href="#">посмотрите гайд как это работает</a>
          </p>
        </div>
        <div class="modal-content" v-if="modalType == 'EditClass'">
          <div>
            <h3 class="heading-2">Настройка классов</h3>
            <p class="text-4">
              Здесь задаются задания для разных классов, а также просматривают данные
            </p>
          </div>
          <div class="grid pre-dash">
            <TeacherDashboard :school="school" />
          </div>
        </div>
        <div class="modal-content" v-if="modalType == 'CheckStats'">
          <Statistics />
        </div>
      </div>
    </div>
  </div>
</template>
<script setup>
import Button from '@/components/Button-component.vue';
import Input from '@/components/Input-component.vue';
import Feed from '@/components/Feed-component.vue';
import TeacherDashboard from '@/components/TeacherDashboard-component.vue';
import Statistics from '@/components/Statistics-component.vue';
import LeaderDashboard from '@/components/LeaderDashboard-component.vue';
import Notification from '@/components/layout/Notification-component.vue';
import vSelect from 'vue-select';
import 'vue-select/dist/vue-select.css';
import { sendRequest } from '@/utils/utils';
import { user, userToken, onError, currentLocalization } from '@/App.vue';
import { ref, reactive, watch } from 'vue';
</script>
<script>
const isSuc = ref(false);
const isErr = ref(false);

const open = ref(false);
const modalType = ref('');

const selectedOptionField = ref('');
const school = ref('');
const teacherClasses = ref([]);

const teacherForm = reactive({
  student: [
    {
      ref: '',
      label: 'Имя ученика',
      name: 'first-name',
      type: 'text',
      placeholder: 'Введите имя ученика',
      rules: [
        {
          id: 1,
          f: (val) => val.length > 0,
          msg: 'Это поле обязательно',
        },
      ],
      error: [],
    },
    {
      ref: '',
      label: 'Фамилия ученика',
      name: 'second-name',
      type: 'text',
      placeholder: 'Введите фамилию ученика',
      rules: [
        {
          id: 2,
          f: (val) => val.length > 0,
          msg: 'Это поле обязательно',
        },
      ],
      error: [],
    },
    {
      ref: '',
      label: 'Год рождения',
      name: 'data-birth',
      type: 'text',
      placeholder: 'Введите год рождения',
      rules: [
        {
          id: 3,
          f: (val) => val.match(/^-?\d+$/),
          msg: 'Поле должно состоять из цифр',
        },
        {
          id: 4,
          f: (val) => val.length == 4,
          msg: 'Поле должно содержать 4 символа',
        },
      ],
      error: [],
    },
    {
      ref: '',
      label: 'Логин',
      name: 'login',
      type: 'text',
      placeholder: 'Введите логин ученика',
      rules: [
        {
          id: 5,
          f: (val) => val.length >= 3 || val.length == 0,
          msg: 'Длина логина должна быть больше 3 символов',
        },
      ],
      error: reactive([]),
    },
    {
      ref: '',
      label: 'Пароль',
      name: 'password',
      type: 'password',
      placeholder: 'Введите ваш пароль',
      rules: [
        {
          id: 6,
          f: (val) => val.length >= 3 || val.length == 0,
          msg: 'Длина пароля должна быть больше 3 символов',
        },
      ],
      error: reactive([]),
    },
    {
      ref: '',
      label: 'Подтвердите пароль',
      name: 'confirm-password',
      type: 'password',
      placeholder: 'Подтвердите ваш пароль',
      rules: [
        {
          id: 7,
          f: (val) => val === teacherForm.student[4].ref,
          msg: 'Пароли не совпадают',
        },
      ],
      error: reactive([]),
    },
  ],
});
export default {
  mounted() {
    user.value = localStorage.getItem('admin') ? 'admin' : localStorage.getItem('userType') || '';
    userToken.value = localStorage.getItem('admin') || localStorage.getItem('userToken') || '';
    selectedOptionField.value = '';

    watch(
      teacherForm,
      () => {
        for (const key in teacherForm) {
          for (const item of teacherForm[key]) {
            if (item.rules) {
              for (const rule of item.rules) {
                if (!rule.f(item.ref)) {
                  if (item.error.find((el) => el.$uid === rule.id)) {
                    continue;
                  }
                  item.error.push({ $uid: rule.id, $message: rule.msg });
                } else if (item.error.length > 0) {
                  const idx = item.error.findIndex((el) => el.$uid === rule.id);
                  item.error.splice(idx, 1);
                }
              }
            }
          }
        }
      },
      { deep: true },
    );
  },
  methods: {
    toggleModal(modal) {
      isErr.value = false;
      isSuc.value = false;
      modalType.value = modal;
      open.value = !open.value;

      if (!open.value) {
        modalType.value = '';
      }
    },

    // teacher view
    createStudent() {
      isSuc.value = false;
      isErr.value = false;

      const json = {
        first_name: teacherForm.student[0].ref,
        second_name: teacherForm.student[1].ref,
        // date_birth: teacherForm.student[2].ref,
        username: teacherForm.student[3].ref,
        password: teacherForm.student[4].ref,
        class_id: +selectedOptionField.value.split('.')[0],
      };

      sendRequest('/api/student/create', 'POST', json, userToken.value)
        .then(() => {
          isSuc.value = true;
        })
        .catch((err) => {
          isErr.value = true;
        });
    },
    // teacher view
    getTeacherClasses() {
      sendRequest('/api/class/all', 'GET', null, userToken.value)
        .then((data) => {
          if (data) {
            data.classes.forEach((el, i) => {
              teacherClasses.value[i] = el.class.id + '. ' + el.class.grade + ' ' + el.class.name;
            });
          }
        })
        .catch((err) => {
          onError.value = 'Невозможно получить данные классов, пожалуйста повторите позднее';
          console.log(err);

          setTimeout(() => {
            onError.value = null;
          }, 5000);
        });
    },
    // teacher view
    getClassForTeacher() {
      isSuc.value = false;
      isErr.value = false;

      sendRequest('/api/class/all', 'GET', null, userToken.value)
        .then((data) => {
          if (data) {
            school.value = data;
            console.log(data);
          }
        })
        .catch((err) => {
          onError.value =
            'Невозможно получить данные для отображения, пожалуйста повторите позднее';
          console.log(err);

          setTimeout(() => {
            onError.value = null;
          }, 5000);
        });
    },
  },
  components: {
    Button,
    Input,
    Feed,
    TeacherDashboard,
    Statistics,
    LeaderDashboard,
    vSelect,
  },
};
</script>
<style scoped lang="scss">
.teacher {
  position: relative;
}

.student {
  & .justify {
    margin-bottom: 20px;
  }

  & .row {
    gap: 20px;
  }
}

section > .grid {
  max-width: 300px;
  width: 100%;
  text-align: left;
  gap: 10px;
  margin-bottom: 60px;
}

.feed {
  margin-top: 40px;
}
</style>
