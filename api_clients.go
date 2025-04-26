package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// ClientsAPIService ClientsAPI service
type ClientsAPIService service

type ApiBulkClientsRequest struct {
	ctx               context.Context
	ApiService        *ClientsAPIService
	xAPITOKEN         *string
	xRequestedWith    *string
	genericBulkAction *GenericBulkAction
	index             *string
	include           *string
}

// The API token to be used for authentication
func (r ApiBulkClientsRequest) XAPITOKEN(xAPITOKEN string) ApiBulkClientsRequest {
	r.xAPITOKEN = &xAPITOKEN
	return r
}

// Used to send the XMLHttpRequest header
func (r ApiBulkClientsRequest) XRequestedWith(xRequestedWith string) ApiBulkClientsRequest {
	r.xRequestedWith = &xRequestedWith
	return r
}

// Bulk action array
func (r ApiBulkClientsRequest) GenericBulkAction(genericBulkAction GenericBulkAction) ApiBulkClientsRequest {
	r.genericBulkAction = &genericBulkAction
	return r
}

// Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;
func (r ApiBulkClientsRequest) Index(index string) ApiBulkClientsRequest {
	r.index = &index
	return r
}

// Include child relationships of the Client Object. ie ?include&#x3D;documents,system_logs  &#x60;&#x60;&#x60;html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  &#x60;&#x60;&#x60;  Usage:  &#x60;&#x60;&#x60;html /api/v1/clients?include&#x3D;contacts,documents,activities &#x60;&#x60;&#x60;
func (r ApiBulkClientsRequest) Include(include string) ApiBulkClientsRequest {
	r.include = &include
	return r
}

func (r ApiBulkClientsRequest) Execute() (*Client, *http.Response, error) {
	return r.ApiService.BulkClientsExecute(r)
}

/*
BulkClients Bulk client actions

##  POST /api/v1/clients/bulk

# Bulk actions allow to make changes to multiple clients in a single request the following actions are supported

- archive
- restore
- delete
- template _requires template,template_id properties also_
- assign_group _requires group_settings_id also_
- bulk_update _requires column,new_value also_

All of these actions require an array of client ids to perform the requested action on ie.

"ids":['id1','id2']

actions such as template, assign_group and bulk_update also require additional properties to be passed in the request

- template

The template bulk action allows the creation of a custom template using the provided template_id to be run against the array of clients provided.

- assign_group

# Allows setting multiple clients to a single group

- bulk_update

Allows updating certain columns on the client model in bulk. the current list of supported columns that can be updated archived_at:

- public_notes
- industry_id
- size_id
- country_id
- custom_value1
- custom_value2
- custom_value3
- custom_value4

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiBulkClientsRequest
*/
func (a *ClientsAPIService) BulkClients(ctx context.Context) ApiBulkClientsRequest {
	return ApiBulkClientsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Client
func (a *ClientsAPIService) BulkClientsExecute(r ApiBulkClientsRequest) (*Client, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *Client
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientsAPIService.BulkClients")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/clients/bulk"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAPITOKEN == nil {
		return localVarReturnValue, nil, reportError("xAPITOKEN is required and must be specified")
	}
	if r.xRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xRequestedWith is required and must be specified")
	}
	if r.genericBulkAction == nil {
		return localVarReturnValue, nil, reportError("genericBulkAction is required and must be specified")
	}

	if r.index != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "index", r.index, "form", "")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-API-TOKEN", r.xAPITOKEN, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Requested-With", r.xRequestedWith, "simple", "")
	// body params
	localVarPostBody = r.genericBulkAction
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-TOKEN"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthenticationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v AuthorizationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v RateLimiterError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiClientStatementRequest struct {
	ctx                    context.Context
	ApiService             *ClientsAPIService
	xAPITOKEN              *string
	xRequestedWith         *string
	clientStatementRequest *ClientStatementRequest
	index                  *string
	include                *string
}

// The API token to be used for authentication
func (r ApiClientStatementRequest) XAPITOKEN(xAPITOKEN string) ApiClientStatementRequest {
	r.xAPITOKEN = &xAPITOKEN
	return r
}

// Used to send the XMLHttpRequest header
func (r ApiClientStatementRequest) XRequestedWith(xRequestedWith string) ApiClientStatementRequest {
	r.xRequestedWith = &xRequestedWith
	return r
}

// Statement Options
func (r ApiClientStatementRequest) ClientStatementRequest(clientStatementRequest ClientStatementRequest) ApiClientStatementRequest {
	r.clientStatementRequest = &clientStatementRequest
	return r
}

// Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;
func (r ApiClientStatementRequest) Index(index string) ApiClientStatementRequest {
	r.index = &index
	return r
}

// Include child relationships of the Client Object. ie ?include&#x3D;documents,system_logs  &#x60;&#x60;&#x60;html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  &#x60;&#x60;&#x60;  Usage:  &#x60;&#x60;&#x60;html /api/v1/clients?include&#x3D;contacts,documents,activities &#x60;&#x60;&#x60;
func (r ApiClientStatementRequest) Include(include string) ApiClientStatementRequest {
	r.include = &include
	return r
}

func (r ApiClientStatementRequest) Execute() (*Client, *http.Response, error) {
	return r.ApiService.ClientStatementExecute(r)
}

