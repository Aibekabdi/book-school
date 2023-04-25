<template>
  <Notification class="notification" v-if="onError" :message="onError" />
  <div class="admin" v-if="!user">
    <h2 class="heading-2">{{ t(currentLocalization, 'SIAD') }}</h2>
    <section>
      <component
        v-for="item in adminLogin"
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

    <p v-if="isErr" class="error-message">{{ t(currentLocalization, 'INCORRECT_LOGIN') }}</p>
  </div>
  <div class="admin max-w" v-if="user === 'admin'">
    <div class="sun"></div>
    <div class="robot"></div>
    <h2 class="heading-2">{{ t(currentLocalization, 'ADMIN_PANEL') }}</h2>
    <div class="grid school">
      <nav class="row justify">
        <div class="row">
          <Button
            type="Router"
            link="/create"
            :label="t(currentLocalization, 'ADD_BOOK')"
            icon="note_add"
            color="primary"
          />
          <Button
            type="Button"
            @click="toggleModal('AddQuiz')"
            :label="t(currentLocalization, 'ADD_TEST')"
            icon="quiz"
            color="second"
          />
          <Button
            type="Button"
            @click="toggleModal('AddQuestion')"
            :label="t(currentLocalization, 'ADD_QUESTION')"
            icon="help"
            color="success"
          />
          <Button
            type="Button"
            @click="toggleModal('EditData'), getAllInfo()"
            :label="t(currentLocalization, 'EDIT_DATA')"
            icon="edit_note"
            color="info"
          />
        </div>
        <div class="row">
          <Button
            type="Button"
            @click="toggleModal('Tariffing')"
            :label="t(currentLocalization, 'TARIFFING')"
            icon="payments"
            color="danger"
          />
        </div>
      </nav>
      <div class="dashboard"></div>
    </div>
    <h5 class="heading-2">{{ t(currentLocalization, 'ALL_BOOKS') }}</h5>
    <Feed />
    <div class="modal" v-if="open">
      <div>
        <div class="close">
          <Button type="Button" @click="toggleModal()" icon="close" :rounded="true" color="black" />
        </div>
        <div class="modal-content" v-if="modalType == 'AddQuiz'">
          <div>
            <h3 class="heading-2">{{ t(currentLocalization, 'CREATE_TEST') }}</h3>
            <p class="text-4">{{ t(currentLocalization, 'HCTFB') }}</p>
          </div>
          <CreateTest />
        </div>
        <div class="modal-content" v-if="modalType == 'AddQuestion'">
          <div>
            <h3 class="heading-2">{{ t(currentLocalization, 'CREATE_QUESTION') }}</h3>
            <p class="text-4">{{ t(currentLocalization, 'HCBQFT') }}</p>
          </div>
          <CreateCreativeTask />
        </div>
        <div class="modal-content" v-if="modalType == 'EditData'">
          <div>
            <h3 class="heading-2">{{ t(currentLocalization, 'ALL_USERS_TABLE') }}</h3>
            <p class="text-4">{{ t(currentLocalization, 'STCSS') }}</p>
          </div>
          <div class="grid pre-dash">
            <AdminDashboard :schools="schools" />
          </div>
        </div>
      </div>
    </div>
    <ModalUnavailable />
  </div>
</template>
<script setup>
import { ref, reactive, watch } from 'vue';
import ModalUnavailable from '@/components/modals/Modal-unavailable.vue';
import Input from '@/components/Input-component.vue';
import Button from '@/components/Button-component.vue';
import Feed from '@/components/Feed-component.vue';
import AdminDashboard from '@/components/AdminDashboard-component.vue';
import CreateTest from '@/components/CreateTest-component.vue';
import CreateCreativeTask from '@/components/CreateCreativeTask-component.vue';
import Notification from '@/components/layout/Notification-component.vue';
import { sendRequest } from '@/utils/utils';
import { user, userToken, onError, currentLocalization } from '@/App.vue';
import { t, tErr } from '@/utils/i18n.js';
</script>
<script>
const open = ref(false);
const modalType = ref('');
const isErr = ref(false);
const schools = ref({});

const adminLogin = reactive([
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
  {
    ref: '',
    label: 'Подтвердите пароль',
    name: 'confirm-password',
    type: 'password',
    placeholder: 'CONFIRM_PASSWORD',
    rules: [
      {
        id: 3,
        f: (val) => val === adminLogin[1].ref,
        msg: 'PDNM',
      },
    ],
    error: reactive([]),
  },
]);

watch(adminLogin, () => {
  for (const item of adminLogin) {
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
  name: 'PaymentView',
  components: {
    Button,
    Input,
    ModalUnavailable,
    Feed,
    AdminDashboard,
    CreateTest,
    CreateCreativeTask,
  },
  mounted() {
    console.log(user.value);
  },
  methods: {
    toggleModal(modal) {
      modalType.value = modal;
      open.value = !open.value;

      if (!open.value) {
        modalType.value = '';
      }
    },
    login() {
      isErr.value = false;
      const json = {
        name: adminLogin[0].ref,
        password: adminLogin[1].ref,
      };
      console.log(json);
      sendRequest('/auth/admin/sign-in', 'POST', json)
        .then((data) => {
          if (data) {
            user.value = 'admin';
            userToken.value = data.token;
            localStorage.setItem('admin', userToken.value);
          }
        })
        .catch((err) => {
          isErr.value = true;
        });
    },
    getAllInfo() {
      isErr.value = false;
      sendRequest('/api/school/all', 'GET', null, userToken.value)
        .then((data) => {
          if (data) {
            schools.value = data;
          }
          console.log(data);
        })
        .catch((err) => {
          onError.value = 'ITRD';
          console.log(err);

          setTimeout(() => {
            onError.value = null;
          }, 5000);
        });
    },
  },
};
</script>

<style scoped lang="scss">
section {
  text-align: center;
  height: fit-content;
}

.heading-2 {
  text-align: center;
  margin-bottom: 40px;
}
</style>
