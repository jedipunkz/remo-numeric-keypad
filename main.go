package main

import (
	"context"
	"fmt"
	"os"

	mykeyboard "github.com/jedipunkz/mygobot/pkg/keyboard"
	myremo "github.com/jedipunkz/myremo/pkg/remo"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/keyboard"
)

const (
	confFile = ".remo-key"
)

type keys struct {
	key0     aplSig
	key1     aplSig
	key2     aplSig
	key3     aplSig
	key4     aplSig
	key5     aplSig
	key6     aplSig
	key7     aplSig
	key8     aplSig
	key9     aplSig
	esc      aplSig
	minus    aplSig
	asterisk aplSig
	plus     aplSig
	divide   aplSig
	dot      aplSig
}

type aplSig struct {
	apl string
	sig string
}

func newKeys() *keys {
	k := new(keys)
	k.key0.apl = viper.GetString("keys.0.appliance")
	k.key1.apl = viper.GetString("keys.1.appliance")
	k.key2.apl = viper.GetString("keys.2.appliance")
	k.key3.apl = viper.GetString("keys.3.appliance")
	k.key4.apl = viper.GetString("keys.4.appliance")
	k.key5.apl = viper.GetString("keys.5.appliance")
	k.key6.apl = viper.GetString("keys.6.appliance")
	k.key7.apl = viper.GetString("keys.7.appliance")
	k.key8.apl = viper.GetString("keys.8.appliance")
	k.key9.apl = viper.GetString("keys.9.appliance")
	k.esc.apl = viper.GetString("keys.esc.appliance")
	k.minus.apl = viper.GetString("keys.minus.appliance")
	k.asterisk.apl = viper.GetString("keys.asterisk.appliance")
	k.plus.apl = viper.GetString("keys.plus.appliance")
	k.divide.apl = viper.GetString("keys.divide.appliance")
	k.dot.apl = viper.GetString("keys.dot.appliance")
	k.key0.sig = viper.GetString("keys.0.signal")
	k.key1.sig = viper.GetString("keys.1.signal")
	k.key2.sig = viper.GetString("keys.2.signal")
	k.key3.sig = viper.GetString("keys.3.signal")
	k.key4.sig = viper.GetString("keys.4.signal")
	k.key5.sig = viper.GetString("keys.5.signal")
	k.key6.sig = viper.GetString("keys.6.signal")
	k.key7.sig = viper.GetString("keys.7.signal")
	k.key8.sig = viper.GetString("keys.8.signal")
	k.key9.sig = viper.GetString("keys.9.signal")
	k.esc.sig = viper.GetString("keys.esc.signal")
	k.minus.sig = viper.GetString("keys.minus.signal")
	k.asterisk.sig = viper.GetString("keys.asterisk.signal")
	k.plus.sig = viper.GetString("keys.plus.signal")
	k.divide.sig = viper.GetString("keys.divide.signal")
	k.dot.sig = viper.GetString("keys.dot.signal")
	return k
}

func init() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.AddConfigPath(home)
	viper.SetConfigName(confFile)

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file: ", viper.ConfigFileUsed())
	}
}

func main() {
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
			case keyboard.Minus:
				r.SendSignalByAplSig(k.minus.apl, k.minus.sig, ctx)
			case keyboard.Asterisk:
				r.SendSignalByAplSig(k.asterisk.apl, k.asterisk.sig, ctx)
			case keyboard.Plus:
				r.SendSignalByAplSig(k.plus.apl, k.plus.sig, ctx)
			case keyboard.Divide:
				r.SendSignalByAplSig(k.divide.apl, k.divide.sig, ctx)
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
}
