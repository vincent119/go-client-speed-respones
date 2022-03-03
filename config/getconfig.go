package config

import (
	//"log"
	"os"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg = pflag.StringP("Config","c","","./config/")
)

type AppConf struct {
	Port string `mapstructure:"port"`
	LogPath string `mapstructure:"logpath"`
	LogFile string `mapstructure:"logfile"`
	Ukey  string  `mapstructure:"ukey"`
	Slat string `mapstructure:"slat"`
}
type Url1 struct {
	LogName string `mapstructure:"logname"`
}
type RedisConfig struct {
	Port string `mapstructure:"port"`
	Auth string `mapstructure:"auth"`
	Host string `mapstructure:"ip"`
	TTL int `mapstructure:"ttl"`
}

type Config struct {
	App AppConf `mapstructure:"app"`
	Uri1 Url1 `mapstructure:"url1"`
	RedisCfg RedisConfig `mapstructure:"RedisConfig"`
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
func GetServerUkey() (x string){
	return Conf.App.Ukey
}
func GetUrl1LogFile() (x string){
  return Conf.Uri1.LogName
}
func RedisPort() (x string){
  return Conf.RedisCfg.Port
}
func RedisHost() (x string){
  return Conf.RedisCfg.Host
}
func RedisAuth() (x string){
	return Conf.RedisCfg.Auth
}
func RedisTtl() (x int){
	return Conf.RedisCfg.TTL
}



