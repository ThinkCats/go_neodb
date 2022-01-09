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
