package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/sensors"
	"github.com/zcalusic/sysinfo"
)

type hardwareUssageStruct struct {
	CpuPercUsed float64
	RamPercUsed float64
	Temperature float64
}

func (app *application) getHardwareUssage(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Tokenstring")
	if !app.verifyAndUpdateBearerToken(tokenString) {
		log.Printf("<WARNING>\t\t[(Harware-Ussage)invalid bearer token]\n%s\n\n", tokenString)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	virtMem, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("<ERROR>\t\t[(Hardware-Ussage) failed to get memory ussage]\n%s\n\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	cpuPerc, err := cpu.Percent(0, false)
	if err != nil {
		log.Printf("<ERROR>\t\t[(Hardware-Ussage) failed to get cpu ussage]\n%s\n\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	hardware := hardwareUssageStruct{
		CpuPercUsed: cpuPerc[0],
		RamPercUsed: virtMem.UsedPercent,
	}
	temps, err := sensors.SensorsTemperatures()
	if err != nil {
		log.Printf("<ERROR>\t\t[(Hardware-Ussage) failed to get cpu temp]\n%s\n\n", err)
	}
	for i := 0; i < len(temps); i++ {
		if temps[i].SensorKey == "coretemp_package_id_0" {
			hardware.Temperature = temps[i].Temperature
			break
		}
	}
	log.Printf("<INFO>\t\t[(Harware-Ussage) success]\n\n")
	json.NewEncoder(w).Encode(hardware)
}

func (app *application) getHardwareDetails(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Tokenstring")
	if !app.verifyAndUpdateBearerToken(tokenString) {
		log.Printf("<WARNING>\t\t[(Harware-Details)invalid bearer token]\n%s\n\n", tokenString)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	var si sysinfo.SysInfo
	si.GetSysInfo()

	data, err := json.Marshal(&si)
	if err != nil {
		log.Printf("<ERROR>\t\t[(Harware-Details)marshalleling failed]\n%s\n\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("<INFO>\t\t[(Harware-Details) success]\n\n")
	w.Write(data)
}
