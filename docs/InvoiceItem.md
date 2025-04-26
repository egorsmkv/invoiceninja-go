# InvoiceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | Pointer to **float32** | The quantity of the product offered for this line item | [optional] 
**Cost** | Pointer to **float32** | The cost of the product offered for this line item | [optional] 
**ProductKey** | Pointer to **string** | The product key of the product offered for this line item (Referred to as Product in the product tab) | [optional] 
**ProductCost** | Pointer to **float32** | The cost of the product offered for this line item (Referred to as Cost in the product tab) | [optional] 
**Notes** | Pointer to **string** | The notes/description for the product offered for this line item | [optional] 
**Discount** | Pointer to **float32** | The discount applied to the product offered for this line item | [optional] 
**IsAmountDiscount** | Pointer to **bool** | Indicates whether the discount applied to the product offered for this line item is a fixed amount or a percentage | [optional] 
**TaxName1** | Pointer to **string** | The name of the first tax applied to the product offered for this line item | [optional] 
**TaxRate1** | Pointer to **float32** | The rate of the first tax applied to the product offered for this line item | [optional] 
**TaxName2** | Pointer to **string** | The name of the second tax applied to the product offered for this line item | [optional] 
**TaxRate2** | Pointer to **float32** | The rate of the second tax applied to the product offered for this line item | [optional] 
**TaxName3** | Pointer to **string** | The name of the third tax applied to the product offered for this line item | [optional] 
**TaxRate3** | Pointer to **float32** | The rate of the third tax applied to the product offered for this line item | [optional] 
**SortId** | Pointer to **string** | Deprecated | [optional] 
**LineTotal** | Pointer to **float32** | The total amount of the product offered for this line item | [optional] [readonly] 
**GrossLineTotal** | Pointer to **float32** | The total amount of the product offered for this line item before discounts | [optional] [readonly] 
**TaxAmount** | Pointer to **float32** | The total amount of tax applied to the product offered for this line item | [optional] [readonly] 
**Date** | Pointer to **time.Time** | Deprecated | [optional] 
**CustomValue1** | Pointer to **string** | The first custom value of the product offered for this line item | [optional] 
**CustomValue2** | Pointer to **string** | The second custom value of the product offered for this line item | [optional] 
**CustomValue3** | Pointer to **string** | The third custom value of the product offered for this line item | [optional] 
**CustomValue4** | Pointer to **string** | The fourth custom value of the product offered for this line item | [optional] 
**TypeId** | Pointer to **string** | 1 &#x3D; product, 2 &#x3D; service, 3 unpaid gateway fee, 4 paid gateway fee, 5 late fee, 6 expense | [optional] [default to "1"]
**TaxId** | Pointer to **string** | The tax ID of the product: 1 product, 2 service, 3 digital, 4 shipping, 5 exempt, 5 reduced tax, 7 override, 8 zero rate, 9 reverse tax | [optional] [default to "1"]

## Methods

### NewInvoiceItem

`func NewInvoiceItem() *InvoiceItem`

NewInvoiceItem instantiates a new InvoiceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceItemWithDefaults

`func NewInvoiceItemWithDefaults() *InvoiceItem`

NewInvoiceItemWithDefaults instantiates a new InvoiceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *InvoiceItem) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InvoiceItem) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InvoiceItem) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *InvoiceItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetCost

`func (o *InvoiceItem) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *InvoiceItem) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *InvoiceItem) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *InvoiceItem) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetProductKey

`func (o *InvoiceItem) GetProductKey() string`

GetProductKey returns the ProductKey field if non-nil, zero value otherwise.

### GetProductKeyOk

`func (o *InvoiceItem) GetProductKeyOk() (*string, bool)`

GetProductKeyOk returns a tuple with the ProductKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductKey

`func (o *InvoiceItem) SetProductKey(v string)`

SetProductKey sets ProductKey field to given value.

### HasProductKey

`func (o *InvoiceItem) HasProductKey() bool`

HasProductKey returns a boolean if a field has been set.

### GetProductCost

`func (o *InvoiceItem) GetProductCost() float32`

GetProductCost returns the ProductCost field if non-nil, zero value otherwise.

### GetProductCostOk

`func (o *InvoiceItem) GetProductCostOk() (*float32, bool)`

GetProductCostOk returns a tuple with the ProductCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCost

`func (o *InvoiceItem) SetProductCost(v float32)`

SetProductCost sets ProductCost field to given value.

### HasProductCost

`func (o *InvoiceItem) HasProductCost() bool`

HasProductCost returns a boolean if a field has been set.

### GetNotes

`func (o *InvoiceItem) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *InvoiceItem) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *InvoiceItem) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *InvoiceItem) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetDiscount

