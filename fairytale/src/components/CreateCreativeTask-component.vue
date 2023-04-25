<template>
  <div class="global">
    <div class="grid create">
      <div class="grid">
        <Input
          :label="t(currentLocalization, 'TITLE')"
          name="headline"
          :placeholder="t(currentLocalization, 'ENOOQ')"
          v-model:value="creativeTask.question"
          width="100%"
        />
      </div>
      <div class="grid">
        <div class="row centerize">
          <span class="heading-3">{{ t(currentLocalization, 'ART_TASK') }}:</span>
          <input type="checkbox" name="checkbox" id="checkbox" v-model="creativeTask.isArt" />
        </div>
        <span class="text-4">{{
          creativeTask.isArt ? t(currentLocalization, 'CAT') : t(currentLocalization, 'COQ')
        }}</span>
      </div>
      <div class="grid">
        <span class="heading-3">{{ t(currentLocalization, 'SELECT_CATEGORY') }}:</span>
        <v-select
          :options="categories.map((e) => t(currentLocalization, e))"
          v-model="creativeTask.category"
        ></v-select>
      </div>
      <Button
        type="Button"
        @click="createCreativeTask"
        color="success"
        :label="t(currentLocalization, 'CREATE_QUESTION')"
        width="100%"
      />
    </div>
    <div class="messages">
      <span v-if="isSuc === 'создан'" class="success-message"
        >{{ t(currentLocalization, 'OQCS') }}!</span
      >
      <span v-else-if="isSuc === 'удален'" class="success-message"
        >{{ t(currentLocalization, 'OQDS') }}!</span
      >
      <span v-else-if="isErr" class="error-message">{{ t(currentLocalization, 'REG_ERR') }}</span>
    </div>
    <div class="grid dashboard">
      <div class="row justify" v-for="(question, idx) in creativeTasks" :key="question">
        <div class="grid grid-name" :id="idx + 1 + ' line'">{{ idx + 1 }}</div>
        <div class="grid grid-name" :id="idx + 1 + ' line'">{{ question.category }}</div>
        <div class="grid grid-name" :id="idx + 1 + ' line'">
          {{ question.is_art ? t(currentLocalization, 'AT') : t(currentLocalization, 'OQ') }}
        </div>
        <div class="grid grid-name" :id="idx + 1 + ' line'">{{ question.question }}</div>
        <div class="grid grid-name">
          <div class="row">
            <Button type="Button" @click="showUpdateModal(idx + 1)" color="primary" icon="edit" />
            <div class="inner-modal" v-if="open">
              <div>
                <div class="close">
                  <Button
                    type="Button"
                    @click="showUpdateModal(idx + 1)"
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
                      <span class="text-4">{{ t(currentLocalization, 'CATEGORY') }}:</span>
                    </div>
                    <div class="grid">
                      <span class="text-4">{{ t(currentLocalization, 'TYPE') }}:</span>
                    </div>
                    <div class="grid">
                      <span class="text-4">{{ t(currentLocalization, 'QUESTION') }}:</span>
                    </div>
                  </div>
                  <div class="row justify">
                    <div class="grid grid-text">{{ updateTask.id }}</div>
                    <div class="grid grid-text">{{ updateTask.category }}</div>
                    <div class="grid grid-text">
                      {{
                        updateTask.isArt
                          ? t(currentLocalization, 'AT')
                          : t(currentLocalization, 'OQ')
                      }}
                    </div>
                    <div class="grid grid-text">{{ updateTask.question }}</div>
                  </div>
                  <div class="grid create">
                    <div class="grid">
                      <Input
                        :label="t(currentLocalization, 'TITLE')"
                        name="headline"
                        :placeholder="t(currentLocalization, 'ENOOQ')"
                        v-model:value="creativeTask.question"
                        width="100%"
                      />
                    </div>
                    <div class="grid">
                      <div class="row centerize">
                        <span class="heading-3">{{ t(currentLocalization, 'AT') }}:</span>
                        <input
                          type="checkbox"
                          name="checkbox"
                          id="checkbox"
                          v-model="creativeTask.isArt"
                        />
                      </div>
                      <span class="text-4">{{
                        creativeTask.isArt
                          ? t(currentLocalization, 'CAT')
                          : t(currentLocalization, 'COQ')
                      }}</span>
                    </div>
                    <div class="grid">
                      <span class="heading-3"
                        >{{ t(currentLocalization, 'SELECT_CATEGORY') }}:</span
                      >
                      <v-select
                        :options="categories.map((e) => t(currentLocalization, e))"
                        v-model="creativeTask.category"
                      ></v-select>
                    </div>
                    <Button
                      type="Button"
                      @click="updateCreativeTask"
                      color="success"
                      :label="t(currentLocalization, 'UPDATE_QUESTION')"
                      width="100%"
                    />
                  </div>
                  <div class="messages">
                    <span v-if="isSuc" class="success-message"
                      >{{ t(currentLocalization, 'OQUS') }}!</span
                    >
                    <span v-else-if="isErr" class="error-message">{{
                      t(currentLocalization, 'REG_ERR')
                    }}</span>
                  </div>
                </div>
              </div>
            </div>
            <Button
              type="Button"
              @click="deleteCreativeTask(creativeTasks[idx])"
              color="second"
              icon="delete"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup>
