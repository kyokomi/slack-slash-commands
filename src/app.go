package app

import (
	"net/http"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine"
)

func init() {
	http.HandleFunc("/v1/cmd", slackSlashCommands)
}

func slackSlashCommands(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	r.ParseForm()
	log.Debugf(ctx, "postForm %#v\n", r.PostForm)

	w.Write([]byte("Hello World slack-slash-commands"))
}
