package symbol

import "fmt"

type TruthFunc func() bool

//Sentence is a human readable text
type Sentence struct {
	Description string
	truthValue  TruthFunc
}

func NewSentence(description string, truthValue bool) *Sentence {
	return &Sentence{
		Description: description,
		truthValue:  func() bool { return truthValue },
	}
}

func (s Sentence) String() string {
	return s.Description + fmt.Sprintf(" has the truth value: %v", s.truthValue())
}

//ChangeTruth changes the truth value of a sentence
func (s *Sentence) ChangeTruth(truth bool) {
	s.truthValue = func() bool { return truth }
}
