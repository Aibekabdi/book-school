<template>
  <Notification class="notification" v-if="onError" :message="onError" />
  <div class="book max-w">
    <div class="sun"></div>
    <div class="grid">
      <nav class="row justify">
        <Button
          v-if="user === 'admin'"
          type="Router"
          link="/admin"
          :label="t(currentLocalization, 'TO_ADMIN_PANEL')"
          color="primary"
          icon="admin_panel_settings"
        />
        <Button
          v-else-if="user !== 'admin'"
          type="Router"
          link="/home"
          :label="t(currentLocalization, 'HOME_PAGE')"
          color="primary"
          icon="home"
        />
        <div class="row">
          <Button
            v-if="selectedTab === 'quiz' && user === 'admin'"
            type="Button"
            @click="deleteTest"
            :label="t(currentLocalization, 'DELETE_TEST')"
            color="danger"
            icon="remove"
          />
          <Button
            v-if="user === 'admin'"
            type="Button"
            @click="deleteBook"
            :label="t(currentLocalization, 'DELETE_BOOK')"
            color="danger"
            icon="delete"
          />
        </div>
      </nav>
      <div class="messages">
        <span v-if="isSuc" class="success-message">{{
          t(currentLocalization, 'CREATE_BOOK_SUC')
        }}</span>
        <span v-else-if="isErr" class="error-message">{{ t(currentLocalization, 'REG_ERR') }}</span>
      </div>
      <Tabs
        :names="tTabs(currentLocalization, tabs)"
        :selectedTab="selectedTab"
        @changeTab="changeTab"
      >
        <div v-if="selectedTab === 'book'">
          <div class="grid" v-if="bookData.title">
            <div class="row justify">
              <div class="image">
                <img :src="bookData.preview" />
              </div>
              <div class="row">
                <Button
                  v-if="current === bookData.text.length - 1"
                  type="Button"
                  @click="complete('books')"
                  :label="t(currentLocalization, 'IREAD')"
                  color="danger"
                  icon="book"
                  size="large"
                />
              </div>
            </div>
            <p class="text-1" v-if="bookData.title">{{ bookData.title }}</p>
            <p class="text-4" v-if="bookData.text" v-html="bookData.text[current]"></p>
          </div>
          <div class="grid" v-else>
            <p class="text-1">{{ t(currentLocalization, 'CANT_LOAD') }}</p>
            <p class="heading-3">{{ t(currentLocalization, 'CHECK_CON') }}</p>
          </div>
          <Pagination
            v-if="bookData.text"
            @any-change="handleCurrent"
            :pages="bookData.text.length"
          />
        </div>
        <div v-if="selectedTab === 'audio_book'">
          <div class="grid" v-if="bookData.title">
            <div class="row justify">
              <div class="image">
                <img :src="bookData.preview" />
              </div>
              <div class="row">
                <audio
                  controls
                  :src="bookData.audio[current]"
                  id="audio"
                  style="display: none"
                ></audio>
                <Button
                  v-if="current !== bookData.text.length - 1"
                  type="Button"
                  @click="playAudio()"
                  :label="t(currentLocalization, 'LISTEN')"
                  color="info"
                  icon="play_arrow"
                  size="large"
                />
                <Button
                  v-else
                  type="Button"
                  @click="playAudio(), complete('audio')"
                  :label="t(currentLocalization, 'LISTEN')"
                  color="info"
                  icon="play_arrow"
                  size="large"
                />
              </div>
            </div>
            <p class="text-1" v-if="bookData.title">{{ bookData.title }}</p>
            <p class="text-4" v-if="bookData.text" v-html="bookData.text[current]"></p>
          </div>
          <div class="grid" v-else>
            <p class="text-1">{{ t(currentLocalization, 'CANT_LOAD') }}</p>
            <p class="heading-3">{{ t(currentLocalization, 'CHECK_CON') }}</p>
          </div>
          <Pagination
            v-if="bookData.text"
            @any-change="handleCurrent"
            :pages="bookData.text.length"
          />
        </div>
        <div v-else-if="selectedTab === 'quiz'">
          <p class="text-2">{{ t(currentLocalization, 'TEST') }}</p>
          <Test />
        </div>
        <div v-else-if="selectedTab === 'studio'">
          <p class="text-2">{{ t(currentLocalization, 'ART_STUDIO') }}</p>
          <CreativeTask :category="bookData.category" />
        </div>
        <div v-else-if="selectedTab === 'journal'">
          <p class="text-2">{{ t(currentLocalization, 'JOURNAL') }}</p>
          <div v-if="user === 'teacher'">
            <Journal />
          </div>
          <div v-else>
            <p class="text-4 error-message">{{ t(currentLocalization, 'NO_DATA') }}</p>
          </div>
        </div>
      </Tabs>
    </div>
  </div>
