package main

import (
	_ "embed"
	"fmt"
	"strings"
	"text/template"
)

//go:embed tmpl.go.tmpl
var tmplStr string
var tmpl *template.Template

func init() {
	tmpl = template.Must(template.New("selector").Parse(tmplStr))
}

type serviceDesc struct {
	ServiceType string // Greeter
	ServiceName string // helloworld.Greeter
	Metadata    string // api/helloworld/helloworld.proto
	Handlers    []*handler
	Tags        []*tag
	TagSet      map[string]*tag
	Maps        map[string][]string
	Methods     map[string][]string
}

type handler struct {
	Name     string
	Select   string
	Property int32
	Enabled  bool
	Global   bool
}

type tag struct {
	Name      string
	methodSet map[string]struct{}
}

func (s *serviceDesc) touchTag(name string) *tag {
	t, ok := s.TagSet[name]
	if !ok {
		t = &tag{
			Name:      name,
			methodSet: make(map[string]struct{}),
		}

		s.TagSet[name] = t
		s.Tags = append(s.Tags, t)
	}

	return t
}

func (t *tag) addMethod(method string) {
	t.methodSet[method] = struct{}{}
}

func (s *serviceDesc) execute() string {
	for _, h := range s.Handlers {
		var c []string

		methodSet := make(map[string]struct{})
		for _, t := range s.Tags {
			if match(h.Select, t.Name) {
			skip:
				for m := range t.methodSet {
					for _, tag := range s.Methods[m] {
						if !match(h.Select, tag) {
							continue skip
						}
					}

					methodSet[m] = struct{}{}
				}
			}
		}

		for m := range methodSet {
			c = append(c, fmt.Sprintf("/%s/%s", s.ServiceName, m))
		}

		s.Maps[h.Name] = c
	}

	var buf strings.Builder
	if err := tmpl.Execute(&buf, s); err != nil {
		panic(err)
	}
	return strings.Trim(buf.String(), "\r\n")
}

func match(selector, name string) bool {
	verbs := strings.Split(selector, "|")
	for _, verb := range verbs {
		verb = strings.TrimSpace(verb)
		switch {
		case verb == "$any":
			return true
		case verb[:1] == "!":
			if verb[1:] != name {
				return true
			} else {
				return false
			}
		default:
			if verb == name {
				return true
			}
		}
	}

	return false
}
