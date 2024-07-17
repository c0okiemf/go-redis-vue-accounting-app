import type { Transaction } from '@/types/transaction';
import { onMounted, ref } from 'vue';

export const ITEMS_PER_PAGE = 10;

export const useFetchTransactions = () => {
  const transactions = ref<Transaction[]>([]);
  const numTransactons = ref<number>(0);
  const page = ref<number>(1);

  const fetchTransactions = () =>
    fetch(`http://localhost:8081/api/v1/transactions?page=${page.value}&per_page=${ITEMS_PER_PAGE}`)
      .then((response) => response.json())
      .then((data: { transactions: Transaction[]; total: number }) => {
        transactions.value = data.transactions;
        numTransactons.value = data.total;
      });

  onMounted(() => {
    fetchTransactions();
  });

  const setPage = (newPage: number) => {
    page.value = newPage;
    fetchTransactions();
  };

  return {
    transactions,
    numTransactons,
    page,
    fetchTransactions,
    setPage
  };
};
