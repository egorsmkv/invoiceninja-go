# RecurringInvoiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** | The user hashed id | [optional] 
**AssignedUserId** | Pointer to **string** | The assigned user hashed id | [optional] 
**ClientId** | Pointer to **string** | The client hashed id | [optional] 
**FrequencyId** | Pointer to **float32** | The recurring invoice frequency | [optional] 
**RemainingCycles** | Pointer to **float32** | The number of invoices left to be generated | [optional] 
**Number** | Pointer to **string** | The recurringinvoice number - is a unique alpha numeric number per invoice per company | [optional] 
**PoNumber** | Pointer to **string** | The purchase order associated with this recurring invoice | [optional] 
**Terms** | Pointer to **string** | The invoice terms | [optional] 
**PublicNotes** | Pointer to **string** | The public notes of the invoice | [optional] 
**PrivateNotes** | Pointer to **string** | The private notes of the invoice | [optional] 
**Footer** | Pointer to **string** | The invoice footer notes | [optional] 
**CustomValue1** | Pointer to **string** | A custom field value | [optional] 
**CustomValue2** | Pointer to **string** | A custom field value | [optional] 
**CustomValue3** | Pointer to **string** | A custom field value | [optional] 
**CustomValue4** | Pointer to **string** | A custom field value | [optional] 
**TaxName1** | Pointer to **string** | The tax name | [optional] 
**TaxName2** | Pointer to **string** | The tax name | [optional] 
**TaxRate1** | Pointer to **float32** | The tax rate | [optional] 
**TaxRate2** | Pointer to **float32** | The tax rate | [optional] 
**TaxName3** | Pointer to **string** | The tax name | [optional] 
**TaxRate3** | Pointer to **float32** | The tax rate | [optional] 
**LineItems** | Pointer to **map[string]interface{}** | An array of objects which define the line items of the invoice | [optional] 
**Discount** | Pointer to **float32** | The invoice discount, can be an amount or a percentage | [optional] 
**Partial** | Pointer to **float32** | The deposit/partial amount | [optional] 
**IsAmountDiscount** | Pointer to **bool** | Flag determining if the discount is an amount or a percentage | [optional] 
**UsesInclusiveTaxes** | Pointer to **bool** | Defines the type of taxes used as either inclusive or exclusive | [optional] 
**Date** | Pointer to **string** | The Invoice Date | [optional] 
**PartialDueDate** | Pointer to **string** | The due date for the deposit/partial amount | [optional] 
**DueDate** | Pointer to **string** | The due date of the invoice | [optional] 
**CustomSurcharge1** | Pointer to **float32** | First Custom Surcharge | [optional] 
**CustomSurcharge2** | Pointer to **float32** | Second Custom Surcharge | [optional] 
**CustomSurcharge3** | Pointer to **float32** | Third Custom Surcharge | [optional] 
**CustomSurcharge4** | Pointer to **float32** | Fourth Custom Surcharge | [optional] 
**CustomSurchargeTax1** | Pointer to **bool** | Toggles charging taxes on custom surcharge amounts | [optional] 
**CustomSurchargeTax2** | Pointer to **bool** | Toggles charging taxes on custom surcharge amounts | [optional] 
**CustomSurchargeTax3** | Pointer to **bool** | Toggles charging taxes on custom surcharge amounts | [optional] 
**CustomSurchargeTax4** | Pointer to **bool** | Toggles charging taxes on custom surcharge amounts | [optional] 
**LocationId** | Pointer to **string** | The client location id that this invoice relates to | [optional] 

## Methods

### NewRecurringInvoiceRequest

`func NewRecurringInvoiceRequest() *RecurringInvoiceRequest`

NewRecurringInvoiceRequest instantiates a new RecurringInvoiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurringInvoiceRequestWithDefaults

`func NewRecurringInvoiceRequestWithDefaults() *RecurringInvoiceRequest`

NewRecurringInvoiceRequestWithDefaults instantiates a new RecurringInvoiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *RecurringInvoiceRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RecurringInvoiceRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RecurringInvoiceRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *RecurringInvoiceRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAssignedUserId

`func (o *RecurringInvoiceRequest) GetAssignedUserId() string`

GetAssignedUserId returns the AssignedUserId field if non-nil, zero value otherwise.

### GetAssignedUserIdOk

`func (o *RecurringInvoiceRequest) GetAssignedUserIdOk() (*string, bool)`

GetAssignedUserIdOk returns a tuple with the AssignedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUserId

`func (o *RecurringInvoiceRequest) SetAssignedUserId(v string)`

SetAssignedUserId sets AssignedUserId field to given value.

### HasAssignedUserId

`func (o *RecurringInvoiceRequest) HasAssignedUserId() bool`

HasAssignedUserId returns a boolean if a field has been set.

### GetClientId

`func (o *RecurringInvoiceRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *RecurringInvoiceRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *RecurringInvoiceRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *RecurringInvoiceRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetFrequencyId

`func (o *RecurringInvoiceRequest) GetFrequencyId() float32`

GetFrequencyId returns the FrequencyId field if non-nil, zero value otherwise.

### GetFrequencyIdOk

`func (o *RecurringInvoiceRequest) GetFrequencyIdOk() (*float32, bool)`

GetFrequencyIdOk returns a tuple with the FrequencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyId

`func (o *RecurringInvoiceRequest) SetFrequencyId(v float32)`

SetFrequencyId sets FrequencyId field to given value.

### HasFrequencyId

`func (o *RecurringInvoiceRequest) HasFrequencyId() bool`

HasFrequencyId returns a boolean if a field has been set.

### GetRemainingCycles

`func (o *RecurringInvoiceRequest) GetRemainingCycles() float32`

GetRemainingCycles returns the RemainingCycles field if non-nil, zero value otherwise.

### GetRemainingCyclesOk

`func (o *RecurringInvoiceRequest) GetRemainingCyclesOk() (*float32, bool)`

GetRemainingCyclesOk returns a tuple with the RemainingCycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingCycles

`func (o *RecurringInvoiceRequest) SetRemainingCycles(v float32)`

SetRemainingCycles sets RemainingCycles field to given value.

### HasRemainingCycles

`func (o *RecurringInvoiceRequest) HasRemainingCycles() bool`

HasRemainingCycles returns a boolean if a field has been set.

### GetNumber

`func (o *RecurringInvoiceRequest) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *RecurringInvoiceRequest) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *RecurringInvoiceRequest) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *RecurringInvoiceRequest) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetPoNumber

`func (o *RecurringInvoiceRequest) GetPoNumber() string`

GetPoNumber returns the PoNumber field if non-nil, zero value otherwise.

### GetPoNumberOk

`func (o *RecurringInvoiceRequest) GetPoNumberOk() (*string, bool)`

GetPoNumberOk returns a tuple with the PoNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoNumber

`func (o *RecurringInvoiceRequest) SetPoNumber(v string)`

SetPoNumber sets PoNumber field to given value.

### HasPoNumber

`func (o *RecurringInvoiceRequest) HasPoNumber() bool`

HasPoNumber returns a boolean if a field has been set.

### GetTerms

