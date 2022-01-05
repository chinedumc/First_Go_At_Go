package main 

import "fmt"

type Doctor struct {
	number int
	actorName string
	companions [] string
}

func main() {
	aDoctor := Doctor{
		number: 3,
		actorName: "Jon Pertwee",
		companions: [] string {
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[1])
	
	// Shorthand: No explicit/external struct type declaration made. Used in relatively few situations though
	// Anonymous struct
	aDoctors := struct{name string} {name: "Johns Pertwees"}
	fmt.Println(aDoctors)
	
	anotherDoctor := aDoctors
	anotherDoctor.name = "Tom Baker"
	fmt.Println(aDoctors)
	fmt.Println(anotherDoctor)

	// To make both aDoctors and anotherDoctor the same, use the & operator(the pointer)
	// anotherDoctor := &aDoctors
}