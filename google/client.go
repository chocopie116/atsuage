package google

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type ImageClient interface {
	Search(string) (*ImageResponse, error)
}

type googleCustomSearchApiResponse struct {
	//TODO Error区別したい
	//Error struct {
	//	Code   int `json:"code"`
	//	Errors []struct {
	//		Domain       string `json:"domain"`
	//		ExtendedHelp string `json:"extendedHelp"`
	//		Message      string `json:"message"`
	//		Reason       string `json:"reason"`
	//	} `json:"errors"`
	//	Message string `json:"message"`
	//} `json:"error"`
	Items []struct {
		Pagemap          struct {
			CseThumbnail []struct {
				Src    string `json:"src"`
			} `json:"cse_thumbnail"`
		} `json:"pagemap"`
	} `json:"items"`
}

type ImageResponse struct {
	Url string
}

func (r *ImageResponse) IsEmpty() (bool){
	return r.Url == ""
}

func (r *ImageResponse) FetchImageUrl() (string) {
	return r.Url
}

type ImageClientImpl struct {
}

func (c *ImageClientImpl) Search(q string) (*ImageResponse, error) {
	u := fmt.Sprintf("https://www.googleapis.com/customsearch/v1?key=%s&cx=%s&q=%s",
		os.Getenv("GOOGLE_IMAGE_API_KEY"),
		os.Getenv("GOOGLE_IMAGE_SEARCH_ENGINE_ID"),
		q)

	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var gr googleCustomSearchApiResponse
	if err := json.NewDecoder(res.Body).Decode(&gr); err != nil {
		return nil, err
	}

	//TODO API RateLimitの時と検索結果が0件の時を区別したい
	if len(gr.Items) == 0 {
		return &ImageResponse{}, nil
	}

	var ir ImageResponse
	//TODO 雑に検索結果のうちの10件を全部代入してランダムに1件選ぶようにしたい
	ir.Url = gr.Items[0].Pagemap.CseThumbnail[0].Src

	return &ir, nil
}