/*
ClientStatement Client statement PDF

##  POST /api/v1/client_statement

# Return a PDF of the client statement

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiClientStatementRequest
*/
func (a *ClientsAPIService) ClientStatement(ctx context.Context) ApiClientStatementRequest {
	return ApiClientStatementRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Client
func (a *ClientsAPIService) ClientStatementExecute(r ApiClientStatementRequest) (*Client, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *Client
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientsAPIService.ClientStatement")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/client_statement"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAPITOKEN == nil {
		return localVarReturnValue, nil, reportError("xAPITOKEN is required and must be specified")
	}
	if r.xRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xRequestedWith is required and must be specified")
	}
	if r.clientStatementRequest == nil {
		return localVarReturnValue, nil, reportError("clientStatementRequest is required and must be specified")
	}

	if r.index != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "index", r.index, "form", "")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-API-TOKEN", r.xAPITOKEN, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Requested-With", r.xRequestedWith, "simple", "")
	// body params
	localVarPostBody = r.clientStatementRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-TOKEN"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthenticationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v AuthorizationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v RateLimiterError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteClientRequest struct {
	ctx            context.Context
	ApiService     *ClientsAPIService
	xAPITOKEN      *string
	xRequestedWith *string
	index          *string
	include        *string
	id             string
}

// The API token to be used for authentication
func (r ApiDeleteClientRequest) XAPITOKEN(xAPITOKEN string) ApiDeleteClientRequest {
	r.xAPITOKEN = &xAPITOKEN
	return r
}

// Used to send the XMLHttpRequest header
func (r ApiDeleteClientRequest) XRequestedWith(xRequestedWith string) ApiDeleteClientRequest {
	r.xRequestedWith = &xRequestedWith
	return r
}

// Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;
func (r ApiDeleteClientRequest) Index(index string) ApiDeleteClientRequest {
	r.index = &index
	return r
}

// Include child relationships of the Client Object. ie ?include&#x3D;documents,system_logs  &#x60;&#x60;&#x60;html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  &#x60;&#x60;&#x60;  Usage:  &#x60;&#x60;&#x60;html /api/v1/clients?include&#x3D;contacts,documents,activities &#x60;&#x60;&#x60;
func (r ApiDeleteClientRequest) Include(include string) ApiDeleteClientRequest {
	r.include = &include
	return r
}

func (r ApiDeleteClientRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteClientExecute(r)
}

/*
DeleteClient Delete client

##  DELETE /api/v1/clients/{id}

# Handles the deletion of a client by id

> ❗ Note
Deleting a client does not purge the client from the system. The delete action will remove the clients data from all
views in the application but keep it on file in case it needs to be restored.

A Client can be later restored reversing this action.

To permanently wipe a client and all of their records from the system, use the [/purge route](/#tag/clients/POST/api/v1/clients/{id}/purge)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The Client Hashed ID
	@return ApiDeleteClientRequest
*/
func (a *ClientsAPIService) DeleteClient(ctx context.Context, id string) ApiDeleteClientRequest {
	return ApiDeleteClientRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *ClientsAPIService) DeleteClientExecute(r ApiDeleteClientRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientsAPIService.DeleteClient")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/clients/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAPITOKEN == nil {
		return nil, reportError("xAPITOKEN is required and must be specified")
	}
	if r.xRequestedWith == nil {
		return nil, reportError("xRequestedWith is required and must be specified")
	}

	if r.index != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "index", r.index, "form", "")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-API-TOKEN", r.xAPITOKEN, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Requested-With", r.xRequestedWith, "simple", "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-TOKEN"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthenticationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v AuthorizationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v RateLimiterError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiEditClientRequest struct {
	ctx            context.Context
	ApiService     *ClientsAPIService
	xAPITOKEN      *string
	xRequestedWith *string
	index          *string
	include        *string
	id             string
}

// The API token to be used for authentication
func (r ApiEditClientRequest) XAPITOKEN(xAPITOKEN string) ApiEditClientRequest {
	r.xAPITOKEN = &xAPITOKEN
	return r
}

// Used to send the XMLHttpRequest header
func (r ApiEditClientRequest) XRequestedWith(xRequestedWith string) ApiEditClientRequest {
	r.xRequestedWith = &xRequestedWith
	return r
}

// Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;
func (r ApiEditClientRequest) Index(index string) ApiEditClientRequest {
	r.index = &index
	return r
}

// Include child relationships of the Client Object. ie ?include&#x3D;documents,system_logs  &#x60;&#x60;&#x60;html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  &#x60;&#x60;&#x60;  Usage:  &#x60;&#x60;&#x60;html /api/v1/clients?include&#x3D;contacts,documents,activities &#x60;&#x60;&#x60;
func (r ApiEditClientRequest) Include(include string) ApiEditClientRequest {
	r.include = &include
	return r
}

func (r ApiEditClientRequest) Execute() (*Client, *http.Response, error) {
	return r.ApiService.EditClientExecute(r)
}

/*
EditClient Edit Client

##  GET /api/v1/clients/{id}/edit
Displays a client by id, essentially an alias of the show route

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The Client Hashed ID
	@return ApiEditClientRequest
*/
func (a *ClientsAPIService) EditClient(ctx context.Context, id string) ApiEditClientRequest {
	return ApiEditClientRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return Client
func (a *ClientsAPIService) EditClientExecute(r ApiEditClientRequest) (*Client, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *Client
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientsAPIService.EditClient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/clients/{id}/edit"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAPITOKEN == nil {
		return localVarReturnValue, nil, reportError("xAPITOKEN is required and must be specified")
	}
	if r.xRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xRequestedWith is required and must be specified")
	}

	if r.index != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "index", r.index, "form", "")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-API-TOKEN", r.xAPITOKEN, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Requested-With", r.xRequestedWith, "simple", "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-TOKEN"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthenticationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v AuthorizationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v RateLimiterError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetClientsRequest struct {
	ctx                  context.Context
	ApiService           *ClientsAPIService
	xAPITOKEN            *string
	xRequestedWith       *string
	include              *string
	index                *string
	status               *string
	createdAt            *int32
	updatedAt            *int32
	isDeleted            *bool
	filterDeletedClients *string
	name                 *string
	balance              *string
	betweenBalance       *string
	email                *string
	idNumber             *string
	number               *string
	filter               *string
	sort                 *string
	group                *string
	clientId             *string
}

