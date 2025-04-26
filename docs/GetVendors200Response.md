# GetVendors200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Vendor**](Vendor.md) |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 

## Methods

### NewGetVendors200Response

`func NewGetVendors200Response() *GetVendors200Response`

NewGetVendors200Response instantiates a new GetVendors200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVendors200ResponseWithDefaults

`func NewGetVendors200ResponseWithDefaults() *GetVendors200Response`

NewGetVendors200ResponseWithDefaults instantiates a new GetVendors200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetVendors200Response) GetData() []Vendor`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetVendors200Response) GetDataOk() (*[]Vendor, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetVendors200Response) SetData(v []Vendor)`

SetData sets Data field to given value.

### HasData

`func (o *GetVendors200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetVendors200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetVendors200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetVendors200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetVendors200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


