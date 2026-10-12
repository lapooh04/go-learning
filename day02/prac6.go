//defer statement: function call is not executed until the surrounding function returns
//stacking defers: all the defer function calls are pushed into stack (and popped in LIFO)

package main

import "fmt"

func main() {
	defer fmt.Println("world")
	fmt.Println("Hello")

	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}