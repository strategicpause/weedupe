package director

import (
	"sync"
	"weedupe/combiner"
	"weedupe/mapper"
	"weedupe/provider"
	"weedupe/reducer"
)

type Director struct {
	provider provider.InputProvider
	mapper mapper.Mapper
	reducer reducer.Reducer
	combiner combiner.Combiner
	resultChan chan map[string]int
	files []string
	waitGroup sync.WaitGroup
}

func NewDirector(p provider.InputProvider, m mapper.Mapper, r reducer.Reducer, c combiner.Combiner) *Director {
	d := Director {
		provider: p,
		mapper: m,
		reducer: r,
		combiner: c,
		resultChan: make(chan map[string]int),
	}
	go d.handler()

	return &d
}

func (d *Director) Load(files []string) {
	d.files = files
}

func (d *Director) Start() error {
	for d.provider.HasNext() {
		content, err := d.provider.GetNext()
		if err != nil {
			return err
		}
		d.waitGroup.Add(1)
		go d.MapReduce(content)
	}
	d.waitGroup.Wait()
	return nil
}

func (d *Director) MapReduce(content string) {
	pairs := d.mapper.Map(content)
	d.resultChan <- d.reducer.Reduce(pairs)
}

func (d *Director) handler() {
	for {
		select {
		case data := <-d.resultChan:
			d.combiner.Combine(data)
			d.waitGroup.Done()
		}
	}
}