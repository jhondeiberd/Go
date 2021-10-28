package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//var age int
	age := 40
	var name string
	name = "JHON"
	fmt.Print("hello world. my name is " + strings.ToUpper(name) + ". I am " + strconv.Itoa(age))
	fmt.Print(" years old.")
}

//hello world. my name is DARA. I am 24 years old.
//                        string     int
