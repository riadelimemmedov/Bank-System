package api

import (
	"simplebank/util"

	"github.com/go-playground/validator/v10"
)

// !validCurrencies
var validCurrencies validator.Func = func(fieldLavel validator.FieldLevel) bool {
	if currency, ok := fieldLavel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
