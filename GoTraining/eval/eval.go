package main

import (
	"fmt"
	"strconv"
	"strings"
	"errors"
	)

func plus(a any, b any) (any, error) {
	a1, okA := a.(int)
	b1, okB := b.(int)
	if okA == true && okB == true {
		return a1 + b1, nil
	}

	var a2, b2 float64
	if okA == true {
		a2 = float64(a1)
	} else {
		if a2,okA = a.(float64); okA == false {
			return nil, errors.New("Can't convert to number A")
		}	
	}

	if okB == true {
		b2 = float64(b1)
	} else {
		if b2, okB = b.(float64); okB == false {
			return nil, errors.New("Can't convert to number B")
		}	
	}

	return a2 + b2, nil
}

func minus(a any, b any) (any, error) {
	a1, okA := a.(int)
	b1, okB := b.(int)
	if okA == true && okB == true {
		return a1 + b1, nil
	}

	var a2, b2 float64
	if okA == true {
		a2 = float64(a1)
	} else {
		if a2,okA = a.(float64); okA == false {
			return nil, errors.New("Can't convert to number A")
		}	
	}

	if okB == true {
		b2 = float64(b1)
	} else {
		if b2, okB = b.(float64); okB == false {
			return nil, errors.New("Can't convert to number B")
		}	
	}

	return a2 - b2, nil
}

var operations = map[string]func(any, any) (any, error) {
	"plus": plus,
	"minus": minus,
}

func eval(parsed []string) (any, error) {
	
	if len(parsed) == 0 {
		return 0, errors.New("string size == 0")
	}
	if len(parsed)%2 == 0 {
		return 0, errors.New("wrong number of arguments( even, must be odd)")
	}

	var firstArg any
	var secondArg any
	var err error
	
	if firstArg, err = strconv.Atoi(parsed[0]); err!=nil {
		firstArg, err  = strconv.ParseFloat(parsed[0], 64)
		if firstArg, err  = strconv.ParseFloat(parsed[0], 64); err!=nil {
			return nil, errors.New("wrong type of element: " + parsed[0])
		}
	}		
	
	if len(parsed) == 1 {
		return firstArg, nil
	}

	
	for i := 1; i < len(parsed); {
			
		if secondArg, err = strconv.Atoi(parsed[i]); err!=nil {
			if secondArg, err  = strconv.ParseFloat(parsed[i], 64); err!=nil {
				return nil, errors.New("wrong type of element: " + parsed[i])
			}
		}		
		i++
	
		if operation, ok := operations[parsed[i]]; !ok {
			return nil, errors.New("Wrong operation: " + parsed[i])
		} else {
			if firstArg, err = operation(firstArg, secondArg); err!=nil {
				return nil, err
			}
		}
		i++
	}
	
	return firstArg, nil
}

func main() {

    fmt.Print("Enter string: ")
	var input string
	
    _, err := fmt.Scan(&input)
		if err != nil {
		fmt.Println("Scan error: ", err)
	}
	
	parsed := strings.Split(input, ",")
	
	if res, err := eval(parsed); err != nil {
		fmt.Println("Eval error: ",err)
	} else {
		fmt.Printf("Result: %v\n", res)
	}
}
