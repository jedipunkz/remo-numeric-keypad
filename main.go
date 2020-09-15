package main

import (
	"context"
	"fmt"
	"log"
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
	key1 aplSig
	key2 aplSig
	key3 aplSig
	key4 aplSig
	key5 aplSig
}

type aplSig struct {
	apl string
	sig string
}

func newKeys() *keys {
	k := new(keys)
	k.key1.apl = viper.GetString("key1.apl")
	k.key2.apl = viper.GetString("key2.apl")
	k.key3.apl = viper.GetString("key3.apl")
	k.key4.apl = viper.GetString("key4.apl")
	k.key5.apl = viper.GetString("key5.apl")
	k.key1.sig = viper.GetString("key1.sig")
	k.key2.sig = viper.GetString("key2.sig")
	k.key3.sig = viper.GetString("key3.sig")
	k.key4.sig = viper.GetString("key4.sig")
	k.key5.sig = viper.GetString("key5.sig")
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
			key := data.(keyboard.KeyEvent)

			if key.Key == keyboard.One {
				fmt.Println(k.key1.apl)
				fmt.Println(k.key1.sig)
				if err := r.SendSignalByAplSig(k.key1.apl, k.key1.sig, ctx); err != nil {
					log.Fatal(err)
				}
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
