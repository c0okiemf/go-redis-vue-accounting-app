package helpers

import (
	"accounting/types"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func MakeAddTransactionUserFacingErrors(err error, account types.TransactionInput) []types.ValidationError {
	userFacingErrors := []types.ValidationError{}
	// Handling validation errors
	if validationErr, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErr {
			inputType := reflect.TypeOf(account)
			field, _ := inputType.FieldByName(fieldErr.StructField())
			userFacingErrors = append(userFacingErrors, types.ValidationError{
				Key:    field.Tag.Get("json"),
				Reason: fieldErr.Tag(),
			})
		}
	}
	println(err.Error())
	// Handling transaction type error
	if _, ok := err.(*types.TransactionTypeError); ok {
		userFacingErrors = append(userFacingErrors, types.ValidationError{
			Key:    "transactionType",
			Reason: "enum",
		})
	}
	return userFacingErrors
}
