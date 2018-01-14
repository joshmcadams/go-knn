package knn

// DistanceFunc is responsible for calculaing the distance between two neighbors.
type DistanceFunc func(a, b interface{}) float64

// ClassificationFunc returns the classification label for a given neighbor.
type ClassificationFunc func(t interface{}) interface{}

// DataSet provides the classifer the ability to iterate over data being used for training and/or
// classification.
type DataSet interface {
	Iterate() <-chan interface{}
}

// Classify is the main entry point for the KNN processing for an individual piece of data that
// needs to be classified. It requires the number of neighbors needed for classification, the data
// point being classified, the training data used to perform the classification, a method for
// calculating distance between neighbors, and a method for converting between classification
// types.
func Classify(numNeighbors int, point interface{}, training DataSet, df DistanceFunc, cf ClassificationFunc) interface{} {
	lst, err := NewTopNList(numNeighbors)
	if err != nil {
		panic(err)
	}

	for neighbor := range training.Iterate() {
		lst.Add(df(point, neighbor), neighbor)
	}

	return lst.GetClassification(cf)
}

// KNN performs K-Nearest Neighbors analysis on a testing data set given a training data set.
func KNN(numNeighbors int, testing, training DataSet, df DistanceFunc, cf ClassificationFunc) []interface{} {
	classifications := make([]interface{}, 0, 0)
	for datum := range testing.Iterate() {
		classifications = append(classifications, Classify(5, datum, training, df, cf))
	}
	return classifications
}
