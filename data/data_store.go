package textdata

import (
	"fmt"
)

type DataStore interface {
	ReadChar(x, y int) (byte, error)
	UpdateChar(x, y int, c byte) error
	AppendChar(x int, c byte) error
	DeleteChar(x, y int) error
	AppendRow() error
	InsertRow(x int) error
	Reset() error
}

type ArrayStore struct {
	LineArray [][]byte
	Rows      int
}

func NewArrayStore() DataStore {
	return &ArrayStore{
		LineArray: make([][]byte, 0),
		Rows:      0,
	}
}

func (a *ArrayStore) ReadChar(row, column int) (byte, error) {
	err := a.checkIfCharExists(row, column)
	if err != nil {
		return 0, err
	}
	return a.LineArray[row][column], nil
}

func (a *ArrayStore) UpdateChar(row, column int, c byte) error {
	err := a.checkIfCharExists(row, column)
	if err != nil {
		return err
	}
	a.LineArray[row][column] = c
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

	if len(a.LineArray[row]) == 0 {
		//delete row itself
		a.LineArray = append(a.LineArray[:row], a.LineArray[row+1:]...)
		a.Rows--
		return nil
	} else {
		if column == 0 {
			//merge this row to previous row
			a.LineArray[row-1] = append(a.LineArray[row-1], a.LineArray[row]...)
			a.LineArray = append(a.LineArray[:row], a.LineArray[row+1:]...)
			a.Rows--
			return nil
		} else {
			//just remove the character
			a.LineArray[row] = append(a.LineArray[row][:column], a.LineArray[row][column+1:]...)
			return nil
		}
	}
}

func (a *ArrayStore) AppendChar(row int, c byte) error {
	if row >= a.Rows {
		return fmt.Errorf("arrayStore: row %d requested to be read does not exist", row)
	}
	a.LineArray[row] = append(a.LineArray[row], c)
	return nil
}

func (a *ArrayStore) AppendRow() error {
	a.Rows++
	a.LineArray = append(a.LineArray, make([]byte, 0))
	return nil
}

func (a *ArrayStore) InsertRow(row int) error {
	a.Rows++
	a.LineArray = append(a.LineArray[:row], a.LineArray[row+1:]...)
	return nil
}

func (a *ArrayStore) checkIfCharExists(row, column int) error {
	if row >= a.Rows {
		return fmt.Errorf("arrayStore: row %d requested to be read does not exist", row)
	}
	r := a.LineArray[row]

	if column >= len(r) {
		return fmt.Errorf("arrayStore: column %d of row %d requested to be read does not exist", column, row)
	}
	return nil
}

func (a *ArrayStore) Reset() error {
	a.Rows = 0
	a.LineArray = make([][]byte, 0)
	return nil
}
