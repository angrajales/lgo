package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func twoval(ival *int) {
	*ival = 2
}

func main() {

	i := 1
	fmt.Println(i)

	zeroval(i)
	fmt.Println(i)

	twoval(&i)
	fmt.Println(i)

	fmt.Println(&i)
}
