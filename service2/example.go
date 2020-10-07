package service2

func sum(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}

func pow(a, b int) int {
	result := 1
	for i := 0; i < b; i++ {
		result += result * a
	}

}
