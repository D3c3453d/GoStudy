package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type DBConfig struct {
	Username string
	Password string
	DBName   string
}

func NewDBConf(fileName string) *DBConfig {
	viper.SetConfigFile(fileName)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		logrus.Panic("Read file error", err)
	}
	var conf DBConfig
	if err := viper.Unmarshal(&conf); err != nil {
		logrus.Panic("Parse file error", err)
	}
	return &conf
}
