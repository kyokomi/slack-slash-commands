package plugins

import (
	"fmt"
	"net/http"
	"slack"
)

type SlashCommand interface {
	AddPlugin(cmd string, plugin Plugin)
	Execute(req slack.SlashCommandRequest) string
}

type slashCommand struct {
	slash   *slack.SlashCommandService
	plugins map[string]Plugin
}

func (c *slashCommand) AddPlugin(cmd string, plugin Plugin) {
	c.plugins[cmd] = plugin
}

func (c *slashCommand) Execute(req slack.SlashCommandRequest) string {
	cmd, _ := req.CmdArgs()
	p, ok := c.plugins[cmd]
	if !ok {
		return fmt.Sprintf("%s command not found", cmd)
	}

	resp, err := c.slash.Reply(req, p.Do(req))
	if err != nil {
		return err.Error()
	}
	resp.Body.Close()

	return ""
}

func New(client *http.Client, plugins map[string]Plugin) SlashCommand {
	return &slashCommand{
		slash:   slack.New(client).SlashCommand,
		plugins: plugins,
	}
}

type Plugin interface {
	Do(req slack.SlashCommandRequest) slack.SlashCommandMessage
}
