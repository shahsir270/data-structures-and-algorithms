// Factorial of a non-negative integer is the multiplication of all positive integers smaller than or equal to n. For example factorial of 6 is 6*5*4*3*2*1 which is 720.

package main

import "fmt"

// recursion type
func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

// iterative type
func fact(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}
func main() {
	fmt.Println(factorial(5))
	fmt.Println(fact(5))
}
