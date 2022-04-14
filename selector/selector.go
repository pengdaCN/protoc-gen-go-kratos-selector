package selector

import "github.com/go-kratos/kratos/v2/middleware"

type Selector interface {
	Global() []string
	GetOperations(name string) []string
	GetHandlers(name string) []middleware.Middleware
	Build() middleware.Middleware
}

func Chain(s ...Selector) middleware.Middleware {
	return nil
}
