<template>
  <Notification class="notification" v-if="onError" :message="onError" />
  <div class="dropdown-container" v-if="user === 'teacher'">
    <v-select class="dropdown" :options="students" v-model="selectedStudent"></v-select>
    <Button
      type="Button"
      @click="showTest"
      :label="t(currentLocalization, 'CHECK_TEST')"
      color="success"
      icon="priority"
    />
  </div>
  <div v-if="this.testData.id">
    <img
      v-if="open_url"
      class="full-image"
      @click="openImage()"
      :src="open_url"
      :title="t(currentLocalization, 'PTD')"
    />
    <div class="grid grid-main">
      <span class="heading-3 error-message" v-if="completed">{{
        t(currentLocalization, 'LEAVEPAGE')
      }}</span>
      <span class="heading-3" v-if="completed"
        >{{ t(currentLocalization, 'YCAO') }} - {{ points === 0 ? '0' : points / 10 }}
        {{ t(currentLocalization, 'ONQUESTIONS') }}</span
      >
      <span class="heading-3" v-if="!completed"
        >{{ t(currentLocalization, 'AORQ') }} - {{ points }}</span
      >
      <div class="row row-main">
        <div class="grid">
          <div class="row row-secondary">
            <audio :src="question.audio" :id="question_i + '-audio'" style="display: none"></audio>
            <Button
              type="Button"
              @click="playAudio(`${question_i}-audio`)"
              color="info"
              icon="play_arrow"
              size="small"
              width="40px"
            />
            <span class="text-4 mb-2">{{ question.question }}</span>
          </div>
          <!-- <audio controls>
            <source :src="question.audio" type="audio/mpeg" />
          </audio> -->
          <img
            v-if="question.image"
            @click="openImage(question.image)"
            class="image"
            :src="question.image"
            :title="t(currentLocalization, 'PTI')"
          />
        </div>
      </div>

      <div class="row row-secondary" v-for="(answer, answer_i) in question.answers" :key="answer">
        <input
          v-on:change="chooseAnswer(question_i + 1, question.id, answer.id, answer_i)"
          type="radio"
          :name="question.id + 'question'"
          :id="question.id + 'question'"
          :class="'answer-chose'"
        />
        <div class="grid">
          <div class="row justify">
            <div class="row">
              <audio
                :src="answer.audio"
                :id="question_i + '-' + answer_i + '-audio'"
                style="display: none"
              ></audio>
              <Button
                type="Button"
                @click="playAudio(`${question_i}-${answer_i}-audio`)"
                color="info"
                icon="play_arrow"
                size="small"
                width="40px"
              />
              <span class="text-4">{{ answer.answer }}</span>
            </div>
            <!-- <audio controls>
              <source :src="answer.audio" type="audio/mpeg" />
            </audio> -->
            <div class="row">
              <span
                class="text-4"
                title="Правильный ответ"
                v-if="answer.correct && user === 'teacher'"
              >
                <span class="material-symbols-outlined success-message"> done </span>
              </span>
              <span
                class="text-4"
                title="Выбранный учеником ответ"
                v-if="answer.is_students_answer && user === 'teacher'"
              >
                <span class="material-symbols-outlined error-message">
                  radio_button_unchecked
                </span></span
              >
            </div>
          </div>
          <img
            v-if="answer.image"
            @click="openImage(answer.image)"
            class="image"
            :src="answer.image"
            :title="t(currentLocalization, 'PTI')"
          />
        </div>
        <div class="grid" v-if="completed">
          <span
            v-if="answer.correct && complete_test.correct_answers[0]"
            class="material-symbols-outlined success-message"
          >
            check_small
          </span>
          <!-- <span title="" v-else class="material-symbols-outlined error-message"> close </span> -->
        </div>
      </div>
    </div>
    <div class="grid grid-main">
      <div class="row">
        <div class="col" v-for="(q, q_index) in testData.questions" :key="q">
          <Button
            type="Button"
            color="success"
            width="40px"
            :label="(q_index + 1).toString()"
            @click="changeQuestion(q, q_index)"
          />
        </div>
      </div>
    </div>
    <div class="row">
      <Button
        v-if="user === 'student'"
        type="Button"
        @click="passTest"
        :label="t(currentLocalization, 'END_TEST')"
        color="success"
        icon="check_small"
        :disabled="completed"
      />
      <Button
        v-if="completed"
        type="Button"
        @click="rePassTest"
        :label="t(currentLocalization, 'RETAKETEST')"
        color="success"
      >
        <span class="material-symbols-outlined"> restart_alt </span>
      </Button>
    </div>
  </div>
  <div v-else>
    <h4 class="heading-3">{{ t(currentLocalization, 'NO_DATA') }}</h4>
  </div>
