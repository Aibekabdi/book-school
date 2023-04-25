<template>
  <Notification class="notification" v-if="onError" :message="onError" />
  <div class="grid school">
    <h3 class="heading-3">School View</h3>
    <div class="grid">
      <nav class="row justify">
        <div class="row">
          <Button
            type="Button"
            @click="toggleModal('CreateTeacher')"
            label="Создать учителя"
            icon="add_circle"
            color="primary"
          />
          <Button
            type="Button"
            @click="toggleModal('CreateClass'), getAllTeachers()"
            label="Создать класс"
            icon="add"
            color="second"
          />
          <Button
            type="Button"
            @click="toggleModal('ConfigureClasses'), getClassForSchool()"
            label="Управление классами"
            icon="school"
            color="info"
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
    </div>
    <div class="modal" v-if="open">
      <div>
        <div class="close">
          <Button type="Button" @click="toggleModal()" icon="close" :rounded="true" color="black" />
        </div>
        <div class="modal-content" v-if="modalType == 'CreateTeacher'">
          <div>
            <h3 class="heading-2">Создать учителя</h3>
            <p class="text-4">Введите ФИО, логин и пароль учителя</p>
            <span v-if="isSuc" class="success-message">Учитель был успешно создан</span>
            <span v-else-if="isErr" class="error-message">Ошибка. Попробуйте снова</span>
          </div>
          <section>
            <component
              v-for="item in adminForm.teacher"
              :key="item"
              :is="Input"
              :label="item.label"
              :name="item.name"
              :type="item.type"
              v-model:value="item.ref"
              :placeholder="item.placeholder"
              :error="item.error"
            />
            <Button
              type="Button"
              @click="createTeacher"
              label="Создать"
              icon="add"
              color="success"
              width="100%"
            />
          </section>
          <p class="text-5">
            Здесь вы можете создать учителя, <a href="#">посмотрите гайд как это работает</a>
          </p>
        </div>
        <div class="modal-content" v-if="modalType == 'CreateClass'">
          <div>
            <h3 class="heading-2">Создать класс</h3>
            <p class="text-4">Введите цифру, букву и учителя</p>
            <span v-if="isSuc" class="success-message">Класс был успешно создан</span>
            <span v-else-if="isErr" class="error-message">Ошибка. Попробуйте снова</span>
          </div>
          <section>
            <div class="grid">
              <span class="heading-3">Выберите класс:</span>
              <v-select :options="classes" v-model="selectedOptionField2"></v-select>
            </div>
            <component
              v-for="item in adminForm.class"
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
              <span class="heading-3">Выберите учителя:</span>
              <v-select :options="allTeachers" v-model="selectedOptionField"></v-select>
            </div>
            <Button
              type="Button"
              @click="createClass"
              label="Создать"
              icon="add"
              color="success"
              width="100%"
            />
          </section>
          <p class="text-5">
            Здесь вы можете создать класс, <a href="#">посмотрите гайд как это работает</a>
          </p>
        </div>
        <div class="modal-content" v-if="modalType == 'ConfigureClasses'">
          <div>
            <h3 class="heading-2">Управление классами</h3>
            <p class="text-4">Просмотр учителей, классов и учеников. Статистика по школе</p>
          </div>
          <div class="grid pre-dash">
            <SchoolDashboard :school="school" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup>
import Button from '@/components/Button-component.vue';
import Input from '@/components/Input-component.vue';
import SchoolDashboard from '@/components/SchoolDashboard-component.vue';
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
const selectedOptionField2 = ref(null);
const school = ref('');
const allTeachers = ref([]);
const classes = ['2 год', '3 год', '4 год', '5 год', '1 класс', '2 класс', '3 класс', '4 класс'];

const adminForm = reactive({
  teacher: [
    {
      ref: '',
      label: 'Имя учителя',
      name: 'first-name',
      type: 'text',
      placeholder: 'Введите имя учителя',
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
      label: 'Фамилия учителя',
      name: 'second-name',
      type: 'text',
      placeholder: 'Введите фамилию учителя',
      rules: [
        {
          id: 2,
          f: (val) => val.length > 0,
          msg: 'Это поле обязательно',
        },
      ],
      error: reactive([]),
    },
    {
      ref: '',
      label: 'Логин',
      name: 'login',
      type: 'text',
      placeholder: 'Введите логин учителя',
      rules: [
        {
          id: 3,
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
          id: 4,
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
          id: 5,
          f: (val) => val === adminForm.teacher[3].ref,
          msg: 'Пароли не совпадают',
        },
      ],
      error: reactive([]),
    },
  ],
  class: [
    {
      ref: '',
      label: 'Буква класса',
      name: 'letter',
      type: 'text',
      placeholder: 'Введите букву класса',
      rules: [
        {
          id: 3,
          f: (val) => val.length === 1,
          msg: 'Поле должно содержать 1 символ',
        },
        {
          id: 4,
          f: (val) => val.match(/[a-z]/i),
          msg: 'Поле должно состоять из буквы',
        },
      ],
      error: [],
    },
  ],
});
export default {
  mounted() {
    user.value = localStorage.getItem('admin') ? 'admin' : localStorage.getItem('userType') || '';
    userToken.value = localStorage.getItem('admin') || localStorage.getItem('userToken') || '';
    selectedOptionField.value = '';

    watch(
      adminForm,
      () => {
        for (const key in adminForm) {
          for (const item of adminForm[key]) {
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

    // school view
    createTeacher() {
      isSuc.value = false;
      isErr.value = false;
      const json = {
        first_name: adminForm.teacher[0].ref,
        second_name: adminForm.teacher[1].ref,
        username: adminForm.teacher[2].ref,
        password: adminForm.teacher[3].ref,
      };

      console.log(json);

      sendRequest('/api/teacher/create', 'POST', json, userToken.value)
        .then(() => {
          isSuc.value = true;
        })
        .catch((err) => {
          isErr.value = true;
        });
    },
    // school view
    createClass() {
      isSuc.value = false;
      isErr.value = false;
      const json = {
        grade: selectedOptionField2.value,
        name: adminForm.class[0].ref,
        teacher_id: +selectedOptionField.value.split('.')[0],
      };
      sendRequest('/api/class/create', 'POST', json, userToken.value)
        .then(() => {
          isSuc.value = true;
        })
        .catch((err) => {
          isErr.value = true;
        });
    },
    // school view
    getClassForSchool() {
      isSuc.value = false;
      isErr.value = false;

      sendRequest('/api/teacher/all', 'GET', null, userToken.value)
        .then((data) => {
          if (data) {
            school.value = data;
            console.log(school.value);
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
    getAllTeachers() {
      sendRequest('/api/teacher/all', 'GET', null, userToken.value)
        .then((data) => {
          if (data) {
            data.teachers.forEach((el, i) => {
              allTeachers.value[i] =
                +el.teacher.id +
                '. ' +
                el.teacher.first_name +
                ' ' +
                el.teacher.second_name +
                ` [${el.teacher.username}]`;
            });
          }
        })
        .catch((err) => {
          onError.value = 'Невозможно получить данные учителей, пожалуйста повторите позднее';
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
    SchoolDashboard,
    LeaderDashboard,
    vSelect,
  },
};
</script>
<style scoped lang="scss">
.school {
  position: relative;
}

section > .grid {
  max-width: 300px;
  width: 100%;
  text-align: left;
  gap: 10px;
  margin-bottom: 60px;
}
</style>
