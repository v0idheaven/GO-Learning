package main

import "fmt"

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

// func main() {
// 	length := 10
// 	width := 5

// 	fmt.Println("Area:", length*width)
// }

func main() {
	var student_name string
	var marks_math int
	var marks_physics int
	var marks_chemistry int
	var marks_biology int
	var marks_english int

	fmt.Print("Enter student name: ")
	fmt.Scanln(&student_name)
	fmt.Print("Enter marks for Math: ")
	fmt.Scanln(&marks_math)
	fmt.Print("Enter marks for Physics: ")
	fmt.Scanln(&marks_physics)
	fmt.Print("Enter marks for Chemistry: ")
	fmt.Scanln(&marks_chemistry)
	fmt.Print("Enter marks for Biology: ")
	fmt.Scanln(&marks_biology)
	fmt.Print("Enter marks for English: ")
	fmt.Scanln(&marks_english)
	total_marks := marks_math + marks_physics + marks_chemistry + marks_biology + marks_english
	fmt.Println("Total Marks:", total_marks)
	fmt.Println("Percentage:", float64(total_marks)/500*100, "%")
}
