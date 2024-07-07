package main

import (
	"fmt"
	"net/http"
)

const (
	sockPath = "/tmp/uds.sock"
	bufSize  = 4096
)

type RedqReqArg struct {
	Type   string
	Action string
	Value  string
}

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

func (app *application) getTotalBandwidthUssage(w http.ResponseWriter, r *http.Request) {
	// conn, err := net.Dial("unix", sockPath)
	// if err != nil {
	// 	log.Fatalf("<FATAL>\t\tTotal Bandwith Ussage\n%s\n\n", err)
	// }

	// arg := RedqReqArg{
	// 	Type:   "TotalBandwidthUssage",
	// 	Action: "",
	// 	Value:  "",
	// }
	// json.NewEncoder(w).Encode(hardware)
	// conn.Write()
}
