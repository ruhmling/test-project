package ejemplo_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPayment_IsValid(t *testing.T) {
	_, e := validateProcessPayment(&Payment{Amount: 10, Status: "pending", Type: "test"})
	assert.Nil(t, e)
}

func TestPayment_IsInvalidAmount(t *testing.T) {
	_, e := validateProcessPayment(&Payment{Amount: 0, Status: "pending", Type: "test"})
	assert.NotNil(t, e)
}

func TestPayment_IsInvalidAmount2(t *testing.T) {
	_, e := validateProcessPayment(&Payment{Amount: 6000, Status: "pending", Type: "test"})
	assert.NotNil(t, e)
}

func TestPayment_IsInvalidType(t *testing.T) {
	_, e := validateProcessPayment(&Payment{Amount: 10, Status: "pending"})
	assert.NotNil(t, e)
}

func TestPayment_CheckTypeIsInvalidType(t *testing.T) {
	p, e := validateProcessPayment(&Payment{Amount: 10, Status: "pending", Type: "CH1"})
	assert.Nil(t, e)
	assert.Equal(t, "to_review", p.Detail)
}
