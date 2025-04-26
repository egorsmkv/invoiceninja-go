# Credit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique hashed ID of the credit | [optional] 
**UserId** | Pointer to **string** | The unique hashed ID of the user associated with the credit | [optional] 
**AssignedUserId** | Pointer to **string** | The unique hashed ID of the assigned user responsible for the credit | [optional] 
**ClientId** | Pointer to **string** | The unique hashed ID of the client associated with the credit | [optional] 
**StatusId** | Pointer to **string** | The ID representing the current status of the credit | [optional] 
**InvoiceId** | Pointer to **string** | The unique hashed ID of the linked invoice to which the credit is applied | [optional] 
**Number** | Pointer to **string** | The unique alphanumeric credit number per company | [optional] 
**PoNumber** | Pointer to **string** | The purchase order number referred to by the credit | [optional] 
**Terms** | Pointer to **string** | The terms associated with the credit | [optional] 
**PublicNotes** | Pointer to **string** | Public notes for the credit | [optional] 
**PrivateNotes** | Pointer to **string** | Private notes for internal use, not visible to the client | [optional] 
**Footer** | Pointer to **string** | The footer text for the credit | [optional] 
**CustomValue1** | Pointer to **string** | Custom value 1 for additional credit information | [optional] 
**CustomValue2** | Pointer to **string** | Custom value 2 for additional credit information | [optional] 
**CustomValue3** | Pointer to **string** | Custom value 3 for additional credit information | [optional] 
**CustomValue4** | Pointer to **string** | Custom value 4 for additional credit information | [optional] 
**TaxName1** | Pointer to **string** | The name of the first tax applied to the credit | [optional] 
**TaxName2** | Pointer to **string** | The name of the second tax applied to the credit | [optional] 
**TaxRate1** | Pointer to **float32** | The rate of the first tax applied to the credit | [optional] 
**TaxRate2** | Pointer to **float32** | The rate of the second tax applied to the credit | [optional] 
**TaxName3** | Pointer to **string** | The name of the third tax applied to the credit | [optional] 
**TaxRate3** | Pointer to **float32** | The rate of the third tax applied to the credit | [optional] 
**TotalTaxes** | Pointer to **float32** | The total amount of taxes for the credit | [optional] 
**LineItems** | Pointer to [**[]InvoiceItem**](InvoiceItem.md) | An array of objects which define the line items of the credit | [optional] 
**Amount** | Pointer to **float32** | The total amount of the credit | [optional] 
**Balance** | Pointer to **float32** | The outstanding balance of the credit | [optional] 
**PaidToDate** | Pointer to **float32** | The total amount paid to date for the credit | [optional] 
**Discount** | Pointer to **float32** | The discount applied to the credit | [optional] 
**Partial** | Pointer to **float32** | The partial amount applied to the credit | [optional] 
**IsAmountDiscount** | Pointer to **bool** | Indicates whether the discount applied is a fixed amount or a percentage | [optional] 
**IsDeleted** | Pointer to **bool** | Indicates whether the credit has been deleted | [optional] 
**UsesInclusiveTaxes** | Pointer to **bool** | Indicates whether the tax rates applied to the credit are inclusive or exclusive | [optional] 
**Date** | Pointer to **string** | The date the credit was issued | [optional] 
**LastSentDate** | Pointer to **string** | The date the credit was last sent out | [optional] 
**NextSendDate** | Pointer to **string** | The next scheduled date for sending a credit reminder | [optional] 
**PartialDueDate** | Pointer to **string** | The due date for the partial amount of the credit | [optional] 
**DueDate** | Pointer to **string** | The due date for the total amount of the credit | [optional] 
**Settings** | Pointer to [**CompanySettings**](CompanySettings.md) |  | [optional] 
**LastViewed** | Pointer to **float32** | The timestamp of the last time the credit was viewed | [optional] 
**UpdatedAt** | Pointer to **float32** | The timestamp of the last time the credit was updated | [optional] 
**ArchivedAt** | Pointer to **float32** | The timestamp of the last time the credit was archived | [optional] 
**CustomSurcharge1** | Pointer to **float32** | First custom surcharge amount | [optional] 
**CustomSurcharge2** | Pointer to **float32** | Second Custom Surcharge | [optional] 
**CustomSurcharge3** | Pointer to **float32** | Third Custom Surcharge | [optional] 
**CustomSurcharge4** | Pointer to **float32** | Fourth Custom Surcharge | [optional] 
**CustomSurchargeTax1** | Pointer to **bool** | Toggles charging taxes on custom surcharge amounts | [optional] 
**CustomSurchargeTax2** | Pointer to **bool** | Toggles charging taxes on custom surcharge amounts | [optional] 
**CustomSurchargeTax3** | Pointer to **bool** | Toggles charging taxes on custom surcharge amounts | [optional] 
**CustomSurchargeTax4** | Pointer to **bool** | Toggles charging taxes on custom surcharge amounts | [optional] 
**LocationId** | Pointer to **string** | The client location id that this invoice relates to | [optional] 

