// Package converter implementiert die Logik, um eine
// Eingabedatei einzulesen.
package converter

import (
	"bufio"
	"os"
	"strconv"
)

// Convert nimmt als Parameter den Dateipfad zu der Datei mit
// den Glückszahlen der Teilnehmer entgegen. Alle Zahlen aus
// dieser path-Datei werden als Array zurückgegeben.
func Convert (path string) []int {
	r, err := os.Open(path) // io.Reader r erstellen

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int

	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		result = append(result, x)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	return result

}

// Quelle: https://stackoverflow.com/questions/9862443/golang-is-there-a-better-way-read-a-file-of-integers-into-an-array

