package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

func getLinesChannel(conn net.Conn) <-chan string {
	var currentLine string
	ch := make(chan string)

	go func() {
		defer close(ch)
		for {
			data := make([]byte, 8)
			n, err := conn.Read(data)
			if err != nil {
				if err != io.EOF {
					fmt.Println("error: ", err)
				}
				if currentLine != "" {
					ch <- currentLine
				}
				break
			}
			split := strings.Split(string(data[:n]), "\n")
			if len(split) > 1 {
				currentLine += split[0]

				ch <- currentLine

				currentLine = split[1]
				continue
			}
			currentLine += strings.Join(split, "")
		}
	}()

	return ch
}

func main() {
	listner, err := net.Listen("tcp", ":42069")
	if err != nil {
		panic(err)
	}
	defer listner.Close()

	for {
		conn, err := listner.Accept()
		if err != nil {
			panic(err)
		}
		ch := getLinesChannel(conn)

		for line := range ch {
			fmt.Printf("%s\r\n", line)
		}

		conn.Close()
	}
}
