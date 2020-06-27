package textdata

import (
	"fmt"
)

type DataStore interface {
	ReadChar(col, row int) (byte, error)
	UpdateChar(col, row int, c byte) error
	AppendChar(row int, c byte) error
	DeleteChar(column, row int) error
	AppendRow() error
	Reset() error
	GetWindow() (col, row int)
	UpdateWindow(col, row int) error
	SetCursor(col, row int ) error
	GetCursor() (col, row int)
}

type ArrayStore struct {
	lineArray    [][]byte
	rows         int
	windowRows   int
	windowCols   int
	cursorRow    int
	cursorColumn int
}

func NewArrayStore() DataStore {
	return &ArrayStore{
		lineArray: make([][]byte, 0),
		rows:      0,
		cursorColumn:      0,
		cursorRow:        0,
	}
}

func (a *ArrayStore) ReadChar(row, column int) (byte, error) {
	err := a.checkIfCharExists(row, column)
	if err != nil {
		return 0, err
	}
	return a.lineArray[row][column], nil
}

func (a *ArrayStore) UpdateChar(row, column int, c byte) error {
	err := a.checkIfCharExists(row, column)
	if err != nil {
		return err
	}
	a.lineArray[row][column] = c
	return nil
}

func (a *ArrayStore) DeleteChar(row, column int) error {
	err := a.checkIfCharExists(row, column)
	if err != nil {
		return err
	}
	if row == 0 && column == 0 {
		return nil
	}

	if len(a.lineArray[row]) == 0 {
		//delete row itself
		a.lineArray = append(a.lineArray[:row], a.lineArray[row+1:]...)
		a.rows--
		return nil
	} else {
		if column == 0 {
			//merge this row to previous row
			a.lineArray[row-1] = append(a.lineArray[row-1], a.lineArray[row]...)
			a.lineArray = append(a.lineArray[:row], a.lineArray[row+1:]...)
			a.rows--
			return nil
		} else {
			//just remove the character
			a.lineArray[row] = append(a.lineArray[row][:column], a.lineArray[row][column+1:]...)
			return nil
		}
	}
}

func (a *ArrayStore) AppendChar(row int, c byte) error {
	if row >= a.rows {
		return fmt.Errorf("arrayStore: row %d requested to be read does not exist", row)
	}
	a.lineArray[row] = append(a.lineArray[row], c)
	return nil
}

func (a *ArrayStore) AppendRow() error {
	a.rows++
	a.lineArray = append(a.lineArray, make([]byte, 0))
	return nil
}

func (a *ArrayStore) Reset() error {
	a.rows = 0
	a.lineArray = make([][]byte, 0)
	return nil
}

func (a *ArrayStore) GetWindow() (x, y int) {
	return a.windowRows, a.windowCols
}
func (a *ArrayStore) GetCursor() (column, row int) {
	return a.cursorColumn, a.cursorRow
}

func (a *ArrayStore) SetCursor(column, row int) error {
	a.cursorRow = row
	a.cursorColumn = column
	return nil
}

func (a *ArrayStore) UpdateWindow(x, y int) error {
	a.windowRows = x
	a.windowCols = y
	return nil
}

func (a *ArrayStore) checkIfCharExists(row, column int) error {
	if row >= a.rows {
		return fmt.Errorf("arrayStore: row %d requested to be read does not exist", row)
	}
	r := a.lineArray[row]

	if column >= len(r) {
		return fmt.Errorf("arrayStore: column %d of row %d requested to be read does not exist", column, row)
	}
	return nil
}
