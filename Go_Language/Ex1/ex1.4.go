package main

import "fmt"

func main() {

	var lenght float64;
	var width float64
	var height float64

	fmt.Print("Enter Lenght- ")
	fmt.Scan(&lenght)

	fmt.Print("Enter Width- ")
	fmt.Scan(&width)

	fmt.Print("Enter Height- ")
	fmt.Scan(&height)


	Volume:=lenght*width*height;

	fmt.Println("Volume for prism is ", float64(Volume));


}
