package selector

import "github.com/go-kratos/kratos/v2/middleware"

type Selector interface {
	// Global 返回所有全局的处理器
	Global() []GlobalHandler
	// HalfBuild 只会构建非全局的中间件
	HalfBuild() middleware.Middleware
	Operations() []string
}

type GlobalHandler interface {
	Name() string
	Property() int32
	Handlers() []middleware.Middleware
	Match(string) bool
}

func Chain(s ...Selector) middleware.Middleware {
	return nil
}
