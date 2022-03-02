package model


type ClientPingStatus struct {
  ClinetIP string `json:"clientIp"`
  Domain string `json:"domain"`
  Remote string `json:"Remote"`
  PingMax float32  `json:"max"`
  PingMin float32  `json:"min"`
  PingAvg float32  `json:"avg"`
  PostStatus int `json:"post"`
  PackLoss int `json:"paxkloss"`
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
  Domain string `json:"domain"`
  Remote string `json:"Remote"`
  Connect string `json:"Connect"`
  TimeStamp string `json:"timestamp"`
  ConnStatus string `json:"status"`
}