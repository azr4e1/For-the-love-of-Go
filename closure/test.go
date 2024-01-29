package main

import "fmt"

func main() {
	funcs := []func(){}
	for _, v := range []int{1, 2, 3} {
		funcs = append(funcs, func() {
			fmt.Println(v)
		})
	}

	for _, f := range funcs {
		f()
	}
}
