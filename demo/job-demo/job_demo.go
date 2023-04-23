package job_demo

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func TJob() *cron.Cron {
	c := cron.New(cron.WithSeconds())

	_, err := c.AddFunc("*/1 * * * * *", fmtF)
	if err != nil {
		panic("here is err" + err.Error())
	}
	c.Start()
	return c
}


func fmtF(){
	fmt.Println("here is test job ", time.Now().Unix())
}
