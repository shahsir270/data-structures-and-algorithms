package main

import "fmt"

func count(n int64) int {
	var c int
	for {
		if n == 0 && c == 0 {
			c = 1
			break
		} else if n == 0 {
			break
		}
		c++
		n = n / 10
	}
	return c
}

func main() {
	fmt.Println(count(20))
}
