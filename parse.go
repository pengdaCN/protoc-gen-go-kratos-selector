package main

import (
	"fmt"
	"gitee.com/pengdacn/protoc-gen-go-kratos-selector/selector"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

type serviceDesc struct {
	selectInfo    *selector.Selector
	name          string
	protoName     string
	protoFullName string
	funcs         []funcDesc
}

type funcDesc struct {
	name          string
	protoName     string
	protoFullName string
	tags          []string
}

func parse(file *protogen.File) ([]serviceDesc, error) {
	if len(file.Services) == 0 {
		return nil, nil
	}

	var r []serviceDesc
	for _, service := range file.Services {
		descs, err := parseService(service)
		if err != nil {
			return nil, err
		}

		if descs.selectInfo == nil {
			continue
		}

		r = append(r, descs)
	}

	return r, nil
}

func parseService(srv *protogen.Service) (descSrv serviceDesc, err error) {
	if len(srv.Methods) == 0 {
		return
	}

	extension, ok := proto.GetExtension(srv.Desc.Options(), selector.E_Select).(*selector.Selector)
	if extension == nil || !ok {
		return
	}

	if extension.Use == "" {
		return descSrv, fmt.Errorf("%s use is nil", srv.Desc.FullName())
	}

	useageIds := make(map[string]struct{})
	for _, verb := range extension.Verbs {
		_, ok := useageIds[verb.Id]
		if ok {
			return descSrv, fmt.Errorf("%s repeat id:%s", srv.Desc.FullName(), verb.Id)
		}

		useageIds[verb.Id] = struct{}{}
	}

	descSrv.selectInfo = extension
	descSrv.name = srv.GoName
	descSrv.protoName = string(srv.Desc.Name())
	descSrv.protoFullName = string(srv.Desc.FullName())
	for _, method := range srv.Methods {
		desc, err := parseMethod(method)
		if err != nil {
			return descSrv, err
		}

		descSrv.funcs = append(descSrv.funcs, desc)
	}

	return
}

func parseMethod(method *protogen.Method) (desc funcDesc, err error) {
	desc.name = method.GoName
	desc.protoName = string(method.Desc.Name())
	desc.protoFullName = string(method.Desc.FullName())

	extension, ok := proto.GetExtension(method.Desc.Options(), selector.E_Tag).(*selector.Tag)
	if extension == nil || !ok {
		return
	}

	tags := map[string]struct{}{
		desc.name: {},
	}

	desc.tags = append(desc.tags, extension.Name)
	for _, name := range extension.Additional {
		_, ok := tags[name]
		if ok {
			return desc, fmt.Errorf("%s repeat tag name: %s", method.Desc.FullName(), name)
		}

		desc.tags = append(desc.tags, name)
		tags[name] = struct{}{}
	}

	return
}
