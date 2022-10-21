package ejemplo_3

import "fmt"

func checkAndNormalizePaymentType(payment *Payment) error {
	payment.SanitizeType()

	if payment.Type == "" {
		return fmt.Errorf("invalid payment type")
	}

	if payment.Type == "CH1" {
		println("pm type ch1")
		payment.Detail = "to_review"
	}

	return nil
}
