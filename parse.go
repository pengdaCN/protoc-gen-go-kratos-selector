package main

import (
	"fmt"
	"gitee.com/pengdacn/protoc-gen-go-kratos-selector/selector"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

const (
	suffix               = "_selector.pb.go"
	contextPackage       = protogen.GoImportPath("context")
	transportHTTPPackage = protogen.GoImportPath("github.com/go-kratos/kratos/v2/transport/http")
	bindingPackage       = protogen.GoImportPath("github.com/go-kratos/kratos/v2/transport/http/binding")
)

// generateFile generates a _selector.pb.go file
func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	if len(file.Services) == 0 {
		return nil
	}
	filename := file.GeneratedFilenamePrefix + suffix
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by proto-gen-go-kratos-selector. DO NOT EDIT.")
	g.P("// versions:")
	g.P(fmt.Sprintf("// proto-gen-go-kratos-selector %s", release))
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	generateFileContent(gen, file, g)
	return g
}

func generateFileContent(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile) {
	if len(file.Services) == 0 {
		return
	}
	g.P("// This is a compile-time assertion to ensure that this generated file")
	g.P("// is compatible with the kratos package it is being compiled against.")
	g.P("var _ = new(", contextPackage.Ident("Context"), ")")
	g.P()

	for _, service := range file.Services {
		genService(gen, file, g, service)
	}
}

func genService(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service) {
	if service.Desc.Options().(*descriptorpb.ServiceOptions).GetDeprecated() {
		g.P("//")
		g.P(deprecationComment)
	}

	handlers, ok := proto.GetExtension(service.Desc.Options(), selector.E_Handlers).(*selector.Handlers)
	if !ok {
		return
	}

	sd := &serviceDesc{
		ServiceType: service.GoName,
		ServiceName: string(service.Desc.FullName()),
		Metadata:    file.Desc.Path(),
		TagSet:      make(map[string]*tag),
		Maps:        make(map[string][]string),
	}

	for _, h := range handlers.Handlers {
		sd.Handlers = append(sd.Handlers, &handler{
			Name:     h.Name,
			Select:   h.Select,
			Property: h.Property,
			Enabled:  h.Enabled,
			Global:   h.Global,
		})
	}

	// 创建tag
	for _, method := range service.Methods {
		if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
			continue
		}
		grp, ok := proto.GetExtension(method.Desc.Options(), selector.E_Group).(*selector.Group)
		if grp != nil && ok {
			t := sd.touchTag(grp.Name)

			t.addMethod(string(method.Desc.Name()))

			for _, g := range grp.Additional {
				t := sd.touchTag(g)

				t.addMethod(string(method.Desc.Name()))
			}
		} else {
			sd.touchTag("").addMethod(string(method.Desc.Name()))
		}
	}

	g.P(sd.execute())
}

const deprecationComment = "// Deprecated: Do not use."
