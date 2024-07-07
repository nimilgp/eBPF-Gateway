package main

import (
	"log"
	"net"
	"os"
)

const SockAddr = "/tmp/uds.sock"

func helloFunc(c net.Conn) {
	buf := make([]byte, 512)
	nr, err := c.Read(buf)
	if err != nil {
		return
	}
	data := buf[0:nr]
	println("Server got:", string(data))
	sendtoclient := "Echo from server :" + string(data) + "\n"
	// log.Printf(sendtoclient)
	_, err = c.Write([]byte(sendtoclient))
	if err != nil {
		log.Fatal("Write: ", err)
	}
}

func main() {
	if err := os.RemoveAll(SockAddr); err != nil {
		log.Fatal(err)
	}
	listener, err := net.Listen("unix", SockAddr)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer listener.Close()
	fd, err := listener.Accept()
	if err != nil {
		log.Fatal("accept error:", err)
	}
	for {
		helloFunc(fd)
	}
}
