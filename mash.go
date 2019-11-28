package uheprng

import (
	"math"
)

type Mash struct {
	state float64
}

func NewMash() *Mash {
	return &Mash{
		state: 0xefc8249d,
	}
}

func (m *Mash) Init() {
	m.state = 0xefc8249d
}

func (m *Mash) Next(data string) float64 {
	n := float64(m.state)
	for _, s := range data {
		n = n + float64(s)
		h := 0.02519603282416938 * n
		n = math.Floor(h)
		h -= n
		h *= n
		n = math.Floor(h)
		h -= n
		n += h * 0x100000000
	}
	m.state = n
	res := float64(int(n)%4294967296) * 2.3283064365386963e-10 // 2^-32
	return res
}
