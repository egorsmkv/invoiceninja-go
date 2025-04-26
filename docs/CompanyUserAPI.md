# \CompanyUserAPI

All URIs are relative to *https://demo.invoiceninja.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateCompanyUser**](CompanyUserAPI.md#UpdateCompanyUser) | **Post** /api/v1/company_users | Update a company user record



## UpdateCompanyUser

> CompanyUser UpdateCompanyUser(ctx).XRequestedWith(xRequestedWith).Execute()

Update a company user record



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
	resp, r, err := apiClient.CompanyUserAPI.UpdateCompanyUser(context.Background()).XRequestedWith(xRequestedWith).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompanyUserAPI.UpdateCompanyUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCompanyUser`: CompanyUser
	fmt.Fprintf(os.Stdout, "Response from `CompanyUserAPI.UpdateCompanyUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCompanyUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

### Return type

[**CompanyUser**](CompanyUser.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

