package iris

// Species represents the iris species.
type Species int

const (
	// Unknown represents an unknow iris species.
	Unknown = Species(iota)
	// Setosa represents the I. setosa iris species.
	Setosa = Species(iota)
	// Versicolor represents the I. versicolor iris species.
	Versicolor = Species(iota)
	// Virginica represents the I. virginica iris species.
	Virginica = Species(iota)
)

// String implements the Stringer interface for Species.
func (s Species) String() string {
	switch s {
	case Setosa:
		return "I. setosa"
	case Versicolor:
		return "I. versicolor"
	case Virginica:
		return "I. virginica"
	default:
		return ""
	}
}
