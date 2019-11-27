package main

import (
	"fmt"
)

func main() {
	/*
		There is a strict rule of the main entry. The package name must be "main" so the main function.
		Compilation error if violated

		The bracket "{" must be in the same line with the function name
	*/
	fmt.Println("My very first Go application. Just follow whatever written on the internet")

	// This is an integer
	a := 10
	fmt.Println(a)

	persons := [4]string{
		"Thai Anh Duc",
		"Dao Dang Khanh Ngoc",
		"Thai Ngoc Diep",
		"Thai Diep Chi"}

	var i int
	for i = 0; i < len(persons); i++ {
		fmt.Println(persons[i])
	}

	// Runtime array size
	tasks := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(tasks))
	fmt.Println("Print an array: ", tasks)

	// Prepare defer functions
	defer deferCallWorkLikeCallBack()

	// Let's call some user defined functions
	sum, div := calculate(2, 4)
	fmt.Println("Call func result: ", sum, div)
}

func calculate(a int, b int) (int, int) {
	// Can return multiple values
	// if the function name starts with Capital letter -> public method
	// otherwise, it is private
	// So Calculate will allow other packages to access it
	return a + b, a - b
}

func deferCallWorkLikeCallBack() {
	// Is it possible to pass the returned value from calling method to the deferred function?
	var number int = 1
	// This will pass 1 to the deferred method. So not a reference
	defer deferInDefer(number)
	// This will pass the latest number, which is 13, to the deferred function
	// Pass the address (pointer) of the number variable
	defer deferWithPointer(&number)
	fmt.Println("Hi I am a deferred call. You can think of me as a callback in your word")
	number = 13
}

func deferInDefer(number int) {
	fmt.Println("Hi you passed the number:", number)
}

func deferWithPointer(number *int) {
	// This allows the function to access the latest value passed from the calling function
	fmt.Println("Hi you passed the pointer to number:", *number)
}
