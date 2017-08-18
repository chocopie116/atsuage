package bot

import (
	"strings"

	"github.com/chocopie116/atsuage/google"
)

type ImageCmd struct {
	GoogleClient google.GoogleImageClient
}

func (i ImageCmd) Match(st BotStatement) (bool){
	return strings.HasPrefix(st.Text, "jpi ")
}

func (i ImageCmd) Action(st BotStatement) (BotResponse, error) {
	return BotResponse{Text: st.Text}, nil
}

