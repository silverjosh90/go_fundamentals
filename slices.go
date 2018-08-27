package main

import (
	"fmt"
)


// Slices are passed as reference by default

func sliceBasics() {
	// Initializing with make. Type, len, capacity
	myCourse := make([]string, 5, 10)
	fmt.Println("This is my len", len(myCourse), cap(myCourse))

	// Initializing with values
	course := []string{"Course1", "Course2", "Course3"}
	fmt.Println("These are my courses", course) 

	// Getting a value
	fmt.Println(course[2]) // Equals Course3
	
	// Change Value
	course[2] = "someOtherCourse"

	// Slice of a slice
	sliceOfCourses := course[0:2]
	fmt.Println(sliceOfCourses)

	// Slices are inherently passed as references
	newArray := append(sliceOfCourses, "additionalCourseAtEnd") // When capacity is reached it is automatically doubled
	fmt.Println("Using append method", newArray)

	// Slices can be appeneded to eachother with an elipses
	newArray = append(newArray, sliceOfCourses...)
	fmt.Println(newArray)
}