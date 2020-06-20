package renderer

import (
	"log"
	"strconv"
	"strings"
	"text-editor/data"
	"text-editor/netclient"
)

const (
	row_char = 14
	col_char = 8
)

type Renderer interface {
	RefreshScreen(dataStore textdata.DataStore, windowLength, windowHeight int) error
}

type SimpleRenderer struct {
	client netclient.Client
}

func NewRenderer(client netclient.Client) Renderer {
	return &SimpleRenderer{
		client: client,
	}
}

func (s *SimpleRenderer) RefreshScreen(dataStore textdata.DataStore, windowLength, windowHeight int) error {
	rowPixel := 0
	colPixel := 0

	rowSize := windowLength / 8
	colSize := windowHeight / 8

	serverString := make([]string, 5)
	serverString[0] = "text"
	serverString[3] = "#000000"

	for {
		value, err := dataStore.ReadChar(rowPixel, colPixel)
		if err != nil {
			return nil
		}
		err = s.send(rowPixel, colPixel, value)
		if err != nil {
			log.Fatalf("Renderer : Error in sending data to server. Error : %+v", err)
		}
		rowPixel, colPixel = updatePosition(rowPixel, colPixel, colSize)
		if rowPixel > rowSize {
			return nil
		}
	}
}

func (s *SimpleRenderer) send(rowPixel int, colPixel int, value byte) error {
	serverString := make([]string, 5)
	serverString[0] = "text"
	serverString[3] = "#000000"
	serverString[1] = strconv.Itoa(colPixel * col_char)
	serverString[2] = strconv.Itoa(rowPixel * row_char)
	serverString[4] = string(value)
	err := s.client.Send([]byte(strings.Join(serverString[:], ",") + "\n"))
	if err != nil {
		return err
	}
	return nil
}

func updatePosition(rowPixel, colPixel, colSize int) (int, int) {
	colPixel = colPixel + 1
	if colPixel > colSize {
		return rowPixel + 1, 0
	} else {
		return rowPixel, colPixel
	}
}
