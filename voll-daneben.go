package main

import (
	"fmt"
	"github.com/benCoder01/bwinf37A3/alcapone"
	"github.com/benCoder01/bwinf37A3/converter"
	"github.com/benCoder01/bwinf37A3/sorter"
	"os"
)

func main() {
	// Kommandozeilenargumente einlesen
	// z.B.: ./_example_data/a3-Voll_daneben_beispieldaten_beispiel2.txt
	commandLineArgs := os.Args

	if len(commandLineArgs) != 2 {
		panic("Falsche Eingabe: ./main <Pfad-zur-Datei>")
	}

	filePath := os.Args[1]

	fmt.Println("Die AlCapone-Zahlen für die Datei: ", filePath)

	participantNumbers := converter.Convert(filePath)
	sorter.Sort(participantNumbers)

	// AlCapone-Zahlen und Gewinn berechnen
	caponeNumbers, win := alcapone.Choose(participantNumbers)

	// Ergebnisse ausgeben
	fmt.Println("Die von AlCapone zu wählenden Zahlen:")
	printArray(caponeNumbers)

	fmt.Println("Gewinn: ", win)
}

func printArray(values []int) {
	for _, element := range values {
		fmt.Println(element)
	}
}
