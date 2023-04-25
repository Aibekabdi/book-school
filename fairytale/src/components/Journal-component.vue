<template>
  <Notification class="notification" v-if="onError" :message="onError" />
  <div v-if="usersList">
    <div class="grid">
      <div class="row justify">
        <h6 class="heading-3">{{ t(currentLocalization, 'LOS') }}:</h6>
        <div class="row">
          <div class="grid width">
            <span class="heading-3">{{ t(currentLocalization, 'SELECT_STUDENT') }}:</span>
            <v-select :options="usersOption" v-model="selectedUser"></v-select>
            <div class="btn-container">
              <Button
                type="Button"
                @click="openModal(selectedUser.split('.')[0])"
                icon="check"
                :label="t(currentLocalization, 'VIEW')"
                color="info"
                width="100%"
              />
            </div>
          </div>
        </div>
      </div>
      <div class="table" v-for="(classes, class_i) in usersList.classes" :key="classes">
        <div class="grid">
          <p class="text-3">
            {{ class_i + 1 }}. {{ classes.class.grade }} {{ classes.class.name }}
          </p>
          <span>{{ t(currentLocalization, 'LASTNAME') }} {{ t(currentLocalization, 'NAME') }}</span>
        </div>
        <div class="row justify" v-for="(user, i) in classes.students" :key="user">
          <div class="grid">
            <div class="row">
              <div class="grid">
                <span>{{ i + 1 }}) {{ user.second_name }} {{ user.first_name }}</span>
              </div>
              <div class="grid">
                <span>{{ user.username }}</span>
              </div>
            </div>
          </div>
          <div class="grid">
            <Button
              type="Button"
              @click="openModal(user.id, user.question_id)"
              icon="check_circle"
              :label="t(currentLocalization, 'CHECK')"
              color="success"
            />
            <div class="modal" v-if="open">
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
                <div class="pre-modal" v-if="creativeAnswers">
                  <div class="modal-content" v-for="(item, index) in creativeAnswers" :key="item">
                    <h2 class="heading-2 handle" v-if="item.question">{{ item.question }}</h2>
                    <h2 class="heading-2 handle error-message" v-else>
                      Ученик пока не ответил на вопрос
                    </h2>
                    <br />
                    <h3 class="heading-2 handle" v-if="!item.is_art">
                      {{ index + 1 }}. {{ item.Answer }}
                    </h3>
                    <div class="image" v-else-if="item.is_art">
                      <img :src="item.Answer" />
                    </div>
                    <div class="row">
                      <input
                        type="range"
                        name="range"
                        id="range"
                        min="0"
                        max="100"
                        v-model="points"
                      />
                      <span class="text-2">{{ points }}</span>
                    </div>
                    <h4 class="heading-3">{{ t(currentLocalization, 'WAFS') }}:</h4>
                    <textarea
                      v-model="textarea"
                      :placeholder="t(currentLocalization, 'SWH') + '...'"
                    ></textarea>
                    <div class="messages">
                      <span v-if="isSuc" class="success-message"
                        >{{ t(currentLocalization, 'AWSS') }}!</span
                      >
                      <span v-else-if="isErr" class="error-message">{{
                        t(currentLocalization, 'REG_ERR')
                      }}</span>
                    </div>
                    <Button
                      type="Button"
                      @click="submitComment(item.student_id, item.id)"
                      :label="t(currentLocalization, 'SEND')"
                      color="primary"
                      icon="send"
                      size="large"
                    />
                    <hr />
                  </div>
                </div>
                <div class="inside" v-else>
                  <div class="modal-content">
                    <span class="heading-1 material-symbols-outlined"> visibility_off </span>
                    <h3 class="heading-1">{{ t(currentLocalization, 'NTSH') }}</h3>
                    <p class="heading-2 error-message">{{ t(currentLocalization, 'NO_DATA') }}</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div v-else>
    <h3 class="heading-3 error-message">{{ t(currentLocalization, 'NO_DATA') }}</h3>
  </div>
</template>
<script setup>
import vSelect from 'vue-select';
import Tabs from '@/components/Tab-component.vue';
import Input from '@/components/Input-component.vue';
import Button from '@/components/Button-component.vue';
import Notification from '@/components/layout/Notification-component.vue';
import { ref } from 'vue';
import { sendRequest } from '@/utils/utils';
import { userToken, onError, currentLocalization } from '@/App.vue';
import { t } from '@/utils/i18n.js';
</script>
<script>
const selectedUser = ref('');
const usersList = ref([]);
const usersOption = ref([]);

const open = ref(false);

const isSuc = ref(false);
const isErr = ref(false);

const creativeAnswers = ref([]);

