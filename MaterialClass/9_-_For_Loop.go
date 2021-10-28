package main

import (
	"fmt"
	"strconv"
)

func main() {
	var total string
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			for i := 0; i < 5; i++ {
				total = total + " " + strconv.Itoa(i)
			}
			fmt.Println(total)
		}
	}

}
