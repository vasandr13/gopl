package main

import (
	"fmt"
	"time"
)

func main() {

	current_time := time.Now()

	fmt.Println("\nHello from Rebrain!")

	fmt.Printf("%02d.%02d.%02d %02d:%02d\n",
		current_time.Day(), current_time.Month(), current_time.Year(),
		current_time.Hour(), current_time.Minute())

	fmt.Println()

}
