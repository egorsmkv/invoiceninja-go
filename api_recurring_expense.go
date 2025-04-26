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

// RecurringExpenseAPIService RecurringExpenseAPI service
type RecurringExpenseAPIService service

type ApiUploadRecurringExpenseRequest struct {
	ctx            context.Context
	ApiService     *RecurringExpenseAPIService
	xAPITOKEN      *string
	xRequestedWith *string
	id             string
	include        *string
	method         *string
	documents      []*os.File
}

// The API token to be used for authentication
func (r ApiUploadRecurringExpenseRequest) XAPITOKEN(xAPITOKEN string) ApiUploadRecurringExpenseRequest {
	r.xAPITOKEN = &xAPITOKEN
	return r
}

// Used to send the XMLHttpRequest header
func (r ApiUploadRecurringExpenseRequest) XRequestedWith(xRequestedWith string) ApiUploadRecurringExpenseRequest {
	r.xRequestedWith = &xRequestedWith
	return r
}

// Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes
func (r ApiUploadRecurringExpenseRequest) Include(include string) ApiUploadRecurringExpenseRequest {
	r.include = &include
	return r
}

func (r ApiUploadRecurringExpenseRequest) Method(method string) ApiUploadRecurringExpenseRequest {
	r.method = &method
	return r
}

func (r ApiUploadRecurringExpenseRequest) Documents(documents []*os.File) ApiUploadRecurringExpenseRequest {
	r.documents = documents
	return r
}

func (r ApiUploadRecurringExpenseRequest) Execute() (*RecurringExpense, *http.Response, error) {
	return r.ApiService.UploadRecurringExpenseExecute(r)
}

/*
UploadRecurringExpense Uploads a document to a recurring_expense

Handles the uploading of a document to a recurring_expense

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The RecurringExpense Hashed ID
	@return ApiUploadRecurringExpenseRequest
*/
func (a *RecurringExpenseAPIService) UploadRecurringExpense(ctx context.Context, id string) ApiUploadRecurringExpenseRequest {
	return ApiUploadRecurringExpenseRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return RecurringExpense
func (a *RecurringExpenseAPIService) UploadRecurringExpenseExecute(r ApiUploadRecurringExpenseRequest) (*RecurringExpense, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RecurringExpense
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecurringExpenseAPIService.UploadRecurringExpense")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/recurring_expenses/{id}/upload"
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

	documentsLocalVarFormFileName = "documents"
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
