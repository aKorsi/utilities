package utilities

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"math"
	"runtime"
)

func ServerInfo() (info string, err error) {
	lc, err := cpu.Counts(true)
	if err != nil {
		return
	}
	cpuInfo, err := cpu.Info()
	if err != nil {
		return
	}
	hostInfo, err := host.Info()
	if err != nil {
		return
	}
	sv, err := mem.VirtualMemory()
	if err != nil {
		return
	}
	sm, err := mem.SwapMemory()
	if err != nil {
		return
	}
	info += fmt.Sprintln("CPU Model:", cpuInfo[0].ModelName)
	info += fmt.Sprintln("CPU Frequency:", cpuInfo[0].Mhz/1000, "GHz")
	info += fmt.Sprintln("CPU Logical Cores:", lc)
	info += fmt.Sprintln("Ram Total:", float64(sv.Total)/math.Pow(1024, 3), "GB")
	info += fmt.Sprintln("Ram Available:", float64(sv.Available)/math.Pow(1024, 3), "GB")
	info += fmt.Sprintln("Ram UsedPercent:", sv.UsedPercent, "%")
	info += fmt.Sprintln("Swap Total:", float64(sm.Total)/math.Pow(1024, 3), "GB")
	info += fmt.Sprintln("Swap Available:", float64(sm.Free)/math.Pow(1024, 3), "GB")
	info += fmt.Sprintln("Swap UsedPercent:", sm.UsedPercent, "%")
	info += fmt.Sprintln("OS:", hostInfo.OS)
	info += fmt.Sprintln("OS Architecture:", runtime.GOARCH)
	info += fmt.Sprintln("OS Hostname:", hostInfo.Hostname)
	info += fmt.Sprintln("OS KernelVersion:", hostInfo.KernelVersion)
	info += fmt.Sprintln("OS Platform:", hostInfo.Platform)
	info += fmt.Sprintln("OS PlatformFamily:", hostInfo.PlatformFamily)
	info += fmt.Sprintln("OS PlatformVersion:", hostInfo.PlatformVersion)

	return
}
