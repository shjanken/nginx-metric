package metric

import (
	"fmt"
	"log"
)

// DataProvider provide the data
type DataProvider interface {
	// Read the data from data provider
	// data provider send the data to the channel one by one
	ReadData(ch chan *Item)
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
	Read() ([]Item, error)
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

	ch := make(chan *Item)
	go ser.provider.ReadData(ch)

	var logs []Log
	for item := range ch {

		// 从 channel 接受到数据以后
		// 将数据接受到一个数组里面，如果数组的长度到达 1000 了，则调用函数存放
		logs = append(logs, item.Log)
		if len(logs) == 1000 {
			fmt.Printf("has 1000 rows: %v\n", logs)
			if err := ser.repo.Insert(logs); err != nil {
				log.Fatalf("insert log failure. %v", err) // record error
			}
			logs = nil
		}

		// 如果数据少于1000条，直接插入
		if err := ser.repo.Insert(logs); err != nil {
			log.Fatalf("insert log failure. %v", err)
		}
	}
	return nil
}

// Read the data from data providr
func (ser *service) Read() ([]Item, error) {
	// 判断 service 结构体里是否有 provider. 如果没有则 panic
	if ser.provider == nil {
		panic(fmt.Sprintf("data provider is null. cant not read data"))
	}

	ch := make(chan *Item)
	go ser.provider.ReadData(ch)

	var items []Item
	for item := range ch {
		items = append(items, *item)
	}
	return items, nil
}
