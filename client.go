package openapi

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	JsonCheck       = regexp.MustCompile(`(?i:(?:application|text)/(?:[^;]+\+)?json)`)
	XmlCheck        = regexp.MustCompile(`(?i:(?:application|text)/(?:[^;]+\+)?xml)`)
	queryParamSplit = regexp.MustCompile(`(^|&)([^&]+)`)
	queryDescape    = strings.NewReplacer("%5B", "[", "%5D", "]")
)

// APIClient manages communication with the Invoice Ninja API Reference. API v5.11.48
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	ActivitiesAPI *ActivitiesAPIService

	AuthAPI *AuthAPIService

	BankIntegrationsAPI *BankIntegrationsAPIService

	BankTransactionRulesAPI *BankTransactionRulesAPIService

	BankTransactionsAPI *BankTransactionsAPIService

	ChartsAPI *ChartsAPIService

	ClaimLicenseAPI *ClaimLicenseAPIService

	ClientGatewayTokensAPI *ClientGatewayTokensAPIService

	ClientsAPI *ClientsAPIService

	CompaniesAPI *CompaniesAPIService

	CompanyGatewaysAPI *CompanyGatewaysAPIService

	CompanyLedgerAPI *CompanyLedgerAPIService

	CompanyUserAPI *CompanyUserAPIService

	ConnectedAccountAPI *ConnectedAccountAPIService

	CreditsAPI *CreditsAPIService

	DesignsAPI *DesignsAPIService

	DocumentsAPI *DocumentsAPIService

	EmailsAPI *EmailsAPIService

	ExpenseAPI *ExpenseAPIService

	ExpenseCategoriesAPI *ExpenseCategoriesAPIService

	ExpensesAPI *ExpensesAPIService

	ExportAPI *ExportAPIService

	GroupSettingsAPI *GroupSettingsAPIService

	HealthCheckAPI *HealthCheckAPIService

	ImportAPI *ImportAPIService

	ImportsAPI *ImportsAPIService

	InvoicesAPI *InvoicesAPIService

	LocationsAPI *LocationsAPIService

	LogoutAPI *LogoutAPIService

	MigrationAPI *MigrationAPIService

	OneTimeTokenAPI *OneTimeTokenAPIService

	PaymentTermsAPI *PaymentTermsAPIService

	PaymentTermssAPI *PaymentTermssAPIService

	PaymentsAPI *PaymentsAPIService

	PingAPI *PingAPIService

	PostmarkAPI *PostmarkAPIService

	PreviewAPI *PreviewAPIService

	ProductsAPI *ProductsAPIService

	ProjectsAPI *ProjectsAPIService

	PurchaseOrdersAPI *PurchaseOrdersAPIService

	QuotesAPI *QuotesAPIService

	RecurringExpenseAPI *RecurringExpenseAPIService

	RecurringExpensesAPI *RecurringExpensesAPIService

	RecurringInvoicesAPI *RecurringInvoicesAPIService

	RecurringQuotesAPI *RecurringQuotesAPIService

	RefreshAPI *RefreshAPIService

	ReportsAPI *ReportsAPIService

	StaticsAPI *StaticsAPIService

	SubscriptionsAPI *SubscriptionsAPIService

	SupportAPI *SupportAPIService

	SystemLogsAPI *SystemLogsAPIService

	TaskSchedulersAPI *TaskSchedulersAPIService

	TaskStatusAPI *TaskStatusAPIService

	TaskStatussAPI *TaskStatussAPIService

	TasksAPI *TasksAPIService

	TaxRatesAPI *TaxRatesAPIService

	TemplatesAPI *TemplatesAPIService

	TokensAPI *TokensAPIService

	UpdateAPI *UpdateAPIService

	UsersAPI *UsersAPIService

	VendorsAPI *VendorsAPIService

	WebcronAPI *WebcronAPIService

	WebhooksAPI *WebhooksAPIService

	YodleeAPI *YodleeAPIService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.ActivitiesAPI = (*ActivitiesAPIService)(&c.common)
	c.AuthAPI = (*AuthAPIService)(&c.common)
	c.BankIntegrationsAPI = (*BankIntegrationsAPIService)(&c.common)
	c.BankTransactionRulesAPI = (*BankTransactionRulesAPIService)(&c.common)
	c.BankTransactionsAPI = (*BankTransactionsAPIService)(&c.common)
	c.ChartsAPI = (*ChartsAPIService)(&c.common)
	c.ClaimLicenseAPI = (*ClaimLicenseAPIService)(&c.common)
	c.ClientGatewayTokensAPI = (*ClientGatewayTokensAPIService)(&c.common)
	c.ClientsAPI = (*ClientsAPIService)(&c.common)
	c.CompaniesAPI = (*CompaniesAPIService)(&c.common)
	c.CompanyGatewaysAPI = (*CompanyGatewaysAPIService)(&c.common)
	c.CompanyLedgerAPI = (*CompanyLedgerAPIService)(&c.common)
	c.CompanyUserAPI = (*CompanyUserAPIService)(&c.common)
	c.ConnectedAccountAPI = (*ConnectedAccountAPIService)(&c.common)
	c.CreditsAPI = (*CreditsAPIService)(&c.common)
	c.DesignsAPI = (*DesignsAPIService)(&c.common)
	c.DocumentsAPI = (*DocumentsAPIService)(&c.common)
	c.EmailsAPI = (*EmailsAPIService)(&c.common)
	c.ExpenseAPI = (*ExpenseAPIService)(&c.common)
	c.ExpenseCategoriesAPI = (*ExpenseCategoriesAPIService)(&c.common)
	c.ExpensesAPI = (*ExpensesAPIService)(&c.common)
	c.ExportAPI = (*ExportAPIService)(&c.common)
	c.GroupSettingsAPI = (*GroupSettingsAPIService)(&c.common)
	c.HealthCheckAPI = (*HealthCheckAPIService)(&c.common)
	c.ImportAPI = (*ImportAPIService)(&c.common)
	c.ImportsAPI = (*ImportsAPIService)(&c.common)
	c.InvoicesAPI = (*InvoicesAPIService)(&c.common)
	c.LocationsAPI = (*LocationsAPIService)(&c.common)
	c.LogoutAPI = (*LogoutAPIService)(&c.common)
	c.MigrationAPI = (*MigrationAPIService)(&c.common)
	c.OneTimeTokenAPI = (*OneTimeTokenAPIService)(&c.common)
	c.PaymentTermsAPI = (*PaymentTermsAPIService)(&c.common)
	c.PaymentTermssAPI = (*PaymentTermssAPIService)(&c.common)
	c.PaymentsAPI = (*PaymentsAPIService)(&c.common)
	c.PingAPI = (*PingAPIService)(&c.common)
	c.PostmarkAPI = (*PostmarkAPIService)(&c.common)
	c.PreviewAPI = (*PreviewAPIService)(&c.common)
	c.ProductsAPI = (*ProductsAPIService)(&c.common)
	c.ProjectsAPI = (*ProjectsAPIService)(&c.common)
	c.PurchaseOrdersAPI = (*PurchaseOrdersAPIService)(&c.common)
	c.QuotesAPI = (*QuotesAPIService)(&c.common)
	c.RecurringExpenseAPI = (*RecurringExpenseAPIService)(&c.common)
	c.RecurringExpensesAPI = (*RecurringExpensesAPIService)(&c.common)
	c.RecurringInvoicesAPI = (*RecurringInvoicesAPIService)(&c.common)
	c.RecurringQuotesAPI = (*RecurringQuotesAPIService)(&c.common)
	c.RefreshAPI = (*RefreshAPIService)(&c.common)
	c.ReportsAPI = (*ReportsAPIService)(&c.common)
	c.StaticsAPI = (*StaticsAPIService)(&c.common)
	c.SubscriptionsAPI = (*SubscriptionsAPIService)(&c.common)
	c.SupportAPI = (*SupportAPIService)(&c.common)
	c.SystemLogsAPI = (*SystemLogsAPIService)(&c.common)
	c.TaskSchedulersAPI = (*TaskSchedulersAPIService)(&c.common)
	c.TaskStatusAPI = (*TaskStatusAPIService)(&c.common)
	c.TaskStatussAPI = (*TaskStatussAPIService)(&c.common)
	c.TasksAPI = (*TasksAPIService)(&c.common)
	c.TaxRatesAPI = (*TaxRatesAPIService)(&c.common)
	c.TemplatesAPI = (*TemplatesAPIService)(&c.common)
	c.TokensAPI = (*TokensAPIService)(&c.common)
	c.UpdateAPI = (*UpdateAPIService)(&c.common)
	c.UsersAPI = (*UsersAPIService)(&c.common)
	c.VendorsAPI = (*VendorsAPIService)(&c.common)
	c.WebcronAPI = (*WebcronAPIService)(&c.common)
	c.WebhooksAPI = (*WebhooksAPIService)(&c.common)
	c.YodleeAPI = (*YodleeAPIService)(&c.common)

	return c
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insensitive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.EqualFold(a, needle) {
			return true
		}
	}
	return false
}

