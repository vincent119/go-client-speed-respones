package model


type UrlClinetrsp struct {
  ClinetIP string `json:"ip"`
  Domain string `json:"domain"`
  PingStaus string `json:"status"`
  PingMax int `json:"max"`
  PingMin int `json:"min"`
  PingAvg int `json:"avg"`
}