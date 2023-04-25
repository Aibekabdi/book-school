<template>
  <Notification class="notification" v-if="onError" :message="onError" />
  <h3 class="heading-3">Student View</h3>
  <div class="grid student">
    <nav class="row justify">
      <h5 class="text-3">
        {{ studentData.login }} {{ studentData.firstName }} {{ studentData.secondName }}
      </h5>
      <h6 class="text-3" v-if="myStats != null">
        {{ myStats.total_points }} {{ t(currentLocalization, 'POINTS') }}
      </h6>
    </nav>
    <nav class="row justify">
      <div class="row">
        <Button
          type="Button"
          @click="toggleModal('CheckStatistics')"
          :label="t(currentLocalization, 'YOUR_STATS')"
          icon="show_chart"
          color="second"
        />
        <Button
          type="Button"
          @click="toggleModal('NotificationsList')"
          label="Уведомления"
          icon="notifications"
          color="second"
        />
      </div>
      <div class="row">
        <Button
          type="Router"
          link="/shop"
          :label="t(currentLocalization, 'SHOP')"
          icon="sentiment_satisfied"
          color="primary"
        />
      </div>
    </nav>
    <div class="grid-of-4">
      <a href="#feed" class="grid-square">
        <img width="130" height="130" src="@/assets/student/books.png" alt="" />
        <span v-if="myStats != null"> {{ myStats.book_points }}</span>
      </a>
      <a href="#feed" class="grid-square">
        <img width="130" height="130" src="@/assets/student/headphones.png" alt="" />
        <span v-if="myStats != null"> {{ myStats.audio_points }}</span>
      </a>
      <a href="#feed" class="grid-square">
        <img width="130" height="130" src="@/assets/student/test.png" alt="" />
        <span v-if="myStats != null"> {{ myStats.test_points }}</span>
      </a>
      <a href="#feed" class="grid-square">
        <img width="130" height="130" src="@/assets/student/paint.png" alt="" />
        <span v-if="myStats != null"> {{ myStats.creative_task_points }}</span>
      </a>
      <a href="#feed" class="grid-square">
        <img width="130" height="130" src="@/assets/student/open_question.png" alt="" />
        <span v-if="myStats != null"> {{ myStats.open_points }}</span>
      </a>
    </div>
    <div class="background-wave">
      <img class="bg-inner" src="@/assets/student/background-wave.png" alt="" />
      <h3 class="text-2">Геймификация</h3>
      <h6 class="text-4 points">ЗАРАБОТАЙ БАЛЛЫ</h6>
      <h6 class="text-4 character">СОЗДАЙ СВОЕГО ПЕРСОНАЖ</h6>
      <Button type="Link" link="/shop" size="large" label="ИГРАТЬ" color="success" />
      <img class="controller" src="@/assets/student/controller.png" alt="" />
      <img class="boards" src="@/assets/student/boards.png" alt="" />
      <img class="red-controller" src="@/assets/student/red-controller.png" alt="" />
      <img class="books-list" src="@/assets/student/books-list.png" alt="" />
      <img class="kitaplab" src="@/assets/student/plane.png" alt="" />
      <img class="new-robot" src="@/assets/aim-robot.png" alt="" />
    </div>
    <div id="feed">
      <h4 class="text-2">{{ t(currentLocalization, 'YOUR_BOOKS') }}</h4>
      <Feed class="feed" />
    </div>
    <div class="modal" v-if="open">
      <div>
        <div class="close">
          <Button type="Button" @click="toggleModal()" icon="close" :rounded="true" color="black" />
        </div>
        <div class="modal-content" v-if="modalType === 'CheckStatistics'">
          <MyStatistics />
        </div>
        <div class="modal-content" v-else-if="modalType === 'NotificationsList'">
          <NotificationsList />
        </div>
      </div>
    </div>
  </div>
