<template>
  <Notification class="notification" v-if="onError" :message="onError" />
  <div v-if="this.schools.length > 0">
    <div class="row">
      <Tabs
        :names="tTabs(currentLocalization, tabs)"
        :selectedTab="selectedTab"
        @changeTab="changeTab"
      >
        <div class="dashboard" v-if="selectedTab === 'general'">
          <div class="table" v-for="(school, school_id) in this.schools" :key="school">
            <div class="grid">
              <div class="row">
                <span class="heading-2"
                  >{{ school_id + 1 }}. {{ t(currentLocalization, 'SCHOOL') }} -
                  {{ school.school.name }}
                </span>
                <Button
                  type="Button"
                  @click="showUpdateModal(this.schools[school_id])"
                  color="info"
                  icon="manage_accounts"
                />
              </div>
              <div class="grid">
                <span class="text-6"
                  >{{ t(currentLocalization, 'MAX_CLASSES') }}:
                  {{ school.school.class_count }}</span
                >
                <span class="text-6"
                  >{{ t(currentLocalization, 'CURR_CLASSES') }}:
                  {{ school.total_classes ? school.total_classes : 0 }}</span
                >
                <span class="text-6"
                  >{{ t(currentLocalization, 'AOS') }}:
                  {{ school.total_students ? school.total_students : 0 }}</span
                >
                <span class="text-6"
                  >{{ t(currentLocalization, 'AOT') }}: {{ school.teachers.length }}</span
                >
              </div>
            </div>
            <div class="row">
              <!-- TEACHER -->
              <div class="col">
                <span class="th-1">{{ t(currentLocalization, 'TEACHER') }}:</span>
                <div class="row">
                  <div class="col">
                    <span class="th-2"
                      >{{ t(currentLocalization, 'LASTNAME') }}{{ ' '
                      }}{{ t(currentLocalization, 'NAME') }}</span
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
                      >{{ t(currentLocalization, 'LASTNAME') }}{{ ' '
                      }}{{ t(currentLocalization, 'NAME') }}</span
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
                <div class="row">
                  <div class="col">
                    <span
                      >{{ teacher.teacher.id }}. {{ teacher.teacher.first_name }}
                      {{ teacher.teacher.second_name }}</span
                    >
                  </div>
                  <div class="col">{{ teacher.teacher.username }}</div>
                </div>
              </div>
              <div class="col row-card">
                <div class="row" v-for="classroom in teacher.classes" :key="classroom">
                  <div class="col">{{ classroom.class.grade }} {{ classroom.class.name }}</div>
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
        <div class="dashboard" v-else-if="selectedTab === 'classes'">
          <div class="table" v-for="(school, school_id) in this.schools" :key="school">
            <div class="grid">
              <div class="row">
                <span class="heading-2"
                  >{{ school_id + 1 }}. {{ t(currentLocalization, 'SCHOOL') }} -
                  {{ school.school.name }}</span
                >
                <Button
                  type="Button"
                  @click="showUpdateModal(this.schools[school_id])"
                  color="info"
                  icon="manage_accounts"
                />
              </div>
            </div>
            <div class="row">
              <!-- TEACHER -->
              <div class="col">
                <span class="th-1">{{ t(currentLocalization, 'TEACHER') }}:</span>
                <div class="row">
                  <div class="col">
                    <span class="th-2"
                      >{{ t(currentLocalization, 'LASTNAME') }}{{ ' '
                      }}{{ t(currentLocalization, 'NAME') }}</span
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
            <div class="row margin" v-for="teacher in school.teachers" :key="teacher">
              <div class="col">
                <div class="row">
                  <div class="col">
                    <span>{{ teacher.teacher.first_name }} {{ teacher.teacher.second_name }}</span>
                  </div>
                  <div class="col">
                    <span>{{ teacher.teacher.username }}</span>
                  </div>
                </div>
              </div>
              <div class="col">
                <div class="row" v-for="classroom in teacher.classes" :key="classroom">
                  <div class="col">{{ classroom.class.grade }} {{ classroom.class.name }}</div>
                  <div class="col">{{ classroom.students ? classroom.students.length : 0 }}</div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="dashboard" v-else-if="selectedTab === 'students'">
          <div class="table" v-for="(school, school_id) in this.schools" :key="school">
            <div class="grid">
              <div class="row">
                <span class="heading-2"
                  >{{ school_id + 1 }}. {{ t(currentLocalization, 'SCHOOL') }} -
                  {{ school.school.name }}</span
                >
                <Button
                  type="Button"
                  @click="showUpdateModal(this.schools[school_id])"
                  color="info"
                  icon="manage_accounts"
                />
              </div>
            </div>
            <div class="row">
              <!-- TEACHER -->
              <div class="col">
                <span class="th-1">{{ t(currentLocalization, 'TEACHER') }}:</span>
                <div class="row">
                  <div class="col">
                    <span class="th-2"
                      >{{ t(currentLocalization, 'LASTNAME') }}{{ ' '
                      }}{{ t(currentLocalization, 'NAME') }}</span
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
                      >{{ t(currentLocalization, 'LASTNAME') }}{{ ' '
                      }}{{ t(currentLocalization, 'NAME') }}</span
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
                <div class="row">
                  <div class="col">
                    <span
                      >{{ teacher.teacher.id }}. {{ teacher.teacher.first_name }}
                      {{ teacher.teacher.second_name }}</span
                    >
                  </div>
                  <div class="col">{{ teacher.teacher.username }}</div>
                </div>
              </div>
              <div class="col row-card">
                <div class="row" v-for="classroom in teacher.classes" :key="classroom">
                  <div class="col">{{ classroom.class.grade }} {{ classroom.class.name }}</div>
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
      </Tabs>
      <div class="inner-modal" v-if="open">
        <div>
          <div class="close">
            <Button
              type="Button"
              @click="open = !open"
              icon="close"
              :rounded="true"
              color="black"
            />
          </div>
          <div class="modal-content">
            <h4 class="heading-3">{{ t(currentLocalization, 'COD') }}</h4>
            <div class="row justify">
              <div class="grid">
                <span class="text-4">ID:</span>
              </div>
              <div class="grid">
                <span class="text-4">{{ t(currentLocalization, 'LOGIN') }}:</span>
              </div>
              <div class="grid">
                <span class="text-4">Пароль:</span>
              </div>
            </div>
            <div class="row justify">
              <div class="grid grid-text">{{ updatedUser.id }}</div>
              <div class="grid grid-text">{{ updatedUser.login }}</div>
              <div class="grid grid-text">{{ updatedUser.password }}</div>
            </div>
            <div class="row justify">
              <div class="grid">
                <span class="text-4">Классов:</span>
              </div>
              <div class="grid">
                <span class="text-4">Учеников:</span>
              </div>
              <div class="grid">
                <span class="text-4">Учителей:</span>
              </div>
            </div>
            <div class="row justify">
              <div class="grid grid-text">{{ updatedUser.classes }}</div>
              <div class="grid grid-text">{{ updatedUser.students }}</div>
              <div class="grid grid-text">{{ updatedUser.teachers }}</div>
            </div>
            <div class="grid create">
              <Input
                :label="t(currentLocalization, 'LOGIN')"
                name="login"
                type="text"
                :placeholder="t(currentLocalization, 'ENTER_SCHOOL_LOGIN')"
                v-model:value="updateUser.login"
                width="100%"
              />
              <Input
                :label="t(currentLocalization, 'CLASS_AMOUNT')"
                name="class_count"
                type="number"
                :placeholder="t(currentLocalization, 'ENTER_CLASSES')"
                v-model:value="updateUser.class_count"
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
                @click="updateSchool()"
                color="success"
                :label="t(currentLocalization, 'UD')"
                width="100%"
              />
            </div>
            <Button
              type="Button"
              @click="deleteSchool()"
              color="danger"
              label="Удалить школу"
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
</template>

