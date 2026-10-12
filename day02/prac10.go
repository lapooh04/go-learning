package main

import "fmt"

var pow = []int{1,2,4, 8, 16, 32, 64, 128}

//using a range instead of a for loop for iterating over slices and maps
//range always returns index and value where iterating over 
// we can use blank identifier "_", if we wish to not use any of the values
func main() {
	//it will return index and value
	for i,v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

}