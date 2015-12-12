package echo

import (
	"plugins"
	"slack"
	"strings"
)

type plugin struct {
}

func New() plugins.Plugin {
	return &plugin{}
}

func (p *plugin) Do(req slack.SlashCommandRequest) slack.SlashCommandMessage {
	_, args := req.CmdArgs()
	return slack.NewMessage(
		"Echo Message " + strings.Join(args, ""),
	)
}
