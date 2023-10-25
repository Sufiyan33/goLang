package main

type tank interface {
	Tarea() float64
	Volume() float64
}

type myValue struct {
	radius float64
	height float64
}

func (m myValue) Tarea() float64 {
	return 2*m.radius*m.height + 2*3.14*m.radius*m.radius
}

func (m myValue) Volume() float64 {
	return 3.14 * m.radius * m.radius * m.height
}
