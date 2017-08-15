package slack

type ChatMessage struct {
	Token string `json:"token"`
	TeamId string `json:"team_id"`
	TeamDomain string `json:"team_domain"`
	ChannelId string `json:"channel_id"`
	ChannelName string `json:"channel_name"`
	Timestamp string `json:"timestamp"`
	UserId string `json:"user_id"`
	UserName string `json:"user_name"`
	Text string `json:"text"`
	TriggerWord string `json:"trigger_word"`
}


type BotStatement struct {
	Text string
}

func (m ChatMessage) createBotStatement() BotStatement {
	return BotStatement{Text: m.Text}
}