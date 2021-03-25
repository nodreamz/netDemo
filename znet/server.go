package znet

import (
	"fmt"
	"net"
	"zinx/utils"
	"zinx/ziface"
)

//IServer的接口实现，定义一个Server的服务器模块
type Server struct {
	//服务器的名称
	Name string
	//服务器绑定的ip版本
	IPVersion string
	//服务器监听的IP
	IP string
	//服务器监听的端口
	Port int
	//当前的Server添加一个router，server注册的链接对应的处理业务
	Router ziface.IRouter
}


func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("Add Router success")
}

//启动服务器
func (s *Server) Start() {
	//获取一个TCP的Addr
	fmt.Printf("[Zinx] Server Name：%s, listener at IP: %s, Port: %d is starting\n", utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	fmt.Printf("[Zinx] Version: %s, MaxConn: %d, MaxPackeetSize: %d\n", utils.GlobalObject.Version, utils.GlobalObject.MaxConn, utils.GlobalObject.MaxPackageSize)
	fmt.Printf("[start] Serve Listener at IP: %s, Port:%d is starting\n", s.IP, s.Port)

	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolveTcpAddr error", err)
			return
		}
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("net.Listen error", err)
			return
		}

		fmt.Println("start Zinx server succ,", s.Name, "succ, Listening...")
		var cid uint32
		cid = 0
		//监听服务器的地址
		//阻塞的等待客户端链接，处理客户端链接业务（读写）
		for {
			//如果有客户端链接过来，阻塞会返回
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("accept error", err)
				continue
			}

			//已经与客户端建立链接，做一些业务
			//将处理新连接的业务方法和conn惊醒绑定得到我们的链接模块
			dealConn := NewConnection(conn, cid, s.Router)
			cid++

			//启动当前的链接业务处理
			go dealConn.Start()
		}
	}()


}

//停止服务器
func (s *Server) Stop() {
	//TODO 将一些服务器的资源、状态或者一些已经开辟的链接信息 进行停止或者回收
}

//运行服务器
func (s *Server) Serve() {
	//启动server的服务功能
	s.Start()
	//TODO 做一些启动服务器之后的额外业务
	select{}

}


//初始化Server模块的方法
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name: utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP: utils.GlobalObject.Host,
		Port: utils.GlobalObject.TcpPort,
		Router: nil,

	}
	return s
}
