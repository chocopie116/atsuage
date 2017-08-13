package slack

import (
	"testing"
)

type TestCmd struct {
}

func (t TestCmd) Match (message ChatMessage) (bool, error){
	return true, nil
}

func (t TestCmd) Action (message ChatMessage) (BotResponse, error) {
	return BotResponse{Text: "TestCmd Response text is here"}, nil
}

func TestNewBot(t *testing.T) {
	commands := []BotCmd{&TestCmd{}}
	b:= NewBot(commands)

	m := ChatMessage{
		Token: "sometoken",
		TeamId: "1234",
		TeamDomain: "test",
		ChannelId: "1234",
		ChannelName: "general",
		Timestamp: "123456",
		UserId: "123",
		UserName: "chocopie116",
		Text: "atsuage",
		TriggerWord: "command",
	}
	r, err := b.Parse(m)

	if err != nil {
		t.Fatal("majide")
	}

	if r.Text != "TestCmd Response text is here" {
		t.Fatal("BodResponse.Text is wrong")
	}
}

type NotFoundTestCmd struct {
}

func (t NotFoundTestCmd) Match (message ChatMessage) (bool, error){
	return false, nil
}

func (t NotFoundTestCmd) Action (message ChatMessage) (BotResponse, error) {
	var r BotResponse
	return r, nil
}

func TestNewBot_nothing_matched(t *testing.T) {
	commands := []BotCmd{&NotFoundTestCmd{}}
	b:= NewBot(commands)

	m := ChatMessage{
		Token: "sometoken",
		TeamId: "1234",
		TeamDomain: "test",
		ChannelId: "1234",
		ChannelName: "general",
		Timestamp: "123456",
		UserId: "123",
		UserName: "chocopie116",
		Text: "atsuage",
		TriggerWord: "command",
	}
	_, err := b.Parse(m)

	if err == nil {
		t.Fatal("must return err because NotFoundTestCmd is never matched")
	}
}
