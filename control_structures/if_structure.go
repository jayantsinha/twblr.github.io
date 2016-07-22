package control_structures

func fizzBuzz(i int32) (r string) {
	//r string
	if i%2 == 0 {
		r = "2"
	}
	if i%3 == 0 {
		if i%5 == 0 {
			r = "FizzBuzz"
		} else {
			r = "Fizz"
		}
	} 
	if i%5 == 0 && i%3 != 0 {
		r = "Buzz"
	}
	return
}
