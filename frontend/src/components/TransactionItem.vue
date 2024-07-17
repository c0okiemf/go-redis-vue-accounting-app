<script setup lang="ts">
import { TransactionType, type Transaction } from '@/types/transaction';
import { ref } from 'vue';

defineProps<{
  transaction: Transaction;
  highlightedColumnIndex: number | null;
}>();

defineEmits<{
  (event: 'onCellHover', payload: number): void;
  (event: 'onCellLeave'): void;
}>();

const isRowHighlighted = ref<boolean>(false);

const onRowHover = () => {
  isRowHighlighted.value = true;
};

const onRowLeave = () => {
  isRowHighlighted.value = false;
};
</script>

<template>
  <div
    :class="[
      'row',
      transaction.transactionType === TransactionType.Received && 'received-transaction',
      isRowHighlighted && 'highlighted'
    ]"
    @mouseover="onRowHover"
    @mouseleave="onRowLeave"
  >
    <div
      v-for="(prop, index) in Object.values(transaction)"
      :key="index"
      @mouseover="$emit('onCellHover', index)"
      @mouseleave="$emit('onCellLeave')"
      :class="['cell', index === highlightedColumnIndex && 'highlighted']"
    >
      {{ prop }}
    </div>
  </div>
</template>

<style scoped>
.row {
  display: table-row;
  width: 100%;
}

.received-transaction {
  background: var(--background-row-received);
}

.cell {
  display: table-cell;
  flex: 1;
  padding: 1rem;
  border: 1px solid var(--color-border);
  white-space: nowrap;
}

.highlighted {
  background: var(--background-highlighted);
  color: var(--background-highlighted-text);
}
</style>
