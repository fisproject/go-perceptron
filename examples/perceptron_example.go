package main

import (
	"../"
	"fmt"
	"io/ioutil"
)

func main() {

	// A perceptron learns to perform a binary NAND function
	training_set, err := ioutil.ReadFile("../data/input.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Initialise the weights and the threshold
	th := 0.5
	eta := 0.1
	w := []float64{0.0, 0.0, 0.0}
	p := perceptron.Perceptron{th, eta, w}

	train_data, err := p.Load(training_set)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Train
	err = p.Train(train_data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\nTrain Done!")

	// Predict unknown input
	x := []float64{1, 0, 1}
	y := p.Predict(x)
	fmt.Println(y)
}
