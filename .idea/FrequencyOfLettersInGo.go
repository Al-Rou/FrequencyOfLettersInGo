package main

import (
	"fmt"
	"os"
)
import "bufio"

func main() {
	fmt.Println("Enter your sentence: ")
	var input string
	input, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	var frequenyCounter = [256] int {}
	for i := 0; i < len(input); i++{
		frequenyCounter[input[i]]++
	}
	for i := 0; i < len(frequenyCounter); i++{
		if frequenyCounter[i] != 0 {
			if ((i != 10) && (i != 32)){
				fmt.Printf("%v appears %v times\n", string(i), frequenyCounter[i])
			}
		}
	}
}