## Methods

### NewCredit

`func NewCredit() *Credit`

NewCredit instantiates a new Credit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditWithDefaults

`func NewCreditWithDefaults() *Credit`

NewCreditWithDefaults instantiates a new Credit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Credit) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Credit) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Credit) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Credit) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *Credit) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Credit) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Credit) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Credit) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAssignedUserId

`func (o *Credit) GetAssignedUserId() string`

GetAssignedUserId returns the AssignedUserId field if non-nil, zero value otherwise.

### GetAssignedUserIdOk

`func (o *Credit) GetAssignedUserIdOk() (*string, bool)`

GetAssignedUserIdOk returns a tuple with the AssignedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUserId

`func (o *Credit) SetAssignedUserId(v string)`

SetAssignedUserId sets AssignedUserId field to given value.

### HasAssignedUserId

`func (o *Credit) HasAssignedUserId() bool`

HasAssignedUserId returns a boolean if a field has been set.

### GetClientId

`func (o *Credit) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Credit) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Credit) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Credit) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetStatusId

`func (o *Credit) GetStatusId() string`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *Credit) GetStatusIdOk() (*string, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *Credit) SetStatusId(v string)`

SetStatusId sets StatusId field to given value.

### HasStatusId

`func (o *Credit) HasStatusId() bool`

HasStatusId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *Credit) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *Credit) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *Credit) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *Credit) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetNumber

`func (o *Credit) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Credit) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Credit) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Credit) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetPoNumber

`func (o *Credit) GetPoNumber() string`

GetPoNumber returns the PoNumber field if non-nil, zero value otherwise.

### GetPoNumberOk

`func (o *Credit) GetPoNumberOk() (*string, bool)`

GetPoNumberOk returns a tuple with the PoNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoNumber

`func (o *Credit) SetPoNumber(v string)`

SetPoNumber sets PoNumber field to given value.

### HasPoNumber

`func (o *Credit) HasPoNumber() bool`

HasPoNumber returns a boolean if a field has been set.

### GetTerms

`func (o *Credit) GetTerms() string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *Credit) GetTermsOk() (*string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *Credit) SetTerms(v string)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *Credit) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### GetPublicNotes

`func (o *Credit) GetPublicNotes() string`

GetPublicNotes returns the PublicNotes field if non-nil, zero value otherwise.

### GetPublicNotesOk

`func (o *Credit) GetPublicNotesOk() (*string, bool)`

GetPublicNotesOk returns a tuple with the PublicNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNotes

`func (o *Credit) SetPublicNotes(v string)`

SetPublicNotes sets PublicNotes field to given value.

### HasPublicNotes

`func (o *Credit) HasPublicNotes() bool`

HasPublicNotes returns a boolean if a field has been set.

### GetPrivateNotes

`func (o *Credit) GetPrivateNotes() string`

