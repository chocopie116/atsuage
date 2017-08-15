package bot

import (
	"fmt"
)

type Bot interface {
	Parse(BotStatement) (BotResponse, error)
}

func NewBot(commands [] BotCmd) Bot{
	return BotImpl{commands}
}

type BotImpl struct {
	commands [] BotCmd
}

type BotStatement struct {
	Text string
}

type BotResponse struct {
	Text string
}

type BotCmd interface {
	Match(BotStatement)(bool, error)
	Action(BotStatement)(BotResponse, error)
}

func (b BotImpl) Parse(st BotStatement) (BotResponse, error){
	var r BotResponse
	for _, c := range b.commands {
		matched, err := c.Match(st)
		if err != nil {
			return r, err
		}

		if matched == false{
			continue
		}

		r, err := c.Action(st)
		if err != nil {
			return r, err
		}

		return r, nil
	}

	return r,fmt.Errorf("Nothing BodCmd matched. plz check injected BotCmd")
}

