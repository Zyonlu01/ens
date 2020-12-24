package main

import (
	"net"
)

func main() {
	for {

		abc := make([]byte, 1000000)
		conn, _ := net.Dial("tcp", "192.168.1.1:80")
		conn.Write(abc)
	}

}
//hi knk