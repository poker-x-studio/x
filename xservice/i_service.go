/*
功能：服务接口
说明：
*/
package xservice

type IService interface {
	Service_start() error
	Service_status() Status
	Service_stop() error
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
