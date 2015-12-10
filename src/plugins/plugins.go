package plugins

import "fmt"

type command string

type SlashCommand interface {
	AddPlugin(cmd string, plugin Plugin)
	Execute(cmd string, args ...string) string
}

type slashCommandContext struct {
	Plugins map[command]Plugin
}

func (c *slashCommandContext) AddPlugin(cmd string, plugin Plugin) {
	c.Plugins[command(cmd)] = plugin
}

func (c *slashCommandContext) Execute(cmd string, args ...string) string {
	if len(cmd) == 0 {
		return "TODO: help"
	}

	p, ok := c.Plugins[command(cmd)]
	if !ok {
		return fmt.Sprintf("%s command not found", string(cmd))
	}
	return p.Do(cmd, args...)
}

func New() SlashCommand {
	return &slashCommandContext{
		Plugins: map[command]Plugin{},
	}
}

type Plugin interface {
	Do(cmd string, args ...string) (responseMessage string)
}
