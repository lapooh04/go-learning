//pointers: memory addr of a value
package main

import "fmt"

func main() {
	i, j := 42, 2701
	p := &i  //a pointers pointing to variable i
	fmt.Println(*p) //reading i through the pointer
	*p = 21 //set i through the pointer
	fmt.Println("i is now:", i)

	p = &j //p now points to j
	*p = *p/37  //divide j through pointer
	fmt.Println("new values of j:", j)
}