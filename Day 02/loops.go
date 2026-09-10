// Question: Sum of Numbers from 1 to N
// Problem Statement
// Take a number N and calculate the sum of numbers from 1 to N.

package main

// import "fmt"

// func main() {
// 	var n, sum int
// 	fmt.Print("Enter a number: ")
// 	fmt.Scan(&n)

// 	for i := 1; i <= n; i++ {
// 		sum += i
// 	}

// 	fmt.Println("Sum of numbers from 1 to", n, "is:", sum)
// }

// ------------------------------------------------------------------------------------------

// Question 8: Sum of Even Numbers
// Problem Statement
// Take a number N and calculate the sum of all even numbers from 1 to N

// func main() {
// 	var n, sum int
// 	fmt.Print("Enter a number: ")
// 	fmt.Scan(&n)

// 	for i := 2; i <= n; i += 2 {
// 		sum += i
// 	}
// 	fmt.Println("Sum of even numbers from 1 to", n, "is:", sum)
// }

// ------------------------------------------------------------------------------------------

// Question: Print Even and Odd Numbers
// Problem Statement
// Write a Go program to:
// 1. Take a positive integer N.
// 2. Print all even numbers.
// 3. Print all odd numbers.

// func main() {
// 	var n int
// 	fmt.Print("Enter N: ")
// 	fmt.Scan(&n)
// 	fmt.Println("Even Numbers:")
// 	for i := 2; i <= n; i += 2 {
// 		fmt.Print(i, " ")
// 	}
// 	fmt.Println("\nOdd Numbers:")
// 	for i := 1; i <= n; i += 2 {
// 		fmt.Print(i, " ")
// 	}
// }

// ------------------------------------------------------------------------------------------

// Question: Multiplication Table
// Generator
// Problem Statement
// Write a Go program to:
// 1. Take a number.
// 2. Print its multiplication table from 1 to 10.

// func main() {
// var n int
// fmt.Print("Enter Number: ")
// fmt.Scan(&n)
// for i := 1; i <= 10; i++ {
// fmt.Printf("%d x %d = %d\n", n, i, n*i)
// }
// }

// ------------------------------------------------------------------------------------------

// Question: Login Attempt System
// Problem Statement
// A website allows only 3 login attempts.
// Write a Go program to:
// 1. Store the correct password as 1234.
// 2. Ask the user to enter the password.
// 3. Give a maximum of 3 attempts.
// 4. If the password is correct, print Login Successful.
// 5. Otherwise, after 3 incorrect attempts, print Account Locked.

// func main() {
// const correctPassword = 1234
// var password int
// for attempt := 1; attempt <= 3; attempt++ {
// fmt.Print("Enter Password: ")
// fmt.Scan(&password)
// if password == correctPassword {
// fmt.Println("Login Successful")
// return
// } else {
// fmt.Println("Incorrect Password")
// }
// }
// fmt.Println("Account Locked")
// }
