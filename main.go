package main

import (
	"fmt"
	"weedupe/mapper"
	"weedupe/combiner"
	"weedupe/director"
	"weedupe/provider"
	"weedupe/reducer"
)

func main() {
	p := provider.NewFileInputProvider([]string {"data/alice.txt", "data/dracula.txt", "data/gatsby.txt"})
	m := mapper.NewWordCountMapper()
	r := reducer.NewWordCountReducer()
	c := combiner.NewWordCountCombiner()

	director := director.NewDirector(p, m, r, c)
	err := director.Start()
	if err != nil {
		fmt.Println(err)
	}
	results := c.GetResults()
	fmt.Println(results)
}
