package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/cli"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const (
	confFile = ".remo-numpad"
	version  = "0.0.1"
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
	hyphen   aplSig
	asterisk aplSig
	plus     aplSig
	slash    aplSig
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
	k.hyphen.apl = viper.GetString("keys.hyphen.appliance")
	k.asterisk.apl = viper.GetString("keys.asterisk.appliance")
	k.plus.apl = viper.GetString("keys.plus.appliance")
	k.slash.apl = viper.GetString("keys.slash.appliance")
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
	k.hyphen.sig = viper.GetString("keys.hyphen.signal")
	k.asterisk.sig = viper.GetString("keys.asterisk.signal")
	k.plus.sig = viper.GetString("keys.plus.signal")
	k.slash.sig = viper.GetString("keys.slash.signal")
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
	c := cli.NewCLI("remo-numpad", version)

	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"server": func() (cli.Command, error) {
			return &serverCommand{}, nil
		},
		"scan": func() (cli.Command, error) {
			return &scanCommand{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitStatus)
}
