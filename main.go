package main

import (
	"fmt"
	"github.com/benCoder01/bwinf37A3/alcapone"
	"github.com/benCoder01/bwinf37A3/converter"
	"github.com/benCoder01/bwinf37A3/sorter"
	"github.com/benCoder01/bwinf37A3/win"
)

func main() {
	makeNumbers()

}

func makeNumbers() {
	participantNumbers := converter.Convert("./example_data/a3-Voll_daneben_beispieldaten_beispiel1.txt")
	//participantNumbers := converter.Convert("./example_data/custom.txt")

	sorter.Sort(participantNumbers)


	alCaponeNumbers := alcapone.Choose(participantNumbers)
	printArray(alCaponeNumbers)

	win := win.CalculateCaponeWin(participantNumbers, alCaponeNumbers)
	println("Gewinn: ", win)
}

func printArray(values []int) {
	for _, element := range values {
		fmt.Println(element)
	}
}
