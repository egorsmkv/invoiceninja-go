# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The invoice hashed id | [optional] [readonly] 
**UserId** | Pointer to **string** | The user hashed id | [optional] [readonly] 
**AssignedUserId** | Pointer to **string** | The assigned user hashed id | [optional] 
**ClientId** | Pointer to **string** | The client hashed id | [optional] 
**StatusId** | Pointer to **string** | The invoice status variable | [optional] 
**Number** | Pointer to **string** | The invoice number - is a unique alpha numeric number per invoice per company | [optional] 
**PoNumber** | Pointer to **string** | The purchase order associated with this invoice | [optional] 
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
**TotalTaxes** | Pointer to **float32** | The total taxes for the invoice | [optional] 
**LineItems** | Pointer to [**[]InvoiceItem**](InvoiceItem.md) | An array of objects which define the line items of the invoice | [optional] 
**Invitations** | Pointer to [**[]InvoiceInvitation**](InvoiceInvitation.md) | An array of objects which define the invitations of the invoice | [optional] 
**Amount** | Pointer to **float32** | The invoice amount | [optional] 
**Balance** | Pointer to **float32** | The invoice balance | [optional] 
**PaidToDate** | Pointer to **float32** | The amount paid on the invoice to date | [optional] 
**Discount** | Pointer to **float32** | The invoice discount, can be an amount or a percentage | [optional] 
**Partial** | Pointer to **float32** | The deposit/partial amount | [optional] 
**IsAmountDiscount** | Pointer to **bool** | Flag determining if the discount is an amount or a percentage | [optional] 
**IsDeleted** | Pointer to **bool** | Defines if the invoice has been deleted | [optional] 
**UsesInclusiveTaxes** | Pointer to **bool** | Defines the type of taxes used as either inclusive or exclusive | [optional] 
**Date** | Pointer to **string** | The Invoice Date | [optional] 
**LastSentDate** | Pointer to **string** | The last date the invoice was sent out | [optional] 
**NextSendDate** | Pointer to **string** | The Next date for a reminder to be sent | [optional] 
**PartialDueDate** | Pointer to **string** | The due date for the deposit/partial amount | [optional] 
**DueDate** | Pointer to **string** | The due date of the invoice | [optional] 
**LastViewed** | Pointer to **float32** | Timestamp | [optional] 
**UpdatedAt** | Pointer to **float32** | Timestamp | [optional] 
**ArchivedAt** | Pointer to **float32** | Timestamp | [optional] 
**CustomSurcharge1** | Pointer to **float32** | First Custom Surcharge | [optional] 
**CustomSurcharge2** | Pointer to **float32** | Second Custom Surcharge | [optional] 
**CustomSurcharge3** | Pointer to **float32** | Third Custom Surcharge | [optional] 
**CustomSurcharge4** | Pointer to **float32** | Fourth Custom Surcharge | [optional] 
**CustomSurchargeTax1** | Pointer to **bool** | Toggles charging taxes on custom surcharge amounts | [optional] 
**CustomSurchargeTax2** | Pointer to **bool** | Toggles charging taxes on custom surcharge amounts | [optional] 
**CustomSurchargeTax3** | Pointer to **bool** | Toggles charging taxes on custom surcharge amounts | [optional] 
**CustomSurchargeTax4** | Pointer to **bool** | Toggles charging taxes on custom surcharge amounts | [optional] 
**ProjectId** | Pointer to **string** | The project associated with this invoice | [optional] 
**AutoBillTries** | Pointer to **int32** | The number of times the invoice has attempted to be auto billed | [optional] [readonly] 
**AutoBillEnabled** | Pointer to **bool** | Boolean flag determining if the invoice is set to auto bill | [optional] 
**SubscriptionId** | Pointer to **string** | The subscription associated with this invoice | [optional] 
**LocationId** | Pointer to **string** | The client location id that this invoice relates to | [optional] 

## Methods

### NewInvoice

`func NewInvoice() *Invoice`

NewInvoice instantiates a new Invoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWithDefaults

`func NewInvoiceWithDefaults() *Invoice`

