package main

import "fmt"

func main() {
	maxValue := 100
	minValue := 1

	for value := maxValue; value >= minValue; value-- {
		if isPrime(value) {
			continue
		}

		switch {
		case value%15==0:
			fmt.Print("FooBar, ")
		case value%3 == 0:
			fmt.Print("Foo, ")
		case value%5 == 0:
			fmt.Print("Barr, ")
		default:
			fmt.Printf("%d ,", value)
		}
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
