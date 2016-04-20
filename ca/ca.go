package ca

import (
	"fmt"
	"math/rand"

	"github.com/jgcarvalho/zeca2/rule"
)

type Config struct {
	InitState []string
	EndState  []string
	// 	TransStates    []string `toml:"transition-states"`
	// 	Hydrophobicity string   `toml:"hydrophobicity"`
	// 	R              int      `toml:"r"`
	Steps int `toml:"steps"`
	// Consensus int `toml:"consensus"`
	IgnoreSteps int `toml:"ignore-steps"`
}

func (conf Config) Run(rule rule.Rule) {
	var init, end, previous, current []string
	init = make([]string, len(conf.InitState))
	end = make([]string, len(conf.EndState))
	copy(init, conf.InitState)
	copy(end, conf.EndState)
	if len(init) != len(end) {
		panic("Init and End States have diffent lenghts")
	}
	previous = make([]string, len(conf.InitState))
	copy(previous, init)
	current = make([]string, len(init))

	// set begin and end equals to # (static states)
	current[0], current[len(init)-1] = "#", "#"

	for i := 0; i < conf.Steps; i++ {

		if i%2 == 0 {
			step(&previous, &current, &init, &end, &rule)
			fmt.Println(current)
		} else {
			step(&current, &previous, &init, &end, &rule)
			fmt.Println(previous)
		}
	}
}

func step(previous, current, init, end *[]string, ru *rule.Rule) {
	var rnd float64
	var state string
	for c := 1; c < len(*init)-1; c++ {
		rnd = rand.Float64()
		for k, v := range (*ru)[rule.Pattern{(*previous)[c-1], (*previous)[c], (*previous)[c+1]}] {
			if v > rnd {
				state = k
				break
			} else {
				rnd -= v
			}
		}
		(*current)[c] = state
	}
}
