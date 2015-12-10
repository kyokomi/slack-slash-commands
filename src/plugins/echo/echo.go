package echo

import (
	"plugins"
	"strings"
)

type plugin struct {
}

func New() plugins.Plugin {
	return &plugin{}
}

func (p *plugin) Do(cmd string, args ...string) (responseMessage string) {
	return "Echo Message " + strings.Join(args, "")
}
