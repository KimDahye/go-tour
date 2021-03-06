package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x1 := 0
	x2 := 1
	return func() int {
		res := x1
		x3 := x1 + x2
		x1 = x2
		x2 = x3
		return res
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
