package control_structures

import "fmt"

func fizzBuzz(i int32) string {
	var result string = "";
	
	if (i%3 == 0 && i%5 == 0) {
		result = "FizzBuzz"
	}else if i%5 == 0 {
		result = "Buzz" 
	}else if i%3 == 0 {
		result = "Fizz" 
	} else {
		result = fmt.Sprintf("%d", i)
	}

	return result
}