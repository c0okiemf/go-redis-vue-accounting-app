package types

import (
	"encoding/json"
	"fmt"
)

type TransactionType int

const (
	TransactionTypeSent TransactionType = iota
	TransactionTypeReceived
)

var transactionTypeNames = map[TransactionType]string{
	TransactionTypeSent:     "Sent",
	TransactionTypeReceived: "Received",
}

func (t TransactionType) String() string {
	if name, ok := transactionTypeNames[t]; ok {
		return name
	}
	return "Unknown"
}

type TransactionTypeError struct {
	Message string
}

func (e *TransactionTypeError) Error() string {
	return fmt.Sprintf("TransactionTypeError: %s", e.Message)
}

func (t *TransactionType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	switch s {
	case "Received":
		*t = TransactionTypeReceived
	case "Sent":
		*t = TransactionTypeSent
	default:
		return &TransactionTypeError{
			Message: "unknown transaction type",
		}
	}

	return nil
}

func (t TransactionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(transactionTypeNames[t])
}

type TransactionInput struct {
	AccountNumber   string          `json:"accountNumber" binding:"required"`
	AccountName     string          `json:"accountName"`
	IBAN            string          `json:"iban" binding:"required"`
	Address         string          `json:"address"`
	Amount          int             `json:"amount" binding:"required,min=1"`
	TransactionType TransactionType `json:"transactionType" binding:"omitempty"`
}

type Transaction struct {
	ID string `json:"id"`
	TransactionInput
}
