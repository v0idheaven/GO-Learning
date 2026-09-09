package main

import "fmt"

// -----------------------------------------------------------------------------------------
// Question 1: Employee Details
// Create variables for:
// ● Employee Name
// ● Employee ID
// ● Department
// ● Salary
// Print all the details.

// func main() {
// 	fmt.Println("Hello, World!")
// }

// func main() {
// 	employee_name := "John Doe"
// 	employee_age := 30
// 	employee_salary := 50000.0

// 	fmt.Println("Name:", employee_name)
// 	fmt.Println("Age:", employee_age)
// 	fmt.Println("Salary:", employee_salary)
// }

// -----------------------------------------------------------------------------------------

// Question 2: Rectangle Calculator
// Create variables for the length and width of a rectangle.
// Calculate and print:
// ● Area
// ● Perimeter

// func main() {
// 	length := 10
// 	width := 5

// 	fmt.Println("Area:", length*width)
// }

// // -----------------------------------------------------------------------------------------

// Question 3:Shopping Bill Calculator
// Problem Statement
// A customer purchases three products from a store.
// Write a Go program to:
// 1. Create three variables to store the prices of the products.
// 2. Calculate the Total Bill.
// 3. Calculate the GST (18%) on the total bill.
// 4. Calculate the Final Amount to be paid

// func main() {
// 	price1 := 500.0
// 	price2 := 1000.0
// 	price3 := 750.0

// 	totalBill := price1 + price2 + price3

// 	gst := totalBill * 0.18

// 	finalAmount := totalBill + gst

// 	fmt.Println("Price of Product 1:", price1)
// 	fmt.Println("Price of Product 2:", price2)
// 	fmt.Println("Price of Product 3:", price3)
// 	fmt.Println("Total Bill:", totalBill)
// 	fmt.Println("GST (18%):", gst)
// 	fmt.Println("Final Amount:", finalAmount)
// }

// -----------------------------------------------------------------------------------------

// Question 4:
// A company wants to calculate the Gross Salary of an employee.
// Write a Go program to:
// 1. Create a variable to store the Basic Salary.
// 2. Calculate the House Rent Allowance (HRA) as 20% of the Basic Salary.
// 3. Calculate the Dearness Allowance (DA) as 10% of the Basic Salary.

// func main() {
// 	basicSalary := 30000.0

// 	hra := basicSalary * 0.20

// 	da := basicSalary * 0.10

// 	grossSalary := basicSalary + hra + da

// 	fmt.Println("Basic Salary:", basicSalary)
// 	fmt.Println("HRA (20%):", hra)
// 	fmt.Println("DA (10%):", da)
// 	fmt.Println("Gross Salary:", grossSalary)
// }

// -----------------------------------------------------------------------------------------

// Question 5: Student Percentage Calculator
// Write a Go program to:
// ● Take the student's name as input.
// ● Take marks of 5 subjects.
// ● Calculate the total marks.
// ● Calculate the percentage.
// ● Display the result.

// func main() {
// 	var name string
// 	var m1, m2, m3, m4, m5 float64

// 	fmt.Print("Enter student's name: ")
// 	fmt.Scanln(&name)

// 	fmt.Print("Enter marks of 5 subjects: ")
// 	fmt.Scan(&m1, &m2, &m3, &m4, &m5)

// 	total := m1 + m2 + m3 + m4 + m5

// 	percentage := total / 5

// 	fmt.Println("\nStudent Name:", name)
// 	fmt.Println("Total Marks:", total)
// 	fmt.Println("Percentage:", percentage, "%")
// }

// -----------------------------------------------------------------------------------------

// Question 6: Simple Interest Calculator
// Write a program that takes:
// ● Principal Amount
// ● Rate of Interest
// ● Time (Years)
// Calculate the Simple Interest using:
// SI = (P × R × T) / 100
// Print the result.

func main() {
	var principal, rate, time float64

	fmt.Print("Enter Principal Amount: ")
	fmt.Scan(&principal)

	fmt.Print("Enter Rate of Interest: ")
	fmt.Scan(&rate)

	fmt.Print("Enter Time (Years): ")
	fmt.Scan(&time)

	si := (principal * rate * time) / 100

	fmt.Println("Simple Interest:", si)
}
