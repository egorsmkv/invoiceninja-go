# GetCompanyGateways200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CompanyGateway**](CompanyGateway.md) |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 

## Methods

### NewGetCompanyGateways200Response

`func NewGetCompanyGateways200Response() *GetCompanyGateways200Response`

NewGetCompanyGateways200Response instantiates a new GetCompanyGateways200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCompanyGateways200ResponseWithDefaults

`func NewGetCompanyGateways200ResponseWithDefaults() *GetCompanyGateways200Response`

NewGetCompanyGateways200ResponseWithDefaults instantiates a new GetCompanyGateways200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetCompanyGateways200Response) GetData() []CompanyGateway`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCompanyGateways200Response) GetDataOk() (*[]CompanyGateway, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCompanyGateways200Response) SetData(v []CompanyGateway)`

SetData sets Data field to given value.

### HasData

`func (o *GetCompanyGateways200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetCompanyGateways200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetCompanyGateways200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetCompanyGateways200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetCompanyGateways200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


