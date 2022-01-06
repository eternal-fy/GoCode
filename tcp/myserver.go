package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("connect error ", err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("error", err)
		}
		go run(conn)
	}

}

func run(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	buf := make([]byte, 2048)
	n, err := reader.Read(buf[:])
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println("tcp端传来数据：", string(buf[:n]))
	conn.Write(buf[:n])
}
