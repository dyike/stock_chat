package job

import (
    "time"
    "log"
    "github.com/robfig/cron"
    "wechat/stock"
)

var JobService = cron.New()
var feedList = []string{"US_BILI", "US_APLE", "hk01810"}

func GetStockDetails() {
    JobService.AddFunc("*/10 * * * *", func() {
        log.Println("Run get feed data")
        stock.GetFeed(feedList)
    })

    JobService.Start()
    t1 := time.NewTimer(time.Second * 10)
    for {
        select {
        case <-t1.C:
            t1.Reset(time.Second * 10)
        }
    }
}