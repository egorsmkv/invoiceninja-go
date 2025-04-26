# \MigrationAPI

All URIs are relative to *https://demo.invoiceninja.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostPurgeCompany**](MigrationAPI.md#PostPurgeCompany) | **Post** /api/v1/migration/purge/{company} | Attempts to purge a company record and all its child records
[**PostPurgeCompanySaveSettings**](MigrationAPI.md#PostPurgeCompanySaveSettings) | **Post** /api/v1/migration/purge_save_settings/{company} | Attempts to purge a companies child records but save the company record and its settings
[**PostStartMigration**](MigrationAPI.md#PostStartMigration) | **Post** /api/v1/migration/start | Starts the migration from previous version of Invoice Ninja



## PostPurgeCompany

> PostPurgeCompany(ctx, company).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Execute()

Attempts to purge a company record and all its child records



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	xAPITOKEN := "TOKEN" // string | The API token to be used for authentication
	xRequestedWith := "XMLHttpRequest" // string | Used to send the XMLHttpRequest header
	company := "D2J234DFA" // string | The Company Hashed ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MigrationAPI.PostPurgeCompany(context.Background(), company).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MigrationAPI.PostPurgeCompany``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**company** | **string** | The Company Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPurgeCompanyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPurgeCompanySaveSettings

> PostPurgeCompanySaveSettings(ctx, company).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Execute()

Attempts to purge a companies child records but save the company record and its settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	xAPITOKEN := "TOKEN" // string | The API token to be used for authentication
	xRequestedWith := "XMLHttpRequest" // string | Used to send the XMLHttpRequest header
	company := "D2J234DFA" // string | The Company Hashed ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MigrationAPI.PostPurgeCompanySaveSettings(context.Background(), company).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MigrationAPI.PostPurgeCompanySaveSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**company** | **string** | The Company Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostPurgeCompanySaveSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 


### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostStartMigration

> PostStartMigration(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).XAPIPASSWORD(xAPIPASSWORD).Migration(migration).Execute()

Starts the migration from previous version of Invoice Ninja



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	xAPITOKEN := "TOKEN" // string | The API token to be used for authentication
	xRequestedWith := "XMLHttpRequest" // string | Used to send the XMLHttpRequest header
	xAPIPASSWORD := "supersecretpassword" // string | The login password when challenged on certain protected routes
	migration := map[string]interface{}{ ... } // map[string]interface{} | The migraton file

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MigrationAPI.PostStartMigration(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).XAPIPASSWORD(xAPIPASSWORD).Migration(migration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MigrationAPI.PostStartMigration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostStartMigrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **xAPIPASSWORD** | **string** | The login password when challenged on certain protected routes | 
 **migration** | [**map[string]interface{}**](map[string]interface{}.md) | The migraton file | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

