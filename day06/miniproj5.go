//mini project 5
//code clock cli
//building a cli application that tracks the amount of time spent on tasks, starts a timer, stops the times, saves the session to a json file
//prompts the user with a task, initalize a reader, start the clock, pause the program when prompted
//parse the json content and load into slice, write it to the disk 
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)
//define the blueprint for the json data to capture
type Session struct {
	Task string `json:"task"`
	Date string `json:"date"`
	Duration string `json:"duration"`
}

const dataFile = "codeclock.json"

//function to save the session into the file
func saveSession(newSession Session) {
	var sessions []Session

	//reading the file to get existing data
	fileData, err := os.ReadFile(dataFile)

	//if err is nil then the file exists we have to unmarshall old data into our slice
	if err == nil {
		json.Unmarshal(fileData, &sessions)
	}

	//append the new session to the list
	sessions = append(sessions, newSession)

	//convert the slice back to the json format
	updatedJSON, _ := json.MarshalIndent(sessions, "", " ")

	//saving the json file to the hard drive with permissions
	err = os.WriteFile(dataFile, updatedJSON, 0644)

	if err != nil {
		fmt.Printf("error saving file: %v\n", err)
	} else {
		fmt.Printf("session saved to %s!\n", dataFile)
	}
}

func main() {
	fmt.Println("============")
	fmt.Println("Code Clock")
	fmt.Println("==========")

	//start the clock
	fmt.Print("What task are you starting? : ")

	//create a new reader to read data from standard input and assign it to a variable
	inputReader := bufio.NewReader(os.Stdin)

	taskInput, _ := inputReader.ReadString('\n')
	//clean up the input to remove hidden \n or characters
	taskName := strings.TrimSpace(taskInput)
    
	//start the timer
	startTime := time.Now()
    
	fmt.Printf("\nClocked IN at %s\n", startTime.Format("15:04:05"))
	fmt.Println("Press ENTER when you are finished...")

	inputReader.ReadString('\n')

	//stop the timer and calc duration
	exactDuration := time.Since(startTime)
	duration := exactDuration.Round(time.Second)
	fmt.Printf("Clocked OUT! Time Spent: %s\n", duration)
	
	//create new session record
	newSession := Session{
		Task: taskName,
		Date: time.Now().Format("2006-01-02"),
		Duration: duration.String(),
	}

	//saving the session data to the file
	saveSession(newSession)
}
