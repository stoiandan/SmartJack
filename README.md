# SmartJack

SmartJack is a small knowledge-Based System (think Prolog, but dumber). That's able to take a KB and respond to queries with either `true`, `false` or `inconcludent`


## Examples

For example let's see the engine in action by creating a world with two variables

```
  let d = "Dan is playing at the PC"
  let a = "Anatasia is playing a the"
```

Given our two varaibles, what we now want is a K.B. (kwnowledge base) in which we have *facts*, i.e. true and false sentences about our world:

```
 let KB []Sentences // a knowledge base consists of a series of facts
 append(KB, d OR a) // We know that either d or a are TRUE
```

Given our KB and our variables, we can know use the engine to query, i.e. ask questions about our world.
The engine will return one of three answers:

- *false* -This means that given what we know (the KB ), we can reason that the answer to the question is false
- *true* - same as above, only this time the answer is true
- *inconcludent* - The answer is uncertain, this means there are possible worlds, aka possible states of our varaibles, that don't contradict the KB, for which the answer is true, but not for *all* of them

There is a concrete example in the `main.go` file where the present example is written down in Golang.