</template>
<script setup>
import { ref, reactive } from 'vue';
import Button from '@/components/Button-component.vue';
import Tabs from '@/components/Tab-component.vue';
import Pagination from '@/components/Pagination-component.vue';
import Test from '@/components/Test-component.vue';
import CreativeTask from '@/components/CreativeTask-component.vue';
import Journal from '@/components/Journal-component.vue';
import Notification from '@/components/layout/Notification-component.vue';
import { sendRequest } from '@/utils/utils';
import { user, userToken, onError, currentLocalization } from '@/App.vue';
import { t, tTabs } from '@/utils/i18n.js';
</script>
<script>
const current = ref(0);

const bookData = reactive({
  title: '',
  text: '',
  category: '',
  audio: '',
  id: document.URL.split('/content/')[1],
});

const isSuc = ref(false);
const isErr = ref(false);

let tabs = [
  { name: 'book', label: 'BOOK' },
  { name: 'audio_book', label: 'AUDIO_BOOK' },
  { name: 'quiz', label: 'TEST' },
  { name: 'studio', label: 'ART_STUDIO' },
];

const selectedTab = ref('book');
const changeTab = (tabName) => {
  selectedTab.value = tabName;
};

export default {
  name: 'ContentView',
  components: { Button, Tabs, Pagination, Test, CreativeTask, Journal, Notification },
  destroyed() {
    document.body.classList.remove('change-back');
  },
  mounted() {
    document.body.classList.add('change-back');

    user.value = localStorage.getItem('admin') ? 'admin' : localStorage.getItem('userType') || '';
    userToken.value = localStorage.getItem('admin') || localStorage.getItem('userToken') || '';

    tabs = [
      { name: 'book', label: 'BOOK' },
      { name: 'audio_book', label: 'AUDIO_BOOK' },
      { name: 'quiz', label: 'TEST' },
      { name: 'studio', label: 'ART_STUDIO' },
    ];

    if (user.value === 'teacher') {
      tabs = [
        { name: 'book', label: 'BOOK' },
        { name: 'audio_book', label: 'AUDIO_BOOK' },
        { name: 'quiz', label: 'TEST' },
        { name: 'studio', label: 'ART_STUDIO' },
        { name: 'journal', label: 'JOURNAL' },
      ];
    }

    selectedTab.value = 'book';

    this.getBookById();
  },
  methods: {
    handleCurrent(s) {
      current.value = s;
      console.log('current:', current.value);
    },
    getBookById() {
      bookData.id = document.URL.split('/content/')[1];
      sendRequest(`/api/books/${bookData.id}`, 'GET', null, userToken.value)
        .then((data) => {
          if (data) {
            bookData.title = data.name;
            bookData.text = data.pages;
            bookData.preview = data.preview;
            bookData.category = data.category;
            bookData.audio = data.audio;
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
    deleteBook() {
      isSuc.value = false;
      isErr.value = false;
      sendRequest(`/api/books/delete/${bookData.id}`, 'DELETE', null, userToken.value)
        .then(() => {
          isSuc.value = true;
          this.$router.push({ path: '/admin' });
        })
        .catch(() => {
          console.log(bookData.id);
          isErr.value = true;
        });
    },
    playAudio() {
      console.log(bookData.audio[2]);
      document.getElementById('audio').play();
    },
    complete(type) {
      sendRequest(`/api/${type}/complete/${bookData.id}`, 'POST', null, userToken.value)
        .then((data) => {
          console.log(data);
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
};
</script>
<style scoped lang="scss">
.book {
  position: relative;
}

img {
  width: 100%;
  height: auto;
  border-radius: 12px;
  margin: 20px 0;
}

.image {
  max-width: 280px;
  width: 100%;
}

.row {
  margin-bottom: 20px;
}

.text-1,
.text-4 {
  font-size: 32px;
  max-width: 1440px;
  word-break: break-all;
}

.messages {
  text-align: center;
  margin: 20px 0;
}

img {
  margin: 0;
  width: 100%;
  max-width: 100% !important;
}
</style>
