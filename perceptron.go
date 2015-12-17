package perceptron

import (
	"encoding/json"
	"fmt"
	"math"
)

type Data struct {
	Feats []float64 `json:"feats"`
	Label int       `json:"label"`
}

type DataSet struct {
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

func (self *Perceptron) Load(training_set_byte []byte) (d DataSet, err error) {
	err = json.Unmarshal([]byte(training_set_byte), &d)
	return d, err
}

func (self *Perceptron) Predict(input []float64) float64 {
	return inner_product(input, self.Weights)
}

func (self *Perceptron) Train(d DataSet) (err error) {

	for true {
		fmt.Println("\n> Next Round")
		error_sum := 0.0

		for i := 0; i < len(d.Training_set); i++ {
			t := d.Training_set[i]

			actual_output := self.Threshold < inner_product(t.Feats, self.Weights) // Network

			e := float64(t.Label - btoi(actual_output))

			if e != 0.0 {
				error_sum += math.Abs(e)
				for j, v := range t.Feats {
					self.Weights[j] += self.Eta * e * v // Update the weights
				}
			}

			fmt.Println("Input vector : ", t.Feats, "Expected : ", t.Label, "Actual : ", btoi(actual_output))
			fmt.Println("Updated Weights", self.Weights)
		}

		fmt.Println("Error Sum : ", error_sum)

		if error_sum == 0.0 {
			break
		}
	}
	return err
}
