package engine

import "fmt"

type Engine struct {
	Scaffold Scaffold
}

func NewEngine(scaffold Scaffold) *Engine {
	return &Engine{Scaffold: scaffold}
}

func (e *Engine) Run() {
	fmt.Println("Running engine")
}
