<template>
  <Notification class="notification" v-if="onError" :message="onError" />
  <div class="dashboard centerize" v-if="leaderboard">
    <div v-if="user === 'teacher'">
      <h3 class="heading-2">{{ t(currentLocalization, 'LEADERBOARD') }}</h3>
      <div class="table">
        <div class="row justify table-data">
          <div class="row">
            <span class="row-center"
              ><span class="material-symbols-outlined"> school </span>
              {{ t(currentLocalization, 'CLASS') }}</span
            >
          </div>
          <div class="row">
            <span class="row-center"
              ><span class="material-symbols-outlined"> account_circle </span>
              {{ t(currentLocalization, 'LASTNAME') }} {{ t(currentLocalization, 'NAME') }}</span
            >
          </div>
          <div class="row">
            <span class="row-center"
              ><span class="material-symbols-outlined"> book </span>
              {{ t(currentLocalization, 'BOOKS') }}</span
            >
          </div>
          <div class="row">
            <span class="row-center"
              ><span class="material-symbols-outlined"> volume_up </span>
              {{ t(currentLocalization, 'AUDIO') }}</span
            >
          </div>
          <div class="row">
            <span class="row-center"
              ><span class="material-symbols-outlined"> emoji_objects </span>
              {{ t(currentLocalization, 'CREATIVE_TASKS') }}</span
            >
          </div>
          <div class="row">
            <span class="row-center"
              ><span class="material-symbols-outlined"> quiz </span>
              {{ t(currentLocalization, 'TEST_TASKS') }}</span
            >
          </div>
          <div class="row">
            <span class="row-center"
              ><span class="material-symbols-outlined"> check_circle </span>
              {{ t(currentLocalization, 'SUMMARY') }}</span
            >
          </div>
        </div>
        <div class="grid" v-for="student in leaderboard.slice(0, 5)" :key="student">
          <div class="grid">
            <div class="row justify table-data">
              <span>{{ student.grade }} {{ student.name }}</span>
              <span>{{ student.student_second_name }} {{ student.student_first_name }}</span>
              <span>{{ student.book_points }}</span>
              <span>{{ student.audio_points }}</span>
              <span>{{ student.creative_task_points }}</span>
              <span>{{ student.test_points }}</span>
              <span>{{ student.total_points }}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div v-else></div>
</template>
<script setup>
import Notification from '@/components/layout/Notification-component.vue';
import { ref } from 'vue';
import { user, userToken, onError, currentLocalization } from '@/App.vue';
import { sendRequest } from '@/utils/utils';
import { t } from '@/utils/i18n.js';
</script>
<script>
const leaderboard = ref([]);

export default {
  mounted() {
    user.value = localStorage.getItem('admin') ? 'admin' : localStorage.getItem('userType') || '';
    userToken.value = localStorage.getItem('admin') || localStorage.getItem('userToken') || '';
    if (user.value === 'teacher') {
      this.getTeacherLeaderBoard();
    }
  },
  methods: {
    getTeacherLeaderBoard() {
      sendRequest('/api/class/stats/total', 'GET', null, userToken.value)
        .then((data) => {
          if (data) {
            console.log(data);
            leaderboard.value = data;
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
  },
};
</script>
<style scoped lang="scss">
.row-center {
  display: flex;
  align-self: center;
  gap: 4px;
}

.table-data {
  & > span {
    width: 100%;
  }
}
</style>
