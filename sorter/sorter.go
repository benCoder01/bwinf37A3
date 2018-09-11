package sorter

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
