# \RecurringInvoicesAPI

All URIs are relative to *https://demo.invoiceninja.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionRecurringInvoice**](RecurringInvoicesAPI.md#ActionRecurringInvoice) | **Get** /api/v1/recurring_invoices/{id}/{action} | Custom recurring invoice action
[**BulkRecurringInvoices**](RecurringInvoicesAPI.md#BulkRecurringInvoices) | **Post** /api/v1/recurring_invoices/bulk | Bulk recurring invoice actions
[**DeleteRecurringInvoice**](RecurringInvoicesAPI.md#DeleteRecurringInvoice) | **Delete** /api/v1/recurring_invoices/{id} | Delete recurring invoice
[**DownloadRecurringInvoice**](RecurringInvoicesAPI.md#DownloadRecurringInvoice) | **Get** /api/v1/recurring_invoice/{invitation_key}/download | Download recurring invoice PDF
[**EditRecurringInvoice**](RecurringInvoicesAPI.md#EditRecurringInvoice) | **Get** /api/v1/recurring_invoices/{id}/edit | Edit recurring invoice
[**GetRecurringInvoices**](RecurringInvoicesAPI.md#GetRecurringInvoices) | **Get** /api/v1/recurring_invoices | List recurring invoices
[**GetRecurringInvoicesCreate**](RecurringInvoicesAPI.md#GetRecurringInvoicesCreate) | **Get** /api/v1/recurring_invoices/create | Blank recurring invoice
[**ShowRecurringInvoice**](RecurringInvoicesAPI.md#ShowRecurringInvoice) | **Get** /api/v1/recurring_invoices/{id} | Show recurring invoice
[**StoreRecurringInvoice**](RecurringInvoicesAPI.md#StoreRecurringInvoice) | **Post** /api/v1/recurring_invoices | Create recurring invoice
[**UpdateRecurringInvoice**](RecurringInvoicesAPI.md#UpdateRecurringInvoice) | **Put** /api/v1/recurring_invoices/{id} | Update recurring invoice
[**UploadRecurringInvoice**](RecurringInvoicesAPI.md#UploadRecurringInvoice) | **Post** /api/v1/recurring_invoices/{id}/upload | Add recurring invoice document



## ActionRecurringInvoice

> RecurringInvoice ActionRecurringInvoice(ctx, id, action).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Custom recurring invoice action



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
	id := "D2J234DFA" // string | The RecurringInvoice Hashed ID
	action := "clone_to_quote" // string | The action string to be performed
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecurringInvoicesAPI.ActionRecurringInvoice(context.Background(), id, action).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecurringInvoicesAPI.ActionRecurringInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionRecurringInvoice`: RecurringInvoice
	fmt.Fprintf(os.Stdout, "Response from `RecurringInvoicesAPI.ActionRecurringInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The RecurringInvoice Hashed ID | 
**action** | **string** | The action string to be performed | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionRecurringInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 


 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**RecurringInvoice**](RecurringInvoice.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkRecurringInvoices

> RecurringInvoice BulkRecurringInvoices(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).BulkRecurringInvoicesRequest(bulkRecurringInvoicesRequest).Index(index).Execute()

Bulk recurring invoice actions



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
	bulkRecurringInvoicesRequest := *openapiclient.NewBulkRecurringInvoicesRequest() // BulkRecurringInvoicesRequest | Bulk action details
	index := "user" // string | Replaces the default response index from data to a user specific string  ie.  ```html   ?index=new_index ```  response is wrapped  ```json   {     'new_index' : [       .....       ]   } ```  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecurringInvoicesAPI.BulkRecurringInvoices(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).BulkRecurringInvoicesRequest(bulkRecurringInvoicesRequest).Index(index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecurringInvoicesAPI.BulkRecurringInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkRecurringInvoices`: RecurringInvoice
	fmt.Fprintf(os.Stdout, "Response from `RecurringInvoicesAPI.BulkRecurringInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkRecurringInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **bulkRecurringInvoicesRequest** | [**BulkRecurringInvoicesRequest**](BulkRecurringInvoicesRequest.md) | Bulk action details | 
 **index** | **string** | Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;  | 

### Return type

[**RecurringInvoice**](RecurringInvoice.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRecurringInvoice

> DeleteRecurringInvoice(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Delete recurring invoice



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
	id := "D2J234DFA" // string | The RecurringInvoice Hashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RecurringInvoicesAPI.DeleteRecurringInvoice(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecurringInvoicesAPI.DeleteRecurringInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The RecurringInvoice Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecurringInvoiceRequest struct via the builder pattern


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


## DownloadRecurringInvoice

> DownloadRecurringInvoice(ctx, invitationKey).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Download recurring invoice PDF



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
	invitationKey := "D2J234DFA" // string | The Recurring Invoice Invitation Key
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RecurringInvoicesAPI.DownloadRecurringInvoice(context.Background(), invitationKey).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecurringInvoicesAPI.DownloadRecurringInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationKey** | **string** | The Recurring Invoice Invitation Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadRecurringInvoiceRequest struct via the builder pattern


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


## EditRecurringInvoice

> RecurringInvoice EditRecurringInvoice(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Edit recurring invoice



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
	id := "D2J234DFA" // string | The RecurringInvoice Hashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecurringInvoicesAPI.EditRecurringInvoice(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecurringInvoicesAPI.EditRecurringInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditRecurringInvoice`: RecurringInvoice
	fmt.Fprintf(os.Stdout, "Response from `RecurringInvoicesAPI.EditRecurringInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The RecurringInvoice Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditRecurringInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**RecurringInvoice**](RecurringInvoice.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecurringInvoices

> GetRecurringInvoices200Response GetRecurringInvoices(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).ClientId(clientId).CreatedAt(createdAt).UpdatedAt(updatedAt).IsDeleted(isDeleted).FilterDeletedClients(filterDeletedClients).VendorId(vendorId).Filter(filter).ClientStatus(clientStatus).Sort(sort).Number(number).ProductKey(productKey).NextSendBetween(nextSendBetween).FrequencyId(frequencyId).Execute()

List recurring invoices



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
	clientId := "?client_id={client_id}" // string | Filters the entity list by client_id. Suitable when you only want the entities of a specific client.  ```html GET /api/v1/invoices?client_id=AxB7Hjk9 Returns only invoices for the specified client ```  (optional)
	createdAt := int32(?created_at=2022-01-10) // int32 | Filters the entity list by the created at timestamp. Parameter value can be a datetime string or unix timestamp  ```html GET /api/v1/invoices?created_at=2022-01-10 Returns entities created on January 10th, 2022 ```  (optional)
	updatedAt := int32(?updated_at=2022-01-10) // int32 | Filters the entity list by the updated at timestamp. Parameter value can be a datetime string or unix timestamp  ```html GET /api/v1/invoices?updated_at=2022-01-10 Returns entities last updated on January 10th, 2022 ```  (optional)
	isDeleted := false // bool | Filters the entity list by entities that have been deleted.  ```html GET /api/v1/invoices?is_deleted=true Returns only soft-deleted entities ```  (optional)
	filterDeletedClients := "?filter_deleted_clients=true" // string | Filters the entity list and only returns entities for clients that have not been deleted  ```html GET /api/v1/invoices?filter_deleted_clients=true Returns only invoices for active (non-deleted) clients ```  (optional)
	vendorId := "?vendor_id={vendor_id}" // string | Filters the entity list by an associated vendor  ```html GET /api/v1/purchases?vendor_id=AxB7Hjk9 Returns only purchases for the specified vendor ```  (optional)
	filter := "?filter=bob" // string | Searches across a range of columns including:   - custom_value1   - custom_value2   - custom_value3   - custom_value4  (optional)
	clientStatus := "?client_status=active,paused" // string | A comma separated list of invoice status strings. Valid options include:   - all - active   - paused   - completed    (optional)
	sort := "id|desc number|desc balance|asc" // string | Returns the list sorted by column in ascending or descending order. (optional)
	number := "?number=123" // string | Filters the list by number.  (optional)
	productKey := "?product_key=123" // string | Filters the list by product_key.  (optional)
	nextSendBetween := "?next_send_between=2025-01-01|2025-01-31" // string | Filters the list by next_send_between.  (optional)
	frequencyId := "?frequency_id=1,2,3" // string | Filters the list by frequency_id.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecurringInvoicesAPI.GetRecurringInvoices(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).ClientId(clientId).CreatedAt(createdAt).UpdatedAt(updatedAt).IsDeleted(isDeleted).FilterDeletedClients(filterDeletedClients).VendorId(vendorId).Filter(filter).ClientStatus(clientStatus).Sort(sort).Number(number).ProductKey(productKey).NextSendBetween(nextSendBetween).FrequencyId(frequencyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecurringInvoicesAPI.GetRecurringInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecurringInvoices`: GetRecurringInvoices200Response
	fmt.Fprintf(os.Stdout, "Response from `RecurringInvoicesAPI.GetRecurringInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecurringInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 
 **clientId** | **string** | Filters the entity list by client_id. Suitable when you only want the entities of a specific client.  &#x60;&#x60;&#x60;html GET /api/v1/invoices?client_id&#x3D;AxB7Hjk9 Returns only invoices for the specified client &#x60;&#x60;&#x60;  | 
 **createdAt** | **int32** | Filters the entity list by the created at timestamp. Parameter value can be a datetime string or unix timestamp  &#x60;&#x60;&#x60;html GET /api/v1/invoices?created_at&#x3D;2022-01-10 Returns entities created on January 10th, 2022 &#x60;&#x60;&#x60;  | 
 **updatedAt** | **int32** | Filters the entity list by the updated at timestamp. Parameter value can be a datetime string or unix timestamp  &#x60;&#x60;&#x60;html GET /api/v1/invoices?updated_at&#x3D;2022-01-10 Returns entities last updated on January 10th, 2022 &#x60;&#x60;&#x60;  | 
 **isDeleted** | **bool** | Filters the entity list by entities that have been deleted.  &#x60;&#x60;&#x60;html GET /api/v1/invoices?is_deleted&#x3D;true Returns only soft-deleted entities &#x60;&#x60;&#x60;  | 
 **filterDeletedClients** | **string** | Filters the entity list and only returns entities for clients that have not been deleted  &#x60;&#x60;&#x60;html GET /api/v1/invoices?filter_deleted_clients&#x3D;true Returns only invoices for active (non-deleted) clients &#x60;&#x60;&#x60;  | 
 **vendorId** | **string** | Filters the entity list by an associated vendor  &#x60;&#x60;&#x60;html GET /api/v1/purchases?vendor_id&#x3D;AxB7Hjk9 Returns only purchases for the specified vendor &#x60;&#x60;&#x60;  | 
 **filter** | **string** | Searches across a range of columns including:   - custom_value1   - custom_value2   - custom_value3   - custom_value4  | 
 **clientStatus** | **string** | A comma separated list of invoice status strings. Valid options include:   - all - active   - paused   - completed    | 
 **sort** | **string** | Returns the list sorted by column in ascending or descending order. | 
 **number** | **string** | Filters the list by number.  | 
 **productKey** | **string** | Filters the list by product_key.  | 
 **nextSendBetween** | **string** | Filters the list by next_send_between.  | 
 **frequencyId** | **string** | Filters the list by frequency_id.  | 

### Return type

[**GetRecurringInvoices200Response**](GetRecurringInvoices200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecurringInvoicesCreate

> RecurringInvoice GetRecurringInvoicesCreate(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Blank recurring invoice



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
	resp, r, err := apiClient.RecurringInvoicesAPI.GetRecurringInvoicesCreate(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecurringInvoicesAPI.GetRecurringInvoicesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecurringInvoicesCreate`: RecurringInvoice
	fmt.Fprintf(os.Stdout, "Response from `RecurringInvoicesAPI.GetRecurringInvoicesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecurringInvoicesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**RecurringInvoice**](RecurringInvoice.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowRecurringInvoice

> RecurringInvoice ShowRecurringInvoice(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Show recurring invoice



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
	id := "D2J234DFA" // string | The RecurringInvoice Hashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecurringInvoicesAPI.ShowRecurringInvoice(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecurringInvoicesAPI.ShowRecurringInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowRecurringInvoice`: RecurringInvoice
	fmt.Fprintf(os.Stdout, "Response from `RecurringInvoicesAPI.ShowRecurringInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The RecurringInvoice Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowRecurringInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**RecurringInvoice**](RecurringInvoice.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreRecurringInvoice

> RecurringInvoice StoreRecurringInvoice(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).RecurringInvoiceRequest(recurringInvoiceRequest).Include(include).Execute()

Create recurring invoice



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
	recurringInvoiceRequest := *openapiclient.NewRecurringInvoiceRequest() // RecurringInvoiceRequest | 
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecurringInvoicesAPI.StoreRecurringInvoice(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).RecurringInvoiceRequest(recurringInvoiceRequest).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecurringInvoicesAPI.StoreRecurringInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoreRecurringInvoice`: RecurringInvoice
	fmt.Fprintf(os.Stdout, "Response from `RecurringInvoicesAPI.StoreRecurringInvoice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreRecurringInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **recurringInvoiceRequest** | [**RecurringInvoiceRequest**](RecurringInvoiceRequest.md) |  | 
 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**RecurringInvoice**](RecurringInvoice.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRecurringInvoice

> RecurringInvoice UpdateRecurringInvoice(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Update recurring invoice



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
	id := "D2J234DFA" // string | The RecurringInvoice Hashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecurringInvoicesAPI.UpdateRecurringInvoice(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecurringInvoicesAPI.UpdateRecurringInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRecurringInvoice`: RecurringInvoice
	fmt.Fprintf(os.Stdout, "Response from `RecurringInvoicesAPI.UpdateRecurringInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The RecurringInvoice Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRecurringInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**RecurringInvoice**](RecurringInvoice.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadRecurringInvoice

> RecurringInvoice UploadRecurringInvoice(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Method(method).Documents(documents).Execute()

Add recurring invoice document



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
	id := "D2J234DFA" // string | The RecurringInvoice Hashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)
	method := "method_example" // string |  (optional)
	documents := []*os.File{"TODO"} // []*os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecurringInvoicesAPI.UploadRecurringInvoice(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Method(method).Documents(documents).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecurringInvoicesAPI.UploadRecurringInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadRecurringInvoice`: RecurringInvoice
	fmt.Fprintf(os.Stdout, "Response from `RecurringInvoicesAPI.UploadRecurringInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The RecurringInvoice Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadRecurringInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 
 **method** | **string** |  | 
 **documents** | **[]*os.File** |  | 

### Return type

[**RecurringInvoice**](RecurringInvoice.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

