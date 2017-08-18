package google

import (
	"testing"
	"log"
)
func TestSearchQuery(t *testing.T) {
	g := GoogleImageClient{}
	r, err := g.Search("Test")
	if err != nil {
		t.Errorf("want nil but got %+v", err)
	}

	if r.IsEmpty() == true {
		t.Errorf("want true but got %+v", r.IsEmpty())
	}

	log.Print(r.fetchImageSrc())
}
