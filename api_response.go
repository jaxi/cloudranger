/*
 * CloudRanger API
 *
 *  # Welcome to CloudRanger  Here are some instructions on  how to use the API  ## Authentication and Authorization   * Access to the API is controlled by your API key generated in the CloudRanger dashboard and token  * The token is returned by calling the /authorize endpoint and supplying the x-api-key header  * All other calls use the x-api-key header and the Authorzation header with Bearer followed by your token 
 *
 * API version: 2018-05
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cloudranger

import (
	"net/http"
)

type APIResponse struct {
	*http.Response `json:"-"`
	Message string `json:"message,omitempty"`
	// Operation is the name of the swagger operation.
	Operation       string `json:"operation,omitempty"`
	// RequestURL is the request URL. This value is always available, even if the
	// embedded *http.Response is nil.
	RequestURL      string `json:"url,omitempty"`
	// Method is the HTTP method used for the request.  This value is always
	// available, even if the embedded *http.Response is nil.
	Method          string `json:"method,omitempty"`
	// Payload holds the contents of the response body (which may be nil or empty).
	// This is provided here as the raw response.Body() reader will have already
	// been drained.
	Payload []byte `json:"-"`
}

func NewAPIResponse(r *http.Response) *APIResponse {

	response := &APIResponse{Response: r}
	return response
}

func NewAPIResponseWithError(errorMessage string) *APIResponse {

	response := &APIResponse{Message: errorMessage}
	return response
}