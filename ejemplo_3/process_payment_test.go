package ejemplo_3

import (
	"fmt"
	"github.com/cucumber/godog"
	"testing"
	"time"
)

type feature struct{}

type context struct {
	payment    *Payment
	gotPayment *Payment
	gotError   error
}

func (c *context) whenProcessThePayment() {
	c.gotPayment, c.gotError = ValidateProcessPayment(c.payment)
}

func (c *context) thenReturnsAPaymentWithDetail(detail string) error {
	if c.gotPayment.Detail != detail {
		return fmt.Errorf("RESPONSE: Got different detail, expected %s - got %s", detail, c.gotPayment.Detail)
	}

	return nil
}

func (c *context) givenPaymentWithAmount(amount float64) {
	c.payment.Amount = amount
}

func (c *context) givenPaymentWithStatus(status string) {
	c.payment.Status = status
}

func (c *context) givenNewPayment() {
	c.payment = new(Payment)
}

func (c *context) thenReturnsANonError() error {
	return c.gotError
}

func (c *context) givenPaymentWithType(paymentType string) {
	c.payment.Type = paymentType
}

func (feature) InitializeScenario(ctx *godog.ScenarioContext) {
	contextFeature := new(context)

	ctx.Step(`^a New Payment$`, contextFeature.givenNewPayment)
	ctx.Step(`^the payment has amount (\d+)$`, contextFeature.givenPaymentWithAmount)
	ctx.Step(`^the payment has status "([^"]*)"$`, contextFeature.givenPaymentWithStatus)
	ctx.Step(`^validate process the payment$`, contextFeature.whenProcessThePayment)
	ctx.Step(`^returns a payment with detail "([^"]*)"$`, contextFeature.thenReturnsAPaymentWithDetail)
	ctx.Step(`^returns a non error$`, contextFeature.thenReturnsANonError)
	ctx.Step(`^the payment has type "([^"]*)"$`, contextFeature.givenPaymentWithType)
}

func TestPaymentValidationFeature(t *testing.T) {
	testFeature := godog.TestSuite{
		Name: "payment_process",
		ScenarioInitializer: func(ctx *godog.ScenarioContext) {
			updateTransactionFeature := new(feature)
			updateTransactionFeature.InitializeScenario(ctx)
		},
		Options: &godog.Options{
			Format:    "pretty",
			Paths:     []string{fmt.Sprintf("./%s.feature", "payment_process")},
			Tags:      "~@wip",
			Randomize: time.Now().UTC().UnixNano(),
			Strict:    true,
		},
	}

	if testFeature.Run() != 0 {
		t.Fail()
	}
}
