package restimpl

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/shinsuke-nara/echo-server/restapi/operations"
)

func PostEcho(params operations.PostEchoParams) middleware.Responder {
	return &operations.PostEchoOK{Payload: params.Message}
}