// The API token to be used for authentication
func (r ApiGetClientsRequest) XAPITOKEN(xAPITOKEN string) ApiGetClientsRequest {
	r.xAPITOKEN = &xAPITOKEN
	return r
}

// Used to send the XMLHttpRequest header
func (r ApiGetClientsRequest) XRequestedWith(xRequestedWith string) ApiGetClientsRequest {
	r.xRequestedWith = &xRequestedWith
	return r
}

// Include child relationships of the Client Object. ie ?include&#x3D;documents,system_logs  &#x60;&#x60;&#x60;html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  &#x60;&#x60;&#x60;  Usage:  &#x60;&#x60;&#x60;html /api/v1/clients?include&#x3D;contacts,documents,activities &#x60;&#x60;&#x60;
func (r ApiGetClientsRequest) Include(include string) ApiGetClientsRequest {
	r.include = &include
	return r
}

// Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;
func (r ApiGetClientsRequest) Index(index string) ApiGetClientsRequest {
	r.index = &index
	return r
}

// Filter the entity based on their status. ie active / archived / deleted. Format is a comma separated string with any of the following options:   - active - archived - deleted    &#x60;&#x60;&#x60;html GET /api/v1/invoices?status&#x3D;archived,deleted Returns only archived and deleted invoices &#x60;&#x60;&#x60;
func (r ApiGetClientsRequest) Status(status string) ApiGetClientsRequest {
	r.status = &status
	return r
}

// Filters the entity list by the created at timestamp. Parameter value can be a datetime string or unix timestamp  &#x60;&#x60;&#x60;html GET /api/v1/invoices?created_at&#x3D;2022-01-10 Returns entities created on January 10th, 2022 &#x60;&#x60;&#x60;
func (r ApiGetClientsRequest) CreatedAt(createdAt int32) ApiGetClientsRequest {
	r.createdAt = &createdAt
	return r
}

// Filters the entity list by the updated at timestamp. Parameter value can be a datetime string or unix timestamp  &#x60;&#x60;&#x60;html GET /api/v1/invoices?updated_at&#x3D;2022-01-10 Returns entities last updated on January 10th, 2022 &#x60;&#x60;&#x60;
func (r ApiGetClientsRequest) UpdatedAt(updatedAt int32) ApiGetClientsRequest {
	r.updatedAt = &updatedAt
	return r
}

// Filters the entity list by entities that have been deleted.  &#x60;&#x60;&#x60;html GET /api/v1/invoices?is_deleted&#x3D;true Returns only soft-deleted entities &#x60;&#x60;&#x60;
func (r ApiGetClientsRequest) IsDeleted(isDeleted bool) ApiGetClientsRequest {
	r.isDeleted = &isDeleted
	return r
}

// Filters the entity list and only returns entities for clients that have not been deleted  &#x60;&#x60;&#x60;html GET /api/v1/invoices?filter_deleted_clients&#x3D;true Returns only invoices for active (non-deleted) clients &#x60;&#x60;&#x60;
func (r ApiGetClientsRequest) FilterDeletedClients(filterDeletedClients string) ApiGetClientsRequest {
	r.filterDeletedClients = &filterDeletedClients
	return r
}

// Filter by client name    &#x60;&#x60;&#x60;html ?name&#x3D;bob &#x60;&#x60;&#x60;
func (r ApiGetClientsRequest) Name(name string) ApiGetClientsRequest {
	r.name = &name
	return r
}

// Filter by client balance, format uses an operator and value separated by a colon. lt,lte, gt, gte, eq  &#x60;&#x60;&#x60;html ?balance&#x3D;lt:10 &#x60;&#x60;&#x60;  ie all clients whose balance is less than 10
func (r ApiGetClientsRequest) Balance(balance string) ApiGetClientsRequest {
	r.balance = &balance
	return r
}

// Filter between client balances, format uses two values separated by a colon  &#x60;&#x60;&#x60;html ?between_balance&#x3D;10:100 &#x60;&#x60;&#x60;
func (r ApiGetClientsRequest) BetweenBalance(betweenBalance string) ApiGetClientsRequest {
	r.betweenBalance = &betweenBalance
	return r
}

// Filter by client email  &#x60;&#x60;&#x60;html ?email&#x3D;bob@gmail.com &#x60;&#x60;&#x60;
func (r ApiGetClientsRequest) Email(email string) ApiGetClientsRequest {
	r.email = &email
	return r
}

// Filter by client id_number  &#x60;&#x60;&#x60;html ?id_number&#x3D;0001 &#x60;&#x60;&#x60;
func (r ApiGetClientsRequest) IdNumber(idNumber string) ApiGetClientsRequest {
	r.idNumber = &idNumber
	return r
}

// Filter by client number  &#x60;&#x60;&#x60;html ?number&#x3D;0002 &#x60;&#x60;&#x60;
func (r ApiGetClientsRequest) Number(number string) ApiGetClientsRequest {
	r.number = &number
	return r
}

// Broad filter which targets multiple client columns:      &#x60;&#x60;&#x60;html     name,      id_number,      contact.first_name      contact.last_name,      contact.email,      contact.phone     custom_value1,     custom_value2,     custom_value3,     custom_value4,   &#x60;&#x60;&#x60;    &#x60;&#x60;&#x60;html   ?filter&#x3D;Bobby   &#x60;&#x60;&#x60;
func (r ApiGetClientsRequest) Filter(filter string) ApiGetClientsRequest {
	r.filter = &filter
	return r
}

// Returns the list sorted by column in ascending or descending order.  Ensure you use column | direction, ie.  &#x60;&#x60;&#x60;html   ?sort&#x3D;id|desc &#x60;&#x60;&#x60;
func (r ApiGetClientsRequest) Sort(sort string) ApiGetClientsRequest {
	r.sort = &sort
	return r
}

