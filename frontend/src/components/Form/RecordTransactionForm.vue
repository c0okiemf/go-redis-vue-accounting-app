<script setup lang="ts">
import { useRecordTransactionForm } from '@/hooks/useRecordTransactionForm';
import { TransactionType } from '@/types/transaction';
import FormGroup from './FormGroup.vue';

const { transaction, fieldLabels, validationErrors, recordTransaction } =
  useRecordTransactionForm();
</script>

<template>
  <div class="form">
    <form @submit.prevent="recordTransaction">
      <FormGroup :label="fieldLabels.accountNumber" :error="validationErrors.accountName">
        <input v-model="transaction.accountNumber" type="text" id="accountNumber" required />
      </FormGroup>
      <FormGroup :label="fieldLabels.accountName" :error="validationErrors.accountName">
        <input v-model="transaction.accountName" type="text" id="accountName" />
      </FormGroup>
      <FormGroup :label="fieldLabels.iban" :error="validationErrors.iban">
        <input v-model="transaction.iban" type="text" id="iban" required />
      </FormGroup>
      <FormGroup :label="fieldLabels.address" :error="validationErrors.address">
        <input v-model="transaction.address" type="text" id="address" />
      </FormGroup>
      <FormGroup :label="fieldLabels.amount" :error="validationErrors.amount">
        <input v-model="transaction.amount" type="number" id="amount" required />
      </FormGroup>
      <FormGroup :label="fieldLabels.transactionType" :error="validationErrors.transactionType">
        <select v-model="transaction.transactionType" id="transactionType" required>
          <option :value="TransactionType.Sent">Sent</option>
          <option :value="TransactionType.Received">Received</option>
        </select>
      </FormGroup>
      <button type="submit">Record Transaction</button>
    </form>
  </div>
</template>

<style lang="scss" scoped>
.form {
  max-width: 400px;
  margin: 0 auto;
  width: 100%;
}

form {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

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
