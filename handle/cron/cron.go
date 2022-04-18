package cron

import (
	//"time"
	"fmt"

	//"github.com/robfig/cron"
	"github.com/robfig/cron/v3"
)

func newWithSeconds() *cron.Cron {
	secondParser := cron.NewParser(
		cron.SecondOptional | cron.Minute | cron.Hour |
			cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)
	return cron.New(cron.WithParser(secondParser), cron.WithChain())
}

func Crontab(){
	fmt.Println("[Cron] Starting...")
 
	c := newWithSeconds()
 
	spec := "* */5 * * * *" //每5秒执行一次
 
	c.AddFunc(spec, func() {
		fmt.Println("[Cron] Run models.CleanAllTag...")
		//models.CleanAllTag()
	})
 
	c.AddFunc(spec, func() {
		fmt.Println("[Cron] Run models.CleanAllArticle...")
		//models.CleanAllArticle()
	})
 
	c.Start()
 
	select {}
}