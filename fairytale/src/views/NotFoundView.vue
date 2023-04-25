<template>
  <div class="error max-w">
    <h2 class="heading-1">404</h2>
    <h3 class="heading-2">{{ t(currentLocalization, '404') }}</h3>
    <div class="btn-container">
      <Button
        type="Link"
        link="/"
        :label="t(currentLocalization, 'GO_MAIN')"
        color="second"
        icon="home"
        :rounded="true"
        size="large"
      />
      <Button
        type="Router"
        link="/login"
        :label="t(currentLocalization, 'SIGN_IN')"
        color="black"
        icon="login"
        :rounded="true"
        size="large"
      />
    </div>
  </div>
</template>
<script setup>
import { currentLocalization } from '@/App.vue';
import { t } from '@/utils/i18n.js';
</script>
<script>
import Button from '@/components/Button-component.vue';
export default {
  name: 'NotFoundView',
  components: { Button },
};
</script>
<style scoped lang="scss">
.error {
  width: 500px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  text-align: center;
}

h2.heading-1 {
  font-size: 54px;
  margin: 0;
}

h3.heading-2 {
  font-size: 34px;
}

.btn-container {
  width: 100%;
  justify-content: center;
}
</style>
