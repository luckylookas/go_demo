package main

import (
	"bytes"
	"fmt"
	"net"
	"sync"
)

//does nto work yet, but demonstrates the intention

var sep = [...]byte{'\r', '\n', '\r', '\n'}
var http_ok = [...]byte{'H', 'T', 'T', 'P', '/', '1', '.', '1', ' ', '2', '0', '0'}
var http_error = [...]byte{'H', 'T', 'T', 'P', '/', '1', '.', '1', ' ', '5', '0', '0'}

var buffers []map[bool][]byte
var mutex = sync.Mutex{}

//pool of buffers to read content from tcp, just a demo
func reserveBuffer() ([]byte, int) {
	mutex.Lock()
	defer mutex.Unlock()
	for idx, buf := range buffers {
		if it, ok := buf[true]; ok {
			buf[false] = buf[true]
			buf[true] = nil
			return it, idx
		}
	}
	return nil, -1
}

func returnBuffer(index int) {
	mutex.Lock()
	defer mutex.Unlock()
	buffers[index][true] = buffers[index][false]
	buffers[index][false] = nil
}

func handleConnection(conn net.Conn) {

	buffer, index := reserveBuffer()
	defer returnBuffer(index)
	defer conn.Close()

	_, _ = conn.Read(buffer)

	if buffer[0] == 'P' {
		relevant := buffer[5:bytes.IndexByte(buffer, '\n')]
		req := relevant[:bytes.IndexByte(relevant, ' ')]
		fmt.Println("req:", string(req))
		last := false
		body := false

		for index, _ := conn.Read(buffer); index == 50; index, _ = conn.Read(buffer) {
			body, last = read(buffer[:index], body, last)
		}
		fmt.Println("MMMMMM")
		read(buffer[:index], body, last)


		conn.Write(http_ok[:])
	} else {
		conn.Write(http_error[:])
	}
}

func read(buffer []byte, body bool, last bool) (bool, bool) {
	if body {
		fmt.Print(string(buffer))
	} else if (last && buffer[0] == '\n') {
		body = true;
		fmt.Print(buffer[1:])
	} else {
		split := bytes.LastIndex(buffer, sep[:])
		if split > 0 {
			body = true
			fmt.Print(string(buffer[split:]))
		} else {
			last = buffer[49] == '\n'
		}
	}
	return body, last
}

func main() {
	buffers = make([]map[bool][]byte, 10)
	for i := 0; i < 10; i++ {
		buffers[i] = make(map[bool][]byte)
		buffers[i][true] = make([]byte, 50)
	}
	ln, _ := net.Listen("tcp", ":8080")

	for {
		conn, _ := ln.Accept()
		go handleConnection(conn)
	}
}
