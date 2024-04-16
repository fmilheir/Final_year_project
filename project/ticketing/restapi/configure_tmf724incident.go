// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"time"
	"fmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	"fmilheir/test_swagger/restapi/operations"
	"fmilheir/test_swagger/restapi/operations/diagnose_incident"
	"fmilheir/test_swagger/restapi/operations/events_subscription"
	"fmilheir/test_swagger/restapi/operations/incident"
	"fmilheir/test_swagger/restapi/operations/notification_listeners_client_side"
	"fmilheir/test_swagger/restapi/operations/resolve_incident"
	"fmilheir/test_swagger/models"
	"fmilheir/test_swagger/database"
)

//go:generate swagger generate server --target ..\..\test_swaggeer --name Tmf724incident --spec ..\Documents\TMF724-Incident-v4.0.1.swagger.json --principal interface{}

func configureFlags(api *operations.Tmf724incidentAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.Tmf724incidentAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.DiagnoseIncidentCreateDiagnoseIncidentHandler == nil {
		api.DiagnoseIncidentCreateDiagnoseIncidentHandler = diagnose_incident.CreateDiagnoseIncidentHandlerFunc(func(params diagnose_incident.CreateDiagnoseIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation diagnose_incident.CreateDiagnoseIncident has not yet been implemented")
		})
	}
	if api.IncidentCreateIncidentHandler == nil {
		api.IncidentCreateIncidentHandler = incident.CreateIncidentHandlerFunc(func(params incident.CreateIncidentParams) middleware.Responder {
			rootEventIDs := make([]*models.ResourceEntity, len(params.Incident.RootEventID))
			for i, v := range params.Incident.RootEventID {
				rootEventIDs[i] = v
			}

			fmt.Println(len(rootEventIDs))
			
			fmt.Println("Start CreateIncidentHandlerFunc")
			fmt.Println(*params.Incident)
			fmt.Println(*params.Incident.Name)
			fmt.Println(*params.Incident.Category)
			fmt.Println(*params.Incident.Priority)
			fmt.Println(*params.Incident.State)
			fmt.Println(*params.Incident.AckState)
			//fmt.Println(*sourceObjects[0])
			//fmt.Println(*sourceObjects[1])
			fmt.Println(rootEventIDs)
			fmt.Println(*rootEventIDs[0])
			fmt.Println(params.Incident.SourceObject)
			fmt.Println("My values are: ",*params.Incident.SourceObject[0])

			
			//fmt.Println(*params.Incident.SourceObject)

			
			name := params.Incident.Name
			category := params.Incident.Category
			priority := params.Incident.Priority
			state := params.Incident.State
			ackState := params.Incident.AckState
			occurTime := strfmt.DateTime(time.Now())
			domain := "IT"
			sourceObject := params.Incident.SourceObject
			detail := params.Incident.IncidentDetail

			newIncident := &models.Incident{
				Name:         *name,
				Category:     *category,
				Priority:     *priority,
				State:        *state,
				AckState:     *ackState,
				OccurTime:    occurTime,
				Domain:       domain,
				RootEventID:  rootEventIDs,
				SourceObject: sourceObject,
				IncidentDetail: detail,
			}
			fmt.Println("New Incident: ", newIncident)

			if err := database.DB.Db.Create(newIncident); err != nil {
				// Log or print the error message
				fmt.Println("Error creating incident:", err)
				return incident.NewCreateIncidentInternalServerError().WithPayload(&models.Error{Message: "Failed to create incident"})
			}

			return incident.NewCreateIncidentCreated().WithPayload(newIncident)
		})
	}
	if api.ResolveIncidentCreateResolveIncidentHandler == nil {
		api.ResolveIncidentCreateResolveIncidentHandler = resolve_incident.CreateResolveIncidentHandlerFunc(func(params resolve_incident.CreateResolveIncidentParams) middleware.Responder {
			if params.ResolveIncident.Incident == nil || params.ResolveIncident.Incident.ID == nil || params.ResolveIncident.Incident.IncidentResolutionSuggestion == ""{
				return resolve_incident.NewCreateResolveIncidentBadRequest().WithPayload(&models.Error{Message: "Incident is required"})
			}
			
			ID:= *params.ResolveIncident.Incident.ID
			IncidentResolutionSuggestion := params.ResolveIncident.Incident.IncidentResolutionSuggestion

			fmt.Println("Start CreateResolveIncidentHandlerFunc")
			//update incident whre id = ID changing the state to RESOLVED and adding the IncidentResolutionSuggestion
			if database.DB.Db.Model(&models.Incident{}).Where("id = ?", ID).Update("state", "RESOLVED").Error != nil {
				return resolve_incident.NewCreateResolveIncidentInternalServerError().WithPayload(&models.Error{Message: "Failed to update incident"})
			}
			
			if database.DB.Db.Model(&models.Incident{}).Where("id = ?", ID).Update("IncidentResolutionSuggestion", IncidentResolutionSuggestion).Error != nil {
				return resolve_incident.NewCreateResolveIncidentInternalServerError().WithPayload(&models.Error{Message: "Failed to update incident with resolution"})
			}

			return resolve_incident.NewCreateResolveIncidentCreated().WithPayload(&models.ResolveIncident{ID: ID})
		})
	}
	if api.DiagnoseIncidentListDiagnoseIncidentHandler == nil {
		api.DiagnoseIncidentListDiagnoseIncidentHandler = diagnose_incident.ListDiagnoseIncidentHandlerFunc(func(params diagnose_incident.ListDiagnoseIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation diagnose_incident.ListDiagnoseIncident has not yet been implemented")
		})
	}
	if api.IncidentListIncidentHandler == nil {
		api.IncidentListIncidentHandler = incident.ListIncidentHandlerFunc(func(params incident.ListIncidentParams) middleware.Responder {
			var incidents []*models.Incident
			
			// Get all incidents from the database
			if err := database.DB.Db.Find(&incidents).Error; err != nil {
				fmt.Println("Error getting incidents:", err)
				return incident.NewListIncidentInternalServerError().WithPayload(&models.Error{Message: "Failed to get incidents"})
			}

			return incident.NewListIncidentOK().WithPayload(incidents)
		})
	}
	if api.ResolveIncidentListResolveIncidentHandler == nil {
		api.ResolveIncidentListResolveIncidentHandler = resolve_incident.ListResolveIncidentHandlerFunc(func(params resolve_incident.ListResolveIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation resolve_incident.ListResolveIncident has not yet been implemented")
		})
	}
	if api.NotificationListenersClientSideListenToDiagnoseIncidentCreateEventHandler == nil {
		api.NotificationListenersClientSideListenToDiagnoseIncidentCreateEventHandler = notification_listeners_client_side.ListenToDiagnoseIncidentCreateEventHandlerFunc(func(params notification_listeners_client_side.ListenToDiagnoseIncidentCreateEventParams) middleware.Responder {
			return middleware.NotImplemented("operation notification_listeners_client_side.ListenToDiagnoseIncidentCreateEvent has not yet been implemented")
		})
	}
	if api.NotificationListenersClientSideListenToDiagnoseIncidentStateChangeEventHandler == nil {
		api.NotificationListenersClientSideListenToDiagnoseIncidentStateChangeEventHandler = notification_listeners_client_side.ListenToDiagnoseIncidentStateChangeEventHandlerFunc(func(params notification_listeners_client_side.ListenToDiagnoseIncidentStateChangeEventParams) middleware.Responder {
			return middleware.NotImplemented("operation notification_listeners_client_side.ListenToDiagnoseIncidentStateChangeEvent has not yet been implemented")
		})
	}
	if api.NotificationListenersClientSideListenToIncidentCreateEventHandler == nil {
		api.NotificationListenersClientSideListenToIncidentCreateEventHandler = notification_listeners_client_side.ListenToIncidentCreateEventHandlerFunc(func(params notification_listeners_client_side.ListenToIncidentCreateEventParams) middleware.Responder {
			return middleware.NotImplemented("operation notification_listeners_client_side.ListenToIncidentCreateEvent has not yet been implemented")
		})
	}
	if api.NotificationListenersClientSideListenToIncidentStateChangeEventHandler == nil {
		api.NotificationListenersClientSideListenToIncidentStateChangeEventHandler = notification_listeners_client_side.ListenToIncidentStateChangeEventHandlerFunc(func(params notification_listeners_client_side.ListenToIncidentStateChangeEventParams) middleware.Responder {
			return middleware.NotImplemented("operation notification_listeners_client_side.ListenToIncidentStateChangeEvent has not yet been implemented")
		})
	}
	if api.NotificationListenersClientSideListenToResolveIncidentCreateEventHandler == nil {
		api.NotificationListenersClientSideListenToResolveIncidentCreateEventHandler = notification_listeners_client_side.ListenToResolveIncidentCreateEventHandlerFunc(func(params notification_listeners_client_side.ListenToResolveIncidentCreateEventParams) middleware.Responder {
			return middleware.NotImplemented("operation notification_listeners_client_side.ListenToResolveIncidentCreateEvent has not yet been implemented")
		})
	}
	if api.NotificationListenersClientSideListenToResolveIncidentStateChangeEventHandler == nil {
		api.NotificationListenersClientSideListenToResolveIncidentStateChangeEventHandler = notification_listeners_client_side.ListenToResolveIncidentStateChangeEventHandlerFunc(func(params notification_listeners_client_side.ListenToResolveIncidentStateChangeEventParams) middleware.Responder {
			return middleware.NotImplemented("operation notification_listeners_client_side.ListenToResolveIncidentStateChangeEvent has not yet been implemented")
		})
	}
	if api.EventsSubscriptionRegisterListenerHandler == nil {
		api.EventsSubscriptionRegisterListenerHandler = events_subscription.RegisterListenerHandlerFunc(func(params events_subscription.RegisterListenerParams) middleware.Responder {
			return middleware.NotImplemented("operation events_subscription.RegisterListener has not yet been implemented")
		})
	}
	if api.DiagnoseIncidentRetrieveDiagnoseIncidentHandler == nil {
		api.DiagnoseIncidentRetrieveDiagnoseIncidentHandler = diagnose_incident.RetrieveDiagnoseIncidentHandlerFunc(func(params diagnose_incident.RetrieveDiagnoseIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation diagnose_incident.RetrieveDiagnoseIncident has not yet been implemented")
		})
	}
	if api.IncidentRetrieveIncidentHandler == nil {
		api.IncidentRetrieveIncidentHandler = incident.RetrieveIncidentHandlerFunc(func(params incident.RetrieveIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation incident.RetrieveIncident has not yet been implemented")
		})
	}
	if api.ResolveIncidentRetrieveResolveIncidentHandler == nil {
		api.ResolveIncidentRetrieveResolveIncidentHandler = resolve_incident.RetrieveResolveIncidentHandlerFunc(func(params resolve_incident.RetrieveResolveIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation resolve_incident.RetrieveResolveIncident has not yet been implemented")
		})
	}
	if api.EventsSubscriptionUnregisterListenerHandler == nil {
		api.EventsSubscriptionUnregisterListenerHandler = events_subscription.UnregisterListenerHandlerFunc(func(params events_subscription.UnregisterListenerParams) middleware.Responder {
			return middleware.NotImplemented("operation events_subscription.UnregisterListener has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
