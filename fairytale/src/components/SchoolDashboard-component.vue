<template>
  <Notification class="notification" v-if="onError" :message="onError" />
  <div class="grid pre-dash">
    <div v-if="this.school">
      <div class="row">
        <Tabs
          :names="tTabs(currentLocalization, tabs)"
          :selectedTab="selectedTab"
          @changeTab="changeTab"
        >
          <div class="dashboard" v-if="selectedTab === 'general'">
            <div class="table">
              <div class="grid">
                <span class="heading-2"
                  >{{ this.school.school.id }}. {{ t(currentLocalization, 'SCHOOL') }} -
                  {{ this.school.school.name }}</span
                >
                <div class="grid" v-if="this.school.teachers">
                  <span class="text-6"
                    >{{ t(currentLocalization, 'MAX_CLASSES') }}:
                    {{ this.school.school.class_count }}</span
                  >
                  <span class="text-6"
                    >{{ t(currentLocalization, 'CURR_CLASSES') }}:
                    {{ this.school.total_classes ? this.school.total_classes : 0 }}</span
                  >
                  <span class="text-6"
                    >{{ t(currentLocalization, 'AOS') }}:
                    {{ this.school.total_students ? this.school.total_students : 0 }}</span
                  >
                  <span class="text-6"
                    >{{ t(currentLocalization, 'AOT') }}:
                    {{ this.school.teachers ? this.school.teachers.length : 0 }}</span
                  >
                </div>
                <div class="row">
                  <!-- TEACHER -->
                  <div class="col">
                    <span class="th-1">{{ t(currentLocalization, 'TEACHERS') }}:</span>
                    <div class="row">
                      <div class="col">
                        <span class="th-2"
                          >{{ t(currentLocalization, 'LASTNAME') }}{{ ' ' }}
                          {{ t(currentLocalization, 'NAME') }}</span
                        >
                      </div>
                      <div class="col">
                        <span class="th-2">{{ t(currentLocalization, 'LOGIN') }}</span>
                      </div>
                    </div>
                  </div>
                  <!-- CLASS -->
                  <div class="col">
                    <span class="th-1">{{ t(currentLocalization, 'CLASSES') }}:</span>
                    <div class="row">
                      <div class="col">
                        <span class="th-2">{{ t(currentLocalization, 'NLC') }}</span>
                      </div>
                      <div class="col">
                        <span class="th-2">{{ t(currentLocalization, 'AOS') }}</span>
                      </div>
                    </div>
                  </div>
                  <!-- STUDENTS -->
                  <div class="col">
                    <span class="th-1">{{ t(currentLocalization, 'STUDENTS') }}:</span>
                    <div class="row">
                      <div class="col">
                        <span class="th-2"
                          >{{ t(currentLocalization, 'LASTNAME') }}{{ ' ' }}
                          {{ t(currentLocalization, 'NAME') }}</span
                        >
                      </div>
                      <div class="col">
                        <span class="th-2">{{ t(currentLocalization, 'LOGIN') }}</span>
                      </div>
                    </div>
                  </div>
                </div>
                <div
                  class="row margin"
                  v-for="(teacher, teacher_id) in this.school.teachers"
                  :key="teacher"
                >
                  <div class="col">
                    <div class="row">
                      <div class="col">
                        <span
                          >{{ teacher_id + 1 }}. {{ teacher.teacher.first_name }}
                          {{ teacher.teacher.second_name }}</span
                        >
                      </div>
                      <div class="col">
                        <div class="row username">
                          {{ teacher.teacher.username }}
                          <Button
                            type="Button"
                            @click="showUpdateModal(this.school.teachers[teacher_id], 'teacher')"
                            color="info"
                            icon="manage_accounts"
                          />
                        </div>
                      </div>
                    </div>
                  </div>
                  <div class="col row-card">
                    <div
                      class="row"
                      v-for="(classroom, class_id) in teacher.classes"
                      :key="classroom"
                    >
                      <div class="col col-row">
                        <div class="row">
                          <span>{{ classroom.class.grade }} {{ classroom.class.name }}</span>
                          <Button
                            type="Button"
                            @click="
                              showUpdateModal(
                                this.school.teachers[teacher_id].classes[class_id],
                                'classes',
                              )
                            "
                            color="info"
                            icon="settings"
                          />
                        </div>
                      </div>
                      <div class="col">
                        {{ classroom.students ? classroom.students.length : 0 }}
                      </div>
                      <div class="col row-card">
                        <div class="row" v-for="student in classroom.students" :key="student">
                          <div class="col">{{ student.first_name }} {{ student.second_name }}</div>
                          <div class="col">{{ student.username }}</div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="dashboard" v-else-if="selectedTab === 'classes'">
            <div class="table">
              <div class="grid">
                <span class="heading-2"
                  >{{ school.school.id }}. {{ t(currentLocalization, 'SCHOOL') }} -
                  {{ school.school.name }}</span
                >
              </div>
              <div class="row">
                <!-- TEACHER -->
                <div class="col">
                  <span class="th-1">{{ t(currentLocalization, 'TEACHERS') }}:</span>
                  <div class="row">
                    <div class="col">
                      <span class="th-2"
                        >{{ t(currentLocalization, 'LASTNAME') }}{{ ' ' }}
                        {{ t(currentLocalization, 'NAME') }}</span
                      >
                    </div>
                    <div class="col">
                      <span class="th-2">{{ t(currentLocalization, 'LOGIN') }}</span>
                    </div>
                  </div>
                </div>
                <!-- CLASS -->
                <div class="col">
                  <span class="th-1">{{ t(currentLocalization, 'CLASSES') }}:</span>
                  <div class="row">
                    <div class="col">
                      <span class="th-2">{{ t(currentLocalization, 'NLC') }}</span>
                    </div>
                    <div class="col">
                      <span class="th-2">{{ t(currentLocalization, 'AOS') }}</span>
                    </div>
                  </div>
                </div>
              </div>
              <div
                class="row margin"
                v-for="(teacher, teacher_id) in school.teachers"
                :key="teacher"
              >
                <div class="col">
                  <div class="row">
                    <div class="col">
                      <span
                        >{{ teacher.teacher.first_name }} {{ teacher.teacher.second_name }}</span
                      >
                    </div>
                    <div class="col">
                      <div class="row username">
                        {{ teacher.teacher.username }}
                        <Button
                          type="Button"
                          @click="showUpdateModal(this.school.teachers[teacher_id], 'teacher')"
                          color="info"
                          icon="manage_accounts"
                        />
                      </div>
                    </div>
                  </div>
                </div>
                <div class="col">
                  <div
                    class="row"
                    v-for="(classroom, class_id) in teacher.classes"
                    :key="classroom"
                  >
                    <div class="col col-row">
                      <div class="row">
                        <span>{{ classroom.class.grade }} {{ classroom.class.name }}</span>
                        <Button
                          type="Button"
                          @click="
                            showUpdateModal(
                              this.school.teachers[teacher_id].classes[class_id],
                              'classes',
                            )
                          "
                          color="info"
                          icon="settings"
                        />
                      </div>
                    </div>
                    <div class="col">{{ classroom.students ? classroom.students.length : 0 }}</div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="dashboard" v-else-if="selectedTab === 'students'">
            <div class="table">
              <div class="grid">
                <span class="heading-2"
                  >{{ school.school.id }}. {{ t(currentLocalization, 'SCHOOL') }} -
                  {{ school.school.name }}</span
                >
              </div>
              <div class="row">
                <!-- CLASS -->
                <div class="col">
                  <span class="th-1">{{ t(currentLocalization, 'CLASSES') }}:</span>
                  <div class="row">
                    <div class="col">
                      <span class="th-2">{{ t(currentLocalization, 'NLC') }}</span>
                    </div>
                    <div class="col">
                      <span class="th-2">{{ t(currentLocalization, 'AOS') }}</span>
                    </div>
                  </div>
                </div>
                <!-- STUDENTS -->
                <div class="col">
                  <span class="th-1">{{ t(currentLocalization, 'STUDENTS') }}:</span>
                  <div class="row">
                    <div class="col">
                      <span class="th-2"
                        >{{ t(currentLocalization, 'LASTNAME') }}{{ ' ' }}
                        {{ t(currentLocalization, 'NAME') }}</span
                      >
                    </div>
                    <div class="col">
                      <span class="th-2">{{ t(currentLocalization, 'LOGIN') }}</span>
                    </div>
                  </div>
                </div>
              </div>
              <div class="row margin" v-for="teacher in school.teachers" :key="teacher">
                <div class="col">
                  <div
                    class="row"
                    v-for="(classroom, class_id) in teacher.classes"
                    :key="classroom"
                  >
                    <div class="col">
                      <div class="row">
                        <div class="col col-row">
                          <div class="row">
                            <span>{{ classroom.class.grade }} {{ classroom.class.name }}</span>
                            <Button
                              type="Button"
                              @click="showUpdateModal(this.school.teachers[teacher_id], 'classes')"
                              color="info"
                              icon="settings"
                            />
                          </div>
                        </div>
                        <div class="col">
                          {{ classroom.students ? classroom.students.length : 0 }}
                        </div>
                      </div>
                    </div>
                    <div class="col">
                      <div class="row" v-for="student in classroom.students" :key="student">
                        <div class="col">{{ student.first_name }} {{ student.second_name }}</div>
                        <div class="col">{{ student.username }}</div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </Tabs>
        <div class="inner-modal" v-if="open">
          <div>
            <div class="close">
              <Button type="Button" @click="open = ''" icon="close" :rounded="true" color="black" />
            </div>
            <div class="modal-content" v-if="open === 'teacher'">
              <h4 class="heading-3">{{ t(currentLocalization, 'COD') }}</h4>
              <div class="row justify">
                <div class="grid">
                  <span class="text-4">ID:</span>
                </div>
                <div class="grid">
                  <span class="text-4"
                    >{{ t(currentLocalization, 'LASTNAME') }}
                    {{ t(currentLocalization, 'NAME') }}:</span
                  >
                </div>
                <div class="grid">
                  <span class="text-4">{{ t(currentLocalization, 'LOGIN') }}:</span>
                </div>
                <div class="grid">
                  <span class="text-4">{{ t(currentLocalization, 'PWD') }}:</span>
                </div>
              </div>
              <div class="row justify">
                <div class="grid grid-text">{{ updatedUser.id }}</div>
                <div class="grid grid-text">{{ updatedUser.name }}</div>
                <div class="grid grid-text">{{ updatedUser.login }}</div>
                <div class="grid grid-text">{{ updatedUser.password }}</div>
              </div>
              <div class="grid create">
                <Input
                  :label="t(currentLocalization, 'USER_NAME')"
                  name="first_name"
                  type="text"
                  :placeholder="t(currentLocalization, 'ENTER_USER_NAME')"
                  v-model:value="updateUser.first_name"
                  width="100%"
                />
                <Input
                  :label="t(currentLocalization, 'USER_LASTNAME')"
                  name="second_name"
                  type="text"
                  :placeholder="t(currentLocalization, 'ENTER_USER_LASTNAME')"
                  v-model:value="updateUser.second_name"
                  width="100%"
                />
                <Input
                  :label="t(currentLocalization, 'USER_LOGIN')"
                  name="login"
                  type="text"
                  :placeholder="t(currentLocalization, 'ENTER_USER_LOGIN')"
                  v-model:value="updateUser.login"
                  width="100%"
                />
                <Input
                  :label="t(currentLocalization, 'PWD')"
                  name="password"
                  type="password"
                  :placeholder="t(currentLocalization, 'ENTER_PWD')"
                  v-model:value="updateUser.password"
                  width="100%"
                />
                <Button
                  type="Button"
                  @click="updateTeacher()"
                  color="success"
                  :label="t(currentLocalization, 'UD')"
                  width="100%"
                />
              </div>
              <Button
                type="Button"
                @click="deleteTeacher()"
                color="danger"
                :label="t(currentLocalization, 'DELETE_TEACHER')"
                width="100%"
              />
              <div class="messages">
                <span v-if="isSuc" class="success-message"
                  >{{ t(currentLocalization, 'UDUS') }}!</span
                >
                <span v-else-if="isErr" class="error-message">{{
                  t(currentLocalization, 'REG_ERR')
                }}</span>
              </div>
            </div>
            <div class="modal-content" v-else-if="open === 'classes'">
              <h4 class="heading-3">{{ t(currentLocalization, 'CLASS_DELETION') }}</h4>
              <Button
                type="Button"
                @click="deleteClass()"
                color="danger"
                :label="t(currentLocalization, 'DELETE_CLASS')"
                width="400px"
              />
            </div>
          </div>
        </div>
      </div>
    </div>
    <div v-else>
      <span class="heading-3 error-message">{{ t(currentLocalization, 'NO_DATA') }}</span>
      <div class="dashboard not-found">
        <span class="material-symbols-outlined">visibility_off</span>
        <span>{{ t(currentLocalization, 'TTCTACORP') }}</span>
      </div>
    </div>
  </div>
