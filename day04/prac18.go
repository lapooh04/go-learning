//error is also a built in interface in go (like stringer)
//fmt package looks for this interface when prniting values
package main 

import (
	"fmt"
	"time"
	"io"
	"strings"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it did not work",
	}
}

//reader is also an interface from the io interface which has a method called read

 
// a nil error denotes success and a non-nil error denotes failure
func main() {

	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8) //where the data will be read into (8 bytes at a time)
	for {
		n, errd := r.Read(b)  //using the read method from the reader interface
		fmt.Printf("n = %v err = %v b = %v\n", n, errd, b)
		if errd == io.EOF {
		break
	    }   //it stops when the byte stream ends
	}
	
	if err := run(); err != nil {
		fmt.Println(err)
	}
}