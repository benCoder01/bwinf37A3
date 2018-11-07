// Package alcapone implementiert den Algorithmus zum Berechnen
// der Zahlen, die AlCapone Jr. wählen muss. Dabei wird
// versucht, dass AlCapone Jr. einen möglichst hohen Gewinn
// erzielt.
package alcapone

import "github.com/benCoder01/bwinf37A3/win"

// values enthält ein Array aus Glückszahlen.
// Dies sind die aufsteigend sortiert.
var values []int

// groups enthält ein Array aus Gruppen, die aus den
// Glückszahlen gebildet wurden.
var groups []Group

// Group repräsentiert eine Gruppe aus Glückszahlen. Anfangs-
// und Endposition sind inklusive.
// Außerdem sind die Methoden valid() und find().
type Group struct {
	startPos     int // Anfangsposition der Gruppe
	endPos       int // Endpossition
	caponeNumber int // Ausgewählte Zahl aus der Gruppe
	size         int // Anzahl der Zahlen in einer Gruppe
}

// GroupPair bezeichnet ein Paar aus Gruppen. Weiterhin sind
// für GroupPair die Methoden merge(), win() und median()
// vorhanden. group1 und group2 sind Zeiger, die jeweils auf
// ein Group Objekt verweisen.
type GroupPair struct {
	group1, group2 *Group
}

// Choose bekommt die sortierten Glückszahlen der Teilnehmer
// als Array übergeben und berechnet daraus die Zahlen, die
// AlCapone Jr. wählen sollte. Dabei ist ein Gewinn nicht
// garantiert, sondern es wird lediglich versucht einen
// möglichst hohen Gewinn zu erzielen. Bei einer Länge von
// maximal 10 werden die Glückszahlen als AlCapone-Zahlen
// zurückgegeben. Vorher wird jedoch die Länge auf 10 erhöht.
// Ansonsten vergleicht die Funktion den Gewinn zweier
// Vorgehensweisen: Zum einen das Bilden von Gruppen, zum
// anderem das Errechnen der Zahlen aus dem Durchschnitt der
// Glückszahlen in gleich großen Abschnitten. Die Werte der
// Vorgehensweise mit dem höheren Gewinn werden dann mit
// dem Gewinn von AlCapone Jr. zurückgegeben.
func Choose(pValues []int) ([]int, int) {
	values = pValues

	for len(values) < 10 {
		values = append(values, 1)
	}

	if len(values) == 10 {
		return values, win.Calculate(pValues, values)
	}

	numbersAverage := numbersFromAverage()
	numbersGroups := numbersFromGroups()

	winAverage := win.Calculate(values, numbersAverage)
	winGroups := win.Calculate(values, numbersGroups)
	if winAverage > winGroups {
		return numbersAverage, winAverage
	} else {
		return numbersGroups, winGroups
	}

}

// numbersFromGroups berechnet mit Bildung von Gruppen die
// AlCapone-Zahlen.
func numbersFromGroups() [] int {
	if len(values) == 0 {
		return []int{}
	}
	groups = divide()

	for len(groups) > 10 {
		gp := createGroupPair()
		gp.merge()
	}

	var numbers []int

	for _, group := range groups {
		numbers = append(numbers, group.caponeNumber)
	}

	return numbers
}

// numbersFromAverage teilt alle Glückszahlen der Teilnehmer
// in Abschnitte ein. Für jeden Abschnitte wird dann aus den
// Glückszahlen in dem Abschnitte ein Durchschnitt bestimmt.
// Sollte keine Glückszahlen in einem Abschnitt enthalten sein,
// so wird der Mittelwert aus oberer und unterer Grenze des
// Abschnittes als Zahl für den Abschnitt gewählt. Die
// Durchschnitte werden dann in einem Array zurückgegeben.
func numbersFromAverage() []int {
	if len(values) == 0 {
		return []int{}
	}

	sectionDistance := (values[len(values)-1] - values[0]) / 10
	var numbers []int
	i := 0 // Zähler über das Array values

	for section := 1; section <= 10; section++ {
		sum := 0
		size := 0

		/*
		section*sectionDistance berechnet die Obergrenze des
		aktuellen Bereiches. Dabei hat man die aktuelle Nummer
		der Sektion und den Abstand zwischen den Abschnitten.
		*/
		for ; values[i] < section*sectionDistance; i++ {
			sum += values[i]
			size++
		}

		if size == 0 {
			average := ((section-1)*sectionDistance + section*sectionDistance) / 2
			numbers = append(numbers, average)
		} else {
			numbers = append(numbers, sum/size)
		}
	}

	return numbers
}

