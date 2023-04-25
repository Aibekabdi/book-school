<template>
  <Notification class="notification" v-if="onError" :message="onError" />
  <div class="shop max-w">
    <nav class="row justify">
      <div class="row">
        <Button
          type="Router"
          link="/home"
          :label="t(currentLocalization, 'HOME')"
          icon="home"
          color="primary"
        />
      </div>
    </nav>
    <div class="row">
      <h2 class="heading-2">{{ t(currentLocalization, 'SHOP') }}</h2>
      <h2 class="heading-2" v-if="myStats !== null">
        {{ myStats.total_points }} {{ t(currentLocalization, 'POINTS') }}
      </h2>
    </div>
    <div class="row shop-wrapper">
      <div class="background">
        <!-- <p>{{ newSet.head.image_url }}</p> -->
        <!-- <p>{{ newSet.chest.image_url }}</p> -->
        <!-- <p>{{ newSet.arms.image_url }}</p> -->
        <!-- <p>{{ newSet.legs.image_url }}</p> -->
        <img :src="newSet.head.image_url" />
        <img :src="newSet.chest.image_url" />
        <img :src="newSet.arms.image_url" />
        <img :src="newSet.legs.image_url" />
      </div>
      <div class="grid shopping-grid">
        <div class="row">
          <div v-for="item in heads">
            <div class="item-card">
              <div :key="item.id" class="img-container" @click="wearBody('head', item)">
                <img v-if="item.image_url" :src="item.image_icon_url" />
                <span v-else>{{ item.id }}</span>
                <div class="buyed" v-if="item.buyed">
                  <span class="material-symbols-outlined"> done </span>
                </div>
              </div>
              <p>{{ item.buyed ? 'Куплено' : item.price }}</p>
            </div>
          </div>
        </div>
        <div class="row">
          <div v-for="item in chest">
            <div class="item-card">
              <div :key="item.id" class="img-container" @click="wearBody('chest', item)">
                <img v-if="item.image_url" :src="item.image_icon_url" />
                <span v-else>{{ item.id }}</span>
                <div class="buyed" v-if="item.buyed">
                  <span class="material-symbols-outlined"> done </span>
                </div>
              </div>
              <p>{{ item.buyed ? 'Куплено' : item.price }}</p>
            </div>
          </div>
        </div>
        <div class="row">
          <div v-for="item in arms">
            <div class="item-card">
              <div :key="item.id" class="img-container" @click="wearBody('arms', item)">
                <img v-if="item.image_url" :src="item.image_icon_url" />
                <span v-else>{{ item.id }}</span>
                <div class="buyed" v-if="item.buyed">
                  <span class="material-symbols-outlined"> done </span>
                </div>
              </div>
              <p>{{ item.buyed ? 'Куплено' : item.price }}</p>
            </div>
          </div>
        </div>
        <div class="row">
          <div v-for="item in legs">
            <div class="item-card">
              <div :key="item.id" class="img-container" @click="wearBody('legs', item)">
                <img v-if="item.image_url" :src="item.image_icon_url" />
                <span v-else>{{ item.id }}</span>
                <div class="buyed" v-if="item.buyed">
                  <span class="material-symbols-outlined"> done </span>
                </div>
              </div>
              <p>{{ item.buyed ? 'Куплено' : item.price }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="buttons-container">
      <!-- <button v-if="!newSet.head.buyed" @click="buyBody('head', newSet.head.id)">
        Купить голову
      </button>
      <button v-if="!newSet.chest.buyed" @click="buyBody('chest', newSet.chest.id)">
        Купить тело
      </button>
      <button v-if="!newSet.arms.buyed" @click="buyBody('arms', newSet.arms.id)">
        Купить руки
      </button>
      <button v-if="!newSet.legs.buyed" @click="buyBody('legs', newSet.legs.id)">
        Купить ноги
      </button> -->
      <Button
        type="Button"
        v-if="!newSet.head.buyed"
        @click="buyBody('head', newSet.head.id)"
        :label="t(currentLocalization, 'BUY_HEAD')"
        color="success"
      />
      <Button
        type="Button"
        v-if="!newSet.chest.buyed"
        @click="buyBody('chest', newSet.chest.id)"
        :label="t(currentLocalization, 'BUY_CHEST')"
        color="success"
      />
      <Button
        type="Button"
        v-if="!newSet.arms.buyed"
        @click="buyBody('arms', newSet.arms.id)"
        :label="t(currentLocalization, 'BUY_ARMS')"
        color="success"
      />
      <Button
        type="Button"
        v-if="!newSet.legs.buyed"
        @click="buyBody('legs', newSet.legs.id)"
        :label="t(currentLocalization, 'BUY_LEGS')"
        color="success"
      />
    </div>

    <div class="buttons-container">
      <Button
        type="Button"
        @click="discardChanges()"
        :label="t(currentLocalization, 'DISCARD')"
        color="danger"
      />
      <Button
        type="Button"
        :disabled="!canApplyNewSet()"
        @click="applyNewSet()"
        :label="t(currentLocalization, 'APPLY')"
        color="warning"
      />
    </div>
  </div>
</template>
<script setup>
import Button from '@/components/Button-component.vue';
import { t } from '@/utils/i18n.js';
import { sendRequest } from '@/utils/utils.js';
import { onMounted, ref, reactive } from 'vue';
import { userToken, onError, currentLocalization } from '@/App.vue';
import Notification from '@/components/layout/Notification-component.vue';

const myStats = ref(null);

const currentSet = reactive({
  head: { id: 'head1', buyed: true },
  chest: { id: 'chest1', buyed: true },
  arms: { id: 'arms1', buyed: true },
  legs: { id: 'legs1', buyed: true },
});
const newSet = reactive({
  head: { id: 'head1', buyed: true },
  chest: { id: 'chest1', buyed: true },
  arms: { id: 'arms1', buyed: true },
  legs: { id: 'legs1', buyed: true },
});
const heads = ref([
  { id: 'head1', buyed: true, price: 0 },
  { id: 'head2', buyed: false, price: 0 },
  { id: 'head3', buyed: false, price: 0 },
  { id: 'head4', buyed: false, price: 0 },
  { id: 'head5', buyed: false, price: 0 },
  { id: 'head6', buyed: false, price: 0 },
  { id: 'head7', buyed: false, price: 0 },
  { id: 'head8', buyed: false, price: 0 },
  { id: 'head9', buyed: false, price: 0 },
  { id: 'head10', buyed: false, price: 0, price: 0 },
]);
const chest = ref([
  { id: 'chest1', buyed: true, price: 0 },
  { id: 'chest2', buyed: false, price: 0 },
  { id: 'chest3', buyed: false, price: 0 },
  { id: 'chest4', buyed: false, price: 0 },
  { id: 'chest5', buyed: false, price: 0 },
  { id: 'chest6', buyed: false, price: 0 },
  { id: 'chest7', buyed: false, price: 0 },
  { id: 'chest8', buyed: false, price: 0 },
  { id: 'chest9', buyed: false, price: 0 },
  { id: 'chest10', buyed: false, price: 0 },
]);
const arms = ref([
  { id: 'arms1', buyed: true, price: 0 },
  { id: 'arms2', buyed: false, price: 0 },
  { id: 'arms3', buyed: false, price: 0 },
  { id: 'arms4', buyed: false, price: 0 },
  { id: 'arms5', buyed: false, price: 0 },
  { id: 'arms6', buyed: false, price: 0 },
  { id: 'arms7', buyed: false, price: 0 },
  { id: 'arms8', buyed: false, price: 0 },
  { id: 'arms9', buyed: false, price: 0 },
  { id: 'arms10', buyed: false, price: 0 },
]);
const legs = ref([
  { id: 'legs1', buyed: true, price: 0 },
  { id: 'legs2', buyed: false, price: 0 },
  { id: 'legs3', buyed: false, price: 0 },
  { id: 'legs4', buyed: false, price: 0 },
  { id: 'legs5', buyed: false, price: 0 },
  { id: 'legs6', buyed: false, price: 0 },
  { id: 'legs7', buyed: false, price: 0 },
  { id: 'legs8', buyed: false, price: 0 },
  { id: 'legs9', buyed: false, price: 0 },
  { id: 'legs10', buyed: false, price: 0 },
]);
const wearBody = (slot, body) => {
  newSet[slot] = body;
};
const buyBody = (slot, id) => {
  sendRequest(`/api/shop/buy/${id}`, 'POST', null, localStorage.getItem('userToken'))
    .then(() => {
      newSet[slot].buyed = true;
    })
    .catch((e) => (onError.value = 'YABTPB'));
};
const applyNewSet = () => {
  Object.keys(newSet).forEach((slot) => {
    const oldId = currentSet[slot].id;
    const newId = newSet[slot].id;
    if (oldId !== newId) {
      updateBody(oldId, newId);
    }
  });
};
const discardChanges = () => {
  Object.keys(newSet).forEach((slot) => (newSet[slot] = currentSet[slot]));
};
const canApplyNewSet = () => {
  return (
    Object.keys(newSet).every((slot) => newSet[slot].buyed) &&
    Object.keys(newSet).some((slot) => currentSet[slot].id !== newSet[slot].id)
  );
};
const updateBody = (oldId, newId) => {
  sendRequest(
    `/api/student/body/update/${oldId}/${newId}`,
    'PATCH',
    null,
    localStorage.getItem('userToken'),
  ).catch((e) => (onError.value = e.msg));
};
const loadCurrentBody = () => {
  sendRequest('/api/student/body/current', 'GET', null, localStorage.getItem('userToken'))
    .then((e) => {
      currentSet.head = e.heads[0];
      currentSet.chest = e.chest[0];
      currentSet.arms = e.arms[0];
      currentSet.legs = e.legs[0];
      newSet.head = e.heads[0];
      newSet.chest = e.chest[0];
      newSet.arms = e.arms[0];
      newSet.legs = e.legs[0];
      console.log(e);
    })
    .catch((e) => console.log(e));
};
const getStudentStats = () => {
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
};
onMounted(() => {
  getStudentStats();
  sendRequest('/api/shop/all', 'GET', null, localStorage.getItem('userToken'))
    .then((e) => {
      heads.value = e.heads;
      chest.value = e.chest;
      arms.value = e.arms;
      legs.value = e.legs;
      console.log(e);
    })
    .catch((e) => console.log(e));
  loadCurrentBody();
});
</script>
<style scoped lang="scss">
.shop-wrapper {
  display: flex;
  flex-flow: row nowrap;
}

.img-container {
  position: relative;
  width: 64px;
  height: 64px;
  border: 1px solid var(--black-hover);
  cursor: pointer;

  transition: all 0.4s;

  & img {
    width: 100%;
    height: 100%;
  }

  & .buyed {
    position: absolute;
    right: 0;
    bottom: 0;

    width: 16px;
    height: 16px;
    z-index: -1;

    & > .material-symbols-outlined {
      color: var(--success);
      font-size: 16px;
    }
  }

  &:hover {
    border-radius: 8px;
    border: 1px solid var(--light-blue);
  }
}

.shopping-grid {
  justify-content: center;
  width: 600px;
  max-width: 600px;
  overflow: scroll;
  padding-bottom: 20px;
}

.background {
  width: 600px;
  min-width: 600px;
  height: 400px;
  border: 1px solid var(--black);
  position: relative;

  z-index: 1;

  background-color: var(--white-hover);
  // background: center center url('@/assets/background/back-1.PNG');

  & img {
    position: absolute;
    left: 50%;
    top: 50%;

    width: 216px;
    height: 373px;

    transform: translate(-50%, -50%);
  }

  & img:nth-child(1) {
    z-index: 9;
  }

  & img:nth-child(2) {
    z-index: 8;
  }

  & img:nth-child(3) {
    z-index: 7;
  }

  & img:nth-child(4) {
    z-index: 6;
  }
}

.buttons-container {
  margin-top: 40px;
  display: flex;
  flex-flow: row wrap;

  gap: 10px;
}

.item-card {
  display: flex;
  flex-direction: column;

  & p {
    text-align: center;
    margin-top: 4px;
    margin-bottom: 0;
  }
}

.grid {
  gap: 30px;
}

@media only screen and (max-width: 1440px) {
  .shop-wrapper {
    flex-direction: column;
  }

  .background {
    margin-bottom: 70px;
  }
}
</style>
