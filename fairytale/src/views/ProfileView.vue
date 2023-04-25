<template>
  <div class="profile max-w">
    <h2 class="heading-2">{{ t(currentLocalization, 'YOUR_PROFILE') }}</h2>
    <div v-if="user === 'school'">
      <ProfileSchool />
    </div>
    <div v-else-if="user === 'teacher'">
      <ProfileTeacher />
    </div>
  </div>
</template>
<script setup>
import { user, userToken, currentLocalization } from '@/App.vue';
import ProfileSchool from './Subviews/Profile-School.vue';
import ProfileTeacher from './Subviews/Profile-Teacher.vue';
import { t } from '@/utils/i18n.js';
</script>
<script>
export default {
  name: 'ProfileView',
  components: { ProfileSchool, ProfileTeacher },
};
</script>
<!-- remove scoped for subviews -->
<style lang="scss">
.profile {
  & .grid > .grid {
    align-items: center;
  }

  & nav.row {
    width: 100%;
  }

  & section {
    width: 300px;
  }

  & .success-message,
  & .error-message {
    margin-top: 20px;
  }

  & .heading-3 {
    font-weight: normal;
    overflow-x: scroll;
    padding-bottom: 10px;
    scrollbar-width: thin;
    & span {
      font-weight: 700;
      color: var(--black-hover);
    }
  }
}
</style>
