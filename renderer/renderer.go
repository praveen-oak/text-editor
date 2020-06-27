package renderer

import (
	"fmt"
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
	RefreshScreen(dataStore textdata.DataStore) error
}

type SimpleRenderer struct {
	client netclient.Client
}

func NewRenderer(client netclient.Client) Renderer {
	return &SimpleRenderer{
		client: client,
	}
}

func (s *SimpleRenderer) RefreshScreen(dataStore textdata.DataStore) error {
	rowPixel := 0
	colPixel := 0
	windowLength, windowHeight := dataStore.GetWindow()
	rowSize := windowLength / 8
	colSize := windowHeight / 8
	//cursorRendered := false

	serverString := make([]string, 5)
	serverString[0] = "text"
	serverString[3] = "#000000"

	cCol, cRow := dataStore.GetCursor()
	for {
		value, err := dataStore.ReadChar(rowPixel, colPixel)
		if err != nil {
			break
		}

		if rowPixel == cRow && colPixel == cCol {
			//cursorRendered = true
			//err = s.sendCursor(rowPixel, colPixel)
		} else {
			err = s.sendChar(rowPixel, colPixel, value)
		}
		if err != nil {
			log.Fatalf("Renderer : Error in sending data to server. Error : %+v", err)
		}
		rowPixel, colPixel = updatePosition(rowPixel, colPixel, colSize)
		if rowPixel > rowSize {
			break
		}
	}
	//if !cursorRendered {
	//	s.sendCursor(rowPixel, colPixel)
	//}
	return nil
}

func (s *SimpleRenderer) sendChar(rowPixel int, colPixel int, value byte) error {
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

func (s *SimpleRenderer) sendCursor(rowPixel int, colPixel int) error {
	serverString := make([]string, 6)
	serverString[0] = "rect"
	serverString[5] = "#000000"
	serverString[1] = strconv.Itoa(colPixel * col_char)
	serverString[2] = strconv.Itoa(rowPixel * row_char)
	serverString[3] = strconv.Itoa(col_char)
	serverString[4] = strconv.Itoa(row_char)
	fmt.Print("\n"+strings.Join(serverString[:], ",") + "\n")
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
