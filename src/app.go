package app

import (
	"net/http"

	"plugins"
	"plugins/echo"
	"plugins/time"
	"slack"

	"github.com/unrolled/render"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func init() {
	slackCmd := plugins.New()
	slackCmd.AddPlugin("echo", echo.New())
	slackCmd.AddPlugin("time", time.New())

	renderer := render.New(render.Options{})

	http.HandleFunc("/v1/cmd", func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)

		req, err := slack.ParseFormSlashCommandRequest(r)
		if err != nil {
			renderer.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		client := slack.New(urlfetch.Client(ctx))

		renderer.Text(w, http.StatusOK, slackCmd.Execute(client.SlashCommand, req))
	})
}
