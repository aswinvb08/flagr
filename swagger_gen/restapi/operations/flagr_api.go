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

	"github.com/openflagr/flagr/swagger_gen/restapi/operations/constraint"
	"github.com/openflagr/flagr/swagger_gen/restapi/operations/distribution"
	"github.com/openflagr/flagr/swagger_gen/restapi/operations/evaluation"
	"github.com/openflagr/flagr/swagger_gen/restapi/operations/export"
	"github.com/openflagr/flagr/swagger_gen/restapi/operations/flag"
	"github.com/openflagr/flagr/swagger_gen/restapi/operations/health"
	"github.com/openflagr/flagr/swagger_gen/restapi/operations/segment"
	"github.com/openflagr/flagr/swagger_gen/restapi/operations/tag"
	"github.com/openflagr/flagr/swagger_gen/restapi/operations/variant"
)

// NewFlagrAPI creates a new Flagr instance
func NewFlagrAPI(spec *loads.Document) *FlagrAPI {
	return &FlagrAPI{
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

		BinProducer:  runtime.ByteStreamProducer(),
		JSONProducer: runtime.JSONProducer(),

		ConstraintCreateConstraintHandler: constraint.CreateConstraintHandlerFunc(func(params constraint.CreateConstraintParams) middleware.Responder {
			return middleware.NotImplemented("operation constraint.CreateConstraint has not yet been implemented")
		}),
		FlagCreateFlagHandler: flag.CreateFlagHandlerFunc(func(params flag.CreateFlagParams) middleware.Responder {
			return middleware.NotImplemented("operation flag.CreateFlag has not yet been implemented")
		}),
		SegmentCreateSegmentHandler: segment.CreateSegmentHandlerFunc(func(params segment.CreateSegmentParams) middleware.Responder {
			return middleware.NotImplemented("operation segment.CreateSegment has not yet been implemented")
		}),
		TagCreateTagHandler: tag.CreateTagHandlerFunc(func(params tag.CreateTagParams) middleware.Responder {
			return middleware.NotImplemented("operation tag.CreateTag has not yet been implemented")
		}),
		VariantCreateVariantHandler: variant.CreateVariantHandlerFunc(func(params variant.CreateVariantParams) middleware.Responder {
			return middleware.NotImplemented("operation variant.CreateVariant has not yet been implemented")
		}),
		ConstraintDeleteConstraintHandler: constraint.DeleteConstraintHandlerFunc(func(params constraint.DeleteConstraintParams) middleware.Responder {
			return middleware.NotImplemented("operation constraint.DeleteConstraint has not yet been implemented")
		}),
		FlagDeleteFlagHandler: flag.DeleteFlagHandlerFunc(func(params flag.DeleteFlagParams) middleware.Responder {
			return middleware.NotImplemented("operation flag.DeleteFlag has not yet been implemented")
		}),
		SegmentDeleteSegmentHandler: segment.DeleteSegmentHandlerFunc(func(params segment.DeleteSegmentParams) middleware.Responder {
			return middleware.NotImplemented("operation segment.DeleteSegment has not yet been implemented")
		}),
		TagDeleteTagHandler: tag.DeleteTagHandlerFunc(func(params tag.DeleteTagParams) middleware.Responder {
			return middleware.NotImplemented("operation tag.DeleteTag has not yet been implemented")
		}),
		VariantDeleteVariantHandler: variant.DeleteVariantHandlerFunc(func(params variant.DeleteVariantParams) middleware.Responder {
			return middleware.NotImplemented("operation variant.DeleteVariant has not yet been implemented")
		}),
		TagFindAllTagsHandler: tag.FindAllTagsHandlerFunc(func(params tag.FindAllTagsParams) middleware.Responder {
			return middleware.NotImplemented("operation tag.FindAllTags has not yet been implemented")
		}),
		ConstraintFindConstraintsHandler: constraint.FindConstraintsHandlerFunc(func(params constraint.FindConstraintsParams) middleware.Responder {
			return middleware.NotImplemented("operation constraint.FindConstraints has not yet been implemented")
		}),
		DistributionFindDistributionsHandler: distribution.FindDistributionsHandlerFunc(func(params distribution.FindDistributionsParams) middleware.Responder {
			return middleware.NotImplemented("operation distribution.FindDistributions has not yet been implemented")
		}),
		FlagFindFlagsHandler: flag.FindFlagsHandlerFunc(func(params flag.FindFlagsParams) middleware.Responder {
			return middleware.NotImplemented("operation flag.FindFlags has not yet been implemented")
		}),
		SegmentFindSegmentsHandler: segment.FindSegmentsHandlerFunc(func(params segment.FindSegmentsParams) middleware.Responder {
			return middleware.NotImplemented("operation segment.FindSegments has not yet been implemented")
		}),
		TagFindTagsHandler: tag.FindTagsHandlerFunc(func(params tag.FindTagsParams) middleware.Responder {
			return middleware.NotImplemented("operation tag.FindTags has not yet been implemented")
		}),
		VariantFindVariantsHandler: variant.FindVariantsHandlerFunc(func(params variant.FindVariantsParams) middleware.Responder {
			return middleware.NotImplemented("operation variant.FindVariants has not yet been implemented")
		}),
		ExportGetExportEvalCacheJSONHandler: export.GetExportEvalCacheJSONHandlerFunc(func(params export.GetExportEvalCacheJSONParams) middleware.Responder {
			return middleware.NotImplemented("operation export.GetExportEvalCacheJSON has not yet been implemented")
		}),
		ExportGetExportSqliteHandler: export.GetExportSqliteHandlerFunc(func(params export.GetExportSqliteParams) middleware.Responder {
			return middleware.NotImplemented("operation export.GetExportSqlite has not yet been implemented")
		}),
		FlagGetFlagHandler: flag.GetFlagHandlerFunc(func(params flag.GetFlagParams) middleware.Responder {
			return middleware.NotImplemented("operation flag.GetFlag has not yet been implemented")
		}),
		FlagGetFlagEntityTypesHandler: flag.GetFlagEntityTypesHandlerFunc(func(params flag.GetFlagEntityTypesParams) middleware.Responder {
			return middleware.NotImplemented("operation flag.GetFlagEntityTypes has not yet been implemented")
		}),
		FlagGetFlagSnapshotsHandler: flag.GetFlagSnapshotsHandlerFunc(func(params flag.GetFlagSnapshotsParams) middleware.Responder {
			return middleware.NotImplemented("operation flag.GetFlagSnapshots has not yet been implemented")
		}),
		HealthGetHealthHandler: health.GetHealthHandlerFunc(func(params health.GetHealthParams) middleware.Responder {
			return middleware.NotImplemented("operation health.GetHealth has not yet been implemented")
		}),
		EvaluationPostEvaluationHandler: evaluation.PostEvaluationHandlerFunc(func(params evaluation.PostEvaluationParams) middleware.Responder {
			return middleware.NotImplemented("operation evaluation.PostEvaluation has not yet been implemented")
		}),
		EvaluationPostEvaluationBatchHandler: evaluation.PostEvaluationBatchHandlerFunc(func(params evaluation.PostEvaluationBatchParams) middleware.Responder {
			return middleware.NotImplemented("operation evaluation.PostEvaluationBatch has not yet been implemented")
		}),
		ConstraintPutConstraintHandler: constraint.PutConstraintHandlerFunc(func(params constraint.PutConstraintParams) middleware.Responder {
			return middleware.NotImplemented("operation constraint.PutConstraint has not yet been implemented")
		}),
		DistributionPutDistributionsHandler: distribution.PutDistributionsHandlerFunc(func(params distribution.PutDistributionsParams) middleware.Responder {
			return middleware.NotImplemented("operation distribution.PutDistributions has not yet been implemented")
		}),
		FlagPutFlagHandler: flag.PutFlagHandlerFunc(func(params flag.PutFlagParams) middleware.Responder {
			return middleware.NotImplemented("operation flag.PutFlag has not yet been implemented")
		}),
		SegmentPutSegmentHandler: segment.PutSegmentHandlerFunc(func(params segment.PutSegmentParams) middleware.Responder {
			return middleware.NotImplemented("operation segment.PutSegment has not yet been implemented")
		}),
		SegmentPutSegmentsReorderHandler: segment.PutSegmentsReorderHandlerFunc(func(params segment.PutSegmentsReorderParams) middleware.Responder {
			return middleware.NotImplemented("operation segment.PutSegmentsReorder has not yet been implemented")
		}),
		VariantPutVariantHandler: variant.PutVariantHandlerFunc(func(params variant.PutVariantParams) middleware.Responder {
			return middleware.NotImplemented("operation variant.PutVariant has not yet been implemented")
		}),
		FlagRestoreFlagHandler: flag.RestoreFlagHandlerFunc(func(params flag.RestoreFlagParams) middleware.Responder {
			return middleware.NotImplemented("operation flag.RestoreFlag has not yet been implemented")
		}),
		FlagSetFlagEnabledHandler: flag.SetFlagEnabledHandlerFunc(func(params flag.SetFlagEnabledParams) middleware.Responder {
			return middleware.NotImplemented("operation flag.SetFlagEnabled has not yet been implemented")
		}),
	}
}

