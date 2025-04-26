# BankTransactionRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The bank transaction rules hashed id | [optional] 
**UserId** | Pointer to **string** | The user hashed id | [optional] 
**Name** | Pointer to **string** | The name of the transaction | [optional] 
**Rules** | Pointer to [**[]BTRules**](BTRules.md) | A mapped collection of the sub rules for the BankTransactionRule | [optional] 
**AutoConvert** | Pointer to **bool** | Flags whether the rule converts the transaction automatically | [optional] 
**MatchesOnAll** | Pointer to **bool** | Flags whether all subrules are required for the match | [optional] 
**AppliesTo** | Pointer to **string** | Flags whether the rule applies to a CREDIT or DEBIT | [optional] 
**ClientId** | Pointer to **string** | The client hashed id | [optional] 
**VendorId** | Pointer to **string** | The vendor hashed id | [optional] 
**CategoryId** | Pointer to **string** | The category hashed id | [optional] 

## Methods

### NewBankTransactionRule

`func NewBankTransactionRule() *BankTransactionRule`

NewBankTransactionRule instantiates a new BankTransactionRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankTransactionRuleWithDefaults

`func NewBankTransactionRuleWithDefaults() *BankTransactionRule`

NewBankTransactionRuleWithDefaults instantiates a new BankTransactionRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BankTransactionRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BankTransactionRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BankTransactionRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BankTransactionRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *BankTransactionRule) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BankTransactionRule) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BankTransactionRule) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *BankTransactionRule) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetName

`func (o *BankTransactionRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BankTransactionRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BankTransactionRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BankTransactionRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRules

`func (o *BankTransactionRule) GetRules() []BTRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *BankTransactionRule) GetRulesOk() (*[]BTRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *BankTransactionRule) SetRules(v []BTRules)`

SetRules sets Rules field to given value.

### HasRules

`func (o *BankTransactionRule) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetAutoConvert

`func (o *BankTransactionRule) GetAutoConvert() bool`

GetAutoConvert returns the AutoConvert field if non-nil, zero value otherwise.

### GetAutoConvertOk

`func (o *BankTransactionRule) GetAutoConvertOk() (*bool, bool)`

GetAutoConvertOk returns a tuple with the AutoConvert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoConvert

`func (o *BankTransactionRule) SetAutoConvert(v bool)`

SetAutoConvert sets AutoConvert field to given value.

### HasAutoConvert

`func (o *BankTransactionRule) HasAutoConvert() bool`

HasAutoConvert returns a boolean if a field has been set.

### GetMatchesOnAll

`func (o *BankTransactionRule) GetMatchesOnAll() bool`

GetMatchesOnAll returns the MatchesOnAll field if non-nil, zero value otherwise.

### GetMatchesOnAllOk

`func (o *BankTransactionRule) GetMatchesOnAllOk() (*bool, bool)`

GetMatchesOnAllOk returns a tuple with the MatchesOnAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchesOnAll

`func (o *BankTransactionRule) SetMatchesOnAll(v bool)`

SetMatchesOnAll sets MatchesOnAll field to given value.

### HasMatchesOnAll

`func (o *BankTransactionRule) HasMatchesOnAll() bool`

HasMatchesOnAll returns a boolean if a field has been set.

### GetAppliesTo

`func (o *BankTransactionRule) GetAppliesTo() string`

GetAppliesTo returns the AppliesTo field if non-nil, zero value otherwise.

### GetAppliesToOk

`func (o *BankTransactionRule) GetAppliesToOk() (*string, bool)`

GetAppliesToOk returns a tuple with the AppliesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesTo

`func (o *BankTransactionRule) SetAppliesTo(v string)`

SetAppliesTo sets AppliesTo field to given value.

### HasAppliesTo

`func (o *BankTransactionRule) HasAppliesTo() bool`

HasAppliesTo returns a boolean if a field has been set.

### GetClientId

`func (o *BankTransactionRule) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *BankTransactionRule) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *BankTransactionRule) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *BankTransactionRule) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetVendorId

`func (o *BankTransactionRule) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *BankTransactionRule) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *BankTransactionRule) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *BankTransactionRule) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetCategoryId

`func (o *BankTransactionRule) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *BankTransactionRule) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *BankTransactionRule) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *BankTransactionRule) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


