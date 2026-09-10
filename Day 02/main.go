// Question 1: Library Management System
// Problem Statement
// A library wants to calculate the total cost of books.
// Write a Go program to:
// 1. Take the Book Name as input.
// 2. Take the Author Name as input.
// 3. Take the Book Price as input.
// 4. Take the Number of Copies as input.
// 5. Calculate the Total Cost.
// 6. Display all the details.

package main

import "fmt"

// func main() {
// 	var bookName, authorName string
// 	var price float64
// 	var copies int
// 	fmt.Print("Enter Book Name: ")
// 	fmt.Scan(&bookName)
// 	fmt.Print("Enter Author Name: ")
// 	fmt.Scan(&authorName)
// 	fmt.Print("Enter Book Price: ")
// 	fmt.Scan(&price)
// 	fmt.Print("Enter Number of Copies: ")
// 	fmt.Scan(&copies)
// 	totalCost := price * float64(copies)
// 	fmt.Println("\n----- Library Details -----")
// 	fmt.Println("Book Name:", bookName)
// 	fmt.Println("Author Name:", authorName)
// 	fmt.Println("Book Price:", price)
// 	fmt.Println("Copies:", copies)
// 	fmt.Println("Total Cost:", totalCost)
// }

//-----------------------------------------------------------------------------------------

// Question 2: Hotel Bill Generator
// Problem Statement
// A hotel wants to generate a customer's bill.
// Write a Go program to:
// 1. Take the Customer Name.
// 2. Take the Room Charge.
// 3. Take the Food Bill.
// 4. Take the Laundry Charge.
// 5. Calculate the Total Bill.
// 6. Calculate GST (12%).
// 7. Calculate the Final Amount to be Paid.
// None
// 8. Display the complete bill.

// func main() {
// 	var customerName string
// 	var roomCharge, foodBill, laundryCharge float64

// 	fmt.Print("Enter Customer Name: ")
// 	fmt.Scanln(&customerName)
// 	fmt.Print("Enter Room Charge: ")
// 	fmt.Scanln(&roomCharge)
// 	fmt.Print("Enter Food Bill: ")
// 	fmt.Scanln(&foodBill)
// 	fmt.Print("Enter Laundry Charge: ")
// 	fmt.Scanln(&laundryCharge)
// 	totalBill := roomCharge + foodBill + laundryCharge
// 	gst := totalBill * 0.12
// 	finalAmount := totalBill + gst

// 	fmt.Println("\n----- Hotel Bill -----")
// 	fmt.Println("Customer Name:", customerName)
// 	fmt.Println("Room Charge:", roomCharge)
// 	fmt.Println("Food Bill:", foodBill)
// 	fmt.Println("Laundry Charge:", laundryCharge)
// 	fmt.Println("Total Bill:", totalBill)
// 	fmt.Println("GST (12%):", gst)
// 	fmt.Println("Final Amount to be Paid:", finalAmount)
// }

// -----------------------------------------------------------------------------------------

// Question 3: Laptop Purchase Invoice
// Problem Statement
// A customer purchases a laptop from an electronics store.
// Write a Go program to:
// 1. Take the Customer Name.
// 2. Take the Laptop Brand.
// 3. Take the Laptop Price.
// 4. Calculate a Discount (15%).
// 5. Calculate GST (18%) on the discounted price.
// 6. Calculate the Final Amount to Pay.
// 7. Display the complete invoice.

// func main() {
// 	var customerName, laptopBrand string
// 	var laptopPrice float64

// 	fmt.Print("Enter Customer Name: ")
// 	fmt.Scanln(&customerName)
// 	fmt.Print("Enter Laptop Brand: ")
// 	fmt.Scanln(&laptopBrand)
// 	fmt.Print("Enter Laptop Price: ")
// 	fmt.Scanln(&laptopPrice)

// 	discount := laptopPrice * 0.15
// 	discountedPrice := laptopPrice - discount
// 	gst := discountedPrice * 0.18
// 	finalAmount := discountedPrice + gst

// 	fmt.Println("\n----- Laptop Purchase Invoice -----")
// 	fmt.Println("Customer Name:", customerName)
// 	fmt.Println("Laptop Brand:", laptopBrand)
// 	fmt.Println("Laptop Price:", laptopPrice)
// 	fmt.Println("Discount (15%):", discount)
// 	fmt.Println("Discounted Price:", discountedPrice)
// 	fmt.Println("GST (18%):", gst)
// 	fmt.Println("Final Amount to Pay:", finalAmount)
// }

// -----------------------------------------------------------------------------------------

// Question 4: Scholarship Eligibility System
// Problem Statement
// A university awards scholarships based on the following conditions:
// ● Percentage must be 85 or above.
// ● Attendance must be 75% or above.
// Scholarship Categories:
// ● Percentage ≥ 95 and Attendance ≥ 90 → 100% Scholarship
// ● Percentage ≥ 90 and Attendance ≥ 80 → 75% Scholarship
// ● Percentage ≥ 85 and Attendance ≥ 75 → 50% Scholarship
// ● Otherwise → Not Eligible
// Write a Go program to determine the scholarship.

