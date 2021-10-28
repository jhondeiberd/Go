package abstractFactory

import (
	"myProject7/PizzaBuilder"
	"myProject7/SandwichBuilder"
)

type VEGGYDealer struct{}

func (V VEGGYDealer) CreatePizza() PizzaBuilder.Pizza {
	veggyBuilder := PizzaBuilder.GetBuilder("veggy")
	waitress := PizzaBuilder.NewWaitressPizza(veggyBuilder)
	veggyPizza := waitress.BuildPizza()

	return veggyPizza
}

func (V VEGGYDealer) CreateSandwich() SandwichBuilder.Sandwich {

	veggyBuilder := SandwichBuilder.GetBuilder("veggy")
	waitress := SandwichBuilder.NewWaitressSandwich(veggyBuilder)
	veggySandwich := waitress.BuildSandwich()

	return veggySandwich
}

