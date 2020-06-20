package parser

import (
	"fmt"
	"strconv"
	"strings"
	"text-editor/engine"
	"text-editor/netclient"
)

type Parser interface {
	Parse(commandString string) (engine.Command, error)
}

type SimpleParser struct {
	client *netclient.Client
}

func NewSimpleParser() SimpleParser {
	return SimpleParser{}
}

func (s *SimpleParser) Parse(input []byte) (*engine.Command, error) {
	commandString := string(input)
	commandString = strings.Split(commandString, "\n")[0]
	sArr := strings.Split(commandString, ",")
	if len(sArr) == 0 {
		return nil, fmt.Errorf("Unable to parse command %s", commandString)
	}
	t := engine.CommandType(sArr[0])
	switch t {
	case engine.Mousedown, engine.Mouseup, engine.MouseMove:

		row, err := strconv.Atoi(sArr[2])
		if err != nil {
			return nil, fmt.Errorf("Unable to parse row %s", commandString)
		}
		col, err := strconv.Atoi(sArr[1])
		if err != nil {
			return nil, fmt.Errorf("Unable to parse col %s", commandString)
		}
		return engine.NewCommand(t, row, col, ""), nil
	case engine.Keydown, engine.Keyup:
		return engine.NewCommand(t, -1, -1, sArr[1]), nil
	case engine.Resize:
		row, err := strconv.Atoi(sArr[2])
		if err != nil {
			return nil, fmt.Errorf("Unable to parse row %s", commandString)
		}
		col, err := strconv.Atoi(sArr[1])
		if err != nil {
			return nil, fmt.Errorf("Unable to parse col %s", commandString)
		}
		return engine.NewCommand(t, row, col, ""), nil
	default:
		return nil, fmt.Errorf("Command type not recognized. String : %s", commandString)
	}
}
