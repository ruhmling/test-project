package ejemplo_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessPaymentAmountZero_ReturnsInvalidAmountError(t *testing.T) {
	p := newPayment()
	_, err := processPayment(p)

	assert.Equal(t, "invalid amount", err.Error())
}

func TestProcessPaymentAmountNegative_ReturnsInvalidAmountError(t *testing.T) {
	p := newPayment()
	p.Amount = -123.5
	_, err := processPayment(p)

	assert.Equal(t, "invalid amount", err.Error())
}

func TestProcessPaymentAmountTooHigh_ReturnsInvalidAmountError(t *testing.T) {
	p := newPayment()
	p.Amount = 6000
	_, err := processPayment(p)

	assert.Equal(t, "invalid amount", err.Error())
}

func TestProcessPaymentInvalidType_ReturnsInvalidAmountError(t *testing.T) {
	p := newPayment()
	p.Amount = 10
	p.Type = ""
	_, err := processPayment(p)

	assert.Equal(t, "invalid payment type", err.Error())
}

func TestProcessValidPayment_ReturnsValidPaymentAndNilError(t *testing.T) {
	p := newPayment()
	p.Amount = 10
	p.Type = "foo "
	payment, err := processPayment(p)

	assert.Nil(t, err)
	assert.NotNil(t, payment)
	assert.Equal(t, "foo", p.Type)
}
