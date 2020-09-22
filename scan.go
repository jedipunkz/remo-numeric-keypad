package main

import (
	"fmt"

	mykeyboard "github.com/jedipunkz/mygobot/pkg/keyboard"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/keyboard"
)

type scanCommand struct{}

func (c *scanCommand) Help() string {
	return "Usage: remo-numpad scan"
}

func (c *scanCommand) Synopsis() string {
	return "scan your key inputs"
}

func (c *scanCommand) Run(args []string) int {
	mykeyboard := mykeyboard.NewMyKeyboard()

	work := func() {
		mykeyboard.KeyDriver.On(keyboard.Key, func(data interface{}) {
			pushedKey := data.(keyboard.KeyEvent)
			fmt.Println("Your Key: ", pushedKey, " ", pushedKey.Char, " ", pushedKey.Bytes)
		})
	}

	robot := gobot.NewRobot("keyboardbot",
		[]gobot.Connection{},
		[]gobot.Device{mykeyboard.KeyDriver},
		work,
	)

	robot.Start()
	return 0
}
