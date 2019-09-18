package demo

import "github.com/itcast/zinx/net"

func main()  {
	server:=net.NewServer()
	server.Start()
}