GetPrivateNotes returns the PrivateNotes field if non-nil, zero value otherwise.

### GetPrivateNotesOk

`func (o *Credit) GetPrivateNotesOk() (*string, bool)`

GetPrivateNotesOk returns a tuple with the PrivateNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNotes

`func (o *Credit) SetPrivateNotes(v string)`

SetPrivateNotes sets PrivateNotes field to given value.

### HasPrivateNotes

`func (o *Credit) HasPrivateNotes() bool`

HasPrivateNotes returns a boolean if a field has been set.

### GetFooter

`func (o *Credit) GetFooter() string`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *Credit) GetFooterOk() (*string, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *Credit) SetFooter(v string)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *Credit) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetCustomValue1

`func (o *Credit) GetCustomValue1() string`

GetCustomValue1 returns the CustomValue1 field if non-nil, zero value otherwise.

### GetCustomValue1Ok

`func (o *Credit) GetCustomValue1Ok() (*string, bool)`

GetCustomValue1Ok returns a tuple with the CustomValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue1

`func (o *Credit) SetCustomValue1(v string)`

SetCustomValue1 sets CustomValue1 field to given value.

### HasCustomValue1

`func (o *Credit) HasCustomValue1() bool`

HasCustomValue1 returns a boolean if a field has been set.

### GetCustomValue2

`func (o *Credit) GetCustomValue2() string`

GetCustomValue2 returns the CustomValue2 field if non-nil, zero value otherwise.

### GetCustomValue2Ok

`func (o *Credit) GetCustomValue2Ok() (*string, bool)`

GetCustomValue2Ok returns a tuple with the CustomValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue2

`func (o *Credit) SetCustomValue2(v string)`

SetCustomValue2 sets CustomValue2 field to given value.

### HasCustomValue2

`func (o *Credit) HasCustomValue2() bool`

HasCustomValue2 returns a boolean if a field has been set.

### GetCustomValue3

`func (o *Credit) GetCustomValue3() string`

GetCustomValue3 returns the CustomValue3 field if non-nil, zero value otherwise.

### GetCustomValue3Ok

`func (o *Credit) GetCustomValue3Ok() (*string, bool)`

GetCustomValue3Ok returns a tuple with the CustomValue3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue3

`func (o *Credit) SetCustomValue3(v string)`

SetCustomValue3 sets CustomValue3 field to given value.

### HasCustomValue3

`func (o *Credit) HasCustomValue3() bool`

HasCustomValue3 returns a boolean if a field has been set.

### GetCustomValue4

`func (o *Credit) GetCustomValue4() string`

GetCustomValue4 returns the CustomValue4 field if non-nil, zero value otherwise.

### GetCustomValue4Ok

`func (o *Credit) GetCustomValue4Ok() (*string, bool)`

GetCustomValue4Ok returns a tuple with the CustomValue4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue4

`func (o *Credit) SetCustomValue4(v string)`

SetCustomValue4 sets CustomValue4 field to given value.

### HasCustomValue4

`func (o *Credit) HasCustomValue4() bool`

HasCustomValue4 returns a boolean if a field has been set.

### GetTaxName1

`func (o *Credit) GetTaxName1() string`

GetTaxName1 returns the TaxName1 field if non-nil, zero value otherwise.

### GetTaxName1Ok

`func (o *Credit) GetTaxName1Ok() (*string, bool)`

GetTaxName1Ok returns a tuple with the TaxName1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName1

`func (o *Credit) SetTaxName1(v string)`

SetTaxName1 sets TaxName1 field to given value.

### HasTaxName1

`func (o *Credit) HasTaxName1() bool`

HasTaxName1 returns a boolean if a field has been set.

### GetTaxName2

`func (o *Credit) GetTaxName2() string`

GetTaxName2 returns the TaxName2 field if non-nil, zero value otherwise.

### GetTaxName2Ok

`func (o *Credit) GetTaxName2Ok() (*string, bool)`

