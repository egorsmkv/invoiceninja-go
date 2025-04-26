# \ReportsAPI

All URIs are relative to *https://demo.invoiceninja.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClientReport**](ReportsAPI.md#GetClientReport) | **Post** /api/v1/reports/clients | Client reports
[**GetContactReport**](ReportsAPI.md#GetContactReport) | **Post** /api/v1/reports/contacts | Contact reports
[**GetCreditReport**](ReportsAPI.md#GetCreditReport) | **Post** /api/v1/reports/credit | Credit reports
[**GetDocumentReport**](ReportsAPI.md#GetDocumentReport) | **Post** /api/v1/reports/documents | Document reports
[**GetExpenseReport**](ReportsAPI.md#GetExpenseReport) | **Post** /api/v1/reports/expense | Expense reports
[**GetInvoiceItemReport**](ReportsAPI.md#GetInvoiceItemReport) | **Post** /api/v1/reports/invoice_items | Invoice item reports
[**GetInvoiceReport**](ReportsAPI.md#GetInvoiceReport) | **Post** /api/v1/reports/invoices | Invoice reports
[**GetPaymentReport**](ReportsAPI.md#GetPaymentReport) | **Post** /api/v1/reports/payments | Payment reports
[**GetProductReport**](ReportsAPI.md#GetProductReport) | **Post** /api/v1/reports/products | Product reports
[**GetProductSalesReport**](ReportsAPI.md#GetProductSalesReport) | **Post** /api/v1/reports/product_sales | Product Salesreports
[**GetProfitLossReport**](ReportsAPI.md#GetProfitLossReport) | **Post** /api/v1/reports/profitloss | Profit loss reports
[**GetQuoteItemReport**](ReportsAPI.md#GetQuoteItemReport) | **Post** /api/v1/reports/quote_items | Quote item reports
[**GetQuoteReport**](ReportsAPI.md#GetQuoteReport) | **Post** /api/v1/reports/quotes | Quote reports
[**GetRecurringInvoiceReport**](ReportsAPI.md#GetRecurringInvoiceReport) | **Post** /api/v1/reports/recurring_invoices | Recurring Invoice reports
[**GetTaskReport**](ReportsAPI.md#GetTaskReport) | **Post** /api/v1/reports/tasks | Task reports



## GetClientReport

> GetClientReport(ctx).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()

Client reports



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
	genericReportSchema := *openapiclient.NewGenericReportSchema() // GenericReportSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.GetClientReport(context.Background()).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetClientReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClientReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **genericReportSchema** | [**GenericReportSchema**](GenericReportSchema.md) |  | 

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


## GetContactReport

> GetContactReport(ctx).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()

Contact reports



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
	genericReportSchema := *openapiclient.NewGenericReportSchema() // GenericReportSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.GetContactReport(context.Background()).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetContactReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContactReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **genericReportSchema** | [**GenericReportSchema**](GenericReportSchema.md) |  | 

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


## GetCreditReport

> GetCreditReport(ctx).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()

Credit reports



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
	genericReportSchema := *openapiclient.NewGenericReportSchema() // GenericReportSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.GetCreditReport(context.Background()).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetCreditReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCreditReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **genericReportSchema** | [**GenericReportSchema**](GenericReportSchema.md) |  | 

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


## GetDocumentReport

> GetDocumentReport(ctx).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()

Document reports



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
	genericReportSchema := *openapiclient.NewGenericReportSchema() // GenericReportSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.GetDocumentReport(context.Background()).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetDocumentReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **genericReportSchema** | [**GenericReportSchema**](GenericReportSchema.md) |  | 

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


## GetExpenseReport

> GetExpenseReport(ctx).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()

Expense reports



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
	genericReportSchema := *openapiclient.NewGenericReportSchema() // GenericReportSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.GetExpenseReport(context.Background()).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetExpenseReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExpenseReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **genericReportSchema** | [**GenericReportSchema**](GenericReportSchema.md) |  | 

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


## GetInvoiceItemReport

> GetInvoiceItemReport(ctx).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()

Invoice item reports



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
	genericReportSchema := *openapiclient.NewGenericReportSchema() // GenericReportSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.GetInvoiceItemReport(context.Background()).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetInvoiceItemReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceItemReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **genericReportSchema** | [**GenericReportSchema**](GenericReportSchema.md) |  | 

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


## GetInvoiceReport

> GetInvoiceReport(ctx).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()

Invoice reports



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
	genericReportSchema := *openapiclient.NewGenericReportSchema() // GenericReportSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.GetInvoiceReport(context.Background()).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetInvoiceReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **genericReportSchema** | [**GenericReportSchema**](GenericReportSchema.md) |  | 

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


## GetPaymentReport

> GetPaymentReport(ctx).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()

Payment reports



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
	genericReportSchema := *openapiclient.NewGenericReportSchema() // GenericReportSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.GetPaymentReport(context.Background()).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetPaymentReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **genericReportSchema** | [**GenericReportSchema**](GenericReportSchema.md) |  | 

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


## GetProductReport

> GetProductReport(ctx).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()

Product reports



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
	genericReportSchema := *openapiclient.NewGenericReportSchema() // GenericReportSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.GetProductReport(context.Background()).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetProductReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProductReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **genericReportSchema** | [**GenericReportSchema**](GenericReportSchema.md) |  | 

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


## GetProductSalesReport

> GetProductSalesReport(ctx).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()

Product Salesreports



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
	genericReportSchema := *openapiclient.NewGenericReportSchema() // GenericReportSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.GetProductSalesReport(context.Background()).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetProductSalesReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProductSalesReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **genericReportSchema** | [**GenericReportSchema**](GenericReportSchema.md) |  | 

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


## GetProfitLossReport

> GetProfitLossReport(ctx).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()

Profit loss reports



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
	genericReportSchema := *openapiclient.NewGenericReportSchema() // GenericReportSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.GetProfitLossReport(context.Background()).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetProfitLossReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProfitLossReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **genericReportSchema** | [**GenericReportSchema**](GenericReportSchema.md) |  | 

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


## GetQuoteItemReport

> GetQuoteItemReport(ctx).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()

Quote item reports



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
	genericReportSchema := *openapiclient.NewGenericReportSchema() // GenericReportSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.GetQuoteItemReport(context.Background()).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetQuoteItemReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQuoteItemReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **genericReportSchema** | [**GenericReportSchema**](GenericReportSchema.md) |  | 

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


## GetQuoteReport

> GetQuoteReport(ctx).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()

Quote reports



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
	genericReportSchema := *openapiclient.NewGenericReportSchema() // GenericReportSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.GetQuoteReport(context.Background()).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetQuoteReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQuoteReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **genericReportSchema** | [**GenericReportSchema**](GenericReportSchema.md) |  | 

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


## GetRecurringInvoiceReport

> GetRecurringInvoiceReport(ctx).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()

Recurring Invoice reports



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
	genericReportSchema := *openapiclient.NewGenericReportSchema() // GenericReportSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.GetRecurringInvoiceReport(context.Background()).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetRecurringInvoiceReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecurringInvoiceReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **genericReportSchema** | [**GenericReportSchema**](GenericReportSchema.md) |  | 

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


## GetTaskReport

> GetTaskReport(ctx).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()

Task reports



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
	genericReportSchema := *openapiclient.NewGenericReportSchema() // GenericReportSchema | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ReportsAPI.GetTaskReport(context.Background()).XRequestedWith(xRequestedWith).GenericReportSchema(genericReportSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetTaskReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **genericReportSchema** | [**GenericReportSchema**](GenericReportSchema.md) |  | 

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

