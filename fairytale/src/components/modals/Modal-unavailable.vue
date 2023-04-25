<script setup>
import { currentLocalization } from '@/App.vue';
import { t } from '@/utils/i18n.js';
</script>
<template>
  <div class="modal-unavailable">
    <div>
      <span>📖</span>
      <h3 class="heading-3">{{ t(currentLocalization, 'DEVICE_UNAV') }}</h3>
      <p>{{ t(currentLocalization, 'PLS_GO_TO_PC') }}</p>
    </div>
  </div>
</template>
<style scoped lang="scss">
.modal-unavailable {
  display: none;
}

@media only screen and (max-width: 1200px) {
  .modal-unavailable {
    display: block;
    position: fixed;
    top: 0;
    left: 0;
    bottom: 0;
    right: 0;
    padding: 40px;
    background-color: rgba(0, 0, 0, 0.5);
    box-shadow: 0 0 50px 50px rgba(0, 0, 0, 0.8);
    z-index: 100;

    & div {
      width: 100%;
      height: 100%;
      text-align: center;
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
      background-color: var(--white-hover);
      border-radius: 20px;
      padding: 40px;

      & span {
        font-size: 82px;
      }
    }
  }
}
</style>
