package win

func ComputeWin(participantNumbers []int, alCaponeNumbers []int) int {
	win := 0

	for _, number := range participantNumbers {
		difference := getDifferenceToNearestNumber(number,alCaponeNumbers)
		win += 25 - difference // Differenz zwischen Einsatz und Auszahlung
	}

	return win
}

func getDifferenceToNearestNumber(number int, alCaponeNumbers []int) int {
	nearestNumber := alCaponeNumbers[0]

	for _, caponeNumber := range alCaponeNumbers {
		if difference(caponeNumber, number) < difference(nearestNumber, number) {
			nearestNumber = caponeNumber
		}
	}

	return difference(nearestNumber, number)
}

func difference(number1 int, number2 int) int {
	difference := number1 - number2

	if difference < 0 {
		difference *= -1
	}

	return difference
}