/*FlagrAPI Flagr is a feature flagging, A/B testing and dynamic configuration microservice. The base path for all the APIs is "/api/v1".
 */
type FlagrAPI struct {
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

	// BinProducer registers a producer for the following mime types:
	//   - application/octet-stream
	BinProducer runtime.Producer
	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// ConstraintCreateConstraintHandler sets the operation handler for the create constraint operation
	ConstraintCreateConstraintHandler constraint.CreateConstraintHandler
	// FlagCreateFlagHandler sets the operation handler for the create flag operation
	FlagCreateFlagHandler flag.CreateFlagHandler
	// SegmentCreateSegmentHandler sets the operation handler for the create segment operation
	SegmentCreateSegmentHandler segment.CreateSegmentHandler
	// TagCreateTagHandler sets the operation handler for the create tag operation
	TagCreateTagHandler tag.CreateTagHandler
	// VariantCreateVariantHandler sets the operation handler for the create variant operation
	VariantCreateVariantHandler variant.CreateVariantHandler
	// ConstraintDeleteConstraintHandler sets the operation handler for the delete constraint operation
	ConstraintDeleteConstraintHandler constraint.DeleteConstraintHandler
	// FlagDeleteFlagHandler sets the operation handler for the delete flag operation
	FlagDeleteFlagHandler flag.DeleteFlagHandler
	// SegmentDeleteSegmentHandler sets the operation handler for the delete segment operation
	SegmentDeleteSegmentHandler segment.DeleteSegmentHandler
	// TagDeleteTagHandler sets the operation handler for the delete tag operation
	TagDeleteTagHandler tag.DeleteTagHandler
	// VariantDeleteVariantHandler sets the operation handler for the delete variant operation
	VariantDeleteVariantHandler variant.DeleteVariantHandler
	// TagFindAllTagsHandler sets the operation handler for the find all tags operation
	TagFindAllTagsHandler tag.FindAllTagsHandler
	// ConstraintFindConstraintsHandler sets the operation handler for the find constraints operation
	ConstraintFindConstraintsHandler constraint.FindConstraintsHandler
	// DistributionFindDistributionsHandler sets the operation handler for the find distributions operation
	DistributionFindDistributionsHandler distribution.FindDistributionsHandler
	// FlagFindFlagsHandler sets the operation handler for the find flags operation
	FlagFindFlagsHandler flag.FindFlagsHandler
	// SegmentFindSegmentsHandler sets the operation handler for the find segments operation
	SegmentFindSegmentsHandler segment.FindSegmentsHandler
	// TagFindTagsHandler sets the operation handler for the find tags operation
	TagFindTagsHandler tag.FindTagsHandler
	// VariantFindVariantsHandler sets the operation handler for the find variants operation
	VariantFindVariantsHandler variant.FindVariantsHandler
	// ExportGetExportEvalCacheJSONHandler sets the operation handler for the get export eval cache JSON operation
	ExportGetExportEvalCacheJSONHandler export.GetExportEvalCacheJSONHandler
	// ExportGetExportSqliteHandler sets the operation handler for the get export sqlite operation
	ExportGetExportSqliteHandler export.GetExportSqliteHandler
	// FlagGetFlagHandler sets the operation handler for the get flag operation
	FlagGetFlagHandler flag.GetFlagHandler
	// FlagGetFlagEntityTypesHandler sets the operation handler for the get flag entity types operation
	FlagGetFlagEntityTypesHandler flag.GetFlagEntityTypesHandler
	// FlagGetFlagSnapshotsHandler sets the operation handler for the get flag snapshots operation
	FlagGetFlagSnapshotsHandler flag.GetFlagSnapshotsHandler
	// HealthGetHealthHandler sets the operation handler for the get health operation
	HealthGetHealthHandler health.GetHealthHandler
	// EvaluationPostEvaluationHandler sets the operation handler for the post evaluation operation
	EvaluationPostEvaluationHandler evaluation.PostEvaluationHandler
	// EvaluationPostEvaluationBatchHandler sets the operation handler for the post evaluation batch operation
	EvaluationPostEvaluationBatchHandler evaluation.PostEvaluationBatchHandler
	// ConstraintPutConstraintHandler sets the operation handler for the put constraint operation
	ConstraintPutConstraintHandler constraint.PutConstraintHandler
	// DistributionPutDistributionsHandler sets the operation handler for the put distributions operation
	DistributionPutDistributionsHandler distribution.PutDistributionsHandler
	// FlagPutFlagHandler sets the operation handler for the put flag operation
	FlagPutFlagHandler flag.PutFlagHandler
	// SegmentPutSegmentHandler sets the operation handler for the put segment operation
	SegmentPutSegmentHandler segment.PutSegmentHandler
	// SegmentPutSegmentsReorderHandler sets the operation handler for the put segments reorder operation
	SegmentPutSegmentsReorderHandler segment.PutSegmentsReorderHandler
	// VariantPutVariantHandler sets the operation handler for the put variant operation
	VariantPutVariantHandler variant.PutVariantHandler
	// FlagRestoreFlagHandler sets the operation handler for the restore flag operation
	FlagRestoreFlagHandler flag.RestoreFlagHandler
	// FlagSetFlagEnabledHandler sets the operation handler for the set flag enabled operation
	FlagSetFlagEnabledHandler flag.SetFlagEnabledHandler

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
func (o *FlagrAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *FlagrAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *FlagrAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *FlagrAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *FlagrAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *FlagrAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *FlagrAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *FlagrAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *FlagrAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the FlagrAPI
func (o *FlagrAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.BinProducer == nil {
		unregistered = append(unregistered, "BinProducer")
	}
	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.ConstraintCreateConstraintHandler == nil {
		unregistered = append(unregistered, "constraint.CreateConstraintHandler")
	}
	if o.FlagCreateFlagHandler == nil {
		unregistered = append(unregistered, "flag.CreateFlagHandler")
	}
	if o.SegmentCreateSegmentHandler == nil {
		unregistered = append(unregistered, "segment.CreateSegmentHandler")
	}
	if o.TagCreateTagHandler == nil {
		unregistered = append(unregistered, "tag.CreateTagHandler")
	}
	if o.VariantCreateVariantHandler == nil {
		unregistered = append(unregistered, "variant.CreateVariantHandler")
	}
	if o.ConstraintDeleteConstraintHandler == nil {
		unregistered = append(unregistered, "constraint.DeleteConstraintHandler")
	}
	if o.FlagDeleteFlagHandler == nil {
		unregistered = append(unregistered, "flag.DeleteFlagHandler")
	}
	if o.SegmentDeleteSegmentHandler == nil {
		unregistered = append(unregistered, "segment.DeleteSegmentHandler")
	}
	if o.TagDeleteTagHandler == nil {
		unregistered = append(unregistered, "tag.DeleteTagHandler")
	}
	if o.VariantDeleteVariantHandler == nil {
		unregistered = append(unregistered, "variant.DeleteVariantHandler")
	}
	if o.TagFindAllTagsHandler == nil {
		unregistered = append(unregistered, "tag.FindAllTagsHandler")
	}
	if o.ConstraintFindConstraintsHandler == nil {
		unregistered = append(unregistered, "constraint.FindConstraintsHandler")
	}
	if o.DistributionFindDistributionsHandler == nil {
		unregistered = append(unregistered, "distribution.FindDistributionsHandler")
	}
	if o.FlagFindFlagsHandler == nil {
		unregistered = append(unregistered, "flag.FindFlagsHandler")
	}
	if o.SegmentFindSegmentsHandler == nil {
		unregistered = append(unregistered, "segment.FindSegmentsHandler")
	}
	if o.TagFindTagsHandler == nil {
		unregistered = append(unregistered, "tag.FindTagsHandler")
	}
	if o.VariantFindVariantsHandler == nil {
		unregistered = append(unregistered, "variant.FindVariantsHandler")
	}
	if o.ExportGetExportEvalCacheJSONHandler == nil {
		unregistered = append(unregistered, "export.GetExportEvalCacheJSONHandler")
	}
	if o.ExportGetExportSqliteHandler == nil {
		unregistered = append(unregistered, "export.GetExportSqliteHandler")
	}
	if o.FlagGetFlagHandler == nil {
		unregistered = append(unregistered, "flag.GetFlagHandler")
	}
	if o.FlagGetFlagEntityTypesHandler == nil {
		unregistered = append(unregistered, "flag.GetFlagEntityTypesHandler")
	}
	if o.FlagGetFlagSnapshotsHandler == nil {
		unregistered = append(unregistered, "flag.GetFlagSnapshotsHandler")
	}
	if o.HealthGetHealthHandler == nil {
		unregistered = append(unregistered, "health.GetHealthHandler")
	}
	if o.EvaluationPostEvaluationHandler == nil {
		unregistered = append(unregistered, "evaluation.PostEvaluationHandler")
	}
	if o.EvaluationPostEvaluationBatchHandler == nil {
		unregistered = append(unregistered, "evaluation.PostEvaluationBatchHandler")
	}
	if o.ConstraintPutConstraintHandler == nil {
		unregistered = append(unregistered, "constraint.PutConstraintHandler")
	}
	if o.DistributionPutDistributionsHandler == nil {
		unregistered = append(unregistered, "distribution.PutDistributionsHandler")
	}
	if o.FlagPutFlagHandler == nil {
		unregistered = append(unregistered, "flag.PutFlagHandler")
	}
	if o.SegmentPutSegmentHandler == nil {
		unregistered = append(unregistered, "segment.PutSegmentHandler")
	}
	if o.SegmentPutSegmentsReorderHandler == nil {
		unregistered = append(unregistered, "segment.PutSegmentsReorderHandler")
	}
	if o.VariantPutVariantHandler == nil {
		unregistered = append(unregistered, "variant.PutVariantHandler")
	}
	if o.FlagRestoreFlagHandler == nil {
		unregistered = append(unregistered, "flag.RestoreFlagHandler")
	}
	if o.FlagSetFlagEnabledHandler == nil {
		unregistered = append(unregistered, "flag.SetFlagEnabledHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *FlagrAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *FlagrAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *FlagrAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *FlagrAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
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
func (o *FlagrAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/octet-stream":
			result["application/octet-stream"] = o.BinProducer
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
func (o *FlagrAPI) HandlerFor(method, path string) (http.Handler, bool) {
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

// Context returns the middleware context for the flagr API
func (o *FlagrAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *FlagrAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/flags/{flagID}/segments/{segmentID}/constraints"] = constraint.NewCreateConstraint(o.context, o.ConstraintCreateConstraintHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/flags"] = flag.NewCreateFlag(o.context, o.FlagCreateFlagHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/flags/{flagID}/segments"] = segment.NewCreateSegment(o.context, o.SegmentCreateSegmentHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/flags/{flagID}/tags"] = tag.NewCreateTag(o.context, o.TagCreateTagHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/flags/{flagID}/variants"] = variant.NewCreateVariant(o.context, o.VariantCreateVariantHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/flags/{flagID}/segments/{segmentID}/constraints/{constraintID}"] = constraint.NewDeleteConstraint(o.context, o.ConstraintDeleteConstraintHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/flags/{flagID}"] = flag.NewDeleteFlag(o.context, o.FlagDeleteFlagHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/flags/{flagID}/segments/{segmentID}"] = segment.NewDeleteSegment(o.context, o.SegmentDeleteSegmentHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/flags/{flagID}/tags/{tagID}"] = tag.NewDeleteTag(o.context, o.TagDeleteTagHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/flags/{flagID}/variants/{variantID}"] = variant.NewDeleteVariant(o.context, o.VariantDeleteVariantHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/tags"] = tag.NewFindAllTags(o.context, o.TagFindAllTagsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/flags/{flagID}/segments/{segmentID}/constraints"] = constraint.NewFindConstraints(o.context, o.ConstraintFindConstraintsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/flags/{flagID}/segments/{segmentID}/distributions"] = distribution.NewFindDistributions(o.context, o.DistributionFindDistributionsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/flags"] = flag.NewFindFlags(o.context, o.FlagFindFlagsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/flags/{flagID}/segments"] = segment.NewFindSegments(o.context, o.SegmentFindSegmentsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/flags/{flagID}/tags"] = tag.NewFindTags(o.context, o.TagFindTagsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/flags/{flagID}/variants"] = variant.NewFindVariants(o.context, o.VariantFindVariantsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/export/eval_cache/json"] = export.NewGetExportEvalCacheJSON(o.context, o.ExportGetExportEvalCacheJSONHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/export/sqlite"] = export.NewGetExportSqlite(o.context, o.ExportGetExportSqliteHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/flags/{flagID}"] = flag.NewGetFlag(o.context, o.FlagGetFlagHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/flags/entity_types"] = flag.NewGetFlagEntityTypes(o.context, o.FlagGetFlagEntityTypesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/flags/{flagID}/snapshots"] = flag.NewGetFlagSnapshots(o.context, o.FlagGetFlagSnapshotsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/health"] = health.NewGetHealth(o.context, o.HealthGetHealthHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/evaluation"] = evaluation.NewPostEvaluation(o.context, o.EvaluationPostEvaluationHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/evaluation/batch"] = evaluation.NewPostEvaluationBatch(o.context, o.EvaluationPostEvaluationBatchHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/flags/{flagID}/segments/{segmentID}/constraints/{constraintID}"] = constraint.NewPutConstraint(o.context, o.ConstraintPutConstraintHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/flags/{flagID}/segments/{segmentID}/distributions"] = distribution.NewPutDistributions(o.context, o.DistributionPutDistributionsHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/flags/{flagID}"] = flag.NewPutFlag(o.context, o.FlagPutFlagHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/flags/{flagID}/segments/{segmentID}"] = segment.NewPutSegment(o.context, o.SegmentPutSegmentHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/flags/{flagID}/segments/reorder"] = segment.NewPutSegmentsReorder(o.context, o.SegmentPutSegmentsReorderHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/flags/{flagID}/variants/{variantID}"] = variant.NewPutVariant(o.context, o.VariantPutVariantHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/flags/{flagID}/restore"] = flag.NewRestoreFlag(o.context, o.FlagRestoreFlagHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/flags/{flagID}/enabled"] = flag.NewSetFlagEnabled(o.context, o.FlagSetFlagEnabledHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *FlagrAPI) Serve(builder middleware.Builder) http.Handler {
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
func (o *FlagrAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *FlagrAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *FlagrAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *FlagrAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
