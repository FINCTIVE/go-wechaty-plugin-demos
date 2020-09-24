package xp_plugin

import (
	"github.com/wechaty/go-wechaty/wechaty"
	"github.com/wechaty/go-wechaty/wechaty/user"
	"log"
)

func NewXPPlugin() *wechaty.Plugin{
	p := wechaty.NewPlugin()
	var xp int = 0
	p.OnMessage(func(context *wechaty.Context, message *user.Message) {
		count := context.GetData("WordCount").(int)
		if count > 15 {
			_,err := message.Say("xp+1")
			if err != nil {
				log.Println(err)
				return
			}
			xp++
		}
	})
	return p
}