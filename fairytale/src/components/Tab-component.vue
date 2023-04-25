<template>
  <div class="tab-nav">
    <span
      v-for="tab in names"
      :key="tab.name"
      :class="['tab-nav__item', { selected: tab.name === selectedTab }]"
      @click="clickOnTab(tab.name)"
      >{{ tab.label }}
    </span>
  </div>
  <div class="tab-content" :style="{ width: width }">
    <slot />
  </div>
</template>
<script setup>
const props = defineProps({
  names: {
    type: Array,
    required: true,
  },
  selectedTab: {
    type: String,
    required: false,
  },
  width: {
    type: String,
    required: false,
    default: '100%',
  },
});
const emit = defineEmits(['changeTab']);
const clickOnTab = (tabName) => {
  emit('changeTab', tabName);
};
</script>
<style scoped lang="scss">
.tab {
  &-nav {
    display: flex;
    align-items: center;
    margin-bottom: 20px;

    &__item {
      height: 40px;
      display: flex;
      align-items: center;
      justify-content: center;
      margin-right: 10px;
      border-radius: 7px;
      cursor: pointer;
      border: 1px solid var(--success);
      padding: 0 20px;
      font-size: 15px;

      &:hover {
        background: radial-gradient(
              96.87% 124.51% at 45.33% 23.44%,
              rgba(255, 255, 255, 0.49) 11.84%,
              rgba(255, 224, 224, 0) 63.33%,
              rgba(255, 224, 224, 0) 100%
            )
            /* warning: gradient uses a rotation that is not supported by CSS and may not behave as expected */,
          var(--success-hover);
        box-shadow: 1.08px 0.76px 2.24px rgba(0, 0, 0, 0.56);
        color: #fff;
        transition: 0.2s;
      }
      &.selected {
        background: radial-gradient(
              96.87% 124.51% at 45.33% 23.44%,
              rgba(255, 255, 255, 0.49) 11.84%,
              rgba(255, 224, 224, 0) 63.33%,
              rgba(255, 224, 224, 0) 100%
            )
            /* warning: gradient uses a rotation that is not supported by CSS and may not behave as expected */,
          var(--success);
        box-shadow: 1.08px 0.76px 2.24px rgba(0, 0, 0, 0.56);
      }
    }
  }

  &-content {
    padding: 30px;
    border-radius: 12px;
  }
}
</style>
