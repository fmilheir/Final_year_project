// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"time"
	"fmt"
	"log"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/fmilheir/final_year_project/backend/restapi/operations"
	"github.com/fmilheir/final_year_project/backend/restapi/operations/diagnose_incident"
	"github.com/fmilheir/final_year_project/backend/restapi/operations/events_subscription"
	"github.com/fmilheir/final_year_project/backend/restapi/operations/incident"
	"github.com/fmilheir/final_year_project/backend/restapi/operations/notification_listeners_client_side"
	"github.com/fmilheir/final_year_project/backend/restapi/operations/resolve_incident"

)

//go:generate swagger generate server --target ..\..\backend --name Incident --spec ..\..\..\Documents\TMF724-Incident-v4.0.1.swagger.json --principal interface{}

func configureFlags(api *operations.IncidentAPI) {
	 //api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.IncidentAPI) http.Handler {
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
			return middleware.NotImplemented("operation incident.CreateIncident has not yet been implemented")
		})
	}
	if api.ResolveIncidentCreateResolveIncidentHandler == nil {
		api.ResolveIncidentCreateResolveIncidentHandler = resolve_incident.CreateResolveIncidentHandlerFunc(func(params resolve_incident.CreateResolveIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation resolve_incident.CreateResolveIncident has not yet been implemented")
		})
	}
	if api.DiagnoseIncidentListDiagnoseIncidentHandler == nil {
		api.DiagnoseIncidentListDiagnoseIncidentHandler = diagnose_incident.ListDiagnoseIncidentHandlerFunc(func(params diagnose_incident.ListDiagnoseIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation diagnose_incident.ListDiagnoseIncident has not yet been implemented")
		})
	}
	if api.IncidentListIncidentHandler == nil {
		api.IncidentListIncidentHandler = incident.ListIncidentHandlerFunc(func(params incident.ListIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation incident.ListIncident has not yet been implemented")
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
	tlsConfig.MinVersion = tls.VersionTLS12
    tlsConfig.CipherSuites = []uint16{
        tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
        tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
    }
	 tlsConfig.NextProtos = append(tlsConfig.NextProtos, "h2")
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
	if scheme == "https" {
        s.ReadTimeout = 10 * time.Second
        s.WriteTimeout = 10 * time.Second
    }

}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Log request, validate, etc.
        handler.ServeHTTP(w, r)
    })

}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	fmt.Print("I am here")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %+v", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		handler.ServeHTTP(w, r)
	})
}
