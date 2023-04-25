<template>
  <Notification class="notification" v-if="onError" :message="onError" />
  <div class="grid centerize">
    <div class="dropdown">
      <span class="heading-3">{{ t(currentLocalization, 'SELECT_BOOK') }}:</span>
      <v-select :options="booksList" v-model="selectedBook"></v-select>
    </div>
    <Tabs
      :names="tTabs(currentLocalization, tabs)"
      :selectedTab="selectedTab"
      @changeTab="changeTab"
      width="900px"
    >
      <div class="grid centerize" v-if="selectedTab === 'Kindergarden'">
        <div class="grid">
          <div class="grid">
            <div class="grid relative" v-for="(question, idx) in questionsList" :key="idx">
              <div class="row">
                <div class="btn-remove">
                  <Button
                    type="Button"
                    @click="removeQuestion"
                    color="danger"
                    icon="delete"
                    :disabled="idx === 0"
                    width="40px"
                  />
                </div>
                <component
                  :is="Input"
                  :key="question"
                  :name="idx + 1 + ' question'"
                  :label="idx + 1 + '. Вопрос'"
                  v-model:value="question.ref"
                  :placeholder="t(currentLocalization, 'ENOQ')"
                  width="632px"
                />
                <div class="input-container">
                  <label
                    tabindex="0"
                    :for="question.id + ' input'"
                    :class="
                      questionsList[idx].withImage ? 'uploaded-file-trigger' : 'upload-file-trigger'
                    "
                    ><span class="material-symbols-outlined">
                      {{ questionsList[idx].withImage ? 'download_done' : 'upload_file' }}
                    </span></label
                  >
                  <input
                    type="file"
                    class="upload-file"
                    :id="question.id + ' input'"
                    accept="image/*"
                    title=""
                    @change="
                      setQuestionFiles($event.target.name, $event.target.files, idx),
                        setQuestionWithImage(idx)
                    "
                  />
                </div>
              </div>
              <div class="grid">
                <!-- <div class="check-answer">
                  <span class="material-symbols-outlined"> radio_button_checked </span>
                </div> -->
                <div class="grid" v-for="(answer, index) in question.answers" :key="index">
                  <div class="row justify" v-if="index < 3">
                    <component
                      :is="Input"
                      :key="answer"
                      :name="index + 1 + ' answer'"
                      :label="index + 1 + '. Ответ'"
                      v-model:value="answer.ref"
                      :placeholder="t(currentLocalization, 'ENOA')"
                      width="632px"
                    />
                    <div class="input-container">
                      <label
                        tabindex="0"
                        :for="answer.id + ' upload'"
                        :class="
                          questionsList[idx].answers[index].withImage
                            ? 'uploaded-file-trigger'
                            : 'upload-file-trigger'
                        "
                        ><span class="material-symbols-outlined">
                          {{
                            questionsList[idx].answers[index].withImage
                              ? 'download_done'
                              : 'upload_file'
                          }}
                        </span></label
                      >
                      <input
                        type="file"
                        class="upload-file"
                        :id="answer.id + ' upload'"
                        accept="image/*"
                        :title="t(currentLocalization, 'PTCPTU')"
                        @change="
                          setAnswerFiles($event.target.name, $event.target.files, idx, index),
                            setAnswerWithImage(idx, index)
                        "
                      />
                    </div>
                    <div class="radio-container">
                      <Button
                        type="Button"
                        :icon="questionsList[idx].answers[index].radio ? 'done' : 'remove'"
                        :color="questionsList[idx].answers[index].radio ? 'success' : 'warning'"
                        size="small"
                        width="40px"
                        @click="choseAnswer(idx, index)"
                      />
                      <!-- <input
                        type="radio"
                        :name="idx + '.' + 'answer'"
                        :id="idx + '.' + 'answer'"
                        v-model="answer.radio"
                      /> -->
                    </div>
                  </div>
                </div>
              </div>
              <hr />
            </div>
          </div>
        </div>
        <Button
          type="Button"
          @click="addQuestion"
          color="success"
          :label="t(currentLocalization, 'CREATE_QUESTION')"
          width="200px"
        />
      </div>
      <div class="grid centerize" v-if="selectedTab === 'Pupils'">
        <div class="grid">
          <div class="grid">
            <div class="grid relative" v-for="(question, idx) in questionsList" :key="idx">
              <div class="row">
                <div class="btn-remove">
                  <Button
                    type="Button"
                    @click="removeQuestion"
                    color="danger"
                    icon="delete"
                    :disabled="idx === 0"
                    width="40px"
                  />
                </div>
                <component
                  :is="Input"
                  :key="question"
                  :name="idx + 1 + ' question'"
                  :label="idx + 1 + '. Вопрос'"
                  v-model:value="question.ref"
                  :placeholder="t(currentLocalization, 'ENOQ')"
                  width="632px"
                />
                <div class="input-container">
                  <label
                    tabindex="0"
                    :for="question.id + ' input'"
                    :class="
                      questionsList[idx].withImage ? 'uploaded-file-trigger' : 'upload-file-trigger'
                    "
                    ><span class="material-symbols-outlined">
                      {{ questionsList[idx].withImage ? 'download_done' : 'upload_file' }}
                    </span></label
                  >
                  <input
                    type="file"
                    class="upload-file"
                    :id="question.id + ' input'"
                    accept="image/*"
                    :title="t(currentLocalization, 'PTCPTU')"
                    @change="
                      setQuestionFiles($event.target.name, $event.target.files, idx),
                        setQuestionWithImage(idx)
                    "
                  />
                </div>
              </div>
              <div class="grid">
                <div class="grid" v-for="(answer, index) in question.answers" :key="index">
                  <div class="row justify">
                    <component
                      :is="Input"
                      :key="answer"
                      :name="index + 1 + ' answer'"
                      :label="index + 1 + '. Ответ'"
                      v-model:value="answer.ref"
                      :placeholder="t(currentLocalization, 'ENOA')"
                      width="632px"
                    />
                    <div class="input-container">
                      <label
                        tabindex="0"
                        :for="answer.id + ' upload'"
                        :class="
                          questionsList[idx].answers[index].withImage
                            ? 'uploaded-file-trigger'
                            : 'upload-file-trigger'
                        "
                        ><span class="material-symbols-outlined">
                          {{
                            questionsList[idx].answers[index].withImage
                              ? 'download_done'
                              : 'upload_file'
                          }}
                        </span></label
                      >
                      <input
                        type="file"
                        class="upload-file"
                        :id="answer.id + ' upload'"
                        accept="image/*"
                        :title="t(currentLocalization, 'PTCPTU')"
                        @change="
                          setAnswerFiles($event.target.name, $event.target.files, idx, index),
                            setAnswerWithImage(idx, index)
                        "
                      />
                    </div>
                    <div class="radio-container">
                      <!-- <Button :name="idx + '.' + 'answer'" :id="idx + '.' + 'answer'" @click="" /> -->
                      <Button
                        type="Button"
                        :icon="questionsList[idx].answers[index].radio ? 'done' : 'remove'"
                        :color="questionsList[idx].answers[index].radio ? 'success' : 'warning'"
                        size="small"
                        width="40px"
                        @click="choseAnswer(idx, index)"
                      />
                      <!-- <input
                        type="radio"
                        :name="idx + '.' + 'answer'"
                        :id="idx + '.' + 'answer'"
                        v-model="answer.radio"
                      /> -->
                    </div>
                  </div>
                </div>
              </div>
              <hr />
            </div>
          </div>
        </div>
        <Button
          type="Button"
          @click="addQuestion"
          color="success"
          :label="t(currentLocalization, 'CREATE_QUESTION')"
          width="200px"
        />
      </div>
    </Tabs>
    <div class="grid create">
      <Button
        type="Button"
        @click="createTest"
        color="warning"
        :label="t(currentLocalization, 'CREATE_TEST')"
        width="200px"
      />
    </div>
    <div class="messages">
      <span v-if="isSuc" class="success-message">{{ t(currentLocalization, 'AWSS') }}!</span>
      <span v-else-if="isErr" class="error-message">{{ t(currentLocalization, 'REG_ERR') }}</span>
    </div>
  </div>
