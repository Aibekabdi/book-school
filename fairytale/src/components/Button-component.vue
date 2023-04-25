<template>
  <button
    v-if="type == 'Button'"
    :class="[
      'btn',
      `btn_${color}`,
      { btn_rounded: rounded },
      { btn_outlined: outlined },
      { btn_large: size === 'large' },
    ]"
    :disabled="disabled"
    @click="clickOnButton"
    :style="{ width: width }"
  >
    <span v-if="icon && !label" class="material-symbols-outlined no-mr">{{ icon }}</span>
    <span v-if="icon && label" class="material-symbols-outlined">{{ icon }}</span>
    <span v-if="label" class="label">{{ label }}</span>
  </button>
  <a
    v-if="type == 'Link'"
    :class="[
      'btn',
      `btn_${color}`,
      { btn_rounded: rounded },
      { btn_outlined: outlined },
      { btn_large: size === 'large' },
    ]"
    :disabled="disabled"
    :href="`${link}`"
    :style="{ width: width }"
  >
    <span v-if="icon && !label" class="material-symbols-outlined no-mr">{{ icon }}</span>
    <span v-if="icon && label" class="material-symbols-outlined">{{ icon }}</span>
    <span v-if="label" class="label">{{ label }}</span>
  </a>
  <router-link
    v-if="type == 'Router'"
    :class="[
      'btn',
      `btn_${color}`,
      { btn_rounded: rounded },
      { btn_outlined: outlined },
      { btn_large: size === 'large' },
    ]"
    :disabled="disabled"
    :to="`${link}`"
    :style="{ width: width }"
  >
    <span v-if="icon && !label" class="material-symbols-outlined no-mr">{{ icon }}</span>
    <span v-if="icon && label" class="material-symbols-outlined">{{ icon }}</span>
    <span v-if="label" class="label">{{ label }}</span></router-link
  >
