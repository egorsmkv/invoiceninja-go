# PurchaseOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique hashed identifier for the purchase order | [optional] 
**UserId** | Pointer to **string** | The unique hashed identifier for the user who created the purchase order | [optional] 
**AssignedUserId** | Pointer to **string** | The unique hashed identifier for the user assigned to the purchase order | [optional] 
**VendorId** | Pointer to **string** | The unique hashed identifier for the vendor associated with the purchase order | [optional] 
**StatusId** | Pointer to **string** | The status of the purchase order represented by a unique identifier | [optional] 
**Number** | Pointer to **string** | The unique alpha-numeric purchase order number per company | [optional] 
**QuoteNumber** | Pointer to **string** | The quote number associated with this purchase order | [optional] 
**Terms** | Pointer to **string** | The terms and conditions for the purchase order | [optional] 
**PublicNotes** | Pointer to **string** | Publicly visible notes associated with the purchase order | [optional] 
**PrivateNotes** | Pointer to **string** | Privately visible notes associated with the purchase order, not disclosed to the vendor | [optional] 
**Footer** | Pointer to **string** | The footer text of the purchase order | [optional] 
**CustomValue1** | Pointer to **string** | First custom value field for additional information | [optional] 
**CustomValue2** | Pointer to **string** | Second custom value field for additional information | [optional] 
**CustomValue3** | Pointer to **string** | Third custom value field for additional information | [optional] 
**CustomValue4** | Pointer to **string** | Fourth custom value field for additional information | [optional] 
**TaxName1** | Pointer to **string** | The name of the first tax applied to the purchase order | [optional] 
**TaxName2** | Pointer to **string** | The name of the second tax applied to the purchase order | [optional] 
**TaxRate1** | Pointer to **float32** | The rate of the first tax applied to the purchase order | [optional] 
**TaxRate2** | Pointer to **float32** | The rate of the second tax applied to the purchase order | [optional] 
**TaxName3** | Pointer to **string** | The name of the third tax applied to the purchase order | [optional] 
**TaxRate3** | Pointer to **float32** | The rate of the third tax applied to the purchase order | [optional] 
**TotalTaxes** | Pointer to **float32** | The total amount of taxes applied to the purchase order | [optional] 
**LineItems** | Pointer to [**[]InvoiceItem**](InvoiceItem.md) | An array of objects which define the line items of the purchase order | [optional] 
**Amount** | Pointer to **float32** | The total amount of the purchase order before taxes and discounts | [optional] 
**Balance** | Pointer to **float32** | The balance due for the purchase order after accounting for payments | [optional] 
**PaidToDate** | Pointer to **float32** | The total amount paid on the purchase order so far | [optional] 
**Discount** | Pointer to **float32** | The discount amount or percentage applied to the purchase order | [optional] 
**Partial** | Pointer to **float32** | The partial or deposit amount for the purchase order | [optional] 
**IsAmountDiscount** | Pointer to **bool** | Boolean flag indicating if the discount is a fixed amount or a percentage | [optional] 
**IsDeleted** | Pointer to **bool** | Boolean flag indicating if the purchase order has been deleted | [optional] 
**UsesInclusiveTaxes** | Pointer to **bool** | Boolean flag indicating if the taxes used are inclusive or exclusive | [optional] 
**Date** | Pointer to **string** | The date the purchase order was created | [optional] 
**LastSentDate** | Pointer to **string** | The last date the purchase order was sent to the vendor | [optional] 
**NextSendDate** | Pointer to **string** | The next scheduled date for sending a reminder for the purchase order | [optional] 
**PartialDueDate** | Pointer to **string** | The due date for the partial or deposit amount | [optional] 
**DueDate** | Pointer to **string** | The due date for the total amount of the purchase order | [optional] 
**Settings** | Pointer to [**CompanySettings**](CompanySettings.md) |  | [optional] 
**LastViewed** | Pointer to **float32** | Timestamp | [optional] 
**UpdatedAt** | Pointer to **float32** | Timestamp | [optional] 
**ArchivedAt** | Pointer to **float32** | Timestamp | [optional] 
**CustomSurcharge1** | Pointer to **float32** | First custom surcharge amount for the purchase order | [optional] 
**CustomSurcharge2** | Pointer to **float32** | Second custom surcharge amount for the purchase order | [optional] 
**CustomSurcharge3** | Pointer to **float32** | Third custom surcharge amount for the purchase order | [optional] 
**CustomSurcharge4** | Pointer to **float32** | Fourth custom surcharge amount for the purchase order | [optional] 
**CustomSurchargeTax1** | Pointer to **bool** | Boolean flag indicating if taxes are charged on the first custom surcharge amount | [optional] 
**CustomSurchargeTax2** | Pointer to **bool** | Boolean flag indicating if taxes are charged on the second custom surcharge amount | [optional] 
**CustomSurchargeTax3** | Pointer to **bool** | Boolean flag indicating if taxes are charged on the third custom surcharge amount | [optional] 
**CustomSurchargeTax4** | Pointer to **bool** | Boolean flag indicating if taxes are charged on the fourth custom surcharge amount | [optional] 

