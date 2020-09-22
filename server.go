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
				r.SendSignalByAplSig(k.key0.apl, k.key0.sig, ctx)
			case keyboard.One:
				r.SendSignalByAplSig(k.key1.apl, k.key1.sig, ctx)
			case keyboard.Two:
				r.SendSignalByAplSig(k.key2.apl, k.key2.sig, ctx)
			case keyboard.Three:
				r.SendSignalByAplSig(k.key3.apl, k.key3.sig, ctx)
			case keyboard.Four:
				r.SendSignalByAplSig(k.key4.apl, k.key4.sig, ctx)
			case keyboard.Five:
				r.SendSignalByAplSig(k.key5.apl, k.key5.sig, ctx)
			case keyboard.Six:
				r.SendSignalByAplSig(k.key6.apl, k.key6.sig, ctx)
			case keyboard.Seven:
				r.SendSignalByAplSig(k.key7.apl, k.key7.sig, ctx)
			case keyboard.Eight:
				r.SendSignalByAplSig(k.key8.apl, k.key8.sig, ctx)
			case keyboard.Nine:
				r.SendSignalByAplSig(k.key9.apl, k.key9.sig, ctx)
			case keyboard.Escape:
				r.SendSignalByAplSig(k.esc.apl, k.esc.sig, ctx)
			case keyboard.Hyphen:
				r.SendSignalByAplSig(k.hyphen.apl, k.hyphen.sig, ctx)
			case keyboard.Asterisk:
				r.SendSignalByAplSig(k.asterisk.apl, k.asterisk.sig, ctx)
			case keyboard.Plus:
				r.SendSignalByAplSig(k.plus.apl, k.plus.sig, ctx)
			case keyboard.Slash:
				r.SendSignalByAplSig(k.slash.apl, k.slash.sig, ctx)
			case keyboard.Dot:
				r.SendSignalByAplSig(k.dot.apl, k.dot.sig, ctx)
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
