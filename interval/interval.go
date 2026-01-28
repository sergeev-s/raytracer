package interval

import (
	"math"
)

type Interval struct {
	Min, Max float64
}

func NewInterval(min, max float64) Interval {
	return Interval{min, max}
}

func GetEmpty() Interval {
	return Interval{Min: math.Inf(1), Max: math.Inf(-1)}
}

func GetUniverse() Interval {
	return Interval{Min: math.Inf(-1), Max: math.Inf(1)}
}

func Size(interval Interval) float64 {
	return interval.Max - interval.Min
}

func (interval Interval) Contains(x float64) bool {
	return interval.Min <= x && x <= interval.Max
}

func (interval Interval) Surrounds(x float64) bool {
	return interval.Min < x && x < interval.Max
}

func (interval Interval) Clamp(x float64) float64 {
	if x < interval.Min {
		return interval.Min
	}
	if x > interval.Max {
		return interval.Max
	}
	return x
}
