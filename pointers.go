package workshop

import (
	"math"
)

type avg []float64

func (nums *avg) average() float64 {
	var sum float64
	for _, num := range *nums {
		sum += num
	}
	return round(sum / float64(len((*nums))), 2)
}

func round(val float64, prec int) float64 {
	var rounder float64
	intermed := val * math.Pow(10, float64(prec))

	if val >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}
	return rounder / math.Pow(10, float64(prec))
}

func pointers() {
	i, j := 3, 4
	var a = &i
	var b = &j
	assert(i == *a)
	assert(j == *b)

	*a = 10
	assert(i == *a)

	var num *int
	assert(num == nil)
	var numbers = avg{2.3, 3.5, 7.5, 2.5, 7.8}
	assert(numbers.average() == 4.73)
}
