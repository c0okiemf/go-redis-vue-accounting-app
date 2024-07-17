export enum TransactionType {
  Sent = 'Sent',
  Received = 'Received'
}

export interface Transaction {
  id?: number;
  accountNumber: string;
  accountName: string;
  iban: string;
  address: string;
  amount: number;
  transactionType: TransactionType;
}
