# GetQuotes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Quote**](Quote.md) |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 

## Methods

### NewGetQuotes200Response

`func NewGetQuotes200Response() *GetQuotes200Response`

NewGetQuotes200Response instantiates a new GetQuotes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetQuotes200ResponseWithDefaults

`func NewGetQuotes200ResponseWithDefaults() *GetQuotes200Response`

NewGetQuotes200ResponseWithDefaults instantiates a new GetQuotes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetQuotes200Response) GetData() []Quote`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetQuotes200Response) GetDataOk() (*[]Quote, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetQuotes200Response) SetData(v []Quote)`

SetData sets Data field to given value.

### HasData

`func (o *GetQuotes200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetQuotes200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetQuotes200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetQuotes200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetQuotes200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


