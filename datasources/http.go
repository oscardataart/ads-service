package datasources

import (
	"context"
)

type AuthorizationType string

var (
	BadGateway = Error{Msg: "BAD_GATEWAY_ERROR"}
)

type HTTPClient interface {
	// Performs HTTP request via REST.
	Request(ctx context.Context, info *RequestInfo) (*ResponseInfo, error)
}

type RequestInfo struct {
	Url        string
	HTTPMethod string
}

type ResponseInfo struct {
	StatusCode int
	Body       []byte
}

type Error struct {
	Msg string
}

func (e Error) Error() string {
	return e.Msg
}
