//function closures
// a closure happens when an inner function references a variable from an outer function
//in go a function can actually return another function
//the inner function remembers the state of the variable and each time we call it it remembers and can update

package main 

import "fmt"

//this is a function that returns a function (a function that takes as integer and returns an integer)
func adder() func(int) int {
	sum := 0
	//this function references a variable outside of its own body
	return func(x int) int {
		sum += x
		return sum
	}
}

//the outer function will set up the initial state and the inner function will update the state each time
func fibonacci() func() int{
	//n = (n-1) + (n-2)
	a := 0
	b := 1
	//this inner function saves the state of the outside variables
	return func() int {
		current := a
		a, b = b, a+b

		return current
	}
}

func createCounter() func() int {
	cnt := 0
	return func() int {
		cnt++
		return cnt
	}
}

func main() {
	fmt.Println("closure function example")
	pos, neg := adder(), adder()
	//each time this function is called it will remember its state
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
    
	fmt.Println("fibonacci sequence")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	fmt.Println("simple counter using function closures")
	c := createCounter()
	for i := 0; i < 5; i++ {
		fmt.Println(c()) //since this is a function closure it will remember its state in every call
	}
}