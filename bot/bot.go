package bot

import (
	"fmt"
	"github.com/chocopie116/atsuage/slack"
)

type Bot interface {
	Parse(slack.ChatMessage) (BotResponse, error)
}

func NewBot(commands [] BotCmd) Bot{
	return BotImpl{commands}
}

type BotImpl struct {
	commands [] BotCmd
}

type BotResponse struct {
	Text string
}

type BotCmd interface {
	Match(slack.ChatMessage)(bool, error)
	Action(slack.ChatMessage)(BotResponse, error)
}

func (b BotImpl) Parse(m slack.ChatMessage) (BotResponse, error){
	var r BotResponse
	for _, c := range b.commands {
		matched, err := c.Match(slack.ChatMessage{})
		if err != nil {
			return r, err
		}

		if matched == false{
			continue
		}

		r, err := c.Action(m)
		if err != nil {
			return r, err
		}

		return r, nil
	}

	return r,fmt.Errorf("Nothing BodCmd matched. plz check injected BotCmd")
}