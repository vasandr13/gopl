
// took it from here:
// https://golang.org/doc/effective_go.html#constants
// took explanation of bitwise operator from here:
// https://stackoverflow.com/questions/5801008/go-and-operators


package main

import "fmt"

const (
	_  = iota             	// 0
	KB = 1 << (10 * iota) 	// 1 << (10 * 1) = (1 times 2) ten times = 1024
	MB                    	// 1 << (10 * 2) = (1 times 2) twenty times = 1048576
	GB			// 1 << (10 * 3) = (1 times 2) thirty times = 1073741824
	TB
	PB
  	EB
)

func main() {
  fmt.Println(KB, MB, GB, TB, PB, EB)
}


// go run ex3_13.go 
// 1024 1048576 1073741824 1099511627776 1125899906842624 1152921504606846976
