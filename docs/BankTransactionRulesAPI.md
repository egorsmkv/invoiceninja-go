# \BankTransactionRulesAPI

All URIs are relative to *https://demo.invoiceninja.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkBankTransactionRules**](BankTransactionRulesAPI.md#BulkBankTransactionRules) | **Post** /api/v1/bank_transation_rules/bulk | Performs bulk actions on an array of bank_transation rules
[**DeleteBankTransactionRule**](BankTransactionRulesAPI.md#DeleteBankTransactionRule) | **Delete** /api/v1/bank_transaction_rules/{id} | Deletes a bank_transaction rule
[**EditBankTransactionRule**](BankTransactionRulesAPI.md#EditBankTransactionRule) | **Get** /api/v1/bank_transaction_rules/{id}/edit | Shows a bank_transaction for editing
[**GetBankTransactionRules**](BankTransactionRulesAPI.md#GetBankTransactionRules) | **Get** /api/v1/bank_transaction_rules | Gets a list of bank_transaction_rules
[**GetBankTransactionRulesCreate**](BankTransactionRulesAPI.md#GetBankTransactionRulesCreate) | **Get** /api/v1/bank_transaction_rules/create | Gets a new blank bank_transaction rule object
[**ShowBankTransactionRule**](BankTransactionRulesAPI.md#ShowBankTransactionRule) | **Get** /api/v1/bank_transaction_rules/{id} | Shows a bank_transaction
[**StoreBankTransactionRule**](BankTransactionRulesAPI.md#StoreBankTransactionRule) | **Post** /api/v1/bank_transaction_rules | Adds a bank_transaction rule
[**UpdateBankTransactionRule**](BankTransactionRulesAPI.md#UpdateBankTransactionRule) | **Put** /api/v1/bank_transaction_rules/{id} | Updates a bank_transaction Rule



## BulkBankTransactionRules

> BulkBankTransactionRules(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).RequestBody(requestBody).Index(index).Execute()

Performs bulk actions on an array of bank_transation rules



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
	r, err := apiClient.BankTransactionRulesAPI.BulkBankTransactionRules(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).RequestBody(requestBody).Index(index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankTransactionRulesAPI.BulkBankTransactionRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkBankTransactionRulesRequest struct via the builder pattern


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


## DeleteBankTransactionRule

> DeleteBankTransactionRule(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Deletes a bank_transaction rule



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
	id := "D2J234DFA" // string | The Bank Transaction Rule Hashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BankTransactionRulesAPI.DeleteBankTransactionRule(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankTransactionRulesAPI.DeleteBankTransactionRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Bank Transaction Rule Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBankTransactionRuleRequest struct via the builder pattern


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


## EditBankTransactionRule

> BankTransactionRule EditBankTransactionRule(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Shows a bank_transaction for editing



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
	id := "D2J234DFA" // string | The Bank Transaction Rule Hashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankTransactionRulesAPI.EditBankTransactionRule(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankTransactionRulesAPI.EditBankTransactionRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditBankTransactionRule`: BankTransactionRule
	fmt.Fprintf(os.Stdout, "Response from `BankTransactionRulesAPI.EditBankTransactionRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Bank Transaction Rule Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditBankTransactionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**BankTransactionRule**](BankTransactionRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankTransactionRules

> GetBankTransactionRules200Response GetBankTransactionRules(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Index(index).Rows(rows).Execute()

Gets a list of bank_transaction_rules



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
	rows := float32(50) // float32 | The number of bank integrations to return (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankTransactionRulesAPI.GetBankTransactionRules(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Index(index).Rows(rows).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankTransactionRulesAPI.GetBankTransactionRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBankTransactionRules`: GetBankTransactionRules200Response
	fmt.Fprintf(os.Stdout, "Response from `BankTransactionRulesAPI.GetBankTransactionRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBankTransactionRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 
 **index** | **string** | Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;  | 
 **rows** | **float32** | The number of bank integrations to return | 

### Return type

[**GetBankTransactionRules200Response**](GetBankTransactionRules200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBankTransactionRulesCreate

> BankTransactionRule GetBankTransactionRulesCreate(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Gets a new blank bank_transaction rule object



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
	resp, r, err := apiClient.BankTransactionRulesAPI.GetBankTransactionRulesCreate(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankTransactionRulesAPI.GetBankTransactionRulesCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBankTransactionRulesCreate`: BankTransactionRule
	fmt.Fprintf(os.Stdout, "Response from `BankTransactionRulesAPI.GetBankTransactionRulesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBankTransactionRulesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**BankTransactionRule**](BankTransactionRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowBankTransactionRule

> BankTransactionRule ShowBankTransactionRule(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Shows a bank_transaction



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
	id := "D2J234DFA" // string | The Bank Transaction RuleHashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankTransactionRulesAPI.ShowBankTransactionRule(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankTransactionRulesAPI.ShowBankTransactionRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShowBankTransactionRule`: BankTransactionRule
	fmt.Fprintf(os.Stdout, "Response from `BankTransactionRulesAPI.ShowBankTransactionRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Bank Transaction RuleHashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowBankTransactionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**BankTransactionRule**](BankTransactionRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreBankTransactionRule

> BankTransactionRule StoreBankTransactionRule(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Adds a bank_transaction rule



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
	resp, r, err := apiClient.BankTransactionRulesAPI.StoreBankTransactionRule(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankTransactionRulesAPI.StoreBankTransactionRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoreBankTransactionRule`: BankTransactionRule
	fmt.Fprintf(os.Stdout, "Response from `BankTransactionRulesAPI.StoreBankTransactionRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreBankTransactionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**BankTransactionRule**](BankTransactionRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBankTransactionRule

> BankTransactionRule UpdateBankTransactionRule(ctx, id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Updates a bank_transaction Rule



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
	id := "D2J234DFA" // string | The Bank Transaction Rule Hashed ID
	include := "first_load" // string | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BankTransactionRulesAPI.UpdateBankTransactionRule(context.Background(), id).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BankTransactionRulesAPI.UpdateBankTransactionRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBankTransactionRule`: BankTransactionRule
	fmt.Fprintf(os.Stdout, "Response from `BankTransactionRulesAPI.UpdateBankTransactionRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Bank Transaction Rule Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBankTransactionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**BankTransactionRule**](BankTransactionRule.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

