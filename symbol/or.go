package symbol

//Or is the logical OR operator binding multiple sentneces and returning a new sentence represeting the input
// connected with OR
func Or(sentences ...*Sentence) *Sentence {

	orTruth := func(sentences ...*Sentence) bool {
		for _, sent := range sentences {
			if sent.truthValue() == true {
				return true
			}
		}
		return false
	}
	return &Sentence{
		Description: JoinSentences(" V ", sentences...),
		truthValue:  func() bool { return orTruth(sentences...) },
	}
}
