package perceptron

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

var _weights []float64

type Datas struct {
	Input_vector   string `json:"input_vector"`
	Desired_output int    `json:"desired_output"`
}

type JSON struct {
	Training_set_count int     `json:"training_set_count"`
	Training_set       []Datas `json:"training_set"`
}

/*  string to []string to []float64  */
func stof(value string) (val []float64, err error) {
	val = make([]float64, 0)
	s := strings.Split(value, ",")
	for _, t := range s {
		if t != "" && t != "\t" {
			n, err := strconv.ParseFloat(t, 64)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			val = append(val, n)
		}
	}
	return val, nil
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

	fmt.Println("input values", inputValues)
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
		fmt.Println("> next step")
		error_count := 0

		for i := 0; i < training_data.Training_set_count; i++ {

			t := training_data.Training_set[i]
			value, err := stof(t.Input_vector)
			if err != nil {
				return err
			}
			result := dot_product(value, _weights)
			actual_output := result > threshold
			fmt.Println("desired_output: ", t.Desired_output, "actual_output: ", btoi(actual_output))

			/* judgement */
			error := t.Desired_output - btoi(actual_output)

			if error != 0 {
				error_count += 1
				for j, v := range value {
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
