package metric

import "fmt"

// DataProvider provide the data
type DataProvider interface {
	// Read the data from data provider
	// data provider send the data to the channel one by one
	ReadData() (<-chan Item, error)
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
	ch, err := ser.provider.ReadData()
	if err != nil {
		return err
	}

	for item := range ch {
		// TODO: 处理log
		err := ser.repo.Insert(&item)
		if err != nil {
			// TODO: gather error
		}
	}
	return nil
}

// Read the data from data provider
func (ser *service) Read() ([]Item, error) {
	ch, err := ser.provider.ReadData()
	if err != nil {
		return nil, fmt.Errorf("read data failure. %v", err)
	}

	var items []Item
	for item := range ch {
		items = append(items, item)
	}
	return items, nil
}
