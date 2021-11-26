package combiner

type Combiner interface {
	Combine(map[string]int)
	GetResults() map[string]int
}

type  WordCountCombiner struct {
	results map[string]int
}

func NewWordCountCombiner() Combiner {
	return &WordCountCombiner{
		results: map[string]int{},
	}
}

func (w *WordCountCombiner) Combine(d1 map[string]int) {
	for k, v := range d1 {
		w.results[k] = v
	}
}

func (w *WordCountCombiner) GetResults() map[string]int {
	return w.results
}