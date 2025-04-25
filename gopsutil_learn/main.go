package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"time"
)

// CpuInfoStruct 定义CPU信息的结构体
type CpuInfoStruct struct {
	CpuPercent float64 `json:"cpu_percent"` // CPU使用率
	UserTime   float64 `json:"user_time"`   // 用户CPU时间
	SystemTime float64 `json:"system_time"` // 系统CPU时间
	IdleTime   float64 `json:"idle_time"`   // 空闲CPU时间
}

// MemInfoStruct 定义内存信息的结构体
type MemInfoStruct struct {
	Total           uint64  `json:"total"`             // 总内存
	Used            uint64  `json:"used"`              // 已用内存
	Available       uint64  `json:"available"`         // 可用内存
	UsedPercent     float64 `json:"used_percent"`      // 内存使用率
	SwapTotal       uint64  `json:"swap_total"`        // 总交换分区
	SwapUsed        uint64  `json:"swap_used"`         // 已用交换分区
	SwapAvailable   uint64  `json:"swap_available"`    // 可用交换分区
	SwapUsedPercent float64 `json:"swap_used_percent"` // 交换分区使用率
}

// DiskInfoStruct 定义磁盘信息的结构体
type DiskInfoStruct struct {
	Mountpoint  string  `json:"mountpoint"`   // 挂载点
	Total       uint64  `json:"total"`        // 总磁盘大小
	Used        uint64  `json:"used"`         // 已用磁盘大小
	Free        uint64  `json:"free"`         // 可用磁盘大小
	UsedPercent float64 `json:"used_percent"` // 磁盘使用率
}

// NetInfoStruct 定义网络信息的结构体
type NetInfoStruct struct {
	BytesSent   uint64 `json:"bytes_sent"`   // 发送字节数
	BytesRecv   uint64 `json:"bytes_recv"`   // 接收字节数
	PacketsSent uint64 `json:"packets_sent"` // 发送包数
	PacketsRecv uint64 `json:"packets_recv"` // 接收包数
}

// GetCpuInfo 获取CPU信息并返回结构体
func GetCpuInfo() CpuInfoStruct {
	percent, _ := cpu.Percent(time.Second, false)
	times, _ := cpu.Times(false) // 获取总的CPU时间
	return CpuInfoStruct{
		CpuPercent: percent[0],
		UserTime:   times[0].User,
		SystemTime: times[0].System,
		IdleTime:   times[0].Idle,
	}
}

// GetMemInfo 获取内存信息并返回结构体
func GetMemInfo() MemInfoStruct {
	memInfo, _ := mem.VirtualMemory()
	swapInfo, _ := mem.SwapMemory()
	return MemInfoStruct{
		Total:           memInfo.Total / 1024 / 1024 / 1024,
		Used:            memInfo.Used / 1024 / 1024,
		Available:       memInfo.Available / 1024 / 1024,
		UsedPercent:     memInfo.UsedPercent,
		SwapTotal:       swapInfo.Total / 1024 / 1024 / 1024,
		SwapUsed:        swapInfo.Used / 1024 / 1024,
		SwapAvailable:   swapInfo.Free / 1024 / 1024,
		SwapUsedPercent: swapInfo.UsedPercent,
	}
}

// GetDiskInfo 获取磁盘信息并返回结构体
func GetDiskInfo() []DiskInfoStruct {
	parts, _ := disk.Partitions(true)
	var diskInfos []DiskInfoStruct
	for _, part := range parts {
		diskUsage, _ := disk.Usage(part.Mountpoint)
		diskInfo := DiskInfoStruct{
			Mountpoint:  part.Mountpoint,
			Total:       diskUsage.Total / 1024 / 1024 / 1024,
			Used:        diskUsage.Used / 1024 / 1024,
			Free:        diskUsage.Free / 1024 / 1024,
			UsedPercent: diskUsage.UsedPercent,
		}
		diskInfos = append(diskInfos, diskInfo)
	}
	return diskInfos
}

// GetRootDirectoryDiskInfo 获取根目录磁盘信息并返回结构体
func GetRootDirectoryDiskInfo() []DiskInfoStruct {
	// 指定挂载点为根目录
	diskUsage, _ := disk.Usage("/")
	diskInfo := DiskInfoStruct{
		Mountpoint:  "/",
		Total:       diskUsage.Total / 1024 / 1024 / 1024, // 转换为 GB
		Used:        diskUsage.Used / 1024 / 1024,         // 转换为 MB
		Free:        diskUsage.Free / 1024 / 1024,         // 转换为 MB
		UsedPercent: diskUsage.UsedPercent,
	}

	// 返回包含根目录磁盘信息的切片
	return []DiskInfoStruct{diskInfo}
}

// GetNetInfo 获取网络信息并返回结构体
func GetNetInfo() NetInfoStruct {
	netIO, _ := net.IOCounters(false)
	return NetInfoStruct{
		BytesSent:   netIO[0].BytesSent,
		BytesRecv:   netIO[0].BytesRecv,
		PacketsSent: netIO[0].PacketsSent,
		PacketsRecv: netIO[0].PacketsRecv,
	}
}

// PrintCpuInfo 打印CPU信息
func PrintCpuInfo(cpuInfo CpuInfoStruct) {
	fmt.Printf("CPU使用率: %.2f%%, 用户时间: %.2fs, 系统时间: %.2fs, 空闲时间: %.2fs\n",
		cpuInfo.CpuPercent, cpuInfo.UserTime, cpuInfo.SystemTime, cpuInfo.IdleTime)
}

// PrintMemInfo 打印内存信息
func PrintMemInfo(memInfo MemInfoStruct) {
	fmt.Printf("总内存: %v GB, 已用内存: %v MB, 可用内存: %v MB, 内存使用率: %.3f %%\n",
		memInfo.Total, memInfo.Used, memInfo.Available, memInfo.UsedPercent)
	fmt.Printf("总交换分区: %v GB, 已用交换分区: %v MB, 可用交换分区: %v MB, 交换分区使用率: %.3f %%\n",
		memInfo.SwapTotal, memInfo.SwapUsed, memInfo.SwapAvailable, memInfo.SwapUsedPercent)
}

// PrintDiskInfo 打印磁盘信息
func PrintDiskInfo(diskInfos []DiskInfoStruct) {
	for _, diskInfo := range diskInfos {
		fmt.Printf("挂载点: %v, 总磁盘大小: %v GB, 已用磁盘: %v MB, 可用磁盘: %v MB, 磁盘使用率: %.2f %%\n",
			diskInfo.Mountpoint, diskInfo.Total, diskInfo.Used, diskInfo.Free, diskInfo.UsedPercent)
	}
}

// PrintNetInfo 打印网络信息
func PrintNetInfo(netInfo NetInfoStruct) {
	fmt.Printf("发送字节数: %v, 接收字节数: %v, 发送包数: %v, 接收包数: %v\n",
		netInfo.BytesSent, netInfo.BytesRecv, netInfo.PacketsSent, netInfo.PacketsRecv)
}

func main() {
	// 获取CPU信息
	cpuInfo := GetCpuInfo()
	PrintCpuInfo(cpuInfo)

	// 获取内存信息
	memInfo := GetMemInfo()
	PrintMemInfo(memInfo)

	// 获取磁盘信息
	//diskInfos := GetDiskInfo()
	//PrintDiskInfo(diskInfos)

	// 获取根目录磁盘信息
	rootDiskInfo := GetRootDirectoryDiskInfo()
	PrintDiskInfo(rootDiskInfo)

	// 获取网络信息
	netInfo := GetNetInfo()
	PrintNetInfo(netInfo)
}
