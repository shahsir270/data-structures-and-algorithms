// Given an integer, write a function that returns true if the given number is palindrome, else false. For example, 12321 is palindrome, but 1451 is not palindrome.

package main

import "fmt"

func palindrome(n int64) bool {
	var reverse int64
	number := n
	for {
		if n == 0 {
			break
		}
		reverse = reverse*10 + n%10
		n = n / 10
	}
	if number == reverse {
		return true
	}
	return false
}
func main() {
	fmt.Println(palindrome(202020020202))
}
