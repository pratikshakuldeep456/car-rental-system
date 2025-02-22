package pkg

// The PaymentProcessor interface defines
// the contract for payment processing,
// and the CreditCardPaymentProcessor and P
// PayPalPaymentProcessor classes are concrete
// implementations of the payment processor.

type PaymentProcessor interface {
	Pay(amount float64) error
}

type CreditCardPaymentProcess struct {
}

func (c *CreditCardPaymentProcess) Pay(amount float64) error {
	return nil
}

type PayPalPaymentProcessor struct {
}

func (p *PayPalPaymentProcessor) Pay(amount float64) error {
	return nil
}

type PaymentService struct {
	processor PaymentProcessor
}

func NewPaymentService(p PaymentProcessor) *PaymentService {
	return &PaymentService{processor: p}
}

func (ps *PaymentService) ProcessPayment(amount float64) error {
	return ps.processor.Pay(amount)
}
