package injection

import (
	"go/ast"
	"unicode"
	"unicode/utf8"
)

type StructInfo struct {
	PackageName string
	StructName  string
}

func (s StructInfo) CamelCaseStructName() string {
	return lowerFirst(s.StructName)
}

func (s StructInfo) String() string {
	return s.PackageName + "." + s.StructName
}

type InjectObject struct {
	StructInfo
	FilePath string
	Deps     []StructInfo
	FuncDecl *ast.FuncDecl
	AstFile  *ast.File
}

func (io InjectObject) FullName() string {
	return io.PackageName + "." + io.StructName
}

func (io InjectObject) FuncFullName() string {
	return io.PackageName + "." + io.FuncDecl.Name.Name
}

func lowerFirst(s string) string {
	if s == "" {
		return s
	}
	r, size := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[size:]
}
