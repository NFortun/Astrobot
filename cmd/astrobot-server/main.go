package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/flowchartsman/swaggerui"
	"github.com/go-openapi/loads"

	"github.com/NFortun/Astrobot/internal/astrobin"
	"github.com/NFortun/Astrobot/restapi/operations"

	restapi "github.com/NFortun/Astrobot/restapi"
)

var (
	port = flag.Int("port", 8080, "default server port")
)

func main() {

	flag.Parse()
	astrobin.LoadConfig()

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewAstrobotAPI(swaggerSpec)
	server := restapi.NewServer(api)
	server.Port = *port
	defer server.Shutdown()

	go func() {
		d, err := swaggerSpec.Raw().MarshalJSON()
		if err != nil {
			panic(err)
		}
		http.Handle("/swagger/", http.StripPrefix("/swagger", swaggerui.Handler(d)))
		log.Println("serving on :8081")
		log.Fatal(http.ListenAndServe(":8081", nil))
	}()

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
