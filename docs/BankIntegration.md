# BankIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The bank integration hashed id | [optional] 
**UserId** | Pointer to **string** | The user hashed id | [optional] 
**ProviderBankName** | Pointer to **string** | The providers bank name | [optional] 
**BankAccountId** | Pointer to **int32** | The bank account id | [optional] 
**BankAccountName** | Pointer to **string** | The name of the account | [optional] 
**BankAccountNumber** | Pointer to **string** | The account number | [optional] 
**BankAccountStatus** | Pointer to **string** | The status of the bank account | [optional] 
**BankAccountType** | Pointer to **string** | The type of account | [optional] 
**Balance** | Pointer to **float32** | The current bank balance if available | [optional] 
**Currency** | Pointer to **string** | iso_3166_3 code | [optional] 

## Methods

### NewBankIntegration

`func NewBankIntegration() *BankIntegration`

NewBankIntegration instantiates a new BankIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankIntegrationWithDefaults

`func NewBankIntegrationWithDefaults() *BankIntegration`

NewBankIntegrationWithDefaults instantiates a new BankIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BankIntegration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BankIntegration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BankIntegration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BankIntegration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *BankIntegration) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BankIntegration) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BankIntegration) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BankIntegration) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetProviderBankName

`func (o *BankIntegration) GetProviderBankName() string`

GetProviderBankName returns the ProviderBankName field if non-nil, zero value otherwise.

### GetProviderBankNameOk

`func (o *BankIntegration) GetProviderBankNameOk() (*string, bool)`

GetProviderBankNameOk returns a tuple with the ProviderBankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderBankName

`func (o *BankIntegration) SetProviderBankName(v string)`

SetProviderBankName sets ProviderBankName field to given value.

### HasProviderBankName

`func (o *BankIntegration) HasProviderBankName() bool`

HasProviderBankName returns a boolean if a field has been set.

### GetBankAccountId

`func (o *BankIntegration) GetBankAccountId() int32`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *BankIntegration) GetBankAccountIdOk() (*int32, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *BankIntegration) SetBankAccountId(v int32)`

SetBankAccountId sets BankAccountId field to given value.

### HasBankAccountId

`func (o *BankIntegration) HasBankAccountId() bool`

HasBankAccountId returns a boolean if a field has been set.

### GetBankAccountName

`func (o *BankIntegration) GetBankAccountName() string`

GetBankAccountName returns the BankAccountName field if non-nil, zero value otherwise.

### GetBankAccountNameOk

`func (o *BankIntegration) GetBankAccountNameOk() (*string, bool)`

GetBankAccountNameOk returns a tuple with the BankAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountName

`func (o *BankIntegration) SetBankAccountName(v string)`

SetBankAccountName sets BankAccountName field to given value.

### HasBankAccountName

`func (o *BankIntegration) HasBankAccountName() bool`

HasBankAccountName returns a boolean if a field has been set.

### GetBankAccountNumber

`func (o *BankIntegration) GetBankAccountNumber() string`

GetBankAccountNumber returns the BankAccountNumber field if non-nil, zero value otherwise.

### GetBankAccountNumberOk

`func (o *BankIntegration) GetBankAccountNumberOk() (*string, bool)`

GetBankAccountNumberOk returns a tuple with the BankAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountNumber

`func (o *BankIntegration) SetBankAccountNumber(v string)`

SetBankAccountNumber sets BankAccountNumber field to given value.

### HasBankAccountNumber

`func (o *BankIntegration) HasBankAccountNumber() bool`

HasBankAccountNumber returns a boolean if a field has been set.

### GetBankAccountStatus

`func (o *BankIntegration) GetBankAccountStatus() string`

GetBankAccountStatus returns the BankAccountStatus field if non-nil, zero value otherwise.

### GetBankAccountStatusOk

`func (o *BankIntegration) GetBankAccountStatusOk() (*string, bool)`

GetBankAccountStatusOk returns a tuple with the BankAccountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountStatus

`func (o *BankIntegration) SetBankAccountStatus(v string)`

SetBankAccountStatus sets BankAccountStatus field to given value.

### HasBankAccountStatus

`func (o *BankIntegration) HasBankAccountStatus() bool`

HasBankAccountStatus returns a boolean if a field has been set.

### GetBankAccountType

`func (o *BankIntegration) GetBankAccountType() string`

GetBankAccountType returns the BankAccountType field if non-nil, zero value otherwise.

### GetBankAccountTypeOk

`func (o *BankIntegration) GetBankAccountTypeOk() (*string, bool)`

GetBankAccountTypeOk returns a tuple with the BankAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountType

`func (o *BankIntegration) SetBankAccountType(v string)`

SetBankAccountType sets BankAccountType field to given value.

### HasBankAccountType

`func (o *BankIntegration) HasBankAccountType() bool`

HasBankAccountType returns a boolean if a field has been set.

### GetBalance

`func (o *BankIntegration) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *BankIntegration) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *BankIntegration) SetBalance(v float32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *BankIntegration) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetCurrency

`func (o *BankIntegration) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BankIntegration) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BankIntegration) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BankIntegration) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


