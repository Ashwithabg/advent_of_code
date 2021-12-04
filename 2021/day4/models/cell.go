package models

import "fmt"

type Cell struct {
	Value int
}

func (cell *Cell) markDone() {
	cell.Value = -1
}

func (cell *Cell) isMarked() bool{
	return cell.Value == -1
}

func (cell *Cell) display() {
	fmt.Print(" ", cell.Value)
}

func (cell *Cell) isValue(inputValue int) bool{
	return cell.Value == inputValue
}
