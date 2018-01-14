package iris

import (
	"fmt"
	"math"
)

// Datum contains information observed about an individual iris plant.
type Datum struct {
	SepalLength, SepalWidth, PetalLength, PetalWidth float32
	Species                                          Species
}

// String implements the Stringer interface for Datum.
func (d Datum) String() string {
	return fmt.Sprintf(
		"Sepal Length: %0.1f, Sepal Width: %0.1f, Petal Length: %0.1f, Petal Width: %0.1f, Species: %q",
		d.SepalLength, d.SepalWidth, d.PetalLength, d.PetalWidth, d.Species.String())
}

// DistanceFrom returns the Euclidean distance between two Datums.
func (d Datum) DistanceFrom(other Datum) float64 {
	return math.Sqrt(
		math.Pow(float64(d.SepalLength-other.SepalLength), 2) +
			math.Pow(float64(d.SepalWidth-other.SepalWidth), 2) +
			math.Pow(float64(d.PetalLength-other.PetalLength), 2) +
			math.Pow(float64(d.PetalWidth-other.PetalWidth), 2))
}

// DistanceBetween returns the distance between two pieces of iris data.
func DistanceBetween(left, right interface{}) float64 {
	l := left.(Datum)
	r := right.(Datum)
	return l.DistanceFrom(r)
}

// Classification returns the species of a given iris.
func Classification(t interface{}) interface{} {
	d := t.(Datum)
	return d.Species
}
