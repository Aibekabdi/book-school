<template>
  <div class="homepage max-w">
    <div class="sun"></div>
    <div class="robot"></div>
    <h2 class="heading-2">Домашная страница</h2>
    <div v-if="user === 'school'">
      <HomepageSchool />
    </div>
    <div v-else-if="user === 'teacher'">
      <HomepageTeacher />
    </div>
    <div v-else-if="user === 'student'">
      <HomepageStudent />
    </div>
    <ModalUnavailable />
  </div>
</template>
<script setup>
import HomepageSchool from './Subviews/Homepage-School.vue';
import HomepageTeacher from './Subviews/Homepage-Teacher.vue';
import HomepageStudent from './Subviews/Homepage-Student.vue';
import ModalUnavailable from '@/components/modals/Modal-unavailable.vue';
import { user } from '@/App.vue';
</script>
<script>
export default {
  name: 'HomepageView',

  mounted() {
    document.body.classList.add('change-back');
  },
  destroyed() {
    document.body.classList.remove('change-back');
  },
};
</script>

<style scoped lang="scss">
.homepage {
  position: relative;
}

.change-back {
  background-color: var(--light-blue);
}
</style>
