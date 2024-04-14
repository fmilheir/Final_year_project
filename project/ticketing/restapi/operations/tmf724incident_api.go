// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"fmilheir/test_swagger/restapi/operations/diagnose_incident"
	"fmilheir/test_swagger/restapi/operations/events_subscription"
	"fmilheir/test_swagger/restapi/operations/incident"
	"fmilheir/test_swagger/restapi/operations/notification_listeners_client_side"
	"fmilheir/test_swagger/restapi/operations/resolve_incident"
)

// NewTmf724incidentAPI creates a new Tmf724incident instance
func NewTmf724incidentAPI(spec *loads.Document) *Tmf724incidentAPI {
	return &Tmf724incidentAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		DiagnoseIncidentCreateDiagnoseIncidentHandler: diagnose_incident.CreateDiagnoseIncidentHandlerFunc(func(params diagnose_incident.CreateDiagnoseIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation diagnose_incident.CreateDiagnoseIncident has not yet been implemented")
		}),
	
		ResolveIncidentCreateResolveIncidentHandler: resolve_incident.CreateResolveIncidentHandlerFunc(func(params resolve_incident.CreateResolveIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation resolve_incident.CreateResolveIncident has not yet been implemented")
		}),
		DiagnoseIncidentListDiagnoseIncidentHandler: diagnose_incident.ListDiagnoseIncidentHandlerFunc(func(params diagnose_incident.ListDiagnoseIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation diagnose_incident.ListDiagnoseIncident has not yet been implemented")
		}),
		IncidentListIncidentHandler: incident.ListIncidentHandlerFunc(func(params incident.ListIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation incident.ListIncident has not yet been implemented")
		}),
		ResolveIncidentListResolveIncidentHandler: resolve_incident.ListResolveIncidentHandlerFunc(func(params resolve_incident.ListResolveIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation resolve_incident.ListResolveIncident has not yet been implemented")
		}),
		NotificationListenersClientSideListenToDiagnoseIncidentCreateEventHandler: notification_listeners_client_side.ListenToDiagnoseIncidentCreateEventHandlerFunc(func(params notification_listeners_client_side.ListenToDiagnoseIncidentCreateEventParams) middleware.Responder {
			return middleware.NotImplemented("operation notification_listeners_client_side.ListenToDiagnoseIncidentCreateEvent has not yet been implemented")
		}),
		NotificationListenersClientSideListenToDiagnoseIncidentStateChangeEventHandler: notification_listeners_client_side.ListenToDiagnoseIncidentStateChangeEventHandlerFunc(func(params notification_listeners_client_side.ListenToDiagnoseIncidentStateChangeEventParams) middleware.Responder {
			return middleware.NotImplemented("operation notification_listeners_client_side.ListenToDiagnoseIncidentStateChangeEvent has not yet been implemented")
		}),
		NotificationListenersClientSideListenToIncidentCreateEventHandler: notification_listeners_client_side.ListenToIncidentCreateEventHandlerFunc(func(params notification_listeners_client_side.ListenToIncidentCreateEventParams) middleware.Responder {
			return middleware.NotImplemented("operation notification_listeners_client_side.ListenToIncidentCreateEvent has not yet been implemented")
		}),
		NotificationListenersClientSideListenToIncidentStateChangeEventHandler: notification_listeners_client_side.ListenToIncidentStateChangeEventHandlerFunc(func(params notification_listeners_client_side.ListenToIncidentStateChangeEventParams) middleware.Responder {
			return middleware.NotImplemented("operation notification_listeners_client_side.ListenToIncidentStateChangeEvent has not yet been implemented")
		}),
		NotificationListenersClientSideListenToResolveIncidentCreateEventHandler: notification_listeners_client_side.ListenToResolveIncidentCreateEventHandlerFunc(func(params notification_listeners_client_side.ListenToResolveIncidentCreateEventParams) middleware.Responder {
			return middleware.NotImplemented("operation notification_listeners_client_side.ListenToResolveIncidentCreateEvent has not yet been implemented")
		}),
		NotificationListenersClientSideListenToResolveIncidentStateChangeEventHandler: notification_listeners_client_side.ListenToResolveIncidentStateChangeEventHandlerFunc(func(params notification_listeners_client_side.ListenToResolveIncidentStateChangeEventParams) middleware.Responder {
			return middleware.NotImplemented("operation notification_listeners_client_side.ListenToResolveIncidentStateChangeEvent has not yet been implemented")
		}),
		EventsSubscriptionRegisterListenerHandler: events_subscription.RegisterListenerHandlerFunc(func(params events_subscription.RegisterListenerParams) middleware.Responder {
			return middleware.NotImplemented("operation events_subscription.RegisterListener has not yet been implemented")
		}),
		DiagnoseIncidentRetrieveDiagnoseIncidentHandler: diagnose_incident.RetrieveDiagnoseIncidentHandlerFunc(func(params diagnose_incident.RetrieveDiagnoseIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation diagnose_incident.RetrieveDiagnoseIncident has not yet been implemented")
		}),
		IncidentRetrieveIncidentHandler: incident.RetrieveIncidentHandlerFunc(func(params incident.RetrieveIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation incident.RetrieveIncident has not yet been implemented")
		}),
		ResolveIncidentRetrieveResolveIncidentHandler: resolve_incident.RetrieveResolveIncidentHandlerFunc(func(params resolve_incident.RetrieveResolveIncidentParams) middleware.Responder {
			return middleware.NotImplemented("operation resolve_incident.RetrieveResolveIncident has not yet been implemented")
		}),
		EventsSubscriptionUnregisterListenerHandler: events_subscription.UnregisterListenerHandlerFunc(func(params events_subscription.UnregisterListenerParams) middleware.Responder {
			return middleware.NotImplemented("operation events_subscription.UnregisterListener has not yet been implemented")
		}),
	}
}

