# \BankIntegrationsAPI

All URIs are relative to *https://demo.invoiceninja.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkBankIntegrations**](BankIntegrationsAPI.md#BulkBankIntegrations) | **Post** /api/v1/bank_integrations/bulk | Performs bulk actions on an array of bank_integrations
[**DeleteBankIntegration**](BankIntegrationsAPI.md#DeleteBankIntegration) | **Delete** /api/v1/bank_integrations/{id} | Deletes a bank_integration
[**EditBankIntegration**](BankIntegrationsAPI.md#EditBankIntegration) | **Get** /api/v1/bank_integrations/{id}/edit | Shows a bank_integration for editing
[**GetAccountTransactions**](BankIntegrationsAPI.md#GetAccountTransactions) | **Post** /api/v1/bank_integrations/get_transactions/account_id | Retrieve transactions for a account
[**GetBankIntegrations**](BankIntegrationsAPI.md#GetBankIntegrations) | **Get** /api/v1/bank_integrations | Returns a list of Bank Integrations
[**GetBankIntegrationsCreate**](BankIntegrationsAPI.md#GetBankIntegrationsCreate) | **Get** /api/v1/bank_integrations/create | Gets a new blank bank_integration object
[**GetRefreshAccounts**](BankIntegrationsAPI.md#GetRefreshAccounts) | **Post** /api/v1/bank_integrations/refresh_accounts | Gets the list of accounts from the remote server
[**GetRemoveAccount**](BankIntegrationsAPI.md#GetRemoveAccount) | **Post** /api/v1/bank_integrations/remove_account/account_id | Removes an account from the integration
[**ShowBankIntegration**](BankIntegrationsAPI.md#ShowBankIntegration) | **Get** /api/v1/bank_integrations/{id} | Shows a bank_integration
[**StoreBankIntegration**](BankIntegrationsAPI.md#StoreBankIntegration) | **Post** /api/v1/bank_integrations | Adds a bank_integration
[**UpdateBankIntegration**](BankIntegrationsAPI.md#UpdateBankIntegration) | **Put** /api/v1/bank_integrations/{id} | Updates a bank_integration



## BulkBankIntegrations

> BulkBankIntegrations(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).RequestBody(requestBody).Index(index).Execute()

Performs bulk actions on an array of bank_integrations



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
	requestBody := []int32{int32(123)} // []int32 | Action paramters
	index := "user" // string | Replaces the default response index from data to a user specific string  ie.  ```html   ?index=new_index ```  response is wrapped  ```json   {     'new_index' : [       .....       ]   } ```  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BankIntegrationsAPI.BulkBankIntegrations(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).RequestBody(requestBody).Index(index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankIntegrationsAPI.BulkBankIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkBankIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **requestBody** | **[]int32** | Action paramters | 
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


## DeleteBankIntegration

> DeleteBankIntegration(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Deletes a bank_integration



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
	id := "D2J234DFA" // string | The BankIntegration Hashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BankIntegrationsAPI.DeleteBankIntegration(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankIntegrationsAPI.DeleteBankIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The BankIntegration Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBankIntegrationRequest struct via the builder pattern


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


## EditBankIntegration

> BankIntegration EditBankIntegration(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Shows a bank_integration for editing



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
	id := "D2J234DFA" // string | The BankIntegration Hashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankIntegrationsAPI.EditBankIntegration(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankIntegrationsAPI.EditBankIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditBankIntegration`: BankIntegration
	fmt.Fprintf(os.Stdout, "Response from `BankIntegrationsAPI.EditBankIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The BankIntegration Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditBankIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**BankIntegration**](BankIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountTransactions

> BankIntegration GetAccountTransactions(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Retrieve transactions for a account



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
	resp, r, err := apiClient.BankIntegrationsAPI.GetAccountTransactions(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankIntegrationsAPI.GetAccountTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountTransactions`: BankIntegration
	fmt.Fprintf(os.Stdout, "Response from `BankIntegrationsAPI.GetAccountTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**BankIntegration**](BankIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankIntegrations

> GetBankIntegrations200Response GetBankIntegrations(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Index(index).Rows(rows).Execute()

Returns a list of Bank Integrations



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
	include := "company" // string | Include child relations of the BankIntegration object. Format is comma separated. (optional)
	index := "user" // string | Replaces the default response index from data to a user specific string  ie.  ```html   ?index=new_index ```  response is wrapped  ```json   {     'new_index' : [       .....       ]   } ```  (optional)
	rows := float32(50) // float32 | The number of bank integrations to return (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankIntegrationsAPI.GetBankIntegrations(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Index(index).Rows(rows).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankIntegrationsAPI.GetBankIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBankIntegrations`: GetBankIntegrations200Response
	fmt.Fprintf(os.Stdout, "Response from `BankIntegrationsAPI.GetBankIntegrations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBankIntegrationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **include** | **string** | Include child relations of the BankIntegration object. Format is comma separated. | 
 **index** | **string** | Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;  | 
 **rows** | **float32** | The number of bank integrations to return | 

### Return type

[**GetBankIntegrations200Response**](GetBankIntegrations200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankIntegrationsCreate

> BankIntegration GetBankIntegrationsCreate(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Gets a new blank bank_integration object



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
	resp, r, err := apiClient.BankIntegrationsAPI.GetBankIntegrationsCreate(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankIntegrationsAPI.GetBankIntegrationsCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBankIntegrationsCreate`: BankIntegration
	fmt.Fprintf(os.Stdout, "Response from `BankIntegrationsAPI.GetBankIntegrationsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBankIntegrationsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**BankIntegration**](BankIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRefreshAccounts

> BankIntegration GetRefreshAccounts(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Gets the list of accounts from the remote server



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
	resp, r, err := apiClient.BankIntegrationsAPI.GetRefreshAccounts(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankIntegrationsAPI.GetRefreshAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRefreshAccounts`: BankIntegration
	fmt.Fprintf(os.Stdout, "Response from `BankIntegrationsAPI.GetRefreshAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRefreshAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**BankIntegration**](BankIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoveAccount

> BankIntegration GetRemoveAccount(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Removes an account from the integration



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
	resp, r, err := apiClient.BankIntegrationsAPI.GetRemoveAccount(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankIntegrationsAPI.GetRemoveAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemoveAccount`: BankIntegration
	fmt.Fprintf(os.Stdout, "Response from `BankIntegrationsAPI.GetRemoveAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoveAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**BankIntegration**](BankIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowBankIntegration

> BankIntegration ShowBankIntegration(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Shows a bank_integration



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
	id := "D2J234DFA" // string | The BankIntegration Hashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankIntegrationsAPI.ShowBankIntegration(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankIntegrationsAPI.ShowBankIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowBankIntegration`: BankIntegration
	fmt.Fprintf(os.Stdout, "Response from `BankIntegrationsAPI.ShowBankIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The BankIntegration Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowBankIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**BankIntegration**](BankIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreBankIntegration

> BankIntegration StoreBankIntegration(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Adds a bank_integration



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
	resp, r, err := apiClient.BankIntegrationsAPI.StoreBankIntegration(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankIntegrationsAPI.StoreBankIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoreBankIntegration`: BankIntegration
	fmt.Fprintf(os.Stdout, "Response from `BankIntegrationsAPI.StoreBankIntegration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreBankIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**BankIntegration**](BankIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBankIntegration

> BankIntegration UpdateBankIntegration(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Updates a bank_integration



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
	id := "D2J234DFA" // string | The BankIntegration Hashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankIntegrationsAPI.UpdateBankIntegration(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankIntegrationsAPI.UpdateBankIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBankIntegration`: BankIntegration
	fmt.Fprintf(os.Stdout, "Response from `BankIntegrationsAPI.UpdateBankIntegration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The BankIntegration Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBankIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**BankIntegration**](BankIntegration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

