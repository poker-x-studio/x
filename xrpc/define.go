/*
功能：定义
说明：
*/
package xrpc

import "context"

type Args struct {
	Top_up_amount float64 // 充值额度
	User_id       string  //用户id
}

type Reply struct {
	OK bool //是否成功
}

type Deposit int

//充值
func (d *Deposit) Top_up(ctx context.Context, args *Args, reply *Reply) error {
	return nil
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