/*
Tmf724incidentAPI **TMF API Reference : TMF - 724 Incident**

**Release : 22.5 - November 2022**

The Incident API provides a standardized client interface to Incident Management Systems for creating, tracking and managing Incidents as a result of an issue or problem identified by a customer or another system.
Examples of Incident API originators (clients) include CRM applications, network management or fault management systems, or other Incident management systems (e.g. B2B).

The API supports the ability to send requests to create a new Incident specifying the nature and severity of the trouble or issue as well as all necessary related information. The API also includes mechanisms to search for and update existing Incidents. Notifications are defined to provide information when a Incident has been updated, including status changes. A basic set of states of a Incident has been specified (as an example) to handle Incident lifecycle management.
Incident API manages Incident resource:

  - A Incident represents a record, or an issue raised by requestor that need to be solved, used for reporting and managing the resolution of problems, incidents or request -	The main Incident attributes are its name, priority, type,  dateTime attributes (occurTime, expected resolution, resolution), state and related information (change reason and change date), related parties (originator, owner, reviser, etc.), related entities (product, product order, customer bill) and notes Incident API performs the following operations on Incident -	Retrieval of an Incident or a collection of Incident depending on filter criteria -	Partial update of a Incident -	Creation of an Incident -	Notification of events on Incident: o	Incident state change o	Incident change o	Incident resolved o	Incident created o	Incident Information required

Copyright © TM Forum 2022. All Rights Reserved
*/
type Tmf724incidentAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// DiagnoseIncidentCreateDiagnoseIncidentHandler sets the operation handler for the create diagnose incident operation
	DiagnoseIncidentCreateDiagnoseIncidentHandler diagnose_incident.CreateDiagnoseIncidentHandler
	// IncidentCreateIncidentHandler sets the operation handler for the create incident operation
	IncidentCreateIncidentHandler incident.CreateIncidentHandler
	// ResolveIncidentCreateResolveIncidentHandler sets the operation handler for the create resolve incident operation
	ResolveIncidentCreateResolveIncidentHandler resolve_incident.CreateResolveIncidentHandler
	// DiagnoseIncidentListDiagnoseIncidentHandler sets the operation handler for the list diagnose incident operation
	DiagnoseIncidentListDiagnoseIncidentHandler diagnose_incident.ListDiagnoseIncidentHandler
	// IncidentListIncidentHandler sets the operation handler for the list incident operation
	IncidentListIncidentHandler incident.ListIncidentHandler
	// ResolveIncidentListResolveIncidentHandler sets the operation handler for the list resolve incident operation
	ResolveIncidentListResolveIncidentHandler resolve_incident.ListResolveIncidentHandler
	// NotificationListenersClientSideListenToDiagnoseIncidentCreateEventHandler sets the operation handler for the listen to diagnose incident create event operation
	NotificationListenersClientSideListenToDiagnoseIncidentCreateEventHandler notification_listeners_client_side.ListenToDiagnoseIncidentCreateEventHandler
	// NotificationListenersClientSideListenToDiagnoseIncidentStateChangeEventHandler sets the operation handler for the listen to diagnose incident state change event operation
	NotificationListenersClientSideListenToDiagnoseIncidentStateChangeEventHandler notification_listeners_client_side.ListenToDiagnoseIncidentStateChangeEventHandler
	// NotificationListenersClientSideListenToIncidentCreateEventHandler sets the operation handler for the listen to incident create event operation
	NotificationListenersClientSideListenToIncidentCreateEventHandler notification_listeners_client_side.ListenToIncidentCreateEventHandler
	// NotificationListenersClientSideListenToIncidentStateChangeEventHandler sets the operation handler for the listen to incident state change event operation
	NotificationListenersClientSideListenToIncidentStateChangeEventHandler notification_listeners_client_side.ListenToIncidentStateChangeEventHandler
	// NotificationListenersClientSideListenToResolveIncidentCreateEventHandler sets the operation handler for the listen to resolve incident create event operation
	NotificationListenersClientSideListenToResolveIncidentCreateEventHandler notification_listeners_client_side.ListenToResolveIncidentCreateEventHandler
	// NotificationListenersClientSideListenToResolveIncidentStateChangeEventHandler sets the operation handler for the listen to resolve incident state change event operation
	NotificationListenersClientSideListenToResolveIncidentStateChangeEventHandler notification_listeners_client_side.ListenToResolveIncidentStateChangeEventHandler
	// EventsSubscriptionRegisterListenerHandler sets the operation handler for the register listener operation
	EventsSubscriptionRegisterListenerHandler events_subscription.RegisterListenerHandler
	// DiagnoseIncidentRetrieveDiagnoseIncidentHandler sets the operation handler for the retrieve diagnose incident operation
	DiagnoseIncidentRetrieveDiagnoseIncidentHandler diagnose_incident.RetrieveDiagnoseIncidentHandler
	// IncidentRetrieveIncidentHandler sets the operation handler for the retrieve incident operation
	IncidentRetrieveIncidentHandler incident.RetrieveIncidentHandler
	// ResolveIncidentRetrieveResolveIncidentHandler sets the operation handler for the retrieve resolve incident operation
	ResolveIncidentRetrieveResolveIncidentHandler resolve_incident.RetrieveResolveIncidentHandler
	// EventsSubscriptionUnregisterListenerHandler sets the operation handler for the unregister listener operation
	EventsSubscriptionUnregisterListenerHandler events_subscription.UnregisterListenerHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *Tmf724incidentAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *Tmf724incidentAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *Tmf724incidentAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *Tmf724incidentAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *Tmf724incidentAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *Tmf724incidentAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *Tmf724incidentAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *Tmf724incidentAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *Tmf724incidentAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the Tmf724incidentAPI
