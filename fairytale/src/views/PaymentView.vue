<template>
  <div class="payment max-w">
    <div class="grid" v-if="!payment">
      <h2 class="heading-1">{{ t(currentLocalization, 'PAY_NOW') }}</h2>
      <div class="grid">
        <!-- CONTENT -->
      </div>
    </div>
    <div class="grid" v-else-if="payment">
      <h2 class="heading-2">
        {{ t(currentLocalization, 'REG_FOR') }}
        {{
          selectedTab === 'parent'
            ? t(currentLocalization, 'FOR_PARENT')
            : selectedTab === 'teacher'
            ? t(currentLocalization, 'FOR_TEACHER')
            : t(currentLocalization, 'FOR_SCHOOL')
        }}
      </h2>
      <div class="grid">
        <Tabs
          :names="tTabs(currentLocalization, tabs)"
          :selectedTab="selectedTab"
          @changeTab="changeTab"
          width="720px"
        >
          <div class="payment__container">
            <div v-if="selectedTab === 'parent'">
              <section>
                <component
                  v-for="item in registerForm.teacher"
                  :key="item"
                  :is="Input"
                  :label="item.label"
                  :name="item.name"
                  :type="item.type"
                  v-model:value="item.ref"
                  :placeholder="t(currentLocalization, item.placeholder)"
                  :error="tErr(currentLocalization, item.error)"
                />
                <Button
                  type="Button"
                  @click="registerTeacher"
                  :label="t(currentLocalization, 'SIGN_UP')"
                  icon="person_add"
                  color="success"
                  width="100%"
                />
              </section>
            </div>
            <div v-if="selectedTab === 'teacher'">
              <section>
                <component
                  v-for="item in registerForm.teacher"
                  :key="item"
                  :is="Input"
                  :label="item.label"
                  :name="item.name"
                  :type="item.type"
                  v-model:value="item.ref"
                  :placeholder="t(currentLocalization, item.placeholder)"
                  :error="tErr(currentLocalization, item.error)"
                />
                <Button
                  type="Button"
                  @click="registerTeacher"
                  :label="t(currentLocalization, 'SIGN_UP')"
                  icon="person_add"
                  color="success"
                  width="100%"
                />
              </section>
            </div>
            <div v-if="selectedTab === 'school'">
              <section>
                <component
                  v-for="item in registerForm.school"
                  :key="item"
                  :is="Input"
                  :label="item.label"
                  :name="item.name"
                  :type="item.type"
                  v-model:value="item.ref"
                  :placeholder="t(currentLocalization, item.placeholder)"
                  :error="tErr(currentLocalization, item.error)"
                />
                <Button
                  type="Button"
                  @click="registerSchool"
                  :label="t(currentLocalization, 'SIGN_UP')"
                  icon="person_add"
                  color="success"
                  width="100%"
                />
              </section>
            </div>
            <div class="payment__img"><img src="../assets/image 22.png" alt="" /></div>
          </div>
        </Tabs>
        <span v-if="isSuc" class="success-message">{{ t(currentLocalization, 'REG_SUC') }}</span>
        <span v-if="isErr" class="error-message">{{ t(currentLocalization, 'REG_ERR') }}</span>
      </div>
    </div>
  </div>
</template>
<script setup>
import { ref, reactive, watch } from 'vue';
import Input from '@/components/Input-component.vue';
import Button from '@/components/Button-component.vue';
import Tabs from '@/components/Tab-component.vue';
import { sendRequest } from '@/utils/utils';
import { currentLocalization } from '@/App.vue';
import { t, tErr, tTabs } from '@/utils/i18n.js';
</script>
<script>
const tabs = [
  { name: 'parent', label: 'PARENT' },
  { name: 'teacher', label: 'TEACHER' },
  { name: 'school', label: 'SCHOOL' },
];
const selectedTab = ref('teacher');
const changeTab = (tabName) => {
  selectedTab.value = tabName;
};

const payment = ref(true);
const isSuc = ref(false);
const isErr = ref(false);

