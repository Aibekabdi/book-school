<template>
  <Notification class="notification" v-if="onError" :message="onError" />
  <div class="grid" v-if="categories.length !== 0">
    <div class="content-area" v-for="category in categories" :key="category">
      <div v-if="category.books !== null">
        <h2 class="heading-2">{{ category.name }}</h2>
        <div class="row">
          <div class="row card" v-for="book in category.books" :key="book.id">
            <router-link
              class="grid"
              :key="book"
              :to="{ name: 'content', params: { content_id: book.id } }"
              s
            >
              <img :src="book.preview" :alt="book.name" width="120" height="120" loading="lazy" />
              <h4 class="heading-3">{{ book.name }}</h4>
            </router-link>
          </div>
        </div>
      </div>
      <!-- <router-link class="row card" v-for="books in bookData" :key="books" to="/home" s>
        <img :src="books.preview" :alt="books.title" width="120" height="120" loading="lazy" />
        <div class="grid">
          <div class="grid">
            <h4 class="heading-3">{{ books.title }}</h4>
            <span class="class">{{ books.class }}</span>
            <span class="category">{{ books.category }}</span>
          </div>
        </div>
      </router-link> -->
    </div>
  </div>
  <div class="centerize" v-else>
    <h4 class="heading-3 error-message">{{ t(currentLocalization, 'NO_DATA') }}</h4>
    <h5 class="text-4" v-if="user === 'admin'">{{ t(currentLocalization, 'FCAB') }}</h5>
  </div>
</template>

<script setup>
import Notification from '@/components/layout/Notification-component.vue';
import { ref } from 'vue';
import { sendRequest } from '@/utils/utils';
import { user, userToken, onError, currentLocalization } from '@/App.vue';
import { t } from '@/utils/i18n.js';
</script>
<script>
const categories = ref([]);

export default {
  methods: {
    // getPages() {
    //   sendRequest('/api/books/total', 'GET', null, userToken.value)
    //     .then((data) => {
    //       if (data) {
    //         numberOfPages.value = data.status;
    //       }
    //     })
    //     .catch((err) => {
    //       onError.value = 'ITRD';
    //       console.log(err);

    //       setTimeout(() => {
    //         onError.value = null;
    //       }, 5000);
    //     });
    // },
    getBooks() {
      sendRequest(`/api/books/all/all`, 'GET', null, userToken.value)
        .then((data) => {
          console.log(data);
          data.forEach((el) => {
            if (el.books !== null) {
              categories.value.push({
                name: el.categories,
                books: el.books,
              });
            }
          });
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
  mounted() {
    user.value = localStorage.getItem('admin') ? 'admin' : localStorage.getItem('userType') || '';
    userToken.value = localStorage.getItem('admin') || localStorage.getItem('userToken') || '';

    this.getBooks();
    // this.getPages();
    categories.value = [];
  },
};
</script>
<style scoped lang="scss">
.content-area {
  display: flex;
  flex-flow: row wrap;
  justify-content: center;
  align-items: center;
  gap: 20px;
  margin: 20px 0 60px;

  & > div {
    width: 100%;

    & > .row {
      flex-wrap: wrap;
      transition: all 0.4s;

      & > .card {
        width: 320px;
        min-width: 320px;
        height: 480px;
        background-color: var(--light-blue);
        border-radius: 12px;
        padding: 12px;
        cursor: pointer;
        transition: all 0.4s;

        & > a.grid {
          width: 100%;

          overflow: hidden;

          text-overflow: ellipsis;
        }

        & img {
          width: 100%;
          height: 100%;
          border-radius: 10px;
          padding: 2px;
          border: 2px solid var(--white);
        }

        & .heading-3 {
          text-align: center;
          color: var(--black-hover);
        }
      }
    }
  }
}
</style>
