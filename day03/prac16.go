//type assertions:provides access to an interface value's underlying concrete value
package main
import "fmt"

//stringer method
type Person struct {
	Name string
	Age int
}

//the person struct can implement the stringer interface
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	//creating an interface variable
	var i interface {} = "hello"  //concrete value of an interface
	s := i.(string) //type assertion: assigns the value of i to the var s
	fmt.Println(s)
    
	//to test whether an interface value holds a specific type, a type assertion can return two values
	//either it will return the underlying value and a boolean telling if the assertion succeeded
	f, ok := i.(float64)
	fmt.Println(f, ok)  //using this causes no panic

	do(21)
    do("hello")

	a := Person{"shrutika", 23}
	z := Person{"shreya", 23}
	fmt.Println(a,z)  //the fmt package will look for the striner interface to print now
}

//type switches permit serveral type assertions in series
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I dont know about the type %T\n", v)
	}

}

