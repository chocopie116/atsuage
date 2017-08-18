package bot

import (
	"testing"
	"github.com/chocopie116/atsuage/google"
	"encoding/json"
	"log"
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

type MockGoogleImageClient struct {
}

func(m MockGoogleImageClient) Search (q string) (*google.GoogleImageSearchResponse, error) {
	b := []byte("{ 'items': [ { 'pagemap': { 'cse_thumbnail': [ { 'src': 'https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcSEnXvseTrR-L9bKZpR6P5SKP76QCVreYou4goWjqbZQpPXmivc_i5YidE'} ] } }]}")
	var g google.GoogleImageSearchResponse
	json.Unmarshal(b, &g)
	log.Print(g)

	return &g, nil
}

func TestMatch_Ok_SearchQuery(t *testing.T) {
	c := ImageCmd{google.GoogleImageClient{}}
	st := BotStatement{Text: "jpi majide"}
	_, err := c.Action(st)

	if err!= nil {
		t.Errorf("want nil but got %+v", err)
	}
}

