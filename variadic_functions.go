package main

import "fmt"

func sum(args ...int) int {
	fmt.Println(args)
	res := 0
	for i := 0; i < len(args); i++ {
		res = res + args[i]
	}
	return res
}

func main() {
	fmt.Println(sum(7, 1, -3, 4, 8, 7, 4))
}
