//exercise : maps
//word count function
//we use strings package to use the fields function to split function for string
package main

import (
	"fmt"
	"strings"
)


func main() {
	//another way of function declaration
	wc := func (s string)map[string]int {
	// create the map to use
	mymap := make(map[string]int)
	// initialize the slice we will be using for storing the words of the string
	slc := make([]string, 0)
    slc = strings.Fields(s)
	// iterate over the slice and increase its frequency in the map
    for _, v := range slc {
		mymap[v]++
	}

	return mymap
}
	fmt.Println(wc("I love Go Lets Go"))
}