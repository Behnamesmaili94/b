package main

import (
	. "log"
)

type school struct {
	math     *float32
	biology  *float32
	chimical *float32
	physics  *float32
	english  *float32
	history  *float32
	m        *int
}

func main() {

	Println("Main is running")

	alex := school{
		math:     new(float32),
		biology:  new(float32),
		chimical: new(float32),
		physics:  new(float32),
		english:  new(float32),
		history:  new(float32),
		m:        new(int),
	}

	sam := school{
		math:     new(float32),
		biology:  new(float32),
		chimical: new(float32),
		physics:  new(float32),
		english:  new(float32),
		history:  new(float32),
		m:        new(int),
	}

	if alex.averag() > sam.averag() {
		Println("alex has Mvp", sam.averag())
	} else {
		Println("sam has Mvp")
	}

	alex.change(18)
	sam.change(19)

}

func (set school) setValue(biology float32, math float32, physics float32, chimical float32, history float32, english float32) {
	*set.biology = biology
	*set.math = math
	*set.physics = physics
	*set.chimical = chimical
	*set.history = history
	*set.english = english
}

func (get school) averag() float32 {
	result := (*get.biology + *get.chimical + *get.english + *get.history + *get.math + *get.physics) / 6
	return result
}

func (s school) change(arg1 int) {
	*s.m = arg1
	Println("change running m : ", *s.m)

}
