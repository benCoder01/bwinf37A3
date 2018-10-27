
package win

import "fmt"

// CalculateCaponeWin berechnet den Gesamtgewinn von AlCapone Jr.
// Dafür werden zum einen die Zahl der Teilnehmer und die Zahlen von AlCapone eingelesen.
func CalculateCaponeWin(participantNumbers []int, alCaponeNumbers []int) int {
	if len(alCaponeNumbers) != 10 {
		fmt.Println("Error: Die Anzahl der Zahlen ist nicht 10")
		return 0
	}

	payout := 0 // Nettoverlust

	for _, number := range participantNumbers {
		nearestNumber := findNearestNumber(number, alCaponeNumbers)
		payout += difference(number, nearestNumber)
	}

	return (len(participantNumbers) * 25) - payout // Nettoverlust mit den Einstätzen verrechnet
}

// findNearestNumber findet in den von AlCapone Jr. gewählten Zahlen die Zahl, die am nächsten zu number dran ist.
func findNearestNumber(number int, alCaponeNumbers []int) int {
	nearestNumber := alCaponeNumbers[0] //alCaponeNumbers haben immer eine Länge von 10

	for _, element := range alCaponeNumbers {
		if difference(number, element) < difference(number, nearestNumber) {
			// die Zahl in Element ist näher an number dran
			nearestNumber = element
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
