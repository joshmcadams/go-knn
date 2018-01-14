package knn

import (
	"container/list"
	"fmt"
)

// TopNList is a data structure that stores the a classification target's top-N neighbors weighted
// by minium distance.
type TopNList struct {
	n   int
	lst *list.List
}

// le contains a "list element" that is stored in the TopNList. The element consists of the
// distance used to weight the item in the list and the datum containing the information about the
// neighbor.
type le struct {
	distance float64
	neighbor interface{}
}

// NewTopNList creates a new TopNList that will retain up to n neighbors, keep the closest
// neighbors.
func NewTopNList(n int) (TopNList, error) {
	if n <= 0 {
		return TopNList{}, fmt.Errorf("got n=%d, must be >0", n)
	}
	return TopNList{n, list.New()}, nil
}

// Add adds a neighbor to the TopNList if the neighbor is one of the n-closest neighbors.
func (l *TopNList) Add(dist float64, neighbor interface{}) {
	if l.lst.Len() == 0 {
		l.lst.PushFront(le{dist, neighbor})
		return
	}
	for e := l.lst.Front(); e != nil; e = e.Next() {
		v := e.Value
		existing := v.(le)
		if existing.distance > dist {
			l.lst.InsertBefore(le{dist, neighbor}, e)
			break
		}
	}
	if l.lst.Len() > l.n {
		l.lst.Remove(l.lst.Back())
	}
}

// Iterate allows looping over the neighbors contained in the TopNList.
func (l *TopNList) Iterate() <-chan interface{} {
	c := make(chan interface{})
	go func() {
		for e := l.lst.Front(); e != nil; e = e.Next() {
			v := e.Value
			existing := v.(le)
			c <- existing.neighbor
		}
		close(c)
	}()
	return c
}

// GetClassification polls the top-n neighbors for their classification and then returns the most
// common classification found among the neighbors. The ClassificationExtractor is used to extract
// the classification from the neighbor.
func (l *TopNList) GetClassification(cf ClassificationFunc) interface{} {
	m := make(map[interface{}]int)
	for e := range l.Iterate() {
		m[cf(e)]++
	}
	max := 0
	var ret interface{}
	for f, count := range m {
		if count > max {
			ret = f
			max = count
		}
	}
	return ret
}
