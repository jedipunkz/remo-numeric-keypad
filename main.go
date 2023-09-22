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

type KeyConfig struct {
	Appliance string `mapstructure:"appliance"`
	Signal    string `mapstructure:"signal"`
}

type Config struct {
	Token    string               `mapstructure:"token"`
	Keys     map[string]KeyConfig `mapstructure:"keys"`
	Plus     KeyConfig            `mapstructure:"plus"`
	Slash    KeyConfig            `mapstructure:"slash"`
	Asterisk KeyConfig            `mapstructure:"asterisk"`
	Hyphen   KeyConfig            `mapstructure:"hyphen"`
	Esc      KeyConfig            `mapstructure:"esc"`
	Dot      KeyConfig            `mapstructure:"dot"`
}

func newKeys() *Config {
	var myConfig Config
	if err := viper.Unmarshal(&myConfig); err != nil {
		fmt.Println("Unable to decode into struct,", err)
	}

	return &myConfig
}

func init() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	viper.AddConfigPath(home)
	viper.SetConfigName(confFile)
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file: ", viper.ConfigFileUsed())
	}

	var myConfig Config
	if err := viper.Unmarshal(&myConfig); err != nil {
		fmt.Println("Unable to decode into struct,", err)
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
