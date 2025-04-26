# GetExpenses200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Expense**](Expense.md) |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 

## Methods

### NewGetExpenses200Response

`func NewGetExpenses200Response() *GetExpenses200Response`

NewGetExpenses200Response instantiates a new GetExpenses200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetExpenses200ResponseWithDefaults

`func NewGetExpenses200ResponseWithDefaults() *GetExpenses200Response`

NewGetExpenses200ResponseWithDefaults instantiates a new GetExpenses200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetExpenses200Response) GetData() []Expense`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetExpenses200Response) GetDataOk() (*[]Expense, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetExpenses200Response) SetData(v []Expense)`

SetData sets Data field to given value.

### HasData

`func (o *GetExpenses200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetExpenses200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetExpenses200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetExpenses200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetExpenses200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


