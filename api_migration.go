package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// MigrationAPIService MigrationAPI service
type MigrationAPIService service

type ApiPostPurgeCompanyRequest struct {
	ctx            context.Context
	ApiService     *MigrationAPIService
	xAPITOKEN      *string
	xRequestedWith *string
	company        string
}

// The API token to be used for authentication
func (r ApiPostPurgeCompanyRequest) XAPITOKEN(xAPITOKEN string) ApiPostPurgeCompanyRequest {
	r.xAPITOKEN = &xAPITOKEN
	return r
}

// Used to send the XMLHttpRequest header
func (r ApiPostPurgeCompanyRequest) XRequestedWith(xRequestedWith string) ApiPostPurgeCompanyRequest {
	r.xRequestedWith = &xRequestedWith
	return r
}

func (r ApiPostPurgeCompanyRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostPurgeCompanyExecute(r)
}

/*
PostPurgeCompany Attempts to purge a company record and all its child records

Attempts to purge a company record and all its child records

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param company The Company Hashed ID
	@return ApiPostPurgeCompanyRequest
*/
func (a *MigrationAPIService) PostPurgeCompany(ctx context.Context, company string) ApiPostPurgeCompanyRequest {
	return ApiPostPurgeCompanyRequest{
		ApiService: a,
		ctx:        ctx,
		company:    company,
	}
}

// Execute executes the request
func (a *MigrationAPIService) PostPurgeCompanyExecute(r ApiPostPurgeCompanyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MigrationAPIService.PostPurgeCompany")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/migration/purge/{company}"
	localVarPath = strings.Replace(localVarPath, "{"+"company"+"}", url.PathEscape(parameterValueToString(r.company, "company")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAPITOKEN == nil {
		return nil, reportError("xAPITOKEN is required and must be specified")
	}
	if r.xRequestedWith == nil {
		return nil, reportError("xRequestedWith is required and must be specified")
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

type ApiPostPurgeCompanySaveSettingsRequest struct {
	ctx            context.Context
	ApiService     *MigrationAPIService
	xAPITOKEN      *string
	xRequestedWith *string
	company        string
}

// The API token to be used for authentication
func (r ApiPostPurgeCompanySaveSettingsRequest) XAPITOKEN(xAPITOKEN string) ApiPostPurgeCompanySaveSettingsRequest {
	r.xAPITOKEN = &xAPITOKEN
	return r
}

// Used to send the XMLHttpRequest header
func (r ApiPostPurgeCompanySaveSettingsRequest) XRequestedWith(xRequestedWith string) ApiPostPurgeCompanySaveSettingsRequest {
	r.xRequestedWith = &xRequestedWith
	return r
}

func (r ApiPostPurgeCompanySaveSettingsRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostPurgeCompanySaveSettingsExecute(r)
}

/*
PostPurgeCompanySaveSettings Attempts to purge a companies child records but save the company record and its settings

Attempts to purge a companies child records but save the company record and its settings

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param company The Company Hashed ID
	@return ApiPostPurgeCompanySaveSettingsRequest
*/
func (a *MigrationAPIService) PostPurgeCompanySaveSettings(ctx context.Context, company string) ApiPostPurgeCompanySaveSettingsRequest {
	return ApiPostPurgeCompanySaveSettingsRequest{
		ApiService: a,
		ctx:        ctx,
		company:    company,
	}
}

// Execute executes the request
func (a *MigrationAPIService) PostPurgeCompanySaveSettingsExecute(r ApiPostPurgeCompanySaveSettingsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MigrationAPIService.PostPurgeCompanySaveSettings")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/migration/purge_save_settings/{company}"
	localVarPath = strings.Replace(localVarPath, "{"+"company"+"}", url.PathEscape(parameterValueToString(r.company, "company")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAPITOKEN == nil {
		return nil, reportError("xAPITOKEN is required and must be specified")
	}
	if r.xRequestedWith == nil {
		return nil, reportError("xRequestedWith is required and must be specified")
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

type ApiPostStartMigrationRequest struct {
	ctx            context.Context
	ApiService     *MigrationAPIService
	xAPITOKEN      *string
	xRequestedWith *string
	xAPIPASSWORD   *string
	migration      *map[string]interface{}
}

// The API token to be used for authentication
func (r ApiPostStartMigrationRequest) XAPITOKEN(xAPITOKEN string) ApiPostStartMigrationRequest {
	r.xAPITOKEN = &xAPITOKEN
	return r
}

// Used to send the XMLHttpRequest header
func (r ApiPostStartMigrationRequest) XRequestedWith(xRequestedWith string) ApiPostStartMigrationRequest {
	r.xRequestedWith = &xRequestedWith
	return r
}

// The login password when challenged on certain protected routes
func (r ApiPostStartMigrationRequest) XAPIPASSWORD(xAPIPASSWORD string) ApiPostStartMigrationRequest {
	r.xAPIPASSWORD = &xAPIPASSWORD
	return r
}

// The migraton file
func (r ApiPostStartMigrationRequest) Migration(migration map[string]interface{}) ApiPostStartMigrationRequest {
	r.migration = &migration
	return r
}

func (r ApiPostStartMigrationRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostStartMigrationExecute(r)
}

/*
PostStartMigration Starts the migration from previous version of Invoice Ninja

Starts the migration from previous version of Invoice Ninja

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPostStartMigrationRequest
*/
func (a *MigrationAPIService) PostStartMigration(ctx context.Context) ApiPostStartMigrationRequest {
	return ApiPostStartMigrationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *MigrationAPIService) PostStartMigrationExecute(r ApiPostStartMigrationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MigrationAPIService.PostStartMigration")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/migration/start"

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
	if r.migration == nil {
		return nil, reportError("migration is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "migration", r.migration, "form", "")
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
