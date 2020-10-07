package service1

func sum(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func pow(base, power int) int {
	result := 1
	for i := 0; i < power; i++ {
		result = result * base
	}

	return result

}

func doNothing() {
	a := 2

	_ = a

}
