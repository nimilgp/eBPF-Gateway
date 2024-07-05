package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/shirou/gopsutil/v4/mem"
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

func (app *application) getHardwareMemoryUssage(w http.ResponseWriter, r *http.Request) {
	virtMem, _ := mem.VirtualMemory()

	hTotalMem := humanReadableBytes(virtMem.Total)
	hAvailibleMem := humanReadableBytes(virtMem.Available)
	hUsedMem := humanReadableBytes(virtMem.Used)
	hFreeMem := humanReadableBytes(virtMem.Free)

	log.Printf("Total Memory : %s", hTotalMem)
	log.Printf("Availible Memory : %s", hAvailibleMem)
	log.Printf("Used Memory : %s", hUsedMem)
	log.Printf("Free Memory : %s", hFreeMem)
	log.Printf("Used Percentage: %f", virtMem.UsedPercent)
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

// {
// 	"sysinfo": {
// 		"version": "1.1.0",
// 		"timestamp": "2024-07-05T11:10:43.073358992+05:30"
// 	},
// 	"node": {
// 		"hostname": "nixdesk",
// 		"machineid": "4e1417435c17417e9437101eb93f6337"
// 	},
// 	"os": {
// 		"name": "NixOS 24.05 (Uakari)",
// 		"vendor": "nixos",
// 		"version": "24.05",
// 		"architecture": "amd64"
// 	},
// 	"kernel": {
// 		"release": "6.6.34",
// 		"version": "#1-NixOS SMP PREEMPT_DYNAMIC Sun Jun 16 11:47:49 UTC 2024",
// 		"architecture": "x86_64"
// 	},
// 	"product": {
// 		"name": "B760M DS3H AX DDR4",
// 		"vendor": "Gigabyte Technology Co., Ltd.",
// 		"version": "Default string",
// 		"serial": "Default string",
// 		"uuid": "03560274-043c-0527-0606-2a0700080009",
// 		"sku": "Default string"
// 	},
// 	"board": {
// 		"name": "B760M DS3H AX DDR4",
// 		"vendor": "Gigabyte Technology Co., Ltd.",
// 		"version": "x.x",
// 		"serial": "Default string",
// 		"assettag": "Default string"
// 	},
// 	"chassis": {
// 		"type": 3,
// 		"vendor": "Default string",
// 		"version": "Default string",
// 		"serial": "Default string",
// 		"assettag": "Default string"
// 	},
// 	"bios": {
// 		"vendor": "American Megatrends International, LLC.",
// 		"version": "F1",
// 		"date": "10/12/2022"
// 	},
// 	"cpu": {
// 		"vendor": "GenuineIntel",
// 		"model": "13th Gen Intel(R) Core(TM) i5-13500",
// 		"speed": 2475,
// 		"cache": 24576,
// 		"cpus": 1,
// 		"cores": 14,
// 		"threads": 20
// 	},
// 	"memory": {
// 		"type": "DDR4",
// 		"speed": 3200,
// 		"size": 16384
// 	},
// 	"storage": [
// 		{
// 			"name": "nvme0n1",
// 			"model": "CT1000P3SSD8",
// 			"serial": "2317E6CF33A3",
// 			"size": 1000
// 		},
// 		{
// 			"name": "nvme1n1",
// 			"model": "CT500P3SSD8",
// 			"serial": "2243E67D1C0C",
// 			"size": 500
// 		},
// 		{
// 			"name": "sda",
// 			"driver": "sd",
// 			"vendor": "BR25",
// 			"model": "UDISK",
// 			"serial": "1120030306090106"
// 		}
// 	],
// 	"network": [
// 		{
// 			"name": "enp7s0",
// 			"driver": "r8169",
// 			"macaddress": "74:56:3c:27:06:2a",
// 			"port": "tp/mii",
// 			"speed": 1000
// 		},
// 		{
// 			"name": "wlp6s0",
// 			"driver": "iwlwifi",
// 			"macaddress": "72:7c:06:65:bc:6e"
// 		}
// 	]
// }