<script setup>
import { ref, reactive, defineProps } from 'vue';
import Button from './Button-component.vue';
import Input from './Input-component.vue';
import Tabs from '@/components/Tab-component.vue';
import Notification from '@/components/layout/Notification-component.vue';
import { sendRequest } from '@/utils/utils';
import { userToken, onError, currentLocalization } from '@/App.vue';
import { t, tTabs } from '@/utils/i18n.js';

defineProps({
  schools: {
    type: Object,
    required: true,
  },
});
</script>
<script>
const isSuc = ref(false);
const isErr = ref(false);

const open = ref(false);

const tabs = [
  { name: 'general', label: 'GENERAL' },
  { name: 'classes', label: 'CLASSES' },
  { name: 'students', label: 'PUPIL' },
];

const updateUser = reactive({
  class_count: null,
  login: '',
  password: '',
});

const updatedUser = reactive({
  id: null,
  classes: '',
  students: null,
  teachers: null,
  login: '',
  password: '',
});

const selectedTab = ref('general');
const changeTab = (tabName) => {
  selectedTab.value = tabName;
};

export default {
  components: { Tabs, Button, Input, Notification },
  methods: {
    showUpdateModal(obj = null) {
      isSuc.value = false;
      isErr.value = false;
      open.value = !open.value;
      console.log(obj);

      updatedUser.id = obj.school.id;
      updatedUser.classes = obj.total_classes + ' / ' + obj.school.class_count;
      updatedUser.students = obj.total_students;
      updatedUser.teachers = obj.teachers.length;
      updatedUser.login = obj.school.name;
      updatedUser.password = obj.school.password;
    },
    updateSchool() {
      const json = {
        school_id: updatedUser.id,
        class_count: +updateUser.class_count,
        name: updateUser.login,
        password: updateUser.password,
      };
      console.log(json);
      sendRequest('/api/school/profile', 'PATCH', json, userToken.value)
        .then(() => {
          isSuc.value = true;
        })
        .catch((err) => {
          console.log(err);
          isErr.value = true;
        });
    },
    deleteSchool() {
      sendRequest(`/api/school/delete/${updatedUser.id}`, 'DELETE', null, userToken.value)
        .then(() => {
          open.value = !open.value;
        })
        .catch((err) => {
          onError.value = 'ITCU';
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
