package bot

type ImageCmd struct {
}

func (d DefaultCmd) Match(st BotStatement) (bool, error){
	return true, nil
}

func (d DefaultCmd) Action(st BotStatement) (BotResponse, error) {
	return BotResponse{Text: st.Text}, nil
}
