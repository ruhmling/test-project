package ejemplo_2

import "fmt"

func checkAndNormalizeType(payment *Payment) error {
	payment.SanitizeType()

	if payment.Type == "" {
		return fmt.Errorf("invalid payment type")
	}

	if payment.Type == "CH1" {
		payment.Detail = "to_review"
	}

	return nil
}
