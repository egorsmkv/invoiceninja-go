# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The account hashed id | [optional] 
**AccountSmsVerified** | Pointer to **string** | Boolean flag if the account has been verified by sms | [optional] 

## Methods

### NewAccount

`func NewAccount() *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Account) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Account) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Account) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Account) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountSmsVerified

`func (o *Account) GetAccountSmsVerified() string`

GetAccountSmsVerified returns the AccountSmsVerified field if non-nil, zero value otherwise.

### GetAccountSmsVerifiedOk

`func (o *Account) GetAccountSmsVerifiedOk() (*string, bool)`

GetAccountSmsVerifiedOk returns a tuple with the AccountSmsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSmsVerified

`func (o *Account) SetAccountSmsVerified(v string)`

SetAccountSmsVerified sets AccountSmsVerified field to given value.

### HasAccountSmsVerified

`func (o *Account) HasAccountSmsVerified() bool`

HasAccountSmsVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


