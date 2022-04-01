package cron

import (
	//"time"
	"fmt"

	"github.com/robfig/cron"
	//"github.com/robfig/cron/v3"
)

// func newWithSeconds() *cron.Cron {
// 	secondParser := cron.NewParser(cron.Second | cron.Minute |
// 		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
// 	return cron.New(cron.WithParser(secondParser), cron.WithChain())
// }

func Crontab(){
//	i := 0
  c := cron.New()
	spec := "0 */1 * * *"
	c.AddFunc(spec, func() { fmt.Println("OK") })
	// _ ,err := c.AddFunc(spec,func() {
	// 		i++
	// 		fmt.Println("cron running.....")
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	c.Start()
	select {}
}