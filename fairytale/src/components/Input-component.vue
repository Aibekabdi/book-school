<template>
  <div class="form-input" :style="{ width: width }">
    <input
      class="input-text"
      :type="type"
      :name="name"
      :id="name"
      :placeholder="placeholder"
      :value="value"
      @input="updateValue"
    />
    <!-- <label :for="name" class="input-label">{{ label }}</label> -->
    <TransitionGroup>
      <div class="form-error" v-for="element of error" :key="element.$uid">
        <div class="form-error__message">{{ element.$message }}</div>
      </div>
    </TransitionGroup>
  </div>
</template>

<script setup>
const emit = defineEmits(['update:value']);
defineProps({
  error: {
    type: Array,
    required: false,
  },
  value: {
    type: String,
    default: '',
  },
  name: {
    type: String,
    required: true,
  },
  type: {
    type: String,
    default: 'text',
  },
  placeholder: {
    type: String,
    required: true,
  },
  label: {
    type: String,
    required: true,
  },
  width: {
    type: String,
    default: '300px',
  },
});
const updateValue = (e) => {
  emit('update:value', e.target.value);
};
</script>

<style lang="scss" scoped>
.form {
  &-input {
    margin-bottom: 30px;
    position: relative;
  }
  &-error {
    background: var(--danger);
    margin-top: 8px;
    border-radius: 7px;
    font-size: 13px;
    color: #fff;
    padding: 5px;
    text-align: center;
  }
}
.input {
  &-text {
    font-weight: 700;
    font-size: 14px;
    letter-spacing: 0.02em;
    color: var(--black);
    width: 100%;
    border: none;
    border-bottom: 2px solid var(--black-hover);
    padding-bottom: 10px;
    background: none;
    outline: none;

    transition: all 0.4s;

    &:focus {
      color: var(--black-hover);
      border-bottom-color: var(--success);
    }
    // font-weight: 400;
    // font-family: 'Montserrat', sans-serif;
    // border: 1px solid var(--primary);
    // padding: 0 10px;
    // height: 40px;
    // border-radius: 7px;
    // font-size: 15px;
    // width: 100%;
    // position: relative;
    // z-index: 1;
    // &:focus {
    //   outline: 1px solid var(--second);
    //   & + .input-label {
    //     z-index: 1;
    //     opacity: 1;
    //     top: -20px;
    //   }
    // }
    // &:not(:placeholder-shown) {
    //   & + .input-label {
    //     z-index: 1;
    //     opacity: 1;
    //     top: -20px;
    //   }
    // }
  }
  // &-label {
  //   font-weight: bold;
  //   display: block;
  //   position: absolute;
  //   top: 0;
  //   left: 10px;
  //   z-index: -1;
  //   transition: 0.3s;
  //   font-size: 13px;
  //   color: var(--primary);
  // }

  &-text::placeholder {
    font-weight: bold;
    font-size: 14px;
    letter-spacing: 0.02em;
    color: var(--black-hover);
    opacity: 1;
  }
}
.v-enter-active,
.v-leave-active {
  transition: opacity 0.5s ease;
}
.v-enter-from,
.v-leave-to {
  opacity: 0;
}
</style>
