package basics

import (
	"errors"
	"fmt"
)

func Big(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func Eggs() {
	eggs := 42

	fmt.Println(eggs)

	eggs = 50

	fmt.Println(eggs)
}

func XOR(a, b bool) bool {
	if a && b {
		return false
	}
	if (!a) && (!b) {
		return false
	}

	return true
}

func Greet(name string) string {
	switch name {
	case "Alice":
		return "Hey, Alice"
	case "Bob":
		return "What's up Bob?"
	default:
		return "Hello, stranger"
	}
}

func Total(list []int) int {
	var total int

	for _, el := range list {
		total += el
	}

	return total
}

func Evens() {
	for x := 0; x <= 100; x++ {
		if x%2 == 0 {
			fmt.Println(x)
		}
	}
}

func NonNegative(numbers []int) {
	for _, el := range numbers {
		if el < 0 {
			continue
		}
		fmt.Println(el)
	}
}

func Withdraw(balance, withdraw int) (int, error) {
	value := balance - withdraw
	if value < 0 {
		return 0, errors.New("The withdraw would leave you overdrawn")
	}

	return value, nil
}
