//using maps: key-val pairs
//map[key]value
//using make function -> mp = make(map[string]int)
 
package main 

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	//using make function to create a map of a given type that is initialized and ready to use
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	//map literals -> maps with hardcoded values
	mymap := map[string]Vertex{
		"Bell Labs": Vertex{40.68433, -74.39967},
		"Google": Vertex{37.42202, -122.08408},
	}

	fmt.Println(mymap)
	  
	elements := map[string]string{} //empty map literal
	elements["H"] = "Hydrogen" 
	fmt.Println(elements)

	//insert or update an element in map
	elements["O"] = "Oxygen"
	//retriving an element from map
	elem := elements["O"]
	fmt.Println(elem)
	//delete an element : delete(mapname, key)
	delete(elements, "O")
	fmt.Println(elements)
}