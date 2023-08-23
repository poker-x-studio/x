/*
功能：服务状态
说明：
*/
package xservice

// 服务状态
type Status int

const (
	STATUS_NIL     Status = 0 //初始
	STATUS_RUNNING Status = 1 //运行中
	STATUS_DEAD    Status = 2 //关闭
)

type value struct {
	satus Status
	txt   string
}

var status_map map[Status]value

func init() {
	values := []value{
		{STATUS_NIL, "nil"},
		{STATUS_RUNNING, "running"},
		{STATUS_DEAD, "dead"},
	}
	status_map = make(map[Status]value, 0)
	for _, v := range values {
		status_map[v.satus] = v
	}
}

//String 转字符串
func (s Status) String() string {
	value, ok := status_map[s]
	if ok {
		return value.txt
	}
	return ""
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
