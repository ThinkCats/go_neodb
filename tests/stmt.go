package tests

import (
	"github.com/pingcap/tidb/util/parser"
	//	"github.com/pingcap/tidb/util/parser/ast"
)

func ParseCreateTable(sql string) {
	p := parser.New()
	stmtNode, _ := p.ParseOneStmt(sql, "", "")
}
