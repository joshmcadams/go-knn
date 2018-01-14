package crossvalidation

import (
	"sync"
	"testing"
)

type splitable []int

func (s splitable) Length() int {
	return len(s)
}

func (s splitable) ElementAt(i int) interface{} {
	return s[i]
}

func TestTrainTesteSplit(t *testing.T) {
	data := make([]int, 1000, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = i
	}

	tests := []struct {
		name                string
		testPercent         float32
		trainSize, testSize int
	}{
		{"even split", 0.5, 500, 500},
		{"all test", 1.0, 0, 1000},
		{"all train", 0.0, 1000, 0},
		{"75% training", 0.25, 750, 250},
	}

	for _, test := range tests {
		trainingSet := make([]int, 0, test.trainSize)
		testingSet := make([]int, 0, test.testSize)

		trainChan := make(chan interface{})
		testChan := make(chan interface{})

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			for i := range trainChan {
				if ii, ok := i.(int); ok {
					trainingSet = append(trainingSet, ii)
				}
			}
		}()

		go func() {
			defer wg.Done()
			for i := range testChan {
				if ii, ok := i.(int); ok {
					testingSet = append(testingSet, ii)
				}
			}
		}()

		if err := TrainTestSplit(splitable(data), test.testPercent, 123, trainChan, testChan); err != nil {
			t.Errorf("got unexpected error: %v", err)
		}

		wg.Wait()
		if len(trainingSet) != test.trainSize {
			t.Errorf("got %d training elements when expecting %d", len(trainingSet), test.trainSize)
		}
		if len(testingSet) != test.testSize {
			t.Errorf("got %d testing elements when expecting %d", len(testingSet), test.testSize)
		}
	}
}
