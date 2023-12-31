/*
功能：系统信息
说明：
1 Golang实践录：获取系统信息 https://juejin.cn/post/7104069654733144078
*/
package xdebug

import (
	"fmt"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/v3/mem"
)

type LSysInfo struct {
	MemAll         uint64
	MemFree        uint64
	MemUsed        uint64
	MemUsedPercent float64
	Days           int64
	Hours          int64
	Minutes        int64
	Seconds        int64

	CpuUsedPercent float64
	OS             string
	Arch           string
	CpuCores       int
}

// Get_sys_info 获取系统信息
func Get_sys_info() (info LSysInfo) {
	unit := uint64(1024 * 1024) // MB

	v, _ := mem.VirtualMemory()

	info.MemAll = v.Total
	info.MemFree = v.Free
	info.MemUsed = info.MemAll - info.MemFree
	// 注：使用SwapMemory或VirtualMemory，在不同系统中使用率不一样，因此直接计算一次
	info.MemUsedPercent = float64(info.MemUsed) / float64(info.MemAll) * 100.0 // v.UsedPercent
	info.MemAll /= unit
	info.MemUsed /= unit
	info.MemFree /= unit

	info.OS = runtime.GOOS
	info.Arch = runtime.GOARCH
	info.CpuCores = runtime.GOMAXPROCS(0)

	// 获取200ms内的CPU信息，太短不准确，也可以获几秒内的，但这样会有延时，因为要等待
	cc, _ := cpu.Percent(time.Millisecond*200, false)
	info.CpuUsedPercent = cc[0]

	// 获取开机时间
	boottime, _ := host.BootTime()
	ntime := time.Now().Unix()
	btime := time.Unix(int64(boottime), 0).Unix()
	deltatime := ntime - btime

	info.Seconds = int64(deltatime)
	info.Minutes = info.Seconds / 60
	info.Seconds -= info.Minutes * 60
	info.Hours = info.Minutes / 60
	info.Minutes -= info.Hours * 60
	info.Days = info.Hours / 24
	info.Hours -= info.Days * 24

	if Is_debug() {
		fmt.Printf("info: %#v\n", info)
	}
	return
}

// startup_output 启动输出
func startup_output() {
	c, _ := cpu.Info()
	cc, _ := cpu.Percent(time.Second, false) // 1秒
	d, _ := disk.Usage("/")
	n, _ := host.Info()
	nv, _ := net.IOCounters(true)
	physicalCnt, _ := cpu.Counts(false)
	logicalCnt, _ := cpu.Counts(true)
	if len(c) > 1 {
		for _, sub_cpu := range c {
			modelname := sub_cpu.ModelName
			cores := sub_cpu.Cores
			if Is_debug() {
				fmt.Printf("CPUs: %v   %v cores \n", modelname, cores)
			}
		}
	} else {
		sub_cpu := c[0]
		modelname := sub_cpu.ModelName
		cores := sub_cpu.Cores
		if Is_debug() {
			fmt.Printf("CPU: %v   %v cores \n", modelname, cores)
		}
	}
	if Is_debug() {
		fmt.Printf("physical count:%d logical count:%d\n", physicalCnt, logicalCnt)
		fmt.Printf("CPU Used: used %f%%\n", cc[0])
		fmt.Printf("HD: %v GB Free: %v GB Usage:%f%%\n", d.Total/1024/1024/1024, d.Free/1024/1024/1024, d.UsedPercent)
		fmt.Printf("OS: %v(%v) %v\n", n.Platform, n.PlatformFamily, n.PlatformVersion)
		fmt.Printf("Hostname: %v\n", n.Hostname)
		fmt.Printf("Network: %v bytes / %v bytes\n", nv[0].BytesRecv, nv[0].BytesSent)
	}
}

//-----------------------------------------------
//					the end
//-----------------------------------------------
