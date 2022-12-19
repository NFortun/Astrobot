package main

import (
	"flag"
	"log"
	"os"

	"github.com/go-openapi/loads"
	"github.com/jessevdk/go-flags"

	config "github.com/NFortun/Astrobot/config"
	restapi "github.com/NFortun/Astrobot/restapi"
	"github.com/NFortun/Astrobot/restapi/operations"
)

var port = flag.Int("port", 3000, "default server port")

func main() {

	flag.Parse()
	config.LoadConfig()

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewAstrobotAPI(swaggerSpec)
	server := restapi.NewServer(api)
	server.Port = *port
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Astrobot"
	parser.LongDescription = swaggerSpec.Spec().Info.Description
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
