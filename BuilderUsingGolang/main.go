package main

import (
	"BuilderUsingGolang/PizzaBuilder"
	"BuilderUsingGolang/SandwichBuilder"
	"fmt"
)

func main() {

	fmt.Println("---------------------Pizza---------------------")
	veggyBuilderPizza := PizzaBuilder.GetBuilder("veggy")
	hawaiianBuilderPizza := PizzaBuilder.GetBuilder("hawaiian")

	waitressPizza := PizzaBuilder.NewWaitressPizza(veggyBuilderPizza)
	veggyPizza := waitressPizza.BuildPizza()

	fmt.Printf("veggy pizza dough: %s\n", veggyPizza.Dough)
	fmt.Printf("veggy pizza sauce: %s\n", veggyPizza.Sauce)
	fmt.Printf("veggy pizza topping: %s\n", veggyPizza.Topping)

	waitressPizza.SetBuilder(hawaiianBuilderPizza)
	hawaiianPizza := waitressPizza.BuildPizza()

	fmt.Printf("\nhawaiian pizza dough: %s\n", hawaiianPizza.Dough)
	fmt.Printf("hawaiian pizza sauce: %s\n", hawaiianPizza.Sauce)
	fmt.Printf("hawaiian pizza topping: %s\n", hawaiianPizza.Topping)


	fmt.Println("---------------------Sandwich---------------------")

	veggyBuilderSandwich := SandwichBuilder.GetBuilder("veggy")
	hawaiianBuilderSandwich := SandwichBuilder.GetBuilder("hawaiian")

	waitressSandwich := SandwichBuilder.NewWaitressSandwich(veggyBuilderSandwich)
	veggySandwich := waitressSandwich.BuildSandwich()

	fmt.Printf("veggy Sandwich bread: %s\n", veggySandwich.Bread)
	fmt.Printf("veggy Sandwich salade: %s\n", veggySandwich.Salade)
	fmt.Printf("veggy Sandwich topping: %s\n", veggySandwich.Topping)

	waitressSandwich.SetBuilder(hawaiianBuilderSandwich)
	hawaiianSandwich := waitressSandwich.BuildSandwich()

	fmt.Printf("\nhawaiian Sandwich bread: %s\n", hawaiianSandwich.Bread)
	fmt.Printf("hawaiian Sandwich salade: %s\n", hawaiianSandwich.Salade)
	fmt.Printf("hawaiian Sandwich topping: %s\n", hawaiianSandwich.Topping)



}