// Returns the list of clients assigned to {group_id}  &#x60;&#x60;&#x60;html   ?group&#x3D;X89sjd8 &#x60;&#x60;&#x60;
func (r ApiGetClientsRequest) Group(group string) ApiGetClientsRequest {
	r.group = &group
	return r
}

// Returns the list of clients with {client_id} - proxy call to retrieve a client_id wrapped in an array  &#x60;&#x60;&#x60;html   ?client_id&#x3D;X89sjd8 &#x60;&#x60;&#x60;
func (r ApiGetClientsRequest) ClientId(clientId string) ApiGetClientsRequest {
	r.clientId = &clientId
	return r
}

func (r ApiGetClientsRequest) Execute() (*Client, *http.Response, error) {
	return r.ApiService.GetClientsExecute(r)
}

/*
GetClients List clients

##  GET /api/v1/clients
When retrieving a list of clients you can also chain query parameters in order to filter the dataset that is returned. For example, you can send a request to the following URL to retrieve clients that have a balance greater than 1000

```
/api/v1/clients?balance=gt:1000
```

You can also sort the results by adding a sort parameter. The following example will sort the results by the client name in descending order

```
/api/v1/clients?sort=name|desc
```

You can also combine multiple filters together. The following example will return clients that have a balance greater than 1000 and are not deleted and have a name that starts with "Bob"

```
/api/v1/clients?balance=gt:1000&name=Bob*
```

If you wish to retrieve child relations, you can also combine the query parameter `?include=` with a comma separated list of relationships

```
/api/v1/clients?include=activities,ledger,system_logs'
```

The per_page and page variables allow pagination of the list of clients. The following example will return the second page of clients with 15 clients per page

```
/api/v1/clients?per_page=15&page=2
```

The default per_page value is 20.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetClientsRequest
*/
func (a *ClientsAPIService) GetClients(ctx context.Context) ApiGetClientsRequest {
	return ApiGetClientsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Client
func (a *ClientsAPIService) GetClientsExecute(r ApiGetClientsRequest) (*Client, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *Client
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientsAPIService.GetClients")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/clients"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAPITOKEN == nil {
		return localVarReturnValue, nil, reportError("xAPITOKEN is required and must be specified")
	}
	if r.xRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xRequestedWith is required and must be specified")
	}

	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "form", "")
	}
	if r.index != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "index", r.index, "form", "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	if r.createdAt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "created_at", r.createdAt, "form", "")
	}
	if r.updatedAt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "updated_at", r.updatedAt, "form", "")
	}
	if r.isDeleted != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "is_deleted", r.isDeleted, "form", "")
	}
	if r.filterDeletedClients != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter_deleted_clients", r.filterDeletedClients, "form", "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "form", "")
	}
	if r.balance != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "balance", r.balance, "form", "")
	}
	if r.betweenBalance != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "between_balance", r.betweenBalance, "form", "")
	}
	if r.email != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "form", "")
	}
	if r.idNumber != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "id_number", r.idNumber, "form", "")
	}
	if r.number != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "number", r.number, "form", "")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "form", "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "")
	}
	if r.group != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "group", r.group, "form", "")
	}
	if r.clientId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "client_id", r.clientId, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-API-TOKEN", r.xAPITOKEN, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Requested-With", r.xRequestedWith, "simple", "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-TOKEN"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthenticationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v AuthorizationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v RateLimiterError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetClientsCreateRequest struct {
	ctx            context.Context
	ApiService     *ClientsAPIService
	xAPITOKEN      *string
	xRequestedWith *string
	index          *string
	include        *string
}

// The API token to be used for authentication
func (r ApiGetClientsCreateRequest) XAPITOKEN(xAPITOKEN string) ApiGetClientsCreateRequest {
	r.xAPITOKEN = &xAPITOKEN
	return r
}

// Used to send the XMLHttpRequest header
func (r ApiGetClientsCreateRequest) XRequestedWith(xRequestedWith string) ApiGetClientsCreateRequest {
	r.xRequestedWith = &xRequestedWith
	return r
}

// Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;
func (r ApiGetClientsCreateRequest) Index(index string) ApiGetClientsCreateRequest {
	r.index = &index
	return r
}

// Include child relationships of the Client Object. ie ?include&#x3D;documents,system_logs  &#x60;&#x60;&#x60;html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  &#x60;&#x60;&#x60;  Usage:  &#x60;&#x60;&#x60;html /api/v1/clients?include&#x3D;contacts,documents,activities &#x60;&#x60;&#x60;
func (r ApiGetClientsCreateRequest) Include(include string) ApiGetClientsCreateRequest {
	r.include = &include
	return r
}

func (r ApiGetClientsCreateRequest) Execute() (*Client, *http.Response, error) {
	return r.ApiService.GetClientsCreateExecute(r)
}

/*
GetClientsCreate Blank Client

##  GET /api/v1/clients/create

# Returns a blank object with default values

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetClientsCreateRequest
*/
func (a *ClientsAPIService) GetClientsCreate(ctx context.Context) ApiGetClientsCreateRequest {
	return ApiGetClientsCreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Client
func (a *ClientsAPIService) GetClientsCreateExecute(r ApiGetClientsCreateRequest) (*Client, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *Client
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientsAPIService.GetClientsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/clients/create"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAPITOKEN == nil {
		return localVarReturnValue, nil, reportError("xAPITOKEN is required and must be specified")
	}
	if r.xRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xRequestedWith is required and must be specified")
	}

	if r.index != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "index", r.index, "form", "")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-API-TOKEN", r.xAPITOKEN, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Requested-With", r.xRequestedWith, "simple", "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-TOKEN"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthenticationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v AuthorizationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v RateLimiterError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiMergeClientRequest struct {
	ctx                     context.Context
	ApiService              *ClientsAPIService
	xAPITOKEN               *string
	xRequestedWith          *string
	xAPIPASSWORD            *string
	index                   *string
	include                 *string
	id                      string
	mergeableClientHashedId string
}