NewInvoiceWithDefaults instantiates a new Invoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Invoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Invoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Invoice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Invoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *Invoice) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Invoice) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Invoice) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Invoice) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAssignedUserId

`func (o *Invoice) GetAssignedUserId() string`

GetAssignedUserId returns the AssignedUserId field if non-nil, zero value otherwise.

### GetAssignedUserIdOk

`func (o *Invoice) GetAssignedUserIdOk() (*string, bool)`

GetAssignedUserIdOk returns a tuple with the AssignedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUserId

`func (o *Invoice) SetAssignedUserId(v string)`

SetAssignedUserId sets AssignedUserId field to given value.

### HasAssignedUserId

`func (o *Invoice) HasAssignedUserId() bool`

HasAssignedUserId returns a boolean if a field has been set.

### GetClientId

`func (o *Invoice) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Invoice) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Invoice) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Invoice) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetStatusId

`func (o *Invoice) GetStatusId() string`

GetStatusId returns the StatusId field if non-nil, zero value otherwise.

### GetStatusIdOk

`func (o *Invoice) GetStatusIdOk() (*string, bool)`

GetStatusIdOk returns a tuple with the StatusId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusId

`func (o *Invoice) SetStatusId(v string)`

SetStatusId sets StatusId field to given value.

### HasStatusId

`func (o *Invoice) HasStatusId() bool`

HasStatusId returns a boolean if a field has been set.

### GetNumber

`func (o *Invoice) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Invoice) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Invoice) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Invoice) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetPoNumber

`func (o *Invoice) GetPoNumber() string`

GetPoNumber returns the PoNumber field if non-nil, zero value otherwise.

### GetPoNumberOk

`func (o *Invoice) GetPoNumberOk() (*string, bool)`

GetPoNumberOk returns a tuple with the PoNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoNumber

`func (o *Invoice) SetPoNumber(v string)`

SetPoNumber sets PoNumber field to given value.

### HasPoNumber

`func (o *Invoice) HasPoNumber() bool`

HasPoNumber returns a boolean if a field has been set.

### GetTerms

