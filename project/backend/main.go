package main

import (
	"log"
	"os"
	"fmt"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
 
    "github.com/fmilheir/final_year_project/backend/database"
    "github.com/fmilheir/final_year_project/backend/routes"
	"github.com/fmilheir/final_year_project/backend/restapi"
	"github.com/fmilheir/final_year_project/backend/restapi/operations"
	
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
	AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8080")

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewIncidentAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Incident"
	parser.LongDescription = "**TMF API Reference : TMF - 724 Incident**\n\n**Release : 22.5 - November 2022**\n\nThe Incident API provides a standardized client interface to Incident Management Systems for creating, tracking and managing Incidents as a result of an issue or problem identified by a customer or another system. \nExamples of Incident API originators (clients) include CRM applications, network management or fault management systems, or other Incident management systems (e.g. B2B).\n\nThe API supports the ability to send requests to create a new Incident specifying the nature and severity of the trouble or issue as well as all necessary related information. The API also includes mechanisms to search for and update existing Incidents. Notifications are defined to provide information when a Incident has been updated, including status changes. A basic set of states of a Incident has been specified (as an example) to handle Incident lifecycle management.\nIncident API manages Incident resource:\n\n  -\tA Incident represents a record, or an issue raised by requestor that need to be solved, used for reporting and managing the resolution of problems, incidents or request -\tThe main Incident attributes are its name, priority, type,  dateTime attributes (occurTime, expected resolution, resolution), state and related information (change reason and change date), related parties (originator, owner, reviser, etc.), related entities (product, product order, customer bill) and notes Incident API performs the following operations on Incident -\tRetrieval of an Incident or a collection of Incident depending on filter criteria -\tPartial update of a Incident -\tCreation of an Incident -\tNotification of events on Incident: o\tIncident state change o\tIncident change o\tIncident resolved o\tIncident created o\tIncident Information required\n\n\n\nCopyright © TM Forum 2022. All Rights Reserved\n\n\n"
	fmt.Print("Incident API")
	fmt.Print("I am here!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
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
		fmt.Print("Error")
		log.Fatalln(err)
	}
}