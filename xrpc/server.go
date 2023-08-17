/*
功能：server
说明：
*/
package xrpc

import "github.com/smallnest/rpcx/server"

const (
	TCP = "tcp"
)

// addr形式: "localhost:8972"
func NewServer(service_name string, addr string) error {
	rpcx_server := server.NewServer()
	err := rpcx_server.RegisterName(service_name, new(Deposit), "")
	if err != nil {
		return err
	}
	err = rpcx_server.Serve(TCP, addr)
	if err != nil {
		return err
	}
	return nil
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
