package operator


func Div(x,y int) int {
	if y <= 0 {
		panic("divisor cannot be zero")
	}

	return x/y
}
