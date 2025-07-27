package factory

import (
	"fmt"


)

type PaymentMethod interface {
	Pay(amount int) string
}

type Credit struct{}
type Debit struct{}
type UPI struct{}

func (c *Credit) Pay(amount int) string {
	return fmt.Sprintf("Amount: %d is paid using Credit card",amount)
}

func (c *Debit) Pay(amount int) string {
	return fmt.Sprintf("Amount: %d is paid using Debit card",amount)
}

func (c *UPI) Pay(amount int) string {
	return fmt.Sprintf("Amount: %d is paid using UPI",amount)
}

func PaymentFactory(method string) (PaymentMethod, error) {
	switch method {
	case "credit":
		return &Credit{}, nil
	case "debit":
		return &Debit{}, nil
	case "upi":
		return &UPI{}, nil
	default:
		return nil, fmt.Errorf("incorrect payment method")
	}
}

func FactoryPattern() {
	m, _ := PaymentFactory("upi")
	fmt.Println(m.Pay(400))
	m, _ = PaymentFactory("debit")
	fmt.Println(m.Pay(4033))
}
