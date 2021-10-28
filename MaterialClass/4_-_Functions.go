
package main

import (
    "fmt"
)

func main() {
    printName()
    fmt.Println(Multiply(15, 5))
}

func printName() {
    fmt.Println("jacques marchand")
}

func Multiply(a int, b int) int {
    return a * b
}