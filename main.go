package main

import (
	"fmt"
)

const (
	py = 13
	goScore = 16
	general = 11
)

func input(x []int, err error) []int {
	if err != nil {
		return x
	}
	var d int
	n, err := fmt.Scanf("%.2f", &d)
	if n == 1 {
		x = append(x,d)
	}
	return input(x, err)
}

func main() {
	fmt.Println("Enter input:")
	x := input([]int{}, nil)
	fmt.Println("input",x)
}