package main

import (
	"log"
	"encoding/json"

	"github.com/apex/go-apex"
	"github.com/chocopie116/atsuage/bot"
	"bytes"
)

type slackMessage struct {
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

type slackResponse struct {
	Text string `json:"text`
}

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		log.Printf("original %s\n", event)

		st, err := factoryBotStatement(event)

		b := initialize()

		br, err := b.Parse(st)

		if err != nil {
			return nil, err
		}

		//log.Printf("%+v", rs)
		rs, err := formatSlackResponse(br)
		if err != nil {
			return rs, err
		}

		log.Printf("response %s\n", rs)

		return rs, nil
	})
}

func initialize() bot.Bot {
	commands := []bot.BotCmd{&bot.DefaultCmd{}}

	return bot.NewBot(commands)
}

func factoryBotStatement(event json.RawMessage) (bot.BotStatement, error) {
	var m slackMessage
	var st bot.BotStatement

	if err := json.Unmarshal(event, &m); err != nil {
		return st, err
	}

	return bot.BotStatement{Text: m.Text}, nil
}

func formatSlackResponse(br bot.BotResponse) (string, error) {
	sr := slackResponse{Text: br.Text}

	var buf bytes.Buffer
	b, err := json.Marshal(sr)
	if err != nil {
		return "", err
	}
	buf.Write(b)

	return buf.String(), nil
}
