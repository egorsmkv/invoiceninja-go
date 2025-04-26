# \AuthAPI

All URIs are relative to *https://demo.invoiceninja.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostLogin**](AuthAPI.md#PostLogin) | **Post** /api/v1/login | Login



## PostLogin

> CompanyUser PostLogin(ctx).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).PostLoginRequest(postLoginRequest).XAPISECRET(xAPISECRET).Include(include).IncludeStatic(includeStatic).ClearCache(clearCache).Execute()

Login



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
	postLoginRequest := *openapiclient.NewPostLoginRequest("demo@invoiceninja.com", "Password0") // PostLoginRequest | User credentials  ```json {   \"email\" : \"fred@flintstonze.com\",   \"password\" : \"magicpassword123\",   \"one_time_password\" : \"12345\", } ``` 
	xAPISECRET := "password" // string | The API secret as defined by the .env variable API_SECRET. Only needed for self hosted users, and only applicable on the login route. (optional)
	include := "user" // string | Include child relations of the CompanyUser object, format is comma separated.      <br />  > ### **Note**: it is possible to chain multiple includes together, ie. include=account,token  <br />  ```html  Available includes:    user   company   token   account ```  (optional)
	includeStatic := "include_static=true" // string | This include will return the full set of static variables used in the application including:   - Currencies   - Countries    - Languages   - Payment Types   - Email Templatees   - Industries  (optional)
	clearCache := "clear_cache=true" // string | Clears cache  Clears (and rebuilds) the static variable cache.    (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthAPI.PostLogin(context.Background()).XAPITOKEN(xAPITOKEN).XRequestedWith(xRequestedWith).PostLoginRequest(postLoginRequest).XAPISECRET(xAPISECRET).Include(include).IncludeStatic(includeStatic).ClearCache(clearCache).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthAPI.PostLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLogin`: CompanyUser
	fmt.Fprintf(os.Stdout, "Response from `AuthAPI.PostLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPITOKEN** | **string** | The API token to be used for authentication | 
 **xRequestedWith** | **string** | Used to send the XMLHttpRequest header | 
 **postLoginRequest** | [**PostLoginRequest**](PostLoginRequest.md) | User credentials  &#x60;&#x60;&#x60;json {   \&quot;email\&quot; : \&quot;fred@flintstonze.com\&quot;,   \&quot;password\&quot; : \&quot;magicpassword123\&quot;,   \&quot;one_time_password\&quot; : \&quot;12345\&quot;, } &#x60;&#x60;&#x60;  | 
 **xAPISECRET** | **string** | The API secret as defined by the .env variable API_SECRET. Only needed for self hosted users, and only applicable on the login route. | 
 **include** | **string** | Include child relations of the CompanyUser object, format is comma separated.      &lt;br /&gt;  &gt; ### **Note**: it is possible to chain multiple includes together, ie. include&#x3D;account,token  &lt;br /&gt;  &#x60;&#x60;&#x60;html  Available includes:    user   company   token   account &#x60;&#x60;&#x60;  | 
 **includeStatic** | **string** | This include will return the full set of static variables used in the application including:   - Currencies   - Countries    - Languages   - Payment Types   - Email Templatees   - Industries  | 
 **clearCache** | **string** | Clears cache  Clears (and rebuilds) the static variable cache.    | 

### Return type

[**CompanyUser**](CompanyUser.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

