<template>
  <div class="login">
    <!-- <h2 class="heading-1">{{ t(currentLocalization, 'GET_IN') }}</h2> -->
    <h3 class="text-4">
      {{ t(currentLocalization, 'SIGN_AS') }}
      <span>{{
        selectedUser == 'student'
          ? t(currentLocalization, 'STUDENT')
          : selectedUser == 'teacher'
          ? t(currentLocalization, 'TEACHER')
          : selectedUser == 'school'
          ? t(currentLocalization, 'SCHOOL')
          : ''
      }}</span>
    </h3>
    <div class="btn-container">
      <Button
        type="Button"
        @click="selectUser('student')"
        :label="t(currentLocalization, 'STUDENT')"
        color="success"
      />
      <Button
        type="Button"
        @click="selectUser('teacher')"
        :label="t(currentLocalization, 'TEACHER')"
        color="success"
      />
      <Button
        type="Button"
        @click="selectUser('school')"
        :label="t(currentLocalization, 'SCHOOL')"
        color="success"
      />
    </div>
    <div class="login__container">
      <section>
        <component
          v-for="item in userLogin"
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
          @click="login"
          :label="t(currentLocalization, 'SIGN_IN')"
          icon="login"
          color="success"
          width="100%"
        />
      </section>
      <div class="login__img"><img src="../assets/image 22.png" alt="" /></div>
    </div>
    <p v-if="isErr" class="error-message">{{ t(currentLocalization, 'WRONG_LOGIN') }}</p>
  </div>
</template>

<script setup>
import { ref, reactive, watch, nextTick } from 'vue';
import Input from '@/components/Input-component.vue';
import Button from '@/components/Button-component.vue';
import { sendRequest } from '@/utils/utils';
import { user, currentLocalization } from '@/App.vue';
import { t, tErr } from '@/utils/i18n.js';
</script>
<script>
const selectedUser = ref(localStorage.getItem('userType') || 'student');
const userToken = ref(localStorage.getItem('userToken') || '');

const isErr = ref(false);

const userLogin = reactive([
  {
    ref: '',
    label: 'Логин',
    name: 'login',
    type: 'text',
    placeholder: 'ENTER_LOGIN',
    rules: [
      {
        id: 1,
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
    placeholder: 'ENTER_PWD',
    rules: [
      {
        id: 2,
        f: (val) => val.length >= 3 || val.length == 0,
        msg: 'PWD_LENGTH',
      },
    ],
    error: reactive([]),
  },
]);

watch(userLogin, () => {
  for (const item of userLogin) {
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
  name: 'LoginView',
  components: { Input, Button },
  methods: {
    selectUser(user) {
      selectedUser.value = user;
    },
    async login() {
      isErr.value = false;
      const json = {
        name: userLogin[0].ref,
        password: userLogin[1].ref,
        who: selectedUser.value,
      };
      try {
        const data = await sendRequest('/auth/sign-in', 'POST', json);
        userToken.value = data.token;
        localStorage.setItem('userToken', userToken.value);
        localStorage.setItem('userType', selectedUser.value);
        user.value = selectedUser.value;
        await nextTick();
        this.$router.push({ path: '/home' });
      } catch (err) {
        isErr.value = true;
      }
    },
  },
};
</script>

<style scoped lang="scss">
.login__container {
  display: flex;
  align-items: center;
  background: var(--light-blue);
  border-radius: 10px;
  padding: 55px;
}
.login__img {
  width: 221px;
  height: 215px;
  margin-left: 50px;
  flex-shrink: 0;
}
.login {
  // padding: 20px;
  // color: var(--white-hover);
  // background-color: var(--black-hover);
  // border-radius: 12px;
  text-align: center;

  & .heading-1 {
    color: var(--white);
  }
}

.btn-container {
  width: 100%;
  justify-content: space-evenly;
  margin-bottom: 30px;
}

.text-4 {
  & > span {
    border-bottom: 2px solid var(--success);
  }
}

section {
  text-align: center;
  height: fit-content;
}

@media only screen and (max-width: 1200px) {
  .login__img {
    display: none;
  }
  .login__container {
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
