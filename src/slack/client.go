package slack

import (
	"net/http"
)

type Client struct {
	*http.Client

	SlashCommand *SlashCommandService
}

func New(httpClient *http.Client) Client {
	client := Client{
		Client: httpClient,
	}
	client.SlashCommand = &SlashCommandService{client: &client}

	return client
}
