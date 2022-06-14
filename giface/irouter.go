package giface

/**
路由接口

路由里的数据都是 IRequest

*/

type IRouter interface {
	// 在处理业务之前的操作
	PreHandle(request IRequest)
	// 处理业务的主方法
	Handle(request IRequest)
	// 在处理业务之后的操作
	PostHandle(request IRequest)
}