`func (o *Invoice) GetTerms() string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *Invoice) GetTermsOk() (*string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *Invoice) SetTerms(v string)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *Invoice) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### GetPublicNotes

`func (o *Invoice) GetPublicNotes() string`

GetPublicNotes returns the PublicNotes field if non-nil, zero value otherwise.

### GetPublicNotesOk

`func (o *Invoice) GetPublicNotesOk() (*string, bool)`

GetPublicNotesOk returns a tuple with the PublicNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNotes

`func (o *Invoice) SetPublicNotes(v string)`

SetPublicNotes sets PublicNotes field to given value.

### HasPublicNotes

`func (o *Invoice) HasPublicNotes() bool`

HasPublicNotes returns a boolean if a field has been set.

### GetPrivateNotes

`func (o *Invoice) GetPrivateNotes() string`

GetPrivateNotes returns the PrivateNotes field if non-nil, zero value otherwise.

### GetPrivateNotesOk

`func (o *Invoice) GetPrivateNotesOk() (*string, bool)`

GetPrivateNotesOk returns a tuple with the PrivateNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNotes

`func (o *Invoice) SetPrivateNotes(v string)`

SetPrivateNotes sets PrivateNotes field to given value.

### HasPrivateNotes

`func (o *Invoice) HasPrivateNotes() bool`

HasPrivateNotes returns a boolean if a field has been set.

### GetFooter

`func (o *Invoice) GetFooter() string`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *Invoice) GetFooterOk() (*string, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *Invoice) SetFooter(v string)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *Invoice) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetCustomValue1

`func (o *Invoice) GetCustomValue1() string`

GetCustomValue1 returns the CustomValue1 field if non-nil, zero value otherwise.

### GetCustomValue1Ok

`func (o *Invoice) GetCustomValue1Ok() (*string, bool)`

GetCustomValue1Ok returns a tuple with the CustomValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue1

`func (o *Invoice) SetCustomValue1(v string)`

SetCustomValue1 sets CustomValue1 field to given value.

### HasCustomValue1

`func (o *Invoice) HasCustomValue1() bool`

HasCustomValue1 returns a boolean if a field has been set.

### GetCustomValue2

`func (o *Invoice) GetCustomValue2() string`

GetCustomValue2 returns the CustomValue2 field if non-nil, zero value otherwise.

### GetCustomValue2Ok

`func (o *Invoice) GetCustomValue2Ok() (*string, bool)`

GetCustomValue2Ok returns a tuple with the CustomValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue2

`func (o *Invoice) SetCustomValue2(v string)`

SetCustomValue2 sets CustomValue2 field to given value.

### HasCustomValue2

`func (o *Invoice) HasCustomValue2() bool`

HasCustomValue2 returns a boolean if a field has been set.

### GetCustomValue3

`func (o *Invoice) GetCustomValue3() string`

GetCustomValue3 returns the CustomValue3 field if non-nil, zero value otherwise.

### GetCustomValue3Ok

`func (o *Invoice) GetCustomValue3Ok() (*string, bool)`

GetCustomValue3Ok returns a tuple with the CustomValue3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue3

`func (o *Invoice) SetCustomValue3(v string)`

SetCustomValue3 sets CustomValue3 field to given value.

### HasCustomValue3

`func (o *Invoice) HasCustomValue3() bool`

HasCustomValue3 returns a boolean if a field has been set.

### GetCustomValue4

`func (o *Invoice) GetCustomValue4() string`

GetCustomValue4 returns the CustomValue4 field if non-nil, zero value otherwise.

### GetCustomValue4Ok

`func (o *Invoice) GetCustomValue4Ok() (*string, bool)`

GetCustomValue4Ok returns a tuple with the CustomValue4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue4

`func (o *Invoice) SetCustomValue4(v string)`

SetCustomValue4 sets CustomValue4 field to given value.

### HasCustomValue4

`func (o *Invoice) HasCustomValue4() bool`

HasCustomValue4 returns a boolean if a field has been set.

### GetTaxName1

`func (o *Invoice) GetTaxName1() string`

GetTaxName1 returns the TaxName1 field if non-nil, zero value otherwise.

### GetTaxName1Ok

`func (o *Invoice) GetTaxName1Ok() (*string, bool)`

GetTaxName1Ok returns a tuple with the TaxName1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName1

`func (o *Invoice) SetTaxName1(v string)`

SetTaxName1 sets TaxName1 field to given value.

### HasTaxName1

`func (o *Invoice) HasTaxName1() bool`

HasTaxName1 returns a boolean if a field has been set.

### GetTaxName2

`func (o *Invoice) GetTaxName2() string`

GetTaxName2 returns the TaxName2 field if non-nil, zero value otherwise.

### GetTaxName2Ok

`func (o *Invoice) GetTaxName2Ok() (*string, bool)`

GetTaxName2Ok returns a tuple with the TaxName2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName2

`func (o *Invoice) SetTaxName2(v string)`

SetTaxName2 sets TaxName2 field to given value.

### HasTaxName2

`func (o *Invoice) HasTaxName2() bool`

HasTaxName2 returns a boolean if a field has been set.

### GetTaxRate1

`func (o *Invoice) GetTaxRate1() float32`

GetTaxRate1 returns the TaxRate1 field if non-nil, zero value otherwise.

### GetTaxRate1Ok

`func (o *Invoice) GetTaxRate1Ok() (*float32, bool)`

GetTaxRate1Ok returns a tuple with the TaxRate1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate1

`func (o *Invoice) SetTaxRate1(v float32)`

SetTaxRate1 sets TaxRate1 field to given value.

### HasTaxRate1

`func (o *Invoice) HasTaxRate1() bool`

HasTaxRate1 returns a boolean if a field has been set.

### GetTaxRate2

`func (o *Invoice) GetTaxRate2() float32`

GetTaxRate2 returns the TaxRate2 field if non-nil, zero value otherwise.

### GetTaxRate2Ok

`func (o *Invoice) GetTaxRate2Ok() (*float32, bool)`

GetTaxRate2Ok returns a tuple with the TaxRate2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate2

`func (o *Invoice) SetTaxRate2(v float32)`

SetTaxRate2 sets TaxRate2 field to given value.

### HasTaxRate2

`func (o *Invoice) HasTaxRate2() bool`

HasTaxRate2 returns a boolean if a field has been set.

### GetTaxName3

`func (o *Invoice) GetTaxName3() string`

GetTaxName3 returns the TaxName3 field if non-nil, zero value otherwise.

### GetTaxName3Ok

`func (o *Invoice) GetTaxName3Ok() (*string, bool)`

GetTaxName3Ok returns a tuple with the TaxName3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName3

`func (o *Invoice) SetTaxName3(v string)`

SetTaxName3 sets TaxName3 field to given value.

### HasTaxName3

`func (o *Invoice) HasTaxName3() bool`

HasTaxName3 returns a boolean if a field has been set.

### GetTaxRate3

`func (o *Invoice) GetTaxRate3() float32`

GetTaxRate3 returns the TaxRate3 field if non-nil, zero value otherwise.

### GetTaxRate3Ok

`func (o *Invoice) GetTaxRate3Ok() (*float32, bool)`

GetTaxRate3Ok returns a tuple with the TaxRate3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate3

`func (o *Invoice) SetTaxRate3(v float32)`

SetTaxRate3 sets TaxRate3 field to given value.

### HasTaxRate3

`func (o *Invoice) HasTaxRate3() bool`

HasTaxRate3 returns a boolean if a field has been set.

### GetTotalTaxes

`func (o *Invoice) GetTotalTaxes() float32`

GetTotalTaxes returns the TotalTaxes field if non-nil, zero value otherwise.

### GetTotalTaxesOk

`func (o *Invoice) GetTotalTaxesOk() (*float32, bool)`

GetTotalTaxesOk returns a tuple with the TotalTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTaxes

`func (o *Invoice) SetTotalTaxes(v float32)`

SetTotalTaxes sets TotalTaxes field to given value.

### HasTotalTaxes

`func (o *Invoice) HasTotalTaxes() bool`

HasTotalTaxes returns a boolean if a field has been set.

### GetLineItems

`func (o *Invoice) GetLineItems() []InvoiceItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *Invoice) GetLineItemsOk() (*[]InvoiceItem, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *Invoice) SetLineItems(v []InvoiceItem)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *Invoice) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetInvitations

