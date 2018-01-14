package iris

import (
	"strings"
	"testing"
)

func TestSpecies(t *testing.T) {
	if Unknown.String() != "" {
		t.Errorf("got %q, wanted \"\"", Unknown.String())
	}
	if !strings.Contains(strings.ToLower(Setosa.String()), "setosa") {
		t.Errorf("got %q, wanted \"setosa\" in the species", Setosa.String())
	}
	if !strings.Contains(strings.ToLower(Versicolor.String()), "versicolor") {
		t.Errorf("got %q, wanted \"versicolor\" in the species", Versicolor.String())
	}
	if !strings.Contains(strings.ToLower(Virginica.String()), "virginica") {
		t.Errorf("got %q, wanted \"virginica\" in the species", Virginica.String())
	}
	unk := Species(-1)
	if unk.String() != "" {
		t.Errorf("got %q, wanted \"\"", unk.String())
	}
}
