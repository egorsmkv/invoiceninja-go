# ProductBulkAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The action to perform ie. archive / restore / delete / set_tax_id | 
**Ids** | **[]string** |  | 
**TaxId** | Pointer to **string** | The tax rate id to set on the list of products  The following constants are available (default &#x3D; &#39;1&#39;)  &#x60;&#x60;&#x60; PRODUCT_TYPE_PHYSICAL &#x3D; &#39;1&#39; PRODUCT_TYPE_SERVICE &#x3D; &#39;2&#39; PRODUCT_TYPE_DIGITAL &#x3D; &#39;3&#39; PRODUCT_TYPE_SHIPPING &#x3D; &#39;4&#39; PRODUCT_TYPE_EXEMPT &#x3D; &#39;5&#39; PRODUCT_TYPE_REDUCED_TAX &#x3D; &#39;6&#39; PRODUCT_TYPE_OVERRIDE_TAX &#x3D; &#39;7&#39; PRODUCT_TYPE_ZERO_RATED &#x3D; &#39;8&#39; PRODUCT_TYPE_REVERSE_TAX &#x3D; &#39;9&#39; &#x60;&#x60;&#x60;  | [optional] 

## Methods

### NewProductBulkAction

`func NewProductBulkAction(action string, ids []string, ) *ProductBulkAction`

NewProductBulkAction instantiates a new ProductBulkAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductBulkActionWithDefaults

`func NewProductBulkActionWithDefaults() *ProductBulkAction`

NewProductBulkActionWithDefaults instantiates a new ProductBulkAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ProductBulkAction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ProductBulkAction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ProductBulkAction) SetAction(v string)`

SetAction sets Action field to given value.


### GetIds

`func (o *ProductBulkAction) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *ProductBulkAction) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *ProductBulkAction) SetIds(v []string)`

SetIds sets Ids field to given value.


### GetTaxId

`func (o *ProductBulkAction) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *ProductBulkAction) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *ProductBulkAction) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *ProductBulkAction) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