</template>
<script setup>
import { ref } from 'vue';
import { sendRequest } from '@/utils/utils';
import vSelect from 'vue-select';
import Tabs from '@/components/Tab-component.vue';
import Input from '@/components/Input-component.vue';
import Button from '@/components/Button-component.vue';
import Notification from '@/components/layout/Notification-component.vue';
import { userToken, onError, currentLocalization } from '@/App.vue';
import host from '@/main';
import { t, tTabs } from '@/utils/i18n.js';
</script>
<script>
const isSuc = ref(false);
const isErr = ref(false);

const tabs = [
  { name: 'Kindergarden', label: 'KINDER' },
  { name: 'Pupils', label: 'PUPIL' },
];

const selectedTab = ref('Kindergarden');
const changeTab = (tabName) => {
  selectedTab.value = tabName;
};

const selectedBook = ref('');
const booksList = ref([]);

export default {
  data() {
    return {
      questionFiles: [],
      answerFiles: [],
      questionsList: [],
    };
  },
  mounted() {
    this.getBooksList();
    this.addQuestion();
  },
  methods: {
    addQuestion() {
      if (this.questionsList.length >= 12) {
        onError.value = 'CCMT12Q';

        setTimeout(() => {
          onError.value = null;
        }, 5000);
        return;
      }

      this.questionsList.push({
        id: '',
        ref: '',
        withImage: false,
        answers: [
          {
            id: '',
            ref: '',
            withImage: false,
            radio: false,
          },
          {
            id: '',
            ref: '',
            withImage: false,
            radio: false,
          },
          {
            id: '',
            ref: '',
            withImage: false,
            radio: false,
          },
          {
            id: '',
            ref: '',
            withImage: false,
            radio: false,
          },
          {
            id: '',
            ref: '',
            withImage: false,
            radio: false,
          },
        ],
      });
      console.log(this.questionsList);
    },
    choseAnswer: function (questionId, answerId) {
      for (let index = 0; index < this.questionsList[questionId].answers.length; index++) {
        this.questionsList[questionId].answers[index].radio = false;
      }

      this.questionsList[questionId].answers[answerId].radio = true;
    },
    removeQuestion() {
      this.questionsList.pop();
    },
    getBooksList() {
      sendRequest('/api/books', 'GET', null, userToken.value)
        .then((data) => {
          if (data) {
            data.forEach((book,i) => {
              booksList.value[i] = book.id + '. ' + book.name;
            });
          }
        })
        .catch((err) => {
          console.log(err);
        });
    },
    setQuestionFiles(fieldName, fileList, questionIndex) {
      this.questionFiles[questionIndex] = fileList[0];
    },
    setAnswerFiles(fieldName, fileList, questionIndex, answerIndex) {
      this.answerFiles[questionIndex * 3 + answerIndex] = fileList[0];
    },
    setQuestionWithImage(questionIndex) {
      this.questionsList[questionIndex].withImage = true;
    },
    setAnswerWithImage(questionIndex, answerIndex) {
      this.questionsList[questionIndex].answers[answerIndex].withImage = true;
    },
    async createTest() {
      isSuc.value = false;
      isErr.value = false;

      const form = new FormData();

      const json = {
        book_id: +selectedBook.value.split('.')[0],
        questions: [],
      };

      console.log(this.questionFiles);
      console.log(this.answerFiles);

      if (this.questionFiles) {
        for (let k = 0; k < this.questionFiles.length; k++) {
          if (this.questionFiles[k] === 'undefined') {
            continue;
          }
          form.append('question', this.questionFiles[k]);
        }
      }

      if (this.answerFiles) {
        for (let k = 0; k < this.answerFiles.length; k++) {
          if (this.answerFiles[k] === 'undefined') {
            continue;
          }
          form.append('answer', this.answerFiles[k]);
        }
      }

      for (let i = 0; i < this.questionsList.length; i++) {
        const question = {
          question: this.questionsList[i].ref,
          with_image: this.questionsList[i].withImage,
          answers: [],
        };

        for (let index = 0; index < this.questionsList[i].answers.length; index++) {
          if (selectedTab.value === 'Kindergarden' && index === 3) {
            break;
          }

          question.answers.push({
            answer: this.questionsList[i].answers[index].ref,
            with_image: this.questionsList[i].answers[index].withImage,
            correct: this.questionsList[i].answers[index].radio,
          });
        }
        json.questions.push(question);
      }

      console.log(json);

      form.append('document', JSON.stringify(json));

      console.log(form.getAll('question'));
      console.log(form.getAll('answer'));

      let response = await fetch(`${host}/api/test/create`, {
        method: 'POST',
        body: form,
        headers: {
          Accept: '*/*',
          Authorization: `Bearer ${userToken.value}`,
          'Access-Control-Allow-Origin': `${host}/`,
          'Access-Control-Allow-Methods': 'GET, POST, OPTIONS, PUT, PATCH, DELETE',
          'Access-Control-Allow-Headers': 'origin,X-Requested-With,content-type,accept',
          'Access-Control-Allow-Credentials': 'true',
        },
      });
      let result = await response.json();

      if (result.status != 'OK') {
        onError.value = 'ITSD';
        console.log('NOT OK');
        isErr.value = true;

        setTimeout(() => {
          onError.value = null;
        }, 5000);
      } else {
        console.log('OK');
        isSuc.value = true;
      }
    },
  },
};
</script>
<style scoped lang="scss">
.check-answer {
  position: absolute;
  right: -30px;
  top: 50%;
  color: var(--white);
}

