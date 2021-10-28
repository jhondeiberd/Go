package main

/*
In the old system of calculating sales tax in Québec, the taxes on a product were 7% for the GST, and 7.5%
for the QST (applied after calculating the GST). Make a program that reads the unit price of a product and
the quantity purchased, and that displays the amounts for the GST, the QST, and the total price after taxes.
*/

import (
	"fmt"
	//"strconv"
)

func main() {
	var gst float64 = 0.07
	var qst float64 = 0.075
	var price float64
	var qty float64
	var amount float64
	var valueGst float64
	var valueQst float64
	var total float64

	fmt.Println("**** Exercise 2  *****")
	fmt.Print("Enter the price to product: \n")
	fmt.Scanf("%f", &price)

	fmt.Print("Enter the quantity to product: \n")
	fmt.Scanf("%f", &qty)

	amount = price * qty
	valueGst = amount * gst
	valueQst = (amount + valueGst) * qst
	total = amount + valueGst + valueQst

	fmt.Println("The amount is: \t", float64(amount), "\nGST: \t\t\t", float64(valueGst), "\nQST: \t\t\t", float64(valueQst), "\nThe total is: \t", float64(total))

}