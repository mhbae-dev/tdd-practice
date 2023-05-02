package main

import (
	"fmt"
	"os"
)

// This is where an end to end test would go
func main() {

	ohce := Ohce{
		responseBuilder: &ResponseBuilder{},
		clock:           &Clock{},
	}
	var input string
	fmt.Println("Hi, I'm Ohce! What's your name?")
	fmt.Scan(&input)
	fmt.Print(ohce.Greet(input))
	for {
		fmt.Scan(&input)
		if input == "Stop!" {
			fmt.Println(ohce.Respond(input))
			os.Exit(0)
		}
		fmt.Println(ohce.Respond(input))
	}

}
