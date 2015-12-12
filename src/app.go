package app

import (
	"net/http"

	"plugins"
	"plugins/echo"
	"strings"
)

func init() {
	slackCmd := plugins.New()
	slackCmd.AddPlugin("echo", echo.New())

	http.HandleFunc("/v1/cmd", func(w http.ResponseWriter, r *http.Request) {
		args := strings.Fields(r.FormValue("text"))
		w.Write([]byte(slackCmd.Execute(args[0], args[1:]...)))
	})
}
