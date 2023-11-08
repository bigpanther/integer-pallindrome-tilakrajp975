package main

import "fmt"

func IsPalindrome(input int) bool {
	// Do your implementation here
	// reverse
	// equate

	//2*10+1
	//1*10+2

	//100*2+10*1+5
	//100*5+10*1+2

	i := 1
	newNumber := 0
	for input%i != input {
		newNumber = (newNumber * 10) + ((input / i) % 10)
		i = i * 10
	}
	// split digits using modulus
	// write digits in reverse order
	// equate
	return newNumber == input
}

func main() {
	fmt.Println(123) // false
	fmt.Println(121) // true
	fmt.Println(1)   // true
	fmt.Println(11)  // true
	fmt.Println(12)  // false
}
