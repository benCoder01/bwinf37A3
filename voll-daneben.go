package main

import (
	"fmt"
	"github.com/benCoder01/bwinf37A3/alcapone"
	"github.com/benCoder01/bwinf37A3/converter"
	"github.com/benCoder01/bwinf37A3/sorter"
	"github.com/benCoder01/bwinf37A3/win"
	"os"

	//"github.com/benCoder01/bwinf37A3/win"
)

func main() {
	// Kommandozeilenargumente auslesen
	commandLineArgs := os.Args

	if len(commandLineArgs) != 2 {
		panic("Falsche Eingabe: ./main <Pfad-zur-Datei>") // z.B.: ./_example_data/a3-Voll_daneben_beispieldaten_beispiel2.txt
	}

	filePath := os.Args[1]

	participantNumbers := converter.Convert(filePath)

	sorter.Sort(participantNumbers)

	caponeNumbers := alcapone.Choose(participantNumbers)
	// caponeNumbers := newalcapone.Choose(participantNumbers)


	// Glückszahlen bestimmen
	fmt.Println("Die von AlCapone zu wählenden Glückszahlen:")
	printArray(caponeNumbers)

	// Gewinn ausgeben
	caponeWin := win.CalculateCaponeWin(participantNumbers, caponeNumbers)
	fmt.Println("Gewinn: ", caponeWin)
}

func printArray(values []int) {
	for _, element := range values {
		fmt.Println(element)
	}
}

