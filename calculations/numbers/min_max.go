package numbers

// Berechnet das Minimum von zwei Zahlen.
func Min(a, b int) int {

	if a < b {
		return a
	}
	return b

}

// Berechnet das Maximum von zwei Zahlen.
func Max(a, b int) int {

	if a > b {
		return a
	}
	return b

}

// HINWEISS
// - Nutzen Sie jeweils eine `if`-Anweisung und den `<`-Operator.