package generator

import (
	"math"
)

//TruthTableModels generates all 2*n truth table rows (models) for a n variable truth table
func TruthTableModels(variables int) <-chan []bool {
	end := int(math.Pow(2, float64(variables)))
	ch := make(chan []bool, end)
	defer close(ch)

	generatorHelper(0, make([]bool, variables), ch)
	return ch
}

func generatorHelper(pos int, model []bool, ch chan []bool) {
	if pos == len(model) {
		ans := make([]bool, pos)
		copy(ans, model)
		ch <- ans
		return
	}

	for _, val := range []bool{true, false} {
		model[pos] = val
		generatorHelper(pos+1, model, ch)
	}
}
