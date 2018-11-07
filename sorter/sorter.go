// Package sorter implementiert die Logik, um ein Array aus
// Werten zu sortieren.
package sorter

// Sort sortiert aufsteigend alle Werte aus dem Array values.
// Da nur Werte in values verändert werden, muss das Array
// nicht zurückgegeben werden. Stattdesen werden die Werte im
// Speicher verändert.
func Sort(values []int) {
	for i := 1; i < len(values); i++ {
		for j := i; j > 0; j-- {
			if values[j] < values[j-1] {
				swap(values, j, j-1)
			}
		}
	}
}

func swap(numbers []int, pos1 int, pos2 int) {
	temp := numbers[pos1]
	numbers[pos1] = numbers[pos2]
	numbers[pos2] = temp
}
