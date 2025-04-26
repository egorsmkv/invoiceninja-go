# CompanyToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The token name | [optional] 
**Token** | Pointer to **string** | The token value | [optional] 
**IsSystem** | Pointer to **bool** | Determines whether the token is created by the system rather than a user | [optional] 

## Methods

### NewCompanyToken

`func NewCompanyToken() *CompanyToken`

NewCompanyToken instantiates a new CompanyToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyTokenWithDefaults

`func NewCompanyTokenWithDefaults() *CompanyToken`

NewCompanyTokenWithDefaults instantiates a new CompanyToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CompanyToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyToken) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompanyToken) HasName() bool`

HasName returns a boolean if a field has been set.

### GetToken

`func (o *CompanyToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CompanyToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CompanyToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CompanyToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetIsSystem

`func (o *CompanyToken) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *CompanyToken) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *CompanyToken) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.

### HasIsSystem

`func (o *CompanyToken) HasIsSystem() bool`

HasIsSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


