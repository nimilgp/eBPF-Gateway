package main

import (
	"log"
	"net"
)

func helloFunc(c net.Conn) {
	buf := make([]byte, 512)
	nr, err := c.Read(buf)
	if err != nil {
		return
	}
	data := buf[0:nr]
	println("Server got:", string(data))
	_, err = c.Write(data)
	if err != nil {
		log.Fatal("Write: ", err)
	}
}

func main() {
	listener, err := net.Listen("unix", "/tmp/echo.sock")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	fd, err := listener.Accept()
	if err != nil {
		log.Fatal("accept error:", err)
	}

	helloFunc(fd)
}
