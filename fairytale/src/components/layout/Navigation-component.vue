<template>
  <nav class="nav">
    <div class="max-w nav-wrapper">
      <div class="nav-left">
        <!-- <h1><a href="/">Fairytale Platform</a></h1> -->
        <a href="/" class="nav-logo">
          <div class="ellipse">
            <img src="../../assets/nav-logo.png" alt="" />
          </div>
        </a>
      </div>
      <div class="nav-center">
        <ul class="nav-list">
          <li v-for="link in links" :key="link.id" class="list-item">
            <a class="link" :href="link.url">{{ t(currentLocalization, link.name) }}</a>
          </li>
        </ul>
      </div>
      <div class="btn-container">
        <div class="choose-lang">
          <span
            :class="currentLocalization === 'kz' ? 'selected' : ''"
            @click="toggleLocalizationKz()"
            >🇰🇿 KZ</span
          >
          <span
            :class="currentLocalization === 'ru' ? 'selected' : ''"
            @click="toggleLocalizationRu()"
            >🇷🇺 RU</span
          >
        </div>
        <Button
          v-if="!user"
          type="Router"
          link="/login"
          :label="t(currentLocalization, 'SIGN_IN')"
          color="white"
          :rounded="true"
          icon="login"
        />
        <Button
          v-else
          type="Button"
          @click="logout"
          :label="t(currentLocalization, 'SIGN_OUT')"
          color="white"
          :rounded="true"
          icon="door_open"
        />
      </div>
      <div class="nav-mobile">
        <Button
          v-if="open"
          @click="toggleMenu"
          color="white"
          :rounded="true"
          icon="close"
          :outlined="true"
        />
        <Button
          v-else
          @click="toggleMenu"
          color="white"
          :rounded="true"
          icon="menu"
          :outlined="true"
        />
      </div>
    </div>
    <div class="mobile" v-if="open" @click="toggleMenu">
      <div>
        <ul class="nav-list">
          <li v-for="link in links" :key="link.id" class="list-item">
            <a class="link" :href="link.url">{{ t(currentLocalization, link.name) }}</a>
          </li>
        </ul>
        <div class="grid">
          <div class="row choose-lang">
            <span
              :class="currentLocalization === 'kz' ? 'selected' : ''"
              @click="toggleLocalizationKz()"
              >🇰🇿 KZ</span
            >
            <span
              :class="currentLocalization === 'ru' ? 'selected' : ''"
              @click="toggleLocalizationRu()"
              >🇷🇺 RU</span
            >
          </div>
          <Button
            v-if="!user"
            type="Router"
            link="/login"
            :label="t(currentLocalization, 'SIGN_IN')"
            color="white"
            :rounded="true"
            icon="login"
          />
          <Button
            v-else
            type="Button"
            @click="logout"
            :label="t(currentLocalization, 'SIGN_OUT')"
            color="white"
            :rounded="true"
            icon="door_open"
          />
        </div>
      </div>
    </div>
  </nav>
</template>

<script setup>
import { user } from '@/App.vue';
import { ref, defineComponent } from 'vue';
import { currentLocalization } from '@/App.vue';
import Button from '@/components/Button-component.vue';
import { t } from '@/utils/i18n.js';
</script>

<script>
const open = ref(false);

export default defineComponent({
  methods: {
    toggleMenu() {
      open.value = !open.value;
    },
    logout() {
      user.value = null;
      localStorage.removeItem('admin');
      localStorage.removeItem('userType');
      localStorage.removeItem('userToken');
      this.$router.push({ path: '/' });
    },
    toggleLocalizationRu() {
      if (currentLocalization.value == 'ru') {
        return;
      } else {
        currentLocalization.value = 'ru';
        localStorage.setItem('localization', 'ru');
      }
    },
    toggleLocalizationKz() {
      if (currentLocalization.value == 'kz') {
        return;
      } else {
        currentLocalization.value = 'kz';
        localStorage.setItem('localization', 'kz');
      }
    },
  },
  data() {
    return {
      links: [
        { id: 1, name: 'MAIN', url: '/' },
        { id: 2, name: 'ABOUT_US', url: '/about' },
        { id: 3, name: 'REGISTRATION', url: '/register' },
      ],
    };
  },
});
</script>

<style scoped lang="scss">
.nav-logo {
  width: 110px;
  height: 80px;
  display: inline-block;
}
.ellipse {
  background: #ffffff;
  padding: 7px 15px;
  border-radius: 50%;
}
.nav-logo img {
  width: 100%;
  height: 100%;
  object-fit: contain;
}

.nav {
  position: fixed;
  width: 100%;
  display: flex;
  justify-content: center;
  color: var(--white);
  // background: var(--success);
  background: linear-gradient(0deg, rgba(0, 0, 0, 0.2), rgba(0, 0, 0, 0.2)),
    rgba(255, 255, 255, 0.12);
  backdrop-filter: blur(15px);
  // height: 60px;
  z-index: 100;
  // box-shadow: 0 0 40px 20px rgba(105, 121, 248, 0.7);
}

.nav-wrapper {
  display: flex;
  flex-flow: row nowrap;
  align-items: center;
  justify-content: space-between;
  padding: 6px 40px;
  height: 100%;

  & .btn-container {
    flex-wrap: nowrap;
    align-items: center;
  }
}

h1 {
  display: inline-block;
  margin: 0;
}

.nav-list {
  display: flex;
  list-style: none;
}

.list-item {
  margin: 0 20px;
}

.nav-center {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.choose-lang {
  display: flex;
  align-content: center;
  justify-content: center;
  gap: 20px;
  padding: 0 40px;

  & span {
    cursor: pointer;
    font-family: 'Montserrat', sans-serif;
    font-weight: 700;
    color: var(--white-hover);
    transition: all 0.2s;
  }

  & span:hover {
    opacity: 0.8;
  }

  & span.selected {
    color: var(--white);
    border-bottom: 2px solid var(--white-hover);
  }
}

.nav-mobile {
  display: none;
}

.mobile {
  display: none;
}

@media only screen and (max-width: 1200px) {
  .nav-logo {
    width: 85px;
    height: 60px;
  }
  .nav-wrapper > .nav-center {
    display: none;
  }

  .nav-wrapper > .btn-container {
    display: none;
  }

  .nav-mobile {
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .mobile {
    display: block;
    position: fixed;
    top: 72px;
    left: 0;
    right: 0;
    bottom: 0;
    height: calc(100vh - 72px);
    background-color: rgba(0, 0, 0, 0.4);

    & > div {
      position: fixed;
      top: 72px;
      right: 0;
      width: 300px;
      display: flex;
      flex-direction: column;
      background-color: var(--black-hover);
      border-radius: 20px 0 20px 20px;
    }

    & .nav-list {
      list-style-type: disc;
      display: flex;
      flex-direction: column;
      justify-content: space-evenly;
      padding: 40px 40px 20px 40px;

      & .list-item {
        line-height: 2;
      }
    }

    & .choose-lang {
      padding: 0;
    }

    & .grid {
      gap: 20px;
      padding: 0 40px 40px 40px;
    }
  }
}

@media only screen and (max-width: 768px) {
  h1 > a {
    font-size: 16px;
    height: fit-content;
    display: block;
  }
}
</style>
