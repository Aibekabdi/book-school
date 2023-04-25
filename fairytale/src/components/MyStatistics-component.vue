<template>
  <Notification class="notification" v-if="onError" :message="onError" />
  <div v-if="user === 'student'">
    <div v-if="myStats">
      <h3 class="heading-2">{{ t(currentLocalization, 'YOUR_STATS') }}</h3>
      <h6 class="text-3">{{ myStats.student_second_name }} {{ myStats.student_first_name }}</h6>
      <div class="grid fixed-size">
        <span>{{ t(currentLocalization, 'SAOP') }}</span>
        <div class="card-container">
          <span class="points">{{ myStats.total_points }}</span>
        </div>
        <div class="card-border">
          <div class="row justify">
            <span>{{ t(currentLocalization, 'PFR') }}</span>
            <span>{{ myStats.book_points }}</span>
          </div>
          <div class="row justify">
            <span>{{ t(currentLocalization, 'PFL') }}</span>
            <span>{{ myStats.audio_points }}</span>
          </div>
          <div class="row justify">
            <span>{{ t(currentLocalization, 'PFCT') }}</span>
            <span>{{ myStats.creative_task_points }}</span>
          </div>
          <div class="row justify">
            <span>{{ t(currentLocalization, 'PFOT') }}</span>
            <span>{{ myStats.open_points }}</span>
          </div>
          <div class="row justify">
            <span>{{ t(currentLocalization, 'PFT') }}</span>
            <span>{{ myStats.test_points }}</span>
          </div>
        </div>
      </div>
    </div>
    <div v-else-if="myStats === null">
      <h3 class="heading-2">{{ t(currentLocalization, 'OOPS') }}...</h3>
      <h4 class="heading-3 error-message">{{ t(currentLocalization, 'REG_ERR') }}</h4>
    </div>
  </div>
</template>
<script setup>
import Notification from '@/components/layout/Notification-component.vue';
import { ref } from 'vue';
import { user, userToken, onError, currentLocalization } from '@/App.vue';
import { sendRequest } from '@/utils/utils';
import { t } from '@/utils/i18n.js';
</script>
<script>
const myStats = ref({});

export default {
  mounted() {
    user.value = localStorage.getItem('admin') ? 'admin' : localStorage.getItem('userType') || '';
    userToken.value = localStorage.getItem('admin') || localStorage.getItem('userToken') || '';

    this.getStudentStats();
  },
  methods: {
    getStudentStats() {
      sendRequest('/api/student/stats', 'GET', null, userToken.value)
        .then((data) => {
          if (data) {
            myStats.value = data;
            console.log(data);
          }
        })
        .catch((err) => {
          myStats.value = null;

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
.fixed-size {
  width: 600px;
  margin: 0 auto;
  font-family: 'Montserrat', sans-serif;
  & .card-container {
    width: fit-content;
    min-width: 200px;
    min-height: 79px;
    margin: 12px auto;
    padding: 20px;
    color: var(--light-blue);
    border-radius: 12px;
    background-color: var(--white);

    & .points {
      font-size: 32px;
      font-weight: 700;
    }
  }

  & .card-border {
    width: fit-content;
    min-width: 200px;
    margin: 12px auto;
    padding: 20px;
    color: var(--black-hover);
    border-radius: 12px;
    border: 4px solid var(--black);

    font-weight: 600;
    font-size: 20px;

    & > .justify {
      gap: 60px;
    }
  }
}
</style>
