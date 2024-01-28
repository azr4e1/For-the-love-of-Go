package main

import (
	"basics"
	"fmt"
)

func main() {
	a, b := 2, 3
	x, y := true, false
	fmt.Println(basics.Big(a, b))
	basics.Eggs()
	fmt.Println(basics.XOR(x, y))
	fmt.Println(basics.Greet("Alice"))
	fmt.Println(basics.Total([]int{1, 2, 3}))
	basics.Evens()
	basics.NonNegative([]int{-2, 3, -4, 5})
	fmt.Println(basics.Withdraw(300, 200))
	fmt.Println(basics.Withdraw(200, 300))
}