// The API token to be used for authentication
func (r ApiMergeClientRequest) XAPITOKEN(xAPITOKEN string) ApiMergeClientRequest {
	r.xAPITOKEN = &xAPITOKEN
	return r
}

// Used to send the XMLHttpRequest header
func (r ApiMergeClientRequest) XRequestedWith(xRequestedWith string) ApiMergeClientRequest {
	r.xRequestedWith = &xRequestedWith
	return r
}

// The login password when challenged on certain protected routes
func (r ApiMergeClientRequest) XAPIPASSWORD(xAPIPASSWORD string) ApiMergeClientRequest {
	r.xAPIPASSWORD = &xAPIPASSWORD
	return r
}

// Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;
func (r ApiMergeClientRequest) Index(index string) ApiMergeClientRequest {
	r.index = &index
	return r
}

// Include child relationships of the Client Object. ie ?include&#x3D;documents,system_logs  &#x60;&#x60;&#x60;html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  &#x60;&#x60;&#x60;  Usage:  &#x60;&#x60;&#x60;html /api/v1/clients?include&#x3D;contacts,documents,activities &#x60;&#x60;&#x60;
func (r ApiMergeClientRequest) Include(include string) ApiMergeClientRequest {
	r.include = &include
	return r
}

func (r ApiMergeClientRequest) Execute() (*http.Response, error) {
	return r.ApiService.MergeClientExecute(r)
}

/*
MergeClient Merge client

##  POST /api/v1/clients/{id}/{mergeable_client_hashed_id}/merge

# Handles merging 2 clients

The id parameter is the client that will be the primary client after the merge has completed.

The mergeable_client_hashed_id is the client that will be merged into the primary client, this clients records will be updated and associated with the primary client.

> 🚨 **Important**
This action requires elevated permissions, please note the X-API-PASSWORD header requirements for this route.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The Client Hashed ID
	@param mergeableClientHashedId The Mergeable Client Hashed ID
	@return ApiMergeClientRequest
*/
func (a *ClientsAPIService) MergeClient(ctx context.Context, id string, mergeableClientHashedId string) ApiMergeClientRequest {
	return ApiMergeClientRequest{
		ApiService:              a,
		ctx:                     ctx,
		id:                      id,
		mergeableClientHashedId: mergeableClientHashedId,
	}
}

// Execute executes the request
func (a *ClientsAPIService) MergeClientExecute(r ApiMergeClientRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientsAPIService.MergeClient")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/clients/{id}/{mergeable_client_hashed_id}/merge"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"mergeable_client_hashed_id"+"}", url.PathEscape(parameterValueToString(r.mergeableClientHashedId, "mergeableClientHashedId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAPITOKEN == nil {
		return nil, reportError("xAPITOKEN is required and must be specified")
	}
	if r.xRequestedWith == nil {
		return nil, reportError("xRequestedWith is required and must be specified")
	}
	if r.xAPIPASSWORD == nil {
		return nil, reportError("xAPIPASSWORD is required and must be specified")
	}

	if r.index != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "index", r.index, "form", "")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-API-TOKEN", r.xAPITOKEN, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Requested-With", r.xRequestedWith, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-API-PASSWORD", r.xAPIPASSWORD, "simple", "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-TOKEN"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthenticationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v AuthorizationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v RateLimiterError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiPurgeClientRequest struct {
	ctx            context.Context
	ApiService     *ClientsAPIService
	xAPITOKEN      *string
	xRequestedWith *string
	xAPIPASSWORD   *string
	id             string
}

// The API token to be used for authentication
func (r ApiPurgeClientRequest) XAPITOKEN(xAPITOKEN string) ApiPurgeClientRequest {
	r.xAPITOKEN = &xAPITOKEN
	return r
}

// Used to send the XMLHttpRequest header
func (r ApiPurgeClientRequest) XRequestedWith(xRequestedWith string) ApiPurgeClientRequest {
	r.xRequestedWith = &xRequestedWith
	return r
}

// The login password when challenged on certain protected routes
func (r ApiPurgeClientRequest) XAPIPASSWORD(xAPIPASSWORD string) ApiPurgeClientRequest {
	r.xAPIPASSWORD = &xAPIPASSWORD
	return r
}

func (r ApiPurgeClientRequest) Execute() (*http.Response, error) {
	return r.ApiService.PurgeClientExecute(r)
}

