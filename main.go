package main

import (
	"fmt"
	"github.com/sealbro/go-gen-root/internal/ast"
	"github.com/sealbro/go-gen-root/internal/generator"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := cli.App{
		Name:  "go-root",
		Usage: "Composition root tool",
		Commands: []*cli.Command{
			{
				Name:    "generate",
				Usage:   "Generate composition root",
				Aliases: []string{"gen"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "path",
						Usage:    "path to the project root",
						Aliases:  []string{"p"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "module",
						Usage:    "name of the application's module",
						Aliases:  []string{"mod", "m"},
						Required: true,
					},
					&cli.StringFlag{
						Name:    "package",
						Usage:   "name of the package",
						Aliases: []string{"pkg"},
						Value:   "main",
					},
					&cli.StringFlag{
						Name:    "application",
						Usage:   "application struct name",
						Aliases: []string{"app", "a"},
						Value:   "App",
					},
					&cli.StringFlag{
						Name:    "output",
						Usage:   "composition root file to write the output to",
						Aliases: []string{"o"},
						Value:   "di.go",
					},
				},
				Action: func(c *cli.Context) error {
					params := generator.Params{
						AppName:     c.String("application"),
						ModuleName:  c.String("module"),
						PackageName: c.String("package"),
					}
					projectRoot := c.String("path")
					outputPath := c.String("output")

					return generate(params, projectRoot, outputPath)
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

func generate(params generator.Params, projectRoot, outputPath string) error {
	injectObjects, err := ast.ParseGoProject(projectRoot)
	if err != nil {
		return fmt.Errorf("failed to parse project: %w", err)
	}

	for _, obj := range injectObjects {
		for _, dep := range obj.Deps {
			if _, ok := injectObjects[dep]; !ok {
				return fmt.Errorf("dependency not found: %s", dep.String())
			}
		}
	}

	generatedCode, err := generator.Generate(params, injectObjects)
	if err != nil {
		return fmt.Errorf("failed to generate: %w", err)
	}

	err = os.WriteFile(outputPath, []byte(generatedCode), 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
