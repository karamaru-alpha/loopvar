package loopvar

import (
	"fmt"
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "loopvar",
	Doc:  "loopvar is a linter that detects places where loop variables are copied.",
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.RangeStmt)(nil),
		(*ast.ForStmt)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch node := n.(type) {
		case *ast.RangeStmt:
			checkRangeStmt(pass, node)
		case *ast.ForStmt:
			checkForStmt(pass, node)
		}
	})

	return nil, nil
}

func checkRangeStmt(pass *analysis.Pass, rangeStmt *ast.RangeStmt) {
	key, ok := rangeStmt.Key.(*ast.Ident)
	if !ok {
		return
	}
	var value *ast.Ident
	if rangeStmt.Value != nil {
		value = rangeStmt.Value.(*ast.Ident)
	}
	for _, stmt := range rangeStmt.Body.List {
		assignStmt, ok := stmt.(*ast.AssignStmt)
		if !ok {
			continue
		}
		if assignStmt.Tok != token.DEFINE {
			continue
		}
		for _, rh := range assignStmt.Rhs {
			right, ok := rh.(*ast.Ident)
			if !ok {
				continue
			}
			if right.Name != key.Name && (value != nil && right.Name != value.Name) {
				continue
			}
			pass.Report(analysis.Diagnostic{
				Pos:     assignStmt.Pos(),
				Message: fmt.Sprintf(`The loop variable "%s" should not be copied (Go 1.22~ or Go 1.21 GOEXPERIMENT=loopvar)`, right.Name),
				SuggestedFixes: []analysis.SuggestedFix{{
					TextEdits: []analysis.TextEdit{{
						Pos:     assignStmt.Pos(),
						End:     assignStmt.End(),
						NewText: []byte(""),
					}}},
				},
			})
		}
	}
}

func checkForStmt(pass *analysis.Pass, forStmt *ast.ForStmt) {
	if forStmt.Init == nil {
		return
	}
	initAssignStmt, ok := forStmt.Init.(*ast.AssignStmt)
	if !ok {
		return
	}
	initVarNameMap := make(map[string]interface{}, len(initAssignStmt.Lhs))
	for _, lh := range initAssignStmt.Lhs {
		if initVar, ok := lh.(*ast.Ident); ok {
			initVarNameMap[initVar.Name] = struct{}{}
		}
	}
	for _, stmt := range forStmt.Body.List {
		assignStmt, ok := stmt.(*ast.AssignStmt)
		if !ok {
			continue
		}
		if assignStmt.Tok != token.DEFINE {
			continue
		}
		for _, rh := range assignStmt.Rhs {
			right, ok := rh.(*ast.Ident)
			if !ok {
				continue
			}
			if _, ok := initVarNameMap[right.Name]; !ok {
				continue
			}
			pass.Report(analysis.Diagnostic{
				Pos:     assignStmt.Pos(),
				Message: fmt.Sprintf(`The loop variable "%s" should not be copied (Go 1.22~ or Go 1.21 GOEXPERIMENT=loopvar)`, right.Name),
				SuggestedFixes: []analysis.SuggestedFix{{
					TextEdits: []analysis.TextEdit{{
						Pos:     assignStmt.Pos(),
						End:     assignStmt.End(),
						NewText: []byte(""),
					}}},
				},
			})
		}
	}
}
