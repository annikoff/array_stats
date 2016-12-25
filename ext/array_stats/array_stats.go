package main

import (
	"C"
	"unsafe"
	"sort"
	"math"
)

//export fast_percentile
func fast_percentile(array unsafe.Pointer, size int, percent float64) float64 {
	elements := make([]float64, size)
	copy(elements[:], (*[1<<30]float64)(array)[:])
	sort.Float64s(elements)

	length := float64(size)

	if length == 0 {
		return 0
	}
	rank := (percent / 100) * (length + 1)
	_, rank_frac := math.Modf(rank)
	rank_frac = math.Abs(rank_frac)
	rank_truncated := int(math.Trunc(rank))

	if rank_frac > 0 {
		sample_0 := elements[rank_truncated - 1]
		sample_1 := elements[rank_truncated]
		return (rank_frac * (sample_1 - sample_0)) + sample_0
	}else {
		return elements[int(rank) - 1]
	}
}

func main() {}
