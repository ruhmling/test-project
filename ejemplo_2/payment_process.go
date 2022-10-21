package ejemplo_2

import "fmt"

func processPayment(payment *Payment) (*Payment, error) {
	if payment.IsValid() && payment.Amount <= 5000 {
		if errCheck := checkAndNormalizeType(payment); errCheck != nil {
			return nil, errCheck
		} else {
			return payment, nil
		}
	}

	return nil, fmt.Errorf("invalid amount")
}
