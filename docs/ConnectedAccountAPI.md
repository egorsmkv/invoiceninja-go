# \ConnectedAccountAPI

All URIs are relative to *https://demo.invoiceninja.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectedAccount**](ConnectedAccountAPI.md#ConnectedAccount) | **Post** /api/v1/connected_account | Connect an oauth user to an existing user



## ConnectedAccount

> User ConnectedAccount(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).IncludeStatic(includeStatic).ClearCache(clearCache).Execute()

Connect an oauth user to an existing user



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
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)
	includeStatic := "include_static=true" // string | This include will return the full set of static variables used in the application including:   - Currencies   - Countries    - Languages   - Payment Types   - Email Templatees   - Industries  (optional)
	clearCache := "clear_cache=true" // string | Clears cache  Clears (and rebuilds) the static variable cache.    (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectedAccountAPI.ConnectedAccount(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).IncludeStatic(includeStatic).ClearCache(clearCache).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectedAccountAPI.ConnectedAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectedAccount`: User
	fmt.Fprintf(os.Stdout, "Response from `ConnectedAccountAPI.ConnectedAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConnectedAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 
 **includeStatic** | **string** | This include will return the full set of static variables used in the application including:   - Currencies   - Countries    - Languages   - Payment Types   - Email Templatees   - Industries  | 
 **clearCache** | **string** | Clears cache  Clears (and rebuilds) the static variable cache.    | 

### Return type

[**User**](User.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