`func (o *Invoice) GetInvitations() []InvoiceInvitation`

GetInvitations returns the Invitations field if non-nil, zero value otherwise.

### GetInvitationsOk

`func (o *Invoice) GetInvitationsOk() (*[]InvoiceInvitation, bool)`

GetInvitationsOk returns a tuple with the Invitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitations

`func (o *Invoice) SetInvitations(v []InvoiceInvitation)`

SetInvitations sets Invitations field to given value.

### HasInvitations

`func (o *Invoice) HasInvitations() bool`

HasInvitations returns a boolean if a field has been set.

### GetAmount

`func (o *Invoice) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Invoice) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Invoice) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Invoice) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetBalance

`func (o *Invoice) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *Invoice) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *Invoice) SetBalance(v float32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *Invoice) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetPaidToDate

`func (o *Invoice) GetPaidToDate() float32`

GetPaidToDate returns the PaidToDate field if non-nil, zero value otherwise.

### GetPaidToDateOk

`func (o *Invoice) GetPaidToDateOk() (*float32, bool)`

GetPaidToDateOk returns a tuple with the PaidToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidToDate

`func (o *Invoice) SetPaidToDate(v float32)`

SetPaidToDate sets PaidToDate field to given value.

### HasPaidToDate

`func (o *Invoice) HasPaidToDate() bool`

HasPaidToDate returns a boolean if a field has been set.

### GetDiscount

`func (o *Invoice) GetDiscount() float32`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *Invoice) GetDiscountOk() (*float32, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *Invoice) SetDiscount(v float32)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *Invoice) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetPartial

`func (o *Invoice) GetPartial() float32`

GetPartial returns the Partial field if non-nil, zero value otherwise.

### GetPartialOk

`func (o *Invoice) GetPartialOk() (*float32, bool)`

GetPartialOk returns a tuple with the Partial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartial

`func (o *Invoice) SetPartial(v float32)`

