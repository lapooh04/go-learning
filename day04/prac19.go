//channels and goroutines
//go routines are lightweight threads managed by go runtime
//channels help to synchronize different go routines by sending and receiving values
// c <- and <- c (send to c and recv from c)
//we create and initialize them before use
package main  

import "fmt"
import "time"

func sum(s []int, c chan int) {
	sum := 0
	for _,v := range s {
		sum += v 
	}

	//sending sum to channel c 
	c <- sum
}

//we can also use select statement
//select is like switch but for channel operations (each case is a send or recv channel operation)
//it is a way of implementing timeout for channels
//so we dont have to wait forever for a goroutine. whichever case is ready first is the one that is executed.

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	//run the sum function as two go routines
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	//recv from c the sums of both halves
	x, y := <-c, <-c 

	fmt.Println(x, y, x+y)

	ch1 := make(chan string)
	//go routine
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Database query result"
	}()

	select {
	case res := <-ch1:
		fmt.Println("Received:", res)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout! The database took too long")
	}
}