func parameterValueToString(obj any, key string) string {
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		if actualObj, ok := obj.(interface{ GetActualInstanceValue() any }); ok {
			return fmt.Sprintf("%v", actualObj.GetActualInstanceValue())
		}

		return fmt.Sprintf("%v", obj)
	}
	param, ok := obj.(MappedNullable)
	if !ok {
		return ""
	}
	dataMap, err := param.ToMap()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%v", dataMap[key])
}

// parameterAddToHeaderOrQuery adds the provided object to the request header or url query
// supporting deep object syntax
func parameterAddToHeaderOrQuery(headerOrQueryParams any, keyPrefix string, obj any, style string, collectionType string) {
	v := reflect.ValueOf(obj)
	value := ""
	if v == reflect.ValueOf(nil) {
		value = "null"
	} else {
		switch v.Kind() {
		case reflect.Invalid:
			value = "invalid"

		case reflect.Struct:
			if t, ok := obj.(MappedNullable); ok {
				dataMap, err := t.ToMap()
				if err != nil {
					return
				}
				parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, dataMap, style, collectionType)
				return
			}
			if t, ok := obj.(time.Time); ok {
				parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, t.Format(time.RFC3339Nano), style, collectionType)
				return
			}
			value = v.Type().String() + " value"
		case reflect.Slice:
			indValue := reflect.ValueOf(obj)
			if indValue == reflect.ValueOf(nil) {
				return
			}
			lenIndValue := indValue.Len()
			for i := 0; i < lenIndValue; i++ {
				arrayValue := indValue.Index(i)
				keyPrefixForCollectionType := keyPrefix
				if style == "deepObject" {
					keyPrefixForCollectionType = keyPrefix + "[" + strconv.Itoa(i) + "]"
				}
				parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefixForCollectionType, arrayValue.Interface(), style, collectionType)
			}
			return

		case reflect.Map:
			indValue := reflect.ValueOf(obj)
			if indValue == reflect.ValueOf(nil) {
				return
			}
			iter := indValue.MapRange()
			for iter.Next() {
				k, v := iter.Key(), iter.Value()
				parameterAddToHeaderOrQuery(headerOrQueryParams, fmt.Sprintf("%s[%s]", keyPrefix, k.String()), v.Interface(), style, collectionType)
			}
			return

		case reflect.Interface:
			fallthrough
		case reflect.Ptr:
			parameterAddToHeaderOrQuery(headerOrQueryParams, keyPrefix, v.Elem().Interface(), style, collectionType)
			return

		case reflect.Int, reflect.Int8, reflect.Int16,
			reflect.Int32, reflect.Int64:
			value = strconv.FormatInt(v.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16,
			reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			value = strconv.FormatUint(v.Uint(), 10)
		case reflect.Float32, reflect.Float64:
			value = strconv.FormatFloat(v.Float(), 'g', -1, 32)
		case reflect.Bool:
			value = strconv.FormatBool(v.Bool())
		case reflect.String:
			value = v.String()
		default:
			value = v.Type().String() + " value"
		}
	}

	switch valuesMap := headerOrQueryParams.(type) {
	case url.Values:
		if collectionType == "csv" && valuesMap.Get(keyPrefix) != "" {
			valuesMap.Set(keyPrefix, valuesMap.Get(keyPrefix)+","+value)
		} else {
			valuesMap.Add(keyPrefix, value)
		}
	case map[string]string:
		valuesMap[keyPrefix] = value
	}
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	if c.cfg.Debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	resp, err := c.cfg.HTTPClient.Do(request)
	if err != nil {
		return resp, err
	}

	if c.cfg.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return resp, err
		}
		log.Printf("\n%s\n", string(dump))
	}
	return resp, err
}

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *APIClient) GetConfig() *Configuration {
	return c.cfg
}

