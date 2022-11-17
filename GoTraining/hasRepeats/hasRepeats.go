package main

import "fmt" //used only for input-output

func hasRepeats(str string) bool {
		
	var res int64
	for _,r :=range(str) {
		switch {
			case r >= 'A' && r <= 'Z' :
				if res & (int64(1) << (r - 'A')) != 0 {
					return true
				} else {
					res = res | int64(1) << (r - 'A')
				}
			case r >= 'a' && r <= 'z' :
				if res & (int64(1) << (r - 'a')) !=0 {
					return true
				} else {
					res = res | int64(1) << (r - 'a')
				}				
			default: //ignore other symbols
				continue
		}
	}
	return false
}

func main() {

    fmt.Print("Enter string: ")
	
	var input string
    fmt.Scan(&input) //omitting Scan error handling
	fmt.Println("Read: ", input)
	fmt.Println("Has repeats: ", hasRepeats(input))
}
