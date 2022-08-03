package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type DBConfig struct {
	Username string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
	DBName   string `mapstructure:"POSTGRES_DB"`
}

func (conf *DBConfig) LoadConfig(fileName string) {
	viper.SetConfigFile(fileName)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		logrus.Panic("Read file error: ", err)
	}
	if err := viper.Unmarshal(&conf); err != nil {
		logrus.Panic("Parse file error: ", err)
	}
}