`func (o *RecurringInvoiceRequest) GetTerms() string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *RecurringInvoiceRequest) GetTermsOk() (*string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *RecurringInvoiceRequest) SetTerms(v string)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *RecurringInvoiceRequest) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### GetPublicNotes

`func (o *RecurringInvoiceRequest) GetPublicNotes() string`

GetPublicNotes returns the PublicNotes field if non-nil, zero value otherwise.

### GetPublicNotesOk

`func (o *RecurringInvoiceRequest) GetPublicNotesOk() (*string, bool)`

GetPublicNotesOk returns a tuple with the PublicNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNotes

`func (o *RecurringInvoiceRequest) SetPublicNotes(v string)`

SetPublicNotes sets PublicNotes field to given value.

### HasPublicNotes

`func (o *RecurringInvoiceRequest) HasPublicNotes() bool`

HasPublicNotes returns a boolean if a field has been set.

### GetPrivateNotes

`func (o *RecurringInvoiceRequest) GetPrivateNotes() string`

GetPrivateNotes returns the PrivateNotes field if non-nil, zero value otherwise.

### GetPrivateNotesOk

`func (o *RecurringInvoiceRequest) GetPrivateNotesOk() (*string, bool)`

GetPrivateNotesOk returns a tuple with the PrivateNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNotes

`func (o *RecurringInvoiceRequest) SetPrivateNotes(v string)`

SetPrivateNotes sets PrivateNotes field to given value.

### HasPrivateNotes

`func (o *RecurringInvoiceRequest) HasPrivateNotes() bool`

HasPrivateNotes returns a boolean if a field has been set.

### GetFooter

`func (o *RecurringInvoiceRequest) GetFooter() string`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *RecurringInvoiceRequest) GetFooterOk() (*string, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *RecurringInvoiceRequest) SetFooter(v string)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *RecurringInvoiceRequest) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetCustomValue1

`func (o *RecurringInvoiceRequest) GetCustomValue1() string`

GetCustomValue1 returns the CustomValue1 field if non-nil, zero value otherwise.

### GetCustomValue1Ok

`func (o *RecurringInvoiceRequest) GetCustomValue1Ok() (*string, bool)`

GetCustomValue1Ok returns a tuple with the CustomValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue1

`func (o *RecurringInvoiceRequest) SetCustomValue1(v string)`

SetCustomValue1 sets CustomValue1 field to given value.

### HasCustomValue1

`func (o *RecurringInvoiceRequest) HasCustomValue1() bool`

HasCustomValue1 returns a boolean if a field has been set.

### GetCustomValue2

`func (o *RecurringInvoiceRequest) GetCustomValue2() string`

GetCustomValue2 returns the CustomValue2 field if non-nil, zero value otherwise.

### GetCustomValue2Ok

`func (o *RecurringInvoiceRequest) GetCustomValue2Ok() (*string, bool)`

GetCustomValue2Ok returns a tuple with the CustomValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue2

`func (o *RecurringInvoiceRequest) SetCustomValue2(v string)`

SetCustomValue2 sets CustomValue2 field to given value.

### HasCustomValue2

`func (o *RecurringInvoiceRequest) HasCustomValue2() bool`

HasCustomValue2 returns a boolean if a field has been set.

### GetCustomValue3

`func (o *RecurringInvoiceRequest) GetCustomValue3() string`

GetCustomValue3 returns the CustomValue3 field if non-nil, zero value otherwise.

### GetCustomValue3Ok

`func (o *RecurringInvoiceRequest) GetCustomValue3Ok() (*string, bool)`

GetCustomValue3Ok returns a tuple with the CustomValue3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue3

`func (o *RecurringInvoiceRequest) SetCustomValue3(v string)`

SetCustomValue3 sets CustomValue3 field to given value.

### HasCustomValue3

`func (o *RecurringInvoiceRequest) HasCustomValue3() bool`

HasCustomValue3 returns a boolean if a field has been set.

### GetCustomValue4

`func (o *RecurringInvoiceRequest) GetCustomValue4() string`

GetCustomValue4 returns the CustomValue4 field if non-nil, zero value otherwise.

### GetCustomValue4Ok

`func (o *RecurringInvoiceRequest) GetCustomValue4Ok() (*string, bool)`

GetCustomValue4Ok returns a tuple with the CustomValue4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue4

`func (o *RecurringInvoiceRequest) SetCustomValue4(v string)`

SetCustomValue4 sets CustomValue4 field to given value.

### HasCustomValue4

`func (o *RecurringInvoiceRequest) HasCustomValue4() bool`

HasCustomValue4 returns a boolean if a field has been set.

### GetTaxName1

`func (o *RecurringInvoiceRequest) GetTaxName1() string`

GetTaxName1 returns the TaxName1 field if non-nil, zero value otherwise.

### GetTaxName1Ok

`func (o *RecurringInvoiceRequest) GetTaxName1Ok() (*string, bool)`

GetTaxName1Ok returns a tuple with the TaxName1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName1

`func (o *RecurringInvoiceRequest) SetTaxName1(v string)`

SetTaxName1 sets TaxName1 field to given value.

### HasTaxName1

`func (o *RecurringInvoiceRequest) HasTaxName1() bool`

HasTaxName1 returns a boolean if a field has been set.

### GetTaxName2

`func (o *RecurringInvoiceRequest) GetTaxName2() string`

GetTaxName2 returns the TaxName2 field if non-nil, zero value otherwise.

### GetTaxName2Ok

`func (o *RecurringInvoiceRequest) GetTaxName2Ok() (*string, bool)`

GetTaxName2Ok returns a tuple with the TaxName2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName2

`func (o *RecurringInvoiceRequest) SetTaxName2(v string)`

SetTaxName2 sets TaxName2 field to given value.

### HasTaxName2

`func (o *RecurringInvoiceRequest) HasTaxName2() bool`

HasTaxName2 returns a boolean if a field has been set.

### GetTaxRate1

`func (o *RecurringInvoiceRequest) GetTaxRate1() float32`

GetTaxRate1 returns the TaxRate1 field if non-nil, zero value otherwise.

### GetTaxRate1Ok

`func (o *RecurringInvoiceRequest) GetTaxRate1Ok() (*float32, bool)`

GetTaxRate1Ok returns a tuple with the TaxRate1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate1

`func (o *RecurringInvoiceRequest) SetTaxRate1(v float32)`

SetTaxRate1 sets TaxRate1 field to given value.

### HasTaxRate1

`func (o *RecurringInvoiceRequest) HasTaxRate1() bool`

HasTaxRate1 returns a boolean if a field has been set.

### GetTaxRate2

`func (o *RecurringInvoiceRequest) GetTaxRate2() float32`

GetTaxRate2 returns the TaxRate2 field if non-nil, zero value otherwise.

### GetTaxRate2Ok

`func (o *RecurringInvoiceRequest) GetTaxRate2Ok() (*float32, bool)`

GetTaxRate2Ok returns a tuple with the TaxRate2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate2

`func (o *RecurringInvoiceRequest) SetTaxRate2(v float32)`

SetTaxRate2 sets TaxRate2 field to given value.

### HasTaxRate2

`func (o *RecurringInvoiceRequest) HasTaxRate2() bool`

HasTaxRate2 returns a boolean if a field has been set.

### GetTaxName3

`func (o *RecurringInvoiceRequest) GetTaxName3() string`

GetTaxName3 returns the TaxName3 field if non-nil, zero value otherwise.

### GetTaxName3Ok

`func (o *RecurringInvoiceRequest) GetTaxName3Ok() (*string, bool)`

GetTaxName3Ok returns a tuple with the TaxName3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName3

`func (o *RecurringInvoiceRequest) SetTaxName3(v string)`

SetTaxName3 sets TaxName3 field to given value.

### HasTaxName3

`func (o *RecurringInvoiceRequest) HasTaxName3() bool`

HasTaxName3 returns a boolean if a field has been set.

### GetTaxRate3

`func (o *RecurringInvoiceRequest) GetTaxRate3() float32`

GetTaxRate3 returns the TaxRate3 field if non-nil, zero value otherwise.

### GetTaxRate3Ok

`func (o *RecurringInvoiceRequest) GetTaxRate3Ok() (*float32, bool)`

GetTaxRate3Ok returns a tuple with the TaxRate3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate3

`func (o *RecurringInvoiceRequest) SetTaxRate3(v float32)`

SetTaxRate3 sets TaxRate3 field to given value.

### HasTaxRate3

`func (o *RecurringInvoiceRequest) HasTaxRate3() bool`

HasTaxRate3 returns a boolean if a field has been set.

### GetLineItems

`func (o *RecurringInvoiceRequest) GetLineItems() map[string]interface{}`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *RecurringInvoiceRequest) GetLineItemsOk() (*map[string]interface{}, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *RecurringInvoiceRequest) SetLineItems(v map[string]interface{})`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *RecurringInvoiceRequest) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetDiscount

`func (o *RecurringInvoiceRequest) GetDiscount() float32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *RecurringInvoiceRequest) GetDiscountOk() (*float32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *RecurringInvoiceRequest) SetDiscount(v float32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *RecurringInvoiceRequest) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetPartial

`func (o *RecurringInvoiceRequest) GetPartial() float32`

GetPartial returns the Partial field if non-nil, zero value otherwise.

### GetPartialOk

`func (o *RecurringInvoiceRequest) GetPartialOk() (*float32, bool)`

GetPartialOk returns a tuple with the Partial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartial

`func (o *RecurringInvoiceRequest) SetPartial(v float32)`

SetPartial sets Partial field to given value.

### HasPartial

`func (o *RecurringInvoiceRequest) HasPartial() bool`

HasPartial returns a boolean if a field has been set.

### GetIsAmountDiscount

`func (o *RecurringInvoiceRequest) GetIsAmountDiscount() bool`

GetIsAmountDiscount returns the IsAmountDiscount field if non-nil, zero value otherwise.

### GetIsAmountDiscountOk

`func (o *RecurringInvoiceRequest) GetIsAmountDiscountOk() (*bool, bool)`

GetIsAmountDiscountOk returns a tuple with the IsAmountDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAmountDiscount

`func (o *RecurringInvoiceRequest) SetIsAmountDiscount(v bool)`

SetIsAmountDiscount sets IsAmountDiscount field to given value.

### HasIsAmountDiscount

`func (o *RecurringInvoiceRequest) HasIsAmountDiscount() bool`

HasIsAmountDiscount returns a boolean if a field has been set.

### GetUsesInclusiveTaxes

`func (o *RecurringInvoiceRequest) GetUsesInclusiveTaxes() bool`

GetUsesInclusiveTaxes returns the UsesInclusiveTaxes field if non-nil, zero value otherwise.

### GetUsesInclusiveTaxesOk

`func (o *RecurringInvoiceRequest) GetUsesInclusiveTaxesOk() (*bool, bool)`

GetUsesInclusiveTaxesOk returns a tuple with the UsesInclusiveTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesInclusiveTaxes

`func (o *RecurringInvoiceRequest) SetUsesInclusiveTaxes(v bool)`

SetUsesInclusiveTaxes sets UsesInclusiveTaxes field to given value.

### HasUsesInclusiveTaxes

`func (o *RecurringInvoiceRequest) HasUsesInclusiveTaxes() bool`

HasUsesInclusiveTaxes returns a boolean if a field has been set.

### GetDate

`func (o *RecurringInvoiceRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *RecurringInvoiceRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *RecurringInvoiceRequest) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *RecurringInvoiceRequest) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetPartialDueDate

