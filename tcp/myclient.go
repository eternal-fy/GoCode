package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	defer conn.Close()
	if err != nil {
		fmt.Println("error")
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		str, _ := reader.ReadString('\n')
		str = strings.Trim(str, "\r\n")
		if strings.ToUpper(str) == "Q" {
			return
		}
		str = "context: " + str
		conn.Write([]byte(str))
		buf := make([]byte, 2048)
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("error")
		}
		fmt.Println("收到客户端返回的信息：", string(buf[:n]))
	}

}