// divide teilt aus allen Glückszahlen Gruppen ein. Dabei
// hat jede Gruppe einen Wert, dessen Differenz jeweils
// zum Anfangs- und Endwert nicht größer als 25 ist. Dieser
// Wert wird in jeder Gruppe der caponeNumber zugeordnet.
func divide() []Group {
	if len(values) == 0 {
		panic("Keine Werte als Teilnehmerzahlen enthalten!")
	}

	var groups []Group

	for i := 0; i < len(values); i++ {
		group := Group{}
		group.startPos = i
		firstValue := values[group.startPos]

		for i+1 < len(values) && values[i+1]-firstValue <= 25 {
			i++
		}

		mid := i
		group.caponeNumber = values[mid]

		for i+1 < len(values) && values[i+1]-values[mid] <= 25 {
			i++
		}

		group.endPos = i

		group.size = 1 + (group.endPos - group.startPos)

		groups = append(groups, group)
	}

	return groups
}

// merge fügt ein Gruppenpaar zusammen. Dabei werden die beiden
// Gruppen direkt in dem Array groups ersetzt.
func (gp *GroupPair) merge() {

	if !gp.valid() {
		panic("Gruppen können nicht zusammengefügt werden!")
	}

	var group Group // neue Gruppe
	group.startPos = gp.group1.startPos
	group.endPos = gp.group2.endPos

	group.size = gp.group1.size + gp.group2.size
	group.caponeNumber = gp.median()

	if !group.valid() {
		panic("Fehler beim mergen von Gruppen!")
	}

	// zweite Gruppe durch neue Gruppe ersetzen:
	groups[gp.group2.find()] = group
	// erste Gruppe löschen:
	groups = append(groups[:gp.group1.find()], groups[gp.group2.find():]...)

}

// win berechnet den Gewinn/Verlust, der beim Zusammenfügen
// beider Gruppen aus GroupPair entstehen würde. Dafür wird
// ein Median aus beiden Gruppen berechnet. Die Differenz
// zwischen jeder Zahl in der neuen Gruppe zu deren Median
// wird dann auf den Nettoverlust der Gruppe addiert. Der
// zurückgegebene Gesamtgewinn berechnet sich dann aus den
// Einsätzen der Teilnehmer, deren Zahl jeweils in einer der
// beiden Gruppen ist, und dem berechneten Nettoverlust, der
// von den Einsätzen abgezogen wird.
func (gp *GroupPair) win() int {
	if !gp.valid() {
		panic("Kann den Gewinn einer Gruppe nicht berechnen!")
	}

	sum := 0
	median := gp.median()

	for i := gp.group1.startPos; i <= gp.group2.endPos; i++ {
		sum += difference(median, values[i])
	}

	return (25 * (gp.group1.size + gp.group2.size)) - sum
}

// median berechnet den neuen Median der beiden Gruppen aus dem
// Gruppenpaar. Als Median wird im diesen Fall die Zahl
// bezeichnet, die zwischen beiden Gruppen steht.
func (gp *GroupPair) median() int {
	if !gp.valid() {
		panic("Kann den Median einer Gruppe nicht berechnen!")
	}

	return values[(gp.group1.startPos+gp.group2.endPos)/2]

	sum := 0

	for i := gp.group1.startPos; i <= gp.group2.endPos; i++ {
		sum += values[i]
	}

	return sum / (gp.group1.size + gp.group2.size)

}

// valid überprüft ein Gruppenpaar auf Richtigkeit.
func (gp *GroupPair) valid() bool {
	posGroup1 := gp.group1.find()
	posGroup2 := gp.group2.find()

	return !(
		posGroup1 >= posGroup2 ||
			!gp.group1.valid() ||
			!gp.group2.valid())
}

// createGroupPair such in allen Gruppen nach dem besten
// Gruppenpaar, welches zusammengefügt werden sollte.
// Dabei wird über das gesamte Array groups iteriert, wobei
// immer zwei Gruppen mit dem bisherigen besten Gruppenpaar
// verglichen werden. Dabei wird das Gruppenpaar gewählt, das
// beim Zusammenfügen den niedrigsten Verlust/höchsten Gewinn
// erzielt.
func createGroupPair() GroupPair {
	if len(values) < 1 || len(groups) < 2 {
		panic("Es kann nicht nach Gruppen zum Zusammenfügen gesucht werden")
	}

	gp := GroupPair{&groups[0], &groups[1]}
	for i := 1; i < len(groups)-1; i++ {
		currrentGP := GroupPair{&groups[i], &groups[i+1]}

		if gp.win() < currrentGP.win() {
			gp = currrentGP
		}
	}

	return gp
}

// valid überprüft ob der Bereich, den eine Gruppe abdeckt
// innerhalb dem Bereich des Arrays values ist, auf das sich
// jede Gruppe bezieht.
func (g *Group) valid() bool {

	return !(
		g.endPos < g.startPos ||
			g.startPos < 0 ||
			g.endPos > len(values)-1 ||
			g.size != 1+(g.endPos-g.startPos))
}

// find sucht die Positon der Grupppe g im Array groups.
func (g *Group) find() int {
	for i := 0; i < len(groups); i++ {
		if g == &groups[i] {
			return i
		}
	}

	panic("Eine Gruppe wurde nicht im Array groups gefunden")
}

// difference berechnet den Betrag der Differenz zweier Zahlen.
func difference(number1 int, number2 int) int {
	difference := number1 - number2

	if difference < 0 {
		difference *= -1
	}

	return difference
}

