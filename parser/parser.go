package parser

import (
	"fmt"
	"strconv"
	"strings"
	"text-editor/netclient"
)

const(
	Mousedown CommandType = "mousedown"
	Mouseup = "mouseup"
	Keydown = "keydown"
	Keyup = "keyup"
	Resize = "resize"
	MouseMove = "mousemove"
)

type CommandType string

type Command struct{
	commandType CommandType
	row int
	col int
	char byte
}
type Parser interface {
	Parse(commandString string) (Command, error)
}

type SimpleParser struct{
	client *netclient.Client
}

func NewSimpleParser() SimpleParser {
	return SimpleParser{}
}


func (s *SimpleParser) Parse(input []byte) (*Command, error) {
	commandString := string(input)
	sArr := strings.Split(commandString, ":")
	if len(sArr) == 0 {
		return nil, fmt.Errorf("Unable to parse command %s", commandString)
	}
	t := CommandType(sArr[0])
	switch t {
	case Mousedown, Mouseup, MouseMove:
		pos := strings.Split(sArr[1], ",")
		row, err := strconv.Atoi(pos[1])
		if err != nil {
			return nil, fmt.Errorf("Unable to parse row %s", commandString)
		}
		col, err := strconv.Atoi(pos[0])
		if err != nil {
			return nil, fmt.Errorf("Unable to parse col %s", commandString)
		}
		return &Command{
		commandType: t,
		row: row,
		col: col,
	}, nil
	case Keydown, Keyup:
		s := strings.Trim(sArr[1], "\t \n")
		b := byte([]rune(s)[0])
		return &Command{
			commandType: t,
			char:b,
		}, nil
	case Resize:
		size := strings.Split(sArr[1], ",")
		row, err := strconv.Atoi(size[1])
		if err != nil {
			return nil, fmt.Errorf("Unable to parse row %s", commandString)
		}
		col, err := strconv.Atoi(size[0])
		if err != nil {
			return nil, fmt.Errorf("Unable to parse col %s", commandString)
		}
		return &Command{
			commandType: t,
			row: row,
			col: col,
		}, nil
	default:
		return nil, fmt.Errorf("Command type not recognized. String : %s", commandString)
	}
}


