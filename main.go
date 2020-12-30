package main

import (
	"fmt"
	"sentences/symbol"
)

func main() {
	h := symbol.NewSentence("Harry is smart", false)
	l := symbol.NewSentence("Dumbledor is away", false)
	p := symbol.NewSentence("Hermione is smart", false)

	o := symbol.Or(symbol.Or(h, l), p)
	fmt.Println(o)

	p.ChangeTruth(true)
	fmt.Println(o)

}
