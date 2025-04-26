# \ChartsAPI

All URIs are relative to *https://demo.invoiceninja.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChartTotals**](ChartsAPI.md#GetChartTotals) | **Post** /api/v1/charts/totals | Get chart data



## GetChartTotals

> GetChartTotals(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Index(index).Rows(rows).Execute()

Get chart data



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
	rows := float32(50) // float32 | The chart (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ChartsAPI.GetChartTotals(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Index(index).Rows(rows).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ChartsAPI.GetChartTotals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChartTotalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 
 **index** | **string** | Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;  | 
 **rows** | **float32** | The chart | 

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