`func (o *RecurringInvoiceRequest) GetPartialDueDate() string`

GetPartialDueDate returns the PartialDueDate field if non-nil, zero value otherwise.

### GetPartialDueDateOk

`func (o *RecurringInvoiceRequest) GetPartialDueDateOk() (*string, bool)`

GetPartialDueDateOk returns a tuple with the PartialDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialDueDate

`func (o *RecurringInvoiceRequest) SetPartialDueDate(v string)`

SetPartialDueDate sets PartialDueDate field to given value.

### HasPartialDueDate

`func (o *RecurringInvoiceRequest) HasPartialDueDate() bool`

HasPartialDueDate returns a boolean if a field has been set.

### GetDueDate

`func (o *RecurringInvoiceRequest) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *RecurringInvoiceRequest) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *RecurringInvoiceRequest) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *RecurringInvoiceRequest) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetCustomSurcharge1

`func (o *RecurringInvoiceRequest) GetCustomSurcharge1() float32`

GetCustomSurcharge1 returns the CustomSurcharge1 field if non-nil, zero value otherwise.

### GetCustomSurcharge1Ok

`func (o *RecurringInvoiceRequest) GetCustomSurcharge1Ok() (*float32, bool)`

GetCustomSurcharge1Ok returns a tuple with the CustomSurcharge1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurcharge1

`func (o *RecurringInvoiceRequest) SetCustomSurcharge1(v float32)`

SetCustomSurcharge1 sets CustomSurcharge1 field to given value.

### HasCustomSurcharge1

`func (o *RecurringInvoiceRequest) HasCustomSurcharge1() bool`

HasCustomSurcharge1 returns a boolean if a field has been set.

### GetCustomSurcharge2

`func (o *RecurringInvoiceRequest) GetCustomSurcharge2() float32`

GetCustomSurcharge2 returns the CustomSurcharge2 field if non-nil, zero value otherwise.

### GetCustomSurcharge2Ok

`func (o *RecurringInvoiceRequest) GetCustomSurcharge2Ok() (*float32, bool)`

GetCustomSurcharge2Ok returns a tuple with the CustomSurcharge2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurcharge2

`func (o *RecurringInvoiceRequest) SetCustomSurcharge2(v float32)`

SetCustomSurcharge2 sets CustomSurcharge2 field to given value.

### HasCustomSurcharge2

`func (o *RecurringInvoiceRequest) HasCustomSurcharge2() bool`

HasCustomSurcharge2 returns a boolean if a field has been set.

### GetCustomSurcharge3

`func (o *RecurringInvoiceRequest) GetCustomSurcharge3() float32`

GetCustomSurcharge3 returns the CustomSurcharge3 field if non-nil, zero value otherwise.

### GetCustomSurcharge3Ok

`func (o *RecurringInvoiceRequest) GetCustomSurcharge3Ok() (*float32, bool)`

GetCustomSurcharge3Ok returns a tuple with the CustomSurcharge3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurcharge3

`func (o *RecurringInvoiceRequest) SetCustomSurcharge3(v float32)`

SetCustomSurcharge3 sets CustomSurcharge3 field to given value.

### HasCustomSurcharge3

`func (o *RecurringInvoiceRequest) HasCustomSurcharge3() bool`

HasCustomSurcharge3 returns a boolean if a field has been set.

### GetCustomSurcharge4

`func (o *RecurringInvoiceRequest) GetCustomSurcharge4() float32`

GetCustomSurcharge4 returns the CustomSurcharge4 field if non-nil, zero value otherwise.

### GetCustomSurcharge4Ok

`func (o *RecurringInvoiceRequest) GetCustomSurcharge4Ok() (*float32, bool)`

GetCustomSurcharge4Ok returns a tuple with the CustomSurcharge4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurcharge4

`func (o *RecurringInvoiceRequest) SetCustomSurcharge4(v float32)`

SetCustomSurcharge4 sets CustomSurcharge4 field to given value.

### HasCustomSurcharge4

`func (o *RecurringInvoiceRequest) HasCustomSurcharge4() bool`

HasCustomSurcharge4 returns a boolean if a field has been set.

### GetCustomSurchargeTax1

`func (o *RecurringInvoiceRequest) GetCustomSurchargeTax1() bool`

GetCustomSurchargeTax1 returns the CustomSurchargeTax1 field if non-nil, zero value otherwise.

### GetCustomSurchargeTax1Ok

`func (o *RecurringInvoiceRequest) GetCustomSurchargeTax1Ok() (*bool, bool)`

GetCustomSurchargeTax1Ok returns a tuple with the CustomSurchargeTax1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurchargeTax1

`func (o *RecurringInvoiceRequest) SetCustomSurchargeTax1(v bool)`

SetCustomSurchargeTax1 sets CustomSurchargeTax1 field to given value.

### HasCustomSurchargeTax1

`func (o *RecurringInvoiceRequest) HasCustomSurchargeTax1() bool`

HasCustomSurchargeTax1 returns a boolean if a field has been set.

### GetCustomSurchargeTax2

`func (o *RecurringInvoiceRequest) GetCustomSurchargeTax2() bool`

GetCustomSurchargeTax2 returns the CustomSurchargeTax2 field if non-nil, zero value otherwise.

### GetCustomSurchargeTax2Ok

`func (o *RecurringInvoiceRequest) GetCustomSurchargeTax2Ok() (*bool, bool)`

GetCustomSurchargeTax2Ok returns a tuple with the CustomSurchargeTax2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurchargeTax2

`func (o *RecurringInvoiceRequest) SetCustomSurchargeTax2(v bool)`

SetCustomSurchargeTax2 sets CustomSurchargeTax2 field to given value.

### HasCustomSurchargeTax2

`func (o *RecurringInvoiceRequest) HasCustomSurchargeTax2() bool`

HasCustomSurchargeTax2 returns a boolean if a field has been set.

### GetCustomSurchargeTax3

`func (o *RecurringInvoiceRequest) GetCustomSurchargeTax3() bool`

GetCustomSurchargeTax3 returns the CustomSurchargeTax3 field if non-nil, zero value otherwise.

### GetCustomSurchargeTax3Ok

`func (o *RecurringInvoiceRequest) GetCustomSurchargeTax3Ok() (*bool, bool)`

GetCustomSurchargeTax3Ok returns a tuple with the CustomSurchargeTax3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurchargeTax3

`func (o *RecurringInvoiceRequest) SetCustomSurchargeTax3(v bool)`

SetCustomSurchargeTax3 sets CustomSurchargeTax3 field to given value.

### HasCustomSurchargeTax3

`func (o *RecurringInvoiceRequest) HasCustomSurchargeTax3() bool`

HasCustomSurchargeTax3 returns a boolean if a field has been set.

### GetCustomSurchargeTax4

`func (o *RecurringInvoiceRequest) GetCustomSurchargeTax4() bool`

GetCustomSurchargeTax4 returns the CustomSurchargeTax4 field if non-nil, zero value otherwise.

### GetCustomSurchargeTax4Ok

`func (o *RecurringInvoiceRequest) GetCustomSurchargeTax4Ok() (*bool, bool)`

GetCustomSurchargeTax4Ok returns a tuple with the CustomSurchargeTax4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurchargeTax4

`func (o *RecurringInvoiceRequest) SetCustomSurchargeTax4(v bool)`

SetCustomSurchargeTax4 sets CustomSurchargeTax4 field to given value.

### HasCustomSurchargeTax4

`func (o *RecurringInvoiceRequest) HasCustomSurchargeTax4() bool`

HasCustomSurchargeTax4 returns a boolean if a field has been set.

### GetLocationId

`func (o *RecurringInvoiceRequest) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *RecurringInvoiceRequest) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *RecurringInvoiceRequest) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *RecurringInvoiceRequest) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


