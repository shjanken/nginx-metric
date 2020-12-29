package main

import (
	"strings"

	metric "com.github.shjanken/nginx-metric"
	"com.github.shjanken/nginx-metric/test"
	"github.com/satyrius/gonx"
)

type gnoxProvider struct {
}

func (p *gnoxProvider) ReadData() (<-chan metric.Item, error) {
	ch := make(chan metric.Item)
	logs := strings.NewReader(test.TestLogs)
	config := strings.NewReader(test.TestConfig)

	_, err := gonx.NewNginxReader(logs, config, "main")
	if err != nil {
		return nil, err
	}

	return ch, nil
}
