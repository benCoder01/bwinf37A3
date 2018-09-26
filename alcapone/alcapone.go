package alcapone

var values []int

func Choose(pValues []int) []int {
	// Werte zu globalen Variable hinzufügen
	values = pValues

	// Aufgrund Speichereffizenz Variablen vorher initialisieren

	var lowestDistance int
	// Index, der die erste Position von zwei Zahlen angibt, die in der gesamten Zahlenmenge den geringsten Abstand zueinander haben.

	var average int // Durchschnitt der beiden Zahlen, die durch lowesDistance angegeben werden.

	for len(values) > 10 {
		lowestDistance = findLowestDistance()
		average = findAverage(lowestDistance)
		replaceValues(lowestDistance, average)
	}

	// restlichen Werte in values auffüllen
	for len(values) < 10 {
		values = append(values, 1)
	}

	return values
}

// replaceValues bekommt das Array mit Zahlen und die Position,
// an der eine Zahl ersetzt werden soll. Dabei werden die Positionen pos und pos+1 mit number erstzt.
func replaceValues(pos int, number int) {
	// Die Zahl bei Index pos+1 durch number ersetzten
	values[pos+1] = number
	// Die Zahl bei Index pos Löschen.
	values = append(values[:pos], values[pos+1:]...)

}

func findLowestDistance() int {
	// es wird immer die Distanz von lowest und lowest+1 betrachtet
	lowest := 0

	// Es werden immer zwei Werte überprüft. Daher wird das Array nur bis Länge-1 getestet
	for i := 1; i < len(values)-1; i++ {
		if values[i+1]-values[i] < values[lowest+1]-values[lowest] {
			// ein kleinerer Abstand wurde gefunden
			lowest = i
		}
	}

	// es wird immer die erst-beste Position ausgegeben
	return lowest
}

func findAverage(pos int) int {
	return (values[pos+1] + values[pos]) / 2
}
