package iris

import (
	"testing"
)

func TestDatum(t *testing.T) {
	if len(Data) != 150 {
		t.Errorf("got %d data points, wanted 150", len(Data))
	}

	perSpecies := make(map[Species]int)
	for _, d := range Data {
		perSpecies[d.Species]++
	}
	for s, c := range perSpecies {
		if c != 50 {
			t.Errorf("got %d data points for %q, wanted 50", c, s.String())
		}
	}
}

func TestDistanceFrom(t *testing.T) {
	tests := []struct {
		name   string
		d1, d2 Datum
		want   float64
	}{
		{"zero-value data", Datum{0.0, 0.0, 0.0, 0.0, Setosa}, Datum{0.0, 0.0, 0.0, 0.0, Setosa}, 0.0},
		{"same-value data", Datum{1.0, 1.0, 1.0, 1.0, Setosa}, Datum{1.0, 1.0, 1.0, 1.0, Setosa}, 0.0},
		{"different sepal length", Datum{1.0, 1.0, 1.0, 1.0, Setosa}, Datum{2.0, 1.0, 1.0, 1.0, Setosa}, 1.0},
		{"different sepal width", Datum{1.0, 1.0, 1.0, 1.0, Setosa}, Datum{1.0, 2.0, 1.0, 1.0, Setosa}, 1.0},
		{"different petal length", Datum{1.0, 1.0, 1.0, 1.0, Setosa}, Datum{1.0, 1.0, 2.0, 1.0, Setosa}, 1.0},
		{"different petal width", Datum{1.0, 1.0, 1.0, 1.0, Setosa}, Datum{1.0, 1.0, 1.0, 2.0, Setosa}, 1.0},
		{"everything different", Datum{1.0, 1.0, 1.0, 1.0, Setosa}, Datum{2.0, 2.0, 2.0, 2.0, Setosa}, 2.0},
	}
	for _, test := range tests {
		got := test.d1.DistanceFrom(test.d2)
		if test.want != got {
			t.Errorf("got distance %f, wanted %f", got, test.want)
		}
	}
}
