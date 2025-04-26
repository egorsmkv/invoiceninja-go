# BankTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The bank integration hashed id | [optional] 
**UserId** | Pointer to **string** | The user hashed id | [optional] 
**TransactionId** | Pointer to **int32** | The id of the transaction rule | [optional] 
**Amount** | Pointer to **float32** | The transaction amount | [optional] 
**CurrencyId** | Pointer to **string** | The currency ID of the currency | [optional] 
**AccountType** | Pointer to **string** | The account type | [optional] 
**Description** | Pointer to **string** | The description of the transaction | [optional] 
**CategoryId** | Pointer to **int32** | The category id | [optional] 
**CategoryType** | Pointer to **string** | The category description | [optional] 
**BaseType** | Pointer to **string** | Either CREDIT or DEBIT | [optional] 
**Date** | Pointer to **string** | The date of the transaction | [optional] 
**BankAccountId** | Pointer to **int32** | The ID number of the bank account | [optional] 

## Methods

### NewBankTransaction

`func NewBankTransaction() *BankTransaction`

NewBankTransaction instantiates a new BankTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransactionWithDefaults

`func NewBankTransactionWithDefaults() *BankTransaction`

NewBankTransactionWithDefaults instantiates a new BankTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BankTransaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BankTransaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BankTransaction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BankTransaction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *BankTransaction) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BankTransaction) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BankTransaction) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BankTransaction) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetTransactionId

`func (o *BankTransaction) GetTransactionId() int32`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *BankTransaction) GetTransactionIdOk() (*int32, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *BankTransaction) SetTransactionId(v int32)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *BankTransaction) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetAmount

`func (o *BankTransaction) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BankTransaction) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BankTransaction) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BankTransaction) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrencyId

`func (o *BankTransaction) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *BankTransaction) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *BankTransaction) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *BankTransaction) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetAccountType

`func (o *BankTransaction) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *BankTransaction) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *BankTransaction) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *BankTransaction) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetDescription

`func (o *BankTransaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BankTransaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BankTransaction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BankTransaction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategoryId

`func (o *BankTransaction) GetCategoryId() int32`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *BankTransaction) GetCategoryIdOk() (*int32, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *BankTransaction) SetCategoryId(v int32)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *BankTransaction) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetCategoryType

`func (o *BankTransaction) GetCategoryType() string`

GetCategoryType returns the CategoryType field if non-nil, zero value otherwise.

### GetCategoryTypeOk

`func (o *BankTransaction) GetCategoryTypeOk() (*string, bool)`

GetCategoryTypeOk returns a tuple with the CategoryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryType

`func (o *BankTransaction) SetCategoryType(v string)`

SetCategoryType sets CategoryType field to given value.

### HasCategoryType

`func (o *BankTransaction) HasCategoryType() bool`

HasCategoryType returns a boolean if a field has been set.

### GetBaseType

`func (o *BankTransaction) GetBaseType() string`

GetBaseType returns the BaseType field if non-nil, zero value otherwise.

### GetBaseTypeOk

`func (o *BankTransaction) GetBaseTypeOk() (*string, bool)`

GetBaseTypeOk returns a tuple with the BaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseType

`func (o *BankTransaction) SetBaseType(v string)`

SetBaseType sets BaseType field to given value.

### HasBaseType

`func (o *BankTransaction) HasBaseType() bool`

HasBaseType returns a boolean if a field has been set.

### GetDate

`func (o *BankTransaction) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *BankTransaction) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *BankTransaction) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *BankTransaction) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetBankAccountId

`func (o *BankTransaction) GetBankAccountId() int32`

GetBankAccountId returns the BankAccountId field if non-nil, zero value otherwise.

### GetBankAccountIdOk

`func (o *BankTransaction) GetBankAccountIdOk() (*int32, bool)`

GetBankAccountIdOk returns a tuple with the BankAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccountId

`func (o *BankTransaction) SetBankAccountId(v int32)`

SetBankAccountId sets BankAccountId field to given value.

### HasBankAccountId

`func (o *BankTransaction) HasBankAccountId() bool`

HasBankAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


