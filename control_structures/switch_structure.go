package control_structures

const add string = "+"
const subtract string = "-"
const multiply string = "*"
const divide = "/"

func calculate(op string, arg1, arg2 float32) float32 {
	var result float32 = 0.0
	switch op {
		case add:
			result = arg1+arg2
		case subtract:
			result = arg1-arg2
		case multiply:
			result = arg1*arg2
		case divide:
			result = arg1/arg2
	}
	return result
}
