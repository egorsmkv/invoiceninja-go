package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// AuthAPIService AuthAPI service
type AuthAPIService service

type ApiPostLoginRequest struct {
	ctx              context.Context
	ApiService       *AuthAPIService
	xAPITOKEN        *string
	xRequestedWith   *string
	postLoginRequest *PostLoginRequest
	xAPISECRET       *string
	include          *string
	includeStatic    *string
	clearCache       *string
}

// The API token to be used for authentication
func (r ApiPostLoginRequest) XAPITOKEN(xAPITOKEN string) ApiPostLoginRequest {
	r.xAPITOKEN = &xAPITOKEN
	return r
}

// Used to send the XMLHttpRequest header
func (r ApiPostLoginRequest) XRequestedWith(xRequestedWith string) ApiPostLoginRequest {
	r.xRequestedWith = &xRequestedWith
	return r
}

// User credentials  &#x60;&#x60;&#x60;json {   \&quot;email\&quot; : \&quot;fred@flintstonze.com\&quot;,   \&quot;password\&quot; : \&quot;magicpassword123\&quot;,   \&quot;one_time_password\&quot; : \&quot;12345\&quot;, } &#x60;&#x60;&#x60;
func (r ApiPostLoginRequest) PostLoginRequest(postLoginRequest PostLoginRequest) ApiPostLoginRequest {
	r.postLoginRequest = &postLoginRequest
	return r
}

// The API secret as defined by the .env variable API_SECRET. Only needed for self hosted users, and only applicable on the login route.
func (r ApiPostLoginRequest) XAPISECRET(xAPISECRET string) ApiPostLoginRequest {
	r.xAPISECRET = &xAPISECRET
	return r
}

// Include child relations of the CompanyUser object, format is comma separated.      &lt;br /&gt;  &gt; ### **Note**: it is possible to chain multiple includes together, ie. include&#x3D;account,token  &lt;br /&gt;  &#x60;&#x60;&#x60;html  Available includes:    user   company   token   account &#x60;&#x60;&#x60;
func (r ApiPostLoginRequest) Include(include string) ApiPostLoginRequest {
	r.include = &include
	return r
}

// This include will return the full set of static variables used in the application including:   - Currencies   - Countries    - Languages   - Payment Types   - Email Templatees   - Industries
func (r ApiPostLoginRequest) IncludeStatic(includeStatic string) ApiPostLoginRequest {
	r.includeStatic = &includeStatic
	return r
}

// Clears cache  Clears (and rebuilds) the static variable cache.
func (r ApiPostLoginRequest) ClearCache(clearCache string) ApiPostLoginRequest {
	r.clearCache = &clearCache
	return r
}

func (r ApiPostLoginRequest) Execute() (*CompanyUser, *http.Response, error) {
	return r.ApiService.PostLoginExecute(r)
}

/*
PostLogin Login

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPostLoginRequest
*/
func (a *AuthAPIService) PostLogin(ctx context.Context) ApiPostLoginRequest {
	return ApiPostLoginRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CompanyUser
func (a *AuthAPIService) PostLoginExecute(r ApiPostLoginRequest) (*CompanyUser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CompanyUser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthAPIService.PostLogin")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/login"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAPITOKEN == nil {
		return localVarReturnValue, nil, reportError("xAPITOKEN is required and must be specified")
	}
	if r.xRequestedWith == nil {
		return localVarReturnValue, nil, reportError("xRequestedWith is required and must be specified")
	}
	if r.postLoginRequest == nil {
		return localVarReturnValue, nil, reportError("postLoginRequest is required and must be specified")
	}

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
	if r.xAPISECRET != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-API-SECRET", r.xAPISECRET, "simple", "")
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-API-TOKEN", r.xAPITOKEN, "simple", "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Requested-With", r.xRequestedWith, "simple", "")
	// body params
	localVarPostBody = r.postLoginRequest
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
