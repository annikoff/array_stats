package main

import (
	"C"
	"unsafe"
	"sort"
	"math"
)

//export fast_percentile
func fast_percentile(array unsafe.Pointer, size int, percent float64) float64 {
	elements := (*[1<<30]float64)(array)[:size]
	sort.Float64s(elements)

	rank := (percent / 100) * (float64(size) + 1)
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
