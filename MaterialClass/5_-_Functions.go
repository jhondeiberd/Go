package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(Multiply(4, 5))
	fmt.Println(printName("for dara"))
}

func printName(x string) string {
	return "returns printName " + x
}

func Multiply(a int, b int) string {
	return "returns Multiply " + strconv.Itoa(a*b)

}
