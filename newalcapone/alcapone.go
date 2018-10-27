package newalcapone

import "fmt"

var values []int
var caponeNumbers []int

// Choose bekommt ein Array aus Integern, in dem alle gewählten Zahlen der Teilnehmer enthalten sind. Als Rückgabewert gibt die Methode die von AlCapone Jr. zu wählenden Zahlen als Array aus.
func Choose(pValues []int) []int {
	// Werte der Teilnehmer zu globalen Variable hinzufügen
	values = pValues

	// Auffüllen der capone Numbers bis die Länge 10 beträgt
	for len(caponeNumbers) < 10 {
		caponeNumbers = append(caponeNumbers, 1)
	}

	generateArrays()

	//caponeNumbers := []int{5,10,15,20,25,30,35,40,45,50}
	return caponeNumbers

}

func compareToCurrentLuckyNumbersArray(luckynumbers []int) int {
	differenceSum := 0

	for i := 0; i < len(values); i++ {
		smallestDifference := difference(values[i], luckynumbers[0])
		for j := 1; j < len(luckynumbers); j++ {
			if difference(values[i], luckynumbers[j]) < smallestDifference {
				smallestDifference = difference(values[i], luckynumbers[j])
			}
		}
		differenceSum += smallestDifference
	}

	return differenceSum
}

func generateArrays() {
	array := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	bestArray := append(array[:0:0], array...)
	smallestSum := compareToCurrentLuckyNumbersArray(array)

	for i:= 1; i <= 1000; i++ {
		for j:=0; j < len(array); j++ {
			array[j] = i
			if compareToCurrentLuckyNumbersArray(array) > smallestSum {
				smallestSum = compareToCurrentLuckyNumbersArray(array)
				bestArray = append(array[:0:0], array...)
			}
		}
	}

	fmt.Println("Gewinn:", (25*len(values)) - smallestSum)
	printArray(bestArray)

}



// difference berechnet die Differnz zweier Zahlen, ohne dabei auf weiter Packte zuzugreifen.
func difference(number1 int, number2 int) int {
	difference := number1 - number2

	if difference < 0 {
		difference *= -1
	}

	return difference
}

func printArray(values []int) {
	for _, element := range values {
		fmt.Println(element)
	}
}