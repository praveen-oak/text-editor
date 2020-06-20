package main

import (
	"log"
	engine2 "text-editor/engine"
	"text-editor/netclient"

	textdata "text-editor/data"
	parser "text-editor/parser"
	. "text-editor/renderer"
)

func main() {
	client, err := netclient.NewSocketClient("127.0.0.1", "5005")
	dataStore := textdata.NewArrayStore()
	dataStore.AppendRow()
	arrayRenderer := NewRenderer(client)

	p := parser.NewSimpleParser()
	if err != nil {
		log.Fatalf("Error in creating connection to server. Error : %+v", err)
	}

	engine := engine2.NewSimpleEngine(dataStore, arrayRenderer)
	for {
		input, err := client.Read()
		if err != nil {
			log.Print("Error in reading from server")
		}
		command, err := p.Parse(input)
		if command != nil {
			engine.Run(command)
		}
		arrayRenderer.RefreshScreen(dataStore, 803, 603)
	}
}
