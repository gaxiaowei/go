package main

import (
	"fmt"
	"net"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	var buf = make([]byte, 1024)

	fmt.Println(conn.RemoteAddr())

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn read error:", err)
			return
		}

		fmt.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
	}
}

func main() {
	l, err := net.Listen("tcp", ":8001")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			return
		}

		go handleConn(conn)
	}
}
