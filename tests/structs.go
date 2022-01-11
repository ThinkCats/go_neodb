package tests

import (
	"fmt"
)

type A interface {
	Move()
}

type Cat struct {
	name string
	feet int8
}

func (c *Cat) Move() {
	fmt.Println("Move")
}

//func (c *cat) move() {
//	fmt.Println("Move")
//}

func (c *Cat) Eat(food string) {
	fmt.Printf("eat = %+v\n", food)
}
