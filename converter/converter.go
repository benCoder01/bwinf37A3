package converter

import (
	"bufio"
	"os"
	"strconv"
)

// Convert nimmt als Parameter den Dateipfad zu der Datei mit den gewählten Zahlen der Teilnehmer entgegen. Alle Zahlen aus der path Datei
// werden in ein Array geschrieben. Am Schluss wird das fertige Array zurückgegeben.
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

	// https://stackoverflow.com/questions/9862443/golang-is-there-a-better-way-read-a-file-of-integers-into-an-array
}
