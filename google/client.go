package google

import (
	"fmt"
	"os"
	"net/http"
	"encoding/json"
)

type GoogleImageClient struct {
}

type GoogleImageSearchResponse struct {
	Items []struct {
		Pagemap          struct {
			CseThumbnail []struct {
				Src    string `json:"src"`
			} `json:"cse_thumbnail"`
		} `json:"pagemap"`
	} `json:"items"`
}

func (c GoogleImageClient) Search(q string) (*GoogleImageSearchResponse, error) {
	u := fmt.Sprintf("https://www.googleapis.com/customsearch/v1?key=%s&cx=%s&q=%s",
		os.Getenv("GOOGLE_IMAGE_API_KEY"),
		os.Getenv("GOOGLE_IMAGE_SEARCH_ENGINE_ID"),
		q)

	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var r GoogleImageSearchResponse
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	return &r, nil
}

func (r GoogleImageSearchResponse) IsEmpty() (bool){
	return len(r.Items) == 0
}

func (r GoogleImageSearchResponse) fetchImageSrc() (string) {
	return r.Items[0].Pagemap.CseThumbnail[0].Src
}