`func (o *InvoiceItem) GetDiscount() float32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *InvoiceItem) GetDiscountOk() (*float32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *InvoiceItem) SetDiscount(v float32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *InvoiceItem) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetIsAmountDiscount

`func (o *InvoiceItem) GetIsAmountDiscount() bool`

GetIsAmountDiscount returns the IsAmountDiscount field if non-nil, zero value otherwise.

### GetIsAmountDiscountOk

`func (o *InvoiceItem) GetIsAmountDiscountOk() (*bool, bool)`

GetIsAmountDiscountOk returns a tuple with the IsAmountDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAmountDiscount

`func (o *InvoiceItem) SetIsAmountDiscount(v bool)`

SetIsAmountDiscount sets IsAmountDiscount field to given value.

### HasIsAmountDiscount

`func (o *InvoiceItem) HasIsAmountDiscount() bool`

HasIsAmountDiscount returns a boolean if a field has been set.

### GetTaxName1

`func (o *InvoiceItem) GetTaxName1() string`

GetTaxName1 returns the TaxName1 field if non-nil, zero value otherwise.

### GetTaxName1Ok

`func (o *InvoiceItem) GetTaxName1Ok() (*string, bool)`

GetTaxName1Ok returns a tuple with the TaxName1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName1

`func (o *InvoiceItem) SetTaxName1(v string)`

SetTaxName1 sets TaxName1 field to given value.

### HasTaxName1

`func (o *InvoiceItem) HasTaxName1() bool`

HasTaxName1 returns a boolean if a field has been set.

### GetTaxRate1

`func (o *InvoiceItem) GetTaxRate1() float32`

GetTaxRate1 returns the TaxRate1 field if non-nil, zero value otherwise.

### GetTaxRate1Ok

`func (o *InvoiceItem) GetTaxRate1Ok() (*float32, bool)`

GetTaxRate1Ok returns a tuple with the TaxRate1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate1

`func (o *InvoiceItem) SetTaxRate1(v float32)`

SetTaxRate1 sets TaxRate1 field to given value.

### HasTaxRate1

`func (o *InvoiceItem) HasTaxRate1() bool`

HasTaxRate1 returns a boolean if a field has been set.

### GetTaxName2

`func (o *InvoiceItem) GetTaxName2() string`

GetTaxName2 returns the TaxName2 field if non-nil, zero value otherwise.

### GetTaxName2Ok

`func (o *InvoiceItem) GetTaxName2Ok() (*string, bool)`

GetTaxName2Ok returns a tuple with the TaxName2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName2

`func (o *InvoiceItem) SetTaxName2(v string)`

SetTaxName2 sets TaxName2 field to given value.

### HasTaxName2

`func (o *InvoiceItem) HasTaxName2() bool`

HasTaxName2 returns a boolean if a field has been set.

### GetTaxRate2

`func (o *InvoiceItem) GetTaxRate2() float32`

GetTaxRate2 returns the TaxRate2 field if non-nil, zero value otherwise.

### GetTaxRate2Ok

`func (o *InvoiceItem) GetTaxRate2Ok() (*float32, bool)`

GetTaxRate2Ok returns a tuple with the TaxRate2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate2

`func (o *InvoiceItem) SetTaxRate2(v float32)`

SetTaxRate2 sets TaxRate2 field to given value.

### HasTaxRate2

`func (o *InvoiceItem) HasTaxRate2() bool`

HasTaxRate2 returns a boolean if a field has been set.

### GetTaxName3

`func (o *InvoiceItem) GetTaxName3() string`

GetTaxName3 returns the TaxName3 field if non-nil, zero value otherwise.

### GetTaxName3Ok

`func (o *InvoiceItem) GetTaxName3Ok() (*string, bool)`

GetTaxName3Ok returns a tuple with the TaxName3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName3

`func (o *InvoiceItem) SetTaxName3(v string)`

SetTaxName3 sets TaxName3 field to given value.

### HasTaxName3

`func (o *InvoiceItem) HasTaxName3() bool`

HasTaxName3 returns a boolean if a field has been set.

### GetTaxRate3

`func (o *InvoiceItem) GetTaxRate3() float32`

GetTaxRate3 returns the TaxRate3 field if non-nil, zero value otherwise.

### GetTaxRate3Ok

`func (o *InvoiceItem) GetTaxRate3Ok() (*float32, bool)`

GetTaxRate3Ok returns a tuple with the TaxRate3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate3

`func (o *InvoiceItem) SetTaxRate3(v float32)`

SetTaxRate3 sets TaxRate3 field to given value.

### HasTaxRate3

`func (o *InvoiceItem) HasTaxRate3() bool`

HasTaxRate3 returns a boolean if a field has been set.

### GetSortId

`func (o *InvoiceItem) GetSortId() string`

GetSortId returns the SortId field if non-nil, zero value otherwise.

### GetSortIdOk

`func (o *InvoiceItem) GetSortIdOk() (*string, bool)`

GetSortIdOk returns a tuple with the SortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortId

`func (o *InvoiceItem) SetSortId(v string)`

SetSortId sets SortId field to given value.

### HasSortId

`func (o *InvoiceItem) HasSortId() bool`

HasSortId returns a boolean if a field has been set.

### GetLineTotal

`func (o *InvoiceItem) GetLineTotal() float32`

GetLineTotal returns the LineTotal field if non-nil, zero value otherwise.

### GetLineTotalOk

`func (o *InvoiceItem) GetLineTotalOk() (*float32, bool)`

GetLineTotalOk returns a tuple with the LineTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineTotal

`func (o *InvoiceItem) SetLineTotal(v float32)`

SetLineTotal sets LineTotal field to given value.

### HasLineTotal

`func (o *InvoiceItem) HasLineTotal() bool`

HasLineTotal returns a boolean if a field has been set.

### GetGrossLineTotal

`func (o *InvoiceItem) GetGrossLineTotal() float32`

GetGrossLineTotal returns the GrossLineTotal field if non-nil, zero value otherwise.

### GetGrossLineTotalOk

`func (o *InvoiceItem) GetGrossLineTotalOk() (*float32, bool)`

GetGrossLineTotalOk returns a tuple with the GrossLineTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossLineTotal

`func (o *InvoiceItem) SetGrossLineTotal(v float32)`

SetGrossLineTotal sets GrossLineTotal field to given value.

### HasGrossLineTotal

`func (o *InvoiceItem) HasGrossLineTotal() bool`

HasGrossLineTotal returns a boolean if a field has been set.

### GetTaxAmount

`func (o *InvoiceItem) GetTaxAmount() float32`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *InvoiceItem) GetTaxAmountOk() (*float32, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *InvoiceItem) SetTaxAmount(v float32)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *InvoiceItem) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetDate

`func (o *InvoiceItem) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *InvoiceItem) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *InvoiceItem) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *InvoiceItem) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetCustomValue1

