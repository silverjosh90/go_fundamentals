package main

import (
	"os"
	"runtime"
	"reflect"
	"fmt"
)

var (
	name   = os.Getenv("LOGNAME")  // Name of user
	course = "Course"  // Name of course
	module = 3.2 // Module number
)

func variables() {
	ptr := &module
	fmt.Println("Hello World", runtime.GOOS)
	fmt.Println("Hello World name", name, reflect.TypeOf(name))
	fmt.Println("Hello World course", course, reflect.TypeOf(course))

	fmt.Println("An & can be used to print the hex address of a variable eg:", &module)
	fmt.Println("An * can be used to create a pointer to a variable eg:", *ptr)
	
	a := 10.000 // Type inference float64
	b := 3      // Type inference int

	c := int(a) + b // int() method converts type to int

	fmt.Println("This is the addition", c)

	changeCourse(&course)
	fmt.Println("Course value", course)
}

func changeCourse (course *string) string {
	*course = "Some other value"

	return *course
}