package main

/* Exercise2
We want to determine the height of a building of n floors, knowing that the ground floor has a height
of 6 meters and that the other floors each have a height of 4 meters.
*/

/* Exercise3
An aircraft pilot wants to know the atmospheric pressure, expressed in atmosphere units (atm),
as the weather station only provides pressure data in kilopascal units (kPa).
1 atm is equivalent to 101.325 kPa. Make a program that performs the conversion.
*/

/* Exercise4
In a computer technology course, the following evaluation weights are used:
    • Laboratory work counts for 40% of the grade
    • The midterm exam counts for 25%
    • The final exam counts for 35%
Make a program that calculates the final grade of a student, assuming that each of the three grades the user inputs is out of 100.
*/

/* Exercise5
Monique want to have a little program that allows her to evaluate the total amount of her expenses for a month,
as well as the amount of money she can allocate for her leisure activities and non-essential spending.
The program should read her projections for expenses in the following categories:
Weekly expenses (one time per week; assuming that 1 month = 4 weeks)
    • Food expenses and household expenses
    • Common expenses
Monthly expenses (one time per month)
    • Public transit pass
    • Rent
    • Total of monthly bills
The program should also read the amount of her paycheques. Monique receives two paycheques per month.
The program should then display her total expenses, her total income, and the difference.
*/

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("**** Exercise 2  *****")
	var numberFloor int
	fmt.Print("Enter the numbers of the building's floors: ")
	fmt.Scanf("%d", &numberFloor)
	CalculatedHeight(numberFloor)

	fmt.Println("\n\n**** Exercise 3  *****")
	var kPapressure, atmPressure float64
	const convertion float64 = 101.325
	fmt.Print("Enter the pressure data in kilopascal units: ")
	fmt.Scanf("%f", &kPapressure)
	atmPressure = kPapressure / convertion
	fmt.Println("The atmospheric pressure in atmosphere units is : ", fmt.Sprintf("%.4f", atmPressure))

	fmt.Println("\n\n**** Exercise 4  *****")
	var gradeLaboratory, gradeMidtermExam, gradeFinalExam float64
	fmt.Print("Enter grade to laboratory (1 to 100): ")
	fmt.Scanf("%f", &gradeLaboratory)
	fmt.Print("Enter grade to midterm Exam (1 to 100): ")
	fmt.Scanf("%f", &gradeMidtermExam)
	fmt.Print("Enter grade to final Exam (1 to 100): ")
	fmt.Scanf("%f", &gradeFinalExam)
	CalculatedGrade(gradeLaboratory, gradeMidtermExam, gradeFinalExam)

	fmt.Println("\n\n**** Exercise 5  *****")
	var totExpWeek, totMonExpense, totPaycheque float64
	totPaycheque = Paycheques()
	fmt.Println("The total of Incomes\" are: ", fmt.Sprintf("%0.4f", totPaycheque))
	totExpWeek = WeeklyExpenses()
	fmt.Println("The total of Weekly expenses are: ", fmt.Sprintf("%0.4f", totExpWeek))
	totMonExpense = MonthlyExpenses()
	fmt.Println("The total of Monthly expenses are: ", fmt.Sprintf("%0.4f", totMonExpense))
	fmt.Println("\nThe balance is: ", fmt.Sprintf("%0.4f", totPaycheque - totExpWeek - totMonExpense))
}

func WeeklyExpenses () float64 {
	var foodHousehold, commonExpenses, totFoodHousehold, totCommonExpenses float64
	totFoodHousehold = 0
	totCommonExpenses = 0
	fmt.Println("\nWeekly expenses")
	for i := 0; i < 4; i++ {
		fmt.Print("Input the food expenses and household expenses, week ", i + 1, ": ")
		fmt.Scanf("%f", &foodHousehold)
		totFoodHousehold += foodHousehold
		fmt.Print("Input the Common expenses", i + 1, ": ")
		fmt.Scanf("%f", &commonExpenses)
		totCommonExpenses += commonExpenses
	}
	return totFoodHousehold + totCommonExpenses
}

func MonthlyExpenses () float64 {
	var pubTransitPass, rent, bills float64
	fmt.Println("\nMonthly expenses")
	fmt.Print("Input the value of Public transit pass: ")
	fmt.Scanf("%f", &pubTransitPass)
	fmt.Print("Input the value of Rent: ")
	fmt.Scanf("%f", &rent)
	fmt.Print("Input the value of Total of monthly bills: ")
	fmt.Scanf("%f", &bills)
	return pubTransitPass + rent + bills
}

func Paycheques () float64 {
	var paycheque, totPaycheque float64
	totPaycheque = 0
	fmt.Println("\nTotal incomes")
	for i := 0; i < 2; i++ {
		fmt.Print("Input the value of Paycheque ", i + 1, ": ")
		fmt.Scanf("%f", &paycheque)
		totPaycheque += paycheque
	}
	return totPaycheque
}

func CalculatedGrade (grade1 float64, grade2 float64, grade3 float64){
	const laboratory float64 = 0.4
	const finalExam float64 = 0.35
	const midtermExam float64 = 0.25
	var finalGrade float64
	finalGrade = (grade1 * laboratory) + (grade2 * midtermExam) + (grade3 * finalExam)
	fmt.Println("\nThe final grade of a student is : ", fmt.Sprintf("%0.4f", finalGrade))
}

func CalculatedHeight (numberfloor int){
	var heightBuilding int
	const groundFloor int = 6
	const otherFloor int = 4
	heightBuilding = groundFloor + (otherFloor * (numberfloor -1))
	fmt.Print("The height of the building is " + strconv.Itoa(heightBuilding))
}
