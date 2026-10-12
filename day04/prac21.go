//generics in go
//they are also known as type parameters
//we define type parameters inside square brackets before the function arguments
package main 

import "fmt"

//syntax for generic is: func <func-name> [K comparable, V any] (arguments use K and V instead of type) <return type> {}
func GetKeys[K comparable, V any](m map[K]V) []K {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

func main() {
	stringMap := map[string]int{"cpu":80, "memory":60}
	intMap := map[int]string{404:"Not Found", 500: "Server Error"}
	fmt.Println(GetKeys(stringMap))
	fmt.Println(GetKeys(intMap))
}