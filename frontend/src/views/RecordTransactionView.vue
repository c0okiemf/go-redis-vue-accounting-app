<script setup lang="ts">
import RecordTransactionForm from '@/components/Form/RecordTransactionForm.vue';
</script>

<template>
  <div class="container">
    <RecordTransactionForm />
  </div>
</template>

<style lang="scss" scoped>
.container {
  @media (max-width: $desktop-breakpoint) {
    padding: 4rem 1rem 1rem;
  }
}
</style>
