package main

import "fmt"

func main() {
	var s []int // slice with element of type int
	s = []int{2, 3, 4, 5, 6}
	subset := s[0:3] // {2,3,4} - subset은 array가 아니고 slice 임
	fmt.Println(subset)
}
