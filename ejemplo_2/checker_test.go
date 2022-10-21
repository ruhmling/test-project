package ejemplo_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckAndNormalizeType(t *testing.T) {
	p := newPayment()
	p.Type = "foo "
	checkAndNormalizeType(p)

	assert.Equal(t, "foo", p.Type)
}

func TestCheckInvalidTypeError(t *testing.T) {
	assert.Equal(t, "invalid payment type", checkAndNormalizeType(newPayment()).Error())
}

func TestCheckTypeCH1_ToReview(t *testing.T) {
	p := newPayment()
	p.Type = "CH1"
	err := checkAndNormalizeType(p)

	assert.Nil(t, err)
	assert.Equal(t, "to_review", p.Detail)
}

func TestCheckValidType_NoError(t *testing.T) {
	p := newPayment()
	p.Type = "test"
	err := checkAndNormalizeType(p)

	assert.Nil(t, err)
}
