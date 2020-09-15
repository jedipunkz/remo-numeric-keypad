package mykeyboard

import (
	"gobot.io/x/gobot/platforms/keyboard"
)

// MyKeyboard is struct for my own keyboard
type MyKeyboard struct {
	KeyDriver *keyboard.Driver
}

// NewMyKeyboard is contstructor for my own keyboard
func NewMyKeyboard() *MyKeyboard {
	k := new(MyKeyboard)
	k.KeyDriver = keyboard.NewDriver()
	return k
}
