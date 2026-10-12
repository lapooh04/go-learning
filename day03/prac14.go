//methods in golang
//go doesnt have classes
//but methods are tied to a specific type (like a struct)
//that type is a receiver
//syntax: func <receiver> funcname(parametrs) returntype
package main 

import (
	"fmt"
	"math"
)

type Vertex struct {
	X,Y float64
}

type MyFloat float64 

//passing the receiver as pointer
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	//we refer to it using dot operator from the type
	fmt.Println(v.Abs())
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}



