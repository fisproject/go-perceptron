package perceptron

import (
	"../"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestLearning(t *testing.T) {
	training_set, err := ioutil.ReadFile("../json/input.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	/* Initialise the weights and the threshold */
	threshold := 0.5
	learning_rate := 0.1
	weights := []float64{0.0, 0.0, 0.0}

	actual := perceptron.Learning(threshold, learning_rate, weights, training_set)
	var expected error = nil
	if actual != expected {
		t.Errorf("got %v,want %v", actual, expected)
	}
}
