package model


type UrlClinetrsp struct {
  ClinetIP string `json:"ip"`
  Domain string `json:domain"`
  PingStaus bool `json:status`
}