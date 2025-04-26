# GetBankTransactions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]BankTransaction**](BankTransaction.md) |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 

## Methods

### NewGetBankTransactions200Response

`func NewGetBankTransactions200Response() *GetBankTransactions200Response`

NewGetBankTransactions200Response instantiates a new GetBankTransactions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBankTransactions200ResponseWithDefaults

`func NewGetBankTransactions200ResponseWithDefaults() *GetBankTransactions200Response`

NewGetBankTransactions200ResponseWithDefaults instantiates a new GetBankTransactions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetBankTransactions200Response) GetData() []BankTransaction`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetBankTransactions200Response) GetDataOk() (*[]BankTransaction, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetBankTransactions200Response) SetData(v []BankTransaction)`

SetData sets Data field to given value.

### HasData

`func (o *GetBankTransactions200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetBankTransactions200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetBankTransactions200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetBankTransactions200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetBankTransactions200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


