package main

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
