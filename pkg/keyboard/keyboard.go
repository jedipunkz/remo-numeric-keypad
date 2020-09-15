package mykeyboard

import (
	"gobot.io/x/gobot/platforms/keyboard"
)

type Keyboard struct {
	KeyDriver *keyboard.Driver
}

func NewKeyboard() *Keyboard {
	k := new(Keyboard)
	k.KeyDriver = keyboard.NewDriver()
	return k
}
