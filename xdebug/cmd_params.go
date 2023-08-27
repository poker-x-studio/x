/*
功能：命令行参数
说明：
*/
package xdebug

// CmdParamsBase 命令行参数
type CmdParamsBase struct {
	Debug        bool   `short:"d" long:"debug" description:"是否debug"`
	Log_filename string `short:"l" long:"log_filename" description:"日志文件名"`
	Config_file  string `short:"c" long:"config_file" description:"配置文件"`
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
