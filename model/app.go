package model

type ClientPingStatus struct {
	ClinetIP   string  `json:"clientIp"`
	Domain     string  `json:"domain"`
	PingMax    float32 `json:"max"`
	PingMin    float32 `json:"min"`
	PingAvg    float32 `json:"avg"`
	PostStatus int     `json:"post"`
	PackLoss   int     `json:"paxkloss"`
	TimeStamp  string  `json:"timestamp"`
	Other
}

type ClientDnsStatus struct {
	ClinetIP     string  `json:"clientIp"`
	Domain       string  `json:"domainName"`
	DomainResult string  `json:"domainResult"`
	DoaminResp   float32 `json:"domainresp"`
	TimeStamp    string  `json:"timestamp"`
	Other
}

type Other struct{
	Host string `json:"host"`
	Category string `json:"category"`
}

type ClientConnStatus struct {
	ClinetIP   string `json:"clientIp"`
	Domain     string `json:"domain"`
	Connect    string `json:"Connect"`
	TimeStamp  string `json:"timestamp"`
	ConnStatus string `json:"status"`
	Other
}

type GenTokenString struct {
	ClinetIP string `json:"clientIp"`
	Uid     string `json:"uid"`
	Openid   string `json:"openid"`
}
