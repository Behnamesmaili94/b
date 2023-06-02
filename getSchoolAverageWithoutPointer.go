package main

import (
	. "log"
)

type school struct {
	math     float64
	biology  float64
	chimical float64
	physics  float64
	english  float64
	history  float64
	m        int
}

func main() {

	Println("Main is running")

	var n1 school

	n1.setValue(20.0, 20.0, 20.0, 20.0, 20.0, 18.5)

	alex := school{
		math:     20,
		biology:  19,
		chimical: 18,
		physics:  19,
		english:  20,
		history:  17,
		m:        16,
	}

	sam := school{
		math:     19,
		biology:  18,
		chimical: 19,
		physics:  20,
		english:  20,
		history:  20,
		m:        17,
	}

	if alex.averag() > sam.averag() {
		Println("alex has Mvp", sam.averag())
	} else {
		Println("sam has Mvp")
	}

	alex.change(18)
	sam.change(19)
	println(sam.m, alex.m)
	println("number one is : ", int(n1.math), int(n1.biology), int(n1.english), int(n1.history), n1.m)

}

func (set *school) setValue(biology float64, math float64, physics float64, chimical float64, history float64, english float64) {
	set.biology = biology
	set.math = math
	set.physics = physics
	set.chimical = chimical
	set.history = history
	set.english = english
}

func (get school) averag() float64 {
	result := (get.biology + get.chimical + get.english + get.history + get.math + get.physics) / 6
	return result
}

func (s *school) change(arg1 int) {
	s.m = arg1
}
