# Product

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The hashed product ID. | [optional] [readonly] 
**UserId** | Pointer to **string** | The hashed ID of the user that created this product. | [optional] [readonly] 
**AssignedUserId** | Pointer to **string** | The hashed ID of the user assigned to this product. | [optional] 
**ProjectId** | Pointer to **string** | The hashed ID of the project that this product is associated with. | [optional] 
**VendorId** | Pointer to **string** | The hashed ID of the vendor that this product is associated with. | [optional] 
**CustomValue1** | Pointer to **string** | Custom value field 1. | [optional] 
**CustomValue2** | Pointer to **string** | Custom value field 2. | [optional] 
**CustomValue3** | Pointer to **string** | Custom value field 3. | [optional] 
**CustomValue4** | Pointer to **string** | Custom value field 4. | [optional] 
**ProductKey** | Pointer to **string** | The product key. | [optional] 
**Notes** | Pointer to **string** | Notes about the product. | [optional] 
**Cost** | Pointer to **float64** | The cost of the product. (Your purchase price for this product) | [optional] 
**Price** | Pointer to **float64** | The price of the product that you are charging. | [optional] 
**Quantity** | Pointer to **float64** | The quantity of the product. (used as a default) | [optional] 
**TaxName1** | Pointer to **string** | The name of tax 1. | [optional] 
**TaxRate1** | Pointer to **float64** | The rate of tax 1. | [optional] 
**TaxName2** | Pointer to **string** | The name of tax 2. | [optional] 
**TaxRate2** | Pointer to **float64** | The rate of tax 2. | [optional] 
**TaxName3** | Pointer to **string** | The name of tax 3. | [optional] 
**TaxRate3** | Pointer to **float64** | The rate of tax 3. | [optional] 
**ArchivedAt** | Pointer to **int32** | The timestamp when the product was archived. | [optional] [readonly] 
**CreatedAt** | Pointer to **int32** | The timestamp when the product was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **int32** | Timestamp | [optional] [readonly] 
**IsDeleted** | Pointer to **bool** | Boolean flag determining if the product has been deleted | [optional] [readonly] 
**InStockQuantity** | Pointer to **int32** | The quantity of the product that is currently in stock | [optional] [default to 0]
**StockNotification** | Pointer to **bool** | Indicates whether stock notifications are enabled for this product | [optional] [default to true]
**StockNotificationThreshold** | Pointer to **int32** | The minimum quantity threshold for which stock notifications will be triggered | [optional] [default to 0]
**MaxQuantity** | Pointer to **int32** | The maximum quantity that can be ordered for this product | [optional] 
**ProductImage** | Pointer to **string** | The URL of the product image | [optional] 
**TaxId** | Pointer to **string** | The tax category id for this product.&#39;  The following constants are available (default &#x3D; &#39;1&#39;)  &#x60;&#x60;&#x60; PRODUCT_TYPE_PHYSICAL &#x3D; &#39;1&#39; PRODUCT_TYPE_SERVICE &#x3D; &#39;2&#39; PRODUCT_TYPE_DIGITAL &#x3D; &#39;3&#39; PRODUCT_TYPE_SHIPPING &#x3D; &#39;4&#39; PRODUCT_TYPE_EXEMPT &#x3D; &#39;5&#39; PRODUCT_TYPE_REDUCED_TAX &#x3D; &#39;6&#39; PRODUCT_TYPE_OVERRIDE_TAX &#x3D; &#39;7&#39; PRODUCT_TYPE_ZERO_RATED &#x3D; &#39;8&#39; PRODUCT_TYPE_REVERSE_TAX &#x3D; &#39;9&#39; &#x60;&#x60;&#x60;  | [optional] [default to "1"]

## Methods

### NewProduct

`func NewProduct() *Product`

NewProduct instantiates a new Product object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductWithDefaults

`func NewProductWithDefaults() *Product`

NewProductWithDefaults instantiates a new Product object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Product) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Product) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Product) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Product) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *Product) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Product) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Product) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Product) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAssignedUserId

`func (o *Product) GetAssignedUserId() string`

GetAssignedUserId returns the AssignedUserId field if non-nil, zero value otherwise.

### GetAssignedUserIdOk

`func (o *Product) GetAssignedUserIdOk() (*string, bool)`

GetAssignedUserIdOk returns a tuple with the AssignedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUserId

`func (o *Product) SetAssignedUserId(v string)`

SetAssignedUserId sets AssignedUserId field to given value.

### HasAssignedUserId

`func (o *Product) HasAssignedUserId() bool`

HasAssignedUserId returns a boolean if a field has been set.

### GetProjectId

`func (o *Product) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Product) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Product) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Product) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetVendorId

`func (o *Product) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *Product) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *Product) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *Product) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetCustomValue1

`func (o *Product) GetCustomValue1() string`

GetCustomValue1 returns the CustomValue1 field if non-nil, zero value otherwise.

### GetCustomValue1Ok

