package slack

import "fmt"

type Bot interface {
	Parse(ChatMessage) (BotResponse, error)
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
	Match(ChatMessage)(bool, error)
	Action(ChatMessage)(BotResponse, error)
}

func (b BotImpl) Parse(m ChatMessage) (BotResponse, error){
	var r BotResponse
	for _, c := range b.commands {
		matched, err := c.Match(ChatMessage{})
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