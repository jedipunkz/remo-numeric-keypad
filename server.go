package main

import (
	"context"

	mykeyboard "github.com/jedipunkz/mygobot/pkg/keyboard"
	myremo "github.com/jedipunkz/myremo/pkg/remo"

	"github.com/spf13/viper"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/keyboard"
)

type serverCommand struct{}

func (c *serverCommand) Help() string {
	return "Usage: remo-numpad server"
}

func (c *serverCommand) Synopsis() string {
	return "server boot"
}

func (c *serverCommand) Run(args []string) int {
	token := viper.GetString("token")
	r := myremo.NewRemo(token)

	mykeyboard := mykeyboard.NewMyKeyboard()
	k := newKeys()

	ctx := context.Background()

	work := func() {
		mykeyboard.KeyDriver.On(keyboard.Key, func(data interface{}) {
			pushedKey := data.(keyboard.KeyEvent)

			switch pushedKey.Key {
			case keyboard.Zero:
				r.SendSignalByAplSig(k.Keys["0"].Appliance, k.Keys["0"].Signal, ctx)
			case keyboard.One:
				r.SendSignalByAplSig(k.Keys["1"].Appliance, k.Keys["1"].Signal, ctx)
			case keyboard.Two:
				r.SendSignalByAplSig(k.Keys["2"].Appliance, k.Keys["2"].Signal, ctx)
			case keyboard.Three:
				r.SendSignalByAplSig(k.Keys["3"].Appliance, k.Keys["3"].Signal, ctx)
			case keyboard.Four:
				r.SendSignalByAplSig(k.Keys["4"].Appliance, k.Keys["4"].Signal, ctx)
			case keyboard.Five:
				r.SendSignalByAplSig(k.Keys["5"].Appliance, k.Keys["5"].Signal, ctx)
			case keyboard.Six:
				r.SendSignalByAplSig(k.Keys["6"].Appliance, k.Keys["6"].Signal, ctx)
			case keyboard.Seven:
				r.SendSignalByAplSig(k.Keys["7"].Appliance, k.Keys["7"].Signal, ctx)
			case keyboard.Eight:
				r.SendSignalByAplSig(k.Keys["8"].Appliance, k.Keys["8"].Signal, ctx)
			case keyboard.Nine:
				r.SendSignalByAplSig(k.Keys["9"].Appliance, k.Keys["9"].Signal, ctx)
			case keyboard.Escape:
				r.SendSignalByAplSig(k.Esc.Appliance, k.Esc.Signal, ctx)
			case keyboard.Hyphen:
				r.SendSignalByAplSig(k.Hyphen.Appliance, k.Hyphen.Signal, ctx)
			case keyboard.Asterisk:
				r.SendSignalByAplSig(k.Asterisk.Appliance, k.Asterisk.Signal, ctx)
			case keyboard.Plus:
				r.SendSignalByAplSig(k.Plus.Appliance, k.Plus.Signal, ctx)
			case keyboard.Slash:
				r.SendSignalByAplSig(k.Slash.Appliance, k.Slash.Signal, ctx)
			case keyboard.Dot:
				r.SendSignalByAplSig(k.Dot.Appliance, k.Dot.Signal, ctx)
			}
		})
	}

	robot := gobot.NewRobot("keyboardkbot",
		[]gobot.Connection{},
		[]gobot.Device{mykeyboard.KeyDriver},
		work,
	)

	robot.Start()
	return 0
}
