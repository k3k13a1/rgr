package function

type Function struct {
	A int `json:"a"`
	B int `json:"b"`
	C int `json:"c"`
	D int `json:"d"`
}

func New(a, b, c, d int) *Function {
	return &Function{
		A: a,
		B: b,
		C: c,
		D: d,
	}
}

func (f *Function) FuzzNums() map[int]float32 {

	mu := make(map[int]float32)

	for i := f.A; i <= f.D; i++ {
		mu[i] = f.getMu(i)
	}

	return mu
}

func (f *Function) getMu(x int) float32 {
	if x <= f.A {
		return 0.0
	} else if f.A <= x && x <= f.B {
		return (float32(x) - float32(f.A)) / (float32(f.B) - float32(f.A))
	} else if f.C <= x && x <= f.D {
		return (float32(f.D) - float32(x)) / (float32(f.D) - float32(f.C))
	} else if f.B < x && x < f.C {
		return 1.0
	} else if f.D < x {
		return 0.0
	}
	return 0.0
}