## Methods

### NewPurchaseOrder

`func NewPurchaseOrder() *PurchaseOrder`

NewPurchaseOrder instantiates a new PurchaseOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseOrderWithDefaults

`func NewPurchaseOrderWithDefaults() *PurchaseOrder`

NewPurchaseOrderWithDefaults instantiates a new PurchaseOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PurchaseOrder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PurchaseOrder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PurchaseOrder) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PurchaseOrder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *PurchaseOrder) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PurchaseOrder) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PurchaseOrder) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PurchaseOrder) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAssignedUserId

`func (o *PurchaseOrder) GetAssignedUserId() string`

GetAssignedUserId returns the AssignedUserId field if non-nil, zero value otherwise.

### GetAssignedUserIdOk

`func (o *PurchaseOrder) GetAssignedUserIdOk() (*string, bool)`

GetAssignedUserIdOk returns a tuple with the AssignedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUserId

`func (o *PurchaseOrder) SetAssignedUserId(v string)`

SetAssignedUserId sets AssignedUserId field to given value.

### HasAssignedUserId

`func (o *PurchaseOrder) HasAssignedUserId() bool`

HasAssignedUserId returns a boolean if a field has been set.

### GetVendorId

`func (o *PurchaseOrder) GetVendorId() string`

GetVendorId returns the VendorId field if non-nil, zero value otherwise.

### GetVendorIdOk

`func (o *PurchaseOrder) GetVendorIdOk() (*string, bool)`

GetVendorIdOk returns a tuple with the VendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorId

`func (o *PurchaseOrder) SetVendorId(v string)`

SetVendorId sets VendorId field to given value.

### HasVendorId

`func (o *PurchaseOrder) HasVendorId() bool`

HasVendorId returns a boolean if a field has been set.

### GetStatusId

`func (o *PurchaseOrder) GetStatusId() string`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *PurchaseOrder) GetStatusIdOk() (*string, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *PurchaseOrder) SetStatusId(v string)`

SetStatusId sets StatusId field to given value.

### HasStatusId

`func (o *PurchaseOrder) HasStatusId() bool`

HasStatusId returns a boolean if a field has been set.

### GetNumber

`func (o *PurchaseOrder) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PurchaseOrder) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PurchaseOrder) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *PurchaseOrder) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetQuoteNumber

`func (o *PurchaseOrder) GetQuoteNumber() string`

GetQuoteNumber returns the QuoteNumber field if non-nil, zero value otherwise.

### GetQuoteNumberOk

`func (o *PurchaseOrder) GetQuoteNumberOk() (*string, bool)`

GetQuoteNumberOk returns a tuple with the QuoteNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteNumber

`func (o *PurchaseOrder) SetQuoteNumber(v string)`

SetQuoteNumber sets QuoteNumber field to given value.

### HasQuoteNumber

`func (o *PurchaseOrder) HasQuoteNumber() bool`

HasQuoteNumber returns a boolean if a field has been set.

### GetTerms

