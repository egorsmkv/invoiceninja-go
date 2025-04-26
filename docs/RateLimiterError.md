# RateLimiterError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Rate limit exceeded. | [optional] 
**Errors** | Pointer to [**ValidationErrorErrors**](ValidationErrorErrors.md) |  | [optional] 

## Methods

### NewRateLimiterError

`func NewRateLimiterError() *RateLimiterError`

NewRateLimiterError instantiates a new RateLimiterError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimiterErrorWithDefaults

`func NewRateLimiterErrorWithDefaults() *RateLimiterError`

NewRateLimiterErrorWithDefaults instantiates a new RateLimiterError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *RateLimiterError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RateLimiterError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RateLimiterError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RateLimiterError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrors

`func (o *RateLimiterError) GetErrors() ValidationErrorErrors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *RateLimiterError) GetErrorsOk() (*ValidationErrorErrors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *RateLimiterError) SetErrors(v ValidationErrorErrors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *RateLimiterError) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