SetPartial sets Partial field to given value.

### HasPartial

`func (o *Invoice) HasPartial() bool`

HasPartial returns a boolean if a field has been set.

### GetIsAmountDiscount

`func (o *Invoice) GetIsAmountDiscount() bool`

GetIsAmountDiscount returns the IsAmountDiscount field if non-nil, zero value otherwise.

### GetIsAmountDiscountOk

`func (o *Invoice) GetIsAmountDiscountOk() (*bool, bool)`

GetIsAmountDiscountOk returns a tuple with the IsAmountDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAmountDiscount

`func (o *Invoice) SetIsAmountDiscount(v bool)`

SetIsAmountDiscount sets IsAmountDiscount field to given value.

### HasIsAmountDiscount

`func (o *Invoice) HasIsAmountDiscount() bool`

HasIsAmountDiscount returns a boolean if a field has been set.

### GetIsDeleted

`func (o *Invoice) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *Invoice) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *Invoice) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *Invoice) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetUsesInclusiveTaxes

`func (o *Invoice) GetUsesInclusiveTaxes() bool`

GetUsesInclusiveTaxes returns the UsesInclusiveTaxes field if non-nil, zero value otherwise.

### GetUsesInclusiveTaxesOk

`func (o *Invoice) GetUsesInclusiveTaxesOk() (*bool, bool)`

GetUsesInclusiveTaxesOk returns a tuple with the UsesInclusiveTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesInclusiveTaxes

`func (o *Invoice) SetUsesInclusiveTaxes(v bool)`

SetUsesInclusiveTaxes sets UsesInclusiveTaxes field to given value.

### HasUsesInclusiveTaxes

`func (o *Invoice) HasUsesInclusiveTaxes() bool`

HasUsesInclusiveTaxes returns a boolean if a field has been set.

### GetDate

`func (o *Invoice) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Invoice) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Invoice) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *Invoice) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetLastSentDate

`func (o *Invoice) GetLastSentDate() string`

GetLastSentDate returns the LastSentDate field if non-nil, zero value otherwise.

### GetLastSentDateOk

`func (o *Invoice) GetLastSentDateOk() (*string, bool)`

GetLastSentDateOk returns a tuple with the LastSentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSentDate

`func (o *Invoice) SetLastSentDate(v string)`

SetLastSentDate sets LastSentDate field to given value.

### HasLastSentDate

`func (o *Invoice) HasLastSentDate() bool`

HasLastSentDate returns a boolean if a field has been set.

### GetNextSendDate

`func (o *Invoice) GetNextSendDate() string`

GetNextSendDate returns the NextSendDate field if non-nil, zero value otherwise.

### GetNextSendDateOk

`func (o *Invoice) GetNextSendDateOk() (*string, bool)`

GetNextSendDateOk returns a tuple with the NextSendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSendDate

`func (o *Invoice) SetNextSendDate(v string)`

SetNextSendDate sets NextSendDate field to given value.

### HasNextSendDate

`func (o *Invoice) HasNextSendDate() bool`

HasNextSendDate returns a boolean if a field has been set.

### GetPartialDueDate

`func (o *Invoice) GetPartialDueDate() string`

GetPartialDueDate returns the PartialDueDate field if non-nil, zero value otherwise.

### GetPartialDueDateOk

`func (o *Invoice) GetPartialDueDateOk() (*string, bool)`

GetPartialDueDateOk returns a tuple with the PartialDueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartialDueDate

`func (o *Invoice) SetPartialDueDate(v string)`

SetPartialDueDate sets PartialDueDate field to given value.

### HasPartialDueDate

`func (o *Invoice) HasPartialDueDate() bool`

HasPartialDueDate returns a boolean if a field has been set.

### GetDueDate

