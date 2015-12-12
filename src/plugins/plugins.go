package plugins

import (
	"fmt"
	"slack"
)

type command string

type SlashCommand interface {
	AddPlugin(cmd string, plugin Plugin)
	Execute(slash *slack.SlashCommandService, req slack.SlashCommandRequest) string
}

type slashCommandContext struct {
	Plugins map[command]Plugin
}

func (c *slashCommandContext) AddPlugin(cmd string, plugin Plugin) {
	c.Plugins[command(cmd)] = plugin
}

func (c *slashCommandContext) Execute(slash *slack.SlashCommandService, req slack.SlashCommandRequest) string {
	cmd, _ := req.CmdArgs()
	p, ok := c.Plugins[command(cmd)]
	if !ok {
		return fmt.Sprintf("%s command not found", string(cmd))
	}
	return p.Do(slash, req)
}

func New() SlashCommand {
	return &slashCommandContext{
		Plugins: map[command]Plugin{},
	}
}

type Plugin interface {
	Do(slash *slack.SlashCommandService, req slack.SlashCommandRequest) (responseMessage string)
}
