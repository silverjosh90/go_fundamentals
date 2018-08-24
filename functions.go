package main

import (
	"fmt"
	"strings"
)

func functionsPractice() (string1, string2 string) {
	name := "upperAndTitle"
	fmt.Println("HELLO THERE")

	return upperAndTitle(name)
}

func upperAndTitle(name string) (upper, title string) {
	return strings.ToUpper(name), strings.Title(name)
}

func averageArbitraryNumberOfArgs(someSlice ...int) int {
	return 4
}
