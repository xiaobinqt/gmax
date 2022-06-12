package giface

import "net"

type IConnection interface {
	// 启动连接，让当前连接开始工作
	Start()
	//停止连接，结束当前连接工作
	Stop()
	// 获取当前连接绑定的socket conn
	GetTCPConnection() *net.TCPConn
	// 获取当前连接模块的ID
	GetConnID() uint32
	// 获取远程客户端的TCP状态 IP Port
	GetRemoteAddr() net.Addr
	// 发送数据给远程客户端
	Send(data []byte) error
}

// 定义一个处理连接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
