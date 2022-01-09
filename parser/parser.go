package parser

import (
	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	_ "github.com/pingcap/tidb/types/parser_driver"
)

func ParseNode(sql string) (*ast.StmtNode, error) {
	p := parser.New()
	stmsNode, err := p.ParseOneStmt(sql, "", "")
	if err != nil {
		return nil, err
	}
	return &stmsNode, nil
}
