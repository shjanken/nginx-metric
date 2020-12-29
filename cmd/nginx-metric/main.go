package main

import (
	metric "com.github.shjanken/nginx-metric"
)

func main() {
	p := gnoxProvider{}
	service := metric.NewService(&p)
	service.Save()
}
