package time

import (
	"plugins"
	"slack"
	"time"
)

type plugin struct {
}

func New() plugins.Plugin {
	return &plugin{}
}

func (p *plugin) Do(slash *slack.SlashCommandService, req slack.SlashCommandRequest) (responseMessage string) {
	res, err := slash.Reply(req, slack.NewInChannelMessage(
		time.Now().Format(time.RFC3339),
	))

	if err != nil {
		return err.Error()
	}
	defer res.Body.Close()

	return ""
}
