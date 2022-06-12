package gnet

import (
	"fmt"
	"net"

	"github.com/xiaobinqt/gmax/giface"
)

/**
连接模块
*/
type Connection struct {
	// 当前连接的socket TCPConn
	Conn *net.TCPConn

	// 当前连接的ID
	ConnID uint32

	// 当前连接的状态
	isClosed bool

	// 当前连接所绑定的处理业务的方法
	handleAPI giface.HandleFunc

	// 告知当前链接移除退出 channel
	ExitChan chan bool
}

// 初始化连接模块
func NewConnection(conn *net.TCPConn, connID uint32, callbackAPI giface.HandleFunc) *Connection {
	return &Connection{
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		handleAPI: callbackAPI,
		ExitChan:  make(chan bool, 1),
	}
}

func (c *Connection) Start() {
	fmt.Println("Start connection id = ", c.ConnID)
	// 启动从当前连接读取数据
	go c.StartRead()
}

// 连接的读业务
func (c *Connection) StartRead() {
	fmt.Println("conn StartRead() ...")
	defer fmt.Println("conn StartRead() exit...", c.GetRemoteAddr().String())
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("conn read error:", err)
			continue
		}

		// 将读取到的数据交给业务处理
		err = c.handleAPI(c.Conn, buf[:cnt], cnt)
		if err != nil {
			fmt.Println("conn handleAPI error:", err)
			break
		}
	}
}

func (c *Connection) Stop() {
	fmt.Println("conn stop() ...", c.ConnID)
	if c.isClosed {
		return
	}

	c.isClosed = true
	// 关闭 socket 连接
	c.Conn.Close()

	// 关闭 channel,回收资源
	close(c.ExitChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	return nil
}
