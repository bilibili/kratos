package main

import (
	"fmt"
	"strings"

	"github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2/errors"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

const (
	errorsPackage = protogen.GoImportPath("github.com/go-kratos/kratos/v2/errors")
	fmtPackage    = protogen.GoImportPath("fmt")
)

// generateFile generates a _errors.pb.go file containing kratos errors definitions.
func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	if len(file.Enums) == 0 {
		return nil
	}
	filename := file.GeneratedFilenamePrefix + "_errors.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-errors. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	g.QualifiedGoIdent(fmtPackage.Ident(""))
	generateFileContent(gen, file, g)
	return g
}

// generateFileContent generates the kratos errors definitions, excluding the package statement.
func generateFileContent(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile) {
	if len(file.Enums) == 0 {
		return
	}

	g.P("// This is a compile-time assertion to ensure that this generated file")
	g.P("// is compatible with the kratos package it is being compiled against.")
	g.P("const _ = ", errorsPackage.Ident("SupportPackageIsVersion1"))
	g.P()
	index := 0
	for _, enum := range file.Enums {
		skip := genErrorsReason(gen, file, g, enum)
		if !skip {
			index++
		}
	}
	// If all enums do not contain 'errors.code', the current file is skipped
	if index == 0 {
		g.Skip()
	}
}

func genErrorsReason(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, enum *protogen.Enum) bool {
	defaultCode := proto.GetExtension(enum.Desc.Options(), errors.E_DefaultCode)
	code := 0
	if ok := defaultCode.(int32); ok != 0 {
		code = int(ok)
	}
	// if code > 600 || code < 0 {
	// 	panic(fmt.Sprintf("Enum '%s' range must be greater than 0 and less than or equal to 600", string(enum.Desc.Name())))
	// }
	var ew errorWrapper
	for _, v := range enum.Values {
		message := ""
		msg := proto.GetExtension(v.Desc.Options(), errors.E_Msg)
		if ok := msg.(string); ok != "" {
			message = string(ok)
		}
		enumCode := code
		eCode := proto.GetExtension(v.Desc.Options(), errors.E_Code)
		if ok := eCode.(int32); ok != 0 {
			enumCode = int(ok)
		}
		// If the current enumeration does not contain 'errors.code'
		// or the code value exceeds the range, the current enum will be skipped
		if enumCode > 600 || enumCode < 0 {
			panic(fmt.Sprintf("Enum '%s' range must be greater than 0 and less than or equal to 600", string(v.Desc.Name())))
		}
		if enumCode == 0 {
			continue
		}
		err := &errorInfo{
			Name:       string(enum.Desc.Name()),
			Message:    message,
			Value:      string(v.Desc.Name()),
			CamelValue: Case2Camel(string(v.Desc.Name())),
			HttpCode:   enumCode,
		}
		ew.Errors = append(ew.Errors, err)
	}
	if len(ew.Errors) == 0 {
		return true
	}
	g.P(ew.execute())
	return false
}

func Case2Camel(name string) string {
	if !strings.Contains(name, "_") {
		return name
	}
	name = strings.ToLower(name)
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}
