package win


// CalculateCaponeWin berechnet den Gesamtgewinn von AlCapone Jr.
// Dafür werden zum einen die Zahl der Teilnehmer und die Zahlen von AlCapone eingelesen.
func CalculateCaponeWin(participantNumbers []int, alCaponeNumbers []int) int {
	if len(alCaponeNumbers) != 10 {
		return -1
	}

	payout := 0

	// aus Speichergründen wir nearestNumber vor der Schleife initialisiert
	var nearestNumber int

	for _, number := range participantNumbers {
		// finde in den Zahlen von AlCapone die am besten passende Zahl
		nearestNumber = findNearestNumber(number, alCaponeNumbers)
		payout += calculateDifference(number, nearestNumber)
	}

	// Gib die Einsätze der Teilnehmer - dem was ausgezahlt werden muss.
	return (len(participantNumbers) * 25) - payout

}

// findNearestNumber findet in den von AlCapone Jr. gewählten Zahlen die Zahl, die am nächsten zu number dran ist.
func findNearestNumber(number int, alCaponeNumbers []int) int {
	nearestNumber := alCaponeNumbers[0] //alCaponeNumbers haben immer eine Länge von 10

	for _, element := range alCaponeNumbers {
		if calculateDifference(number, element) < calculateDifference(number, nearestNumber) {
			// die Zahl in Element ist näher an number dran
			nearestNumber = element
		}
	}

	return nearestNumber
}

// calculateDifference berechnet die Differnz zweier Zahlen, ohne dabei auf weiter Packte zuzugreifen.
func calculateDifference(number1 int, number2 int) int {
	difference := number1 - number2

	if difference < 0 {
		difference *= -1
	}

	return difference
}

