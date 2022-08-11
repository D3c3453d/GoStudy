package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Commands Commands
	DBConfig DBConfig
}

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

type DBConfig struct {
	Username string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
	DBName   string `mapstructure:"POSTGRES_DB"`
	Host     string `mapstructure:"POSTGRES_HOST"`
	Port     string `mapstructure:"POSTGRES_PORT"`
}

func (conf *Commands) loadConfig(fileName string) {
	viper.SetConfigFile(fileName)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		logrus.Panic(err)
	}
	if err := viper.Unmarshal(&conf); err != nil {
		logrus.Panic(err)
	}
}

func (conf *DBConfig) loadConfig(fileName string) {
	viper.SetConfigFile(fileName)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		logrus.Panic(err)
	}
	if err := viper.Unmarshal(&conf); err != nil {
		logrus.Panic(err)
	}
}

func NewConfig() *Config {
	var command Commands
	var dbConf DBConfig
	command.loadConfig("./cfg/commands.env")
	dbConf.loadConfig("./cfg/db.env")
	Conf := Config{
		Commands: command,
		DBConfig: dbConf,
	}
	return &Conf
}
