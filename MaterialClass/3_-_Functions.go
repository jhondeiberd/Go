//control flow:
//(1) sequence
//(2) loop : iterative
//(3) conditional

package main

import "fmt"

func main() {
	fmt.Println("I'm in main")

	foo1()

	fmt.Println("I'm in main after foo1 call")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	foo2()
}

func foo1() {
	fmt.Println("I'm in foo1")
}

func foo2() {
	fmt.Println("I'm in foo2")
}


