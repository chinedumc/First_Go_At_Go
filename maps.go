package main

import "fmt"

func main() {
	statePopulations := map[string] int {  // using 'make':  statePopulations := make(map[string]int)
		"California": 392500,
		"Texas": 278625,
		"Florida": 206124,
	}
	fmt.Println(statePopulations)
	
	statePopulations["Georgia"] = 103103 // adds Georgia
	fmt.Println(statePopulations) // Prints out everything with Georgia added
	
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)

	//Use the comma-ok 
	pop, ok := statePopulations["Texs"]
	fmt.Println(pop, ok)

	//  same check as: 
	// _, ok := statePopulations["Texs"]
	// fmt.Println(ok)

	sp := statePopulations
	delete(sp, "Texas")
	fmt.Println(sp)
	fmt.Println(statePopulations)
}