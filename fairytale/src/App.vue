<script>
import Navigation from '@/components/layout/Navigation-component.vue';
import { ref } from 'vue';
export const user = ref(
  localStorage.getItem('admin') ? 'admin' : localStorage.getItem('userType') || '',
);
export const userToken = ref(
  localStorage.getItem('admin') || localStorage.getItem('userToken') || '',
);
export const onError = ref(null);
export const currentLocalization = ref(localStorage.getItem('localization') || 'ru');
export default {
  name: 'AppView',
  components: {
    Navigation,
  },
};
</script>

<template>
  <div class="container">
    <Navigation />
    <div class="content">
      <router-view />
    </div>
  </div>
</template>

<style lang="scss">
@import '@/styles/global.scss';

html {
  // smooth scroll:
  scroll-behavior: smooth;
}

body {
  margin: 0;
  padding: 0;
  overflow-x: hidden;
  background-color: inherit;
  background-color: #c5f2ef;
}

#app {
  font-family: 'Roboto Slab', serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

.sun {
  position: absolute;
  left: 80%;
  top: 0;

  width: 211px;
  height: 191px;

  z-index: -1;

  background: url('@/assets/Group\ 3245.png') no-repeat;
}

.robot {
  position: absolute;
  left: -20px;
  bottom: -60px;

  width: 338px;
  height: 393px;

  z-index: -1;

  background: url('@/assets/robot.png') no-repeat;
}

.change-back {
  background: bottom center url('@/assets/bg-g.png') no-repeat;
  background-size: 100% 70%;
}

h1,
h2,
h3,
h4,
h5,
h6,
a {
  font-family: 'Montserrat', sans-serif;
}

.container {
  display: flex;
  align-items: center;
  justify-content: space-between;
  flex-flow: column;
  min-height: 100vh;
}

.content {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  padding: 0 30px;
  margin: 120px 0 60px 0;
  transition: 0.2s;
  min-height: calc(100vh - 120px - 162px);
}

.grid {
  display: flex;
  flex-flow: column;
}

.row {
  display: flex;
  flex-wrap: nowrap;
  gap: 12px;
}

.max-w {
  width: 100%;
  max-width: 1440px;
}

.centerize {
  text-align: center;
  justify-content: center;
}

.justify {
  justify-content: space-between;
}

.line {
  width: 100%;
  height: 2px;
  margin: 20px 0;
  background: var(--black-hover);
  &_block {
    display: block;
  }
}

.pre-dash {
  & .heading-2 {
    text-align: left;
  }

  & > div {
    display: flex;
    flex-direction: column;
    gap: 20px;
  }

  & > div > .row {
    flex-wrap: wrap;
    align-items: center;
    justify-content: space-between;
  }
}

.dashboard {
  display: flex;
  flex-direction: column;
  gap: 20px;
  background-color: #84eee8;
  border-radius: 12px;
  width: 100%;
  min-height: 360px;
  padding: 20px;

  & .grid {
    width: 100%;
  }
}

.btn-container {
  width: fit-content;
  display: flex;
  flex-flow: row wrap;
  gap: 12px;
}

.success-message,
.error-message {
  font-weight: 700;
  text-align: center;
}

.success-message {
  color: var(--success);
}

.error-message {
  color: var(--danger);
}

.notification {
  position: fixed;
  top: 120px;
  left: 50%;
  transform: translateX(-50%);
}

.table {
  display: flex;
  flex-flow: column nowrap;
  border: 2px solid var(--black-hover);
  width: 100%;
  border-radius: 12px;
  padding: 20px;

  & .row {
    text-align: left;
    width: 100%;
    display: flex;
    flex-wrap: nowrap;
    justify-content: space-between;
    align-items: center;
  }

  & > .row {
    padding: 10px;
  }

  & .row.margin {
    padding: 10px;
  }

  .row.margin:nth-child(even) {
    background-color: #dddddd;
  }

  .row.margin:nth-child(odd) {
    background-color: #f1f1f1;
  }

  & .col {
    width: 100%;
    display: flex;
    flex-direction: column;
  }

  & .th-1 {
    font-weight: 600;
    font-size: 18px;
    border-bottom: 2px solid var(--black-hover);
  }

  & .th-2 {
    font-weight: 600;
  }
}

.modal {
  display: block;
  position: fixed;
  top: 0;
  left: 0;
  bottom: 0;
  right: 0;
  padding: 120px 40px 60px 40px;
  background-color: rgba(0, 0, 0, 0.2);
  box-shadow: 0 0 50px 50px rgba(0, 0, 0, 0.8);
  z-index: 20;

  & > div {
    overflow-x: hidden;
    overflow-y: auto;
    width: 100%;
    height: 100%;
    background-color: var(--light-blue);
    border-radius: 20px;
    padding: 40px;
  }

  & .modal-content {
    text-align: center;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 20px;
    justify-content: space-between;
    min-height: 100%;

    & > div {
      width: 100%;
    }
  }

  & .modal-content a {
    font-weight: 600;
    color: var(--primary);
    border-bottom: 1px solid var(--black-hover);
    transition: all 0.4s;
  }

  & .modal-content a:hover {
    opacity: 0.8;
  }

  & > div > .close {
    position: fixed;
    right: 60px;
    top: 140px;
  }

  & .dashboard {
    background-color: var(--white);
  }

  & .dashboard.not-found {
    flex-flow: column;
    gap: 20px;
    align-items: center;
    justify-content: center;

    & > span:first-child {
      font-size: 32px;
    }
  }
}

.vs__selected-options {
  word-break: break-all;
}

.vs__dropdown-menu {
  z-index: 5 !important;
}

hr {
  width: 100%;
  color: var(--black-hover);
  background-color: var(--black-hover);
}

nav.row.justify {
  margin-bottom: 20px;
}
</style>
