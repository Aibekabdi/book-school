<template>
  <div class="error-container">
    <span class="material-symbols-outlined"> report </span>
    <span class="msg">{{ t(currentLocalization, message) }}</span>
  </div>
</template>
<script setup>
import { ref, onMounted } from 'vue';
import { t } from '@/utils/i18n.js';
import { currentLocalization } from '@/App.vue';
const props = defineProps(['message']);
</script>
<style scoped lang="scss">
.error-container {
  width: fit-content !important;
  display: flex;
  flex-direction: row;
  gap: 40px;
  justify-content: space-between;
  align-items: center;
  position: fixed;
  top: 20px;
  z-index: 150;
  min-width: 400px;
  padding: 20px 30px;
  border-radius: 12px;
  background-color: var(--danger);
  box-shadow: 0 0 40px var(--danger-hover);

  color: var(--white);
  & .material-symbols-outlined {
    font-size: 36px;
  }
  & .msg {
    font-family: 'Montserrat', sans-serif;
    font-weight: 700;
  }
}
</style>
