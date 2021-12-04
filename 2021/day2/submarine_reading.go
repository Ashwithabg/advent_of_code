package main

type CommandName string

const (
	FORWARD CommandName = "forward"
	UP = "up"
	DOWN = "down"
)

type SubmarineReading struct {
	Name  CommandName
	Value int
}

