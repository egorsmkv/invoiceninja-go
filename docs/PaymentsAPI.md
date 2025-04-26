# \PaymentsAPI

All URIs are relative to *https://demo.invoiceninja.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionPayment**](PaymentsAPI.md#ActionPayment) | **Get** /api/v1/payments/{id}/{action} | Custom payment actions
[**BulkPayments**](PaymentsAPI.md#BulkPayments) | **Post** /api/v1/payments/bulk | Bulk payment actions
[**DeletePayment**](PaymentsAPI.md#DeletePayment) | **Delete** /api/v1/payments/{id} | Delete payment
[**EditPayment**](PaymentsAPI.md#EditPayment) | **Get** /api/v1/payments/{id}/edit | Edit payment
[**GetPayments**](PaymentsAPI.md#GetPayments) | **Get** /api/v1/payments | List payments
[**GetPaymentsCreate**](PaymentsAPI.md#GetPaymentsCreate) | **Get** /api/v1/payments/create | Blank payment
[**ShowPayment**](PaymentsAPI.md#ShowPayment) | **Get** /api/v1/payments/{id} | Show payment
[**StorePayment**](PaymentsAPI.md#StorePayment) | **Post** /api/v1/payments | Create payment
[**StoreRefund**](PaymentsAPI.md#StoreRefund) | **Post** /api/v1/payments/refund | Refund payment
[**UpdatePayment**](PaymentsAPI.md#UpdatePayment) | **Put** /api/v1/payments/{id} | Update payment
[**UploadPayment**](PaymentsAPI.md#UploadPayment) | **Post** /api/v1/payments/{id}/upload | Upload a payment document



## ActionPayment

> Payment ActionPayment(ctx, id, action).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Custom payment actions



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
	id := "D2J234DFA" // string | The Payment Hashed ID
	action := "clone_to_quote" // string | The action string to be performed
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.ActionPayment(context.Background(), id, action).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.ActionPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActionPayment`: Payment
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.ActionPayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Payment Hashed ID | 
**action** | **string** | The action string to be performed | 

### Other Parameters

Other parameters are passed through a pointer to a apiActionPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 


 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**Payment**](Payment.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkPayments

> Payment BulkPayments(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).RequestBody(requestBody).Index(index).Execute()

Bulk payment actions



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
	requestBody := []int32{int32(123)} // []int32 | User credentials
	index := "user" // string | Replaces the default response index from data to a user specific string  ie.  ```html   ?index=new_index ```  response is wrapped  ```json   {     'new_index' : [       .....       ]   } ```  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.BulkPayments(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).RequestBody(requestBody).Index(index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.BulkPayments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkPayments`: Payment
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.BulkPayments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **requestBody** | **[]int32** | User credentials | 
 **index** | **string** | Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;  | 

### Return type

[**Payment**](Payment.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePayment

> DeletePayment(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Delete payment



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
	id := "D2J234DFA" // string | The Payment Hashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PaymentsAPI.DeletePayment(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.DeletePayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Payment Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePaymentRequest struct via the builder pattern


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


## EditPayment

> Payment EditPayment(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Edit payment



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
	id := "D2J234DFA" // string | The Payment Hashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.EditPayment(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.EditPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditPayment`: Payment
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.EditPayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Payment Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**Payment**](Payment.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayments

> GetPayments200Response GetPayments(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Status(status).ClientId(clientId).CreatedAt(createdAt).UpdatedAt(updatedAt).IsDeleted(isDeleted).FilterDeletedClients(filterDeletedClients).VendorId(vendorId).Filter(filter).Number(number).Sort(sort).Execute()

List payments



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
	filter := "?filter=10" // string | Searches across a range of columns including:   - amount   - date   - custom_value1   - custom_value2   - custom_value3   - custom_value4  (optional)
	number := "?number=0001" // string | Search payments by payment number   (optional)
	sort := "id|desc number|desc balance|asc" // string | Returns the list sorted by column in ascending or descending order. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.GetPayments(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Status(status).ClientId(clientId).CreatedAt(createdAt).UpdatedAt(updatedAt).IsDeleted(isDeleted).FilterDeletedClients(filterDeletedClients).VendorId(vendorId).Filter(filter).Number(number).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.GetPayments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPayments`: GetPayments200Response
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.GetPayments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentsRequest struct via the builder pattern


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
 **filter** | **string** | Searches across a range of columns including:   - amount   - date   - custom_value1   - custom_value2   - custom_value3   - custom_value4  | 
 **number** | **string** | Search payments by payment number   | 
 **sort** | **string** | Returns the list sorted by column in ascending or descending order. | 

### Return type

[**GetPayments200Response**](GetPayments200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentsCreate

> Payment GetPaymentsCreate(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Blank payment



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
	resp, r, err := apiClient.PaymentsAPI.GetPaymentsCreate(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.GetPaymentsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPaymentsCreate`: Payment
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.GetPaymentsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**Payment**](Payment.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowPayment

> Payment ShowPayment(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Show payment



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
	id := "D2J234DFA" // string | The Payment Hashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.ShowPayment(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.ShowPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowPayment`: Payment
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.ShowPayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Payment Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**Payment**](Payment.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StorePayment

> Payment StorePayment(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).PaymentRequest(paymentRequest).Include(include).EmailReceipt(emailReceipt).Execute()

Create payment



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
	paymentRequest := *openapiclient.NewPaymentRequest() // PaymentRequest | The payment request can be as simple as a client id and amount, or as complex as a full payment with invoices and associated credits that should be applied.
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)
	emailReceipt := true // bool | If true, the payment will be emailed to the client. If false, no email will be sent. This will override any other email settings for payments.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.StorePayment(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).PaymentRequest(paymentRequest).Include(include).EmailReceipt(emailReceipt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.StorePayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StorePayment`: Payment
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.StorePayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStorePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **paymentRequest** | [**PaymentRequest**](PaymentRequest.md) | The payment request can be as simple as a client id and amount, or as complex as a full payment with invoices and associated credits that should be applied. | 
 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 
 **emailReceipt** | **bool** | If true, the payment will be emailed to the client. If false, no email will be sent. This will override any other email settings for payments.  | 

### Return type

[**Payment**](Payment.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreRefund

> Payment StoreRefund(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Payment(payment).Include(include).Execute()

Refund payment



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
	payment := *openapiclient.NewPayment() // Payment | The refund request
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.StoreRefund(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Payment(payment).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.StoreRefund``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoreRefund`: Payment
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.StoreRefund`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreRefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **payment** | [**Payment**](Payment.md) | The refund request | 
 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**Payment**](Payment.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePayment

> Payment UpdatePayment(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Update payment



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
	id := "D2J234DFA" // string | The Payment Hashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.UpdatePayment(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.UpdatePayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePayment`: Payment
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.UpdatePayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Payment Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**Payment**](Payment.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadPayment

> Payment UploadPayment(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Method(method).Documents(documents).Execute()

Upload a payment document



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
	id := "D2J234DFA" // string | The Payment Hashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)
	method := "method_example" // string |  (optional)
	documents := []*os.File{"TODO"} // []*os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.UploadPayment(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Method(method).Documents(documents).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.UploadPayment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadPayment`: Payment
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.UploadPayment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Payment Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 
 **method** | **string** |  | 
 **documents** | **[]*os.File** |  | 

### Return type

[**Payment**](Payment.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

