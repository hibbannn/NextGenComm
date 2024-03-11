package client

import "github.com/hibbannn/grpc-http-boilerplate/app/core/domain"

type Client struct {
	config *domain.Config
}

func NewClient(cfg *domain.Config) *Client {
	return &Client{
		config: cfg,
	}
}