const registerForm = reactive({
  teacher: [
    {
      ref: '',
      label: 'Имя учителя',
      name: 'first-name',
      type: 'text',
      placeholder: 'ENTER_TEACHER_NAME',
      rules: [
        {
          id: 1,
          f: (val) => val.length > 0,
          msg: 'NF',
        },
      ],
      error: reactive([]),
    },
    {
      ref: '',
      label: 'Фамилия учителя',
      name: 'second-name',
      type: 'text',
      placeholder: 'ENTER_TEACHER_LASTNAME',
      rules: [
        {
          id: 2,
          f: (val) => val.length > 0,
          msg: 'NF',
        },
      ],
      error: reactive([]),
    },
    {
      ref: '',
      label: 'Логин',
      name: 'login',
      type: 'text',
      placeholder: 'ENTER_TEACHER_LOGIN',
      rules: [
        {
          id: 3,
          f: (val) => val.length >= 3 || val.length == 0,
          msg: 'LOGIN_LENGTH',
        },
      ],
      error: reactive([]),
    },
    {
      ref: '',
      label: 'Пароль',
      name: 'password',
      type: 'password',
      placeholder: 'ENTER_TEACHER_PASSWORD',
      rules: [
        {
          id: 4,
          f: (val) => val.length >= 3 || val.length == 0,
          msg: 'PWD_LENGTH',
        },
      ],
      error: reactive([]),
    },
    {
      ref: '',
      label: 'Подтвердите пароль',
      name: 'confirm-password',
      type: 'password',
      placeholder: 'CONFIRM_PASSWORD',
      rules: [
        {
          id: 5,
          f: (val) => val === registerForm.teacher[3].ref,
          msg: 'PDNM',
        },
      ],
      error: reactive([]),
    },
  ],
  school: [
    {
      ref: '',
      label: 'Классы',
      name: 'class',
      type: 'number',
      placeholder: 'ENTER_CLASSES',
      rules: [
        {
          id: 6,
          f: (val) => val.match(/^-?\d+$/),
          msg: 'CONTAIN_NUMBER',
        },
        {
          id: 7,
          f: (val) => val > 0,
          msg: 'MORE_THAN_0',
        },
      ],
      error: reactive([]),
    },
    {
      ref: '',
      label: 'Логин',
      name: 'login',
      type: 'text',
      placeholder: 'ENTER_SCHOOL_LOGIN',
      rules: [
        {
          id: 8,
          f: (val) => val.length >= 3 || val.length == 0,
          msg: 'LOGIN_LENGTH',
        },
      ],
      error: reactive([]),
    },
    {
      ref: '',
      label: 'Пароль',
      name: 'password',
      type: 'password',
      placeholder: 'ENTER_SCHOOL_PASSWORD',
      rules: [
        {
          id: 9,
          f: (val) => val.length >= 3 || val.length == 0,
          msg: 'PWD_LENGTH',
        },
      ],
      error: reactive([]),
    },
    {
      ref: '',
      label: 'Подтвердите пароль',
      name: 'confirm-password',
      type: 'password',
      placeholder: 'CONFIRM_PASSWORD',
      rules: [
        {
          id: 10,
          f: (val) => val === registerForm.school[2].ref,
          msg: 'PWD_MATCH',
        },
      ],
      error: reactive([]),
    },
  ],
});

watch(registerForm, () => {
  for (const key in registerForm) {
    for (const item of registerForm[key]) {
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
});

export default {
  name: 'PaymentView',
  components: { Input, Button },
  methods: {
    registerTeacher() {
      isErr.value = false;
      const json = {
        first_name: registerForm.teacher[0].ref,
        second_name: registerForm.teacher[1].ref,
        username: registerForm.teacher[2].ref,
        password: registerForm.teacher[3].ref,
      };

      console.log(json);

      sendRequest('/auth/teacher/sign-up', 'POST', json)
        .then(() => {
          // redirect to login
          this.$router.push({ path: '/login' });
          isSuc.value = true;
        })
        .catch((err) => {
          isErr.value = true;
        });
    },

    registerSchool() {
      isErr.value = false;
      const json = {
        class_count: +registerForm.school[0].ref,
        name: registerForm.school[1].ref,
        password: registerForm.school[2].ref,
      };

      console.log(json);

      sendRequest('/auth/school/sign-up', 'POST', json)
        .then(() => {
          // redirect to landing
          this.$router.push({ path: '/login' });
          isSuc.value = true;
        })
        .catch((err) => {
          isErr.value = true;
        });
    },
  },
  data: () => ({
    //
  }),
};
</script>
<style scoped lang="scss">
.payment {
  text-align: center;
  display: flex;
  justify-content: center;
  & > .grid {
    width: 80%;
    min-width: 300px;
    margin-bottom: 40px;
  }

  & > .grid > .grid {
    align-items: center;
  }
}

.error-message {
  margin-top: 20px;
}

.payment__container {
  display: flex;
  align-items: center;
  background: var(--light-blue);
  border-radius: 10px;
  padding: 55px;
  justify-content: space-between;
}
.payment__img {
  width: 221px;
  height: 215px;
  margin-left: 50px;
  flex-shrink: 0;
}
.payment {
  // padding: 20px;
  // color: var(--white-hover);
  // background-color: var(--black-hover);
  // border-radius: 12px;
  text-align: center;

  & .heading-1 {
    color: var(--white);
  }
}

@media only screen and (max-width: 1200px) {
  .payment__img {
    margin-top: 60px;
  }
  .payment__container {
    display: block;
    padding: 35px;
    width: fit-content;
    margin: 0 auto;
  }
  .btn {
    padding: 10px 30px;
    height: 35px;
    font-size: 14px;
  }
}
</style>
