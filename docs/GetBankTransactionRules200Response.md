# GetBankTransactionRules200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]BankTransactionRule**](BankTransactionRule.md) |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 

## Methods

### NewGetBankTransactionRules200Response

`func NewGetBankTransactionRules200Response() *GetBankTransactionRules200Response`

NewGetBankTransactionRules200Response instantiates a new GetBankTransactionRules200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBankTransactionRules200ResponseWithDefaults

`func NewGetBankTransactionRules200ResponseWithDefaults() *GetBankTransactionRules200Response`

NewGetBankTransactionRules200ResponseWithDefaults instantiates a new GetBankTransactionRules200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetBankTransactionRules200Response) GetData() []BankTransactionRule`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetBankTransactionRules200Response) GetDataOk() (*[]BankTransactionRule, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetBankTransactionRules200Response) SetData(v []BankTransactionRule)`

SetData sets Data field to given value.

### HasData

`func (o *GetBankTransactionRules200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetBankTransactionRules200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetBankTransactionRules200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetBankTransactionRules200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetBankTransactionRules200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


