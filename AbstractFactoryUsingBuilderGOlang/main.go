package main

import (
	"fmt"
	"myProject7/PizzaBuilder"
	"myProject7/SandwichBuilder"
	"myProject7/abstractFactory"
)

///////////////////////////////////////////////

func main() {

	veggyFactory, _ := abstractFactory.GetFoodsFactory("veggy")
	hawaiianFactory, _ := abstractFactory.GetFoodsFactory("hawaiian")

	hawaiianPizza := hawaiianFactory.CreatePizza()
	hawaiianSandwich := hawaiianFactory.CreateSandwich()

	veggyPizza := veggyFactory.CreatePizza()
	veggySandwich := veggyFactory.CreateSandwich()

	printPizza(veggyPizza)
	printPizza(hawaiianPizza)

	printSandwich(veggySandwich)
	printSandwich(hawaiianSandwich)

}

func printPizza(p PizzaBuilder.Pizza){
	fmt.Println("----------------------------")
	fmt.Printf(" pizza dough: %s\n", p.Dough)
	fmt.Printf(" pizza sauce: %s\n", p.Sauce)
	fmt.Printf(" pizza topping: %s\n", p.Topping)
}

func printSandwich(s SandwichBuilder.Sandwich){
	fmt.Println("----------------------------")
	fmt.Printf(" sandwich dough: %s\n", s.Bread)
	fmt.Printf(" sandwich salade: %s\n", s.Salade)
	fmt.Printf(" sandwich topping: %s\n", s.Topping)
}