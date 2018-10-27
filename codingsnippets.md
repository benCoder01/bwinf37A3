```go
// Choose bekommt ein Array aus Integern, in dem alle gewählten Zahlen der Teilnehmer enthalten sind. Als Rückgabewert gibt die Methode die von AlCapone Jr. zu wählenden Zahlen als Array aus.
func Choose(pValues []int) []int {
	// Werte zu globalen Variable hinzufügen
	values = pValues

	// Gruppen finde
	groups := divideGroups()

	var biggestGroup int

	for i := 0; i < 10 && len(values) > 0; i++ {
		biggestGroup = findBiggestGroup(groups)
		caponeNumbers = append(caponeNumbers, values[groups[biggestGroup].caponeNumberPos]) // caponeNumberPos speichert nur die Position der am besten passende Zahl

		// Werte aus values löschen
		values = append(values[:groups[biggestGroup].startPos], values[groups[biggestGroup].endPos + 1:]...)

		// neue Gruppen berechnen
		groups = divideGroups()
	}

	// falls weniger als 10 Zahlen gewählt wurden

	for len(caponeNumbers) < 10 {
		caponeNumbers = append(caponeNumbers, 1)
	}


	return caponeNumbers
}

// findBiggestGroup sucht in einem Array aus Gruppen die größte Gruppe, und gibt die Position im Array als Integer zurück.
// Sollten zwei Gruppen gleich groß sein, wird die Gruppe gewählt, die den größeren Abstand zur nächsten passenden Zahl hat, gewählt.
func findBiggestGroup(groups []group) int{

	// Da es gleich große Gruppen geben kann, wird die Position im Array groups aller größten Gruppen in dem Array biggestGroups gespeichert
	var biggestGroups []int
	highestLength := 0

	// gleich große Gruppen in Array
	for key, group := range groups {
		if group.size > highestLength {
			// eine größere Gruppe wurde gefunden
			biggestGroups = nil
			biggestGroups = append(biggestGroups, key)
			highestLength = group.size
		}else if group.size == highestLength {
			// eine weiter Gruppe der selben Größe wurde gefunden
			biggestGroups = append(biggestGroups, key)
		}
	}

	if len(biggestGroups) < 1 {
		// keine Gruppe wurde gefunden
		return 0
	}

	// beste Gruppe aus dem Array biggestGroups finden. Es muss die Gruppe eingefügt werden, die die weiteteste Entfernung hat. (Die Gruppe, die am einsamsten ist)
	bestGroupPos := biggestGroups[0]

	for _, groupNr := range biggestGroups {

		// die Differenz bei der durch bestGroupPos gewählten Gruppe
		differenceCurrentGroup := getLowestDifference(values[groups[bestGroupPos].caponeNumberPos])

		// die Differenz bei der durch bestGroupPos gewählten Gruppe
		differenceKeyGroup := getLowestDifference(values[groups[groupNr].caponeNumberPos])

		if differenceCurrentGroup < differenceKeyGroup {
			bestGroupPos = groupNr
		}
	}

	return bestGroupPos


}

// getLowestDifference gibt den Abstand zwischen der als number gegebenen Zahl und der am nächsten liegenden Zahl in den caponeNumbers als Integer zurück.
func getLowestDifference(number int) int{

	if len(caponeNumbers) == 0 {
		return 0
	}

	closestNumberPos := 0  // speichert den Index einer Zahl aus caponeNumbers, deren Abstand zu number am geringsten zu allen anderen Zahlen aus caponeNumbers ist.

	for key, caponeNumber := range caponeNumbers {
		if 	difference(number, caponeNumber) > difference(number, caponeNumbers[closestNumberPos]) {
			// es wurde ein kleinerer Abstand gefunden
			closestNumberPos = key // key beschreibt beim Iterieren den Index
		}
	}

	return difference(number, caponeNumbers[closestNumberPos])
}


// findPositionOfClosestGroups findet die zwei Gruppen, deren Capone Zahl den geringsten Abstand zueinander haben. Die Position dieser beiden Gruppen wird dann als Array ausgegeben.
func findPositionOfClosestGroups(groups []group) []int {
	if len(groups) == 2 {
		return []int{}
	}

	var posClosestGroups []int
	// die ersten beiden Gruppen werden als Standardwerte betrachtet. Es immer mindestens zwei Gruppen im Array enthalten
	posClosestGroups = append(posClosestGroups, 0, 1)

	for i := 0; i < len(groups)-1; i++ {
		// die beiden Capone Numbers die im Moment als am nächsten zueinander gesehen werden
		currentCaponeNumber1 := groups[posClosestGroups[0]].caponeNumber
		currentCaponeNumber2 := groups[posClosestGroups[1]].caponeNumber

		// die beiden Capone Numbers deren Differenz zu den aktuelle Nummern verglichen werden sollen.
		possibleCaponeNumber1 := groups[i].caponeNumber
		possibleCaponeNumber2 := groups[i+1].caponeNumber

		differenceCurrent := difference(currentCaponeNumber1, currentCaponeNumber2)
		differencePossible := difference(possibleCaponeNumber1, possibleCaponeNumber2)

		if differenceCurrent > differencePossible {
			posClosestGroups[0] = i
			posClosestGroups[1] = i + 1
		} else if differenceCurrent == differencePossible {
			// Kleinere Gruppen als neue am naheliegensten Gruppen
			currentGroupsSizeSum := groups[posClosestGroups[0]].size + groups[posClosestGroups[1]].size // Gesamtgröße der aktuelle am nähesten zusammen liegenden Gruppen
			possibleGroupsSizeSum := groups[i].size + groups[i+1].size                                  // Gesamtgröße der aktuelle am nähesten zusammen liegenden Gruppen

			if currentGroupsSizeSum > possibleGroupsSizeSum {
				posClosestGroups[0] = i
				posClosestGroups[1] = i + 1
			}
		}
	}

	return posClosestGroups
}


// sortGroups sortiert alle Gruppen aufsteigend der Größe nach
func sortGroups(groups []group) []group {
	for i := 1; i < len(groups); i++ {
		for j := i; j > 0; j-- {
			if groups[j].size < groups[j-1].size {
				swap(groups, j, j-1)
			}
		}
	}

	return groups

}

// swap führt auf dem gegebenen groups Array einen Dreieckstausch mit den Positionen pos1 und pos2 durch.
func swap(groups []group, pos1 int, pos2 int) {
	temp := groups[pos1]
	groups[pos1] = groups[pos2]
	groups[pos2] = temp
}
```