`func (o *InvoiceItem) GetCustomValue1() string`

GetCustomValue1 returns the CustomValue1 field if non-nil, zero value otherwise.

### GetCustomValue1Ok

`func (o *InvoiceItem) GetCustomValue1Ok() (*string, bool)`

GetCustomValue1Ok returns a tuple with the CustomValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue1

`func (o *InvoiceItem) SetCustomValue1(v string)`

SetCustomValue1 sets CustomValue1 field to given value.

### HasCustomValue1

`func (o *InvoiceItem) HasCustomValue1() bool`

HasCustomValue1 returns a boolean if a field has been set.

### GetCustomValue2

`func (o *InvoiceItem) GetCustomValue2() string`

GetCustomValue2 returns the CustomValue2 field if non-nil, zero value otherwise.

### GetCustomValue2Ok

`func (o *InvoiceItem) GetCustomValue2Ok() (*string, bool)`

GetCustomValue2Ok returns a tuple with the CustomValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue2

`func (o *InvoiceItem) SetCustomValue2(v string)`

SetCustomValue2 sets CustomValue2 field to given value.

### HasCustomValue2

`func (o *InvoiceItem) HasCustomValue2() bool`

HasCustomValue2 returns a boolean if a field has been set.

### GetCustomValue3

`func (o *InvoiceItem) GetCustomValue3() string`

GetCustomValue3 returns the CustomValue3 field if non-nil, zero value otherwise.

### GetCustomValue3Ok

`func (o *InvoiceItem) GetCustomValue3Ok() (*string, bool)`

GetCustomValue3Ok returns a tuple with the CustomValue3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue3

`func (o *InvoiceItem) SetCustomValue3(v string)`

SetCustomValue3 sets CustomValue3 field to given value.

### HasCustomValue3

`func (o *InvoiceItem) HasCustomValue3() bool`

HasCustomValue3 returns a boolean if a field has been set.

### GetCustomValue4

`func (o *InvoiceItem) GetCustomValue4() string`

GetCustomValue4 returns the CustomValue4 field if non-nil, zero value otherwise.

### GetCustomValue4Ok

`func (o *InvoiceItem) GetCustomValue4Ok() (*string, bool)`

GetCustomValue4Ok returns a tuple with the CustomValue4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue4

`func (o *InvoiceItem) SetCustomValue4(v string)`

SetCustomValue4 sets CustomValue4 field to given value.

### HasCustomValue4

`func (o *InvoiceItem) HasCustomValue4() bool`

HasCustomValue4 returns a boolean if a field has been set.

### GetTypeId

`func (o *InvoiceItem) GetTypeId() string`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *InvoiceItem) GetTypeIdOk() (*string, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *InvoiceItem) SetTypeId(v string)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *InvoiceItem) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetTaxId

`func (o *InvoiceItem) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *InvoiceItem) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *InvoiceItem) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *InvoiceItem) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