</template>
<script setup>
import Button from './Button-component.vue';
import Input from './Input-component.vue';
import Tabs from '@/components/Tab-component.vue';
import Notification from '@/components/layout/Notification-component.vue';
import { ref, reactive, defineProps } from 'vue';
import { sendRequest } from '@/utils/utils';
import { userToken, onError, currentLocalization } from '@/App.vue';
import { t, tTabs } from '@/utils/i18n.js';

defineProps({
  school: {
    type: Object,
    required: true,
  },
});
</script>
<script>
const isSuc = ref(false);
const isErr = ref(false);

const open = ref('');

const tabs = [
  { name: 'general', label: 'GENERAL' },
  { name: 'classes', label: 'CLASSES' },
  { name: 'students', label: 'STUDENTS' },
];

const selectedTab = ref('general');
const changeTab = (tabName) => {
  selectedTab.value = tabName;
};

const updateUser = reactive({
  first_name: '',
  second_name: '',
  login: '',
  password: '',
});

const updatedUser = reactive({
  id: null,
  name: '',
  login: '',
  password: '',
});

export default {
  components: { Tabs, Button, Input, Notification },
  methods: {
    showUpdateModal(obj = null, str) {
      isSuc.value = false;
      isErr.value = false;

      if (open) {
        open.value = str;
      } else {
        open.value = '';
      }

      console.log(obj);

      if (str === 'teacher') {
        updatedUser.id = obj.teacher.id;
        updatedUser.name = obj.teacher.second_name + ' ' + obj.teacher.first_name;
        updatedUser.login = obj.teacher.username;
        updatedUser.password = obj.teacher.password;
      } else {
        updatedUser.id = obj.class.id;
      }

      console.log(updatedUser.id);
    },
    updateTeacher() {
      const json = {
        teacher_id: updatedUser.id,
        first_name: updateUser.first_name,
        second_name: updateUser.second_name,
        username: updateUser.login,
        password: updateUser.password,
      };
      console.log(json);
      sendRequest('/api/teacher/profile', 'PATCH', json, userToken.value)
        .then(() => {
          isSuc.value = true;
        })
        .catch((err) => {
          console.log(err);
          isErr.value = true;
        });
    },
    deleteTeacher() {
      sendRequest(`/api/teacher/delete/${updatedUser.id}`, 'DELETE', null, userToken.value)
        .then(() => {
          open.value = '';
        })
        .catch((err) => {
          onError.value = 'ITDS';
          console.log(err);

          setTimeout(() => {
            onError.value = null;
          }, 5000);
        });
    },
    deleteClass() {
      sendRequest(`/api/class/delete/${updatedUser.id}`, 'DELETE', null, userToken.value)
        .then(() => {
          open.value = '';
        })
        .catch((err) => {
          onError.value = 'ITDS';
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
.table > .grid > .grid {
  text-align: left;
  margin-bottom: 20px;
}
.username {
  justify-content: flex-start;
}
.col-row {
  margin-bottom: 10px;

  & > .row {
    justify-content: start;
  }
}

.inner-modal {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 600px;
  background-color: var(--black-hover);
  align-self: center;
  justify-self: center;
  padding: 60px 20px;
  border-radius: 24px;
  z-index: 50;

  & .row.justify > .grid {
    width: 100%;
    overflow: scroll;

    & > .text-4 {
      font-weight: 700;
    }
  }

  & .heading-3 {
    color: var(--white);
  }

  & .grid-text,
  & .text-4 {
    color: var(--white-hover);
  }

  & .create {
    border-radius: 12px;
    padding: 20px 10px;
    gap: 20px;
    background-color: var(--white-hover);

    & .heading-3 {
      color: var(--black);
    }
    & .text-4 {
      color: var(--black-hover);
    }
  }

  & .form-input {
    margin-bottom: 0;
  }

  & .v-select {
    background-color: var(--white-hover);
  }

  & .close {
    position: fixed;
    right: 20px;
    top: 20px;
    width: fit-content;
  }
  & .modal-content {
    padding: 20px;
  }
}
.row-card {
  width: 203%;
  margin-bottom: 10px;
}
</style>