</template>
<script setup>
defineProps({
  type: {
    type: String,
    default: 'Button',
  },
  link: {
    type: String,
    required: false,
  },
  label: {
    type: String,
    default: '',
    required: false,
  },
  color: {
    type: String,
    default: 'primary',
  },
  disabled: {
    type: Boolean,
    required: false,
  },
  rounded: {
    type: Boolean,
    required: false,
  },
  outlined: {
    type: Boolean,
    required: false,
  },
  icon: {
    type: String,
    required: false,
  },
  size: {
    type: String,
    default: 'normal',
  },
  width: {
    type: String,
    required: false,
  },
});
const emit = defineEmits(['click']);
const clickOnButton = () => {
  emit('click');
};
</script>
<style lang="scss" scoped>
.btn {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 13px 23px;
  height: 40px;
  // color: #fff;
  letter-spacing: 0.02em;
  font-weight: 500;
  text-transform: uppercase;

  border-radius: 12px;
  border: none;
  cursor: pointer;
  font-size: 14px;
  transition: 0.4s;
  box-shadow: 1.08px 0.76px 2.24px rgba(0, 0, 0, 0.56);
  &_primary {
    background-color: var(--primary);
    // border: 1px solid var(--primary);
    &:hover {
      // background-color: var(--primary-hover);

      background: radial-gradient(
            96.87% 124.51% at 45.33% 23.44%,
            rgba(255, 255, 255, 0.49) 11.84%,
            rgba(255, 224, 224, 0) 63.33%,
            rgba(255, 224, 224, 0) 100%
          )
          /* warning: gradient uses a rotation that is not supported by CSS and may not behave as expected */,
        var(--primary-hover);
      box-shadow: 1.08px 0.76px 2.24px rgba(0, 0, 0, 0.56);
    }
  }
  &_second {
    background-color: var(--second);
    // border: 1px solid var(--second);
    &:hover {
      // background-color: var(--second-hover);

      background: radial-gradient(
            96.87% 124.51% at 45.33% 23.44%,
            rgba(255, 255, 255, 0.49) 11.84%,
            rgba(255, 224, 224, 0) 63.33%,
            rgba(255, 224, 224, 0) 100%
          )
          /* warning: gradient uses a rotation that is not supported by CSS and may not behave as expected */,
        var(--second-hover);
      box-shadow: 1.08px 0.76px 2.24px rgba(0, 0, 0, 0.56);
    }
  }
  &_success {
    background-color: var(--success);
    // border: 1px solid var(--success);
    &:hover {
      // background-color: var(--success-hover);

      background: radial-gradient(
            96.87% 124.51% at 45.33% 23.44%,
            rgba(255, 255, 255, 0.49) 11.84%,
            rgba(255, 224, 224, 0) 63.33%,
            rgba(255, 224, 224, 0) 100%
          )
          /* warning: gradient uses a rotation that is not supported by CSS and may not behave as expected */,
        #61deb9;
      box-shadow: 1.08px 0.76px 2.24px rgba(0, 0, 0, 0.56);
    }
  }
  &_info {
    background-color: var(--info);
    // border: 1px solid var(--info);
    &:hover {
      // background-color: var(--info-hover);

      background: radial-gradient(
            96.87% 124.51% at 45.33% 23.44%,
            rgba(255, 255, 255, 0.49) 11.84%,
            rgba(255, 224, 224, 0) 63.33%,
            rgba(255, 224, 224, 0) 100%
          )
          /* warning: gradient uses a rotation that is not supported by CSS and may not behave as expected */,
        var(--info-hover);
      box-shadow: 1.08px 0.76px 2.24px rgba(0, 0, 0, 0.56);
    }
  }
  &_warning {
    background-color: var(--warning);
    // border: 1px solid var(--warning);
    &:hover {
      // background-color: var(--warning-hover);

      background: radial-gradient(
            96.87% 124.51% at 45.33% 23.44%,
            rgba(255, 255, 255, 0.49) 11.84%,
            rgba(255, 224, 224, 0) 63.33%,
            rgba(255, 224, 224, 0) 100%
          )
          /* warning: gradient uses a rotation that is not supported by CSS and may not behave as expected */,
        var(--warning-hover);
      box-shadow: 1.08px 0.76px 2.24px rgba(0, 0, 0, 0.56);
    }
  }
  &_danger {
    background-color: var(--danger);
    // border: 1px solid var(--danger);
    &:hover {
      // background-color: var(--danger-hover);

      background: radial-gradient(
            96.87% 124.51% at 45.33% 23.44%,
            rgba(255, 255, 255, 0.49) 11.84%,
            rgba(255, 224, 224, 0) 63.33%,
            rgba(255, 224, 224, 0) 100%
          )
          /* warning: gradient uses a rotation that is not supported by CSS and may not behave as expected */,
        var(--danger-hover);
      box-shadow: 1.08px 0.76px 2.24px rgba(0, 0, 0, 0.56);
    }
  }
  &_yellow {
    color: var(--black);
    background-color: var(--yellow);
    // border: 1px solid var(--yellow);
    &:hover {
      // background-color: var(--yellow-hover);

      background: radial-gradient(
            96.87% 124.51% at 45.33% 23.44%,
            rgba(255, 255, 255, 0.49) 11.84%,
            rgba(255, 224, 224, 0) 63.33%,
            rgba(255, 224, 224, 0) 100%
          )
          /* warning: gradient uses a rotation that is not supported by CSS and may not behave as expected */,
        var(--yellow-hover);
      box-shadow: 1.08px 0.76px 2.24px rgba(0, 0, 0, 0.56);
    }
  }
  &_black {
    color: var(--white);
    background-color: var(--black);
    // border: 1px solid var(--black);
    &:hover {
      // background-color: var(--black-hover);

      background: radial-gradient(
            96.87% 124.51% at 45.33% 23.44%,
            rgba(255, 255, 255, 0.49) 11.84%,
            rgba(255, 224, 224, 0) 63.33%,
            rgba(255, 224, 224, 0) 100%
          )
          /* warning: gradient uses a rotation that is not supported by CSS and may not behave as expected */,
        var(--black-hover);
      box-shadow: 1.08px 0.76px 2.24px rgba(0, 0, 0, 0.56);
    }
  }
  &_white {
    color: var(--black);
    background-color: var(--white);
    // border: 1px solid var(--white);
    &:hover {
      // background-color: var(--white-hover);

      background: radial-gradient(
            96.87% 124.51% at 45.33% 23.44%,
            rgba(255, 255, 255, 0.49) 11.84%,
            rgba(255, 224, 224, 0) 63.33%,
            rgba(255, 224, 224, 0) 100%
          )
          /* warning: gradient uses a rotation that is not supported by CSS and may not behave as expected */,
        var(--white-hover);
      box-shadow: 1.08px 0.76px 2.24px rgba(0, 0, 0, 0.56);
    }
  }
  &:disabled {
    opacity: 0.6;
    cursor: not-allowed;
  }
  &_rounded {
    border-radius: 40px;
  }
  &_outlined {
    background-color: transparent;
    color: #000;
    &:hover {
      color: #fff;
    }
  }
  &_small {
    width: 40px;
    height: 40px;
    padding: 8px;
  }
  &_large {
    border-width: 2px;
    font-weight: 600;
    font-size: 16px;
    height: 48px;
    padding: 0 24px;
  }
}

.material-symbols-outlined {
  width: 24px;
  height: 24px;
  margin-right: 6px;
}

.material-symbols-outlined.no-mr {
  margin-right: 0;
}

span {
  height: 24px;
}

.label {
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
