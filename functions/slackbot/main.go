package main

import (
	//"encoding/json"

	"github.com/apex/go-apex"
	"encoding/json"
)

type message struct {
	Value string `json:"message"`
}

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		//var m message

		//if err := json.Unmarshal(event, &m); err != nil {
		//	return nil, err
		//}

		return map[string]string{"text": "test"}, nil
	})
}
