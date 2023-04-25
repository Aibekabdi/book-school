<template>
  <Notification class="notification" v-if="onError" :message="onError" />
  <div>
    <h3 class="heading-2">{{ t(currentLocalization, 'LATEST_MSGS') }}</h3>
    <hr />
    <div class="grid" v-for="notification in notifications" :key="notification">
      <h6 class="text-3">{{ notification.book_name }}</h6>
      <p class="text-4">{{ notification.question }}</p>
      <p class="text-4" v-if="!notification.is_art">{{ notification.answer }}</p>
      <img :src="notification.answer" />
      <p class="text-4">{{ t(currentLocalization, 'COMMENT') }}: {{ notification.comment }}</p>
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
const notifications = ref([]);
export default {
  mounted() {
    this.getStudentNotifications();
  },
  methods: {
    getStudentNotifications() {
      sendRequest('/api/creative/check/comment', 'GET', null, userToken.value)
        .then((data) => {
          if (data) {
            data.forEach((el, i) => {
              notifications.value[i] = el;
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
  },
};
</script>
<style scoped lang="scss"></style>
