<script setup lang="ts">
import { ToastTheme, type ToastProps } from '@/types/toast';
import { onMounted, onUnmounted, ref } from 'vue';

const props = withDefaults(defineProps<ToastProps>(), {
  duration: 5000,
  theme: ToastTheme.Success
});

const THEME_CLASSES = {
  [ToastTheme.Success]: 'success',
  [ToastTheme.Error]: 'error'
};

const visible = ref(true);

onMounted(() => {
  setTimeout(() => {
    visible.value = false;
  }, props.duration);
});

onUnmounted(() => {
  visible.value = true;
});
</script>

<template>
  <transition name="fade">
    <div v-if="visible" :class="['toast', THEME_CLASSES[props.theme]]">{{ message }}</div>
  </transition>
</template>

<style scoped>
.toast {
  position: fixed;
  bottom: 20px;
  right: 20px;
  padding: 10px 20px;
  background-color: #333;
  color: #fff;
  border-radius: 5px;
  opacity: 1;
  transition: opacity 0.5s;
}

.toast.error {
  background-color: #ff0000;
}

.toast.success {
  background-color: #00ff00;
  color: black;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.5s;
}

.fade-enter,
.fade-leave-to {
  opacity: 0;
}
</style>
