package sys

import (
	"math"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/wailsapp/wails"

	"github.com/shirou/gopsutil/mem"
)

// Stats .
type Stats struct {
	log *wails.CustomLogger
}

// CPUUsage .
type CPUUsage struct {
	Average int `json:"avg"`
}

// DISKUsage .
type DISKUsage struct {
	Used int `json:"used"`
	Free int `json:"free"`
}

// MEMORYUsage .
type MEMORYUsage struct {
	Total     int `json:"total"`
	Used      int `json:"used"`
	Free      int `json:"free"`
	Available int `json:"Available"`
}

// WailsInit .
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

// GetCPUUsage .
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

// GetDISKUsage
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

func (s *Stats) GetMemoryUsage() *MEMORYUsage {
	UsageStats, err := mem.VirtualMemory()
	if err != nil {
		s.log.Errorf("unable to get memory stats: %s", err.Error())
		return nil
	}

	return &MEMORYUsage{
		Total:     int(UsageStats.Total),
		Used:      int(UsageStats.Used),
		Free:      int(UsageStats.Free),
		Available: int(UsageStats.Available),
	}
}
