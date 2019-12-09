//Les fonctions
package main

import (
	"errors"
	"fmt"
)

//printNumbers simply prints two numbers
func printNumbers(a int, b int, prefix string) {
	fmt.Printf("%s a = %d and b = %d\n", prefix, a, b)
}

//printNumbers2 does exactly the same thing than printNumber but type of a and b has been factorised
func printNumbers2(a, b int, prefix string) {
	fmt.Printf("%s a = %d and b = %d\n", prefix, a, b)
}

//diff is just a dummy function that computes the difference between a and b
func diff(a, b int) int {
	return a - b
}

//divide is just a dummy function that divides a by b. It returns an error if b is zero.
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

//divide2 does exactly the same thing than divide but returned values are named here.
func divide2(a, b int) (result int, err error) {
	if b == 0 {
		err = errors.New("cannot divide by zero")
		return
	}
	result = a / b
	return
}

//main function
func main() {
	printNumbers(1, 2, "Did you know?")

	fmt.Printf("diff(5, 3) = %d\n", diff(5, 3))

	var result int
	var err error

	result, err = divide(5, 2)
	fmt.Printf("divide(5, 2) = %d, Error = %v\n", result, err)

	result, err = divide(5, 0)
	fmt.Printf("divide(5, 0) = %d, Error = %v\n", result, err)
}
