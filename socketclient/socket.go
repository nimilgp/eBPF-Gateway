package main

import (
	"log"
	"net"
)

const SockAddr = "/tmp/uds.sock"

func hello() {
	conn, err := net.Dial("unix", SockAddr)
	if err != nil {
		log.Fatalf("<FATAL>\t\tTotal Bandwith Ussage\n%s\n\n", err)
	}
	conn.Write([]byte("helloooooooo"))

	buf := make([]byte, 512)
	nr, err := conn.Read(buf)
	if err != nil {
		return
	}
	data := buf[0:nr]
	println(string(data))
	conn.Close()
}

func main() {
	hello()
}
