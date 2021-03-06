package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Command int

const (
	KeyNothing Command = iota
	KeySpace
)

type Input struct {
}

func NewInput() *Input {
	return &Input{}
}

func (i *Input) Update(keys []ebiten.Key) []ebiten.Key {
	return inpututil.AppendPressedKeys(keys)
}

func (i *Input) getCommands() []Command {
	var keys []ebiten.Key
	keys = i.Update(keys)
	var commands []Command
	for _, k := range keys {
		var command Command
		switch k {
		case ebiten.KeySpace:
			command = KeySpace
		}
		commands = append(commands, command)
	}
	return commands
}

func (i *Input) getVector() Vector {
	var keys []ebiten.Key
	keys = i.Update(keys)
	v := NewVector(0.0, 0.0)
	for _, k := range keys {
		switch k {
		case ebiten.KeyArrowLeft:
			v = v.Add(NewVector(-1.0, 0.0))
		case ebiten.KeyArrowRight:
			v = v.Add(NewVector(1.0, 0.0))
		}
	}
	return v
}
