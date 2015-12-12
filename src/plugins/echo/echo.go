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

func (p *plugin) Do(slash *slack.SlashCommandService, req slack.SlashCommandRequest) (responseMessage string) {
	_, args := req.CmdArgs()
	return "Echo Message " + strings.Join(args, "")
}
