package minitask1

var positives = [5]uint8{1, 14, 2, 3, 5}

// var sliceOfPositives = positives[:]

func BigNumber() uint8 {
	var max uint8
	for _, v := range positives {
		if v > max {
			max = v
		}
	}
	return max
}
