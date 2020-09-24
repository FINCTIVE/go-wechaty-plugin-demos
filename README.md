# go-wechaty 插件机制使用方式演示代码

*本仓库仅用于说明展示用途，非最终代码成果。*

详情请见 https://github.com/wechaty/summer-of-code/issues/9

# 插件 Demo 

该部分代码位于 word_counter 文件夹下。

统计限定小时内群成员发言词数。仅统计文字类消息。[插件源码](https://github.com/FINCTIVE/wordcounter)

运行方法

```
export WECHATY_PUPPET_HOSTIE_TOKEN=<Your Token>
go run word_counter/word.go
```

# 多插件演示

该部分代码位于 multi_plugins 文件夹下。用于演示使用多个插件的聊天机器人代码。

- `word-counter-plugin` 统计字数。
- `xp-plugin` 当消息字数大于15，经验加1。
- `tldr-plugin` 如果消息字数大于50，回复“太长不看”。

主要代码 [demo_bot.go](https://github.com/FINCTIVE/go-wechaty-plugin-demos/blob/master/multi_plugins/demo_bot.go)

```go
var bot = wechaty.NewWechaty()

	wordCounterPlugin := word_counter_plugin.NewWordCounterPlugin()
	tldrPlugin := tldr_plugin.NewTLDRPlugin("太长不看")
	xpPlugin := xp_plugin.NewXPPlugin()

	var counter int = 0

	bot.OnMessage(func(ctx *wechaty.Context, message *user.Message) {
		// control plugins
		if strings.Contains(message.Text(), "通知") {
			ctx.DisableOnce(tldrPlugin)
			counter++
		}

		if counter == 3{
			tldrPlugin.SetEnable(false)
		}
	})

	bot.Use(wordCounterPlugin)
	bot.Use(tldrPlugin)
	bot.Use(xpPlugin)

	var err = bot.Start()
	if err != nil {
		panic(err)
	}
```