package restimpl

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/shinsuke-nara/echo-server/restapi/operations"
)

func GetHello(params operations.GetHelloParams) middleware.Responder {
	return &operations.GetHelloOK{Payload: "hello world"}
}
