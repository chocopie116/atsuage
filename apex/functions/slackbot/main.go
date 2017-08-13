package main

import (
	"github.com/apex/go-apex"
	"log"
	"encoding/json"
	"github.com/chocopie116/atsuage/slack"
)

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		var m slack.ChatMessage

		log.Printf("original %s\n", event)
		if err := json.Unmarshal(event, &m); err != nil {
			return nil, err
		}
		//log.Printf("%s\n", m)

		return map[string]string{"text": m.Text}, nil
	})
}
