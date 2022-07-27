package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/shinsuke-nara/echo-server/db"
	"github.com/shinsuke-nara/echo-server/restapi"
	"github.com/shinsuke-nara/echo-server/restapi/operations"
	"github.com/shinsuke-nara/echo-server/restapi/operations/board"
	"github.com/shinsuke-nara/echo-server/restimpl"
	boardimpl "github.com/shinsuke-nara/echo-server/restimpl/board"
)

func main() {
	database, err := db.PrepareDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close(database)

	bs := boardimpl.Server{DB: database}
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewEchoServerAPI(swaggerSpec)
	api.GetHelloHandler = operations.GetHelloHandlerFunc(restimpl.GetHello)
	api.PostEchoHandler = operations.PostEchoHandlerFunc(restimpl.PostEcho)
	api.BoardPostMessageHandler = board.PostMessageHandlerFunc(func(params board.PostMessageParams) middleware.Responder {
		return bs.PostMessage(params)
	})
	api.BoardListMessageHandler = board.ListMessageHandlerFunc(func(params board.ListMessageParams) middleware.Responder {
		return bs.ListMessage(params)
	})
	server := restapi.NewServer(api)
	defer server.Shutdown()
	server.Port = 80

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
