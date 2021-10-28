package main

import (
	"fmt"
	"strconv"
)

func main() {

	var age string
	age = "44"
	var name string
	name = "Dara"
	fmt.Println("hello world ! My name is " + name + ". I am " + age + " years old.")
	s, err := strconv.Atoi(age)

	fmt.Print(s, err)

}

//hello world ! My name is Dara. I am 44 years old.
//                         string     int