/*
PurgeClient Purge client

##  POST /api/v1/clients/{id}/purge

Handles purging a client.

> ❗ Note
This is a destructive action.
This action will remove all data associated with the client and cannot be undone.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The Client Hashed ID
	@return ApiPurgeClientRequest
*/
func (a *ClientsAPIService) PurgeClient(ctx context.Context, id string) ApiPurgeClientRequest {
	return ApiPurgeClientRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *ClientsAPIService) PurgeClientExecute(r ApiPurgeClientRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientsAPIService.PurgeClient")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/clients/{id}/purge"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAPITOKEN == nil {
		return nil, reportError("xAPITOKEN is required and must be specified")
	}
	if r.xRequestedWith == nil {
		return nil, reportError("xRequestedWith is required and must be specified")
	}
	if r.xAPIPASSWORD == nil {
		return nil, reportError("xAPIPASSWORD is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-API-TOKEN", r.xAPITOKEN, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Requested-With", r.xRequestedWith, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-API-PASSWORD", r.xAPIPASSWORD, "simple", "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-TOKEN"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthenticationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v AuthorizationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v RateLimiterError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiReactivateEmailRequest struct {
	ctx            context.Context
	ApiService     *ClientsAPIService
	xAPITOKEN      *string
	xRequestedWith *string
	index          *string
	include        *string
	bounceId       string
}

// The API token to be used for authentication
func (r ApiReactivateEmailRequest) XAPITOKEN(xAPITOKEN string) ApiReactivateEmailRequest {
	r.xAPITOKEN = &xAPITOKEN
	return r
}

// Used to send the XMLHttpRequest header
func (r ApiReactivateEmailRequest) XRequestedWith(xRequestedWith string) ApiReactivateEmailRequest {
	r.xRequestedWith = &xRequestedWith
	return r
}

// Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;
func (r ApiReactivateEmailRequest) Index(index string) ApiReactivateEmailRequest {
	r.index = &index
	return r
}

// Include child relationships of the Client Object. ie ?include&#x3D;documents,system_logs  &#x60;&#x60;&#x60;html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  &#x60;&#x60;&#x60;  Usage:  &#x60;&#x60;&#x60;html /api/v1/clients?include&#x3D;contacts,documents,activities &#x60;&#x60;&#x60;
func (r ApiReactivateEmailRequest) Include(include string) ApiReactivateEmailRequest {
	r.include = &include
	return r
}

func (r ApiReactivateEmailRequest) Execute() (*http.Response, error) {
	return r.ApiService.ReactivateEmailExecute(r)
}

/*
ReactivateEmail Removes email suppression of a user in the system

##  POST /api/v1/reactivate_email/{bounce_id}

Emails are suppressed by PostMark, when they receive a Hard bounce / Spam Complaint. This endpoint allows you to remove the suppression and send emails to the user again.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param bounceId The postmark Bounce ID reference
	@return ApiReactivateEmailRequest
*/
func (a *ClientsAPIService) ReactivateEmail(ctx context.Context, bounceId string) ApiReactivateEmailRequest {
	return ApiReactivateEmailRequest{
		ApiService: a,
		ctx:        ctx,
		bounceId:   bounceId,
	}
}

// Execute executes the request
func (a *ClientsAPIService) ReactivateEmailExecute(r ApiReactivateEmailRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientsAPIService.ReactivateEmail")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/reactivate_email/{bounce_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"bounce_id"+"}", url.PathEscape(parameterValueToString(r.bounceId, "bounceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAPITOKEN == nil {
		return nil, reportError("xAPITOKEN is required and must be specified")
	}
	if r.xRequestedWith == nil {
		return nil, reportError("xRequestedWith is required and must be specified")
	}

	if r.index != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "index", r.index, "form", "")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-API-TOKEN", r.xAPITOKEN, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Requested-With", r.xRequestedWith, "simple", "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-TOKEN"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthenticationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v AuthorizationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v RateLimiterError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiShowClientRequest struct {
	ctx            context.Context
	ApiService     *ClientsAPIService
	xAPITOKEN      *string
	xRequestedWith *string
	index          *string
	include        *string
	id             string
}

// The API token to be used for authentication
func (r ApiShowClientRequest) XAPITOKEN(xAPITOKEN string) ApiShowClientRequest {
	r.xAPITOKEN = &xAPITOKEN
	return r
}

// Used to send the XMLHttpRequest header
func (r ApiShowClientRequest) XRequestedWith(xRequestedWith string) ApiShowClientRequest {
	r.xRequestedWith = &xRequestedWith
	return r
}

// Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;
func (r ApiShowClientRequest) Index(index string) ApiShowClientRequest {
	r.index = &index
	return r
}

// Include child relationships of the Client Object. ie ?include&#x3D;documents,system_logs  &#x60;&#x60;&#x60;html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  &#x60;&#x60;&#x60;  Usage:  &#x60;&#x60;&#x60;html /api/v1/clients?include&#x3D;contacts,documents,activities &#x60;&#x60;&#x60;
func (r ApiShowClientRequest) Include(include string) ApiShowClientRequest {
	r.include = &include
	return r
}

func (r ApiShowClientRequest) Execute() (*Client, *http.Response, error) {
	return r.ApiService.ShowClientExecute(r)
}

/*
ShowClient Show client

##  GET /api/v1/clients/{id}

# Displays a client by id

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The Client Hashed ID
	@return ApiShowClientRequest
*/
func (a *ClientsAPIService) ShowClient(ctx context.Context, id string) ApiShowClientRequest {
	return ApiShowClientRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return Client
func (a *ClientsAPIService) ShowClientExecute(r ApiShowClientRequest) (*Client, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *Client
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientsAPIService.ShowClient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/clients/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAPITOKEN == nil {
		return localVarReturnValue, nil, reportError("xAPITOKEN is required and must be specified")
	}
	if r.xRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xRequestedWith is required and must be specified")
	}

	if r.index != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "index", r.index, "form", "")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-API-TOKEN", r.xAPITOKEN, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Requested-With", r.xRequestedWith, "simple", "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-TOKEN"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthenticationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v AuthorizationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v RateLimiterError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiStoreClientRequest struct {
	ctx            context.Context
	ApiService     *ClientsAPIService
	xAPITOKEN      *string
	xRequestedWith *string
	clientRequest  *ClientRequest
	index          *string
	include        *string
}

// The API token to be used for authentication
func (r ApiStoreClientRequest) XAPITOKEN(xAPITOKEN string) ApiStoreClientRequest {
	r.xAPITOKEN = &xAPITOKEN
	return r
}

// Used to send the XMLHttpRequest header
func (r ApiStoreClientRequest) XRequestedWith(xRequestedWith string) ApiStoreClientRequest {
	r.xRequestedWith = &xRequestedWith
	return r
}

func (r ApiStoreClientRequest) ClientRequest(clientRequest ClientRequest) ApiStoreClientRequest {
	r.clientRequest = &clientRequest
	return r
}

// Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;
func (r ApiStoreClientRequest) Index(index string) ApiStoreClientRequest {
	r.index = &index
	return r
}

// Include child relationships of the Client Object. ie ?include&#x3D;documents,system_logs  &#x60;&#x60;&#x60;html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  &#x60;&#x60;&#x60;  Usage:  &#x60;&#x60;&#x60;html /api/v1/clients?include&#x3D;contacts,documents,activities &#x60;&#x60;&#x60;
func (r ApiStoreClientRequest) Include(include string) ApiStoreClientRequest {
	r.include = &include
	return r
}

func (r ApiStoreClientRequest) Execute() (*Client, *http.Response, error) {
	return r.ApiService.StoreClientExecute(r)
}

/*
StoreClient Create client

##  POST /api/v1/clients
Adds a client to a company

> 🚨 Important
When creating (or updating) a client you must include the child contacts with all mutating requests. Client contacts cannot be modified in isolation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiStoreClientRequest
*/
func (a *ClientsAPIService) StoreClient(ctx context.Context) ApiStoreClientRequest {
	return ApiStoreClientRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Client
func (a *ClientsAPIService) StoreClientExecute(r ApiStoreClientRequest) (*Client, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *Client
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientsAPIService.StoreClient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/clients"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAPITOKEN == nil {
		return localVarReturnValue, nil, reportError("xAPITOKEN is required and must be specified")
	}
	if r.xRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xRequestedWith is required and must be specified")
	}
	if r.clientRequest == nil {
		return localVarReturnValue, nil, reportError("clientRequest is required and must be specified")
	}

	if r.index != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "index", r.index, "form", "")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-API-TOKEN", r.xAPITOKEN, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Requested-With", r.xRequestedWith, "simple", "")
	// body params
	localVarPostBody = r.clientRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-TOKEN"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthenticationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v AuthorizationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v RateLimiterError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateClientRequest struct {
	ctx            context.Context
	ApiService     *ClientsAPIService
	xAPITOKEN      *string
	xRequestedWith *string
	clientRequest  *ClientRequest
	index          *string
	include        *string
	id             string
}