`func (o *Product) GetCustomValue1Ok() (*string, bool)`

GetCustomValue1Ok returns a tuple with the CustomValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue1

`func (o *Product) SetCustomValue1(v string)`

SetCustomValue1 sets CustomValue1 field to given value.

### HasCustomValue1

`func (o *Product) HasCustomValue1() bool`

HasCustomValue1 returns a boolean if a field has been set.

### GetCustomValue2

`func (o *Product) GetCustomValue2() string`

GetCustomValue2 returns the CustomValue2 field if non-nil, zero value otherwise.

### GetCustomValue2Ok

`func (o *Product) GetCustomValue2Ok() (*string, bool)`

GetCustomValue2Ok returns a tuple with the CustomValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue2

`func (o *Product) SetCustomValue2(v string)`

SetCustomValue2 sets CustomValue2 field to given value.

### HasCustomValue2

`func (o *Product) HasCustomValue2() bool`

HasCustomValue2 returns a boolean if a field has been set.

### GetCustomValue3

`func (o *Product) GetCustomValue3() string`

GetCustomValue3 returns the CustomValue3 field if non-nil, zero value otherwise.

### GetCustomValue3Ok

`func (o *Product) GetCustomValue3Ok() (*string, bool)`

GetCustomValue3Ok returns a tuple with the CustomValue3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue3

`func (o *Product) SetCustomValue3(v string)`

SetCustomValue3 sets CustomValue3 field to given value.

### HasCustomValue3

`func (o *Product) HasCustomValue3() bool`

HasCustomValue3 returns a boolean if a field has been set.

### GetCustomValue4

`func (o *Product) GetCustomValue4() string`

GetCustomValue4 returns the CustomValue4 field if non-nil, zero value otherwise.

### GetCustomValue4Ok

`func (o *Product) GetCustomValue4Ok() (*string, bool)`

GetCustomValue4Ok returns a tuple with the CustomValue4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue4

`func (o *Product) SetCustomValue4(v string)`

SetCustomValue4 sets CustomValue4 field to given value.

### HasCustomValue4

`func (o *Product) HasCustomValue4() bool`

HasCustomValue4 returns a boolean if a field has been set.

### GetProductKey

`func (o *Product) GetProductKey() string`

GetProductKey returns the ProductKey field if non-nil, zero value otherwise.

### GetProductKeyOk

`func (o *Product) GetProductKeyOk() (*string, bool)`

GetProductKeyOk returns a tuple with the ProductKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductKey

`func (o *Product) SetProductKey(v string)`

SetProductKey sets ProductKey field to given value.

### HasProductKey

`func (o *Product) HasProductKey() bool`

HasProductKey returns a boolean if a field has been set.

### GetNotes

`func (o *Product) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Product) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Product) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Product) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetCost

`func (o *Product) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *Product) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *Product) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *Product) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetPrice

`func (o *Product) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Product) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Product) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Product) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *Product) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Product) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Product) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *Product) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetTaxName1

`func (o *Product) GetTaxName1() string`

GetTaxName1 returns the TaxName1 field if non-nil, zero value otherwise.

### GetTaxName1Ok

`func (o *Product) GetTaxName1Ok() (*string, bool)`

GetTaxName1Ok returns a tuple with the TaxName1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName1

`func (o *Product) SetTaxName1(v string)`

SetTaxName1 sets TaxName1 field to given value.

### HasTaxName1

`func (o *Product) HasTaxName1() bool`

HasTaxName1 returns a boolean if a field has been set.

### GetTaxRate1

`func (o *Product) GetTaxRate1() float64`

GetTaxRate1 returns the TaxRate1 field if non-nil, zero value otherwise.

### GetTaxRate1Ok

`func (o *Product) GetTaxRate1Ok() (*float64, bool)`

GetTaxRate1Ok returns a tuple with the TaxRate1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate1

`func (o *Product) SetTaxRate1(v float64)`

SetTaxRate1 sets TaxRate1 field to given value.

### HasTaxRate1

`func (o *Product) HasTaxRate1() bool`

HasTaxRate1 returns a boolean if a field has been set.

### GetTaxName2

`func (o *Product) GetTaxName2() string`

GetTaxName2 returns the TaxName2 field if non-nil, zero value otherwise.

### GetTaxName2Ok

`func (o *Product) GetTaxName2Ok() (*string, bool)`

GetTaxName2Ok returns a tuple with the TaxName2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName2

`func (o *Product) SetTaxName2(v string)`

SetTaxName2 sets TaxName2 field to given value.

### HasTaxName2

`func (o *Product) HasTaxName2() bool`

HasTaxName2 returns a boolean if a field has been set.

### GetTaxRate2

`func (o *Product) GetTaxRate2() float64`

GetTaxRate2 returns the TaxRate2 field if non-nil, zero value otherwise.

### GetTaxRate2Ok

`func (o *Product) GetTaxRate2Ok() (*float64, bool)`

GetTaxRate2Ok returns a tuple with the TaxRate2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate2

`func (o *Product) SetTaxRate2(v float64)`

SetTaxRate2 sets TaxRate2 field to given value.

### HasTaxRate2

`func (o *Product) HasTaxRate2() bool`

HasTaxRate2 returns a boolean if a field has been set.

### GetTaxName3

`func (o *Product) GetTaxName3() string`

GetTaxName3 returns the TaxName3 field if non-nil, zero value otherwise.

### GetTaxName3Ok

`func (o *Product) GetTaxName3Ok() (*string, bool)`

GetTaxName3Ok returns a tuple with the TaxName3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName3

`func (o *Product) SetTaxName3(v string)`

SetTaxName3 sets TaxName3 field to given value.

### HasTaxName3

`func (o *Product) HasTaxName3() bool`

HasTaxName3 returns a boolean if a field has been set.

### GetTaxRate3

`func (o *Product) GetTaxRate3() float64`

GetTaxRate3 returns the TaxRate3 field if non-nil, zero value otherwise.

### GetTaxRate3Ok

`func (o *Product) GetTaxRate3Ok() (*float64, bool)`

GetTaxRate3Ok returns a tuple with the TaxRate3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate3

`func (o *Product) SetTaxRate3(v float64)`

SetTaxRate3 sets TaxRate3 field to given value.

### HasTaxRate3

`func (o *Product) HasTaxRate3() bool`

HasTaxRate3 returns a boolean if a field has been set.

### GetArchivedAt

`func (o *Product) GetArchivedAt() int32`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *Product) GetArchivedAtOk() (*int32, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *Product) SetArchivedAt(v int32)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *Product) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Product) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Product) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Product) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Product) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Product) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Product) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Product) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Product) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetIsDeleted

`func (o *Product) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *Product) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *Product) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *Product) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetInStockQuantity

