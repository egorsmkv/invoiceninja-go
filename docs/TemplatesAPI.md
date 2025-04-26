# \TemplatesAPI

All URIs are relative to *https://demo.invoiceninja.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetShowTemplate**](TemplatesAPI.md#GetShowTemplate) | **Post** /api/v1/templates | Returns a entity template with the template variables replaced with the Entities



## GetShowTemplate

> Template GetShowTemplate(ctx).XRequestedWith(xRequestedWith).GetShowTemplateRequest(getShowTemplateRequest).Execute()

Returns a entity template with the template variables replaced with the Entities



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
	getShowTemplateRequest := *openapiclient.NewGetShowTemplateRequest() // GetShowTemplateRequest | The template subject and body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.GetShowTemplate(context.Background()).XRequestedWith(xRequestedWith).GetShowTemplateRequest(getShowTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.GetShowTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShowTemplate`: Template
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.GetShowTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetShowTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **getShowTemplateRequest** | [**GetShowTemplateRequest**](GetShowTemplateRequest.md) | The template subject and body | 

### Return type

[**Template**](Template.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

