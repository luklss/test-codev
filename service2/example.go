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

func pow(base, power int) int {
	result := 1
	for i := 0; i < power; i++ {
		result = result * base
	}

	return result

}
