package parser

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {
	sql := "select a , b from xxx"
	result, err := ParseNode(sql)
	if err != nil {
		fmt.Println("Parse SQL ERROR", err)
	}
	fmt.Printf("result = %v\n", *result)
}

func TestExtra(t *testing.T) {
	createDb := "create database hello"
	r, _ := ParseNode(createDb)
	fmt.Printf("column result = %v\n", r)

	sql := "select a,b,c from xxx"
	result, _ := ParseNode(sql)
	fmt.Printf("column result = %v\n", result)

}
