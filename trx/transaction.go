/*
功能：转账
说明：

1 官方公共节点
https://developers.tron.network/v3.7/docs/official-public-node

2 TRX转账交易 流程 https://cn.developers.tron.network/docs/tron-protocol-transaction
2.1 创建交易:
2.2 签名交易：发送交易前，必须由发送方使用私钥对其进行签名。
2.3 广播交易: 用户将交易发送到节点后，节点会尝试在其本地执行并验证该交易, 有效的交易将被继续广播到其他节点，无效的交易被该节点丢弃，这将有效防止垃圾交易在网络中无效广播，占用网络资源。
2.4 交易确认
*/
package trx

import (
	"math/rand"
	"strconv"

	"github.com/poker-x-studio/x/xdebug"
	"github.com/poker-x-studio/x/xlog"

	"github.com/JFJun/trx-sign-go/grpcs"
	"github.com/JFJun/trx-sign-go/sign"
	"github.com/fbsobreira/gotron-sdk/pkg/common"
)

const (
	SUN_AMOUNT = 1     //转账sun个数,1TRX = 1000000sun
	GRPC_PORT  = 50051 //端口
)

// Transaction trx转账交易
func Transaction(addr_from string, private_key_from string, addr_to string) (string, error) {
	xlog_entry := xlog.New_entry(xdebug.Funcname())
	node_ips := []string{
		"3.225.171.164",
		"52.53.189.99",
		"18.196.99.16",
		"34.253.187.192",
		"52.56.56.149",
		"35.180.51.163",
		"54.252.224.209",
		"18.228.15.36",
		"52.15.93.92",
		"34.220.77.106",
	}

	index := rand.Intn(len(node_ips))
	node_address := node_ips[index] + ":" + strconv.Itoa(GRPC_PORT)
	xlog_entry.Tracef("node_address:%s", node_address)

	client, err := grpcs.NewClient(node_address)
	if err != nil {
		xlog_entry.Errorf("创建client,失败,err:%s", err.Error())
		return "", err
	}

	//创建一个未签名的转账
	tx, err := client.Transfer(addr_from, addr_to, SUN_AMOUNT)
	if err != nil {
		xlog_entry.Errorf("创建一个未签名的转账,失败,err:%s", err.Error())
		return "", err
	}

	//签名交易
	signTx, err := sign.SignTransaction(tx.Transaction, private_key_from)
	if err != nil {
		xlog_entry.Errorf("签名交易,失败,err:%s", err.Error())
		return "", err
	}

	//广播交易
	err = client.BroadcastTransaction(signTx)
	if err != nil {
		xlog_entry.Errorf("广播交易,失败,err:%s", err.Error())
		return "", err
	}

	tx_id := common.BytesToHexString(tx.GetTxid())
	xlog_entry.Tracef("交易完成,tx_id:%s", tx_id)
	return tx_id, err
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