export default {
  components: { Tabs, Input, Button, vSelect, Notification },
  mounted() {
    this.getCreativePassed();
  },
  methods: {
    openModal(studentId) {
      console.log(studentId);
      open.value = !open.value;
      const id = document.URL.split('/content/')[1];
      sendRequest(
        `/api/creative/check/get/student/passes/${studentId}/${id}`,
        'GET',
        null,
        userToken.value,
      )
        .then((data) => {
          console.log(data);
          creativeAnswers.value = data;
        })
        .catch((err) => {
          onError.value = 'ITRD';
          console.log(err);

          setTimeout(() => {
            onError.value = null;
          }, 5000);
        });
    },
    getCreativePassed() {
      const id = document.URL.split('/content/')[1];
      sendRequest(`/api/creative/check/get/all/${id}`, 'GET', null, userToken.value)
        .then((data) => {
          let i = 0;
          usersList.value = data;
          data.classes.forEach((classroom) => {
            classroom.students.forEach((student) => {
              usersOption.value[i] =
                student.id +
                '. ' +
                student.second_name +
                ' ' +
                student.first_name +
                ` [${student.username}]`;
              i++;
            });
          });
          console.warn(usersList.value);
        })
        .catch((err) => {
          onError.value = 'ITRD';
          console.log(err);

          setTimeout(() => {
            onError.value = null;
          }, 5000);
        });
    },
    submitComment(student_id, answer_id) {
      const json = {
        student_id: student_id,
        answer_id: answer_id,
        comment: this.textarea,
        point: +this.points,
      };
      console.log(json);

      sendRequest('/api/creative/check/comment', 'POST', json, userToken.value)
        .then(() => {
          console.log('OK');
        })
        .catch((err) => {
          onError.value = 'ITSD';
          console.log(err);

          setTimeout(() => {
            onError.value = null;
          }, 5000);
        });
    },
  },
  data() {
    return {
      textarea: '',
      points: 0,
    };
  },
};
</script>
<style scoped lang="scss">
textarea {
  width: 100%;
  height: 500px;
  resize: vertical;
  border-radius: 12px;
  padding: 12px;
  border: 2px solid var(--primary-hover);
  font-family: 'Roboto Slab', serif;
}

hr {
  margin-bottom: 20px;
}

.btn-container {
  width: 100%;
  margin: 20px 0;
}

.width {
  min-width: 420px;
}

.handle {
  width: 80%;
  word-wrap: break-word;
}

.table {
  margin-bottom: 20px;
}

.pre-modal {
  width: 800px;
  margin: 0 auto;
}

.image {
  width: 800px;
  height: 600px;
  border: 2px solid var(--primary);
}

.inside {
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-self: center;

  & .modal-content {
    min-height: 0;
  }
}

input[type='range'] {
  height: 32px;
  -webkit-appearance: none;
  margin: 10px 0;
  width: 100%;
  background-color: transparent;
}
input[type='range']:focus {
  outline: none;
}
input[type='range']::-webkit-slider-runnable-track {
  width: 100%;
  height: 10px;
  cursor: pointer;
  animate: 0.2s;
  box-shadow: 0px 0px 0px #000000;
  background: #a5affb;
  border-radius: 50px;
  border: 0px solid #000000;
}
input[type='range']::-webkit-slider-thumb {
  box-shadow: 0px 0px 0px #a5affb;
  border: 2px solid #6979f8;
  height: 24px;
  width: 24px;
  border-radius: 50px;
  cursor: pointer;
  -webkit-appearance: none;
  margin-top: -8px;
}
input[type='range']:focus::-webkit-slider-runnable-track {
  background: #a5affb;
}
input[type='range']::-moz-range-track {
  width: 100%;
  height: 10px;
  cursor: pointer;
  animate: 0.2s;
  box-shadow: 0px 0px 0px #000000;
  background: #a5affb;
  border-radius: 50px;
  border: 0px solid #000000;
}
input[type='range']::-moz-range-thumb {
  box-shadow: 0px 0px 0px #a5affb;
  border: 2px solid #6979f8;
  height: 24px;
  width: 24px;
  border-radius: 50px;
  cursor: pointer;
}
input[type='range']::-ms-track {
  width: 100%;
  height: 10px;
  cursor: pointer;
  animate: 0.2s;
  background: transparent;
  border-color: transparent;
  color: transparent;
}
input[type='range']::-ms-fill-lower {
  background: #a5affb;
  border: 0px solid #000000;
  border-radius: 100px;
  box-shadow: 0px 0px 0px #000000;
}
input[type='range']::-ms-fill-upper {
  background: #a5affb;
  border: 0px solid #000000;
  border-radius: 100px;
  box-shadow: 0px 0px 0px #000000;
}
input[type='range']::-ms-thumb {
  margin-top: 1px;
  box-shadow: 0px 0px 0px #a5affb;
  border: 2px solid #6979f8;
  height: 24px;
  width: 24px;
  border-radius: 50px;
  cursor: pointer;
}
input[type='range']:focus::-ms-fill-lower {
  background: #a5affb;
}
input[type='range']:focus::-ms-fill-upper {
  background: #a5affb;
}
</style>
