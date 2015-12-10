package app

import (
	"net/http"

	"plugins"
	"plugins/echo"
	"strings"
)

func init() {
	slackCmd := plugins.New()
	p := echo.New()

	slackCmd.AddPlugin("echo", p)

	http.HandleFunc("/v1/cmd", func(w http.ResponseWriter, r *http.Request) {
		args := strings.Fields(r.FormValue("text"))
		w.Write([]byte(slackCmd.Execute(args[0], args[1:]...)))
	})
}
