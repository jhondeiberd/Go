package abstractFactory

import (
	"myProject7/PizzaBuilder"
	"myProject7/SandwichBuilder"
)

type HAWAIIANDealer struct{}

func (V HAWAIIANDealer) CreatePizza() PizzaBuilder.Pizza {

	hawaiianBuilder := PizzaBuilder.GetBuilder("hawaiian")
	waitress := PizzaBuilder.NewWaitressPizza(hawaiianBuilder)
	hawaiianPizza := waitress.BuildPizza()

	return hawaiianPizza
}

func (V HAWAIIANDealer) CreateSandwich() SandwichBuilder.Sandwich {
	hawaiianBuilder := SandwichBuilder.GetBuilder("hawaiian")
	waitress := SandwichBuilder.NewWaitressSandwich(hawaiianBuilder)
	hawaiianSandwich := waitress.BuildSandwich()

	return hawaiianSandwich
}
