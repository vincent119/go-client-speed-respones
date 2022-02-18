package model


type ClientPingStatus struct {
  ClinetIP string `json:"clientIp"`
  Domain string `json:"domain"`
  PingStaus string `json:"status"`
  PingMax float32  `json:"max"`
  PingMin float32  `json:"min"`
  PingAvg float32  `json:"avg"`
  PackLoss int `json:"packloss"`
  TimeStamp string `json:"timestamp"`
}

type ClientDnsStatus struct {
  ClinetIP string `json:"clientIp"`
  Domain string `json:"domainName"`
  DomainResult string `json:"domainResult"`
  DoaminResp float32 `json:"domainresp"`
  TimeStamp string `json:"timestamp"`
}

type ClientConnStatus struct {
  ClinetIP string `json:"clientIp"`
  Domain string `json:"domainName"`
  TimeStamp string `json:"timestamp"`
  ConnStatus string `json:"status"`
}