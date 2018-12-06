package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8887")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}

		go serve(connection)
	}
}

func serve(connection net.Conn) {
	defer connection.Close()

	scanner := bufio.NewScanner(connection)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if ln == "" {
			break
		}
	}

	body := "STATUS"

	io.WriteString(connection, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(connection, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(connection, "Content-Type: text/plain\r\n")
	io.WriteString(connection, "\r\n")
	io.WriteString(connection, body)
}
