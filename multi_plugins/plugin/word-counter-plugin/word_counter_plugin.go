package word_counter_plugin

import (
	"github.com/wechaty/go-wechaty/wechaty"
	"github.com/wechaty/go-wechaty/wechaty/user"
)

func NewWordCounterPlugin() *wechaty.Plugin{
	p := wechaty.NewPlugin()
	p.OnMessage(func(context *wechaty.Context, message *user.Message) {
		count := len(message.Text())
		context.SetData("WordCount", count)
	})
	return p
}