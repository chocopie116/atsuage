package main

import (
	"log"
	"encoding/json"

	"github.com/apex/go-apex"
	"github.com/chocopie116/atsuage/slack"
	"github.com/chocopie116/atsuage/bot"
)

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		var m slack.ChatMessage

		log.Printf("original %s\n", event)
		if err := json.Unmarshal(event, &m); err != nil {
			return nil, err
		}

		b := initialize()

		rs, err := b.Parse(m)

		if err != nil {
			return nil, err
		}

		//log.Printf("%+v", rs)

		return map[string]string{"text": rs.Text}, nil
	})
}

func initialize() bot.Bot {
	commands := []bot.BotCmd{&bot.DefaultCmd{}}

	return bot.NewBot(commands)
}
