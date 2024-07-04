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
		return fmt.Sprintf("%v Bytes", bytes)
	case bytes < 1024*1024:
		return fmt.Sprintf("%v KiB", bytes>>10)
	case bytes < 1024*1024*1024:
		return fmt.Sprintf("%v MiB", bytes>>20)
	default:
		return fmt.Sprintf("%v GiB", bytes>>30)
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