// The API token to be used for authentication
func (r ApiUpdateClientRequest) XAPITOKEN(xAPITOKEN string) ApiUpdateClientRequest {
	r.xAPITOKEN = &xAPITOKEN
	return r
}

// Used to send the XMLHttpRequest header
func (r ApiUpdateClientRequest) XRequestedWith(xRequestedWith string) ApiUpdateClientRequest {
	r.xRequestedWith = &xRequestedWith
	return r
}

// Client object that needs to be updated
func (r ApiUpdateClientRequest) ClientRequest(clientRequest ClientRequest) ApiUpdateClientRequest {
	r.clientRequest = &clientRequest
	return r
}

// Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;
func (r ApiUpdateClientRequest) Index(index string) ApiUpdateClientRequest {
	r.index = &index
	return r
}

// Include child relationships of the Client Object. ie ?include&#x3D;documents,system_logs  &#x60;&#x60;&#x60;html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  &#x60;&#x60;&#x60;  Usage:  &#x60;&#x60;&#x60;html /api/v1/clients?include&#x3D;contacts,documents,activities &#x60;&#x60;&#x60;
func (r ApiUpdateClientRequest) Include(include string) ApiUpdateClientRequest {
	r.include = &include
	return r
}

func (r ApiUpdateClientRequest) Execute() (*Client, *http.Response, error) {
	return r.ApiService.UpdateClientExecute(r)
}

/*
UpdateClient Update client

##  PUT /api/v1/clients/{id}

# Handles the updating of a client by id

> 🚨 Important
When creating (or updating) a client you must include the child contacts with all mutating requests. Client contacts cannot be modified in isolation.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The Client Hashed ID
	@return ApiUpdateClientRequest
*/
func (a *ClientsAPIService) UpdateClient(ctx context.Context, id string) ApiUpdateClientRequest {
	return ApiUpdateClientRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return Client
func (a *ClientsAPIService) UpdateClientExecute(r ApiUpdateClientRequest) (*Client, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *Client
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientsAPIService.UpdateClient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/clients/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAPITOKEN == nil {
		return localVarReturnValue, nil, reportError("xAPITOKEN is required and must be specified")
	}
	if r.xRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xRequestedWith is required and must be specified")
	}
	if r.clientRequest == nil {
		return localVarReturnValue, nil, reportError("clientRequest is required and must be specified")
	}

	if r.index != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "index", r.index, "form", "")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-API-TOKEN", r.xAPITOKEN, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Requested-With", r.xRequestedWith, "simple", "")
	// body params
	localVarPostBody = r.clientRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-TOKEN"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthenticationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v AuthorizationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v RateLimiterError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateClientTaxDataRequest struct {
	ctx            context.Context
	ApiService     *ClientsAPIService
	xAPITOKEN      *string
	xRequestedWith *string
	index          *string
	include        *string
	client         string
}

// The API token to be used for authentication
func (r ApiUpdateClientTaxDataRequest) XAPITOKEN(xAPITOKEN string) ApiUpdateClientTaxDataRequest {
	r.xAPITOKEN = &xAPITOKEN
	return r
}

// Used to send the XMLHttpRequest header
func (r ApiUpdateClientTaxDataRequest) XRequestedWith(xRequestedWith string) ApiUpdateClientTaxDataRequest {
	r.xRequestedWith = &xRequestedWith
	return r
}

// Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;
func (r ApiUpdateClientTaxDataRequest) Index(index string) ApiUpdateClientTaxDataRequest {
	r.index = &index
	return r
}

