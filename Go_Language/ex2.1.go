package main

import "fmt"

func main(){

	var Floors int

	fmt.Print("Enter the nunbers of floors -")
	fmt.Scan(&Floors)

	Height:=(Floors*4)+6

	fmt.Println("Total height of bulding is " , Height);
}
