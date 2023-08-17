/*
功能：client
说明：
*/
package xrpc

import "github.com/smallnest/rpcx/client"

func NewClient(service_name string, addr1 string, addr2 string) error {
	d, err := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: addr1}, {Key: addr2}})
	if err != nil {
		return err
	}
	xclient := client.NewXClient(service_name, client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()
	return nil
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
