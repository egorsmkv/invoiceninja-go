# PostLoginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The users email address. | 
**Password** | **string** | The user password. Must meet minimum criteria ~ &gt; 6 characters | 
**OneTimePassword** | Pointer to **string** | The one time password if 2FA is enabled | [optional] 

## Methods

### NewPostLoginRequest

`func NewPostLoginRequest(email string, password string, ) *PostLoginRequest`

NewPostLoginRequest instantiates a new PostLoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostLoginRequestWithDefaults

`func NewPostLoginRequestWithDefaults() *PostLoginRequest`

NewPostLoginRequestWithDefaults instantiates a new PostLoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *PostLoginRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PostLoginRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PostLoginRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *PostLoginRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PostLoginRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PostLoginRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetOneTimePassword

`func (o *PostLoginRequest) GetOneTimePassword() string`

GetOneTimePassword returns the OneTimePassword field if non-nil, zero value otherwise.

### GetOneTimePasswordOk

`func (o *PostLoginRequest) GetOneTimePasswordOk() (*string, bool)`

GetOneTimePasswordOk returns a tuple with the OneTimePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneTimePassword

`func (o *PostLoginRequest) SetOneTimePassword(v string)`

SetOneTimePassword sets OneTimePassword field to given value.

### HasOneTimePassword

`func (o *PostLoginRequest) HasOneTimePassword() bool`

HasOneTimePassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


