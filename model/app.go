package model


type AppConfigStr struct {
  Port string `mapstructure:"port" yaml:"port"`
  Logfile string `mapstructure:logfile yaml:"logfile"`
}