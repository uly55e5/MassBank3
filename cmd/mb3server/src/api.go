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
	"context"
	"net/http"
)

// DefaultAPIRouter defines the required methods for binding the api requests to a responses for the DefaultAPI
// The DefaultAPIRouter implementation should parse necessary information from the http request,
// pass the data to a DefaultAPIServicer to perform the required actions, then write the service results to the http response.
type DefaultAPIRouter interface {
	GetRecords(http.ResponseWriter, *http.Request)
	GetSearchRecords(http.ResponseWriter, *http.Request)
	GetRecord(http.ResponseWriter, *http.Request)
	GetSimpleRecord(http.ResponseWriter, *http.Request)
	GetCount(http.ResponseWriter, *http.Request)
	GetBrowseOptions(http.ResponseWriter, *http.Request)
	GetMetadata(http.ResponseWriter, *http.Request)
	GetSimilarity(http.ResponseWriter, *http.Request)
}

// DefaultAPIServicer defines the api actions for the DefaultAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DefaultAPIServicer interface {
	GetRecords(context.Context, []string, string, []string, string, string, string, float64, string, []string, int32, []string, []string, string, string, []string) (ImplResponse, error)
	GetSearchRecords(context.Context, []string, string, []string, string, string, string, float64, string, []string, int32, []string, []string, string, string, []string) (ImplResponse, error)
	GetRecord(context.Context, string) (ImplResponse, error)
	GetSimpleRecord(context.Context, string) (ImplResponse, error)
	GetCount(context.Context) (ImplResponse, error)
	GetBrowseOptions(context.Context, []string, []string, string, []string) (ImplResponse, error)
	GetMetadata(context.Context) (ImplResponse, error)
	GetSimilarity(context.Context, []string, []string, int32) (ImplResponse, error)
}
