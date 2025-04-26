# AuthorizationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Insufficient permissions for this resource. | [optional] 
**Errors** | Pointer to [**ValidationErrorErrors**](ValidationErrorErrors.md) |  | [optional] 

## Methods

### NewAuthorizationError

`func NewAuthorizationError() *AuthorizationError`

NewAuthorizationError instantiates a new AuthorizationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizationErrorWithDefaults

`func NewAuthorizationErrorWithDefaults() *AuthorizationError`

NewAuthorizationErrorWithDefaults instantiates a new AuthorizationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AuthorizationError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AuthorizationError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AuthorizationError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AuthorizationError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrors

`func (o *AuthorizationError) GetErrors() ValidationErrorErrors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AuthorizationError) GetErrorsOk() (*ValidationErrorErrors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *AuthorizationError) SetErrors(v ValidationErrorErrors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *AuthorizationError) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