func (o *Tmf724incidentAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.DiagnoseIncidentCreateDiagnoseIncidentHandler == nil {
		unregistered = append(unregistered, "diagnose_incident.CreateDiagnoseIncidentHandler")
	}
	if o.IncidentCreateIncidentHandler == nil {
		unregistered = append(unregistered, "incident.CreateIncidentHandler")
	}
	if o.ResolveIncidentCreateResolveIncidentHandler == nil {
		unregistered = append(unregistered, "resolve_incident.CreateResolveIncidentHandler")
	}
	if o.DiagnoseIncidentListDiagnoseIncidentHandler == nil {
		unregistered = append(unregistered, "diagnose_incident.ListDiagnoseIncidentHandler")
	}
	if o.IncidentListIncidentHandler == nil {
		unregistered = append(unregistered, "incident.ListIncidentHandler")
	}
	if o.ResolveIncidentListResolveIncidentHandler == nil {
		unregistered = append(unregistered, "resolve_incident.ListResolveIncidentHandler")
	}
	if o.NotificationListenersClientSideListenToDiagnoseIncidentCreateEventHandler == nil {
		unregistered = append(unregistered, "notification_listeners_client_side.ListenToDiagnoseIncidentCreateEventHandler")
	}
	if o.NotificationListenersClientSideListenToDiagnoseIncidentStateChangeEventHandler == nil {
		unregistered = append(unregistered, "notification_listeners_client_side.ListenToDiagnoseIncidentStateChangeEventHandler")
	}
	if o.NotificationListenersClientSideListenToIncidentCreateEventHandler == nil {
		unregistered = append(unregistered, "notification_listeners_client_side.ListenToIncidentCreateEventHandler")
	}
	if o.NotificationListenersClientSideListenToIncidentStateChangeEventHandler == nil {
		unregistered = append(unregistered, "notification_listeners_client_side.ListenToIncidentStateChangeEventHandler")
	}
	if o.NotificationListenersClientSideListenToResolveIncidentCreateEventHandler == nil {
		unregistered = append(unregistered, "notification_listeners_client_side.ListenToResolveIncidentCreateEventHandler")
	}
	if o.NotificationListenersClientSideListenToResolveIncidentStateChangeEventHandler == nil {
		unregistered = append(unregistered, "notification_listeners_client_side.ListenToResolveIncidentStateChangeEventHandler")
	}
	if o.EventsSubscriptionRegisterListenerHandler == nil {
		unregistered = append(unregistered, "events_subscription.RegisterListenerHandler")
	}
	if o.DiagnoseIncidentRetrieveDiagnoseIncidentHandler == nil {
		unregistered = append(unregistered, "diagnose_incident.RetrieveDiagnoseIncidentHandler")
	}
	if o.IncidentRetrieveIncidentHandler == nil {
		unregistered = append(unregistered, "incident.RetrieveIncidentHandler")
	}
	if o.ResolveIncidentRetrieveResolveIncidentHandler == nil {
		unregistered = append(unregistered, "resolve_incident.RetrieveResolveIncidentHandler")
	}
	if o.EventsSubscriptionUnregisterListenerHandler == nil {
		unregistered = append(unregistered, "events_subscription.UnregisterListenerHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *Tmf724incidentAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *Tmf724incidentAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *Tmf724incidentAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *Tmf724incidentAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *Tmf724incidentAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *Tmf724incidentAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the tmf724incident API
func (o *Tmf724incidentAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *Tmf724incidentAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/diagnoseIncident"] = diagnose_incident.NewCreateDiagnoseIncident(o.context, o.DiagnoseIncidentCreateDiagnoseIncidentHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/incident"] = incident.NewCreateIncident(o.context, o.IncidentCreateIncidentHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/resolveIncident"] = resolve_incident.NewCreateResolveIncident(o.context, o.ResolveIncidentCreateResolveIncidentHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/diagnoseIncident"] = diagnose_incident.NewListDiagnoseIncident(o.context, o.DiagnoseIncidentListDiagnoseIncidentHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/incident"] = incident.NewListIncident(o.context, o.IncidentListIncidentHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/resolveIncident"] = resolve_incident.NewListResolveIncident(o.context, o.ResolveIncidentListResolveIncidentHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/listener/diagnoseIncidentCreateEvent"] = notification_listeners_client_side.NewListenToDiagnoseIncidentCreateEvent(o.context, o.NotificationListenersClientSideListenToDiagnoseIncidentCreateEventHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/listener/diagnoseIncidentStateChangeEvent"] = notification_listeners_client_side.NewListenToDiagnoseIncidentStateChangeEvent(o.context, o.NotificationListenersClientSideListenToDiagnoseIncidentStateChangeEventHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/listener/incidentCreateEvent"] = notification_listeners_client_side.NewListenToIncidentCreateEvent(o.context, o.NotificationListenersClientSideListenToIncidentCreateEventHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/listener/incidentStateChangeEvent"] = notification_listeners_client_side.NewListenToIncidentStateChangeEvent(o.context, o.NotificationListenersClientSideListenToIncidentStateChangeEventHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/listener/resolveIncidentCreateEvent"] = notification_listeners_client_side.NewListenToResolveIncidentCreateEvent(o.context, o.NotificationListenersClientSideListenToResolveIncidentCreateEventHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/listener/resolveIncidentStateChangeEvent"] = notification_listeners_client_side.NewListenToResolveIncidentStateChangeEvent(o.context, o.NotificationListenersClientSideListenToResolveIncidentStateChangeEventHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/hub"] = events_subscription.NewRegisterListener(o.context, o.EventsSubscriptionRegisterListenerHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/diagnoseIncident/{id}"] = diagnose_incident.NewRetrieveDiagnoseIncident(o.context, o.DiagnoseIncidentRetrieveDiagnoseIncidentHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/incident/{id}"] = incident.NewRetrieveIncident(o.context, o.IncidentRetrieveIncidentHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/resolveIncident/{id}"] = resolve_incident.NewRetrieveResolveIncident(o.context, o.ResolveIncidentRetrieveResolveIncidentHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/hub/{id}"] = events_subscription.NewUnregisterListener(o.context, o.EventsSubscriptionUnregisterListenerHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *Tmf724incidentAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *Tmf724incidentAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *Tmf724incidentAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *Tmf724incidentAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *Tmf724incidentAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[um][path] = builder(h)
	}
}