`func (o *Product) GetInStockQuantity() int32`

GetInStockQuantity returns the InStockQuantity field if non-nil, zero value otherwise.

### GetInStockQuantityOk

`func (o *Product) GetInStockQuantityOk() (*int32, bool)`

GetInStockQuantityOk returns a tuple with the InStockQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInStockQuantity

`func (o *Product) SetInStockQuantity(v int32)`

SetInStockQuantity sets InStockQuantity field to given value.

### HasInStockQuantity

`func (o *Product) HasInStockQuantity() bool`

HasInStockQuantity returns a boolean if a field has been set.

### GetStockNotification

`func (o *Product) GetStockNotification() bool`

GetStockNotification returns the StockNotification field if non-nil, zero value otherwise.

### GetStockNotificationOk

`func (o *Product) GetStockNotificationOk() (*bool, bool)`

GetStockNotificationOk returns a tuple with the StockNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockNotification

`func (o *Product) SetStockNotification(v bool)`

SetStockNotification sets StockNotification field to given value.

### HasStockNotification

`func (o *Product) HasStockNotification() bool`

HasStockNotification returns a boolean if a field has been set.

### GetStockNotificationThreshold

`func (o *Product) GetStockNotificationThreshold() int32`

GetStockNotificationThreshold returns the StockNotificationThreshold field if non-nil, zero value otherwise.

### GetStockNotificationThresholdOk

`func (o *Product) GetStockNotificationThresholdOk() (*int32, bool)`

GetStockNotificationThresholdOk returns a tuple with the StockNotificationThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockNotificationThreshold

`func (o *Product) SetStockNotificationThreshold(v int32)`

SetStockNotificationThreshold sets StockNotificationThreshold field to given value.

### HasStockNotificationThreshold

`func (o *Product) HasStockNotificationThreshold() bool`

HasStockNotificationThreshold returns a boolean if a field has been set.

### GetMaxQuantity

`func (o *Product) GetMaxQuantity() int32`

GetMaxQuantity returns the MaxQuantity field if non-nil, zero value otherwise.

### GetMaxQuantityOk

`func (o *Product) GetMaxQuantityOk() (*int32, bool)`

GetMaxQuantityOk returns a tuple with the MaxQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQuantity

`func (o *Product) SetMaxQuantity(v int32)`

SetMaxQuantity sets MaxQuantity field to given value.

### HasMaxQuantity

`func (o *Product) HasMaxQuantity() bool`

HasMaxQuantity returns a boolean if a field has been set.

### GetProductImage

`func (o *Product) GetProductImage() string`

GetProductImage returns the ProductImage field if non-nil, zero value otherwise.

### GetProductImageOk

`func (o *Product) GetProductImageOk() (*string, bool)`

GetProductImageOk returns a tuple with the ProductImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductImage

`func (o *Product) SetProductImage(v string)`

SetProductImage sets ProductImage field to given value.

### HasProductImage

`func (o *Product) HasProductImage() bool`

HasProductImage returns a boolean if a field has been set.

### GetTaxId

`func (o *Product) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *Product) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *Product) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *Product) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