</template>
<script setup>
import Button from '@/components/Button-component.vue';
import Feed from '@/components/Feed-component.vue';
import MyStatistics from '@/components/MyStatistics-component.vue';
import Notification from '@/components/layout/Notification-component.vue';
import NotificationsList from '@/components/NotificationsList-component.vue';
import { sendRequest } from '@/utils/utils';
import { user, userToken, onError, currentLocalization } from '@/App.vue';
import { ref, reactive } from 'vue';
import { t } from '@/utils/i18n.js';
</script>
<script>
const isSuc = ref(false);
const isErr = ref(false);

const open = ref(false);
const modalType = ref('');

const myStats = ref(null);

const studentData = reactive({
  firstName: '',
  secondName: '',
  login: '',
});
export default {
  mounted() {
    user.value = localStorage.getItem('admin') ? 'admin' : localStorage.getItem('userType') || '';
    userToken.value = localStorage.getItem('admin') || localStorage.getItem('userToken') || '';

    this.getStudentData();
    this.getStudentStats();
  },
  methods: {
    toggleModal(modal) {
      isErr.value = false;
      isSuc.value = false;
      modalType.value = modal;
      open.value = !open.value;

      if (!open.value) {
        modalType.value = '';
      }
    },

    // student view
    getStudentData() {
      sendRequest('/api/student/profile', 'GET', null, userToken.value)
        .then((data) => {
          console.log(data);
          if (data) {
            studentData.login = data.username;
            studentData.firstName = data.first_name;
            studentData.secondName = data.second_name;
          }
        })
        .catch((err) => {
          onError.value = 'Невозможно получить данные пользователя, пожалуйста повторите позднее';
          console.log(err);

          setTimeout(() => {
            onError.value = null;
          }, 5000);
        });
    },

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
  components: {
    Button,
    Feed,
    MyStatistics,
  },
};
</script>
<style scoped lang="scss">
.student {
  position: relative;
}

.feed {
  margin-top: 40px;
}

#feed {
  & .text-2 {
    text-align: center;
  }
}

.grid-of-4 {
  display: flex;
  padding: 50px;
  display: flex;
  flex-flow: row nowrap;
  align-items: center;
  justify-content: center;
  gap: 50px;
  width: 100%;

  & .grid-square {
    padding: 40px;
    border-radius: 10px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;

    &:first-child {
      background-color: #4285ff;
    }

    &:nth-child(2) {
      background-color: #76ff40;
    }

    &:nth-child(3) {
      background-color: #ec40ff;
    }

    &:nth-last-child(2) {
      background-color: #ff4440;
    }

    &:last-child {
      background-color: #ffed42;
    }

    & span {
      font-size: 24px;
      font-weight: 700;
      margin-top: 30px;
      color: var(--black);
    }
  }
}

.background-wave {
  position: relative;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 40px;
  height: 500px;
  margin: 100px 0;

  & .text-2 {
    margin-top: 80px;
  }

  & .bg-inner {
    position: absolute;
    width: calc(300vw);
    height: 500px;
    background-size: 100% 503px;
    z-index: -2;
  }

  & .points {
    font-weight: 700;
    position: absolute;
    left: 240px;
    top: 120px;
  }

  & .character {
    font-weight: 700;
    position: absolute;
    left: 302px;
    top: 172px;
  }

  & .new-robot {
    width: 320px;
    height: 400px;
    position: absolute;
    top: 100px;
    right: 0px;
  }

  & .controller {
    position: absolute;
    left: 50%;
    top: 240px;
    z-index: -1;

    transform: translate(-50%, -6%);
  }

  & .boards {
    position: absolute;

    right: 320px;
    top: 80px;
  }

  & .red-controller {
    position: absolute;

    left: 30px;
    bottom: 320px;
  }

  & .books-list {
    position: absolute;

    left: 0;
    bottom: 0;
  }

  & .kitaplab {
    position: absolute;

    left: 100%;
    bottom: 0;
  }
}
</style>
