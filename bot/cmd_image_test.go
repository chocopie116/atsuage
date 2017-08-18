package bot

import (
	"testing"
)

func TestMatch_Ok(t *testing.T) {
	c := ImageCmd{}
	st := BotStatement{Text: "jpi waiwai"}
	is := c.Match(st)

	if is != true {
		t.Errorf("want true but got %+v", is)
	}
}

func TestMatch_Ng_Only_ForwardMatch(t *testing.T) {
	c := ImageCmd{}
	st := BotStatement{Text: "uhi jpi waiwai"}
	is := c.Match(st)

	if is != false {
		t.Errorf("want false but got %+v", is)
	}
}

func TestMatch_Ng_Need_Space_Between_Options(t *testing.T) {
	c := ImageCmd{}
	st := BotStatement{Text: "jpiwaiwai"}
	is := c.Match(st)

	if is != false {
		t.Errorf("want false but got %+v", is)
	}
}