// func main() {
// 	var percentage, attendance float64

// 	fmt.Print("Enter Percentage: ")
// 	fmt.Scanln(&percentage)
// 	fmt.Print("Enter Attendance (%): ")
// 	fmt.Scanln(&attendance)

// 	if percentage >= 95 && attendance >= 90 {
// 		fmt.Println("Congratulations! You are eligible for a 100% Scholarship.")
// 	} else if percentage >= 90 && attendance >= 80 {
// 		fmt.Println("Congratulations! You are eligible for a 75% Scholarship.")
// 	} else if percentage >= 85 && attendance >= 75 {
// 		fmt.Println("Congratulations! You are eligible for a 50% Scholarship.")
// 	} else {
// 		fmt.Println("Sorry, you are not eligible for any scholarship.")
// 	}
// }

// -----------------------------------------------------------------------------------------

// Question 5: Employee Bonus Calculator
// Problem Statement
// A company provides bonuses to its employees based on thei
// years of experience and performance rating.
// Write a Go program to:
// 1. Take the Employee Name as input.
// 2. Take the Basic Salary as input.
// 3. Take the Years of Experience as input.
// 4. Take the Performance Rating (1 to 5) as input.
// 5. Calculate the bonus based on the following rules:
// * Experience ≥ 10 years and Rating ≥ 4.5 + 30% Bonus
// * Experience ≥ 5 years and Rating ≥ 4 20% Bonus
// * Experience ≥ 2 years and Rating ≥ 3 10% Bonus
// * Otherwise No Bonus
// 6. Calculate the Total Salary after adding the bonus.
// 7. Display:
// * Employee Name
// * Basic Salary
// * Years of Experience
// * Performance Rating
// * Bonus Amount
// * Total Salary

func main() {
	var employeeName string
	var basicSalary float64
	var yearsOfExperience, performanceRating int

	fmt.Print("Enter Employee Name: ")
	fmt.Scanln(&employeeName)
	fmt.Print("Enter Basic Salary: ")
	fmt.Scanln(&basicSalary)
	fmt.Print("Enter Years of Experience: ")
	fmt.Scanln(&yearsOfExperience)
	fmt.Print("Enter Performance Rating (1 to 5): ")
	fmt.Scanln(&performanceRating)

	var bonus float64
	if yearsOfExperience >= 10 && performanceRating >= 4 {
		bonus = basicSalary * 0.30
	} else if yearsOfExperience >= 5 && performanceRating >= 4 {
		bonus = basicSalary * 0.20
	} else if yearsOfExperience >= 2 && performanceRating >= 3 {
		bonus = basicSalary * 0.10
	}
	totalSalary := basicSalary + bonus
	fmt.Println("\nEmployee Name:", employeeName)
	fmt.Println("Basic Salary:", basicSalary)
	fmt.Println("Experience:", yearsOfExperience)
	fmt.Println("Rating:", performanceRating)
	fmt.Println("Bonus:", bonus)
	fmt.Println("Total Salary:", totalSalary)
}

// -----------------------------------------------------------------------------------------

// Question 6: Branch Allocation System
// Problem Statement
// A university allocates branches based on student performance.
// Rules:
// ● Percentage ≥ 60
// ● Entrance Score ≥ 50
// Branch Allocation:
// ● Percentage ≥ 95 and Entrance ≥ 90 → Computer Science
// ● Percentage ≥ 90 and Entrance ≥ 85 → Artificial Intelligence
// ● Percentage ≥ 85 and Entrance ≥ 80 → Information Technology
// ● Percentage ≥ 75 and Entrance ≥ 70 → Electronics
// ● Otherwise → Mechanical
// If not eligible but Sports Quota is true → Sports Science.
// Otherwise → Admission Rejected.

// func main() {
// 	var percentage, entranceScore float64
// 	var sportsQuota bool

// 	fmt.Print("Enter Percentage: ")
// 	fmt.Scanln(&percentage)
// 	fmt.Print("Enter Entrance Score: ")
// 	fmt.Scanln(&entranceScore)
// 	fmt.Print("Enter Sports Quota (true/false): ")
// 	fmt.Scanln(&sportsQuota)

// 	if percentage >= 95 && entranceScore >= 90 {
// 		fmt.Println("Branch: Computer Science")
// 	} else if percentage >= 90 && entranceScore >= 85 {
// 		fmt.Println("Branch: Artificial Intelligence")
// 	} else if percentage >= 85 && entranceScore >= 80 {
// 		fmt.Println("Branch: Information Technology")
// 	} else if percentage >= 75 && entranceScore >= 70 {
// 		fmt.Println("Branch: Electronics")
// 	} else if sportsQuota {
// 		fmt.Println("Branch: Sports Science")
// 	} else {
// 		fmt.Println("Admission Rejected")
// 	}
// }
