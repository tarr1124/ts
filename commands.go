package main

import (
	"fmt"
	"github.com/timakin/ts/loader"
	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandAll,
//	commandBiz,
	commandHack,
}

var commandAll = cli.Command{
	Name:  "pop",
	Usage: "",
	Description: "Show today's news from major tech news sites, HN, PH, and subreddit of /programming.",
	Action: doAll,
}

//var commandBiz = cli.Command{
//	Name:  "biz",
//	Usage: "",
//	Description: `
//`,
//	Action: doBiz,
//}
var commandHack = cli.Command{
	Name:  "test",
	Usage: "",
	Description: `
`,
	Action: doHack,
}

func pp(str string) {
  fmt.Printf(str)
}

func displayRSSFeed(name string, uri string) {
	pp(name + "\n")
	loader.GetRSSFeed(uri)
}

func doAll(c *cli.Context) {
		pp("▁ ▂ ▄ ▅ ▆ ▇ █ тecнѕтacĸ █ ▇ ▆ ▅ ▄ ▂ ▁\n\n")
		ph := make(chan loader.ResultData)
		re := make(chan loader.ResultData)
		go loader.GetPHFeed(ph)
		go loader.GetRedditFeed(re)
		phres := <- ph
		reres := <- re
		var PHData loader.Feed = &phres
		var REData loader.Feed = &reres
		PHData.Display()
		REData.Display()
		displayRSSFeed("[HackerNews]", "https://news.ycombinator.com/rss")
		displayRSSFeed("[TechCrunch]", "http://feeds.feedburner.com/TechCrunch/")
		displayRSSFeed("[Mashable]", "http://feeds.mashable.com/Mashable")
		displayRSSFeed("[EchoJS]", "http://www.echojs.com/rss")
		displayRSSFeed("[RubyDaily]", "http://feeds.rubydaily.org/RubyDaily")
//		displayRSSFeed("[Hatena]", "http://b.hatena.ne.jp/search/tag?q=%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0&users=10&mode=rss")
}

//func doBiz(c *cli.Context) {
//}
func doHack(c *cli.Context) {
	displayRSSFeed("[EchoJS]", "http://www.echojs.com/rss")
	displayRSSFeed("[RubyDaily]", "http://feeds.rubydaily.org/RubyDaily")
	displayRSSFeed("[Hatena]", "http://b.hatena.ne.jp/search/tag?q=%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0&users=10&mode=rss")
}
