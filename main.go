package main

import (
	"fmt"
	"sentences/query"
	"sentences/symbol"
)

func main() {
	//example of how to use the KB engine
	d := symbol.NewSentence("Dan is palying at PC", true)
	a := symbol.NewSentence("Anastasia is playing at PC", true)

	kb1 := symbol.Or(d, a)

	kb := []*symbol.Sentence{kb1}

	answer := query.Ask(kb, []*symbol.Sentence{d, a}, d)
	fmt.Println(answer)

}
