package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// ClaimLicenseAPIService ClaimLicenseAPI service
type ClaimLicenseAPIService service

type ApiGetClaimLicenseRequest struct {
	ctx            context.Context
	ApiService     *ClaimLicenseAPIService
	xRequestedWith *string
	licenseKey     *string
	productId      *string
}

// Used to send the XMLHttpRequest header
func (r ApiGetClaimLicenseRequest) XRequestedWith(xRequestedWith string) ApiGetClaimLicenseRequest {
	r.xRequestedWith = &xRequestedWith
	return r
}

// The license hash
func (r ApiGetClaimLicenseRequest) LicenseKey(licenseKey string) ApiGetClaimLicenseRequest {
	r.licenseKey = &licenseKey
	return r
}

// The ID of the product purchased.
func (r ApiGetClaimLicenseRequest) ProductId(productId string) ApiGetClaimLicenseRequest {
	r.productId = &productId
	return r
}

func (r ApiGetClaimLicenseRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetClaimLicenseExecute(r)
}

/*
GetClaimLicense Attempts to claim a white label license

Attempts to claim a white label license

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetClaimLicenseRequest
*/
func (a *ClaimLicenseAPIService) GetClaimLicense(ctx context.Context) ApiGetClaimLicenseRequest {
	return ApiGetClaimLicenseRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *ClaimLicenseAPIService) GetClaimLicenseExecute(r ApiGetClaimLicenseRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClaimLicenseAPIService.GetClaimLicense")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/claim_license"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xRequestedWith == nil {
		return nil, reportError("xRequestedWith is required and must be specified")
	}
	if r.licenseKey == nil {
		return nil, reportError("licenseKey is required and must be specified")
	}
	if r.productId == nil {
		return nil, reportError("productId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "license_key", r.licenseKey, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "product_id", r.productId, "form", "")
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
