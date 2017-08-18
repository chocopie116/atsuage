package main

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/apex/go-apex"
	"github.com/chocopie116/atsuage/bot"
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
	Text string `json:"text"`
}

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		log.Printf("original %s\n", event)

		st, err := factoryBotStatement(event)
		if err != nil {
			return nil, err
		}

		b := initialize()
		br, err := b.Parse(st)
		if err != nil {
			return nil, err
		}

		rs, err := convertSlackResponse(br)
		if err != nil {
			return rs, err
		}
		log.Printf("response %+v\n", rs)

		return rs, nil
	})
}

func initialize() bot.Bot {
	commands := []bot.BotCmd{bot.DefaultCmd{}}

	return bot.NewBot(commands)
}

func factoryBotStatement(event json.RawMessage) (bot.BotStatement, error) {
	var m slackMessage
	var st bot.BotStatement
	var t string

	if err := json.Unmarshal(event, &m); err != nil {
		return st, err
	}

	if m.TriggerWord != "" {
		t = strings.Replace(m.Text, m.TriggerWord, "", 1)
	} else {
		t = m.Text
	}

	return bot.BotStatement{Text: t}, nil
}

func convertSlackResponse(br bot.BotResponse) (slackResponse, error) {
	return slackResponse{Text: br.Text}, nil
}
