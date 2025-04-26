# \EmailsAPI

All URIs are relative to *https://demo.invoiceninja.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendEmailTemplate**](EmailsAPI.md#SendEmailTemplate) | **Post** /api/v1/emails | Sends an email for an entity



## SendEmailTemplate

> Template SendEmailTemplate(ctx).XRequestedWith(xRequestedWith).SendEmailTemplateRequest(sendEmailTemplateRequest).Execute()

Sends an email for an entity



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
	sendEmailTemplateRequest := *openapiclient.NewSendEmailTemplateRequest("Entity_example", "EntityId_example", "invoice") // SendEmailTemplateRequest | Required fields to send an email

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailsAPI.SendEmailTemplate(context.Background()).XRequestedWith(xRequestedWith).SendEmailTemplateRequest(sendEmailTemplateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailsAPI.SendEmailTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendEmailTemplate`: Template
	fmt.Fprintf(os.Stdout, "Response from `EmailsAPI.SendEmailTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendEmailTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **sendEmailTemplateRequest** | [**SendEmailTemplateRequest**](SendEmailTemplateRequest.md) | Required fields to send an email | 

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

