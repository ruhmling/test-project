package ejemplo_1

import (
	"fmt"
	"github.com/ruhmling/test-project/helpers/log"
)

func validateProcessPayment(payment *Payment) (*Payment, error) {
	var errCheck error

	if payment.IsValid() {
		errCheck = isPaymentToProcess(payment)

		if errCheck == nil {
			return payment, nil
		}
	} else {
		errCheck = fmt.Errorf("invalid amount")
	}

	log.Debug("Payment is invalid")

	return payment, errCheck
}

func isPaymentToProcess(payment *Payment) error {
	err := checkType(payment)

	if payment.Amount > 5000 {
		return fmt.Errorf("invalid amount")
	}

	// ToDo: add detail to_process for pending payments
	//if err == nil && payment.Status == "pending" {
	//	payment.Detail = "to_process"
	//}

	return err
}
