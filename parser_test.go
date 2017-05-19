package pbparser_test

import (
	"fmt"
	"testing"

	pbparser "github.com/tallstoat/pbparser"
)

func TestParseFile(t *testing.T) {
	var tests = []struct {
		file string
	}{
		{file: "./resources/service.proto"},
	}

	for i, tt := range tests {
		fmt.Printf("Running test: %v \n", i)
		fmt.Printf("Parsing file: %v \n", tt.file)

		pf, err := pbparser.ParseFile(tt.file)
		if err != nil {
			t.Errorf("%v", err.Error())
		}

		fmt.Println("Syntax: " + pf.Syntax)
		fmt.Println("PackageName: " + pf.PackageName)

		for _, d := range pf.Dependencies {
			fmt.Println()
			fmt.Println("Dependency: " + d)
		}

		for _, d := range pf.PublicDependencies {
			fmt.Println()
			fmt.Println("PublicDependency: " + d)
		}

		for _, m := range pf.Messages {
			fmt.Println()
			fmt.Println("Message: " + m.Name)
			fmt.Println("QualifiedName: " + m.QualifiedName)
			fmt.Println("Doc: " + m.Documentation)
			for _, f := range m.Fields {
				fmt.Println()
				fmt.Println("Field: " + f.Name)
				if f.Label != "" {
					fmt.Println("Label: " + f.Label)
				}
				fmt.Printf("Type: %v\n", f.Type)
				fmt.Printf("Tag: %v\n", f.Tag)
				fmt.Println("Doc: " + f.Documentation)
			}
		}

		for _, ed := range pf.ExtendDeclarations {
			fmt.Println()
			fmt.Println("Extend: " + ed.Name)
			fmt.Println("QualifiedName: " + ed.QualifiedName)
			fmt.Println("Doc: " + ed.Documentation)
			for _, f := range ed.Fields {
				fmt.Println()
				fmt.Println("Field: " + f.Name)
				if f.Label != "" {
					fmt.Println("Label: " + f.Label)
				}
				fmt.Printf("Type: %v\n", f.Type)
				fmt.Printf("Tag: %v\n", f.Tag)
				fmt.Println("Doc: " + f.Documentation)
			}
		}

		for _, s := range pf.Services {
			fmt.Println()
			fmt.Println("Service: " + s.Name)
			fmt.Println("QualifiedName: " + s.QualifiedName)
			fmt.Println("Doc: " + s.Documentation)
			for _, rpc := range s.RPCs {
				fmt.Println()
				fmt.Println("RPC: " + rpc.Name)
				fmt.Println("Doc: " + rpc.Documentation)
				fmt.Println("RequestType: " + rpc.RequestType.Name())
				fmt.Println("ResponseType: " + rpc.ResponseType.Name())
			}
		}

		for _, en := range pf.Enums {
			fmt.Println()
			fmt.Println("Enum: " + en.Name)
			fmt.Println("QualifiedName: " + en.QualifiedName)
			fmt.Println("Doc: " + en.Documentation)
			for _, enc := range en.EnumConstants {
				fmt.Println("Name: " + enc.Name)
				fmt.Printf("Tag: %v \n", enc.Tag)
			}
		}

		fmt.Printf("\nFinished test: %v \n", i)
	}

	fmt.Println("done")
}
