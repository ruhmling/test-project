package ejemplo_1

import (
	"fmt"
	"strings"
)

const (
	errType = "CH1"
)

func checkType(payment *Payment) error {
	paymentType := sanitizeType(payment.Type)

	if paymentType == "" {
		return fmt.Errorf("invalid payment type")
	}

	if strings.Contains(paymentType, errType) {
		payment.Detail = "to_review"
	}

	return nil
}

func sanitizeType(paymentType string) string {
	return strings.TrimSpace(paymentType)
}