GetTaxName2Ok returns a tuple with the TaxName2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName2

`func (o *Credit) SetTaxName2(v string)`

SetTaxName2 sets TaxName2 field to given value.

### HasTaxName2

`func (o *Credit) HasTaxName2() bool`

HasTaxName2 returns a boolean if a field has been set.

### GetTaxRate1

`func (o *Credit) GetTaxRate1() float32`

GetTaxRate1 returns the TaxRate1 field if non-nil, zero value otherwise.

### GetTaxRate1Ok

`func (o *Credit) GetTaxRate1Ok() (*float32, bool)`

GetTaxRate1Ok returns a tuple with the TaxRate1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate1

`func (o *Credit) SetTaxRate1(v float32)`

SetTaxRate1 sets TaxRate1 field to given value.

### HasTaxRate1

`func (o *Credit) HasTaxRate1() bool`

HasTaxRate1 returns a boolean if a field has been set.

### GetTaxRate2

`func (o *Credit) GetTaxRate2() float32`

GetTaxRate2 returns the TaxRate2 field if non-nil, zero value otherwise.

### GetTaxRate2Ok

`func (o *Credit) GetTaxRate2Ok() (*float32, bool)`

GetTaxRate2Ok returns a tuple with the TaxRate2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate2

`func (o *Credit) SetTaxRate2(v float32)`

SetTaxRate2 sets TaxRate2 field to given value.

### HasTaxRate2

`func (o *Credit) HasTaxRate2() bool`

HasTaxRate2 returns a boolean if a field has been set.

### GetTaxName3

`func (o *Credit) GetTaxName3() string`

GetTaxName3 returns the TaxName3 field if non-nil, zero value otherwise.

### GetTaxName3Ok

`func (o *Credit) GetTaxName3Ok() (*string, bool)`

GetTaxName3Ok returns a tuple with the TaxName3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName3

`func (o *Credit) SetTaxName3(v string)`

SetTaxName3 sets TaxName3 field to given value.

### HasTaxName3

`func (o *Credit) HasTaxName3() bool`

HasTaxName3 returns a boolean if a field has been set.

### GetTaxRate3

`func (o *Credit) GetTaxRate3() float32`

GetTaxRate3 returns the TaxRate3 field if non-nil, zero value otherwise.

### GetTaxRate3Ok

`func (o *Credit) GetTaxRate3Ok() (*float32, bool)`

GetTaxRate3Ok returns a tuple with the TaxRate3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate3

`func (o *Credit) SetTaxRate3(v float32)`

SetTaxRate3 sets TaxRate3 field to given value.

### HasTaxRate3

`func (o *Credit) HasTaxRate3() bool`

HasTaxRate3 returns a boolean if a field has been set.

### GetTotalTaxes

`func (o *Credit) GetTotalTaxes() float32`

GetTotalTaxes returns the TotalTaxes field if non-nil, zero value otherwise.

### GetTotalTaxesOk

`func (o *Credit) GetTotalTaxesOk() (*float32, bool)`

GetTotalTaxesOk returns a tuple with the TotalTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxes

`func (o *Credit) SetTotalTaxes(v float32)`

SetTotalTaxes sets TotalTaxes field to given value.

### HasTotalTaxes

`func (o *Credit) HasTotalTaxes() bool`

HasTotalTaxes returns a boolean if a field has been set.

### GetLineItems

`func (o *Credit) GetLineItems() []InvoiceItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *Credit) GetLineItemsOk() (*[]InvoiceItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *Credit) SetLineItems(v []InvoiceItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *Credit) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetAmount

`func (o *Credit) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Credit) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Credit) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Credit) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBalance

`func (o *Credit) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *Credit) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *Credit) SetBalance(v float32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *Credit) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetPaidToDate

`func (o *Credit) GetPaidToDate() float32`

GetPaidToDate returns the PaidToDate field if non-nil, zero value otherwise.

### GetPaidToDateOk

`func (o *Credit) GetPaidToDateOk() (*float32, bool)`

GetPaidToDateOk returns a tuple with the PaidToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidToDate

`func (o *Credit) SetPaidToDate(v float32)`

SetPaidToDate sets PaidToDate field to given value.

### HasPaidToDate

`func (o *Credit) HasPaidToDate() bool`

HasPaidToDate returns a boolean if a field has been set.

### GetDiscount

`func (o *Credit) GetDiscount() float32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *Credit) GetDiscountOk() (*float32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *Credit) SetDiscount(v float32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *Credit) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetPartial

`func (o *Credit) GetPartial() float32`

GetPartial returns the Partial field if non-nil, zero value otherwise.

### GetPartialOk

`func (o *Credit) GetPartialOk() (*float32, bool)`

GetPartialOk returns a tuple with the Partial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartial

`func (o *Credit) SetPartial(v float32)`

SetPartial sets Partial field to given value.

### HasPartial

`func (o *Credit) HasPartial() bool`

HasPartial returns a boolean if a field has been set.

### GetIsAmountDiscount

`func (o *Credit) GetIsAmountDiscount() bool`

GetIsAmountDiscount returns the IsAmountDiscount field if non-nil, zero value otherwise.

### GetIsAmountDiscountOk

`func (o *Credit) GetIsAmountDiscountOk() (*bool, bool)`

GetIsAmountDiscountOk returns a tuple with the IsAmountDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAmountDiscount

`func (o *Credit) SetIsAmountDiscount(v bool)`

SetIsAmountDiscount sets IsAmountDiscount field to given value.

### HasIsAmountDiscount

`func (o *Credit) HasIsAmountDiscount() bool`

HasIsAmountDiscount returns a boolean if a field has been set.

### GetIsDeleted

`func (o *Credit) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *Credit) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *Credit) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *Credit) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetUsesInclusiveTaxes

`func (o *Credit) GetUsesInclusiveTaxes() bool`

GetUsesInclusiveTaxes returns the UsesInclusiveTaxes field if non-nil, zero value otherwise.

### GetUsesInclusiveTaxesOk

`func (o *Credit) GetUsesInclusiveTaxesOk() (*bool, bool)`

GetUsesInclusiveTaxesOk returns a tuple with the UsesInclusiveTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesInclusiveTaxes

`func (o *Credit) SetUsesInclusiveTaxes(v bool)`

SetUsesInclusiveTaxes sets UsesInclusiveTaxes field to given value.

### HasUsesInclusiveTaxes

`func (o *Credit) HasUsesInclusiveTaxes() bool`

HasUsesInclusiveTaxes returns a boolean if a field has been set.

### GetDate

`func (o *Credit) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Credit) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Credit) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *Credit) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetLastSentDate

`func (o *Credit) GetLastSentDate() string`

GetLastSentDate returns the LastSentDate field if non-nil, zero value otherwise.

### GetLastSentDateOk

`func (o *Credit) GetLastSentDateOk() (*string, bool)`

GetLastSentDateOk returns a tuple with the LastSentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSentDate

`func (o *Credit) SetLastSentDate(v string)`

SetLastSentDate sets LastSentDate field to given value.

### HasLastSentDate

`func (o *Credit) HasLastSentDate() bool`

HasLastSentDate returns a boolean if a field has been set.

### GetNextSendDate

`func (o *Credit) GetNextSendDate() string`

GetNextSendDate returns the NextSendDate field if non-nil, zero value otherwise.

### GetNextSendDateOk

`func (o *Credit) GetNextSendDateOk() (*string, bool)`

GetNextSendDateOk returns a tuple with the NextSendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSendDate

`func (o *Credit) SetNextSendDate(v string)`

SetNextSendDate sets NextSendDate field to given value.

### HasNextSendDate

`func (o *Credit) HasNextSendDate() bool`

HasNextSendDate returns a boolean if a field has been set.

### GetPartialDueDate

