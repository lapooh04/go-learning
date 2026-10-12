package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
	i int
	j int = 42
)

const Pi = 3.14

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("Type: %T Value: %v\n", i, i)
	fmt.Printf("Before type conversion, type: %T, value: %v\n", j, j)
    f := float64(j)
	fmt.Printf("After type conversion, type: %T, value: %v\n", f, f)
	fmt.Println("Happy", Pi, "Day")
}

//basic types in go
//varibales declared without an initial value are given their zero values (0 for int, false for bool and "" for empty string)
//type conversion 
//constants (cannot be declared using := syntax)