//for loops : the init and the post statement are optional
// an inifinte loop in go is an empty for loop with no conditions at all
// for both if and for loop the brackets are not required for the conditional statements but the curly braces are needed
// variabled declared in the if or for loop are only scoped to that loop
// in go we can also start the if condition with a small statement to execute before the condition
package main

import "fmt"
import "math"

func sqrt( x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v:= math.Pow(x,n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
    sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
    
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	fmt.Println(sqrt(2),sqrt(-4))

	fmt.Println(pow(3, 2, 10),pow(3, 3, 20))
}
