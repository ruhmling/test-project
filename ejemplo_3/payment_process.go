package ejemplo_3

import "fmt"

func ValidateProcessPayment(payment *Payment) (*Payment, error) {
	if payment.IsValid() && payment.Amount <= 5000 {
		if errCheck := checkAndNormalizePaymentType(payment); errCheck != nil {
			return nil, errCheck
		} else {
			return payment, nil
		}
	}

	return nil, fmt.Errorf("invalid amount")
}
