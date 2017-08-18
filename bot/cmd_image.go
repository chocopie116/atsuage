package bot

import "strings"

type ImageCmd struct {
}

func (i ImageCmd) Match(st BotStatement) (bool){
	return strings.HasPrefix(st.Text, "jpi ")
}

func (i ImageCmd) Action(st BotStatement) (BotResponse, error) {
	return BotResponse{Text: st.Text}, nil
}

