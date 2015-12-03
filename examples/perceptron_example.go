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
	th := 0.5
	eta := 0.1
	w := []float64{0.0, 0.0, 0.0}

	err = perceptron.Learning(th, eta, w, training_set)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Done!")
}
