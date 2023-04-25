<template>
  <Notification class="notification" v-if="onError" :message="onError" />
  <div>
    <div class="grid">
      <h3 class="heading-2">{{ t(currentLocalization, 'STATS') }}</h3>
      <p class="text-4">{{ t(currentLocalization, 'HYCSSOYC') }}</p>
      <div class="grid">
        <div class="grid marginize" v-for="item in leaderboard" :key="classes">
          <div class="row">
            <span>{{ item.class.grade }} {{ item.class.name }}</span>
          </div>
          <div class="table">
            <div class="row justify">
              <span
                >{{ t(currentLocalization, 'LASTNAME') }} {{ t(currentLocalization, 'NAME') }}</span
              >
              <span>{{ t(currentLocalization, 'BOOK') }}</span>
              <span>{{ t(currentLocalization, 'AUDIO') }}</span>
              <span>{{ t(currentLocalization, 'CREATIVE_TASK') }}</span>
              <span>{{ t(currentLocalization, 'TEST_TASK') }}</span>
              <b>{{ t(currentLocalization, 'SA') }}</b>
            </div>
            <hr />
            <div class="row justify" v-for="student in item.stats" :key="student">
              <span>{{ student.student_second_name }} {{ student.student_first_name }}</span>
              <span>{{ student.book_points }}</span>
              <span>{{ student.audio_points }}</span>
              <span>{{ student.creative_task_points }}</span>
              <span>{{ student.test_points }}</span>
              <b>{{ student.total_points }}</b>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script setup>
import Notification from '@/components/layout/Notification-component.vue';
import { ref } from 'vue';
import { sendRequest } from '@/utils/utils';
import { userToken, onError, currentLocalization } from '@/App.vue';
import { t } from '@/utils/i18n.js';
</script>
<script>
const leaderboard = ref([]);

export default {
  mounted() {
    this.getStudentsWork();
  },
  methods: {
    getStudentsWork() {
      sendRequest('/api/class/stats', 'GET', null, userToken.value)
        .then((data) => {
          if (data) {
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
.table > .justify > span,
.table > .justify > span {
  width: 100%;
}

.marginize {
  margin-bottom: 40px;

  & > .row {
    color: var(--primary);
    font-size: 20px;
    font-weight: 700;
    margin-bottom: 10px;

    & > span {
      text-align: center;
      width: 100%;
    }
  }
}
</style>
