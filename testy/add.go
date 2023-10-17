package testy

func add(a, b int) int {
	return a + b
}

// Metoda sloužící k odečtení dvou celých čísel
func substract(a, b int) int {
	return a - b
}

func doMath(a, b int, f func(int, int) int) int {
	return f(a, b)
}
