package ast

import (
	"fmt"
	"github.com/sealbro/go-gen-root/internal/injection"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func ParseGoProject(projectRoot string) (map[injection.StructInfo]*injection.InjectObject, error) {
	files, err := getAllGoFiles(projectRoot)
	if err != nil {
		panic(err)
	}
	astFiles, err := getAstFiles(files)
	if err != nil {
		return nil, err
	}
	return getInjectObjects(astFiles), nil
}

func getInjectObjects(astFiles map[string]*ast.File) map[injection.StructInfo]*injection.InjectObject {
	injectObjects := make(map[injection.StructInfo]*injection.InjectObject)
	for filePath, f := range astFiles {
		ast.Inspect(f, func(n ast.Node) bool {
			switch x := n.(type) {
			case *ast.FuncDecl:
				if x.Recv != nil || x.Doc == nil {
					return false
				}
				if !slices.ContainsFunc(x.Doc.List, func(comment *ast.Comment) bool {
					return strings.Contains(comment.Text, "@inject")
				}) {
					return false
				}

				deps := make([]injection.StructInfo, 0)
				for _, field := range x.Type.Params.List {
					switch p := field.Type.(type) {
					case *ast.StarExpr:
						switch s := p.X.(type) {
						case *ast.SelectorExpr:
							switch ident := s.X.(type) {
							case *ast.Ident:
								deps = append(deps, injection.StructInfo{
									PackageName: ident.Name,
									StructName:  s.Sel.Name,
								})
							}
						}
					}
				}

				var structName string
				for _, field := range x.Type.Results.List {
					switch p := field.Type.(type) {
					case *ast.StarExpr:
						switch s := p.X.(type) {
						case *ast.Ident:
							structName = s.Name
							break
						}
					}
				}

				obj := &injection.InjectObject{
					StructInfo: injection.StructInfo{
						PackageName: f.Name.Name,
						StructName:  structName,
					},
					FilePath: filePath,
					FuncDecl: x,
					AstFile:  f,
					Deps:     deps,
				}
				injectObjects[obj.StructInfo] = obj
			}
			return true
		})
	}

	return injectObjects
}

func getAstFiles(files []string) (map[string]*ast.File, error) {
	astFiles := make(map[string]*ast.File, len(files))

	fileSet := token.NewFileSet()
	for _, filename := range files {
		src, err := os.ReadFile(filename)
		if err != nil {
			return nil, fmt.Errorf("error reading file %s: %w", filename, err)
		}
		f, err := parser.ParseFile(fileSet, filename, src, parser.ParseComments)
		if err != nil {
			return nil, fmt.Errorf("error parsing file %s: %w", filename, err)
		}
		astFiles[filename] = f
	}

	return astFiles, nil
}

func getAllGoFiles(appDir string) ([]string, error) {
	var files []string
	err := filepath.Walk(appDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".go" {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
