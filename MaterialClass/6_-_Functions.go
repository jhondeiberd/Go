package main

import "fmt"

func main() {
	if isHere() {
		fmt.Println("func isHere returns true")
	}
}
func isHere() bool {
	var status bool
	status = false

	// do database verification
	if true {
		status = true
	}

	return status

}
