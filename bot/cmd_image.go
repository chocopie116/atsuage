package bot

import (
	"strings"

	"github.com/chocopie116/atsuage/google"
)

type ImageCmd struct {
	Client google.ImageClient
}

func (i ImageCmd) Match(st BotStatement) (bool){
	return strings.HasPrefix(st.Text, "jpi ")
}

func (i ImageCmd) Action(st BotStatement) (*BotResponse, error) {
	ir, err := i.Client.Search(st.Text)

	if err != nil {
		return nil, err
	}

	return &BotResponse{Text: ir.FetchImageUrl()}, nil
}
