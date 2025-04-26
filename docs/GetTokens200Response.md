# GetTokens200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CompanyToken**](CompanyToken.md) |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 

## Methods

### NewGetTokens200Response

`func NewGetTokens200Response() *GetTokens200Response`

NewGetTokens200Response instantiates a new GetTokens200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTokens200ResponseWithDefaults

`func NewGetTokens200ResponseWithDefaults() *GetTokens200Response`

NewGetTokens200ResponseWithDefaults instantiates a new GetTokens200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetTokens200Response) GetData() []CompanyToken`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTokens200Response) GetDataOk() (*[]CompanyToken, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTokens200Response) SetData(v []CompanyToken)`

SetData sets Data field to given value.

### HasData

`func (o *GetTokens200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetTokens200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTokens200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTokens200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetTokens200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


