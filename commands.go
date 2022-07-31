package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Commands struct {
	Help  string `mapstructure:"HELP"`
	Add   string `mapstructure:"ADD"`
	All   string `mapstructure:"ALL"`
	Desc  string `mapstructure:"DESC"`
	Phone string `mapstructure:"PHONE"`
	Find  string `mapstructure:"FIND"`
	Show  string `mapstructure:"SHOW"`
	Exit  string `mapstructure:"EXIT"`
}

func NewCommandsConf(fileName string) *Commands {
	viper.SetConfigFile(fileName)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		logrus.Panic("Read file error", err)
	}
	var command Commands
	if err := viper.Unmarshal(&command); err != nil {
		logrus.Panic("Parse file error", err)
	}
	return &command
}
