# FillableInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignedUserId** | Pointer to **string** | The assigned user&#39;s hashed ID | [optional] 
**ClientId** | Pointer to **string** | The client&#39;s hashed ID | [optional] 
**Number** | Pointer to **string** | The unique alphanumeric invoice number for each invoice per company | [optional] 
**PoNumber** | Pointer to **string** | The purchase order number associated with the invoice | [optional] 
**Terms** | Pointer to **string** | The terms and conditions for the invoice | [optional] 
**PublicNotes** | Pointer to **string** | Public notes visible to the client on the invoice | [optional] 
**PrivateNotes** | Pointer to **string** | Private notes for internal use only | [optional] 
**Footer** | Pointer to **string** | The footer text displayed on the invoice | [optional] 
**CustomValue1** | Pointer to **string** | First custom value for additional information | [optional] 
**CustomValue2** | Pointer to **string** | Second custom value for additional information | [optional] 
**CustomValue3** | Pointer to **string** | Third custom value for additional information | [optional] 
**CustomValue4** | Pointer to **string** | Fourth custom value for additional information | [optional] 
**TaxName1** | Pointer to **string** | Name of the first tax applied to the invoice | [optional] 
**TaxName2** | Pointer to **string** | Name of the second tax applied to the invoice | [optional] 
**TaxRate1** | Pointer to **float32** | Rate of the first tax applied to the invoice | [optional] 
**TaxRate2** | Pointer to **float32** | Rate of the second tax applied to the invoice | [optional] 
**TaxName3** | Pointer to **string** | Name of the third tax applied to the invoice | [optional] 
**TaxRate3** | Pointer to **float32** | Rate of the third tax applied to the invoice | [optional] 
**LineItems** | Pointer to [**[]InvoiceItem**](InvoiceItem.md) | An array of objects which define the line items of the invoice | [optional] 
**Discount** | Pointer to **float32** | The discount applied to the invoice | [optional] 
**Partial** | Pointer to **float32** | The partial amount applied to the invoice | [optional] 
**IsAmountDiscount** | Pointer to **bool** | Indicates whether the discount applied is a fixed amount or a percentage | [optional] 
**UsesInclusiveTaxes** | Pointer to **bool** | Indicates whether the tax rates applied to the invoice are inclusive or exclusive | [optional] 
**Date** | Pointer to **string** | The date the invoice was issued | [optional] 
**PartialDueDate** | Pointer to **string** | The due date for the partial payment | [optional] 
**DueDate** | Pointer to **string** | The due date for the invoice | [optional] 
**CustomSurcharge1** | Pointer to **float32** | First custom surcharge applied to the invoice | [optional] 
**CustomSurcharge2** | Pointer to **float32** | Second custom surcharge applied to the invoice | [optional] 
**CustomSurcharge3** | Pointer to **float32** | Third custom surcharge applied to the invoice | [optional] 
**CustomSurcharge4** | Pointer to **float32** | Fourth custom surcharge applied to the invoice | [optional] 

## Methods

### NewFillableInvoice

`func NewFillableInvoice() *FillableInvoice`

NewFillableInvoice instantiates a new FillableInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFillableInvoiceWithDefaults

`func NewFillableInvoiceWithDefaults() *FillableInvoice`

NewFillableInvoiceWithDefaults instantiates a new FillableInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignedUserId

`func (o *FillableInvoice) GetAssignedUserId() string`

GetAssignedUserId returns the AssignedUserId field if non-nil, zero value otherwise.

### GetAssignedUserIdOk

`func (o *FillableInvoice) GetAssignedUserIdOk() (*string, bool)`

GetAssignedUserIdOk returns a tuple with the AssignedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUserId

`func (o *FillableInvoice) SetAssignedUserId(v string)`

SetAssignedUserId sets AssignedUserId field to given value.

### HasAssignedUserId

`func (o *FillableInvoice) HasAssignedUserId() bool`

HasAssignedUserId returns a boolean if a field has been set.

### GetClientId

`func (o *FillableInvoice) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *FillableInvoice) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *FillableInvoice) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *FillableInvoice) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetNumber

`func (o *FillableInvoice) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *FillableInvoice) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *FillableInvoice) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *FillableInvoice) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetPoNumber

`func (o *FillableInvoice) GetPoNumber() string`

GetPoNumber returns the PoNumber field if non-nil, zero value otherwise.

### GetPoNumberOk

`func (o *FillableInvoice) GetPoNumberOk() (*string, bool)`

GetPoNumberOk returns a tuple with the PoNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoNumber

`func (o *FillableInvoice) SetPoNumber(v string)`

SetPoNumber sets PoNumber field to given value.

### HasPoNumber

`func (o *FillableInvoice) HasPoNumber() bool`

HasPoNumber returns a boolean if a field has been set.

### GetTerms

