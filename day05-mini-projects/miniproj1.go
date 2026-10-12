//mini project 1: cloud budget estimator
//a simple CLI tool to take input for CPU cores, RAM and storage, does calculations and displays formatted receipt

package main 

import "fmt"

func main() {
	//defining the fixed pricing 
	const pricePerCore = 5.00
	const pricePerGBRam = 2.50
	const pricePerGBStorage = 0.10

	//variable declaration for user inputs
	var cores int
	var ram int
	var storage float64

	fmt.Println("===========================")
	fmt.Println("Cloud Budget Estimator")
	fmt.Println("===========================")

	//collecting user input
	fmt.Print("Enter number of CPU cores: ")
	fmt.Scan(&cores)
	fmt.Print("Enter amount of RAM in GB: ")
	fmt.Scan(&ram)
	fmt.Print("Enter amount of storage: ")
	fmt.Scan(&storage)

	//performing calculations
	//convert to float to calculate final prices
	costCPU := float64(cores)*pricePerCore
	costRAM := float64(ram)*pricePerGBRam
	costStorage := storage*pricePerGBStorage

	baseTotal := costCPU + costRAM + costStorage
	finalTotal := baseTotal
	discount := 0.0

	//applying conditinal logic to give discount only is total is aobve 100
	if baseTotal > 100.00 {
		discount = baseTotal * 0.10
		finalTotal = baseTotal - discount
	}

	//output the formatted receipt
	fmt.Println("\n----------------------------------------")
    fmt.Println("             MONTHLY ESTIMATE           ")
    fmt.Println("----------------------------------------")
	fmt.Printf("CPU Cost: $%.2f\n", costCPU)
	fmt.Printf("RAM Cost: $%.2f\n", costRAM)
	fmt.Printf("Storage Cost: $%.2f\n", costStorage)
	fmt.Println("======================================")
	fmt.Printf("Base Total: $%.2f\n", baseTotal)

	
    if discount > 0 {
        fmt.Printf("Discount     :-$%.2f (10%% Volume)\n", discount)
    }
     
	//printing the final amount in green by passing the ASCI escape code
    fmt.Printf("Final Total  : \033[32m$%.2f\033[0m\n", finalTotal)
    fmt.Println("========================================")
}