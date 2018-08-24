package main

import (
	"os"
	"fmt"
)

func testConditionals() {
	if true == true {
		fmt.Println("Conditionals are working")
	}

	// With initialization statement
	if variableOne, variableTwo := 1, "turkey"; true == true {
		fmt.Println("These are the initializations", variableOne, variableTwo)
	}
}

func switchStatement() {
	switch /* Initialization can go here */ "SomeStatenent" {
	case "SomeStatenent":
			fmt.Println("Got into the  switch")
			// fallthrough keyword runs the case below it.
			fallthrough
	case "taco":
			fmt.Println("Fell through")
	default:
			fmt.Println("Not Working Bruh")
	}
}

func errorHandling() (error) {
	_, err := os.Open("~/fakefile.txt")
	if(err != nil) {
		return err
	} else {
		return nil
	}
}
