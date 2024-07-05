package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pbnjay/memory"
)

func convertToHumanReadable(bytes uint64) string {
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
	totalMem := memory.TotalMemory()
	freeMem := memory.FreeMemory()
	hTotalMem := convertToHumanReadable(totalMem)
	hFreeMem := convertToHumanReadable(freeMem)

	log.Printf("Total Memory : %s", hTotalMem)
	log.Printf("Free Memory : %s", hFreeMem)
}
