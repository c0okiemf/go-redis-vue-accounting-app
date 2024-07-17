import { useToast } from '@/hooks/useToast';
import router from '@/router';
import { ToastTheme } from '@/types/toast';
import { type Transaction, TransactionType } from '@/types/transaction';
import { ref } from 'vue';

export const useRecordTransactionForm = () => {
  const transaction = ref<Transaction>({
    accountNumber: '',
    accountName: '',
    iban: '',
    address: '',
    amount: 0,
    transactionType: TransactionType.Sent
  });

  const fieldLabels = {
    accountNumber: 'Account Number',
    accountName: 'Account Name',
    iban: 'IBAN',
    address: 'Address',
    amount: 'Amount',
    transactionType: 'Transaction Type'
  };

  const validationErrors = ref<Partial<Record<keyof Transaction, string>>>({});

  const { createToast } = useToast();

  const recordTransaction = async () => {
    try {
      const res = await fetch('http://localhost:8081/api/v1/transaction', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(transaction.value)
      });
      if (!res.ok && res.status === 400) {
        const json = await res.json();
        if (json.errors) {
          const newValidationErrors: Partial<Record<keyof Transaction, string>> = {};
          json.errors.map((error: { key: keyof Omit<Transaction, "id">; reason: string }) => {
            switch (error.reason) {
              case 'required': {
                newValidationErrors[error.key] = `${fieldLabels[error.key]} is required`;
                break;
              }
              case 'enum': {
                newValidationErrors[error.key] = `${fieldLabels[error.key]} - value is not valid`;
                break;
              }
              case 'min': {
                newValidationErrors[error.key] =
                  `${fieldLabels[error.key]} - value should be greater than 0`;
                break;
              }
            }
          });
          validationErrors.value = newValidationErrors;
          return;
        } else {
          throw new Error('Unknown error');
        }
      }
      router.push('/');
      createToast({
        message: 'Transaction recorded successfully',
        duration: 5000
      });
    } catch (error) {
      createToast({
        message: 'Unknown error',
        duration: 5000,
        theme: ToastTheme.Error
      });
      console.error(error);
    }
  };

  return {
    transaction,
    fieldLabels,
    validationErrors,
    recordTransaction
  };
};
