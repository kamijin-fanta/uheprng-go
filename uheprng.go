package uheprng

import (
	"math"
	"strconv"
)

const ORDER = 48

type UhePrng struct {
	carry float64
	phase int
	state [ORDER]float64
	mash  *Mash
}

func NewUhePrng() *UhePrng {
	r := &UhePrng{}
	r.Init()
	return r
}

func (prng *UhePrng) Init() {
	if prng.mash == nil {
		prng.mash = NewMash()
	} else {
		prng.mash.Init()
	}
	prng.carry = 1
	prng.phase = ORDER
	for i := range prng.state {
		prng.state[i] = prng.mash.Next(" ")
	}
}
func (prng *UhePrng) Seed(seed string) {
	// todo clean string
	prng.Init()
	prng.mash.Next(seed)
	for _, s := range seed {
		for i := range prng.state {
			prng.state[i] -= prng.mash.Next(strconv.Itoa(int(s)))
			if prng.state[i] < 0 {
				prng.state[i] += 1
			}
		}
	}
}
func (prng *UhePrng) rawPrng() float64 {
	prng.phase++
	if prng.phase >= ORDER {
		prng.phase = 0
	}
	t := 1768863*prng.state[prng.phase] + prng.carry*2.3283064365386963e-10 // 2^-32
	prng.carry = math.Floor(t)
	prng.state[prng.phase] = t - prng.carry
	res := prng.state[prng.phase]
	return res
}

func (prng *UhePrng) Next(numRange int) int {
	r := float64(numRange)
	r = math.Floor(r * (prng.rawPrng() + math.Floor(prng.rawPrng()*0x200000)*1.1102230246251565e-16)) // 2^-53
	return int(r)
}
