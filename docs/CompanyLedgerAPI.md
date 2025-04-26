# \CompanyLedgerAPI

All URIs are relative to *https://demo.invoiceninja.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompanyLedger**](CompanyLedgerAPI.md#GetCompanyLedger) | **Get** /api/v1/company_ledger | Gets a list of company_ledger



## GetCompanyLedger

> GetCompanyLedger200Response GetCompanyLedger(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Gets a list of company_ledger



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
	resp, r, err := apiClient.CompanyLedgerAPI.GetCompanyLedger(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyLedgerAPI.GetCompanyLedger``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCompanyLedger`: GetCompanyLedger200Response
	fmt.Fprintf(os.Stdout, "Response from `CompanyLedgerAPI.GetCompanyLedger`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyLedgerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**GetCompanyLedger200Response**](GetCompanyLedger200Response.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

