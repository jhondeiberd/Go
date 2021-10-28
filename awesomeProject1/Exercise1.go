package main

/* Exercise2
In the old system of calculating sales tax in Québec, the taxes on a product were 7% for the GST, and 7.5%
for the QST (applied after calculating the GST). Make a program that reads the unit price of a product and
the quantity purchased, and that displays the amounts for the GST, the QST, and the total price after taxes.
*/

/* Exercise3
Make a program that calculates an employee’s gross salary for a week, given their hourly rate and
the number of hours worked.
*/

/* Exercise4
Make a program that reads a temperature in degrees Fahrenheit and converts it into degrees Celsius.
The conversion formula is: C = ((F-32)/9)*5
*/

/* Exercise5
Make a program that displays the volume of a rectangular prism after reading the dimensions of the prism.
 */

/* Exercise6
A dealership selling new vehicles asks you to construct a program that calculates the compensation
paid to their salespeople. The base salary for all the salespeople is $400. For each vehicle sold,
the salesperson receives a commission of $200. Also, the salesperson receives a bonus of 5% of the total
amount of all their sales. Make the program for a salesperson.
 */

import (
	"fmt"
	//"strconv"
)

func main() {

	fmt.Println("**** Exercise 2  *****")
	const gst float64 = 0.07
	const qst float64 = 0.075
	var price, qty, amount, valueGst, valueQst, total float64
	fmt.Print("Enter the price to product: ")
	fmt.Scanf("%f", &price)
	fmt.Print("Enter the quantity to product: ")
	fmt.Scanf("%f", &qty)
	amount = price * qty
	valueGst = amount * gst
	valueQst = (amount + valueGst) * qst
	total = amount + valueGst + valueQst
	fmt.Println("The amount is: \t", fmt.Sprintf("%.4f", amount),
				"\nGST: \t\t\t", fmt.Sprintf("%.4f", valueGst),
				"\nQST: \t\t\t", fmt.Sprintf("%.4f", valueQst),
				"\nThe total is: \t", fmt.Sprintf("%.4f", total))

	fmt.Println("\n\n**** Exercise 3  *****")
	var hour, valueHour, grossSalary float64
	fmt.Print("Enter the number of hours worked in the week: ")
	fmt.Scanf("%f", &hour)
	fmt.Print("Enter the value of the hour: ")
	fmt.Scanf("%f", &valueHour)
	grossSalary = hour * valueHour
	fmt.Println("The gross salary by Employed for this week is : ", fmt.Sprintf("%.4f", grossSalary))

	fmt.Println("\n\n**** Exercise 4  *****")
	var F float64
	fmt.Print("Enter temperature in degrees Fahrenheit: \n")
	fmt.Scanf("%f", &F)
	fmt.Println("The temperature in degrees Celsius is : ", fmt.Sprintf("%.2f", ConvertFtoC(F)))

	fmt.Println("\n\n**** Exercise 5  *****")
	var length, width, height float64
	fmt.Print("Enter the length: ")
	fmt.Scanf("%f", &length)
	fmt.Print("Enter the width: ")
	fmt.Scanf("%f", &width)
	fmt.Print("Enter the height: ")
	fmt.Scanf("%f", &height)
	fmt.Println("The volume of a rectangular prism is : ", fmt.Sprintf("%.2f", Volume(length, width, height)))

	fmt.Println("\n\n**** Exercise 6  *****")
	const baseSalary float64 = 400
	const commissionCar float64 = 200
	const commissionSales float64 = 0.05
	var carSold, salary float64
	fmt.Print("Enter number the cars sold: ")
	fmt.Scanf("%f", &carSold)
	salary = baseSalary + (commissionCar * carSold) + (commissionCar * carSold)*commissionSales
	fmt.Println("The compensation pays to salesperson is : ", fmt.Sprintf("%.4f", salary))
}

func ConvertFtoC (temp float64) float64{
	return ((temp-32)/9)*5
}

func Volume (length float64, width float64, height float64) float64 {
	return length * width * height
}
