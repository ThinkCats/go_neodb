package parser

//
//import "github.com/pingcap/parser/ast"
//
//type SqlMode int
//
//const (
//	SELECT SqlMode = iota
//	INSERT
//	UPDATE
//	DELETE
//	CREATE_DB
//	CREATE_TABLE
//	ALTER
//)
//
//type SqlNode struct {
//	sqlMode    SqlMode
//	colNames   []string
//	tableNames []string
//}
//
//func (c *SqlNode) Enter(n ast.Node) (node ast.Node, skipChildren bool) {
//	if name, ok := n.(*ast.ColumnName); ok {
//		c.colNames = append(c.colNames, name.Name.O)
//	}
//	if name, ok := n.(*ast.TableName); ok {
//		c.tableNames = append(c.tableNames, name.Name.O)
//	}
//	return n, false
//}
//
//func (c *SqlNode) Leave(n ast.Node) (node ast.Node, ok bool) {
//	return n, true
//}
//
//func ParseColumn(rootNode *ast.StmtNode) []string {
//	v := &SqlNode{}
//	(*rootNode).Accept(v)
//	return v.colNames
//}
//
//func ParseTable(rootNode *ast.StmtNode) []string {
//	v := &SqlNode{}
//	(*rootNode).Accept(v)
//	return v.tableNames
//}
