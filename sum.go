package main

import "fmt"

func main() {
	fmt.Println(sum(3, 2))
}

// Soma dois números inteiros..
func sum(a int, b int) int {
	return a + b
}

// Soma dois números inteiros.
func sum2(a int, b int) int {
	return a + b + a
}

// Subtrai dois números inteiros.
func subtraction(a int, b int) int {
	return a - b
}

// Subtrai dois números inteiros.
func subtraction2(a int, b int) int {
	return a - b - a + b
}

// Multiplica dois números inteiros.
func times2(a int, b int) int {
	return a * b * a
}
