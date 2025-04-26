# GetGroupSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GroupSetting**](GroupSetting.md) |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 

## Methods

### NewGetGroupSettings200Response

`func NewGetGroupSettings200Response() *GetGroupSettings200Response`

NewGetGroupSettings200Response instantiates a new GetGroupSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGroupSettings200ResponseWithDefaults

`func NewGetGroupSettings200ResponseWithDefaults() *GetGroupSettings200Response`

NewGetGroupSettings200ResponseWithDefaults instantiates a new GetGroupSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetGroupSettings200Response) GetData() []GroupSetting`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetGroupSettings200Response) GetDataOk() (*[]GroupSetting, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetGroupSettings200Response) SetData(v []GroupSetting)`

SetData sets Data field to given value.

### HasData

`func (o *GetGroupSettings200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetGroupSettings200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetGroupSettings200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetGroupSettings200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetGroupSettings200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


