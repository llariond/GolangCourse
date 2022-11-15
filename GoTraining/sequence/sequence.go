package main

import "fmt"

func generator(firstNum int, calcNext func(int) int) func() int {
	a0 := firstNum
	return func() int {
		a0 = calcNext(a0)
		return a0
	}
}

func main() {

	var calls, startNum int

	fmt.Print("Enter number of calls: ")
	fmt.Scan(&calls) //omitting Scan error handling

	fmt.Print("Enter starting number: ")
	fmt.Scan(&startNum) //omitting Scan error handling

	test0 := generator(startNum, func(a int) int { return 2 * a })
	test1 := generator(startNum, func(a int) int { return a + 2 })
	test2 := generator(startNum, func(a int) int { return -1 * a })
	test3 := generator(startNum, func(a int) int { return a * a }) //big numbers!

	//no display for starting number
	for i := 0; i < calls; i++ {
		fmt.Printf("For a%d: test0: %d; test1: %d; test2: %d; test3: %d\n", i+1, test0(), test1(), test2(), test3())
	}
}
