# GetTaskStatuses200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]TaskStatus**](TaskStatus.md) |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 

## Methods

### NewGetTaskStatuses200Response

`func NewGetTaskStatuses200Response() *GetTaskStatuses200Response`

NewGetTaskStatuses200Response instantiates a new GetTaskStatuses200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTaskStatuses200ResponseWithDefaults

`func NewGetTaskStatuses200ResponseWithDefaults() *GetTaskStatuses200Response`

NewGetTaskStatuses200ResponseWithDefaults instantiates a new GetTaskStatuses200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetTaskStatuses200Response) GetData() []TaskStatus`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTaskStatuses200Response) GetDataOk() (*[]TaskStatus, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTaskStatuses200Response) SetData(v []TaskStatus)`

SetData sets Data field to given value.

### HasData

`func (o *GetTaskStatuses200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetTaskStatuses200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTaskStatuses200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTaskStatuses200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetTaskStatuses200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


