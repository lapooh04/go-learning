//waitgroups: waiting for go routines to complete before proceeding
//sync package provides waitgroup
package main 

import (
	"fmt"
	"sync"
	"time"
)

//we can use wg.Add() to set the number of goroutines to wait for
//wg.Done() is called by each goroutine when it finishes (decrement the counter)
//wg.Wait() blocks the main function until the counter hits zero

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()  //this must be called even if the function panics
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)  //we launch one go routine in each iteration so we add it to the wait group
		go worker(i, &wg)

	}

	wg.Wait() //program will pause here until all the worker calls are over
	fmt.Println("All workers finished")
}