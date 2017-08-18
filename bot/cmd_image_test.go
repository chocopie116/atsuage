package bot

import (
	"testing"
)

func TestMatch_Ok(t *testing.T) {
	c := ImageCmd{}
	st := BotStatement{Text: "jpi waiwai"}
	is, err := c.Match(st)
	if err != nil {
			t.Errorf("BotResponse.Text is wrong. got BotResponse is %+v", err)
	}

	if is != true {
		t.Errorf("want true but got %+v", is)
	}
}

func TestMatch_Ng_Only_ForwardMatch(t *testing.T) {
	c := ImageCmd{}
	st := BotStatement{Text: "uhi jpi waiwai"}
	is, err := c.Match(st)
	if err != nil {
		t.Errorf("BotResponse.Text is wrong. got BotResponse is %+v", err)
	}

	if is != true {
		t.Errorf("want true but got %+v", is)
	}
}

