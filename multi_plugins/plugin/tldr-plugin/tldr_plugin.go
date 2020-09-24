package tldr_plugin

import (
	"github.com/wechaty/go-wechaty/wechaty"
	"github.com/wechaty/go-wechaty/wechaty/user"
	"log"
)

func NewTLDRPlugin(replyMessage string) *wechaty.Plugin{
	p := wechaty.NewPlugin()
	p.OnMessage(func(context *wechaty.Context, message *user.Message) {
		count := context.GetData("WordCount").(int)
		if count > 50 {
			_,err := message.Say(replyMessage)
			if err != nil {
				log.Println(err)
				return
			}
			context.Abort()
		}
	})
	return p
}