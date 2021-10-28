/*
implementation of abstract factory on food project
adidas =	veggy
nike =	hawaiian

shoe =	pizza
shirt =	sandwich
*/

package main

import (
	"fmt"
	"AbstractFactoryUsingGolang/abstractFactory"
)

///////////////////////////////////////////////
func main() {
	veggyFactory, _ := abstractFactory.GetFoodsFactory("veggy")
	hawaiianFactory, _ := abstractFactory.GetFoodsFactory("hawaiian")

	hawaiianPizza := hawaiianFactory.MakePizza()
	hawaiianSandwich := hawaiianFactory.MakeSandwich()

	veggyPizza := veggyFactory.MakePizza()
	veggySandwich := veggyFactory.MakeSandwich()

	printPizzaDetails(hawaiianPizza)
	printSandwichDetails(hawaiianSandwich)

	printPizzaDetails(veggyPizza)
	printSandwichDetails(veggySandwich)
}

func printPizzaDetails(s abstractFactory.IPizza) {
	fmt.Println("----------------------")
	fmt.Printf("Type: %s", s.GetTyp())
	fmt.Println()
	fmt.Printf("Dough: %s", s.GetDough())
	fmt.Println()
	fmt.Printf("Sauce: %s", s.GetSauce())
	fmt.Println()
	fmt.Printf("Topping: %s", s.GetTopping())
	fmt.Println()
}

func printSandwichDetails(s abstractFactory.ISandwich) {
	fmt.Println("----------------------")
	fmt.Printf("Type: %s", s.GetTyp())
	fmt.Println()
	fmt.Printf("Bread: %s", s.GetBread())
	fmt.Println()
	fmt.Printf("Salade: %s", s.GetSalade())
	fmt.Println()
	fmt.Printf("Topping: %s", s.GetTopping())
	fmt.Println()
}




