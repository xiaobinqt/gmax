package gnet

import (
	"fmt"
	"net"

	"github.com/xiaobinqt/gmax/giface"
)

// IServer 是一个服务器接口实现
type Server struct {
	// 服务器的名称
	Name string
	// 服务器绑定的 ip 版本
	IPVersion string
	// 服务器监听的 ip
	IP string
	// 服务器监听的端口
	Port int
	// 当前 server 添加一个 router,server 注册的链接对应的处理业务
	Router giface.IRouter
}

// 初始化 Server
func NewServer(name string) giface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
		Router:    nil,
	}

	return s
}

func (s *Server) AddRouter(router giface.IRouter) {
	s.Router = router
	fmt.Println("Server Add Router Success!")
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listenner, ServerName: %s, IP: %s, Port: %d\n", s.Name, s.IP, s.Port)

	go func() {
		// 1. 获取一个 TCP addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("ResolveIPAddr error: ", err.Error())
			return
		}
		// 2. 监听服务器的地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listenTCP error: ", err.Error())
			return
		}

		fmt.Println("Start Server Success!", s.Name, s.IPVersion, s.IP, s.Port)

		var cid uint32
		cid = 0
		// 3. 阻塞等待客户端连接,处理客户端业务
		for {
			// 如果有客户端连接过来,阻塞返回
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept error: ", err.Error())
				continue
			}

			// 客户端已经与服务器建立连接,处理业务
			dealConn := NewConnection(conn, cid, s.Router)
			cid++

			// 启动当前连接的处理业务
			go dealConn.Start()
		}
	}()

}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	// 启动 server 的服务功能
	s.Start()

	// TODO 做一些启动服务器之后的额外功能

	// 阻塞下...
	select {}
}
