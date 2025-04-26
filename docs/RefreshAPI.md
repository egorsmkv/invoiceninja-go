# \RefreshAPI

All URIs are relative to *https://demo.invoiceninja.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Refresh**](RefreshAPI.md#Refresh) | **Post** /api/v1/refresh | Refresh data by timestamp



## Refresh

> CompanyUser Refresh(ctx).UpdatedAt(updatedAt).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).IncludeStatic(includeStatic).ClearCache(clearCache).Execute()

Refresh data by timestamp



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
	updatedAt := float32(1676173763) // float32 | The unix timestamp from which the refreshed data should be sent from, if no value is passed the system will assume you require all data.
	xAPITOKEN := "TOKEN" // string | The API token to be used for authentication
	xRequestedWith := "XMLHttpRequest" // string | Used to send the XMLHttpRequest header
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)
	includeStatic := "include_static=true" // string | This include will return the full set of static variables used in the application including:   - Currencies   - Countries    - Languages   - Payment Types   - Email Templatees   - Industries  (optional)
	clearCache := "clear_cache=true" // string | Clears cache  Clears (and rebuilds) the static variable cache.    (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RefreshAPI.Refresh(context.Background()).UpdatedAt(updatedAt).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).IncludeStatic(includeStatic).ClearCache(clearCache).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RefreshAPI.Refresh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Refresh`: CompanyUser
	fmt.Fprintf(os.Stdout, "Response from `RefreshAPI.Refresh`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updatedAt** | **float32** | The unix timestamp from which the refreshed data should be sent from, if no value is passed the system will assume you require all data. | 
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 
 **includeStatic** | **string** | This include will return the full set of static variables used in the application including:   - Currencies   - Countries    - Languages   - Payment Types   - Email Templatees   - Industries  | 
 **clearCache** | **string** | Clears cache  Clears (and rebuilds) the static variable cache.    | 

### Return type

[**CompanyUser**](CompanyUser.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

