package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/sensors"
	"github.com/zcalusic/sysinfo"
)

func humanReadableBytes(bytes uint64) string {
	switch {
	case bytes < 1024:
		return fmt.Sprintf("%f Bytes", float64(bytes))
	case bytes < 1024*1024:
		return fmt.Sprintf("%.2f KiB", float64(bytes)/1024)
	case bytes < 1024*1024*1024:
		return fmt.Sprintf("%.2f MiB", float64(bytes)/1024/1024)
	default:
		return fmt.Sprintf("%.2f GiB", float64(bytes)/1024/1024/1024)
	}
}

type hardwareUssageStruct struct {
	CpuPercUsed float64
	RamPercUsed float64
	Temperature float64
}

func (app *application) getHardwareUssage(w http.ResponseWriter, r *http.Request) {
	virtMem, _ := mem.VirtualMemory()
	cpuPerc, _ := cpu.Percent(0, false)
	hardware := hardwareUssageStruct{
		CpuPercUsed: cpuPerc[0],
		RamPercUsed: virtMem.UsedPercent,
	}
	temps, _ := sensors.SensorsTemperatures()
	for i := 0; i < len(temps); i++ {
		if temps[i].SensorKey == "coretemp_package_id_0" {
			hardware.Temperature = temps[i].Temperature
			break
		}
	}
	json.NewEncoder(w).Encode(hardware)
}

func (app *application) getHardwareDetails(w http.ResponseWriter, r *http.Request) {
	var si sysinfo.SysInfo
	si.GetSysInfo()

	data, err := json.Marshal(&si)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(data)

}
