# \InvoicesAPI

All URIs are relative to *https://demo.invoiceninja.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionInvoice**](InvoicesAPI.md#ActionInvoice) | **Get** /api/v1/invoices/{id}/{action} | Custom invoice action
[**ApiV1InvoiceInvitationKeyDownloadGet**](InvoicesAPI.md#ApiV1InvoiceInvitationKeyDownloadGet) | **Get** /api/v1/invoice/{invitation_key}/download | Download invoice PDF
[**ApiV1InvoicesCreateGet**](InvoicesAPI.md#ApiV1InvoicesCreateGet) | **Get** /api/v1/invoices/create | Blank invoice
[**ApiV1InvoicesIdDeliveryNoteGet**](InvoicesAPI.md#ApiV1InvoicesIdDeliveryNoteGet) | **Get** /api/v1/invoices/{id}/delivery_note | Download delivery note
[**ApiV1InvoicesIdUploadPost**](InvoicesAPI.md#ApiV1InvoicesIdUploadPost) | **Post** /api/v1/invoices/{id}/upload | Add invoice document
[**BulkInvoices**](InvoicesAPI.md#BulkInvoices) | **Post** /api/v1/invoices/bulk | Bulk invoice actions
[**DeleteInvoice**](InvoicesAPI.md#DeleteInvoice) | **Delete** /api/v1/invoices/{id} | Delete invoice
[**EditInvoice**](InvoicesAPI.md#EditInvoice) | **Get** /api/v1/invoices/{id}/edit | Edit invoice
[**GetInvoices**](InvoicesAPI.md#GetInvoices) | **Get** /api/v1/invoices | List invoices
[**ShowInvoice**](InvoicesAPI.md#ShowInvoice) | **Get** /api/v1/invoices/{id} | Show invoice
[**StoreInvoice**](InvoicesAPI.md#StoreInvoice) | **Post** /api/v1/invoices | Create invoice
[**UpdateInvoice**](InvoicesAPI.md#UpdateInvoice) | **Put** /api/v1/invoices/{id} | Update invoice



## ActionInvoice

> Invoice ActionInvoice(ctx, id, action).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Custom invoice action



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
	id := "D2J234DFA" // string | The Invoice Hashed ID
	action := "clone_to_quote" // string | The action string to be performed
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.ActionInvoice(context.Background(), id, action).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.ActionInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionInvoice`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.ActionInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Invoice Hashed ID | 
**action** | **string** | The action string to be performed | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 


 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1InvoiceInvitationKeyDownloadGet

> ApiV1InvoiceInvitationKeyDownloadGet(ctx, invitationKey).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Download invoice PDF



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
	invitationKey := "D2J234DFA" // string | The Invoice Invitation Key
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InvoicesAPI.ApiV1InvoiceInvitationKeyDownloadGet(context.Background(), invitationKey).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.ApiV1InvoiceInvitationKeyDownloadGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationKey** | **string** | The Invoice Invitation Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1InvoiceInvitationKeyDownloadGetRequest struct via the builder pattern


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


## ApiV1InvoicesCreateGet

> Invoice ApiV1InvoicesCreateGet(ctx).Execute()

Blank invoice



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.ApiV1InvoicesCreateGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.ApiV1InvoicesCreateGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1InvoicesCreateGet`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.ApiV1InvoicesCreateGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1InvoicesCreateGetRequest struct via the builder pattern


### Return type

[**Invoice**](Invoice.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiV1InvoicesIdDeliveryNoteGet

> ApiV1InvoicesIdDeliveryNoteGet(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Download delivery note



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
	id := "D2J234DFA" // string | The Invoice Hahsed Id
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InvoicesAPI.ApiV1InvoicesIdDeliveryNoteGet(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.ApiV1InvoicesIdDeliveryNoteGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Invoice Hahsed Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1InvoicesIdDeliveryNoteGetRequest struct via the builder pattern


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


## ApiV1InvoicesIdUploadPost

> Invoice ApiV1InvoicesIdUploadPost(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Method(method).Documents(documents).Execute()

Add invoice document



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
	id := "D2J234DFA" // string | The Invoice Hashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)
	method := "method_example" // string |  (optional)
	documents := []*os.File{"TODO"} // []*os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.ApiV1InvoicesIdUploadPost(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Method(method).Documents(documents).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.ApiV1InvoicesIdUploadPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiV1InvoicesIdUploadPost`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.ApiV1InvoicesIdUploadPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Invoice Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV1InvoicesIdUploadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 
 **method** | **string** |  | 
 **documents** | **[]*os.File** |  | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkInvoices

> BulkInvoices(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).BulkInvoicesRequest(bulkInvoicesRequest).Index(index).Execute()

Bulk invoice actions



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
	bulkInvoicesRequest := *openapiclient.NewBulkInvoicesRequest() // BulkInvoicesRequest | Bulk action details
	index := "user" // string | Replaces the default response index from data to a user specific string  ie.  ```html   ?index=new_index ```  response is wrapped  ```json   {     'new_index' : [       .....       ]   } ```  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InvoicesAPI.BulkInvoices(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).BulkInvoicesRequest(bulkInvoicesRequest).Index(index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.BulkInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **bulkInvoicesRequest** | [**BulkInvoicesRequest**](BulkInvoicesRequest.md) | Bulk action details | 
 **index** | **string** | Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;  | 

### Return type

 (empty response body)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInvoice

> DeleteInvoice(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Delete invoice



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
	id := "D2J234DFA" // string | The Invoice Hashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InvoicesAPI.DeleteInvoice(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.DeleteInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Invoice Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInvoiceRequest struct via the builder pattern


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


## EditInvoice

> Invoice EditInvoice(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Edit invoice



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
	id := "D2J234DFA" // string | The Invoice Hashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.EditInvoice(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.EditInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditInvoice`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.EditInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Invoice Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoices

> GetInvoices200Response GetInvoices(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Status(status).ClientId(clientId).CreatedAt(createdAt).UpdatedAt(updatedAt).IsDeleted(isDeleted).FilterDeletedClients(filterDeletedClients).VendorId(vendorId).ClientStatus(clientStatus).Number(number).Filter(filter).WithoutDeletedClients(withoutDeletedClients).Overdue(overdue).Payable(payable).Sort(sort).PrivateNotes(privateNotes).Date(date).DateRange(dateRange).StatusId(statusId).Execute()

List invoices



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
	status := "?status=archived,deleted" // string | Filter the entity based on their status. ie active / archived / deleted. Format is a comma separated string with any of the following options:   - active - archived - deleted    ```html GET /api/v1/invoices?status=archived,deleted Returns only archived and deleted invoices ```  (optional)
	clientId := "?client_id={client_id}" // string | Filters the entity list by client_id. Suitable when you only want the entities of a specific client.  ```html GET /api/v1/invoices?client_id=AxB7Hjk9 Returns only invoices for the specified client ```  (optional)
	createdAt := int32(?created_at=2022-01-10) // int32 | Filters the entity list by the created at timestamp. Parameter value can be a datetime string or unix timestamp  ```html GET /api/v1/invoices?created_at=2022-01-10 Returns entities created on January 10th, 2022 ```  (optional)
	updatedAt := int32(?updated_at=2022-01-10) // int32 | Filters the entity list by the updated at timestamp. Parameter value can be a datetime string or unix timestamp  ```html GET /api/v1/invoices?updated_at=2022-01-10 Returns entities last updated on January 10th, 2022 ```  (optional)
	isDeleted := false // bool | Filters the entity list by entities that have been deleted.  ```html GET /api/v1/invoices?is_deleted=true Returns only soft-deleted entities ```  (optional)
	filterDeletedClients := "?filter_deleted_clients=true" // string | Filters the entity list and only returns entities for clients that have not been deleted  ```html GET /api/v1/invoices?filter_deleted_clients=true Returns only invoices for active (non-deleted) clients ```  (optional)
	vendorId := "?vendor_id={vendor_id}" // string | Filters the entity list by an associated vendor  ```html GET /api/v1/purchases?vendor_id=AxB7Hjk9 Returns only purchases for the specified vendor ```  (optional)
	clientStatus := "?client_status=paid,unpaid" // string | A comma separated list of invoice status strings. Valid options include:   - all - paid   - unpaid   - overdue     (optional)
	number := "?number=INV-001" // string | Search invoices by invoice number   (optional)
	filter := "?filter=bob" // string | Searches across a range of columns including:   - number   - po_number   - date   - amount   - balance   - custom_value1   - custom_value2   - custom_value3   - custom_value4 - client.name - client.contacts.[first_name, last_name, email] - line_items.[product_key, notes]  (optional)
	withoutDeletedClients := "?without_deleted_clients=" // string | Returns the invoice list without the invoices of deleted clients.  (optional)
	overdue := "?overdue=" // string | Returns the list of invoices that are overdue  (optional)
	payable := "?payable={client_id}" // string | Returns the invoice list that are payable for a defined client. Please note, you must pass the client_id as the value for this query parameter  (optional)
	sort := "id|desc number|desc balance|asc" // string | Returns the list sorted by column in ascending or descending order. (optional)
	privateNotes := "?private_notes=super secret" // string | Searches on the private_notes field of the invoices  (optional)
	date := "?date=2022-01-01" // string | Filters the invoices by invoice date returns a list of invoices after (and including) the date  (optional)
	dateRange := "?date_range=2022-01-01,2022-01-31" // string | Filters the invoices by invoice date returns a list of invoices between two dates  (optional)
	statusId := int32(?status_id=1) // int32 | Filters the invoices by status id  ```html   STATUS_DRAFT = 1;   STATUS_SENT = 2;   STATUS_PARTIAL = 3;   STATUS_PAID = 4;   STATUS_CANCELLED = 5;   STATUS_REVERSED = 6; ```  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.GetInvoices(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Status(status).ClientId(clientId).CreatedAt(createdAt).UpdatedAt(updatedAt).IsDeleted(isDeleted).FilterDeletedClients(filterDeletedClients).VendorId(vendorId).ClientStatus(clientStatus).Number(number).Filter(filter).WithoutDeletedClients(withoutDeletedClients).Overdue(overdue).Payable(payable).Sort(sort).PrivateNotes(privateNotes).Date(date).DateRange(dateRange).StatusId(statusId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoices`: GetInvoices200Response
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 
 **status** | **string** | Filter the entity based on their status. ie active / archived / deleted. Format is a comma separated string with any of the following options:   - active - archived - deleted    &#x60;&#x60;&#x60;html GET /api/v1/invoices?status&#x3D;archived,deleted Returns only archived and deleted invoices &#x60;&#x60;&#x60;  | 
 **clientId** | **string** | Filters the entity list by client_id. Suitable when you only want the entities of a specific client.  &#x60;&#x60;&#x60;html GET /api/v1/invoices?client_id&#x3D;AxB7Hjk9 Returns only invoices for the specified client &#x60;&#x60;&#x60;  | 
 **createdAt** | **int32** | Filters the entity list by the created at timestamp. Parameter value can be a datetime string or unix timestamp  &#x60;&#x60;&#x60;html GET /api/v1/invoices?created_at&#x3D;2022-01-10 Returns entities created on January 10th, 2022 &#x60;&#x60;&#x60;  | 
 **updatedAt** | **int32** | Filters the entity list by the updated at timestamp. Parameter value can be a datetime string or unix timestamp  &#x60;&#x60;&#x60;html GET /api/v1/invoices?updated_at&#x3D;2022-01-10 Returns entities last updated on January 10th, 2022 &#x60;&#x60;&#x60;  | 
 **isDeleted** | **bool** | Filters the entity list by entities that have been deleted.  &#x60;&#x60;&#x60;html GET /api/v1/invoices?is_deleted&#x3D;true Returns only soft-deleted entities &#x60;&#x60;&#x60;  | 
 **filterDeletedClients** | **string** | Filters the entity list and only returns entities for clients that have not been deleted  &#x60;&#x60;&#x60;html GET /api/v1/invoices?filter_deleted_clients&#x3D;true Returns only invoices for active (non-deleted) clients &#x60;&#x60;&#x60;  | 
 **vendorId** | **string** | Filters the entity list by an associated vendor  &#x60;&#x60;&#x60;html GET /api/v1/purchases?vendor_id&#x3D;AxB7Hjk9 Returns only purchases for the specified vendor &#x60;&#x60;&#x60;  | 
 **clientStatus** | **string** | A comma separated list of invoice status strings. Valid options include:   - all - paid   - unpaid   - overdue     | 
 **number** | **string** | Search invoices by invoice number   | 
 **filter** | **string** | Searches across a range of columns including:   - number   - po_number   - date   - amount   - balance   - custom_value1   - custom_value2   - custom_value3   - custom_value4 - client.name - client.contacts.[first_name, last_name, email] - line_items.[product_key, notes]  | 
 **withoutDeletedClients** | **string** | Returns the invoice list without the invoices of deleted clients.  | 
 **overdue** | **string** | Returns the list of invoices that are overdue  | 
 **payable** | **string** | Returns the invoice list that are payable for a defined client. Please note, you must pass the client_id as the value for this query parameter  | 
 **sort** | **string** | Returns the list sorted by column in ascending or descending order. | 
 **privateNotes** | **string** | Searches on the private_notes field of the invoices  | 
 **date** | **string** | Filters the invoices by invoice date returns a list of invoices after (and including) the date  | 
 **dateRange** | **string** | Filters the invoices by invoice date returns a list of invoices between two dates  | 
 **statusId** | **int32** | Filters the invoices by status id  &#x60;&#x60;&#x60;html   STATUS_DRAFT &#x3D; 1;   STATUS_SENT &#x3D; 2;   STATUS_PARTIAL &#x3D; 3;   STATUS_PAID &#x3D; 4;   STATUS_CANCELLED &#x3D; 5;   STATUS_REVERSED &#x3D; 6; &#x60;&#x60;&#x60;  | 

### Return type

[**GetInvoices200Response**](GetInvoices200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowInvoice

> Invoice ShowInvoice(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Show invoice



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
	id := "D2J234DFA" // string | The Invoice Hashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.ShowInvoice(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.ShowInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowInvoice`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.ShowInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Invoice Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreInvoice

> Invoice StoreInvoice(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).InvoiceRequest(invoiceRequest).Include(include).Execute()

Create invoice



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
	invoiceRequest := *openapiclient.NewInvoiceRequest("Ht5N9cX3jK") // InvoiceRequest | 
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.StoreInvoice(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).InvoiceRequest(invoiceRequest).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.StoreInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoreInvoice`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.StoreInvoice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **invoiceRequest** | [**InvoiceRequest**](InvoiceRequest.md) |  | 
 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInvoice

> Invoice UpdateInvoice(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Update invoice



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
	id := "D2J234DFA" // string | The Invoice Hashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesAPI.UpdateInvoice(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.UpdateInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInvoice`: Invoice
	fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.UpdateInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Invoice Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**Invoice**](Invoice.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

