package main

import (
	"fmt"
)

func main() {

	var kpa float64

	fmt.Print("Enter the kilopascal units- ")
	fmt.Scan(&kpa);

	atm:=kpa/101.325


	fmt.Printf("Atmosphere units is %.2f",atm )


}