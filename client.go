package gosuapiclient

import (
	"context"
	"net/http"

	"go.uber.org/ratelimit"
	"golang.org/x/oauth2"
)

const BASE_URL = "https://osu.ppy.sh/api/v2"

type Client struct {
	authenticatedClient *http.Client
	rateLimiter         ratelimit.Limiter
}

func NewClient(token oauth2.Token, context context.Context) *Client {
	return &Client{
		authenticatedClient: oauth2.NewClient(context, oauth2.StaticTokenSource(&token)),
		rateLimiter:         ratelimit.New(60),
	}
}
