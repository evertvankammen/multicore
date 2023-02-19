package main

import (
	"evertlevert.nl/processes"
	"fmt"
)

func main() {
	// Get a greeting message and print it.
	message := processes.Hello("Gladys")
	fmt.Println(message)

	processes.Values()
	processes.Variables()
	processes.LoopIfElseSWitch()

	nextInt := processes.IntSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}