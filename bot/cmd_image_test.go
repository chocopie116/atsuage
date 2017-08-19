package bot

import (
	"testing"

	"github.com/chocopie116/atsuage/google"
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

type MockImageClient struct {
}

func(m MockImageClient) Search (q string) (*google.ImageResponse, error) {
	return &google.ImageResponse{Url: "https://example.com/static/img/test.png"}, nil
}

func TestMatch_Ok_SearchQuery(t *testing.T) {
	c := ImageCmd{MockImageClient{}}
	st := BotStatement{Text: "jpi majide"}
	r, err := c.Action(st)

	if err!= nil {
		t.Errorf("want nil but got %+v", err)
	}

	if r.Text != "https://example.com/static/img/test.png" {
		t.Errorf("want img url but got %+v", r.Text)
	}
}

