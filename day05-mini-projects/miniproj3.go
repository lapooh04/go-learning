//mini project 3
//CLI typing test that displays a prompt, captures user input and measures the typing time, speed and accuracy
//use a slice for storing the prompt and select a random index using math/rand
//will use a bufio reader attached to os.stdin (as normal fmt.Scan stops taking input when space is hit)
//capture the time as the input begins, we trim the whitespace
//calculate the words per minute, obtian a percentage to get accuracy
package main 

import (
	"fmt"
	"bufio"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	prompts := []string{
		"Go is an open source programming language designed to make it easy to build simple, reliable, and efficient software.",
		"Concurrency is not parallelism, it is a way to structure a program by breaking it into independently executing components.",
		"The Go standard library is powerful enough to build robust web servers with minimal external dependencies.",
		"Channels are used to communicate between goroutines, allowing safe data sharing.",
	}

	//selecting a random prompt: using math/rand we can pick a random index from the prompts slice
	randomPrompt := rand.Intn(len(prompts))
	targetText := prompts[randomPrompt]

	fmt.Println("Terminal Typing Speed Test")
	fmt.Println("Type the following text exactly as shown")
    
	//using ANSI escape code for cyan color
	fmt.Printf("\033[36m%s\033[0m\n\n", targetText)
	fmt.Println("Press ENTER when ready")
    
	//creating a new bufio reader from standard input
	inputReader := bufio.NewReader(os.Stdin)

	inputReader.ReadString('\n')

	fmt.Println("Start typing.....")
    
	//start the timer
	startTime := time.Now()

	userInput, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured: ", err)
		os.Exit(1)
	}
    
	//stop the timer and calc duration
	duration := time.Since(startTime)
	

	//clean up the input to remove hidden \n or characters
	userInput = strings.TrimSpace(userInput)
	
    //calculating the word per minute
	minutes := duration.Minutes()
	charsTyped := len(userInput)
	wpm := 0.0

	if minutes > 0 {
		wpm = (float64(charsTyped)/5.0)/minutes
	}

	//calculate accuracy by comparing arrays of words
	targetWords := strings.Split(targetText, " ")
	userWords := strings.Split(userInput, " ")
	correctWords := 0
	for i := 0; i < len(userWords) && i < len(targetWords); i++ {
		if userWords[i] == targetWords[i] {
			correctWords++
		}
	}

	accuracy := 0.0
	if len(targetWords) > 0 {
		accuracy = (float64(correctWords)/float64(len(targetWords))) * 100
	}

	//output the results on the dashboard
	fmt.Println("\n----------------------------------------")
    fmt.Println("             TEST RESULTS               ")
    fmt.Println("----------------------------------------")
    fmt.Printf("Time Elapsed : %.2f seconds\n", duration.Seconds())
    fmt.Printf("Gross WPM    : %.0f words/minute\n", wpm)
    fmt.Printf("Accuracy     : %.1f%%\n", accuracy)
    fmt.Printf("Words Correct: %d / %d\n", correctWords, len(targetWords))
    fmt.Println("========================================")

}
