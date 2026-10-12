//swtich statement (no explicit break statements)

package main

import (
	"fmt" 
	"runtime"
	"time"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n",os)
	}

	fmt.Println("when is sunday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	//switch with a condition
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Print("Good morning!")
	case t.Hour() == 12:
		fmt.Print("Good Noon!")
	case t.Hour() > 12:
		fmt.Print("Good evening!")
	}
}