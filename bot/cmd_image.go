package bot

type ImageCmd struct {
}

func (i ImageCmd) Match(st BotStatement) (bool, error){
	return true, nil
}

func (i ImageCmd) Action(st BotStatement) (BotResponse, error) {

	return BotResponse{Text: st.Text}, nil
}

