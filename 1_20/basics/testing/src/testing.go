package src

func Even(i int) bool { // Exported function
	return i%2 == 0
}

func Odd(i int) bool { // Exported function
	return i%2 != 0
}

func Mul(a int, b int) int {
	return a * b
}
