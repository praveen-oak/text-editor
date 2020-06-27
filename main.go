package main

import (
	"log"
	"text-editor/engine"
	"text-editor/netclient"

	"text-editor/data"
	"text-editor/parser"
	"text-editor/renderer"
)

func main() {
	client, err := netclient.NewSocketClient("127.0.0.1", "5005")
	dataStore := textdata.NewArrayStore()
	dataStore.AppendRow()
	arrayRenderer := renderer.NewRenderer(client)

	p := parser.NewSimpleParser()
	if err != nil {
		log.Fatalf("Error in creating connection to server. Error : %+v", err)
	}

	engine := engine.NewSimpleEngine(dataStore, arrayRenderer)
	for {
		input, err := client.Read()
		if err != nil {
			log.Print("Error in reading from server")
		}
		command, err := p.Parse(input)
		if command != nil {
			err = engine.Run(command)
			if err != nil {
				log.Print("Error in running command")
			}
		}
	}
}
