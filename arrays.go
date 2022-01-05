package main

import "fmt" 

func main() {
	grades := [...] int {97, 85, 9, 3}
	fmt.Printf("Grades: %v\n", grades)
	fmt.Println(len(grades))

	var identityMatrix [3] [3] int = [3] [3] int { [3]int {1, 0, 0},  [3]int {0, 1, 0},  [3]int {0, 0, 1}}
	fmt.Println(identityMatrix)

	a := [] int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[ : ] //slice of all elements
	c := a[ 3: ] //slice of 4th element to end of array
	d := a[ :6 ] //slice of first 6 elements
	e := a[ 3:6 ] //slice of 4th to the 6th element

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	f := make([] int, 3)
	fmt.Println(f)
	fmt.Printf("Length: %v\n", len(f))
	fmt.Printf("Capacity: %v\n", cap(f)) 

	g :=  [] int{1, 2, 3, 4, 5 }
	fmt.Println(g)
	h := append(g[ : 2], g[3: ]...)  // to remove the 3rd element
	fmt.Println(h)
	fmt.Println(g)
	i := append(g[ : 2], g[3: ]...)  // to remove the 3rd element
	fmt.Println(i)
	fmt.Println(g)

	j := [] int{6, 7, 8, 9, 10}
	fmt.Println(j)
	k := append(j[ 1: ])
	fmt.Println(k)
	fmt.Println(j)
	l := j[: len(j)-1]
	fmt.Println(l)
	fmt.Println(j)

}