// Include child relationships of the Client Object. ie ?include&#x3D;documents,system_logs  &#x60;&#x60;&#x60;html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  &#x60;&#x60;&#x60;  Usage:  &#x60;&#x60;&#x60;html /api/v1/clients?include&#x3D;contacts,documents,activities &#x60;&#x60;&#x60;
func (r ApiUpdateClientTaxDataRequest) Include(include string) ApiUpdateClientTaxDataRequest {
	r.include = &include
	return r
}

func (r ApiUpdateClientTaxDataRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateClientTaxDataExecute(r)
}

/*
UpdateClientTaxData Update tax data

##  POST /api/v1/clients/{client}/updateTaxData

# Updates the clients tax data - if their address has changed

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param client The Client Hashed ID reference
	@return ApiUpdateClientTaxDataRequest
*/
func (a *ClientsAPIService) UpdateClientTaxData(ctx context.Context, client string) ApiUpdateClientTaxDataRequest {
	return ApiUpdateClientTaxDataRequest{
		ApiService: a,
		ctx:        ctx,
		client:     client,
	}
}

// Execute executes the request
func (a *ClientsAPIService) UpdateClientTaxDataExecute(r ApiUpdateClientTaxDataRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientsAPIService.UpdateClientTaxData")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/clients/{client}/updateTaxData"
	localVarPath = strings.Replace(localVarPath, "{"+"client"+"}", url.PathEscape(parameterValueToString(r.client, "client")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAPITOKEN == nil {
		return nil, reportError("xAPITOKEN is required and must be specified")
	}
	if r.xRequestedWith == nil {
		return nil, reportError("xRequestedWith is required and must be specified")
	}

	if r.index != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "index", r.index, "form", "")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-API-TOKEN", r.xAPITOKEN, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Requested-With", r.xRequestedWith, "simple", "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-TOKEN"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthenticationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v AuthorizationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v RateLimiterError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUploadClientRequest struct {
	ctx            context.Context
	ApiService     *ClientsAPIService
	xAPITOKEN      *string
	xRequestedWith *string
	id             string
	index          *string
	include        *string
	method         *string
	documents      []*os.File
}

// The API token to be used for authentication
func (r ApiUploadClientRequest) XAPITOKEN(xAPITOKEN string) ApiUploadClientRequest {
	r.xAPITOKEN = &xAPITOKEN
	return r
}

// Used to send the XMLHttpRequest header
func (r ApiUploadClientRequest) XRequestedWith(xRequestedWith string) ApiUploadClientRequest {
	r.xRequestedWith = &xRequestedWith
	return r
}

// Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;
func (r ApiUploadClientRequest) Index(index string) ApiUploadClientRequest {
	r.index = &index
	return r
}

// Include child relationships of the Client Object. ie ?include&#x3D;documents,system_logs  &#x60;&#x60;&#x60;html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  &#x60;&#x60;&#x60;  Usage:  &#x60;&#x60;&#x60;html /api/v1/clients?include&#x3D;contacts,documents,activities &#x60;&#x60;&#x60;
func (r ApiUploadClientRequest) Include(include string) ApiUploadClientRequest {
	r.include = &include
	return r
}

func (r ApiUploadClientRequest) Method(method string) ApiUploadClientRequest {
	r.method = &method
	return r
}

// Array of files to upload. The files should be sent with the key name &#39;documents[]&#39;.  Supported file types: - PDF (.pdf) - Word (.doc, .docx) - Excel (.xls, .xlsx) - Images (.jpg, .jpeg, .png) - Text (.txt)  Maximum file size: 20MB per file
func (r ApiUploadClientRequest) Documents(documents []*os.File) ApiUploadClientRequest {
	r.documents = documents
	return r
}

func (r ApiUploadClientRequest) Execute() (*Client, *http.Response, error) {
	return r.ApiService.UploadClientExecute(r)
}

/*
UploadClient Add client document

##  POST /api/v1/clients/{id}/upload

Handles the uploading of a document to a client, please note due to a quirk in REST you will need to use a _method parameter with value of POST

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The Client Hashed ID
	@return ApiUploadClientRequest
*/
func (a *ClientsAPIService) UploadClient(ctx context.Context, id string) ApiUploadClientRequest {
	return ApiUploadClientRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return Client
func (a *ClientsAPIService) UploadClientExecute(r ApiUploadClientRequest) (*Client, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *Client
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClientsAPIService.UploadClient")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/clients/{id}/upload"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAPITOKEN == nil {
		return localVarReturnValue, nil, reportError("xAPITOKEN is required and must be specified")
	}
	if r.xRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xRequestedWith is required and must be specified")
	}

	if r.index != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "index", r.index, "form", "")
	}
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-API-TOKEN", r.xAPITOKEN, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Requested-With", r.xRequestedWith, "simple", "")
	if r.method != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "_method", r.method, "", "")
	}
	var documentsLocalVarFormFileName string
	var documentsLocalVarFileName string
	var documentsLocalVarFileBytes []byte

	documentsLocalVarFormFileName = "documents[]"
	documentsLocalVarFile := r.documents

	if documentsLocalVarFile != nil {
		// loop through the array to prepare multiple files upload
		for _, documentsLocalVarFileValue := range documentsLocalVarFile {
			fbs, _ := io.ReadAll(documentsLocalVarFileValue)

			documentsLocalVarFileBytes = fbs
			documentsLocalVarFileName = documentsLocalVarFileValue.Name()
			documentsLocalVarFileValue.Close()
			formFiles = append(formFiles, formFile{fileBytes: documentsLocalVarFileBytes, fileName: documentsLocalVarFileName, formFileName: documentsLocalVarFormFileName})
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-TOKEN"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v AuthenticationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v AuthorizationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v RateLimiterError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
