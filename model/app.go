package model


type ClientPingStatus struct {
  ClinetIP string `json:"ip"`
  Domain string `json:"domain"`
  PingStaus string `json:"status"`
  PingMax int `json:"max"`
  PingMin int `json:"min"`
  PingAvg int `json:"avg"`
  PackLoss int `json:"packloss"`
  TimeStamp string `json:"timestamp"`
}

type ClientDnsStatus struct {
  ClinetIP string `json:"clientIp"`
  Domain string `json:"domainName"`
  DomainResult string `json:"domainResult"`
  DoaminResp int`json:"domainresp"`
  TimeStamp string `json:"time"`
}

type ClientConnStatus struct {
  ClinetIP string `json:"clientIp"`
  Domain string `json:"domainName"`
  TimeStamp string `json:"time"`
  ConnStatus string `json:"status"`
}