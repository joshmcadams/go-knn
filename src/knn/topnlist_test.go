package knn

import (
	"testing"
)

func TestTopNList(t *testing.T) {
	l, err := NewTopNList(3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	l.Add(9.0, "a")
	want(t, l, []string{"a"})
	l.Add(8.0, "b")
	want(t, l, []string{"b", "a"})
	l.Add(7.0, "c")
	want(t, l, []string{"c", "b", "a"})
	l.Add(6.0, "d")
	want(t, l, []string{"d", "c", "b"})
	l.Add(5.0, "e")
	want(t, l, []string{"e", "d", "c"})
	l.Add(4.0, "f")
	want(t, l, []string{"f", "e", "d"})
	l.Add(9.0, "a")
	want(t, l, []string{"f", "e", "d"})
}

func want(t *testing.T, l TopNList, want []string) {
	got := make([]string, 0, 3)
	for e := range l.Iterate() {
		chr, ok := e.(string)
		if !ok {
			t.Fatalf("unable to convert %v to string", e)
		}
		got = append(got, chr)
	}

	if len(got) != len(want) {
		t.Errorf("got %v, want %v", got, want)
		return
	}

	match := true
	for i, v := range want {
		if got[i] != v {
			match = false
		}
	}

	if !match {
		t.Errorf("got %v, want %v", got, want)
	}
}
