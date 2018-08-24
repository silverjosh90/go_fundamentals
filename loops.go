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