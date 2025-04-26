# \ImportAPI

All URIs are relative to *https://demo.invoiceninja.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetImportJson**](ImportAPI.md#GetImportJson) | **Post** /api/v1/import_json | Import data from the system



## GetImportJson

> GetImportJson(ctx).XRequestedWith(xRequestedWith).Execute()

Import data from the system



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
	xRequestedWith := "XMLHttpRequest" // string | Used to send the XMLHttpRequest header

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImportAPI.GetImportJson(context.Background()).XRequestedWith(xRequestedWith).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImportAPI.GetImportJson``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetImportJsonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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

