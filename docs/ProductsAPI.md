# \ProductsAPI

All URIs are relative to *https://demo.invoiceninja.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkProducts**](ProductsAPI.md#BulkProducts) | **Post** /api/v1/products/bulk | Bulk product actions
[**DeleteProduct**](ProductsAPI.md#DeleteProduct) | **Delete** /api/v1/products/{id} | Delete product
[**EditProduct**](ProductsAPI.md#EditProduct) | **Get** /api/v1/products/{id}/edit | Edit product
[**GetProducts**](ProductsAPI.md#GetProducts) | **Get** /api/v1/products | List products
[**GetProductsCreate**](ProductsAPI.md#GetProductsCreate) | **Get** /api/v1/products/create | Blank product
[**ShowProduct**](ProductsAPI.md#ShowProduct) | **Get** /api/v1/products/{id} | Show product
[**StoreProduct**](ProductsAPI.md#StoreProduct) | **Post** /api/v1/products | Create Product
[**UpdateProduct**](ProductsAPI.md#UpdateProduct) | **Put** /api/v1/products/{id} | Update product
[**UploadProduct**](ProductsAPI.md#UploadProduct) | **Post** /api/v1/products/{id}/upload | Add product document



## BulkProducts

> Product BulkProducts(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).ProductBulkAction(productBulkAction).Index(index).Execute()

Bulk product actions



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
	productBulkAction := *openapiclient.NewProductBulkAction("archive", []string{"2J234DFA,D2J234DFA,D2J234DFA"}) // ProductBulkAction | Bulk action array
	index := "user" // string | Replaces the default response index from data to a user specific string  ie.  ```html   ?index=new_index ```  response is wrapped  ```json   {     'new_index' : [       .....       ]   } ```  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.BulkProducts(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).ProductBulkAction(productBulkAction).Index(index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.BulkProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkProducts`: Product
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.BulkProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **productBulkAction** | [**ProductBulkAction**](ProductBulkAction.md) | Bulk action array | 
 **index** | **string** | Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;  | 

### Return type

[**Product**](Product.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProduct

> DeleteProduct(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Delete product



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
	id := "D2J234DFA" // string | The Product Hashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProductsAPI.DeleteProduct(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.DeleteProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Product Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProductRequest struct via the builder pattern


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


## EditProduct

> Product EditProduct(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Edit product



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
	id := "D2J234DFA" // string | The Product Hashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.EditProduct(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.EditProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditProduct`: Product
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.EditProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Product Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**Product**](Product.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProducts

> GetProducts200Response GetProducts(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Status(status).ClientId(clientId).CreatedAt(createdAt).UpdatedAt(updatedAt).IsDeleted(isDeleted).FilterDeletedClients(filterDeletedClients).VendorId(vendorId).Filter(filter).ProductKey(productKey).Sort(sort).Execute()

List products



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
	filter := "bob" // string | Filter by product name (optional)
	productKey := "bob" // string | Filter by product key (optional)
	sort := "id|desc product_key|desc" // string | Returns the list sorted by column in ascending or descending order. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.GetProducts(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Status(status).ClientId(clientId).CreatedAt(createdAt).UpdatedAt(updatedAt).IsDeleted(isDeleted).FilterDeletedClients(filterDeletedClients).VendorId(vendorId).Filter(filter).ProductKey(productKey).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.GetProducts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProducts`: GetProducts200Response
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.GetProducts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProductsRequest struct via the builder pattern


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
 **filter** | **string** | Filter by product name | 
 **productKey** | **string** | Filter by product key | 
 **sort** | **string** | Returns the list sorted by column in ascending or descending order. | 

### Return type

[**GetProducts200Response**](GetProducts200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProductsCreate

> Product GetProductsCreate(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Blank product



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
	resp, r, err := apiClient.ProductsAPI.GetProductsCreate(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.GetProductsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProductsCreate`: Product
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.GetProductsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProductsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**Product**](Product.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowProduct

> Product ShowProduct(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Show product



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
	id := "D2J234DFA" // string | The Product Hashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.ShowProduct(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.ShowProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowProduct`: Product
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.ShowProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Product Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**Product**](Product.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreProduct

> Product StoreProduct(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).ProductRequest(productRequest).Include(include).Execute()

Create Product



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
	productRequest := *openapiclient.NewProductRequest() // ProductRequest | Product object that needs to be added to the company
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.StoreProduct(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).ProductRequest(productRequest).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.StoreProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoreProduct`: Product
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.StoreProduct`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **productRequest** | [**ProductRequest**](ProductRequest.md) | Product object that needs to be added to the company | 
 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**Product**](Product.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProduct

> Product UpdateProduct(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).ProductRequest(productRequest).Include(include).Execute()

Update product



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
	id := "D2J234DFA" // string | The Product Hashed ID
	productRequest := *openapiclient.NewProductRequest() // ProductRequest | Product object that needs to be added to the company
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.UpdateProduct(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).ProductRequest(productRequest).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.UpdateProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProduct`: Product
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.UpdateProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Product Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **productRequest** | [**ProductRequest**](ProductRequest.md) | Product object that needs to be added to the company | 
 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**Product**](Product.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadProduct

> Product UploadProduct(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Method(method).Documents(documents).Execute()

Add product document



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
	id := "D2J234DFA" // string | The Product Hashed ID
	include := "activities" // string | Include child relationships of the Client Object. ie ?include=documents,system_logs  ```html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  ```  Usage:  ```html /api/v1/clients?include=contacts,documents,activities ```  (optional)
	method := "method_example" // string |  (optional)
	documents := []interface{}{interface{}(123)} // []interface{} |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProductsAPI.UploadProduct(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Method(method).Documents(documents).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProductsAPI.UploadProduct``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadProduct`: Product
	fmt.Fprintf(os.Stdout, "Response from `ProductsAPI.UploadProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Product Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **include** | **string** | Include child relationships of the Client Object. ie ?include&#x3D;documents,system_logs  &#x60;&#x60;&#x60;html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  &#x60;&#x60;&#x60;  Usage:  &#x60;&#x60;&#x60;html /api/v1/clients?include&#x3D;contacts,documents,activities &#x60;&#x60;&#x60;  | 
 **method** | **string** |  | 
 **documents** | **[]interface{}** |  | 

### Return type

[**Product**](Product.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

