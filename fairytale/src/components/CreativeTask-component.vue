<template>
  <Notification class="notification" v-if="onError" :message="onError" />
  <div v-if="creativeTaskData">
    <h3 class="heading-2">{{ t(currentLocalization, 'LIST_OF_TASKS') }}:</h3>
    <div v-if="creativeTaskData.length === 0">
      <span class="heading-3 success-message">Здесь пусто, кажется вы все сделали)</span>
    </div>
    <div v-for="(task, i) in creativeTaskData" :key="task">
      <div class="row">
        <span class="text-4"
          ><b>{{ i + 1 }}.</b> {{ task.question }}</span
        >
        <Button
          v-if="!task.is_art && user === 'student'"
          type="Button"
          @click="showModal(task)"
          :label="t(currentLocalization, 'START_WRITE')"
          color="second"
          icon="edit_document"
        />
        <Button
          v-else-if="task.is_art && user === 'student'"
          type="Button"
          @click="showModal(task)"
          :label="t(currentLocalization, 'START_DRAW')"
          color="info"
          icon="draw"
        />
        <Button
          v-else-if="!task.is_art && user !== 'student'"
          type="Button"
          @click="showModal(task)"
          label="Начать писать"
          color="second"
          icon="edit_document"
          :disabled="true"
        />
        <Button
          v-else-if="task.is_art && user !== 'student'"
          type="Button"
          @click="showModal(task)"
          label="Начать рисовать"
          color="info"
          icon="draw"
          :disabled="true"
        />
      </div>
    </div>
    <div class="modal" v-if="open">
      <div>
        <div class="close">
          <Button type="Button" @click="open = !open" icon="close" :rounded="true" color="black" />
        </div>
        <div class="modal-content" v-if="thisTask.is_art">
          <h3 class="heading-2">{{ t(currentLocalization, 'DPFT') }}:</h3>
          <h5 class="heading-3">{{ thisTask.question }}</h5>
          <ArtStudio :BookId="+book_id" :QuestionId="+thisTask.id" />
          <!-- ART COMPONENT -->
          <Button
            v-if="!thisTask.is_art"
            type="Button"
            @click="submitCreativeTask"
            :label="t(currentLocalization, 'SEND')"
            color="primary"
            icon="send"
            size="large"
          />
        </div>
        <div class="modal-content" v-else>
          <h3 class="heading-2">{{ t(currentLocalization, 'WSTFT') }}:</h3>
          <h5 class="heading-3">{{ thisTask.question }}</h5>
          <textarea
            v-model="textarea"
            :placeholder="t(currentLocalization, 'SWH') + '...'"
          ></textarea>
          <div class="messages">
            <span v-if="isSuc" class="success-message">Ответ был успешно отправлен!</span>
            <span v-else-if="isErr" class="error-message">Ошибка. Попробуйте снова</span>
          </div>
          <Button
            type="Button"
            @click="submitCreativeTask"
            :label="t(currentLocalization, 'SEND')"
            color="primary"
            icon="send"
            size="large"
          />
        </div>
      </div>
    </div>
  </div>
  <div v-else>
    <h4 class="heading-3">{{ t(currentLocalization, 'NO_DATA') }}</h4>
  </div>
</template>
<script setup>
import Button from './Button-component.vue';
import ArtStudio from './ArtStudio-component.vue';
import Notification from '@/components/layout/Notification-component.vue';
import { ref, defineProps } from 'vue';
import { sendRequest } from '@/utils/utils';
import { user, userToken, onError, currentLocalization } from '@/App.vue';
import host from '@/main';
import { t } from '@/utils/i18n.js';

defineProps({
  category: {
    type: String,
    required: true,
  },
});
</script>
<script>
const isSuc = ref(false);
const isErr = ref(false);

const open = ref(false);
const creativeTaskData = ref({});
const thisTask = ref({});
const book_id = ref(document.URL.split('/content/')[1]);

export default {
  mounted() {
    book_id.value = document.URL.split('/content/')[1];
    this.getCreativeTask();
  },
  methods: {
    showModal(task) {
      open.value = !open.value;

      thisTask.value = task;
    },

    getCreativeTask() {
      sendRequest(`/api/creative/tasks/get/${this.category}`, 'GET', null, userToken.value)
        .then((data) => {
          creativeTaskData.value = data;
          console.log(creativeTaskData.value);
        })
        .catch((err) => {
          onError.value = 'ITRD';
          console.log(err);

          setTimeout(() => {
            onError.value = null;
          }, 5000);
        });
    },
    async submitCreativeTask() {
      isSuc.value = false;
      isErr.value = false;
      const json = {
        book_id: +book_id.value,
        question_id: +thisTask.value.id,
        is_art: false,
        answer: this.textarea,
      };
      const form = new FormData();

      form.append('document', JSON.stringify(json));

      let response = await fetch(`${host}/api/creative/pass/create`, {
        method: 'POST',
        body: form,
        headers: {
          // 'Content-type': 'multipart/form-data; charset=UTF-8',
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

        setTimeout(() => {
          onError.value = null;
        }, 5000);
      } else {
        console.log('OK');
      }
    },
  },
  data() {
    return {
      textarea: '',
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
.heading-3 {
  margin: 0;
  width: 580px;
  word-wrap: break-word;
}

.text-4 {
  max-width: 720px;
  word-break: break-all;
}
.row {
  margin: 10px 0;
  align-items: center;
}
</style>
