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

func (p *plugin) Do(_ slack.SlashCommandRequest) slack.SlashCommandMessage {
	return slack.NewInChannelMessage(
		time.Now().Format(time.RFC3339),
	)
}
