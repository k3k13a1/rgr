package defuzz

func MiddleMaximum(m map[int]float32) float32 {
	var sum float32
	for i, v := range m {
		sum += v * float32(i)
	}

	var total float32
	for _, v := range m {
		total += v
	}

	return sum / total
}
