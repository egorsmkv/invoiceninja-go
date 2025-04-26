package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// RefreshAPIService RefreshAPI service
type RefreshAPIService service

type ApiRefreshRequest struct {
	ctx            context.Context
	ApiService     *RefreshAPIService
	updatedAt      *float32
	xAPITOKEN      *string
	xRequestedWith *string
	include        *string
	includeStatic  *string
	clearCache     *string
}

// The unix timestamp from which the refreshed data should be sent from, if no value is passed the system will assume you require all data.
func (r ApiRefreshRequest) UpdatedAt(updatedAt float32) ApiRefreshRequest {
	r.updatedAt = &updatedAt
	return r
}

// The API token to be used for authentication
func (r ApiRefreshRequest) XAPITOKEN(xAPITOKEN string) ApiRefreshRequest {
	r.xAPITOKEN = &xAPITOKEN
	return r
}

// Used to send the XMLHttpRequest header
func (r ApiRefreshRequest) XRequestedWith(xRequestedWith string) ApiRefreshRequest {
	r.xRequestedWith = &xRequestedWith
	return r
}

// Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes
func (r ApiRefreshRequest) Include(include string) ApiRefreshRequest {
	r.include = &include
	return r
}

// This include will return the full set of static variables used in the application including:   - Currencies   - Countries    - Languages   - Payment Types   - Email Templatees   - Industries
func (r ApiRefreshRequest) IncludeStatic(includeStatic string) ApiRefreshRequest {
	r.includeStatic = &includeStatic
	return r
}

// Clears cache  Clears (and rebuilds) the static variable cache.
func (r ApiRefreshRequest) ClearCache(clearCache string) ApiRefreshRequest {
	r.clearCache = &clearCache
	return r
}

func (r ApiRefreshRequest) Execute() (*CompanyUser, *http.Response, error) {
	return r.ApiService.RefreshExecute(r)
}

/*
Refresh Refresh data by timestamp

Refreshes the dataset.

This endpoint can be used if you only need to access the most recent data from a certain point in time. For example, if you only want to retrieve The
most recent data from the last time you accessed the system, you would pass the query parameter ?updated_at=1676173763. (unix timestamp)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiRefreshRequest
*/
func (a *RefreshAPIService) Refresh(ctx context.Context) ApiRefreshRequest {
	return ApiRefreshRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CompanyUser
func (a *RefreshAPIService) RefreshExecute(r ApiRefreshRequest) (*CompanyUser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *CompanyUser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RefreshAPIService.Refresh")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/refresh"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.updatedAt == nil {
		return localVarReturnValue, nil, reportError("updatedAt is required and must be specified")
	}
	if r.xAPITOKEN == nil {
		return localVarReturnValue, nil, reportError("xAPITOKEN is required and must be specified")
	}
	if r.xRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xRequestedWith is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "updated_at", r.updatedAt, "form", "")
	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "form", "")
	}
	if r.includeStatic != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_static", r.includeStatic, "form", "")
	}
	if r.clearCache != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clear_cache", r.clearCache, "form", "")
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
