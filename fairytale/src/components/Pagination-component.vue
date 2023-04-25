<template>
  <div class="pagination row centerize">
    <Button
      v-if="current > 0"
      type="Button"
      @click="prevPage"
      color="black"
      :label="t(currentLocalization, 'PREV')"
      icon="chevron_left"
      :rounded="true"
    />
    <Button
      v-else
      type="Button"
      color="black"
      :label="t(currentLocalization, 'PREV')"
      icon="chevron_left"
      :rounded="true"
      :disabled="true"
    />
    <span class="text-6"
      >{{ t(currentLocalization, 'PAGE') }} {{ current + 1 }} /
      {{ this.pages === 0 ? 1 : this.pages }}</span
    >
    <!-- <Button
              type="Button"
              @click="prevPage"
              color="black"
              label="&nbsp 1 &nbsp"
              :rounded="true"
            />
            <Button
              type="Button"
              @click="prevPage"
              color="black"
              label="&nbsp 2 &nbsp"
              :rounded="true"
            />
            <Button
              type="Button"
              @click="prevPage"
              color="black"
              label="&nbsp 3 &nbsp"
              :rounded="true"
            /> -->
    <Button
      v-if="current + 1 < this.pages"
      type="Button"
      @click="nextPage"
      color="black"
      :label="t(currentLocalization, 'NEXT')"
      icon="chevron_right"
      :rounded="true"
    />
    <Button
      v-else
      type="Button"
      color="black"
      :label="t(currentLocalization, 'NEXT')"
      icon="chevron_right"
      :rounded="true"
      :disabled="true"
    />
  </div>
</template>
<script setup>
import { ref, defineProps } from 'vue';
import Button from '@/components/Button-component.vue';
import { currentLocalization } from '@/App.vue';
import { t } from '@/utils/i18n.js';
defineProps({
  pages: {
    type: Number,
    required: true,
  },
});
</script>
<script>
const current = ref(0);
export default {
  components: { Button },
  mounted() {
    current.value = 0;
  },
  methods: {
    prevPage() {
      if (current.value > 0) {
        current.value--;
        this.$emit('any-change', current.value);
      } else {
        current.value = 1;
      }
      console.log(current.value);
    },
    nextPage() {
      if (current.value < this.pages) {
        current.value++;
        this.$emit('any-change', current.value);
      } else {
        current.value = this.pages;
      }
      console.log(this.pages);
    },
  },
};
</script>
<style scoped lang="scss">
.pagination {
  align-items: center;
}
</style>
