package calclib

import (
	"testing"
)

func TestCalculateMin(t *testing.T) {

	items := []float64{50, 20, 10, 40, 5}
	data := InputWrapper{items, 1}
	res := CalculateMin(data)
	if res[0] != 5 {
		t.Fatalf(`TestCalculate fail obtained %f and expected %d`, res[0], 5)
	}
}

func TestCalculateMax(t *testing.T) {

	items := []float64{50, 20, 10, 40, 5}
	data := InputWrapper{items, 1}
	res := CalculateMax(data)
	if res[0] != 50 {
		t.Fatalf(`TestCalculateMax fail obtained %f and expected %d`, res[0], 5)
	}
}

func TestCalculateAVG(t *testing.T) {
	items := []float64{20, 10, 0}
	res := CalculateAvg(items)
	expected := float64(10)
	if res != expected {
		t.Fatalf(`TestCalculateAVG fail obtained %f and expected %f`, res, expected)
	}
}

func TestCalculateMedianOdd(t *testing.T) {
	items := []float64{20, 10, 40}
	res := CalculateMedian(items)
	expected := float64(20)
	if res != expected {
		t.Fatalf(`TestCalculateMedianOdd fail obtained %f and expected %f`, res, expected)
	}
}

func TestCalculateMedianEven(t *testing.T) {
	items := []float64{20, 10, 40, 5}
	res := CalculateMedian(items)
	expected := float64(15)
	if res != expected {
		t.Fatalf(`TestCalculateMedianEven fail obtained %f and expected %f`, res, expected)
	}
}

func TestCalculatePercentileFailEmpty(t *testing.T) {

	items := []float64{}
	data := PercentileInputWrapper{items, 1.0}
	res, err := CalculatePercentile(data)
	if err.Error() != "empty slices" {
		t.Fatalf(`TestCalculatePercentileFailEmpty list is not empty %f`, res)
	}
}

func TestCalculatePercentileFailHighter100(t *testing.T) {

	items := []float64{10, 20, 30}
	data := PercentileInputWrapper{items, 100.1}
	res, err := CalculatePercentile(data)
	if err.Error() != "outside of percentile limites" {
		t.Fatalf(`TestCalculatePercentileFailEmpty list is not empty %f`, res)
	}
}

func TestCalculatePercentileFailLess0(t *testing.T) {

	items := []float64{10, 20, 30}
	data := PercentileInputWrapper{items, -0.1}
	res, err := CalculatePercentile(data)
	if err.Error() != "outside of percentile limites" {
		t.Fatalf(`TestCalculatePercentileFailEmpty list is not empty %f`, res)
	}
}

func TestCalculatePercentile(t *testing.T) {

	items := []float64{10, 20, 30}
	data := PercentileInputWrapper{items, 15}
	res, err := CalculatePercentile(data)
	if res != 10 {
		t.Fatalf(`TestCalculatePercentile fail %q`, err.Error())
	}
}