`func (o *PurchaseOrder) GetTerms() string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *PurchaseOrder) GetTermsOk() (*string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *PurchaseOrder) SetTerms(v string)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *PurchaseOrder) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### GetPublicNotes

`func (o *PurchaseOrder) GetPublicNotes() string`

GetPublicNotes returns the PublicNotes field if non-nil, zero value otherwise.

### GetPublicNotesOk

`func (o *PurchaseOrder) GetPublicNotesOk() (*string, bool)`

GetPublicNotesOk returns a tuple with the PublicNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNotes

`func (o *PurchaseOrder) SetPublicNotes(v string)`

SetPublicNotes sets PublicNotes field to given value.

### HasPublicNotes

`func (o *PurchaseOrder) HasPublicNotes() bool`

HasPublicNotes returns a boolean if a field has been set.

### GetPrivateNotes

`func (o *PurchaseOrder) GetPrivateNotes() string`

GetPrivateNotes returns the PrivateNotes field if non-nil, zero value otherwise.

### GetPrivateNotesOk

`func (o *PurchaseOrder) GetPrivateNotesOk() (*string, bool)`

GetPrivateNotesOk returns a tuple with the PrivateNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNotes

`func (o *PurchaseOrder) SetPrivateNotes(v string)`

SetPrivateNotes sets PrivateNotes field to given value.

### HasPrivateNotes

`func (o *PurchaseOrder) HasPrivateNotes() bool`

HasPrivateNotes returns a boolean if a field has been set.

### GetFooter

`func (o *PurchaseOrder) GetFooter() string`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *PurchaseOrder) GetFooterOk() (*string, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *PurchaseOrder) SetFooter(v string)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *PurchaseOrder) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetCustomValue1

`func (o *PurchaseOrder) GetCustomValue1() string`

GetCustomValue1 returns the CustomValue1 field if non-nil, zero value otherwise.

### GetCustomValue1Ok

`func (o *PurchaseOrder) GetCustomValue1Ok() (*string, bool)`

GetCustomValue1Ok returns a tuple with the CustomValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue1

`func (o *PurchaseOrder) SetCustomValue1(v string)`

SetCustomValue1 sets CustomValue1 field to given value.

### HasCustomValue1

`func (o *PurchaseOrder) HasCustomValue1() bool`

HasCustomValue1 returns a boolean if a field has been set.

### GetCustomValue2

`func (o *PurchaseOrder) GetCustomValue2() string`

GetCustomValue2 returns the CustomValue2 field if non-nil, zero value otherwise.

### GetCustomValue2Ok

`func (o *PurchaseOrder) GetCustomValue2Ok() (*string, bool)`

GetCustomValue2Ok returns a tuple with the CustomValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue2

`func (o *PurchaseOrder) SetCustomValue2(v string)`

SetCustomValue2 sets CustomValue2 field to given value.

### HasCustomValue2

`func (o *PurchaseOrder) HasCustomValue2() bool`

HasCustomValue2 returns a boolean if a field has been set.

### GetCustomValue3

`func (o *PurchaseOrder) GetCustomValue3() string`

GetCustomValue3 returns the CustomValue3 field if non-nil, zero value otherwise.

### GetCustomValue3Ok

`func (o *PurchaseOrder) GetCustomValue3Ok() (*string, bool)`

GetCustomValue3Ok returns a tuple with the CustomValue3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue3

`func (o *PurchaseOrder) SetCustomValue3(v string)`

SetCustomValue3 sets CustomValue3 field to given value.

### HasCustomValue3

`func (o *PurchaseOrder) HasCustomValue3() bool`

HasCustomValue3 returns a boolean if a field has been set.

### GetCustomValue4

`func (o *PurchaseOrder) GetCustomValue4() string`

GetCustomValue4 returns the CustomValue4 field if non-nil, zero value otherwise.

### GetCustomValue4Ok

`func (o *PurchaseOrder) GetCustomValue4Ok() (*string, bool)`

GetCustomValue4Ok returns a tuple with the CustomValue4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue4

`func (o *PurchaseOrder) SetCustomValue4(v string)`

SetCustomValue4 sets CustomValue4 field to given value.

### HasCustomValue4

`func (o *PurchaseOrder) HasCustomValue4() bool`

HasCustomValue4 returns a boolean if a field has been set.

### GetTaxName1

`func (o *PurchaseOrder) GetTaxName1() string`

GetTaxName1 returns the TaxName1 field if non-nil, zero value otherwise.

### GetTaxName1Ok

`func (o *PurchaseOrder) GetTaxName1Ok() (*string, bool)`

GetTaxName1Ok returns a tuple with the TaxName1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName1

`func (o *PurchaseOrder) SetTaxName1(v string)`

SetTaxName1 sets TaxName1 field to given value.

### HasTaxName1

`func (o *PurchaseOrder) HasTaxName1() bool`

HasTaxName1 returns a boolean if a field has been set.

### GetTaxName2

`func (o *PurchaseOrder) GetTaxName2() string`

GetTaxName2 returns the TaxName2 field if non-nil, zero value otherwise.

### GetTaxName2Ok

`func (o *PurchaseOrder) GetTaxName2Ok() (*string, bool)`

GetTaxName2Ok returns a tuple with the TaxName2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName2

`func (o *PurchaseOrder) SetTaxName2(v string)`

SetTaxName2 sets TaxName2 field to given value.

### HasTaxName2

`func (o *PurchaseOrder) HasTaxName2() bool`

HasTaxName2 returns a boolean if a field has been set.

### GetTaxRate1

`func (o *PurchaseOrder) GetTaxRate1() float32`

GetTaxRate1 returns the TaxRate1 field if non-nil, zero value otherwise.

### GetTaxRate1Ok

`func (o *PurchaseOrder) GetTaxRate1Ok() (*float32, bool)`

GetTaxRate1Ok returns a tuple with the TaxRate1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate1

`func (o *PurchaseOrder) SetTaxRate1(v float32)`

SetTaxRate1 sets TaxRate1 field to given value.

### HasTaxRate1

`func (o *PurchaseOrder) HasTaxRate1() bool`

HasTaxRate1 returns a boolean if a field has been set.

### GetTaxRate2

`func (o *PurchaseOrder) GetTaxRate2() float32`

GetTaxRate2 returns the TaxRate2 field if non-nil, zero value otherwise.

### GetTaxRate2Ok

`func (o *PurchaseOrder) GetTaxRate2Ok() (*float32, bool)`

GetTaxRate2Ok returns a tuple with the TaxRate2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate2

`func (o *PurchaseOrder) SetTaxRate2(v float32)`

SetTaxRate2 sets TaxRate2 field to given value.

### HasTaxRate2

`func (o *PurchaseOrder) HasTaxRate2() bool`

HasTaxRate2 returns a boolean if a field has been set.

### GetTaxName3

`func (o *PurchaseOrder) GetTaxName3() string`

GetTaxName3 returns the TaxName3 field if non-nil, zero value otherwise.

### GetTaxName3Ok

`func (o *PurchaseOrder) GetTaxName3Ok() (*string, bool)`

GetTaxName3Ok returns a tuple with the TaxName3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName3

`func (o *PurchaseOrder) SetTaxName3(v string)`

SetTaxName3 sets TaxName3 field to given value.

### HasTaxName3

`func (o *PurchaseOrder) HasTaxName3() bool`

HasTaxName3 returns a boolean if a field has been set.

### GetTaxRate3

`func (o *PurchaseOrder) GetTaxRate3() float32`

GetTaxRate3 returns the TaxRate3 field if non-nil, zero value otherwise.

### GetTaxRate3Ok

`func (o *PurchaseOrder) GetTaxRate3Ok() (*float32, bool)`

GetTaxRate3Ok returns a tuple with the TaxRate3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate3

`func (o *PurchaseOrder) SetTaxRate3(v float32)`

SetTaxRate3 sets TaxRate3 field to given value.

### HasTaxRate3

`func (o *PurchaseOrder) HasTaxRate3() bool`

HasTaxRate3 returns a boolean if a field has been set.

### GetTotalTaxes

`func (o *PurchaseOrder) GetTotalTaxes() float32`

GetTotalTaxes returns the TotalTaxes field if non-nil, zero value otherwise.

### GetTotalTaxesOk

`func (o *PurchaseOrder) GetTotalTaxesOk() (*float32, bool)`

GetTotalTaxesOk returns a tuple with the TotalTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxes

`func (o *PurchaseOrder) SetTotalTaxes(v float32)`

SetTotalTaxes sets TotalTaxes field to given value.

### HasTotalTaxes

`func (o *PurchaseOrder) HasTotalTaxes() bool`

HasTotalTaxes returns a boolean if a field has been set.

### GetLineItems

`func (o *PurchaseOrder) GetLineItems() []InvoiceItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *PurchaseOrder) GetLineItemsOk() (*[]InvoiceItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *PurchaseOrder) SetLineItems(v []InvoiceItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *PurchaseOrder) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetAmount

