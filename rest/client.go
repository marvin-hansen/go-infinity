package rest

import (
	"github.com/valyala/fasthttp"
	"time"
)

const defaultEndpoint = "http://localhost:8185/"
const defaultTimeout = 5 // seconds

type Client struct {
	Endpoint    string
	HTTPC       *fasthttp.Client
	HTTPTimeout time.Duration
}

func NewClient(config *ClientConfig) *Client {
	return &Client{
		Endpoint:    getEndpoint(config),
		HTTPTimeout: getTimeout(config),
		HTTPC:       new(fasthttp.Client),
	}
}
