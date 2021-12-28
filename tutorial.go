package main

import "fmt"

func main() {
	fmt.Println("Hello! ")

	var name string
	// var ok string
	var age uint
	var answer string
	var answer2 string
	
	fmt.Printf("Please enter your name! \n")
	fmt.Scan(&name)
	fmt.Printf("Welcome to the quiz %v \n", name) 

	fmt.Printf("Please enter your age: \n") 
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Please continue")
	} else {
		fmt.Printf("Sorry %v, you are not qualified to play", name)
		return
	}

	fmt.Printf("Which team is better, Chelsea or Arsenal? ")
	fmt.Scan(&answer)

	if answer == "Chelsea" || answer == "chelsea" || answer == "CHELSEA" {
		fmt.Printf("Yes! %v is the correct answer!", answer)
	} else {
		fmt.Println("No way! You missed it!")
		return
	}

	fmt.Printf("Which is easier, Driving or Sleeping?")
	fmt.Scan(&answer2)

	if answer2 == "sleeping" || answer2 == "Sleeping" || answer2 == "SLEEPING" {
		fmt.Printf("Yes! %v is the correct answer!", answer2)
	} else {
		fmt.Println("No way! You missed it!")
		return
	}
	// fmt.Println("Enter OK when you are ready to begin!")
	// fmt.Scan(&ok)

	// if ok == "ok" || ok == "Ok" || ok == "OK" {
	// 	fmt.Println("Question 1")
	// } else {
	// 	fmt.Println("Enter OK when you are ready to begin!")
	// 	fmt.Scan(&ok)
	// }
}