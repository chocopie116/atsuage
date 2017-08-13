package slack

type Bot interface {
	Parse() (BotResponse, error)
}

func NewBot(commands [] BotCmd){
	//TODO Botのcmdを引数で受け取るようにする
	return SlackBotImpl{commands}
}

type BotImpl struct {
	//TODO
	commands [] BotCmd
}

type BotResponse struct {
}

type BotCmd interface {
	//TODO messageを受ける
	match()(bool, error)
	//TODO messageを受ける
	action()(BotResponse, error)
}

//TODO messageを受け取ってResponseを返す
func (b SlackBotImpl) Parse() (BotResponse, error){
	return BotResponse{}, nil
}
