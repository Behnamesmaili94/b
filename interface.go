package main

import "fmt"

type car struct {
	name  string
	price int
	id    int
}

type airplan struct {
	name  string
	price int
	id    int
}

type SetName interface {
	setname(string2 string)
}

type SetPrice interface {
	setPrice(int string)
}

type SetId interface {
	setId(int string)
}

func main() {

	var exmaple SetName

	airplan1 := airplan{name: "super hornet"}
	exmaple = &airplan1
	exmaple.setname("raptor")

	car1 := car{name: "pride"}
	exmaple = &car1
	exmaple.setname("206")

	fmt.Println("name is : ", car1.name)
}

func (c *car) setname(newName string) {
	c.name = newName
	fmt.Println(c.name + " is model")
}

func (c *airplan) setname(newName string) {
	c.name = newName
	fmt.Println(c.name + " is airplan name")
}

func (c *car) setPrice(newPrice int) {
	c.price = newPrice
	fmt.Println(c.price, "car price now")
}

func (c *airplan) setPrice(newPrice int) {
	c.price = newPrice
	fmt.Println(c.price, "airplan price now")
}

func (c *car) setId(newId int) {
	c.id = newId
	fmt.Println(c.id, ":is id")
}

func (c *airplan) setId(newId int) {
	c.id = newId
	fmt.Println(c.id, ":is id")
}
