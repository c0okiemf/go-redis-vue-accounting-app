<script setup lang="ts">
defineProps<{
  label: string;
  error?: string;
}>();
</script>

<template>
  <div class="form-group">
    <label for="accountNumber">{{ label }}</label>
    <slot></slot>
    <span v-if="error" class="error">{{ error }}</span>
  </div>
</template>

<style lang="scss" scoped>
.form-group {
  width: 100%;
  display: flex;
  flex-direction: column;
  position: relative;
}

label {
  font-size: 0.8rem;
}

.error {
  color: red;
  font-size: 0.8rem;
  position: absolute;
  bottom: -1.2rem;
}
</style>
