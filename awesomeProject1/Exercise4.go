package main

/* Exercise24
In a game, the player tosses two coins. Let’s suppose that, if the first and second coin land on heads, the player
wins $10; if the first lands on heads and the second on tails, the player wins $5; otherwise, the player loses.
We want a program that reads the value of the two coins (heads or tails) and determines whether the player has won.
If yes, it should display the amount won.
*/

/* Exercise25
Write a program that reads 3 values, determines the greatest one, and displays it.
 */

/* Exercise26
Write a program that reads three values and displays them in ascending order.
 */

/* Exercise27
The Ministère des Finances of Québec is adopting a project aiming to reduce taxes. Develop an algorithm that
calculates taxes according to the table provided below. In addition, a 2% reduction of the tax rate is granted
if the person is married. Furthermore, a 0.5% reduction is granted for each child. Finally, 8% is subtracted from
the tax rate for those who have newly arrived in the province. Determine the amount of tax to be paid as a function
of the information provided by the user.
Salary							Tax rate
$0.00 to $18,000.00				10%
$18,000.01 to $32,000.00		20%
$32,000.01 to $60,000.00		30%
$60,000.01 and more				40%
*/

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("**** Exercise 24  *****")
	var coins1, coins2 int
	fmt.Println("Does the first coin land on heads (true = 1)?: ")
	fmt.Scanf("%d", &coins1)
	fmt.Println("Does the second coin land on heads (true = 1): ")
	fmt.Scanf("%d", &coins2)
	HasWon(coins1, coins2)

	fmt.Println("\n\n**** Exercise 25  *****")
	var num1, num2, num3 int
	fmt.Println("Enter the first number: ")
	fmt.Scanf("%d", &num1)
	fmt.Println("Enter the second number: ")
	fmt.Scanf("%d", &num2)
	fmt.Println("Enter the third number: ")
	fmt.Scanf("%d", &num3)
	GreatestNumber(num1, num2, num3)

	fmt.Println("\n\n**** Exercise 26  *****")
	var nu1, nu2, nu3 int
	fmt.Println("Enter the first number: ")
	fmt.Scanf("%d", &nu1)
	fmt.Println("Enter the second number: ")
	fmt.Scanf("%d", &nu2)
	fmt.Println("Enter the third number: ")
	fmt.Scanf("%d", &nu3)
	SorterAscNumber(nu1, nu2, nu3)

	fmt.Println("\n\n**** Exercise 27  *****")
	var numChildren, married int
	var newly bool
	var salary, totalTax, amountValueTax float64
	fmt.Print("Input your salary: ")
	fmt.Scanf("%f", &salary)
	fmt.Print("Are you married? (Yes: 1, No: 0): ")
	fmt.Scanf("%d", &married)
	fmt.Print("Input your numbers the children: ")
	fmt.Scanf("%d", &numChildren)
	fmt.Print("Have you just arrived in the province? (Yes: 1, No: 0): ")
	fmt.Scanf("%v", &newly)
	totalTax = TaxRangeSalary(salary) - DiscMarried(married) - HasChildren(numChildren) - DiscNewly(newly)
	amountValueTax = salary * totalTax
	fmt.Println("The total tax to pay is: ", fmt.Sprintf("%.4f", totalTax), ", the total amount to pay is: ",
				fmt.Sprintf("%.4f", amountValueTax))

}

func TaxRangeSalary(amount float64) float64 {
	const taxRate1 float64 = 0.1
	const taxRate2 float64 = 0.2
	const taxRate3 float64 = 0.3
	const taxRate4 float64 = 0.4
	const RangeSal1 float64 = 18000
	const RangeSal2 float64 = 32000
	const RangeSal3 float64 = 60000
	if amount <= RangeSal1 {
		return taxRate1
	} else if amount <= RangeSal2 && amount > RangeSal1 {
		return taxRate2
	} else if amount <= RangeSal3 && amount > RangeSal2 {
		return taxRate3
	} else {
		return taxRate4
	}
}

func DiscMarried(answer int) float64 {
	const percMarried float64 = 0.02
	if answer == 1 {
		return percMarried
	} else {
		return 0
	}
}

func HasChildren(children int) float64 {
	const percChild float64 = 0.005
	if children > 0 {
		return percChild * float64(children)
	} else {
		return 0
	}
}

func DiscNewly(answer bool) float64 {
	const percNewly float64 = 0.08
	if answer == true {
		return percNewly
	} else {
		return 0
	}
}


func HasWon(a int, b int) {
	if  a == 1 && b == 1 {
		fmt.Println("The player wins $10")
	} else if a == 1 {
		fmt.Println("The player wins $5")
	} else{
		fmt.Println("The player loses")
	}
}

func GreatestNumber(n1 int, n2 int, n3 int) {
	if n1 > n2 {
		if n1 > n3 {
			fmt.Println("the largest number is : " + strconv.Itoa(n1))
		} else {
			fmt.Println("the largest number is: " + strconv.Itoa(n3))
		}
	} else if n2 > n3 {
		fmt.Println("the largest number is: " + strconv.Itoa(n2))
	} else {
		fmt.Println("the largest number is: " + strconv.Itoa(n3))
	}
}

func SorterAscNumber(n1 int, n2 int, n3 int) {
	if n1 >= n2{
		if n2 >= n3{
			fmt.Println("The sorted numbers are " + strconv.Itoa(n3) + " " + strconv.Itoa(n2) + " " + strconv.Itoa(n1))
		} else if n3 >= n1{
			fmt.Println("The sorted numbers are " + strconv.Itoa(n2) + " " + strconv.Itoa(n1) + " " + strconv.Itoa(n3))
		} else if n1 > n3{
			fmt.Println("The sorted numbers are " + strconv.Itoa(n2) + " " + strconv.Itoa(n3) + " " + strconv.Itoa(n1))
		}
	}
	if n2 > n1{
		if n3 >= n2{
			fmt.Println("The sorted numbers are " + strconv.Itoa(n1) + " " + strconv.Itoa(n2) + " " + strconv.Itoa(n3))
		} else if n3 >= n1{
			fmt.Println("The sorted numbers are " + strconv.Itoa(n1) + " " + strconv.Itoa(n3) + " " + strconv.Itoa(n2))
		} else if n1 > n3{
			fmt.Println("The sorted numbers are " + strconv.Itoa(n3) + " " + strconv.Itoa(n1) + " " + strconv.Itoa(n2))
		}
	}
}