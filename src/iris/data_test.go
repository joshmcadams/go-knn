package iris

import (
	"testing"
)

func TestTotalExpectedDataPoints(t *testing.T) {
	if len(Data) != 150 {
		t.Errorf("got %d data points, wanted 150", len(Data))
	}
}

func TestTotalDataPointsPerSpecies(t *testing.T) {
	m := make(map[Species]int)
	for _, d := range Data {
		m[d.Species]++
	}
	for s, c := range m {
		if c != 50 {
			t.Errorf("got %d data points for %q, wanted 50", c, s.String())
		}
	}
}

func TestIndividualDataPointIntegrity(t *testing.T) {
	for i, d := range Data {
		if d.SepalLength <= 0.0 || d.SepalLength >= 10.0 {
			t.Errorf("got sepal length %f for sample %d, wanted a value in range (0.0, 10.0)", d.SepalLength, i)
		}
		if d.SepalWidth <= 0.0 || d.SepalWidth >= 10.0 {
			t.Errorf("got sepal width %f for sample %d, wanted a value in range (0.0, 10.0)", d.SepalWidth, i)
		}
		if d.PetalLength <= 0.0 || d.PetalLength >= 10.0 {
			t.Errorf("got petal length %f for sample %d, wanted a value in range (0.0, 10.0)", d.PetalLength, i)
		}
		if d.PetalWidth <= 0.0 || d.PetalWidth >= 10.0 {
			t.Errorf("got petal width %f for sample %d, wanted a value in range (0.0, 10.0)", d.PetalWidth, i)
		}
		switch d.Species {
		case Setosa:
			continue
		case Versicolor:
			continue
		case Virginica:
			continue
		default:
			t.Errorf("got species %d, wanted a known species", d.Species)
		}
	}
}
