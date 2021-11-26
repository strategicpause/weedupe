package reducer

import (
	"weedupe/common"
)

type Reducer interface {
	Reduce([]*common.Pair) map[string]int
}

type WordCountReducer struct {

}

func NewWordCountReducer() Reducer {
	return &WordCountReducer{}
}

func (r *WordCountReducer) Reduce(pairs []*common.Pair) map[string]int {
	data := map[string]int{}
	for _, pair := range pairs {
		key, ok := pair.A.(string)
		if !ok {
			continue
		}
		value, ok := pair.B.(int)
		if !ok {
			continue
		}
		data[key] += value
	}
	return data
}