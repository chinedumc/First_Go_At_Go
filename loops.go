package main

import "fmt"

func main() {
	for i, j := 0, 0; i < 5; i, j = i +1, j+1 {
		fmt.Println(i, j)
	}

	s := [] int {1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}

	str := "hello Me"
	for _, v := range str {
		fmt.Println( string(v))
	}
}