# GetDesigns200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Design**](Design.md) |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 

## Methods

### NewGetDesigns200Response

`func NewGetDesigns200Response() *GetDesigns200Response`

NewGetDesigns200Response instantiates a new GetDesigns200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDesigns200ResponseWithDefaults

`func NewGetDesigns200ResponseWithDefaults() *GetDesigns200Response`

NewGetDesigns200ResponseWithDefaults instantiates a new GetDesigns200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetDesigns200Response) GetData() []Design`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetDesigns200Response) GetDataOk() (*[]Design, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetDesigns200Response) SetData(v []Design)`

SetData sets Data field to given value.

### HasData

`func (o *GetDesigns200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetDesigns200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetDesigns200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetDesigns200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetDesigns200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


