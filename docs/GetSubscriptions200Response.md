# GetSubscriptions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Subscription**](Subscription.md) |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 

## Methods

### NewGetSubscriptions200Response

`func NewGetSubscriptions200Response() *GetSubscriptions200Response`

NewGetSubscriptions200Response instantiates a new GetSubscriptions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSubscriptions200ResponseWithDefaults

`func NewGetSubscriptions200ResponseWithDefaults() *GetSubscriptions200Response`

NewGetSubscriptions200ResponseWithDefaults instantiates a new GetSubscriptions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetSubscriptions200Response) GetData() []Subscription`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetSubscriptions200Response) GetDataOk() (*[]Subscription, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetSubscriptions200Response) SetData(v []Subscription)`

SetData sets Data field to given value.

### HasData

`func (o *GetSubscriptions200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetSubscriptions200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetSubscriptions200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetSubscriptions200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetSubscriptions200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


