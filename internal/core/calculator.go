package core


func Add(a, b float64) (float64) {
	return a + b
}

func Sub(a, b float64) (float64) {
	return a - b
}

func Mul(a, b float64) (float64) {
	return a * b
}

func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	} else {
		return a / b, nil
	}
}
