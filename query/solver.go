package query

import (
	"sentences/generator"
	"sentences/symbol"
)

//Ask queries the the KB for a given question
func Ask(knowledgeBase []*symbol.Sentence,
	variables []*symbol.Sentence,
	question *symbol.Sentence) string {

	ln := len(variables)
	ans := make([]bool, ln)
	possible := false

	for model := range generator.TruthTableModels(ln) {
		for i := 0; i < ln; i++ {
			variables[i].ChangeTruth(model[i])
		}

		if verifyKB(variables) == true {
			possible = true
			ans = append(ans, question.TruthValue())
		} else {
			continue
		}

	}

	if possible == false {
		return "false"
	}

	for _, truth := range ans {
		if truth == false {
			return "inconcludent! Can be possible, but not always, can't know for sure"
		}
	}
	return "true"
}

func verifyKB(kb []*symbol.Sentence) bool {
	for _, sent := range kb {
		if sent.TruthValue() == false {
			return false
		}
	}
	return true
}