import { ref, reactive } from 'vue';
import Button from './Button-component.vue';
import Input from './Input-component.vue';
import vSelect from 'vue-select';
import { sendRequest } from '@/utils/utils';
import { userToken, currentLocalization } from '@/App.vue';
import { t } from '@/utils/i18n.js';
</script>
<script>
const isSuc = ref('');
const isErr = ref(false);

const open = ref(false);

var creativeTasks = ref([]);
const creativeTask = reactive({
  id: null,
  question: '',
  category: '',
  isArt: false,
});

const updateTask = reactive({
  id: 0,
  question: '',
  category: '',
  isArt: false,
});

const categories = [
  'YEAR_2',
  'YEAR_3',
  'YEAR_4',
  'YEAR_5',
  'CLASS_1',
  'CLASS_2',
  'CLASS_3',
  'CLASS_4',
];

export default {
  components: { vSelect, Button, Input },
  mounted() {
    this.getAllCreativeTasks();
  },
  methods: {
    createCreativeTask() {
      const json = {
        category: creativeTask.category,
        question: creativeTask.question,
        is_art: creativeTask.isArt,
      };

      sendRequest('/api/creative/tasks/create', 'POST', json, userToken.value)
        .then((data) => {
          isSuc.value = 'создан';
          json.id = data;
          creativeTasks.value.push(json);
        })
        .catch((err) => {
          isErr.value = true;
        });
    },
    getAllCreativeTasks() {
      isSuc.value = '';
      isErr.value = false;
      creativeTasks.value = [];
      sendRequest('/api/creative/tasks/get', 'GET', null, userToken.value)
        .then((data) => {
          if (data) {
            data.forEach((el) => {
              creativeTasks.value.push(el);
            });
          }
          console.log(creativeTasks.value);
        })
        .catch((err) => {
          isErr.value = true;
        });
    },
    showUpdateModal(index = 0) {
      isSuc.value = '';
      isErr.value = false;
      if (index != 0) {
        updateTask.id = creativeTasks.value[index - 1].id;
        updateTask.category = creativeTasks.value[index - 1].category;
        updateTask.question = creativeTasks.value[index - 1].question;
        updateTask.isArt = creativeTasks.value[index - 1].is_art;
        console.log(updateTask);
      }

      open.value = !open.value;
    },
    updateCreativeTask() {
      isSuc.value = '';
      isErr.value = false;

      const json = {
        id: updateTask.id,
        category: creativeTask.category,
        question: creativeTask.question,
        is_art: creativeTask.isArt,
      };

      sendRequest('/api/creative/tasks/update', 'PATCH', json, userToken.value)
        .then((data) => {
          isSuc.value = 'обновлен';
          open.value = false;
          creativeTask.category = '';
          creativeTask.question = '';
          creativeTask.isArt = false;
          this.getAllCreativeTasks();
        })
        .catch((err) => {
          isErr.value = true;
        });
    },
    deleteCreativeTask(obj) {
      isSuc.value = false;
      isErr.value = false;
      console.log(obj.id);

      sendRequest(`/api/creative/tasks/delete/${obj.id}`, 'DELETE', null, userToken.value)
        .then(() => {
          isSuc.value = 'удален';
        })
        .catch(() => {
          isErr.value = true;
        });
    },
  },
};
</script>
<style scoped lang="scss">
.global {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.create {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  gap: 20px;

  width: 600px;

  margin-bottom: 40px;

  & > .grid {
    width: 100%;
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

  & .heading-3 {
    color: var(--white);
  }

  & .grid-text {
    color: var(--white-hover);
    max-width: 25%;
    word-wrap: break-word;
  }

  & .text-4 {
    color: var(--white-hover);
  }

  & .create {
    border-radius: 12px;
    padding: 20px 10px;
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
.grid-name {
  max-width: 25%;
  word-wrap: break-word;
}

.messages {
  margin-bottom: 20px;
}
.row.justify {
  margin-bottom: 10px;
  padding-bottom: 10px;
  align-items: center;
}
</style>
