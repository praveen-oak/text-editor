package main

import (
	"bufio"
	"log"
	"net"
	textdata "text-editor/data"
	"text-editor/data/renderer"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:5005")
	if err != nil {
		log.Fatalf("Error in creating tcp address. Exiting")
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatalf("Error in creating tcp connection. Exiting")
	}
	data := make([]byte, 1024)
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	dataStore := textdata.NewArrayStore()

	arrayRenderer := renderer.NewRenderer(dataStore, writer)
	for {
		i, err := reader.Read(data)
		if err != nil {
			log.Print("Error in reading from server")
		}
		log.Print("Read bytes from server data", i, string(data[:i]))
		arrayRenderer.RefreshScreen(803, 603)
	}

}
