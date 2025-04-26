# \TaskSchedulersAPI

All URIs are relative to *https://demo.invoiceninja.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkTaskSchedulerActions**](TaskSchedulersAPI.md#BulkTaskSchedulerActions) | **Post** /api/v1/task_schedulers/bulk | Performs bulk actions on an array of task_schedulers
[**CreateTaskScheduler**](TaskSchedulersAPI.md#CreateTaskScheduler) | **Post** /api/v1/task_schedulers/ | Create task scheduler with job 
[**DestroyTaskScheduler**](TaskSchedulersAPI.md#DestroyTaskScheduler) | **Delete** /api/v1/task_schedulers/{id} | Destroy Task Scheduler
[**GetTaskScheduler**](TaskSchedulersAPI.md#GetTaskScheduler) | **Get** /api/v1/task_schedulers/create | Gets a new blank scheduler object
[**GetTaskSchedulers**](TaskSchedulersAPI.md#GetTaskSchedulers) | **Get** /api/v1/task_schedulers/ | Task Scheduler Index
[**ShowTaskScheduler**](TaskSchedulersAPI.md#ShowTaskScheduler) | **Get** /api/v1/task_schedulers/{id} | Show given scheduler
[**UpdateTaskScheduler**](TaskSchedulersAPI.md#UpdateTaskScheduler) | **Put** /api/v1/task_schedulers/{id} | Update task scheduler 



## BulkTaskSchedulerActions

> TaskSchedulerSchema BulkTaskSchedulerActions(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).RequestBody(requestBody).Index(index).Execute()

Performs bulk actions on an array of task_schedulers



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
	requestBody := []int32{int32(123)} // []int32 | array of ids
	index := "user" // string | Replaces the default response index from data to a user specific string  ie.  ```html   ?index=new_index ```  response is wrapped  ```json   {     'new_index' : [       .....       ]   } ```  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaskSchedulersAPI.BulkTaskSchedulerActions(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).RequestBody(requestBody).Index(index).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskSchedulersAPI.BulkTaskSchedulerActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkTaskSchedulerActions`: TaskSchedulerSchema
	fmt.Fprintf(os.Stdout, "Response from `TaskSchedulersAPI.BulkTaskSchedulerActions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkTaskSchedulerActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **requestBody** | **[]int32** | array of ids | 
 **index** | **string** | Replaces the default response index from data to a user specific string  ie.  &#x60;&#x60;&#x60;html   ?index&#x3D;new_index &#x60;&#x60;&#x60;  response is wrapped  &#x60;&#x60;&#x60;json   {     &#39;new_index&#39; : [       .....       ]   } &#x60;&#x60;&#x60;  | 

### Return type

[**TaskSchedulerSchema**](TaskSchedulerSchema.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTaskScheduler

> CreateTaskScheduler(ctx).XRequestedWith(xRequestedWith).TaskSchedulerSchema(taskSchedulerSchema).XAPISECRET(xAPISECRET).Execute()

Create task scheduler with job 



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
	taskSchedulerSchema := *openapiclient.NewTaskSchedulerSchema() // TaskSchedulerSchema | 
	xAPISECRET := "password" // string | The API secret as defined by the .env variable API_SECRET. Only needed for self hosted users, and only applicable on the login route. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaskSchedulersAPI.CreateTaskScheduler(context.Background()).XRequestedWith(xRequestedWith).TaskSchedulerSchema(taskSchedulerSchema).XAPISECRET(xAPISECRET).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskSchedulersAPI.CreateTaskScheduler``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaskSchedulerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **taskSchedulerSchema** | [**TaskSchedulerSchema**](TaskSchedulerSchema.md) |  | 
 **xAPISECRET** | **string** | The API secret as defined by the .env variable API_SECRET. Only needed for self hosted users, and only applicable on the login route. | 

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


## DestroyTaskScheduler

> DestroyTaskScheduler(ctx, id).XRequestedWith(xRequestedWith).Execute()

Destroy Task Scheduler



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
	id := "D2J234DFA" // string | The Scheduler Hashed ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaskSchedulersAPI.DestroyTaskScheduler(context.Background(), id).XRequestedWith(xRequestedWith).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskSchedulersAPI.DestroyTaskScheduler``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Scheduler Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyTaskSchedulerRequest struct via the builder pattern


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


## GetTaskScheduler

> TaskSchedulerSchema GetTaskScheduler(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()

Gets a new blank scheduler object



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
	resp, r, err := apiClient.TaskSchedulersAPI.GetTaskScheduler(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskSchedulersAPI.GetTaskScheduler``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaskScheduler`: TaskSchedulerSchema
	fmt.Fprintf(os.Stdout, "Response from `TaskSchedulersAPI.GetTaskScheduler`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskSchedulerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **include** | **string** | Includes child relationships in the response, format is comma separated. Check each model for the list of associated includes | 

### Return type

[**TaskSchedulerSchema**](TaskSchedulerSchema.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskSchedulers

> GetTaskSchedulers(ctx).XRequestedWith(xRequestedWith).Execute()

Task Scheduler Index



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
	r, err := apiClient.TaskSchedulersAPI.GetTaskSchedulers(context.Background()).XRequestedWith(xRequestedWith).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskSchedulersAPI.GetTaskSchedulers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskSchedulersRequest struct via the builder pattern


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


## ShowTaskScheduler

> ShowTaskScheduler(ctx, id).XRequestedWith(xRequestedWith).Execute()

Show given scheduler



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
	id := "D2J234DFA" // string | The Scheduler Hashed ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaskSchedulersAPI.ShowTaskScheduler(context.Background(), id).XRequestedWith(xRequestedWith).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskSchedulersAPI.ShowTaskScheduler``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Scheduler Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowTaskSchedulerRequest struct via the builder pattern


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


## UpdateTaskScheduler

> UpdateTaskScheduler(ctx, id).XRequestedWith(xRequestedWith).TaskSchedulerSchema(taskSchedulerSchema).XAPISECRET(xAPISECRET).Execute()

Update task scheduler 



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
	id := "D2J234DFA" // string | The Scheduler Hashed ID
	taskSchedulerSchema := *openapiclient.NewTaskSchedulerSchema() // TaskSchedulerSchema | 
	xAPISECRET := "password" // string | The API secret as defined by the .env variable API_SECRET. Only needed for self hosted users, and only applicable on the login route. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TaskSchedulersAPI.UpdateTaskScheduler(context.Background(), id).XRequestedWith(xRequestedWith).TaskSchedulerSchema(taskSchedulerSchema).XAPISECRET(xAPISECRET).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaskSchedulersAPI.UpdateTaskScheduler``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Scheduler Hashed ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTaskSchedulerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 

 **taskSchedulerSchema** | [**TaskSchedulerSchema**](TaskSchedulerSchema.md) |  | 
 **xAPISECRET** | **string** | The API secret as defined by the .env variable API_SECRET. Only needed for self hosted users, and only applicable on the login route. | 

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

