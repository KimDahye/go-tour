package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	for i := 0; i < 1; i++ {
		go send(ch)
	}
	fmt.Println(<-ch)
}

func send(ch chan int) {
	ch <- 0
}
