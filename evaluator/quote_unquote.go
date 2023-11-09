package evaluator

import (
	"github.com/theakhandpatel/interpreter/ast"
	"github.com/theakhandpatel/interpreter/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{
		Node: node,
	}
}
