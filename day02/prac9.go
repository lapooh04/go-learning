//arrays in go: size and type and written together (cant be resized)
//slices: a flexible and dynamic way to view into elements of an array
// slice in go: a[low : high] , where low is inclusive but high is excluded
// var s []int is a slice of the type int
//array literal and slice literal: we provide the hardcoded values with the initialization

package main

import "fmt"

//declaring a slice
var slice []int

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
    
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	slice = primes[2:5]  //using a slice
	fmt.Println(slice)

	//changing the value of a slice also changes the corresponding element of the underlying array
	//any other slice sharing the same underlying array will also change 
	names := [5]string{"john", "bob", "alice", "ken", "barbie"}
	fmt.Println(names)
	var nameslice []string = names[0:3]
	nameslice[1] = "ted"
	fmt.Println(names)

	//array literal and slice literal (slice literal is an array literal without the specified size)
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
    
	//slice defaults
	q = q[:4]
	fmt.Println(q)

	//length and capacity of slice
	fmt.Println(len(q))
	fmt.Println(cap(q))

	//using make to create a slice
	//make function is used to allocate memory and create dynamic arrays
	//make([]type, len(b), cap(b))
	b := make([]int, 5)
	fmt.Println(b)

	//appending to a slice
	//append function: s = append(slice of type T, values to append)
	b = append(b, 2, 3, 5)
	fmt.Println(b)
}