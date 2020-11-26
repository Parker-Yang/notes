package insert

func Insert(data []int) {
	for i := range data {
		preIndex := i - 1
		current := data[i]
		for preIndex >= 0 && data[preIndex] > current {
			data[preIndex+1] = data[preIndex]
			preIndex -= 1
		}
		data[preIndex+1] = current
	}
}
