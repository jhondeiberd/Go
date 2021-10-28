package main

import (
	"fmt"
)

func main()  {
	var Far float32
	fmt.Print("Enter the Fahrenheit- ")
	fmt.Scan(&Far)

	Celcius:=(Far-32)* 5/9;

	println("Celcius is ",float64(Celcius))
}
