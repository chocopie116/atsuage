package main

import (
	"log"
	"encoding/json"

	"github.com/apex/go-apex"
	"github.com/chocopie116/atsuage/bot"
)

type chatMessage struct {
	Token string `json:"token"`
	TeamId string `json:"team_id"`
	TeamDomain string `json:"team_domain"`
	ChannelId string `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	Timestamp string `json:"timestamp"`
	UserId string `json:"user_id"`
	UserName string `json:"user_name"`
	Text string `json:"text"`
	TriggerWord string `json:"trigger_word"`
}

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		log.Printf("original %s\n", event)

		st, err := factoryBotStatement(event)

		b := initialize()

		rs, err := b.Parse(st)

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

func factoryBotStatement(event json.RawMessage) (bot.BotStatement, error) {
	var m chatMessage
	var st bot.BotStatement

	if err := json.Unmarshal(event, &m); err != nil {
		return st, err
	}

	return bot.BotStatement{Text: m.Text}, nil
}
