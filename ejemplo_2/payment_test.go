package ejemplo_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func newPayment() *Payment {
	return new(Payment)
}

func TestCreateNewPayment(t *testing.T) {
	assert.NotNil(t, newPayment())
}

func TestSetPaymentAmount(t *testing.T) {
	p := newPayment()
	p.Amount = 50.4
	assert.Equal(t, 50.4, p.Amount)
}

func TestSetPaymentType(t *testing.T) {
	p := newPayment()
	p.Type = "test-type"
	assert.Equal(t, "test-type", p.Type)
}

func TestSetPaymentStatus(t *testing.T) {
	p := newPayment()
	p.Status = "pending"
	assert.Equal(t, "pending", p.Status)
}

func TestSetPaymentDetail(t *testing.T) {
	p := newPayment()
	p.Detail = "test-detail"
	assert.Equal(t, "test-detail", p.Detail)
}

func TestNewInvalidPayment(t *testing.T) {
	assert.False(t, newPayment().IsValid())
}

func TestNewValidPayment(t *testing.T) {
	p := newPayment()
	p.Amount = 10

	assert.True(t, p.IsValid())
}

func TestSanitizePaymentType(t *testing.T) {
	p := newPayment()
	p.Type = " foo"
	p.SanitizeType()

	assert.Equal(t, "foo", p.Type)
}
