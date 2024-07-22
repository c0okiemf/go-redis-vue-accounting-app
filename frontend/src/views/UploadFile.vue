<script setup lang="ts">
import RecordTransactionForm from '@/components/Form/RecordTransactionForm.vue';
import { onMounted, ref } from 'vue';

const fileNames = ref<string[]>([]);

const fetchTransactions = () =>
  fetch(`http://localhost:8081/api/v1/uploads`)
    .then((response) => response.json())
    .then((data) => {
      fileNames.value = data.files;
    });

onMounted(() => {
  fetchTransactions();
});
</script>

<template>
  <div class="container">
    <div v-for="file in fileNames">{{ file }}</div>
  </div>
</template>

<style lang="scss" scoped>
.container {
  @media (max-width: $desktop-breakpoint) {
    padding: 4rem 1rem 1rem;
  }
}
</style>
