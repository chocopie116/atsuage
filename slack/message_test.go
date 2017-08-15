package slack

import (
	"testing"
)

func TestCreateBotStatement_withoutTriggerWord(t *testing.T) {
	m := ChatMessage{
		Token: "sometoken",
		TeamId: "1234",
		TeamDomain: "test",
		ChannelId: "1234",
		ChannelName: "general",
		Timestamp: "123456",
		UserId: "123",
		UserName: "chocopie116",
		Text: "somekeyword atsuage",
		TriggerWord: "somekeyword",
	}
	s := m.createBotStatement()

	if s.Text != "somekeyword atsuage" {
		t.Errorf("want astuage but %+v", s)
	}
}

