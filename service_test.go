package metric

import "fmt"

type simpleProvider struct {
	data []*Item
}

func (sp *simpleProvider) ReadData(ch chan *Item) {
	var items []*Item
	for i := 0; i < 100; i++ {
		item := &Item{
			Log: Log{
				Request: fmt.Sprintf("fake request %d", i),
			},
			Error: nil,
		}
		ch <- item
		items = append(items, item)
	}
}