`func (o *Credit) GetPartialDueDate() string`

GetPartialDueDate returns the PartialDueDate field if non-nil, zero value otherwise.

### GetPartialDueDateOk

`func (o *Credit) GetPartialDueDateOk() (*string, bool)`

GetPartialDueDateOk returns a tuple with the PartialDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialDueDate

`func (o *Credit) SetPartialDueDate(v string)`

SetPartialDueDate sets PartialDueDate field to given value.

### HasPartialDueDate

`func (o *Credit) HasPartialDueDate() bool`

HasPartialDueDate returns a boolean if a field has been set.

### GetDueDate

`func (o *Credit) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *Credit) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *Credit) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *Credit) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetSettings

`func (o *Credit) GetSettings() CompanySettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *Credit) GetSettingsOk() (*CompanySettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *Credit) SetSettings(v CompanySettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *Credit) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetLastViewed

`func (o *Credit) GetLastViewed() float32`

GetLastViewed returns the LastViewed field if non-nil, zero value otherwise.

### GetLastViewedOk

`func (o *Credit) GetLastViewedOk() (*float32, bool)`

GetLastViewedOk returns a tuple with the LastViewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastViewed

`func (o *Credit) SetLastViewed(v float32)`

SetLastViewed sets LastViewed field to given value.

### HasLastViewed

`func (o *Credit) HasLastViewed() bool`

HasLastViewed returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Credit) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Credit) GetUpdatedAtOk() (*float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Credit) SetUpdatedAt(v float32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Credit) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetArchivedAt

`func (o *Credit) GetArchivedAt() float32`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *Credit) GetArchivedAtOk() (*float32, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *Credit) SetArchivedAt(v float32)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *Credit) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetCustomSurcharge1

`func (o *Credit) GetCustomSurcharge1() float32`

GetCustomSurcharge1 returns the CustomSurcharge1 field if non-nil, zero value otherwise.

### GetCustomSurcharge1Ok

`func (o *Credit) GetCustomSurcharge1Ok() (*float32, bool)`

GetCustomSurcharge1Ok returns a tuple with the CustomSurcharge1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurcharge1

`func (o *Credit) SetCustomSurcharge1(v float32)`

SetCustomSurcharge1 sets CustomSurcharge1 field to given value.

### HasCustomSurcharge1

`func (o *Credit) HasCustomSurcharge1() bool`

HasCustomSurcharge1 returns a boolean if a field has been set.

### GetCustomSurcharge2

`func (o *Credit) GetCustomSurcharge2() float32`

GetCustomSurcharge2 returns the CustomSurcharge2 field if non-nil, zero value otherwise.

### GetCustomSurcharge2Ok

`func (o *Credit) GetCustomSurcharge2Ok() (*float32, bool)`

GetCustomSurcharge2Ok returns a tuple with the CustomSurcharge2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurcharge2

`func (o *Credit) SetCustomSurcharge2(v float32)`

SetCustomSurcharge2 sets CustomSurcharge2 field to given value.

### HasCustomSurcharge2

`func (o *Credit) HasCustomSurcharge2() bool`

HasCustomSurcharge2 returns a boolean if a field has been set.

### GetCustomSurcharge3

`func (o *Credit) GetCustomSurcharge3() float32`

GetCustomSurcharge3 returns the CustomSurcharge3 field if non-nil, zero value otherwise.

### GetCustomSurcharge3Ok

`func (o *Credit) GetCustomSurcharge3Ok() (*float32, bool)`

GetCustomSurcharge3Ok returns a tuple with the CustomSurcharge3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurcharge3

`func (o *Credit) SetCustomSurcharge3(v float32)`

SetCustomSurcharge3 sets CustomSurcharge3 field to given value.

### HasCustomSurcharge3

`func (o *Credit) HasCustomSurcharge3() bool`

HasCustomSurcharge3 returns a boolean if a field has been set.

### GetCustomSurcharge4

`func (o *Credit) GetCustomSurcharge4() float32`

