# TaxRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Thie hashed id of the tax | [optional] [readonly] 
**Name** | Pointer to **string** | The tax name | [optional] 
**Rate** | Pointer to **float32** | The tax rate | [optional] 
**IsDeleted** | Pointer to **bool** | Boolean flag determining if the tax has been deleted | [optional] 

## Methods

### NewTaxRate

`func NewTaxRate() *TaxRate`

NewTaxRate instantiates a new TaxRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxRateWithDefaults

`func NewTaxRateWithDefaults() *TaxRate`

NewTaxRateWithDefaults instantiates a new TaxRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaxRate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxRate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxRate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaxRate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TaxRate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TaxRate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TaxRate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TaxRate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRate

`func (o *TaxRate) GetRate() float32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *TaxRate) GetRateOk() (*float32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *TaxRate) SetRate(v float32)`

SetRate sets Rate field to given value.

### HasRate

`func (o *TaxRate) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetIsDeleted

`func (o *TaxRate) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *TaxRate) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *TaxRate) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *TaxRate) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


