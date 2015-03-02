package perceptron

import (
	"encoding/json"
	"fmt"
)

var _weights []float64

type Data struct {
	Input_vector   []float64 `json:"input_vector"`
	Desired_output int       `json:"desired_output"`
}

type JSON struct {
	Training_set []Data `json:"training_set"`
}

/*  bool to int */
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

/* 内積 */
func dot_product(inputValues []float64, weights []float64) float64 {

	fmt.Println("input vector", inputValues)
	fmt.Println("weights", weights)

	var ret float64 = 0

	for i, v := range inputValues {
		ret += v * weights[i]
	}

	return ret
}

/* Learning */
func Learning(threshold float64, eta float64, weights []float64, training_set_byte []byte) (err error) {
	var condition bool = false

	_weights = weights

	var training_data JSON
	err = json.Unmarshal([]byte(training_set_byte), &training_data)
	if err != nil {
		return err
	}

	for condition == false {
		fmt.Println("\n> next step")
		error_count := 0

		for i := 0; i < len(training_data.Training_set); i++ {

			t := training_data.Training_set[i]
			result := dot_product(t.Input_vector, _weights)
			actual_output := result > threshold
			fmt.Println("expected: ", t.Desired_output, "actual: ", btoi(actual_output))

			error := t.Desired_output - btoi(actual_output)

			if error != 0 {
				error_count += 1
				for j, v := range t.Input_vector {
					// Update the weights
					_weights[j] += eta * float64(error) * v
				}
			}
		}
		if error_count == 0 {
			condition = true
		}
	}
	return nil
}
