# \LocationsAPI

All URIs are relative to *https://demo.invoiceninja.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkLocations**](LocationsAPI.md#BulkLocations) | **Post** /api/v1/locations/bulk | Bulk location actions
[**DeleteLocation**](LocationsAPI.md#DeleteLocation) | **Delete** /api/v1/locations/{id} | Delete location
[**GetLocations**](LocationsAPI.md#GetLocations) | **Get** /api/v1/locations | List locations
[**GetLocationsCreate**](LocationsAPI.md#GetLocationsCreate) | **Get** /api/v1/locations/create | Blank Location
[**ShowLocation**](LocationsAPI.md#ShowLocation) | **Get** /api/v1/locations/{id} | Show location
[**StoreLocation**](LocationsAPI.md#StoreLocation) | **Post** /api/v1/locations | Create location
[**UpdateLocation**](LocationsAPI.md#UpdateLocation) | **Put** /api/v1/locations/{id} | Update location



## BulkLocations

> Location BulkLocations(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).GenericBulkAction(genericBulkAction).Index(index).Execute()

Bulk location actions



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
	genericBulkAction := *openapiclient.NewGenericBulkAction() // GenericBulkAction | Bulk action array
	index := "user" // string | Replaces the default response index from data to a user specific string  ie.  ```html   ?index=new_index ```  response is wrapped  ```json   {     'new_index' : [       .....       ]   } ```  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocationsAPI.BulkLocations(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).GenericBulkAction(genericBulkAction).Index(index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationsAPI.BulkLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkLocations`: Location
	fmt.Fprintf(os.Stdout, "Response from `LocationsAPI.BulkLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **genericBulkAction** | [**GenericBulkAction**](GenericBulkAction.md) | Bulk action array | 
 **index** | **string** | Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;  | 

### Return type

[**Location**](Location.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLocation

> DeleteLocation(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Execute()

Delete location



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
	id := "D2J234DFA" // string | The Location Hashed ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LocationsAPI.DeleteLocation(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationsAPI.DeleteLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Location Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLocationRequest struct via the builder pattern


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


## GetLocations

> GetLocations200Response GetLocations(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Index(index).Name(name).Sort(sort).Execute()

List locations



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
	index := "user" // string | Replaces the default response index from data to a user specific string  ie.  ```html   ?index=new_index ```  response is wrapped  ```json   {     'new_index' : [       .....       ]   } ```  (optional)
	name := "warehouse" // string | Filter by location name  ```html ?name=warehouse ```  (optional)
	sort := "name|desc" // string | Returns the list sorted by column in ascending or descending order.  ```html   ?sort=name|desc ```  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocationsAPI.GetLocations(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Index(index).Name(name).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationsAPI.GetLocations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocations`: GetLocations200Response
	fmt.Fprintf(os.Stdout, "Response from `LocationsAPI.GetLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **index** | **string** | Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;  | 
 **name** | **string** | Filter by location name  &#x60;&#x60;&#x60;html ?name&#x3D;warehouse &#x60;&#x60;&#x60;  | 
 **sort** | **string** | Returns the list sorted by column in ascending or descending order.  &#x60;&#x60;&#x60;html   ?sort&#x3D;name|desc &#x60;&#x60;&#x60;  | 

### Return type

[**GetLocations200Response**](GetLocations200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocationsCreate

> Location GetLocationsCreate(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Index(index).Execute()

Blank Location



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
	index := "user" // string | Replaces the default response index from data to a user specific string  ie.  ```html   ?index=new_index ```  response is wrapped  ```json   {     'new_index' : [       .....       ]   } ```  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocationsAPI.GetLocationsCreate(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Index(index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationsAPI.GetLocationsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLocationsCreate`: Location
	fmt.Fprintf(os.Stdout, "Response from `LocationsAPI.GetLocationsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **index** | **string** | Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;  | 

### Return type

[**Location**](Location.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowLocation

> Location ShowLocation(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Index(index).Execute()

Show location



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
	id := "D2J234DFA" // string | The Location Hashed ID
	index := "user" // string | Replaces the default response index from data to a user specific string  ie.  ```html   ?index=new_index ```  response is wrapped  ```json   {     'new_index' : [       .....       ]   } ```  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocationsAPI.ShowLocation(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Index(index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationsAPI.ShowLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowLocation`: Location
	fmt.Fprintf(os.Stdout, "Response from `LocationsAPI.ShowLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Location Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **index** | **string** | Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;  | 

### Return type

[**Location**](Location.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreLocation

> Location StoreLocation(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).LocationRequest(locationRequest).Index(index).Execute()

Create location



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
	locationRequest := *openapiclient.NewLocationRequest() // LocationRequest | Location object that needs to be added to the company
	index := "user" // string | Replaces the default response index from data to a user specific string  ie.  ```html   ?index=new_index ```  response is wrapped  ```json   {     'new_index' : [       .....       ]   } ```  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocationsAPI.StoreLocation(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).LocationRequest(locationRequest).Index(index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationsAPI.StoreLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoreLocation`: Location
	fmt.Fprintf(os.Stdout, "Response from `LocationsAPI.StoreLocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **locationRequest** | [**LocationRequest**](LocationRequest.md) | Location object that needs to be added to the company | 
 **index** | **string** | Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;  | 

### Return type

[**Location**](Location.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLocation

> Location UpdateLocation(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).LocationRequest(locationRequest).Index(index).Execute()

Update location



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
	id := "D2J234DFA" // string | The Location Hashed ID
	locationRequest := *openapiclient.NewLocationRequest() // LocationRequest | Location object that needs to be updated
	index := "user" // string | Replaces the default response index from data to a user specific string  ie.  ```html   ?index=new_index ```  response is wrapped  ```json   {     'new_index' : [       .....       ]   } ```  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LocationsAPI.UpdateLocation(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).LocationRequest(locationRequest).Index(index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocationsAPI.UpdateLocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLocation`: Location
	fmt.Fprintf(os.Stdout, "Response from `LocationsAPI.UpdateLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Location Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **locationRequest** | [**LocationRequest**](LocationRequest.md) | Location object that needs to be updated | 
 **index** | **string** | Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;  | 

### Return type

[**Location**](Location.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

