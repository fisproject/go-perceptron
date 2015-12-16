package perceptron

import (
	"encoding/json"
	"fmt"
	"math"
)

type Data struct {
	Input_vector []float64 `json:"input_vector"`
	Label        int       `json:"label"`
}

type JSON struct {
	Training_set []Data `json:"training_set"`
}

type Perceptron struct {
	Threshold float64
	Eta       float64
	Weights   []float64
}

/*  bool to int */
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

/* Inner product */
func inner_product(i []float64, w []float64) (r float64) {
	for k, v := range i {
		r += v * w[k]
	}
	return r
}

func (self *Perceptron) Predict(input []float64) float64 {
	return inner_product(input, self.Weights)
}

func (self *Perceptron) Train(training_set_byte []byte) (err error) {

	var data JSON
	err = json.Unmarshal([]byte(training_set_byte), &data)
	if err != nil {
		return err
	}

	for true {
		fmt.Println("\n> Next Round")
		error_sum := 0.0

		for i := 0; i < len(data.Training_set); i++ {
			t := data.Training_set[i]

			actual_output := self.Threshold < inner_product(t.Input_vector, self.Weights) // Network

			e := float64(t.Label - btoi(actual_output))

			if e != 0.0 {
				error_sum += math.Abs(e)
				for j, v := range t.Input_vector {
					self.Weights[j] += self.Eta * e * v // Update the weights
				}
			}

			fmt.Println("Input vector : ", t.Input_vector, "Expected : ", t.Label, "Actual : ", btoi(actual_output))
			fmt.Println("Updated Weights", self.Weights)
		}

		fmt.Println("Error Sum : ", error_sum)

		if error_sum == 0.0 {
			break
		}
	}
	return err
}
