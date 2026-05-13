package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	var currentLine string
	ch := make(chan string)

	go func() {
		defer close(ch)
		for {
			data := make([]byte, 8)
			n, err := f.Read(data)
			if err != nil {
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
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("error")
	}
	defer f.Close()

	ch := getLinesChannel(f)
	for line := range ch {
		fmt.Printf("read: %s\n", line)
	}
}
