package gnet

import "github.com/xiaobinqt/gmax/giface"

// 实现 router 先签入 BaseRouter
// 可以重写方法
type BaseRouter struct {
}

func (br *BaseRouter) PreHandle(request giface.IRequest) {
	// do nothing
}

func (br *BaseRouter) Handle(request giface.IRequest) {
	// do nothing
}

func (br *BaseRouter) PostHandle(request giface.IRequest) {
	// do nothing
}
