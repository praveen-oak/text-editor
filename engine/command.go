package engine

import (
	textdata "text-editor/data"
	"text-editor/renderer"
)

const (
	Mousedown CommandType = "mousedown"
	Mouseup               = "mouseup"
	Keydown               = "keydown"
	Keyup                 = "keyup"
	Resize                = "resize"
	MouseMove             = "mousemove"
)

type Engine interface {
	Run(command Command) error
}

type SimpleEngine struct {
	dataStore textdata.DataStore
	renderer  renderer.Renderer
}

type CommandType string

type Command struct {
	commandType CommandType
	row         int
	col         int
	value       string
}

func NewCommand(commandType CommandType, row int, col int, value string) *Command {
	return &Command{
		commandType: commandType,
		row:         row,
		col:         col,
		value:       value,
	}
}

func NewSimpleEngine(dataStore textdata.DataStore, renderer renderer.Renderer) *SimpleEngine {
	return &SimpleEngine{
		dataStore: dataStore,
		renderer:  renderer,
	}
}

func (s *SimpleEngine) Run(command *Command) error {
	switch command.commandType {
	case Resize:
		s.dataStore.UpdateWindow(command.col, command.row)
		s.renderer.RefreshScreen(s.dataStore)
	case Keydown:
		s.processKeyDown(command.value)
		s.renderer.RefreshScreen(s.dataStore)
	}
	return nil
}

func (s *SimpleEngine) processKeyDown(value string) error {
	cCol,cRow := s.dataStore.GetCursor()
	if value == "Return" {
		s.dataStore.AppendRow()
		s.dataStore.SetCursor(0, cRow+1)
	} else {
		s.dataStore.AppendChar(cRow, value[0])
		s.dataStore.SetCursor(cCol+1, cRow)
	}
	return nil
}
