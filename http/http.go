package demo

import (
	"bufio"
	"fmt"
	_ "io"
	"net"
	"unsafe"
)

type Query struct {
	key string
}
type Response struct {
	key string
}

func Process(c *Cache, conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read failed,err", err)
			break
		}
		recvstr := string(buf[:n])
		if v, ok := c.Map_Cache[recvstr]; ok {
			conn.Write([]byte(v))
			return
		}
		fmt.Println("received:", recvstr)
		conn.Write([]byte(""))
	}
}
func Accept() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go Process(conn) // 启动一个goroutine处理连接
	}
}

func Send(key string) string {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("err", err)
		return ""
	}
	defer conn.Close()
	var req []byte
	req = []byte(key)
	_, err = conn.Write(req)
	if err != nil {
		return ""
	}
	buf := [128]byte{}
	n, err := conn.Read(buf[:])
	if err != nil {
		return ""
	}
	return unsafe.String(&buf[0], n)
}
