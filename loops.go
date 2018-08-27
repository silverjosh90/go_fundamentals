package main

import (
	"fmt"
	"time"
)

func loopBasics() {
	for i:=0 ; i < 10; i++ {
		fmt.Println("Looping", i)
		time.Sleep(1 * time.Second)
	}
}

func loopOverRange() {
	sampleSlice := []int {1,3,4,5,6,7}

	BreakPoint: // label for break
	for index, value := range sampleSlice {
		fmt.Println("My Index: ", index, "My Value: ", value)

		for i:=0 ; i < 10; i++ {
			if (i == 3) {
				break BreakPoint
			} else {
				continue // Will run the next iteration of the loop.
			}
		}
	}
}