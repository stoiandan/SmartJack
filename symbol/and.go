package symbol

//And retruns a Sentence uniting input sentences with logical AND
func And(sentences ...*Sentence) *Sentence {

	andTruth := func(sentences ...*Sentence) bool {
		for _, sent := range sentences {
			if sent.TruthValue() == false {
				return false
			}
		}
		return true
	}

	return &Sentence{
		Description: JoinSentences(" ^ ", sentences...),
		TruthValue:  func() bool { return andTruth(sentences...) },
	}
}