GetCustomSurcharge4 returns the CustomSurcharge4 field if non-nil, zero value otherwise.

### GetCustomSurcharge4Ok

`func (o *Credit) GetCustomSurcharge4Ok() (*float32, bool)`

GetCustomSurcharge4Ok returns a tuple with the CustomSurcharge4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurcharge4

`func (o *Credit) SetCustomSurcharge4(v float32)`

SetCustomSurcharge4 sets CustomSurcharge4 field to given value.

### HasCustomSurcharge4

`func (o *Credit) HasCustomSurcharge4() bool`

HasCustomSurcharge4 returns a boolean if a field has been set.

### GetCustomSurchargeTax1

`func (o *Credit) GetCustomSurchargeTax1() bool`

GetCustomSurchargeTax1 returns the CustomSurchargeTax1 field if non-nil, zero value otherwise.

### GetCustomSurchargeTax1Ok

`func (o *Credit) GetCustomSurchargeTax1Ok() (*bool, bool)`

GetCustomSurchargeTax1Ok returns a tuple with the CustomSurchargeTax1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurchargeTax1

`func (o *Credit) SetCustomSurchargeTax1(v bool)`

SetCustomSurchargeTax1 sets CustomSurchargeTax1 field to given value.

### HasCustomSurchargeTax1

`func (o *Credit) HasCustomSurchargeTax1() bool`

HasCustomSurchargeTax1 returns a boolean if a field has been set.

### GetCustomSurchargeTax2

`func (o *Credit) GetCustomSurchargeTax2() bool`

GetCustomSurchargeTax2 returns the CustomSurchargeTax2 field if non-nil, zero value otherwise.

### GetCustomSurchargeTax2Ok

`func (o *Credit) GetCustomSurchargeTax2Ok() (*bool, bool)`

GetCustomSurchargeTax2Ok returns a tuple with the CustomSurchargeTax2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurchargeTax2

`func (o *Credit) SetCustomSurchargeTax2(v bool)`

SetCustomSurchargeTax2 sets CustomSurchargeTax2 field to given value.

### HasCustomSurchargeTax2

`func (o *Credit) HasCustomSurchargeTax2() bool`

HasCustomSurchargeTax2 returns a boolean if a field has been set.

### GetCustomSurchargeTax3

`func (o *Credit) GetCustomSurchargeTax3() bool`

GetCustomSurchargeTax3 returns the CustomSurchargeTax3 field if non-nil, zero value otherwise.

### GetCustomSurchargeTax3Ok

`func (o *Credit) GetCustomSurchargeTax3Ok() (*bool, bool)`

GetCustomSurchargeTax3Ok returns a tuple with the CustomSurchargeTax3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurchargeTax3

`func (o *Credit) SetCustomSurchargeTax3(v bool)`

SetCustomSurchargeTax3 sets CustomSurchargeTax3 field to given value.

### HasCustomSurchargeTax3

`func (o *Credit) HasCustomSurchargeTax3() bool`

HasCustomSurchargeTax3 returns a boolean if a field has been set.

### GetCustomSurchargeTax4

`func (o *Credit) GetCustomSurchargeTax4() bool`

GetCustomSurchargeTax4 returns the CustomSurchargeTax4 field if non-nil, zero value otherwise.

### GetCustomSurchargeTax4Ok

`func (o *Credit) GetCustomSurchargeTax4Ok() (*bool, bool)`

GetCustomSurchargeTax4Ok returns a tuple with the CustomSurchargeTax4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurchargeTax4

`func (o *Credit) SetCustomSurchargeTax4(v bool)`

SetCustomSurchargeTax4 sets CustomSurchargeTax4 field to given value.

### HasCustomSurchargeTax4

`func (o *Credit) HasCustomSurchargeTax4() bool`

HasCustomSurchargeTax4 returns a boolean if a field has been set.

### GetLocationId

`func (o *Credit) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *Credit) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *Credit) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *Credit) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


