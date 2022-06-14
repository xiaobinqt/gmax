package giface

// 把客户端请求链接信息封装到接口中

type IRequest interface {
	// 得到当前链接
	GetConnection() IConnection
	// 得到请求消息
	GetData() []byte
}
