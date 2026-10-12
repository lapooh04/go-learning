//interfaces: it is a type
//it provides the signatures of methods but no implementation
//a value of interface can hold any value that implements those methods
//an empty interface is an interface that has zero methods
package main

import (
	"fmt"
	//"math"
)

//a way of describing behavior
type Speaker interface {
	Speak() string
}

//concrete types that will implement the methods declared in the interface
type Dog struct{}
type Robot struct{}

func (d Dog) Speak() string {
	return "woof!"
}

func (r Robot) Speak() string {
	return "beep beep!"
}

func main() {
	r := Robot{}
	d := Dog{}
	fmt.Println(r.Speak())
    fmt.Println(d.Speak())
}