package symbol

//And retruns a Sentence uniting input sentences with logical AND
func And(sentences ...*Sentence) *Sentence {

	andTruth := func(sentences ...*Sentence) bool {
		for _, sent := range sentences {
			if sent.truthValue() == false {
				return false
			}
		}
		return true
	}

	return &Sentence{
		Description: JoinSentences(" ^ ", sentences...),
		truthValue:  func() bool { return andTruth(sentences...) },
	}
}
