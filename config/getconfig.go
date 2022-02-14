package config

import (
	"log"
	"os"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)
/*

  test

*/

var (
	cfg = pflag.StringP("Config","c","","./config/")
)

type AppConf struct {
	Port string `mapstructure:"port"`
	LogFile string `mapstructure:"logfile"`
}

type Config struct {
	App AppConf `mapstructure:"app"`
}

type LogrusConfig struct {
	Path         string `ini:"path"`
	Level        string `ini:"level"`
	Formatter    string `ini:"formatter"`
	OutputType   string `ini:"output_type"`
	ReportCaller bool   `ini:"report_caller"`
	Suffix       string `ini:"suffix_format"`
}

var cf *viper.Viper
var Conf Config
func Init(env string) {
	log.Print("This is the environment: ", env)
	pflag.Parse()
	cf = viper.New()
	confPath ,_ := os.Getwd()
	if *cfg != ""{
	  cf.SetConfigFile(*cfg)
	} else {
	  cf.AddConfigPath(".")
	  cf.AddConfigPath(confPath)
	}
	cf.SetConfigType("yaml")
	cf.AutomaticEnv()
	err := cf.ReadInConfig()
	if err != nil {
	  panic(err)
	}
	err = cf.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("Server  port : %s\n",Conf.App.Port)
}

func GetServerPort()(x string) {
  return Conf.App.Port
}

func GetServerLogPath()(cx string){
  return Conf.App.LogFile
}





