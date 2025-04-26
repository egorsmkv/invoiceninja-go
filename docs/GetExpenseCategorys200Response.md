# GetExpenseCategorys200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ExpenseCategory**](ExpenseCategory.md) |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 

## Methods

### NewGetExpenseCategorys200Response

`func NewGetExpenseCategorys200Response() *GetExpenseCategorys200Response`

NewGetExpenseCategorys200Response instantiates a new GetExpenseCategorys200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetExpenseCategorys200ResponseWithDefaults

`func NewGetExpenseCategorys200ResponseWithDefaults() *GetExpenseCategorys200Response`

NewGetExpenseCategorys200ResponseWithDefaults instantiates a new GetExpenseCategorys200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetExpenseCategorys200Response) GetData() []ExpenseCategory`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetExpenseCategorys200Response) GetDataOk() (*[]ExpenseCategory, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetExpenseCategorys200Response) SetData(v []ExpenseCategory)`

SetData sets Data field to given value.

### HasData

`func (o *GetExpenseCategorys200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetExpenseCategorys200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetExpenseCategorys200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetExpenseCategorys200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetExpenseCategorys200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


