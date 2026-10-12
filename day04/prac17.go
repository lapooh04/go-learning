//exercise on stringer interface
package main 

import "fmt"

type IPAddr [4]byte 

//implementing the stringer interface
func (i IPAddr) String() string{
	return fmt.Sprintf("%d.%d.%d.%d", i[0], i[1], i[2], i[3])
}

func main() {
	hosts := map[string]IPAddr{
		"google" : {8, 8, 8, 8},
		"localhost" : {127, 0, 0, 1},
	}

	for domain, ip := range hosts {
		fmt.Printf("%v : %v\n", domain, ip)
	}
}
