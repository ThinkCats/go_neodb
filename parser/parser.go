package parser

import (
	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	_ "github.com/pingcap/tidb/types/parser_driver"
)

func ParseNode(sql string) (*ast.StmtNode, error) {
	p := parser.New()
	stmsNode, _, err := p.Parse(sql, "", "")
	if err != nil {
		return nil, err
	}
	return &stmsNode[0], nil
}
