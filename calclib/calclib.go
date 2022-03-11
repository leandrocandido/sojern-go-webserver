package calclib

import (
	"errors"
	"math"
	"sort"
)

type InputWrapper struct {
	List       []float64
	Quantifier int
}

type PercentileInputWrapper struct {
	List       []float64
	Percentile float64
}

func CalculateMin(data InputWrapper) []float64 {
	sort.Float64s(data.List)
	res := data.List[0:data.Quantifier]
	return res
}

func CalculateMax(data InputWrapper) []float64 {
	sort.Float64s(data.List)
	size := len(data.List)
	res := data.List[size-data.Quantifier : size]
	return res
}

func CalculateAvg(list []float64) float64 {

	size := len(list)
	sum := 0.0
	for i := 0; i < size; i++ {
		sum += list[i]
	}
	avg := (float64(sum)) / (float64(size))
	return avg
}

func CalculateMedian(list []float64) float64 {
	sort.Float64s(list)
	med := len(list)
	res := 0.0

	if med%2 == 0 {
		res = (list[(med-1)/2] + list[med/2]) / 2
	} else {
		res = list[med/2]
	}

	return res
}

func CalculatePercentile(data PercentileInputWrapper) (float64, error) {
	// Find the length of items in the slice
	il := len(data.List)

	// Return an error for empty slices
	if il == 0 {
		return -1, errors.New("empty slices")
	}

	// Return error for less than 0 or greater than 100 percentages
	if data.Percentile < 0 || data.Percentile > 100 {
		return -2, errors.New("outside of percentile limites")
	}

	// Start by sorting a copy of the slice
	sort.Float64s(data.List)

	// Return the last item
	res := 0.0
	if data.Percentile == 100.0 {
		res = data.List[il-1]
	} else {
		// Find ordinal ranking
		or := int(math.Ceil(float64(il) * data.Percentile / 100))

		// Return the item that is in the place of the ordinal rank
		if or == 0 {
			res = data.List[0]
		}
		res = data.List[or-1]
	}
	return res, nil
}
