package main

import (
	"log"

	"github.com/go-openapi/loads"
	"github.com/shinsuke-nara/echo-server/restapi"
	"github.com/shinsuke-nara/echo-server/restapi/operations"
	"github.com/shinsuke-nara/echo-server/restimpl"
)

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewEchoServerAPI(swaggerSpec)
	api.GetHelloHandler = operations.GetHelloHandlerFunc(restimpl.GetHello)
	api.PostEchoHandler = operations.PostEchoHandlerFunc(restimpl.PostEcho)
	server := restapi.NewServer(api)
	defer server.Shutdown()
	server.Port = 80

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
