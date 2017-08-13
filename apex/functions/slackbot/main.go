package main

import (
	"github.com/apex/go-apex"
	"log"
	"encoding/json"
)

type slackChatMessage struct {
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
		var m slackChatMessage

		log.Printf("original %s\n", event)
		if err := json.Unmarshal(event, &m); err != nil {
			return nil, err
		}
		//log.Printf("%s\n", m)

		return map[string]string{"text": m.Text}, nil
	})
}