`func (o *Invoice) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *Invoice) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *Invoice) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *Invoice) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetLastViewed

`func (o *Invoice) GetLastViewed() float32`

GetLastViewed returns the LastViewed field if non-nil, zero value otherwise.

### GetLastViewedOk

`func (o *Invoice) GetLastViewedOk() (*float32, bool)`

GetLastViewedOk returns a tuple with the LastViewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastViewed

`func (o *Invoice) SetLastViewed(v float32)`

SetLastViewed sets LastViewed field to given value.

### HasLastViewed

`func (o *Invoice) HasLastViewed() bool`

HasLastViewed returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Invoice) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Invoice) GetUpdatedAtOk() (*float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Invoice) SetUpdatedAt(v float32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Invoice) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetArchivedAt

`func (o *Invoice) GetArchivedAt() float32`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *Invoice) GetArchivedAtOk() (*float32, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *Invoice) SetArchivedAt(v float32)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *Invoice) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetCustomSurcharge1

`func (o *Invoice) GetCustomSurcharge1() float32`

GetCustomSurcharge1 returns the CustomSurcharge1 field if non-nil, zero value otherwise.

### GetCustomSurcharge1Ok

`func (o *Invoice) GetCustomSurcharge1Ok() (*float32, bool)`

GetCustomSurcharge1Ok returns a tuple with the CustomSurcharge1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurcharge1

`func (o *Invoice) SetCustomSurcharge1(v float32)`

SetCustomSurcharge1 sets CustomSurcharge1 field to given value.

### HasCustomSurcharge1

`func (o *Invoice) HasCustomSurcharge1() bool`

HasCustomSurcharge1 returns a boolean if a field has been set.

### GetCustomSurcharge2

`func (o *Invoice) GetCustomSurcharge2() float32`

GetCustomSurcharge2 returns the CustomSurcharge2 field if non-nil, zero value otherwise.

### GetCustomSurcharge2Ok

`func (o *Invoice) GetCustomSurcharge2Ok() (*float32, bool)`

GetCustomSurcharge2Ok returns a tuple with the CustomSurcharge2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurcharge2

`func (o *Invoice) SetCustomSurcharge2(v float32)`

SetCustomSurcharge2 sets CustomSurcharge2 field to given value.

### HasCustomSurcharge2

`func (o *Invoice) HasCustomSurcharge2() bool`

HasCustomSurcharge2 returns a boolean if a field has been set.

### GetCustomSurcharge3

`func (o *Invoice) GetCustomSurcharge3() float32`

GetCustomSurcharge3 returns the CustomSurcharge3 field if non-nil, zero value otherwise.

### GetCustomSurcharge3Ok

`func (o *Invoice) GetCustomSurcharge3Ok() (*float32, bool)`

GetCustomSurcharge3Ok returns a tuple with the CustomSurcharge3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurcharge3

`func (o *Invoice) SetCustomSurcharge3(v float32)`

SetCustomSurcharge3 sets CustomSurcharge3 field to given value.

### HasCustomSurcharge3

`func (o *Invoice) HasCustomSurcharge3() bool`

HasCustomSurcharge3 returns a boolean if a field has been set.

### GetCustomSurcharge4

`func (o *Invoice) GetCustomSurcharge4() float32`

GetCustomSurcharge4 returns the CustomSurcharge4 field if non-nil, zero value otherwise.

### GetCustomSurcharge4Ok

`func (o *Invoice) GetCustomSurcharge4Ok() (*float32, bool)`

GetCustomSurcharge4Ok returns a tuple with the CustomSurcharge4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurcharge4

`func (o *Invoice) SetCustomSurcharge4(v float32)`

SetCustomSurcharge4 sets CustomSurcharge4 field to given value.

### HasCustomSurcharge4

`func (o *Invoice) HasCustomSurcharge4() bool`

HasCustomSurcharge4 returns a boolean if a field has been set.

### GetCustomSurchargeTax1

`func (o *Invoice) GetCustomSurchargeTax1() bool`

GetCustomSurchargeTax1 returns the CustomSurchargeTax1 field if non-nil, zero value otherwise.

### GetCustomSurchargeTax1Ok

`func (o *Invoice) GetCustomSurchargeTax1Ok() (*bool, bool)`

GetCustomSurchargeTax1Ok returns a tuple with the CustomSurchargeTax1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurchargeTax1

`func (o *Invoice) SetCustomSurchargeTax1(v bool)`

SetCustomSurchargeTax1 sets CustomSurchargeTax1 field to given value.

### HasCustomSurchargeTax1

`func (o *Invoice) HasCustomSurchargeTax1() bool`

HasCustomSurchargeTax1 returns a boolean if a field has been set.

### GetCustomSurchargeTax2

`func (o *Invoice) GetCustomSurchargeTax2() bool`

GetCustomSurchargeTax2 returns the CustomSurchargeTax2 field if non-nil, zero value otherwise.

### GetCustomSurchargeTax2Ok

`func (o *Invoice) GetCustomSurchargeTax2Ok() (*bool, bool)`

GetCustomSurchargeTax2Ok returns a tuple with the CustomSurchargeTax2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurchargeTax2

`func (o *Invoice) SetCustomSurchargeTax2(v bool)`

SetCustomSurchargeTax2 sets CustomSurchargeTax2 field to given value.

### HasCustomSurchargeTax2

`func (o *Invoice) HasCustomSurchargeTax2() bool`

HasCustomSurchargeTax2 returns a boolean if a field has been set.

### GetCustomSurchargeTax3

`func (o *Invoice) GetCustomSurchargeTax3() bool`

GetCustomSurchargeTax3 returns the CustomSurchargeTax3 field if non-nil, zero value otherwise.

### GetCustomSurchargeTax3Ok

`func (o *Invoice) GetCustomSurchargeTax3Ok() (*bool, bool)`

GetCustomSurchargeTax3Ok returns a tuple with the CustomSurchargeTax3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurchargeTax3

`func (o *Invoice) SetCustomSurchargeTax3(v bool)`

SetCustomSurchargeTax3 sets CustomSurchargeTax3 field to given value.

### HasCustomSurchargeTax3

`func (o *Invoice) HasCustomSurchargeTax3() bool`

HasCustomSurchargeTax3 returns a boolean if a field has been set.

### GetCustomSurchargeTax4

`func (o *Invoice) GetCustomSurchargeTax4() bool`

GetCustomSurchargeTax4 returns the CustomSurchargeTax4 field if non-nil, zero value otherwise.

### GetCustomSurchargeTax4Ok

`func (o *Invoice) GetCustomSurchargeTax4Ok() (*bool, bool)`

GetCustomSurchargeTax4Ok returns a tuple with the CustomSurchargeTax4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSurchargeTax4

`func (o *Invoice) SetCustomSurchargeTax4(v bool)`

SetCustomSurchargeTax4 sets CustomSurchargeTax4 field to given value.

### HasCustomSurchargeTax4

`func (o *Invoice) HasCustomSurchargeTax4() bool`

HasCustomSurchargeTax4 returns a boolean if a field has been set.

### GetProjectId

`func (o *Invoice) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Invoice) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Invoice) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Invoice) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetAutoBillTries

`func (o *Invoice) GetAutoBillTries() int32`

GetAutoBillTries returns the AutoBillTries field if non-nil, zero value otherwise.

### GetAutoBillTriesOk

`func (o *Invoice) GetAutoBillTriesOk() (*int32, bool)`

GetAutoBillTriesOk returns a tuple with the AutoBillTries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBillTries

`func (o *Invoice) SetAutoBillTries(v int32)`

SetAutoBillTries sets AutoBillTries field to given value.

### HasAutoBillTries

`func (o *Invoice) HasAutoBillTries() bool`

HasAutoBillTries returns a boolean if a field has been set.

### GetAutoBillEnabled

`func (o *Invoice) GetAutoBillEnabled() bool`

GetAutoBillEnabled returns the AutoBillEnabled field if non-nil, zero value otherwise.

### GetAutoBillEnabledOk

`func (o *Invoice) GetAutoBillEnabledOk() (*bool, bool)`

GetAutoBillEnabledOk returns a tuple with the AutoBillEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBillEnabled

`func (o *Invoice) SetAutoBillEnabled(v bool)`

SetAutoBillEnabled sets AutoBillEnabled field to given value.

### HasAutoBillEnabled

`func (o *Invoice) HasAutoBillEnabled() bool`

HasAutoBillEnabled returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *Invoice) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *Invoice) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *Invoice) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *Invoice) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetLocationId

`func (o *Invoice) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *Invoice) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *Invoice) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *Invoice) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


