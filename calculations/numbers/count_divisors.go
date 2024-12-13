package numbers

// Erwartet eine Zahl n >= 1 und liefert die Anzahl der Teiler dieser Zahl zurück.
func CountDivisors(n int) int {
	result := 0

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			result++
		}
	}

	return result
}
