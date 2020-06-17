package main

import (
	"log"
	"text-editor/netclient"

	textdata "text-editor/data"
	"text-editor/renderer"
)

func main() {
	client, err := netclient.NewSocketClient("127.0.0.1", "5005")
	dataStore := textdata.NewArrayStore()
	dataStore.AppendRow()
	dataStore.AppendChar(0, 'a')
	dataStore.AppendChar(0, 'b')
	dataStore.AppendChar(0, 'c')
	dataStore.AppendChar(0, 'd')
	arrayRenderer := renderer.NewRenderer(dataStore, client)

	if err != nil {
		log.Fatalf("Error in creating connection to server. Error : %+v", err)
	}

	for {
		_, err := client.Read()
		if err != nil {
			log.Print("Error in reading from server")
		}
		arrayRenderer.RefreshScreen(803, 603)
	}
}
