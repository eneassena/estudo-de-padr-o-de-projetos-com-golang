package main

import "fmt"

type Singleton struct {
	Count int
}

var instance *Singleton

func NewSingleton() *Singleton {
	if instance == nil {
		instance = new(Singleton)
	}
	return instance
}

func (s *Singleton) AddOne() int {
	s.Count++
	return s.Count
}

func main() {
	inst1 := NewSingleton()

	fmt.Println(inst1.AddOne())
	fmt.Println(inst1.AddOne())
	fmt.Println(inst1.AddOne())
	fmt.Println(inst1.AddOne())
	fmt.Println()
}
