package metric

import (
	"fmt"
)

// DataProvider provide the data
type DataProvider interface {
	// Read the data from data provider
	// data provider send the data to the channel one by one
	ReadData(ch chan *Item) error
}

// Closer offer close function for repo
type Closer interface {
	// close the repo
	Close() error
}

// Inserter the data to backend
type Inserter interface {
	Insert(logs []Log) error
}

// Repo is backend
type Repo interface {
	Inserter
	Closer
}

// Service is metrics item service
type Service interface {
	Save() error
	Read() ([]*Item, error)
}

type service struct {
	provider DataProvider
	repo     Repo
}

// NewService return a new log service
func NewService(p DataProvider, r Repo) Service {
	return &service{
		provider: p,
		repo:     r,
	}
}

// Save the data read from the data provider
func (ser *service) Save() error {
	if ser.repo == nil {
		panic("the repo backend is nil")
	}
	if ser.provider == nil {
		panic("the data provider is nil")
	}
	defer ser.repo.Close()

	ch := make(chan *Item)
	go ser.provider.ReadData(ch)

	logs := make([]Log, 0, 1000)

	for {
		item, isOpen := <-ch
		// fmt.Println(item)
		if !isOpen && len(logs) != 0 {
			if err := ser.repo.Insert(logs); err != nil {
				return fmt.Errorf("service save data failure. %w", err)
			}
			break
		} else if len(logs) < 1000 {
			logs = append(logs, item.Log)
		} else if len(logs) == 1000 {
			logs = append(logs, item.Log)
			if err := ser.repo.Insert(logs); err != nil {
				return fmt.Errorf("service save data failure. %w", err)
			}
			logs = nil
		}
	}

	return nil
}

// Read the data from data providr
func (ser *service) Read() ([]*Item, error) {
	// 判断 service 结构体里是否有 provider. 如果没有则 panic
	if ser.provider == nil {
		panic(fmt.Sprintf("data provider is null. cant not read data"))
	}

	ch := make(chan *Item)
	go ser.provider.ReadData(ch)

	var items []*Item
	for item := range ch {
		items = append(items, item)
	}
	return items, nil
}
