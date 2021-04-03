package functions

import (
	"github.com/fdaines/arch-go/utils/packages"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func retrieveFunctions(pkg *packages.PackageInfo, mainPackage string) ([]*FunctionDetails, error) {
	functionDetailsCollection := []*FunctionDetails{}
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	packageDir := strings.Replace(pkg.PackageData.ImportPath, mainPackage, path, 1)

	for _, srcFile := range pkg.PackageData.GoFiles {
		srcFilePath := filepath.Join(mainPackage, srcFile)
		data, err := ioutil.ReadFile(filepath.Join(packageDir, srcFile))
		if err != nil {
			return nil, err
		}
		fileset := token.NewFileSet()
		node, err := parser.ParseFile(fileset, srcFile, data, 0)
		if err != nil {
			return nil, err
		}
		ast.Inspect(node, func(n ast.Node) bool {
			switch t := n.(type) {
			case *ast.FuncDecl:
				functionDetails := &FunctionDetails{
					File:     srcFile,
					FilePath: srcFilePath,
					Name:     t.Name.Name,
					IsPublic: unicode.IsUpper([]rune(t.Name.Name)[0]),
				}
				if t.Type.Params != nil {
					functionDetails.NumParams = len(t.Type.Params.List)
				}
				if t.Type.Results != nil {
					functionDetails.NumReturns = len(t.Type.Results.List)
				}
				functionDetailsCollection = append(functionDetailsCollection, functionDetails)
			}
			return true
		})
	}
	return functionDetailsCollection, nil
}
