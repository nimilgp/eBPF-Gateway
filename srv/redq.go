package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

const (
	sockPath = "/tmp/redq_ebpf.sock"
	bufSize  = 4096
)

type RedqReqArg struct {
	Type   string
	Action string
	Value  string
}

// func humanReadableBytes(bytes uint64) string {
// 	switch {
// 	case bytes < 1024:
// 		return fmt.Sprintf("%f Bytes", float64(bytes))
// 	case bytes < 1024*1024:
// 		return fmt.Sprintf("%.2f KiB", float64(bytes)/1024)
// 	case bytes < 1024*1024*1024:
// 		return fmt.Sprintf("%.2f MiB", float64(bytes)/1024/1024)
// 	default:
// 		return fmt.Sprintf("%.2f GiB", float64(bytes)/1024/1024/1024)
// 	}
// }

func (app *application) getRedq(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// tokenString := r.Header.Get("Tokenstring")
	// if !app.verifyAndUpdateBearerToken(tokenString) {
	// 	log.Printf("<WARNING>\t\t[(Redq)invalid bearer token]\n%s\n\n", tokenString)
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	buf := make([]byte, 4096)
	retBuf := make([]byte, 4096)
	conn, err := net.Dial("unix", sockPath)
	if err != nil {
		log.Fatalf("<FATAL>\t\tRedq socket\n%s\n\n", err)
	}
	count, err := r.Body.Read(buf)
	fmt.Println(string(buf))
	conn.Write(buf[:count])
	count, err = conn.Read(retBuf)
	log.Println(string(retBuf[:count]))
	w.Write(retBuf[:count])
}
