package sys

import (
	"math"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/wailsapp/wails"
)

// Struct used for obtain Stats fom CPU.
type Stats struct {
	log *wails.CustomLogger
}

// struct used for obtain CPU Usage.
type CPUUsage struct {
	Average int `json:"avg"`
}

// struct used for obtain Disk Usages.
type DISKUsage struct {
	Used int `json:"used"`
	Free int `json:"free"`
}

// Function return data in a time interval.
/*func (s *Stats) WailsInit(runtime *wails.Runtime) error {
	s.log = runtime.Log.New("Stats")
	return nil
}*/
func (s *Stats) WailsInit(runtime *wails.Runtime) error {
	s.log = runtime.Log.New("Stats")

	go func() {
		for {
			runtime.Events.Emit("cpu_usage", s.GetCPUUsage())
			time.Sleep(1 * time.Second)
		}
	}()

	return nil
}

// Function return data from CPU Usage.
func (s *Stats) GetCPUUsage() *CPUUsage {
	percent, err := cpu.Percent(1*time.Second, false)
	if err != nil {
		s.log.Errorf("unable to get cpu stats: %s", err.Error())
		return nil
	}

	return &CPUUsage{
		Average: int(math.Round(percent[0])),
	}
}

// Function return data from Disk Usage (Used space and Free space).
func (s *Stats) GetDISKUsage() *DISKUsage {
	UsageStats, err := disk.Usage("/")
	if err != nil {
		s.log.Errorf("unable to get disk stats: %s", err.Error())
		return nil
	}

	return &DISKUsage{
		Used: int(UsageStats.Used),
		Free: int(UsageStats.Free),
	}
}
