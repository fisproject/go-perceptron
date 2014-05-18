package perceptron

import (
	"encoding/json"
	"fmt"
	//"math"
	"strconv"
	"strings"
)

var _weights []float64

type Datas struct {
	Input_vector   string
	Desired_output int
}

/* json []byte to golangObj */
type JsonParseObj struct {
	Training_set_count int
	Training_set_1 Datas
	Training_set_2 Datas
	Training_set_3 Datas
	Training_set_4 Datas
}

/*  string to []string to []float64  */
func stof(value string) []float64 {
	val := make([]float64, 0)
	s := strings.Split(value, ",")
	for _, t := range s {
		if t != "" && t != "\t" {
			n, err := strconv.ParseFloat(t, 64)
			if err != nil {
				fmt.Println(err)
				return []float64{-1}
			}
			val = append(val, n)
		}
	}
	return val
}

/*  bool to int */
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

/* getTarget */
func getTarget(obj JsonParseObj, index int) Datas {
	switch {
	case index == 0:
		return obj.Training_set_1
	case index == 1:
		return obj.Training_set_2
	case index == 2:
		return obj.Training_set_3
	case index == 3:
		return obj.Training_set_4
	default:
		return obj.Training_set_1
	}
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
func Learning(threshold float64, eta float64, weights []float64, training_set []byte) bool {

	_weights = weights

	training_data := &JsonParseObj{}
	if err := json.Unmarshal([]byte(training_set), training_data); err != nil {
		fmt.Println("Can't parse json")
		return false
	}

	fmt.Println(training_data)

	condition := false

	for condition == false {
		fmt.Println("-------------------------------------------")
		error_count := 0

		for i := 0; i < training_data.Training_set_count; i++ {

			t := getTarget(*training_data, i)
			value := stof(t.Input_vector)
			result := dot_product(value, _weights)
			actual_output := result > threshold
			fmt.Println("desired_output:", t.Desired_output, "actual_output:", btoi(actual_output))

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

	return condition
}
