package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
	sayMessage("Hello Go!", i)
	add()
	div()
	anon()
}
}

func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is ", idx)
}

func add() {
	sum(1, 2, 3, 4, 5)
}

func sum(values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("The sum is ", result)
}

func div() {
	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

// ANONYMOUS FUNCTION
func anon() {
	func () {
		fmt.Println("Anonymous func")
	}()
}