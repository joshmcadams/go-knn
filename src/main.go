package main

import (
	"flag"
	"fmt"
	"iris"
	"knn"
	"os"
	"time"
)

var neighbors = flag.Int("neighbors", 5, "number of neighbors to poll for classification")
var randomSeed = flag.Int64("random_seed", 0, "seed used for randomization of the data set division")
var percentTest = flag.Float64("percent_test", 0.25, "percent of the data set used for testing")

func main() {
	flag.Parse()

	if *randomSeed == 0 {
		*randomSeed = time.Now().Unix()
	}

	fmt.Printf("Creating test (%0.2f) and trianing (%0.2f) data sets using random seed %d\n", *percentTest, 1.0-*percentTest, *randomSeed)
	train, test, err := iris.CreateTrainingAndTestDataSets(float32(*percentTest), *randomSeed)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Created a test set of size %d and training set of size %d\n", len(test), len(train))

	fmt.Printf("Performing KNN classification using %d neighbors\n", *neighbors)
	classifications := knn.KNN(*neighbors, iris.Irises(test), iris.Irises(train), iris.DistanceBetween, iris.Classification)

	correct := 0
	for i, d := range test {
		if d.Species != classifications[i] {
			fmt.Printf("... got classification of %q, wanted %q\n", classifications[i], d.Species)
		} else {
			correct++
		}
	}
	fmt.Printf("Got %d of %d classifications correct (%0.2f accuracy)\n", correct, len(test), float32(correct)/float32(len(test)))
}
