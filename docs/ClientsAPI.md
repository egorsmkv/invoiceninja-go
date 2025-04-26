# \ClientsAPI

All URIs are relative to *https://demo.invoiceninja.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkClients**](ClientsAPI.md#BulkClients) | **Post** /api/v1/clients/bulk | Bulk client actions
[**ClientStatement**](ClientsAPI.md#ClientStatement) | **Post** /api/v1/client_statement | Client statement PDF
[**DeleteClient**](ClientsAPI.md#DeleteClient) | **Delete** /api/v1/clients/{id} | Delete client
[**EditClient**](ClientsAPI.md#EditClient) | **Get** /api/v1/clients/{id}/edit | Edit Client
[**GetClients**](ClientsAPI.md#GetClients) | **Get** /api/v1/clients | List clients 
[**GetClientsCreate**](ClientsAPI.md#GetClientsCreate) | **Get** /api/v1/clients/create | Blank Client
[**MergeClient**](ClientsAPI.md#MergeClient) | **Post** /api/v1/clients/{id}/{mergeable_client_hashed_id}/merge | Merge client
[**PurgeClient**](ClientsAPI.md#PurgeClient) | **Post** /api/v1/clients/{id}/purge | Purge client
[**ReactivateEmail**](ClientsAPI.md#ReactivateEmail) | **Post** /api/v1/reactivate_email/{bounce_id} | Removes email suppression of a user in the system
[**ShowClient**](ClientsAPI.md#ShowClient) | **Get** /api/v1/clients/{id} | Show client
[**StoreClient**](ClientsAPI.md#StoreClient) | **Post** /api/v1/clients | Create client
[**UpdateClient**](ClientsAPI.md#UpdateClient) | **Put** /api/v1/clients/{id} | Update client
[**UpdateClientTaxData**](ClientsAPI.md#UpdateClientTaxData) | **Post** /api/v1/clients/{client}/updateTaxData | Update tax data
[**UploadClient**](ClientsAPI.md#UploadClient) | **Post** /api/v1/clients/{id}/upload | Add client document



## BulkClients

> Client BulkClients(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).GenericBulkAction(genericBulkAction).Index(index).Include(include).Execute()

Bulk client actions



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
	include := "activities" // string | Include child relationships of the Client Object. ie ?include=documents,system_logs  ```html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  ```  Usage:  ```html /api/v1/clients?include=contacts,documents,activities ```  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.BulkClients(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).GenericBulkAction(genericBulkAction).Index(index).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.BulkClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkClients`: Client
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.BulkClients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **genericBulkAction** | [**GenericBulkAction**](GenericBulkAction.md) | Bulk action array | 
 **index** | **string** | Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;  | 
 **include** | **string** | Include child relationships of the Client Object. ie ?include&#x3D;documents,system_logs  &#x60;&#x60;&#x60;html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  &#x60;&#x60;&#x60;  Usage:  &#x60;&#x60;&#x60;html /api/v1/clients?include&#x3D;contacts,documents,activities &#x60;&#x60;&#x60;  | 

### Return type

[**Client**](Client.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClientStatement

> Client ClientStatement(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).ClientStatementRequest(clientStatementRequest).Index(index).Include(include).Execute()

Client statement PDF



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
	clientStatementRequest := *openapiclient.NewClientStatementRequest() // ClientStatementRequest | Statement Options
	index := "user" // string | Replaces the default response index from data to a user specific string  ie.  ```html   ?index=new_index ```  response is wrapped  ```json   {     'new_index' : [       .....       ]   } ```  (optional)
	include := "activities" // string | Include child relationships of the Client Object. ie ?include=documents,system_logs  ```html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  ```  Usage:  ```html /api/v1/clients?include=contacts,documents,activities ```  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.ClientStatement(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).ClientStatementRequest(clientStatementRequest).Index(index).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.ClientStatement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClientStatement`: Client
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.ClientStatement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClientStatementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **clientStatementRequest** | [**ClientStatementRequest**](ClientStatementRequest.md) | Statement Options | 
 **index** | **string** | Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;  | 
 **include** | **string** | Include child relationships of the Client Object. ie ?include&#x3D;documents,system_logs  &#x60;&#x60;&#x60;html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  &#x60;&#x60;&#x60;  Usage:  &#x60;&#x60;&#x60;html /api/v1/clients?include&#x3D;contacts,documents,activities &#x60;&#x60;&#x60;  | 

### Return type

[**Client**](Client.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClient

> DeleteClient(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Index(index).Include(include).Execute()

Delete client



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
	id := "D2J234DFA" // string | The Client Hashed ID
	index := "user" // string | Replaces the default response index from data to a user specific string  ie.  ```html   ?index=new_index ```  response is wrapped  ```json   {     'new_index' : [       .....       ]   } ```  (optional)
	include := "activities" // string | Include child relationships of the Client Object. ie ?include=documents,system_logs  ```html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  ```  Usage:  ```html /api/v1/clients?include=contacts,documents,activities ```  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientsAPI.DeleteClient(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Index(index).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.DeleteClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Client Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **index** | **string** | Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;  | 
 **include** | **string** | Include child relationships of the Client Object. ie ?include&#x3D;documents,system_logs  &#x60;&#x60;&#x60;html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  &#x60;&#x60;&#x60;  Usage:  &#x60;&#x60;&#x60;html /api/v1/clients?include&#x3D;contacts,documents,activities &#x60;&#x60;&#x60;  | 

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


## EditClient

> Client EditClient(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Index(index).Include(include).Execute()

Edit Client



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
	id := "D2J234DFA" // string | The Client Hashed ID
	index := "user" // string | Replaces the default response index from data to a user specific string  ie.  ```html   ?index=new_index ```  response is wrapped  ```json   {     'new_index' : [       .....       ]   } ```  (optional)
	include := "activities" // string | Include child relationships of the Client Object. ie ?include=documents,system_logs  ```html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  ```  Usage:  ```html /api/v1/clients?include=contacts,documents,activities ```  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.EditClient(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Index(index).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.EditClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditClient`: Client
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.EditClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Client Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **index** | **string** | Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;  | 
 **include** | **string** | Include child relationships of the Client Object. ie ?include&#x3D;documents,system_logs  &#x60;&#x60;&#x60;html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  &#x60;&#x60;&#x60;  Usage:  &#x60;&#x60;&#x60;html /api/v1/clients?include&#x3D;contacts,documents,activities &#x60;&#x60;&#x60;  | 

### Return type

[**Client**](Client.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClients

> Client GetClients(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Index(index).Status(status).CreatedAt(createdAt).UpdatedAt(updatedAt).IsDeleted(isDeleted).FilterDeletedClients(filterDeletedClients).Name(name).Balance(balance).BetweenBalance(betweenBalance).Email(email).IdNumber(idNumber).Number(number).Filter(filter).Sort(sort).Group(group).ClientId(clientId).Execute()

List clients 



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
	include := "activities" // string | Include child relationships of the Client Object. ie ?include=documents,system_logs  ```html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  ```  Usage:  ```html /api/v1/clients?include=contacts,documents,activities ```  (optional)
	index := "user" // string | Replaces the default response index from data to a user specific string  ie.  ```html   ?index=new_index ```  response is wrapped  ```json   {     'new_index' : [       .....       ]   } ```  (optional)
	status := "?status=archived,deleted" // string | Filter the entity based on their status. ie active / archived / deleted. Format is a comma separated string with any of the following options:   - active - archived - deleted    ```html GET /api/v1/invoices?status=archived,deleted Returns only archived and deleted invoices ```  (optional)
	createdAt := int32(?created_at=2022-01-10) // int32 | Filters the entity list by the created at timestamp. Parameter value can be a datetime string or unix timestamp  ```html GET /api/v1/invoices?created_at=2022-01-10 Returns entities created on January 10th, 2022 ```  (optional)
	updatedAt := int32(?updated_at=2022-01-10) // int32 | Filters the entity list by the updated at timestamp. Parameter value can be a datetime string or unix timestamp  ```html GET /api/v1/invoices?updated_at=2022-01-10 Returns entities last updated on January 10th, 2022 ```  (optional)
	isDeleted := false // bool | Filters the entity list by entities that have been deleted.  ```html GET /api/v1/invoices?is_deleted=true Returns only soft-deleted entities ```  (optional)
	filterDeletedClients := "?filter_deleted_clients=true" // string | Filters the entity list and only returns entities for clients that have not been deleted  ```html GET /api/v1/invoices?filter_deleted_clients=true Returns only invoices for active (non-deleted) clients ```  (optional)
	name := "bob" // string | Filter by client name    ```html ?name=bob ```  (optional)
	balance := "lt:10" // string | Filter by client balance, format uses an operator and value separated by a colon. lt,lte, gt, gte, eq  ```html ?balance=lt:10 ```  ie all clients whose balance is less than 10  (optional)
	betweenBalance := "10:100" // string | Filter between client balances, format uses two values separated by a colon  ```html ?between_balance=10:100 ```  (optional)
	email := "bob@gmail.com" // string | Filter by client email  ```html ?email=bob@gmail.com ```  (optional)
	idNumber := "a1039883" // string | Filter by client id_number  ```html ?id_number=0001 ```  (optional)
	number := "a1039883" // string | Filter by client number  ```html ?number=0002 ```  (optional)
	filter := "a1039883" // string | Broad filter which targets multiple client columns:      ```html     name,      id_number,      contact.first_name      contact.last_name,      contact.email,      contact.phone     custom_value1,     custom_value2,     custom_value3,     custom_value4,   ```    ```html   ?filter=Bobby   ```  (optional)
	sort := "id|desc name|desc balance|asc" // string | Returns the list sorted by column in ascending or descending order.  Ensure you use column | direction, ie.  ```html   ?sort=id|desc ```  (optional)
	group := "X89sjd8" // string | Returns the list of clients assigned to {group_id}  ```html   ?group=X89sjd8 ```  (optional)
	clientId := "X89sjd8" // string | Returns the list of clients with {client_id} - proxy call to retrieve a client_id wrapped in an array  ```html   ?client_id=X89sjd8 ```  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.GetClients(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Index(index).Status(status).CreatedAt(createdAt).UpdatedAt(updatedAt).IsDeleted(isDeleted).FilterDeletedClients(filterDeletedClients).Name(name).Balance(balance).BetweenBalance(betweenBalance).Email(email).IdNumber(idNumber).Number(number).Filter(filter).Sort(sort).Group(group).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClients`: Client
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetClients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **include** | **string** | Include child relationships of the Client Object. ie ?include&#x3D;documents,system_logs  &#x60;&#x60;&#x60;html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  &#x60;&#x60;&#x60;  Usage:  &#x60;&#x60;&#x60;html /api/v1/clients?include&#x3D;contacts,documents,activities &#x60;&#x60;&#x60;  | 
 **index** | **string** | Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;  | 
 **status** | **string** | Filter the entity based on their status. ie active / archived / deleted. Format is a comma separated string with any of the following options:   - active - archived - deleted    &#x60;&#x60;&#x60;html GET /api/v1/invoices?status&#x3D;archived,deleted Returns only archived and deleted invoices &#x60;&#x60;&#x60;  | 
 **createdAt** | **int32** | Filters the entity list by the created at timestamp. Parameter value can be a datetime string or unix timestamp  &#x60;&#x60;&#x60;html GET /api/v1/invoices?created_at&#x3D;2022-01-10 Returns entities created on January 10th, 2022 &#x60;&#x60;&#x60;  | 
 **updatedAt** | **int32** | Filters the entity list by the updated at timestamp. Parameter value can be a datetime string or unix timestamp  &#x60;&#x60;&#x60;html GET /api/v1/invoices?updated_at&#x3D;2022-01-10 Returns entities last updated on January 10th, 2022 &#x60;&#x60;&#x60;  | 
 **isDeleted** | **bool** | Filters the entity list by entities that have been deleted.  &#x60;&#x60;&#x60;html GET /api/v1/invoices?is_deleted&#x3D;true Returns only soft-deleted entities &#x60;&#x60;&#x60;  | 
 **filterDeletedClients** | **string** | Filters the entity list and only returns entities for clients that have not been deleted  &#x60;&#x60;&#x60;html GET /api/v1/invoices?filter_deleted_clients&#x3D;true Returns only invoices for active (non-deleted) clients &#x60;&#x60;&#x60;  | 
 **name** | **string** | Filter by client name    &#x60;&#x60;&#x60;html ?name&#x3D;bob &#x60;&#x60;&#x60;  | 
 **balance** | **string** | Filter by client balance, format uses an operator and value separated by a colon. lt,lte, gt, gte, eq  &#x60;&#x60;&#x60;html ?balance&#x3D;lt:10 &#x60;&#x60;&#x60;  ie all clients whose balance is less than 10  | 
 **betweenBalance** | **string** | Filter between client balances, format uses two values separated by a colon  &#x60;&#x60;&#x60;html ?between_balance&#x3D;10:100 &#x60;&#x60;&#x60;  | 
 **email** | **string** | Filter by client email  &#x60;&#x60;&#x60;html ?email&#x3D;bob@gmail.com &#x60;&#x60;&#x60;  | 
 **idNumber** | **string** | Filter by client id_number  &#x60;&#x60;&#x60;html ?id_number&#x3D;0001 &#x60;&#x60;&#x60;  | 
 **number** | **string** | Filter by client number  &#x60;&#x60;&#x60;html ?number&#x3D;0002 &#x60;&#x60;&#x60;  | 
 **filter** | **string** | Broad filter which targets multiple client columns:      &#x60;&#x60;&#x60;html     name,      id_number,      contact.first_name      contact.last_name,      contact.email,      contact.phone     custom_value1,     custom_value2,     custom_value3,     custom_value4,   &#x60;&#x60;&#x60;    &#x60;&#x60;&#x60;html   ?filter&#x3D;Bobby   &#x60;&#x60;&#x60;  | 
 **sort** | **string** | Returns the list sorted by column in ascending or descending order.  Ensure you use column | direction, ie.  &#x60;&#x60;&#x60;html   ?sort&#x3D;id|desc &#x60;&#x60;&#x60;  | 
 **group** | **string** | Returns the list of clients assigned to {group_id}  &#x60;&#x60;&#x60;html   ?group&#x3D;X89sjd8 &#x60;&#x60;&#x60;  | 
 **clientId** | **string** | Returns the list of clients with {client_id} - proxy call to retrieve a client_id wrapped in an array  &#x60;&#x60;&#x60;html   ?client_id&#x3D;X89sjd8 &#x60;&#x60;&#x60;  | 

### Return type

[**Client**](Client.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientsCreate

> Client GetClientsCreate(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Index(index).Include(include).Execute()

Blank Client



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
	include := "activities" // string | Include child relationships of the Client Object. ie ?include=documents,system_logs  ```html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  ```  Usage:  ```html /api/v1/clients?include=contacts,documents,activities ```  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.GetClientsCreate(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Index(index).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.GetClientsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClientsCreate`: Client
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.GetClientsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClientsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **index** | **string** | Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;  | 
 **include** | **string** | Include child relationships of the Client Object. ie ?include&#x3D;documents,system_logs  &#x60;&#x60;&#x60;html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  &#x60;&#x60;&#x60;  Usage:  &#x60;&#x60;&#x60;html /api/v1/clients?include&#x3D;contacts,documents,activities &#x60;&#x60;&#x60;  | 

### Return type

[**Client**](Client.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MergeClient

> MergeClient(ctx, id, mergeableClientHashedId).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).XAPIPASSWORD(xAPIPASSWORD).Index(index).Include(include).Execute()

Merge client



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
	id := "D2J234DFA" // string | The Client Hashed ID
	mergeableClientHashedId := "D2J234DFA" // string | The Mergeable Client Hashed ID
	index := "user" // string | Replaces the default response index from data to a user specific string  ie.  ```html   ?index=new_index ```  response is wrapped  ```json   {     'new_index' : [       .....       ]   } ```  (optional)
	include := "activities" // string | Include child relationships of the Client Object. ie ?include=documents,system_logs  ```html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  ```  Usage:  ```html /api/v1/clients?include=contacts,documents,activities ```  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientsAPI.MergeClient(context.Background(), id, mergeableClientHashedId).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).XAPIPASSWORD(xAPIPASSWORD).Index(index).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.MergeClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Client Hashed ID | 
**mergeableClientHashedId** | **string** | The Mergeable Client Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiMergeClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **xAPIPASSWORD** | **string** | The login password when challenged on certain protected routes | 


 **index** | **string** | Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;  | 
 **include** | **string** | Include child relationships of the Client Object. ie ?include&#x3D;documents,system_logs  &#x60;&#x60;&#x60;html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  &#x60;&#x60;&#x60;  Usage:  &#x60;&#x60;&#x60;html /api/v1/clients?include&#x3D;contacts,documents,activities &#x60;&#x60;&#x60;  | 

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


## PurgeClient

> PurgeClient(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).XAPIPASSWORD(xAPIPASSWORD).Execute()

Purge client



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
	id := "D2J234DFA" // string | The Client Hashed ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientsAPI.PurgeClient(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).XAPIPASSWORD(xAPIPASSWORD).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.PurgeClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Client Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPurgeClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **xAPIPASSWORD** | **string** | The login password when challenged on certain protected routes | 


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


## ReactivateEmail

> ReactivateEmail(ctx, bounceId).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Index(index).Include(include).Execute()

Removes email suppression of a user in the system



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
	bounceId := "123243" // string | The postmark Bounce ID reference
	index := "user" // string | Replaces the default response index from data to a user specific string  ie.  ```html   ?index=new_index ```  response is wrapped  ```json   {     'new_index' : [       .....       ]   } ```  (optional)
	include := "activities" // string | Include child relationships of the Client Object. ie ?include=documents,system_logs  ```html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  ```  Usage:  ```html /api/v1/clients?include=contacts,documents,activities ```  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientsAPI.ReactivateEmail(context.Background(), bounceId).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Index(index).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.ReactivateEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bounceId** | **string** | The postmark Bounce ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiReactivateEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **index** | **string** | Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;  | 
 **include** | **string** | Include child relationships of the Client Object. ie ?include&#x3D;documents,system_logs  &#x60;&#x60;&#x60;html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  &#x60;&#x60;&#x60;  Usage:  &#x60;&#x60;&#x60;html /api/v1/clients?include&#x3D;contacts,documents,activities &#x60;&#x60;&#x60;  | 

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


## ShowClient

> Client ShowClient(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Index(index).Include(include).Execute()

Show client



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
	id := "D2J234DFA" // string | The Client Hashed ID
	index := "user" // string | Replaces the default response index from data to a user specific string  ie.  ```html   ?index=new_index ```  response is wrapped  ```json   {     'new_index' : [       .....       ]   } ```  (optional)
	include := "activities" // string | Include child relationships of the Client Object. ie ?include=documents,system_logs  ```html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  ```  Usage:  ```html /api/v1/clients?include=contacts,documents,activities ```  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.ShowClient(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Index(index).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.ShowClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowClient`: Client
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.ShowClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Client Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **index** | **string** | Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;  | 
 **include** | **string** | Include child relationships of the Client Object. ie ?include&#x3D;documents,system_logs  &#x60;&#x60;&#x60;html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  &#x60;&#x60;&#x60;  Usage:  &#x60;&#x60;&#x60;html /api/v1/clients?include&#x3D;contacts,documents,activities &#x60;&#x60;&#x60;  | 

### Return type

[**Client**](Client.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreClient

> Client StoreClient(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).ClientRequest(clientRequest).Index(index).Include(include).Execute()

Create client



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
	clientRequest := *openapiclient.NewClientRequest([]openapiclient.ClientContactRequest{*openapiclient.NewClientContactRequest()}, int32(123)) // ClientRequest | 
	index := "user" // string | Replaces the default response index from data to a user specific string  ie.  ```html   ?index=new_index ```  response is wrapped  ```json   {     'new_index' : [       .....       ]   } ```  (optional)
	include := "activities" // string | Include child relationships of the Client Object. ie ?include=documents,system_logs  ```html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  ```  Usage:  ```html /api/v1/clients?include=contacts,documents,activities ```  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.StoreClient(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).ClientRequest(clientRequest).Index(index).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.StoreClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoreClient`: Client
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.StoreClient`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **clientRequest** | [**ClientRequest**](ClientRequest.md) |  | 
 **index** | **string** | Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;  | 
 **include** | **string** | Include child relationships of the Client Object. ie ?include&#x3D;documents,system_logs  &#x60;&#x60;&#x60;html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  &#x60;&#x60;&#x60;  Usage:  &#x60;&#x60;&#x60;html /api/v1/clients?include&#x3D;contacts,documents,activities &#x60;&#x60;&#x60;  | 

### Return type

[**Client**](Client.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClient

> Client UpdateClient(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).ClientRequest(clientRequest).Index(index).Include(include).Execute()

Update client



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
	id := "D2J234DFA" // string | The Client Hashed ID
	clientRequest := *openapiclient.NewClientRequest([]openapiclient.ClientContactRequest{*openapiclient.NewClientContactRequest()}, int32(123)) // ClientRequest | Client object that needs to be updated
	index := "user" // string | Replaces the default response index from data to a user specific string  ie.  ```html   ?index=new_index ```  response is wrapped  ```json   {     'new_index' : [       .....       ]   } ```  (optional)
	include := "activities" // string | Include child relationships of the Client Object. ie ?include=documents,system_logs  ```html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  ```  Usage:  ```html /api/v1/clients?include=contacts,documents,activities ```  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.UpdateClient(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).ClientRequest(clientRequest).Index(index).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.UpdateClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateClient`: Client
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.UpdateClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Client Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **clientRequest** | [**ClientRequest**](ClientRequest.md) | Client object that needs to be updated | 
 **index** | **string** | Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;  | 
 **include** | **string** | Include child relationships of the Client Object. ie ?include&#x3D;documents,system_logs  &#x60;&#x60;&#x60;html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  &#x60;&#x60;&#x60;  Usage:  &#x60;&#x60;&#x60;html /api/v1/clients?include&#x3D;contacts,documents,activities &#x60;&#x60;&#x60;  | 

### Return type

[**Client**](Client.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClientTaxData

> UpdateClientTaxData(ctx, client).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Index(index).Include(include).Execute()

Update tax data



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
	client := "V2J234DFA" // string | The Client Hashed ID reference
	index := "user" // string | Replaces the default response index from data to a user specific string  ie.  ```html   ?index=new_index ```  response is wrapped  ```json   {     'new_index' : [       .....       ]   } ```  (optional)
	include := "activities" // string | Include child relationships of the Client Object. ie ?include=documents,system_logs  ```html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  ```  Usage:  ```html /api/v1/clients?include=contacts,documents,activities ```  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClientsAPI.UpdateClientTaxData(context.Background(), client).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Index(index).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.UpdateClientTaxData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**client** | **string** | The Client Hashed ID reference | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClientTaxDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **index** | **string** | Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;  | 
 **include** | **string** | Include child relationships of the Client Object. ie ?include&#x3D;documents,system_logs  &#x60;&#x60;&#x60;html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  &#x60;&#x60;&#x60;  Usage:  &#x60;&#x60;&#x60;html /api/v1/clients?include&#x3D;contacts,documents,activities &#x60;&#x60;&#x60;  | 

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


## UploadClient

> Client UploadClient(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Index(index).Include(include).Method(method).Documents(documents).Execute()

Add client document



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
	id := "D2J234DFA" // string | The Client Hashed ID
	index := "user" // string | Replaces the default response index from data to a user specific string  ie.  ```html   ?index=new_index ```  response is wrapped  ```json   {     'new_index' : [       .....       ]   } ```  (optional)
	include := "activities" // string | Include child relationships of the Client Object. ie ?include=documents,system_logs  ```html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  ```  Usage:  ```html /api/v1/clients?include=contacts,documents,activities ```  (optional)
	method := "method_example" // string |  (optional)
	documents := []*os.File{"TODO"} // []*os.File | Array of files to upload. The files should be sent with the key name 'documents[]'.  Supported file types: - PDF (.pdf) - Word (.doc, .docx) - Excel (.xls, .xlsx) - Images (.jpg, .jpeg, .png) - Text (.txt)  Maximum file size: 20MB per file  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClientsAPI.UploadClient(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Index(index).Include(include).Method(method).Documents(documents).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClientsAPI.UploadClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadClient`: Client
	fmt.Fprintf(os.Stdout, "Response from `ClientsAPI.UploadClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Client Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **index** | **string** | Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;  | 
 **include** | **string** | Include child relationships of the Client Object. ie ?include&#x3D;documents,system_logs  &#x60;&#x60;&#x60;html Available includes:  contacts [All contacts related to the client] documents [All documents related to the client] gateway_tokens [All payment tokens related to the client] activities [All activities related to the client] ledger [The client ledger] system_logs [System logs related to the client] group_settings [The group settings object the client is assigned to]  &#x60;&#x60;&#x60;  Usage:  &#x60;&#x60;&#x60;html /api/v1/clients?include&#x3D;contacts,documents,activities &#x60;&#x60;&#x60;  | 
 **method** | **string** |  | 
 **documents** | **[]*os.File** | Array of files to upload. The files should be sent with the key name &#39;documents[]&#39;.  Supported file types: - PDF (.pdf) - Word (.doc, .docx) - Excel (.xls, .xlsx) - Images (.jpg, .jpeg, .png) - Text (.txt)  Maximum file size: 20MB per file  | 

### Return type

[**Client**](Client.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