.messages {
  margin-top: 20px;
}
div.grid.centerize {
  align-items: center;
  width: fit-content;

  & .relative {
    position: relative;
  }

  & div.grid.centerize {
    width: 100%;

    & .grid-title {
      width: 180px;
      text-align: right;
      justify-content: center;
      margin-bottom: 30px;
    }

    & .grid-input {
      // width: 100%;
    }

    & .grid-image {
      & > .row {
        margin-top: 8px;
      }
    }

    & .grid-select {
      margin-top: 10px;
      display: flex;
    }
  }
}

.btn-remove {
  right: 0;
  top: 0;
  position: absolute;
}

hr {
  width: 100%;
  height: 2px;
  margin-bottom: 30px;
  color: var(--black-hover);
  background-color: var(--black-hover);
}

.create {
  margin-top: 20px;
}

.input-container {
  position: relative;
  width: 40px;
  & .upload-file {
    cursor: pointer;
    position: absolute;
    top: 0;
    left: 0;
    width: 40px;
    opacity: 0;
    height: 40px;
  }

  & .upload-file-trigger {
    display: block;
    padding: 8px;
    color: var(--white);
    border-radius: 12px;
    background-color: var(--primary);
    height: 40px;
    z-index: 10;
  }

  & .uploaded-file-trigger {
    display: block;
    padding: 8px;
    color: var(--white);
    border-radius: 12px;
    background-color: var(--success);
    height: 40px;
    z-index: 10;
  }
}

.dropdown {
  width: 360px;
  margin-bottom: 20px;
  display: flex;
  flex-direction: column;
  gap: 10px;
}
</style>
