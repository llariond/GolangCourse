package main

import "fmt"

func main() {

	var maxSize int

	fmt.Print("Enter max size of slice: ")
	fmt.Scan(&maxSize) //omitting Scan error handling

	myslice := make([]int64, 0, 0)
	curCap := cap(myslice)
	
	for len(myslice) < maxSize {
		myslice = append(myslice, 1)
		if cap(myslice) != curCap {
			fmt.Printf("Changing cap at len: %d from %d to %d\n", 
				len(myslice),
				curCap,
				cap(myslice))
			curCap = cap(myslice)
		}
	}

	fmt.Println("Finished")
}
