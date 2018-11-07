// Package win implementiert die Berechnung des Gewinnes für
// AlCapone Jr. anhand der Zahlen der Teilnehmer und der Zahlen,
// die AlCapoen Jr. gewählt hat.
package win

// Calculate berechnet den Gewinn für AlCapone Jr. Dafür werden
// sowohl die Teilnehmerzahlen, als auch die Zahlen von
// AlCapone Jr. der Funktion übergeben. Für jede Zahl der
// Teilnehmer wird dann die am nähesten liegende Zahl in den
// AlCapone-Zahlen gesucht. Die Differenz zwischen den beiden
// Zahlen wird dann auf die Gesamtauszahlung addiert. Bei der
// Rückgabe des Gewinns wird dann die Gesamtauszahlung mit dem
// Einsatz verrechnet, und man erählt den tatsächlichen Gewinn
// AlCapones.
func Calculate(participant []int, alcapone []int) int {
	if len(alcapone) != 10 {
		panic("Error: Die Anzahl der Zahlen ist nicht 10")
		return 0
	}

	payout := 0

	for _, number := range participant {
		num := closestNumber(number, alcapone)
		payout += difference(number, num)
	}

	return (len(participant) * 25) - payout
}

// closestNumber sucht in den Zahlen von AlCapone die Zahl,
// die am nächsten zu der Zahl number liegt.
func closestNumber(number int, alcapone []int) int {
	nearestNumber := alcapone[0]

	for _, value := range alcapone {
		if difference(value, number) < difference(number, nearestNumber) {
			nearestNumber = value
		}
	}

	return nearestNumber
}

// difference berechnet den Betrag der Differenz zweier Zahlen.
func difference(number1 int, number2 int) int {
	difference := number1 - number2

	if difference < 0 {
		difference *= -1
	}

	return difference
}
