package parser

import (
	"fmt"
	"log"

	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	_ "github.com/pingcap/tidb/types/parser_driver"
)

type SqlMod int

const (
	SELECT SqlMod = iota
	INSERT
	UPDATE
	DELETE
	CREATE_DB
	CREATE_TABLE
	ALTER
)

type Col struct {
	name string
}

type SqlNode struct {
	sqlMod     SqlMod
	colNames   []*Col
	tableNames []string
}

func (c *SqlNode) Enter(n ast.Node) (node ast.Node, skipChildren bool) {
	if createDb, ok := n.(*ast.CreateDatabaseStmt); ok {
		fmt.Println("Test Create DB AST:", createDb.Name)
		c.sqlMod = CREATE_DB
		c.tableNames = append(c.tableNames, createDb.Name)
		return n, false
	}
	if createTable, ok := n.(*ast.CreateTableStmt); ok {
		fmt.Println("Test Create Table AST:", createTable.Table)
		c.sqlMod = CREATE_TABLE
		c.tableNames = append(c.tableNames, createTable.Table.Text())
		return n, false
	}
	if name, ok := n.(*ast.ColumnName); ok {
		col := &Col{name.Name.O}
		c.colNames = append(c.colNames, col)
	}
	if name, ok := n.(*ast.TableName); ok {
		c.tableNames = append(c.tableNames, name.Name.O)
	}
	return n, false
}

func (c *SqlNode) Leave(n ast.Node) (node ast.Node, ok bool) {
	return n, true
}

func ParseNode(sql string) (*SqlNode, error) {
	p := parser.New()
	stmsNode, err := p.ParseOneStmt(sql, "", "")
	if err != nil {
		log.Fatalln("Parse SQL ERROR,", err)
		return nil, err
	}
	v := &SqlNode{}
	(*(&stmsNode)).Accept(v)
	return v, nil
}
