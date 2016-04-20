package main

import (
	"github.com/jgcarvalho/zeca2/ca"
	"github.com/jgcarvalho/zeca2/rule"
)

func main() {
	r := rule.Read("/home/jgcarvalho/sync/data/anglesdb/AAA.probs")
	// fmt.Println(r)
	conf := ca.Config{
		InitState:   []string{"#", "A??", "A??", "A??", "A??", "A??", "A??", "A??", "A??", "A??", "A??", "A??", "A??", "A??", "A??", "A??", "A??", "A??", "A??", "A??", "A??", "#"},
		EndState:    []string{"#", "Aa0", "Aa0", "Aa0", "Aa0", "Aa0", "Aa0", "Aa0", "Aa0", "Aa0", "Aa0", "Aa0", "Aa0", "Aa0", "Aa0", "Aa0", "Aa0", "Aa0", "Aa0", "Aa0", "Aa0", "#"},
		Steps:       100000,
		IgnoreSteps: 1,
	}
	conf.Run(r)
}
