package main

import "fmt"

func main(){

	var veichlesold float64
	var totalsales float64

	fmt.Print("Enter Veichle sold -")
	fmt.Scan(&veichlesold)

	fmt.Print("Enter Sales - ")
	fmt.Scan(&totalsales)

	Commision:=veichlesold *200

	Bonus:=totalsales*5/100

	FinalExpense:=400+Commision+Bonus;

	fmt.Println("Sales person will recive ", float64(FinalExpense))

}