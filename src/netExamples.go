package main

import (
	"bufio"
	"fmt"
	"net"
)

func netEntry() {
	netBasic()
}

func netBasic() {
	conn, err := net.Dial("tcp", "baidu.com:http")
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	_ = status
}

func httpBasic() {

}
