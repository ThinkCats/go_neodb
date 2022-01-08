package parser

import (
	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
)

func ParseSQL(sql string) (*ast.StmtNode, error) {
	p := parser.New()
	stmsNode, err := p.Parse(sql, "", "")
	if err != nil {
		return nil, err
	}
	return &stmsNode, nil
}
