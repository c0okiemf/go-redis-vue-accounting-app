<script setup lang="ts">
import ListPagination from '@/components/ListPagination.vue';
import TransactionItem from '@/components/TransactionItem.vue';
import { ITEMS_PER_PAGE, useFetchTransactions } from '@/hooks/useFetchTransactions';
import { ref } from 'vue';

const highlightedColumnIndex = ref<number | null>(null);

const { transactions, numTransactons, page, setPage } = useFetchTransactions();

const onCellHover = (cellIndex: number) => {
  highlightedColumnIndex.value = cellIndex;
};

const onCellLeave = () => {
  highlightedColumnIndex.value = null;
};
</script>

<template>
  <div class="container">
    <ListPagination
      :total="numTransactons"
      :itemsPerPage="ITEMS_PER_PAGE"
      :currentPage="page"
      @onSetPage="setPage"
    />
    <div class="table-container">
      <div class="table">
        <TransactionItem
          v-for="transaction in transactions"
          :key="transaction.id"
          :transaction="transaction"
          :highlighted-column-index="highlightedColumnIndex"
          @onCellHover="onCellHover"
          @onCellLeave="onCellLeave"
        />
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.container {
  display: flex;
  align-items: center;
  overflow: hidden;

  @media (max-width: $desktop-breakpoint) {
    padding: 4rem 1rem 1rem;
  }
}

.table-container {
  max-height: 100vh;
  overflow: scroll;
  padding: 2rem 0;
  max-width: 100%;
}

.table {
  display: table;
  table-layout: fixed;
  min-width: 100%;
  width: auto;

  @media (max-width: $desktop-breakpoint) {
    width: auto;
    margin: 0 auto;
  }
}
</style>
