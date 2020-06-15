package renderer

import (
	"bufio"
	"log"
	"strconv"
	"strings"
	"text-editor/data"
)

const(
	row_char = 8
	col_char = 14
)

type Renderer interface {
	RefreshScreen(windowLength, windowHeight int) error
}

type SimpleRenderer struct {
	dataStore textdata.DataStore
	writer *bufio.Writer
}

func NewRenderer(dataStore textdata.DataStore, writer *bufio.Writer) Renderer {
	return &SimpleRenderer{
		dataStore: dataStore,
		writer:     writer,
	}
}

func (s *SimpleRenderer) RefreshScreen(windowLength, windowHeight int) error {
	rowPixel := 0
	colPixel := 0

	rowSize := windowLength/8
	colSize := windowHeight/8

	serverString := make([]string, 5)
	serverString[0] = "text"
	serverString[3] = "#000000"

	for {
		value, err := s.dataStore.ReadChar(rowPixel, colPixel)
		if err != nil {
			return nil
		}
		err = s.send(rowPixel, colPixel, value)
		if err != nil {
			log.Fatalf("Renderer : Error in sending data to server. Error : %+v", err)
		}
		rowPixel, colPixel = updatePosition(rowPixel, colPixel, rowSize)
		if colPixel > colSize {
			return nil
		}
	}
}

func (s *SimpleRenderer) send(rowPixel int, colPixel int, value byte) error {
	serverString := make([]string, 5)
	serverString[0] = "text"
	serverString[3] = "#000000"
	serverString[1] = strconv.Itoa(rowPixel*row_char)
	serverString[2] = strconv.Itoa(colPixel*col_char)
	serverString[4] = string(value)
	_, err := s.writer.Write([]byte(strings.Join(serverString[:], ",") + "\n"))
	if err != nil {
		return err
	}

	return nil
}

func updatePosition(rowPixel, colPixel, rowSize int) (int, int) {
	rowPixel = rowPixel + 1
	if rowPixel > rowSize {
		return 0, colPixel+1
	} else {
		return rowPixel, colPixel
	}
}