`func (o *PurchaseOrder) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PurchaseOrder) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PurchaseOrder) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PurchaseOrder) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBalance

`func (o *PurchaseOrder) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *PurchaseOrder) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *PurchaseOrder) SetBalance(v float32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *PurchaseOrder) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetPaidToDate

`func (o *PurchaseOrder) GetPaidToDate() float32`

GetPaidToDate returns the PaidToDate field if non-nil, zero value otherwise.

### GetPaidToDateOk

`func (o *PurchaseOrder) GetPaidToDateOk() (*float32, bool)`

GetPaidToDateOk returns a tuple with the PaidToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidToDate

`func (o *PurchaseOrder) SetPaidToDate(v float32)`

SetPaidToDate sets PaidToDate field to given value.

### HasPaidToDate

`func (o *PurchaseOrder) HasPaidToDate() bool`

HasPaidToDate returns a boolean if a field has been set.

### GetDiscount

`func (o *PurchaseOrder) GetDiscount() float32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *PurchaseOrder) GetDiscountOk() (*float32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *PurchaseOrder) SetDiscount(v float32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *PurchaseOrder) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetPartial

`func (o *PurchaseOrder) GetPartial() float32`

GetPartial returns the Partial field if non-nil, zero value otherwise.

### GetPartialOk

`func (o *PurchaseOrder) GetPartialOk() (*float32, bool)`

GetPartialOk returns a tuple with the Partial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartial

`func (o *PurchaseOrder) SetPartial(v float32)`

SetPartial sets Partial field to given value.

### HasPartial

`func (o *PurchaseOrder) HasPartial() bool`

HasPartial returns a boolean if a field has been set.

### GetIsAmountDiscount

`func (o *PurchaseOrder) GetIsAmountDiscount() bool`

GetIsAmountDiscount returns the IsAmountDiscount field if non-nil, zero value otherwise.

### GetIsAmountDiscountOk

`func (o *PurchaseOrder) GetIsAmountDiscountOk() (*bool, bool)`

GetIsAmountDiscountOk returns a tuple with the IsAmountDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAmountDiscount

`func (o *PurchaseOrder) SetIsAmountDiscount(v bool)`

SetIsAmountDiscount sets IsAmountDiscount field to given value.

### HasIsAmountDiscount

`func (o *PurchaseOrder) HasIsAmountDiscount() bool`

HasIsAmountDiscount returns a boolean if a field has been set.

### GetIsDeleted

`func (o *PurchaseOrder) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *PurchaseOrder) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *PurchaseOrder) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *PurchaseOrder) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetUsesInclusiveTaxes

`func (o *PurchaseOrder) GetUsesInclusiveTaxes() bool`

GetUsesInclusiveTaxes returns the UsesInclusiveTaxes field if non-nil, zero value otherwise.

### GetUsesInclusiveTaxesOk

`func (o *PurchaseOrder) GetUsesInclusiveTaxesOk() (*bool, bool)`

GetUsesInclusiveTaxesOk returns a tuple with the UsesInclusiveTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesInclusiveTaxes

`func (o *PurchaseOrder) SetUsesInclusiveTaxes(v bool)`

SetUsesInclusiveTaxes sets UsesInclusiveTaxes field to given value.

### HasUsesInclusiveTaxes

`func (o *PurchaseOrder) HasUsesInclusiveTaxes() bool`

HasUsesInclusiveTaxes returns a boolean if a field has been set.

### GetDate

`func (o *PurchaseOrder) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PurchaseOrder) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PurchaseOrder) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *PurchaseOrder) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetLastSentDate

`func (o *PurchaseOrder) GetLastSentDate() string`

GetLastSentDate returns the LastSentDate field if non-nil, zero value otherwise.

### GetLastSentDateOk

`func (o *PurchaseOrder) GetLastSentDateOk() (*string, bool)`

GetLastSentDateOk returns a tuple with the LastSentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSentDate

`func (o *PurchaseOrder) SetLastSentDate(v string)`

SetLastSentDate sets LastSentDate field to given value.

### HasLastSentDate

`func (o *PurchaseOrder) HasLastSentDate() bool`

HasLastSentDate returns a boolean if a field has been set.

### GetNextSendDate

`func (o *PurchaseOrder) GetNextSendDate() string`

GetNextSendDate returns the NextSendDate field if non-nil, zero value otherwise.

### GetNextSendDateOk

`func (o *PurchaseOrder) GetNextSendDateOk() (*string, bool)`

GetNextSendDateOk returns a tuple with the NextSendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSendDate

`func (o *PurchaseOrder) SetNextSendDate(v string)`

SetNextSendDate sets NextSendDate field to given value.

### HasNextSendDate

`func (o *PurchaseOrder) HasNextSendDate() bool`

HasNextSendDate returns a boolean if a field has been set.

### GetPartialDueDate

`func (o *PurchaseOrder) GetPartialDueDate() string`

GetPartialDueDate returns the PartialDueDate field if non-nil, zero value otherwise.

### GetPartialDueDateOk

`func (o *PurchaseOrder) GetPartialDueDateOk() (*string, bool)`

GetPartialDueDateOk returns a tuple with the PartialDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialDueDate

`func (o *PurchaseOrder) SetPartialDueDate(v string)`

SetPartialDueDate sets PartialDueDate field to given value.

### HasPartialDueDate

`func (o *PurchaseOrder) HasPartialDueDate() bool`

HasPartialDueDate returns a boolean if a field has been set.

### GetDueDate

`func (o *PurchaseOrder) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *PurchaseOrder) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *PurchaseOrder) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *PurchaseOrder) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetSettings

`func (o *PurchaseOrder) GetSettings() CompanySettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *PurchaseOrder) GetSettingsOk() (*CompanySettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *PurchaseOrder) SetSettings(v CompanySettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *PurchaseOrder) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetLastViewed

`func (o *PurchaseOrder) GetLastViewed() float32`

GetLastViewed returns the LastViewed field if non-nil, zero value otherwise.

### GetLastViewedOk

`func (o *PurchaseOrder) GetLastViewedOk() (*float32, bool)`

GetLastViewedOk returns a tuple with the LastViewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastViewed

`func (o *PurchaseOrder) SetLastViewed(v float32)`

SetLastViewed sets LastViewed field to given value.

### HasLastViewed

`func (o *PurchaseOrder) HasLastViewed() bool`

HasLastViewed returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PurchaseOrder) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PurchaseOrder) GetUpdatedAtOk() (*float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PurchaseOrder) SetUpdatedAt(v float32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PurchaseOrder) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetArchivedAt

`func (o *PurchaseOrder) GetArchivedAt() float32`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *PurchaseOrder) GetArchivedAtOk() (*float32, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *PurchaseOrder) SetArchivedAt(v float32)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *PurchaseOrder) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetCustomSurcharge1

`func (o *PurchaseOrder) GetCustomSurcharge1() float32`

GetCustomSurcharge1 returns the CustomSurcharge1 field if non-nil, zero value otherwise.

### GetCustomSurcharge1Ok

`func (o *PurchaseOrder) GetCustomSurcharge1Ok() (*float32, bool)`

GetCustomSurcharge1Ok returns a tuple with the CustomSurcharge1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurcharge1

`func (o *PurchaseOrder) SetCustomSurcharge1(v float32)`

SetCustomSurcharge1 sets CustomSurcharge1 field to given value.

### HasCustomSurcharge1

`func (o *PurchaseOrder) HasCustomSurcharge1() bool`

HasCustomSurcharge1 returns a boolean if a field has been set.

### GetCustomSurcharge2

`func (o *PurchaseOrder) GetCustomSurcharge2() float32`

GetCustomSurcharge2 returns the CustomSurcharge2 field if non-nil, zero value otherwise.

### GetCustomSurcharge2Ok

`func (o *PurchaseOrder) GetCustomSurcharge2Ok() (*float32, bool)`

GetCustomSurcharge2Ok returns a tuple with the CustomSurcharge2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurcharge2

`func (o *PurchaseOrder) SetCustomSurcharge2(v float32)`

SetCustomSurcharge2 sets CustomSurcharge2 field to given value.

### HasCustomSurcharge2

`func (o *PurchaseOrder) HasCustomSurcharge2() bool`

HasCustomSurcharge2 returns a boolean if a field has been set.

### GetCustomSurcharge3

`func (o *PurchaseOrder) GetCustomSurcharge3() float32`

GetCustomSurcharge3 returns the CustomSurcharge3 field if non-nil, zero value otherwise.

### GetCustomSurcharge3Ok

`func (o *PurchaseOrder) GetCustomSurcharge3Ok() (*float32, bool)`

GetCustomSurcharge3Ok returns a tuple with the CustomSurcharge3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurcharge3

`func (o *PurchaseOrder) SetCustomSurcharge3(v float32)`

SetCustomSurcharge3 sets CustomSurcharge3 field to given value.

### HasCustomSurcharge3

`func (o *PurchaseOrder) HasCustomSurcharge3() bool`

HasCustomSurcharge3 returns a boolean if a field has been set.

### GetCustomSurcharge4

`func (o *PurchaseOrder) GetCustomSurcharge4() float32`

GetCustomSurcharge4 returns the CustomSurcharge4 field if non-nil, zero value otherwise.

### GetCustomSurcharge4Ok

`func (o *PurchaseOrder) GetCustomSurcharge4Ok() (*float32, bool)`

GetCustomSurcharge4Ok returns a tuple with the CustomSurcharge4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurcharge4

`func (o *PurchaseOrder) SetCustomSurcharge4(v float32)`

SetCustomSurcharge4 sets CustomSurcharge4 field to given value.

### HasCustomSurcharge4

`func (o *PurchaseOrder) HasCustomSurcharge4() bool`

HasCustomSurcharge4 returns a boolean if a field has been set.

### GetCustomSurchargeTax1

`func (o *PurchaseOrder) GetCustomSurchargeTax1() bool`

GetCustomSurchargeTax1 returns the CustomSurchargeTax1 field if non-nil, zero value otherwise.

### GetCustomSurchargeTax1Ok

`func (o *PurchaseOrder) GetCustomSurchargeTax1Ok() (*bool, bool)`

GetCustomSurchargeTax1Ok returns a tuple with the CustomSurchargeTax1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurchargeTax1

`func (o *PurchaseOrder) SetCustomSurchargeTax1(v bool)`

SetCustomSurchargeTax1 sets CustomSurchargeTax1 field to given value.

### HasCustomSurchargeTax1

`func (o *PurchaseOrder) HasCustomSurchargeTax1() bool`

HasCustomSurchargeTax1 returns a boolean if a field has been set.

### GetCustomSurchargeTax2

`func (o *PurchaseOrder) GetCustomSurchargeTax2() bool`

GetCustomSurchargeTax2 returns the CustomSurchargeTax2 field if non-nil, zero value otherwise.

### GetCustomSurchargeTax2Ok

`func (o *PurchaseOrder) GetCustomSurchargeTax2Ok() (*bool, bool)`

GetCustomSurchargeTax2Ok returns a tuple with the CustomSurchargeTax2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurchargeTax2

`func (o *PurchaseOrder) SetCustomSurchargeTax2(v bool)`

SetCustomSurchargeTax2 sets CustomSurchargeTax2 field to given value.

### HasCustomSurchargeTax2

`func (o *PurchaseOrder) HasCustomSurchargeTax2() bool`

HasCustomSurchargeTax2 returns a boolean if a field has been set.

### GetCustomSurchargeTax3

`func (o *PurchaseOrder) GetCustomSurchargeTax3() bool`

GetCustomSurchargeTax3 returns the CustomSurchargeTax3 field if non-nil, zero value otherwise.

### GetCustomSurchargeTax3Ok

`func (o *PurchaseOrder) GetCustomSurchargeTax3Ok() (*bool, bool)`

GetCustomSurchargeTax3Ok returns a tuple with the CustomSurchargeTax3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurchargeTax3

`func (o *PurchaseOrder) SetCustomSurchargeTax3(v bool)`

SetCustomSurchargeTax3 sets CustomSurchargeTax3 field to given value.

### HasCustomSurchargeTax3

`func (o *PurchaseOrder) HasCustomSurchargeTax3() bool`

HasCustomSurchargeTax3 returns a boolean if a field has been set.

### GetCustomSurchargeTax4

`func (o *PurchaseOrder) GetCustomSurchargeTax4() bool`

GetCustomSurchargeTax4 returns the CustomSurchargeTax4 field if non-nil, zero value otherwise.

### GetCustomSurchargeTax4Ok

`func (o *PurchaseOrder) GetCustomSurchargeTax4Ok() (*bool, bool)`

GetCustomSurchargeTax4Ok returns a tuple with the CustomSurchargeTax4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurchargeTax4

`func (o *PurchaseOrder) SetCustomSurchargeTax4(v bool)`

SetCustomSurchargeTax4 sets CustomSurchargeTax4 field to given value.

### HasCustomSurchargeTax4

`func (o *PurchaseOrder) HasCustomSurchargeTax4() bool`

HasCustomSurchargeTax4 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


