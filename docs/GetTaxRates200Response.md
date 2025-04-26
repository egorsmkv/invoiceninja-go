# GetTaxRates200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]TaxRate**](TaxRate.md) |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 

## Methods

### NewGetTaxRates200Response

`func NewGetTaxRates200Response() *GetTaxRates200Response`

NewGetTaxRates200Response instantiates a new GetTaxRates200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTaxRates200ResponseWithDefaults

`func NewGetTaxRates200ResponseWithDefaults() *GetTaxRates200Response`

NewGetTaxRates200ResponseWithDefaults instantiates a new GetTaxRates200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetTaxRates200Response) GetData() []TaxRate`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetTaxRates200Response) GetDataOk() (*[]TaxRate, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetTaxRates200Response) SetData(v []TaxRate)`

SetData sets Data field to given value.

### HasData

`func (o *GetTaxRates200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetTaxRates200Response) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTaxRates200Response) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTaxRates200Response) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetTaxRates200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