type formFile struct {
	fileName     string
	formFileName string
	fileBytes    []byte
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody any,
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	formFiles []formFile,
) (localVarRequest *http.Request, err error) {
	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(formFiles) > 0) {
		if body != nil {
			return nil, errors.New("cannot specify postBody and multipart form at the same time")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		for _, formFile := range formFiles {
			if len(formFile.fileBytes) > 0 && formFile.fileName != "" {
				w.Boundary()
				part, err := w.CreateFormFile(formFile.formFileName, filepath.Base(formFile.fileName))
				if err != nil {
					return nil, err
				}
				_, err = part.Write(formFile.fileBytes)
				if err != nil {
					return nil, err
				}
			}
		}

		// Set the Boundary in the Content-Type
		headerParams["Content-Type"] = w.FormDataContentType()

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("cannot specify postBody and x-www-form-urlencoded form at the same time")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		url.Host = c.cfg.Host
	}

	// Override request scheme, if applicable
	if c.cfg.Scheme != "" {
		url.Scheme = c.cfg.Scheme
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = queryParamSplit.ReplaceAllStringFunc(query.Encode(), func(s string) string {
		pieces := strings.Split(s, "=")
		pieces[0] = queryDescape.Replace(pieces[0])
		return strings.Join(pieces, "=")
	})

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers[h] = []string{v}
		}
		localVarRequest.Header = headers
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.
	}

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}
	return localVarRequest, nil
}

