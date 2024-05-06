/*
 * MassBank3 API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package mb3server

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service      DefaultApiServicer
	errorHandler ErrorHandler
}

// DefaultApiOption for how the controller is set up.
type DefaultApiOption func(*DefaultApiController)

// WithDefaultApiErrorHandler inject ErrorHandler into controller
func WithDefaultApiErrorHandler(h ErrorHandler) DefaultApiOption {
	return func(c *DefaultApiController) {
		c.errorHandler = h
	}
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(s DefaultApiServicer, opts ...DefaultApiOption) Router {
	controller := &DefaultApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the DefaultApiController
func (c *DefaultApiController) Routes() Routes {
	return Routes{
		{
			"GetBrowseOptions",
			strings.ToUpper("Get"),
			"/v1/filter/browse",
			c.GetBrowseOptions,
		},
		{
			"GetCount",
			strings.ToUpper("Get"),
			"/v1/records/count",
			c.GetCount,
		},
		{
			"GetFilterOptions",
			strings.ToUpper("Get"),
			"/v1/filter/options",
			c.GetFilterOptions,
		},
		{
			"GetMetadata",
			strings.ToUpper("Get"),
			"/v1/metadata",
			c.GetMetadata,
		},
		{
			"GetRecord",
			strings.ToUpper("Get"),
			"/v1/records/{accession}",
			c.GetRecord,
		},
		{
			"GetRecords",
			strings.ToUpper("Get"),
			"/v1/records",
			c.GetRecords,
		},
		{
			"GetSVG",
			strings.ToUpper("Get"),
			"/v1/records/{accession}/svg",
			c.GetSVG,
		},
	}
}

// GetBrowseOptions - get browse options
func (c *DefaultApiController) GetBrowseOptions(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	instrumentTypeParam := strings.Split(query.Get("instrument_type"), ",")
	msTypeParam := strings.Split(query.Get("ms_type"), ",")
	ionModeParam := query.Get("ion_mode")
	contributorParam := strings.Split(query.Get("contributor"), ",")
	result, err := c.service.GetBrowseOptions(r.Context(), instrumentTypeParam, msTypeParam, ionModeParam, contributorParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetCount - The number of all records
func (c *DefaultApiController) GetCount(w http.ResponseWriter, r *http.Request) {

	result, err := c.service.GetCount(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetFilterOptions - get filter options
func (c *DefaultApiController) GetFilterOptions(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetFilterOptions(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetMetadata - get massbank metadata
func (c *DefaultApiController) GetMetadata(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetMetadata(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetRecord - Get a MassBank record
func (c *DefaultApiController) GetRecord(w http.ResponseWriter, r *http.Request) {
	accessionParam := chi.URLParam(r, "accession")

	result, err := c.service.GetRecord(r.Context(), accessionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetRecords - Get a list of records
func (c *DefaultApiController) GetRecords(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	instrumentTypeParam := strings.Split(query.Get("instrument_type"), ",")
	splashParam := query.Get("splash")
	msTypeParam := strings.Split(query.Get("ms_type"), ",")
	ionModeParam := query.Get("ion-mode")
	compoundNameParam := query.Get("compound_name")
	exactMassParam := query.Get("exact_mass")
	massToleranceParam, err := parseFloat64Parameter(query.Get("mass_tolerance"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	formulaParam := query.Get("formula")
	peaksParam := strings.Split(query.Get("peaks"), ",")
	intensityParam, err := parseInt32Parameter(query.Get("intensity"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	peakDifferencesParam := strings.Split(query.Get("peak_differences"), ",")
	peakListParam := strings.Split(query.Get("peak_list"), ",")
	limitParam, err := parseInt32Parameter(query.Get("limit"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	pageParam, err := parseInt32Parameter(query.Get("page"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	intensityCutoffParam, err := parseInt32Parameter(query.Get("intensity_cutoff"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	inchiKeyParam := query.Get("inchi_key")
	contributorParam := strings.Split(query.Get("contributor"), ",")
	result, err := c.service.GetRecords(r.Context(), instrumentTypeParam, splashParam, msTypeParam, ionModeParam, compoundNameParam, exactMassParam, massToleranceParam, formulaParam, peaksParam, intensityParam, peakDifferencesParam, peakListParam, limitParam, pageParam, intensityCutoffParam, inchiKeyParam, contributorParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetSVG - The SVG image for an accession
func (c *DefaultApiController) GetSVG(w http.ResponseWriter, r *http.Request) {
	accessionParam := chi.URLParam(r, "accession")

	result, err := c.service.GetSVG(r.Context(), accessionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}
