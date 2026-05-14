package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	add, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		panic(err)
	}
	conn, err := net.DialUDP("udp", nil, add)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		_, err = conn.Write([]byte(text))
		if err != nil {
			log.Fatal(err)
		}
	}
}
