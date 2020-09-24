package main

import (
	"plugin/tldr-plugin"
	"plugin/word-counter-plugin"
	"plugin/xp-plugin"
	"fmt"
	"github.com/wechaty/go-wechaty/wechaty"
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
	"github.com/wechaty/go-wechaty/wechaty/user"
	"log"
	"os"
	"os/signal"
	"strings"
)

func main() {
	var bot = wechaty.NewWechaty()
	wordCounterPlugin := word_counter_plugin.NewWordCounterPlugin()
	tldrPlugin := tldr_plugin.NewTLDRPlugin("太长不看")
	xpPlugin := xp_plugin.NewXPPlugin()

	var counter int = 0

	bot.OnScan(func(ctx *wechaty.Context, qrCode string, status schemas.ScanStatus, data string) {
		fmt.Printf("Scan QR Code to login: %v\nhttps://wechaty.github.io/qrcode/%s\n", status, qrCode)
	}).OnLogin(func(ctx *wechaty.Context, user *user.ContactSelf) {
		fmt.Printf("User %s logined\n", user.Name())
	}).OnLogout(func(ctx *wechaty.Context, user *user.ContactSelf, reason string) {
		fmt.Printf("User %s logouted: %s\n", user, reason)
	}).OnMessage(func(ctx *wechaty.Context, message *user.Message) {
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

	var quitSig = make(chan os.Signal)
	signal.Notify(quitSig, os.Interrupt, os.Kill)

	select {
	case <-quitSig:
		log.Fatal("exit.by.signal")
	}
}
