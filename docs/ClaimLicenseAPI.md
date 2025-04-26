# \ClaimLicenseAPI

All URIs are relative to *https://demo.invoiceninja.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClaimLicense**](ClaimLicenseAPI.md#GetClaimLicense) | **Get** /api/v1/claim_license | Attempts to claim a white label license



## GetClaimLicense

> GetClaimLicense(ctx).XRequestedWith(xRequestedWith).LicenseKey(licenseKey).ProductId(productId).Execute()

Attempts to claim a white label license



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
	licenseKey := "d87sh-s755s-s7d76-sdsd8" // string | The license hash
	productId := "1" // string | The ID of the product purchased.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ClaimLicenseAPI.GetClaimLicense(context.Background()).XRequestedWith(xRequestedWith).LicenseKey(licenseKey).ProductId(productId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClaimLicenseAPI.GetClaimLicense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClaimLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **licenseKey** | **string** | The license hash | 
 **productId** | **string** | The ID of the product purchased. | 

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

