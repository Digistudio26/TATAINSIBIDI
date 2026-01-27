package sprint

func main() {
	fmt.Println(RemoveElementsInRange([]float64{10., .8, -.4, 20., 7.7, 3.}, 4, 1))
}

func RemoveElementsInRange(arr []float64, from, to int) []float64 {
	// Ensure correct order
	if from > to {
		from, to = to, from
	}

	
	if from < 0 {
		from = 0
	}
	if to > len(arr) {
		to = len(arr)
	}

	// If nothing to remove, return original array
	if from >= to {
		return arr
	}

	// Remove elements in range [from, to)
	return append(arr[:from], arr[to:]...)
}