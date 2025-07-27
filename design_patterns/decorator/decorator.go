package decorator

import "fmt"

type Decorator interface {
	GetPrice() int
	GetIngredient() string
}

type Coffee struct {
	price          int
	baseIngredient string
}

func (c *Coffee) GetPrice() int {
	return c.price
}

func (c *Coffee) GetIngredient() string {
	return c.baseIngredient
}

type Milk struct {
	baseDecorator Decorator
}

func (m *Milk) GetPrice() int {
	return m.baseDecorator.GetPrice() + 10
}

func (m *Milk) GetIngredient() string {
	return fmt.Sprintf("milk is added to the top of the [%v]", m.baseDecorator.GetIngredient())
}

type Sugar struct {
	baseDecorator Decorator
}

func (m *Sugar) GetPrice() int {
	return m.baseDecorator.GetPrice() + 4
}

func (m *Sugar) GetIngredient() string {
	return fmt.Sprintf("sugar is added to the top of the [%v]", m.baseDecorator.GetIngredient())
}

func DecoratorPattern() {
	coffee := &Coffee{
		baseIngredient: "black coffee",
		price: 20,
	}

	milk := &Milk{coffee}
	fmt.Println(milk.GetPrice())
	fmt.Println(milk.GetIngredient())

	sugar := &Sugar{milk}
	fmt.Println(sugar.GetPrice())
	fmt.Println(sugar.GetIngredient())
}