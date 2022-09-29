package gosuapiclient

import "net/http"

type Client struct {
	authenticatedClient *http.Client
}

func NewClient() *Client {
	return &Client{}
}
