package main


func BalanceOut(arr []bool) []bool {
	trueCount := 0
	falseCount := 0

	for _, v := range arr {
		if v {
			trueCount++
		} else {
			falseCount++
		}
	}

	if trueCount > falseCount {
		diff := trueCount - falseCount
		for i := 0; i < diff; i++ {
			arr = append(arr, false)
		}
	} else if falseCount > trueCount {
		diff := falseCount - trueCount
		for i := 0; i < diff; i++ {
			arr = append(arr, true)
		}
	}

	return arr
}