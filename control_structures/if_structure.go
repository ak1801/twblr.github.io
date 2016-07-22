package control_structures
import "fmt"
func fizzBuzz(i int32) string {
	var result = i;
	
	if i%3 == 0 {
		result = Fizz
	}else if i%5 == 0 {
		result = Buzz 
	}else if (i%3 == 0 && i%5 == 0) {
		result = fizzBuzz 
	} else {
		result = fmt.Sprintf("%d", i)
	}

	return result
}
