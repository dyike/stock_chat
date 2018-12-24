package main

import (
	"fmt"
	// "wechat/stock"
	// "wechat/job"
	// "wechat/chat"
	"wechat/config"
	"flag"
)


func main() {
	var confFile string
	flag.StringVar(&confFile, "conf", "./config.json", "config file name")
	flag.Parse()

	conf, err := config.LoadConfig(confFile)
	if err != nil {
		return
	}
	fmt.Println(conf)
	// cc := chat.New(conf.Wechat.WorkId, conf.Wechat.Secret, conf.Wechat.AgentId)
// 	md := make(map[string]string)
// 	md["content"]  = `
// 您的会议室已经预定，稍后会同步到邮箱
// >**事项详情**
// >事　项：<font color="info">开会</font>
// >组织者：@miglioguan
// >参与者：@miglioguan、@kunliu、@jamdeezhou、@kanexiong、@kisonwang
// >
// >会议室：<font color="info">广州TIT 1楼 301</font>
// >日　期：<font color="warning">2018年5月18日</font>
// >时　间：<font color="comment">上午9:00-11:00</font>
// >
// >请准时参加会议。
// >
// >如需修改会议信息，请点击：[修改会议信息](https://work.weixin.qq.com)
// `
//  	data := map[string]interface{}{
//         "touser":  "@all",
//         "msgtype": "markdown",
//         // "agentid" : 1000003,
//         "markdown": md,
//     }


    // cc.SendMessage(data)
	// fmt.Println(data)

	// getStockDetail()
	// var feedList = []string{"US_BILI", "US_APLE", "hk01810"}
	// res, err := stock.GetFeed(feedList)
	// if err == nil {
	// 	for _, v := range res {
	// 		fmt.Println(v.Name)
	// 	}
	// }

	// job.GetStockDetails()
}
