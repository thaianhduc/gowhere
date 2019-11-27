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

	/*
		This is an integer
	*/
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

	/*
		Runtime array size
	*/
	tasks := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(tasks))
}
