package generator

import (
	"fmt"
	"math"
)

func GenerateModel(variables int) <-chan []bool {
	end := int(math.Pow(2, float64(variables)))
	ch := make(chan []bool, end)
	defer close(ch)

	generatorHelper(0, make([]bool, variables), ch)
	return ch
}

func generatorHelper(pos int, model []bool, ch chan []bool) {
	if pos == len(model) {
		ch <- model
		fmt.Printf(" %v \n", model)
		return
	}

	for _, val := range []bool{true, false} {
		model[pos] = val
		generatorHelper(pos+1, model, ch)
	}
}
