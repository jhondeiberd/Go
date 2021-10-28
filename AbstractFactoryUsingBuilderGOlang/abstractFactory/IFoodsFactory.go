package abstractFactory

import (
	"fmt"
	"myProject7/PizzaBuilder"
	"myProject7/SandwichBuilder"
)

type IFoodsFactory interface {
	CreatePizza() PizzaBuilder.Pizza
	CreateSandwich() SandwichBuilder.Sandwich
}

func GetFoodsFactory(typ string) (IFoodsFactory, error) {
	if typ == "veggy" {
		return VEGGYDealer{} , nil
	}

	if typ == "hawaiian" {
		return HAWAIIANDealer{}, nil
	}

	return nil, fmt.Errorf("Wrong type passed")
}
