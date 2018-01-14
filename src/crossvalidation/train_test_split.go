// Package crossvalidation contains functions similar to the Python scikit cross_validation
// package.
package crossvalidation

import (
	"fmt"
	"math/rand"
)

// Splitable defines a set of data that can be split into a testing and training set.
type Splitable interface {
	ElementAt(index int) interface{}
	Length() int
}

// TrainTestSplit accepts a Splitable data set and divides that data set into a training and test
// data set. The training and test data sets are fed through the provided channels. There is also a
// random seed variable that can be used to reproduce the same set for testing.
//
// Note that the train and test channels passed into this function are closed by the function once
// it is done processing.
func TrainTestSplit(data Splitable, percentTest float32, seed int64, train, test chan interface{}) error {
	if percentTest < 0.0 || percentTest > 100.0 {
		return fmt.Errorf("got percent test %f, wanted value in range [0.0, 100.0]", percentTest)
	}

	testSize := int(float32(data.Length()) * percentTest)

	idxs := make(map[int]bool)
	r := rand.New(rand.NewSource(seed))
	for len(idxs) < testSize {
		idxs[r.Intn(data.Length())] = true
	}

	for i := 0; i < data.Length(); i++ {
		if _, ok := idxs[i]; ok {
			test <- data.ElementAt(i)
		} else {
			train <- data.ElementAt(i)
		}
	}
	close(test)
	close(train)

	return nil
}
