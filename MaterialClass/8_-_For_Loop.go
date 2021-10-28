package main

import (
	"fmt"
)

func main() {

	i := 10
	for {
		fmt.Println("i = ", i)
		if i <= 5 {
			fmt.Println("Loop terminated at 5")
			break
		}
		i -= 1
	}
}