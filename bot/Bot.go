package bot

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/proxy"

	tb "gopkg.in/tucnak/telebot.v2"
)

func Bot() {
	dialer, err := proxy.SOCKS5("tcp", "166.62.117.207:17935", &proxy.Auth{User: "user", Password: "pass"}, proxy.Direct)
	if err != nil {
		log.Fatal(err)
	}

	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}
	httpTransport.Dial = dialer.Dial

	b, err := tb.NewBot(tb.Settings{
		Token:  "********",
		URL:    "https://api.telegram.org",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
		Client: httpClient,
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "hello world")
		p := &tb.Photo{File: tb.FromDisk("/Users/di/Downloads/1.png")}
		b.Send(m.Sender, p)
		fmt.Println("Нам написали!!!")
	})

	b.Start()
}
