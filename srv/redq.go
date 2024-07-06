package main

import (
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

func (app *application) getTotalBandwidthUssage(w http.ResponseWriter, r *http.Request) {
	conn, err := net.Dial("unix", sockPath)
	if err != nil {
		log.Fatalf("<FATAL>\t\tTotal Bandwith Ussage\n%s\n\n", err)
	}

	arg := RedqReqArg{
		Type: "TotalBandwidthUssage",
	}
	conn.Write()
}
