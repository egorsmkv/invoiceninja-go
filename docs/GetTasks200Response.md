# GetTasks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Task**](Task.md) |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 

## Methods

### NewGetTasks200Response

`func NewGetTasks200Response() *GetTasks200Response`

NewGetTasks200Response instantiates a new GetTasks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTasks200ResponseWithDefaults

`func NewGetTasks200ResponseWithDefaults() *GetTasks200Response`

NewGetTasks200ResponseWithDefaults instantiates a new GetTasks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetTasks200Response) GetData() []Task`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTasks200Response) GetDataOk() (*[]Task, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTasks200Response) SetData(v []Task)`

SetData sets Data field to given value.

### HasData

`func (o *GetTasks200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetTasks200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTasks200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTasks200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetTasks200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