func (c *APIClient) decode(v any, b []byte, contentType string) (err error) {
	if len(b) == 0 {
		return nil
	}
	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}
	if _, ok := v.(*os.File); ok {
		tempFile, err := os.CreateTemp("", "HttpClientFile")
		if err != nil {
			return err
		}
		defer tempFile.Close()
		_, err = tempFile.Write(b)
		if err != nil {
			return err
		}
		_, err = tempFile.Seek(0, io.SeekStart)
		return err
	}
	if XmlCheck.MatchString(contentType) {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	if JsonCheck.MatchString(contentType) {
		if actualObj, ok := v.(interface{ GetActualInstance() any }); ok { // oneOf, anyOf schemas
			if unmarshalObj, ok := actualObj.(interface{ UnmarshalJSON([]byte) error }); ok { // make sure it has UnmarshalJSON defined
				if err = unmarshalObj.UnmarshalJSON(b); err != nil {
					return err
				}
			} else {
				return errors.New("unknown type with GetActualInstance but no unmarshalObj.UnmarshalJSON defined")
			}
		} else if err = json.Unmarshal(b, v); err != nil { // simple model
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Set request body from an any
func setBody(body any, contentType string) (bodyBuf *bytes.Buffer, err error) {
	bodyBuf = &bytes.Buffer{}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if fp, ok := body.(*os.File); ok {
		_, err = bodyBuf.ReadFrom(fp)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if JsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if XmlCheck.MatchString(contentType) {
		var bs []byte
		bs, err = xml.Marshal(body)
		if err == nil {
			bodyBuf.Write(bs)
		}
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("invalid body type %s", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body any) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		} else {
			expires = now.Add(lifetime)
		}
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

// GenericOpenAPIError Provides access to the body, error and model on returned errors.
type GenericOpenAPIError struct {
	model any
	error string
	body  []byte
}

// Error returns non-empty string if there was an error.
func (e GenericOpenAPIError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericOpenAPIError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericOpenAPIError) Model() any {
	return e.model
}

// format error message using title and detail when model implements rfc7807
func formatErrorMessage(status string, v any) string {
	str := ""
	metaValue := reflect.ValueOf(v).Elem()

	if metaValue.Kind() == reflect.Struct {
		field := metaValue.FieldByName("Title")
		if field != (reflect.Value{}) {
			str = fmt.Sprintf("%s", field.Interface())
		}

		field = metaValue.FieldByName("Detail")
		if field != (reflect.Value{}) {
			str = fmt.Sprintf("%s (%s)", str, field.Interface())
		}
	}

	return strings.TrimSpace(fmt.Sprintf("%s %s", status, str))
}
