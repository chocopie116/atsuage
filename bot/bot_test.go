package bot

import (
	"testing"

	"github.com/chocopie116/atsuage/slack"
)

type TestCmd struct {
}

func (t TestCmd) Match (st BotStatement) (bool, error){
	return true, nil
}

func (t TestCmd) Action (st BotStatement) (BotResponse, error) {
	return BotResponse{Text: "TestCmd Response text is here"}, nil
}

func TestNewBot_OK(t *testing.T) {
	commands := []BotCmd{&TestCmd{}}
	b:= NewBot(commands)

	m := slack.ChatMessage{
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
		t.Errorf("want nil but got %+v", err)
	}

	if r.Text != "TestCmd Response text is here" {
		t.Errorf("BotResponse.Text is wrong. got BotResponse is %+v", r)
	}
}

type NotFoundTestCmd struct {
}

func (t NotFoundTestCmd) Match (st BotStatement) (bool, error){
	return false, nil
}

func (t NotFoundTestCmd) Action (st BotStatement) (BotResponse, error) {
	var r BotResponse
	return r, nil
}

func TestNewBot_nothing_matched(t *testing.T) {
	commands := []BotCmd{&NotFoundTestCmd{}}
	b:= NewBot(commands)

	m := slack.ChatMessage{
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
		t.Error("must return err but got nil")
	}
}

