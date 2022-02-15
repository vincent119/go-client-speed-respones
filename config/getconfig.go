package config

import (
	//"log"
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
	LogPath string `mapstructure:"logpath"`
	LogFile string `mapstructure:"logfile"`
}
type Url1 struct {
	LogName string `mapstructure:"logname"`
}

type Config struct {
	App AppConf `mapstructure:"app"`
	Uri1 Url1 `mapstructure:"url1"`
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
func Init(){
	//log.Print("This is the environment: ", env)
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
func GetServerLogPath() (x string) {
  return Conf.App.LogPath
}
func GetServerLogFile() (x string){
  return Conf.App.LogFile
}
func GetUrl1LogName() (x string){
  return Conf.Uri1.LogName
}