</template>
<script setup>
import Button from './Button-component.vue';
import Notification from '@/components/layout/Notification-component.vue';
import vSelect from 'vue-select';
import { ref, reactive } from 'vue';
import { user, userToken, onError, currentLocalization } from '@/App.vue';
import { sendRequest } from '@/utils/utils';
import { t } from '@/utils/i18n.js';
import LoginView from '@/views/LoginView.vue';
</script>
<script>
let students = reactive([]);
const selectedStudent = ref('');

let tests = reactive([]);

export default {
  data() {
    return {
      testData: {},
      testId: null,
      completed: false,
      points: 0,
      answers: [],
      open_url: '',
      question_i: 0,
      question: {},
      complete_test: {},
    };
  },
  mounted() {
    if (user.value === 'teacher') {
      tests = [];
      students = [];
      this.getTestForTeacher();
    } else {
      this.getTestByBookId();
    }
  },
  methods: {
    rePassTest() {
      sendRequest(
        `/api/test/repass`,
        'POST',
        {
          complete_test_id: this.complete_test.test_id,
        },
        userToken.value,
      )
        .then((data) => {
          this.complete_test = {};
          this.completed = false;
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
    showTest() {
      const index = parseInt(selectedStudent.value.split('.')[0]);
      this.testData = tests[index - 1];
      this.testId = this.testData.id;
    },
    playAudio(element) {
      document.getElementById(element).play();
    },
    getTestForTeacher() {
      const id = ref(document.URL.split('/content/')[1]);
      sendRequest(`/api/test/info/${id.value}`, 'GET', null, userToken.value)
        .then((data) => {
          if (data) {
            for (let i = 0; i < data.length; i++) {
              students.push(
                `${i + 1}. ${data[i].student.first_name} ${data[i].student.second_name}`,
              );

              tests.push(data[i].test);
            }
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
    async changeQuestion(q, q_index) {
      this.question = q;
      this.question_i = q_index;

      if (this.answers[q_index] !== undefined) {
        await this.sleep(100);
        document.querySelectorAll('.answer-chose')[this.answers[q_index].answer_i].checked = true;
      }
    },
    sleep(ms) {
      return new Promise((resolve) => setTimeout(resolve, ms));
    },
    getTestByBookId() {
      const id = ref(document.URL.split('/content/')[1]);
      sendRequest(`/api/test/${id.value}`, 'GET', null, userToken.value)
        .then((data) => {
          if (data) {
            this.testId = data[0].id;
            this.testData = data[0];
            this.question = this.testData.questions[0];
            this.question_i = 0;
            this.testData.questions.map((q) => {
              q.answers.map((a) => {
                a.chosenAnswer = false;
              });
            });
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
    deleteTest() {
      sendRequest(`/api/test/delete/${this.testId}`, 'DELETE', null, userToken.value)
        .then(() => {
          this.$router.push({ path: '/admin' });
        })
        .catch((err) => {
          onError.value = 'ITDD';
          console.log(err);

          setTimeout(() => {
            onError.value = null;
          }, 5000);
        });
    },
    chooseAnswer(question, questionId, answer, answer_i) {
      this.answers[question - 1] = {
        question_id: questionId,
        answer_id: answer,
        answer_i: answer_i,
      };
    },
    passTest() {
      const json = {
        test_id: this.testData.id,
        answers: [],
      };
      for (let i = 0; i < this.answers.length; i++) {
        json.answers.push({
          answer_id: this.answers[i].answer_id,
          question_id: this.answers[i].question_id,
        });
      }
      console.log(json);
      sendRequest('/api/test/complete', 'POST', json, userToken.value)
        .then((data) => {
          console.log(data);
          this.points = data.points;
          this.completed = true;
          this.complete_test = data;
        })
        .catch((err) => {
          onError.value = 'ITSD';
          console.log(err);

          setTimeout(() => {
            onError.value = null;
          }, 5000);
        });
    },
    openImage(url = '') {
      this.open_url = url;
    },
  },
};
</script>
<style scoped lang="scss">
.dropdown-container {
  display: flex;
  width: 100%;
  justify-content: center;
  align-items: center;
  gap: 20px;

  & .dropdown {
    align-self: center;
    width: 480px;
  }
}

h4.heading-3 {
  text-align: center;
}

.image {
  cursor: pointer;
  width: 480px;
  height: 270px;
  margin: 20px 0;
  border-radius: 12px;
}

.full-image {
  cursor: pointer;
  position: fixed;
  width: 100%;
  height: 100%;
  padding: 40px;
  top: 0;
  left: 0;
  z-index: 50;
  border-radius: 60px;
}

.grid-main {
  margin-bottom: 12px;
  padding-bottom: 12px;
  border-bottom: 2px solid var(--black);

  & .row-main {
    margin-bottom: 4px;
  }

  & .row-secondary {
    margin-bottom: 4px;
  }

  & .row {
    align-items: center;
  }
}
</style>
