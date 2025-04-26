# \TokensAPI

All URIs are relative to *https://demo.invoiceninja.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkTokens**](TokensAPI.md#BulkTokens) | **Post** /api/v1/tokens/bulk | Performs bulk actions on an array of tokens
[**DeleteToken**](TokensAPI.md#DeleteToken) | **Delete** /api/v1/tokens/{id} | Deletes a token
[**EditToken**](TokensAPI.md#EditToken) | **Get** /api/v1/tokens/{id}/edit | Shows a token for editting
[**GetTokens**](TokensAPI.md#GetTokens) | **Get** /api/v1/tokens | Gets a list of company tokens
[**GetTokensCreate**](TokensAPI.md#GetTokensCreate) | **Get** /api/v1/tokens/create | Gets a new blank token object
[**ShowToken**](TokensAPI.md#ShowToken) | **Get** /api/v1/tokens/{id} | Shows a token
[**StoreToken**](TokensAPI.md#StoreToken) | **Post** /api/v1/tokens | Adds a token
[**UpdateToken**](TokensAPI.md#UpdateToken) | **Put** /api/v1/tokens/{id} | Updates a token



## BulkTokens

> CompanyToken BulkTokens(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).RequestBody(requestBody).Index(index).Execute()

Performs bulk actions on an array of tokens



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
	requestBody := []int32{int32(123)} // []int32 | Token ids
	index := "user" // string | Replaces the default response index from data to a user specific string  ie.  ```html   ?index=new_index ```  response is wrapped  ```json   {     'new_index' : [       .....       ]   } ```  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.BulkTokens(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).RequestBody(requestBody).Index(index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.BulkTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkTokens`: CompanyToken
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.BulkTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **requestBody** | **[]int32** | Token ids | 
 **index** | **string** | Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;  | 

### Return type

[**CompanyToken**](CompanyToken.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteToken

> DeleteToken(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Deletes a token



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
	id := "D2J234DFA" // string | The Token Hashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TokensAPI.DeleteToken(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.DeleteToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Token Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

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


## EditToken

> CompanyToken EditToken(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Shows a token for editting



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
	id := "D2J234DFA" // string | The Token Hashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.EditToken(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.EditToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditToken`: CompanyToken
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.EditToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Token Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**CompanyToken**](CompanyToken.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokens

> GetTokens200Response GetTokens(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Index(index).Execute()

Gets a list of company tokens



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
	index := "user" // string | Replaces the default response index from data to a user specific string  ie.  ```html   ?index=new_index ```  response is wrapped  ```json   {     'new_index' : [       .....       ]   } ```  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.GetTokens(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Index(index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.GetTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokens`: GetTokens200Response
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.GetTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 
 **index** | **string** | Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;  | 

### Return type

[**GetTokens200Response**](GetTokens200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokensCreate

> CompanyToken GetTokensCreate(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Gets a new blank token object



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.GetTokensCreate(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.GetTokensCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokensCreate`: CompanyToken
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.GetTokensCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokensCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**CompanyToken**](CompanyToken.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowToken

> CompanyToken ShowToken(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Shows a token



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
	id := "D2J234DFA" // string | The Token Hashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.ShowToken(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.ShowToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowToken`: CompanyToken
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.ShowToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Token Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**CompanyToken**](CompanyToken.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreToken

> CompanyToken StoreToken(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Adds a token



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.StoreToken(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.StoreToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoreToken`: CompanyToken
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.StoreToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**CompanyToken**](CompanyToken.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateToken

> CompanyToken UpdateToken(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Updates a token



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
	id := "D2J234DFA" // string | The Token Hashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.UpdateToken(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.UpdateToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateToken`: CompanyToken
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.UpdateToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Token Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**CompanyToken**](CompanyToken.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

