package alcapone

var values []int
var caponeNumbers []int

type group struct {
	startPos     int // Anfangsposition der Gruppe
	endPos       int // Endpossition (letzte Stelle der Gruppe im  Array)
	caponeNumber int // Mittelwert aus allen Zahlen in der Gruppe
	size         int // Anzahl der Zahlen in einer Gruppe
}

// Choose bekommt ein Array aus Integern, in dem alle gewählten Zahlen der Teilnehmer enthalten sind. Als Rückgabewert gibt die Methode die von AlCapone Jr. zu wählenden Zahlen als Array aus.
func Choose(pValues []int) []int {
	values = pValues

	groups := divideGroups()

	if similarDistance() && len(groups) > 10 {
		caponeNumbers = caponeNumbersFromSections()
	}else {
		for len(groups) > 10 {
			groupsToMerge := findGroupsToMerge(groups)
			groups = merge(groups, groupsToMerge[0], groupsToMerge[1])
		}

		// Alle Zahlen zu dem CaponeNumbers Array hinzufügen
		for _, group := range groups {
			caponeNumbers = append(caponeNumbers, group.caponeNumber)
		}
	}

	// Auffüllen der capone Numbers bis die Länge 10 beträgt
	for len(caponeNumbers) < 10 {
		caponeNumbers = append(caponeNumbers, 1)
	}

	return caponeNumbers
}

// merge fügt zwei Gruppen, deren Position als Parameter angegeben wird, in dem Array groups zusammen.
// Die beiden Gruppen müssen direkt hintereinander liegen.
func merge(groups []group, posGroup1 int, posGroup2 int) []group {
	if len(groups) < posGroup2-1 || posGroup1+1 != posGroup2 {
		return groups
	}

	// neue Gruppe erstellen
	var mergedGroup group

	mergedGroup.startPos = groups[posGroup1].startPos // Anfangsposition der ersten Gruppe
	mergedGroup.endPos = groups[posGroup2].endPos     // Endposition der zweiten Gruppe
	mergedGroup.size = 1 + (mergedGroup.endPos - mergedGroup.startPos)

	mergedGroup.caponeNumber = getAvgFromGroup(mergedGroup)

	// Die beiden Gruppen durch die mergedGroup ersetzen
	groups[posGroup2] = mergedGroup                            // mergedGroup auf die Position der zweiten Gruppen setzen
	groups = append(groups[:posGroup1], groups[posGroup2:]...) // Gruppe an der ersten Position löschen

	return groups

}

// findGroupsToMerge sucht aus allen Gruppen das Gruppenpaar aus, bei dem das Zusammenfügen mit dem geringsten Verlust
// vorgenommen werden kann.
func findGroupsToMerge(groups []group) []int {
	if len(groups) == 0 {
		return []int{}
	}


	bestGroupPos := 0 // Standardwert für die Positino der ersten Gruppe des Gruppenpaares

	averageFirstAndSecondGroup := getAvgFromGroups(groups, bestGroupPos, bestGroupPos+1)
	winBestGroups := calculateWinOfTwoGroups(groups, bestGroupPos, bestGroupPos+1, averageFirstAndSecondGroup)

	// Findet das Gruppenpaar, das nach dem Zusammenfügen den größt möglichen Gewinn hätte.
	for i := 1; i < len(groups)-1; i++ {
		winInGroup := calculateWinOfTwoGroups(groups, i, i+1, getAvgFromGroups(groups, i, i+1))

		if winInGroup > winBestGroups {
			winBestGroups = winInGroup
			bestGroupPos = i
		}
	}

	return []int{bestGroupPos, bestGroupPos + 1}
}

// calculateWinOfTwoGroups berechnet den Gewinn bzw. Verlust, wenn zwei Gruppen zusammengefügt werden würden.
// Die beiden Gruppen müssen dabei direkt hintereinander im Array liegen.
func calculateWinOfTwoGroups(groups []group, pos1 int, pos2 int, number int) int {
	if pos1+1 != pos2 {
		return 0
	}

	sum := 0 // Nettoverlust

	// Nacheinander alle Element aus beiden Gruppen mit der gegebenen Zahl vergleichen
	for i := groups[pos1].startPos; i <= groups[pos2].endPos; i++ {
		sum += difference(number, values[i])
	}

	// Gesamtverlust wird mit dem Einsatz verrechnet zurückgegeben
	return (25 * (groups[pos1].size + groups[pos2].size)) - sum
}

// divideGroups teilt die gewählten Zahlen der Teilnehmer in Gruppen ein, die durch das group-struct dargestellt werden.
func divideGroups() []group {
	if len(values) == 0 {
		return []group{}
	}

	var groups []group

	for i := 0; i < len(values); i++ {
		group := group{}
		group.startPos = i

		// Über alle Werte deren Differenz kleiner als 25 zum ersten Wert der Gruppe ist
		for i+1 < len(values) && values[i]-values[group.startPos] < 25 {
			i++
		}

		medianPos := i // Median als Anhaltspunkt für zweiten Teil der Gruppe

		// Über alle Werte deren Differenz kleiner als 25 zum mittleren Wert der Gruppe ist
		for i+1 < len(values) && values[i]-values[medianPos] < 25 {
			i++
		}

		group.endPos = i // Die Gruppe deckt nun einen Bereich von 50 ab.

		group.size = 1 + (group.endPos - group.startPos)

		group.caponeNumber = getAvgFromGroup(group) // Capone Zahl als Durchschnitt aller Zahlen in einer Gruppe

		groups = append(groups, group)
	}

	return groups
}

// difference berechnet den Betrag der Differenz zweier Zahlen
func difference(number1 int, number2 int) int {
	difference := number1 - number2

	if difference < 0 {
		difference *= -1
	}

	return difference
}

// getAvgFromGroup berechnet den Durchschnitt aller Elemente aus einer Gruppe
func getAvgFromGroup(group group) int {
	sum := 0

	for i := group.startPos; i <= group.endPos; i++ {
		sum += values[i]
	}

	return sum / group.size
}

// getAvgFromGroups berechnet den Durchschnitt aller Zahlen von zwei Gruppen
func getAvgFromGroups(groups []group, pos1 int, pos2 int) int {
	return (getAvgFromGroup(groups[pos1]) + getAvgFromGroup(groups[pos2])) / 2
}

// similarDistance überprüft, ob die gewählten Zahlen jeweils den gleichen Abstand zu ihren Nachbarzahlen haben.
func similarDistance() bool {
	if len(values) <= 1 {
		return true
	}

	distance := difference(values[0], values[1])

	for i := 1; i < len(values) - 1; i++ {
		if difference(values[i], values[i+1]) != distance {
			return false
		}
	}

	return true
}

// caponeNumbersFromSections erstellt die Capone Zahlen durch den Mittelwert aus jeweils 10 Abschnitten.
// Insbesonder bei gleich großen Gruppen wird dies angewendet.
func caponeNumbersFromSections() []int {
	sectionStep := values[len(values)-1] / 10 // Differenz zwischen Anfang und Ende eines Bereiches

	var caponeNumbers []int

	valuesIterator := 0

	for i := 1; i <= 10; i++ {
		sumInSection := 0 // Summe aller Zahlen, die im Bereich enthalten sind
		sectionSize := 0
		// Alle Zahlen, die im i-ten Bereich sind
		for ;values[valuesIterator] <= i*sectionStep; valuesIterator++ {
			sumInSection += values[valuesIterator]
			sectionSize++
		}

		caponeNumbers = append(caponeNumbers, sumInSection/sectionSize)
	}

	return caponeNumbers
}
