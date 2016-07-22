package control_structures

const add string = "+"
const subtract string = "-"
const multiply string = "*"
const divide = "/"

func calculate(op string, arg1, arg2 float32) (a float32) {
	switch op {
	case add:
		a = arg1+arg2
	case subtract:
		a = arg1-arg2
	case multiply:
		a = arg1*arg2
	case divide:
		a = arg1/arg2

	}
	return
}
