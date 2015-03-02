package main

import (
	"../"
	"fmt"
	"io/ioutil"
)

func main() {
	
	/* A perceptron learns to perform a binary NAND function  */
	training_set, err := ioutil.ReadFile("../json/input.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	/* Initialise the weights and the threshold */
	threshold := 0.5
	learning_rate := 0.1
	weights := []float64{0.0, 0.0, 0.0}

	err = perceptron.Learning(threshold, learning_rate, weights, training_set)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Completed!")
}