`func (o *FillableInvoice) GetTerms() string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *FillableInvoice) GetTermsOk() (*string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *FillableInvoice) SetTerms(v string)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *FillableInvoice) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### GetPublicNotes

`func (o *FillableInvoice) GetPublicNotes() string`

GetPublicNotes returns the PublicNotes field if non-nil, zero value otherwise.

### GetPublicNotesOk

`func (o *FillableInvoice) GetPublicNotesOk() (*string, bool)`

GetPublicNotesOk returns a tuple with the PublicNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNotes

`func (o *FillableInvoice) SetPublicNotes(v string)`

SetPublicNotes sets PublicNotes field to given value.

### HasPublicNotes

`func (o *FillableInvoice) HasPublicNotes() bool`

HasPublicNotes returns a boolean if a field has been set.

### GetPrivateNotes

`func (o *FillableInvoice) GetPrivateNotes() string`

GetPrivateNotes returns the PrivateNotes field if non-nil, zero value otherwise.

### GetPrivateNotesOk

`func (o *FillableInvoice) GetPrivateNotesOk() (*string, bool)`

GetPrivateNotesOk returns a tuple with the PrivateNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNotes

`func (o *FillableInvoice) SetPrivateNotes(v string)`

SetPrivateNotes sets PrivateNotes field to given value.

### HasPrivateNotes

`func (o *FillableInvoice) HasPrivateNotes() bool`

HasPrivateNotes returns a boolean if a field has been set.

### GetFooter

`func (o *FillableInvoice) GetFooter() string`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *FillableInvoice) GetFooterOk() (*string, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *FillableInvoice) SetFooter(v string)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *FillableInvoice) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetCustomValue1

`func (o *FillableInvoice) GetCustomValue1() string`

GetCustomValue1 returns the CustomValue1 field if non-nil, zero value otherwise.

### GetCustomValue1Ok

`func (o *FillableInvoice) GetCustomValue1Ok() (*string, bool)`

GetCustomValue1Ok returns a tuple with the CustomValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue1

`func (o *FillableInvoice) SetCustomValue1(v string)`

SetCustomValue1 sets CustomValue1 field to given value.

### HasCustomValue1

`func (o *FillableInvoice) HasCustomValue1() bool`

HasCustomValue1 returns a boolean if a field has been set.

### GetCustomValue2

`func (o *FillableInvoice) GetCustomValue2() string`

GetCustomValue2 returns the CustomValue2 field if non-nil, zero value otherwise.

### GetCustomValue2Ok

`func (o *FillableInvoice) GetCustomValue2Ok() (*string, bool)`

GetCustomValue2Ok returns a tuple with the CustomValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue2

`func (o *FillableInvoice) SetCustomValue2(v string)`

SetCustomValue2 sets CustomValue2 field to given value.

### HasCustomValue2

`func (o *FillableInvoice) HasCustomValue2() bool`

HasCustomValue2 returns a boolean if a field has been set.

### GetCustomValue3

`func (o *FillableInvoice) GetCustomValue3() string`

GetCustomValue3 returns the CustomValue3 field if non-nil, zero value otherwise.

### GetCustomValue3Ok

`func (o *FillableInvoice) GetCustomValue3Ok() (*string, bool)`

GetCustomValue3Ok returns a tuple with the CustomValue3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue3

`func (o *FillableInvoice) SetCustomValue3(v string)`

SetCustomValue3 sets CustomValue3 field to given value.

### HasCustomValue3

`func (o *FillableInvoice) HasCustomValue3() bool`

HasCustomValue3 returns a boolean if a field has been set.

### GetCustomValue4

`func (o *FillableInvoice) GetCustomValue4() string`

GetCustomValue4 returns the CustomValue4 field if non-nil, zero value otherwise.

### GetCustomValue4Ok

`func (o *FillableInvoice) GetCustomValue4Ok() (*string, bool)`

GetCustomValue4Ok returns a tuple with the CustomValue4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue4

`func (o *FillableInvoice) SetCustomValue4(v string)`

SetCustomValue4 sets CustomValue4 field to given value.

### HasCustomValue4

`func (o *FillableInvoice) HasCustomValue4() bool`

HasCustomValue4 returns a boolean if a field has been set.

### GetTaxName1

`func (o *FillableInvoice) GetTaxName1() string`

GetTaxName1 returns the TaxName1 field if non-nil, zero value otherwise.

### GetTaxName1Ok

`func (o *FillableInvoice) GetTaxName1Ok() (*string, bool)`

GetTaxName1Ok returns a tuple with the TaxName1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName1

`func (o *FillableInvoice) SetTaxName1(v string)`

SetTaxName1 sets TaxName1 field to given value.

### HasTaxName1

`func (o *FillableInvoice) HasTaxName1() bool`

HasTaxName1 returns a boolean if a field has been set.

### GetTaxName2

`func (o *FillableInvoice) GetTaxName2() string`

GetTaxName2 returns the TaxName2 field if non-nil, zero value otherwise.

### GetTaxName2Ok

`func (o *FillableInvoice) GetTaxName2Ok() (*string, bool)`

GetTaxName2Ok returns a tuple with the TaxName2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName2

`func (o *FillableInvoice) SetTaxName2(v string)`

SetTaxName2 sets TaxName2 field to given value.

### HasTaxName2

`func (o *FillableInvoice) HasTaxName2() bool`

HasTaxName2 returns a boolean if a field has been set.

### GetTaxRate1

`func (o *FillableInvoice) GetTaxRate1() float32`

GetTaxRate1 returns the TaxRate1 field if non-nil, zero value otherwise.

### GetTaxRate1Ok

`func (o *FillableInvoice) GetTaxRate1Ok() (*float32, bool)`

GetTaxRate1Ok returns a tuple with the TaxRate1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate1

`func (o *FillableInvoice) SetTaxRate1(v float32)`

SetTaxRate1 sets TaxRate1 field to given value.

### HasTaxRate1

`func (o *FillableInvoice) HasTaxRate1() bool`

HasTaxRate1 returns a boolean if a field has been set.

### GetTaxRate2

`func (o *FillableInvoice) GetTaxRate2() float32`

GetTaxRate2 returns the TaxRate2 field if non-nil, zero value otherwise.

### GetTaxRate2Ok

`func (o *FillableInvoice) GetTaxRate2Ok() (*float32, bool)`

GetTaxRate2Ok returns a tuple with the TaxRate2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate2

`func (o *FillableInvoice) SetTaxRate2(v float32)`

SetTaxRate2 sets TaxRate2 field to given value.

### HasTaxRate2

`func (o *FillableInvoice) HasTaxRate2() bool`

HasTaxRate2 returns a boolean if a field has been set.

### GetTaxName3

`func (o *FillableInvoice) GetTaxName3() string`

GetTaxName3 returns the TaxName3 field if non-nil, zero value otherwise.

### GetTaxName3Ok

`func (o *FillableInvoice) GetTaxName3Ok() (*string, bool)`

GetTaxName3Ok returns a tuple with the TaxName3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName3

`func (o *FillableInvoice) SetTaxName3(v string)`

SetTaxName3 sets TaxName3 field to given value.

### HasTaxName3

`func (o *FillableInvoice) HasTaxName3() bool`

HasTaxName3 returns a boolean if a field has been set.

### GetTaxRate3

`func (o *FillableInvoice) GetTaxRate3() float32`

GetTaxRate3 returns the TaxRate3 field if non-nil, zero value otherwise.

### GetTaxRate3Ok

`func (o *FillableInvoice) GetTaxRate3Ok() (*float32, bool)`

GetTaxRate3Ok returns a tuple with the TaxRate3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate3

`func (o *FillableInvoice) SetTaxRate3(v float32)`

SetTaxRate3 sets TaxRate3 field to given value.

### HasTaxRate3

`func (o *FillableInvoice) HasTaxRate3() bool`

HasTaxRate3 returns a boolean if a field has been set.

### GetLineItems

`func (o *FillableInvoice) GetLineItems() []InvoiceItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *FillableInvoice) GetLineItemsOk() (*[]InvoiceItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *FillableInvoice) SetLineItems(v []InvoiceItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *FillableInvoice) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetDiscount

`func (o *FillableInvoice) GetDiscount() float32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *FillableInvoice) GetDiscountOk() (*float32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *FillableInvoice) SetDiscount(v float32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *FillableInvoice) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetPartial

`func (o *FillableInvoice) GetPartial() float32`

GetPartial returns the Partial field if non-nil, zero value otherwise.

### GetPartialOk

`func (o *FillableInvoice) GetPartialOk() (*float32, bool)`

GetPartialOk returns a tuple with the Partial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartial

`func (o *FillableInvoice) SetPartial(v float32)`

SetPartial sets Partial field to given value.

### HasPartial

`func (o *FillableInvoice) HasPartial() bool`

HasPartial returns a boolean if a field has been set.

### GetIsAmountDiscount

`func (o *FillableInvoice) GetIsAmountDiscount() bool`

GetIsAmountDiscount returns the IsAmountDiscount field if non-nil, zero value otherwise.

### GetIsAmountDiscountOk

`func (o *FillableInvoice) GetIsAmountDiscountOk() (*bool, bool)`

GetIsAmountDiscountOk returns a tuple with the IsAmountDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAmountDiscount

`func (o *FillableInvoice) SetIsAmountDiscount(v bool)`

SetIsAmountDiscount sets IsAmountDiscount field to given value.

### HasIsAmountDiscount

`func (o *FillableInvoice) HasIsAmountDiscount() bool`

HasIsAmountDiscount returns a boolean if a field has been set.

### GetUsesInclusiveTaxes

`func (o *FillableInvoice) GetUsesInclusiveTaxes() bool`

GetUsesInclusiveTaxes returns the UsesInclusiveTaxes field if non-nil, zero value otherwise.

### GetUsesInclusiveTaxesOk

`func (o *FillableInvoice) GetUsesInclusiveTaxesOk() (*bool, bool)`

GetUsesInclusiveTaxesOk returns a tuple with the UsesInclusiveTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesInclusiveTaxes

`func (o *FillableInvoice) SetUsesInclusiveTaxes(v bool)`

SetUsesInclusiveTaxes sets UsesInclusiveTaxes field to given value.

### HasUsesInclusiveTaxes

`func (o *FillableInvoice) HasUsesInclusiveTaxes() bool`

HasUsesInclusiveTaxes returns a boolean if a field has been set.

### GetDate

`func (o *FillableInvoice) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *FillableInvoice) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *FillableInvoice) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *FillableInvoice) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetPartialDueDate

`func (o *FillableInvoice) GetPartialDueDate() string`

GetPartialDueDate returns the PartialDueDate field if non-nil, zero value otherwise.

### GetPartialDueDateOk

`func (o *FillableInvoice) GetPartialDueDateOk() (*string, bool)`

GetPartialDueDateOk returns a tuple with the PartialDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialDueDate

`func (o *FillableInvoice) SetPartialDueDate(v string)`

SetPartialDueDate sets PartialDueDate field to given value.

### HasPartialDueDate

`func (o *FillableInvoice) HasPartialDueDate() bool`

HasPartialDueDate returns a boolean if a field has been set.

### GetDueDate

`func (o *FillableInvoice) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *FillableInvoice) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *FillableInvoice) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *FillableInvoice) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetCustomSurcharge1

`func (o *FillableInvoice) GetCustomSurcharge1() float32`

GetCustomSurcharge1 returns the CustomSurcharge1 field if non-nil, zero value otherwise.

### GetCustomSurcharge1Ok

`func (o *FillableInvoice) GetCustomSurcharge1Ok() (*float32, bool)`

GetCustomSurcharge1Ok returns a tuple with the CustomSurcharge1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurcharge1

`func (o *FillableInvoice) SetCustomSurcharge1(v float32)`

SetCustomSurcharge1 sets CustomSurcharge1 field to given value.

### HasCustomSurcharge1

`func (o *FillableInvoice) HasCustomSurcharge1() bool`

HasCustomSurcharge1 returns a boolean if a field has been set.

### GetCustomSurcharge2

`func (o *FillableInvoice) GetCustomSurcharge2() float32`

GetCustomSurcharge2 returns the CustomSurcharge2 field if non-nil, zero value otherwise.

### GetCustomSurcharge2Ok

`func (o *FillableInvoice) GetCustomSurcharge2Ok() (*float32, bool)`

GetCustomSurcharge2Ok returns a tuple with the CustomSurcharge2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurcharge2

`func (o *FillableInvoice) SetCustomSurcharge2(v float32)`

SetCustomSurcharge2 sets CustomSurcharge2 field to given value.

### HasCustomSurcharge2

`func (o *FillableInvoice) HasCustomSurcharge2() bool`

HasCustomSurcharge2 returns a boolean if a field has been set.

### GetCustomSurcharge3

`func (o *FillableInvoice) GetCustomSurcharge3() float32`

GetCustomSurcharge3 returns the CustomSurcharge3 field if non-nil, zero value otherwise.

### GetCustomSurcharge3Ok

`func (o *FillableInvoice) GetCustomSurcharge3Ok() (*float32, bool)`

GetCustomSurcharge3Ok returns a tuple with the CustomSurcharge3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurcharge3

`func (o *FillableInvoice) SetCustomSurcharge3(v float32)`

SetCustomSurcharge3 sets CustomSurcharge3 field to given value.

### HasCustomSurcharge3

`func (o *FillableInvoice) HasCustomSurcharge3() bool`

HasCustomSurcharge3 returns a boolean if a field has been set.

### GetCustomSurcharge4

`func (o *FillableInvoice) GetCustomSurcharge4() float32`

GetCustomSurcharge4 returns the CustomSurcharge4 field if non-nil, zero value otherwise.

### GetCustomSurcharge4Ok

`func (o *FillableInvoice) GetCustomSurcharge4Ok() (*float32, bool)`

GetCustomSurcharge4Ok returns a tuple with the CustomSurcharge4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurcharge4

`func (o *FillableInvoice) SetCustomSurcharge4(v float32)`

SetCustomSurcharge4 sets CustomSurcharge4 field to given value.

### HasCustomSurcharge4

`func (o *FillableInvoice) HasCustomSurcharge4() bool`

HasCustomSurcharge4 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


