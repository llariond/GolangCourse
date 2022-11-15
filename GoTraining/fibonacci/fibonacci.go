package main

import "fmt"

func getNextFibNum() func() (nextNum int, index int) {
	counter := -1
	curNum := 1
	prevNum := 0

	return func() (int, int) {
		counter++
		if counter == 0 {
			return prevNum, counter
		} else if counter == 1 {
			return curNum, counter
		} else {
			curNum = curNum + prevNum
			prevNum = curNum - prevNum
			return curNum, counter
		}
	}
}

func main() {

    fmt.Print("Enter max index: ")
	
	var q int
    fmt.Scan(&q) //omitting Scan error handling

	nextNum := getNextFibNum()
	for i := 0; i < q; i++ {
		num, ind := nextNum()
		fmt.Printf("Fibonacci number[%d] = %d\n", ind, num)
	}
}
