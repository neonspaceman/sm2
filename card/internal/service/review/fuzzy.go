package review

import (
	"math"
	"math/rand"
)

type fuzzyRang struct {
	StartInDays float64
	EndInDays   float64
	Factor      float64
}

type fuzzy struct {
	ranges                []fuzzyRang
	minimumIntervalInDays int64
	maximumIntervalInDays int64
}

func newFuzzy(minimumInterval, maximumInterval int64) *fuzzy {
	return &fuzzy{
		minimumIntervalInDays: minimumInterval,
		maximumIntervalInDays: maximumInterval,
		ranges: []fuzzyRang{
			{StartInDays: 2.5, EndInDays: 7.0, Factor: 0.15},
			{StartInDays: 7.0, EndInDays: 20.0, Factor: 0.1},
			{StartInDays: 20.0, EndInDays: math.Inf(1), Factor: 0.05},
		},
	}
}

func (f *fuzzy) get(intervalInDays int64) int64 {
	if float64(intervalInDays) < 2.5 {
		return intervalInDays
	}

	minInterval, maxInterval := f.getRange(intervalInDays)

	fuzzedInterval := (rand.Float64() * float64(maxInterval-minInterval+1)) + float64(minInterval)

	return min(int64(math.Round(fuzzedInterval)), f.maximumIntervalInDays)
}

func (f *fuzzy) getRange(intervalInDays int64) (int64, int64) {
	delta := 1.0

	for _, val := range f.ranges {
		delta += val.Factor * max(min(float64(intervalInDays), val.EndInDays)-val.StartInDays, 0.0)
	}

	minInterval := int64(math.Round(float64(intervalInDays) - delta))
	maxInterval := int64(math.Round(float64(intervalInDays) + delta))

	// make sure the minInterval and maxInterval fall into a valid range
	minInterval = max(2, minInterval)
	maxInterval = min(maxInterval, f.maximumIntervalInDays)
	minInterval = min(minInterval, maxInterval)

	return minInterval, maxInterval
}
