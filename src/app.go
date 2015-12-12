package app

import (
	"net/http"

	"github.com/kyokomi/goslash/goslash"
	"github.com/kyokomi/goslash/plugins"
	"github.com/kyokomi/goslash/plugins/echo"
	"github.com/kyokomi/goslash/plugins/time"

	"github.com/unrolled/render"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

func init() {
	renderer := render.New(render.Options{})

	slashPlugins := map[string]plugins.Plugin{
		"echo": echo.New(),
		"time": time.New(),
	}

	http.HandleFunc("/v1/cmd", func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)

		req, err := goslash.ParseFormSlashCommandRequest(r)
		if err != nil {
			renderer.JSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		slashCmd := plugins.New(urlfetch.Client(ctx), slashPlugins)

		renderer.Text(w, http.StatusOK, slashCmd.Execute(req))
	})
}
