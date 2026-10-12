//structs in golang: collection of fields
//access struct fields outside using a dot operator
package main 

import "fmt"

type Vertex struct {
	X int
	Y int
}

//struct literals
var (
	v1 = Vertex{1,2} 
	v2 = Vertex{X: 1} //y is 0 by default
	v3 = Vertex{} //both are 0
	p = &Vertex{1,2}
)

func main() {
	fmt.Println(Vertex{1,2})
	v := Vertex{1,2}
	v.X = 4
	fmt.Println(v.X)

	//we can also use struct pointers to access a struct field
	p := &v
	p.Y = 9
	fmt.Println(p.Y)

	fmt.Println(v1, v2, v3, p)
}