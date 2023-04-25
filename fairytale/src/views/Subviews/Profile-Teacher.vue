<template>
  <Notification class="notification" v-if="onError" :message="onError" />
  <div class="grid teacher">
    <h3 class="heading-3">Teacher View</h3>
    <div class="grid">
      <nav class="row justify">
        <div class="row">
          <Button
            type="Router"
            link="/home"
            :label="t(currentLocalization, 'HOME')"
            icon="home"
            color="primary"
          />
        </div>
      </nav>
      <div class="row">
        <section>
          <h5 class="heading-3">
            <span v-if="userData.firstName"
              >{{ t(currentLocalization, 'NAME') }}: {{ userData.firstName }}</span
            >
            <span v-else>{{ t(currentLocalization, 'LOADING') }}...</span>
          </h5>
          <h5 class="heading-3">
            <span v-if="userData.secondName"
              >{{ t(currentLocalization, 'LASTNAME') }}: {{ userData.secondName }}</span
            >
            <span v-else>{{ t(currentLocalization, 'LOADING') }}...</span>
          </h5>
          <h5 class="heading-3">
            <span v-if="userData.login"
              >{{ t(currentLocalization, 'LOGIN') }}: {{ userData.login }}</span
            >
            <span v-else>{{ t(currentLocalization, 'LOADING') }}...</span>
          </h5>
          <!-- <h5 class="heading-3">
            <span v-if="userData.password"
              >{{ t(currentLocalization, 'PWD') }}: {{ userData.password }}</span
            >
            
            <span v-else>{{ t(currentLocalization, 'LOADING') }}...</span>
          </h5> -->
        </section>
        <section>
          <component
            v-for="item in userForm"
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
            @click="changeTeacher"
            :label="t(currentLocalization, 'UD')"
            icon="edit"
            color="success"
            width="100%"
          />
        </section>
      </div>
    </div>
    <span v-if="isSuc" class="success-message">{{ t(currentLocalization, 'CHANGE_SUC') }}</span>
    <span v-else-if="isErr" class="error-message">{{ t(currentLocalization, 'REG_ERR') }}</span>
  </div>
</template>
<script setup>
import { ref, reactive, watch } from 'vue';
import Input from '@/components/Input-component.vue';
import Button from '@/components/Button-component.vue';
import Notification from '@/components/layout/Notification-component.vue';
import { sendRequest } from '@/utils/utils';
import { userToken, onError, currentLocalization } from '@/App.vue';
import { t, tErr } from '@/utils/i18n.js';
</script>
<script>
const userData = reactive({
  id: null,
  firstName: '',
  secondName: '',
  login: '',
  password: '',
  class: '',
});

const isSuc = ref(false);
const isErr = ref(false);

const userForm = reactive([
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
    error: [],
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
        f: (val) => val === userForm[3].ref,
        msg: 'PDNM',
      },
    ],
    error: reactive([]),
  },
]);

watch(userForm, () => {
  for (const item of userForm) {
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
});

export default {
  name: 'ProfileTeacher',
  components: { Input, Button, Notification },
  methods: {
    getUserData() {
      sendRequest('/api/teacher/profile', 'GET', null, userToken.value)
        .then((data) => {
          if (data) {
            userData.id = data.id;
            userData.firstName = data.first_name;
            userData.secondName = data.second_name;
            userData.login = data.username;
            userData.password = data.password;
          }
        })
        .catch((err) => {
          onError.value = 'ITRD';
          console.log(err);

          setTimeout(() => {
            onError.value = null;
          }, 5000);
        });
    },

    changeTeacher() {
      if (
        userForm[0].ref === '' &&
        userForm[1].ref === '' &&
        userForm[2].ref === '' &&
        userForm[3].ref === ''
      ) {
        isErr.value = true;
        return;
      }
      if (userForm[2].ref === '') {
        userForm[2].ref = userData.login;
      }
      if (userForm[3].ref === '') {
        userForm[3].ref = userData.password;
        userForm[4].ref = userForm[3].ref;
      } else {
        if (userForm[4].ref === '') {
          isErr.value = true;
          return;
        }
      }
      if (userForm[0].ref === '') {
        userForm[0].ref = userData.firstName;
      }
      if (userForm[1].ref === '') {
        userForm[1].ref = userData.secondName;
      }
      const json = {
        teacher_id: userData.id,
        first_name: userForm[0].ref,
        second_name: userForm[1].ref,
        username: userForm[2].ref,
        password: userForm[3].ref,
      };
      sendRequest('/api/teacher/profile', 'PATCH', json, userToken.value)
        .then(() => {
          isSuc.value = true;
          this.getUserData();
        })
        .catch((err) => {
          console.log(err);
          isErr.value = true;
        });
    },
  },
  mounted() {
    this.getUserData();
    userToken.value = localStorage.getItem('admin') || localStorage.getItem('userToken') || '';
    console.log(userToken.value);
  },
};
</script>
<style scoped lang="scss"></style>
