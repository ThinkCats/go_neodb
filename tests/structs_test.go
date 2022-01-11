package tests

import (
	"log"
	"testing"
)

func TestMove(t *testing.T) {
	var a1 A
	c1 := &Cat{"tom", 4}
	c2 := Cat{"jery", 4}

	a1 = c1
	log.Println(a1)

	a1 = &c2
	log.Println(a1)

	a1.Move()
}
