# Payment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The payment hashed id | [optional] 
**ClientId** | Pointer to **string** | The client hashed id | [optional] 
**InvitationId** | Pointer to **string** | The invitation hashed id | [optional] 
**ClientContactId** | Pointer to **string** | The client contact hashed id | [optional] 
**UserId** | Pointer to **string** | The user hashed id | [optional] 
**TypeId** | Pointer to [**PaymentType**](PaymentType.md) |  | [optional] 
**Date** | Pointer to **string** | The Payment date | [optional] 
**TransactionReference** | Pointer to **string** | The transaction reference as defined by the payment gateway | [optional] 
**AssignedUserId** | Pointer to **string** | The assigned user hashed id | [optional] 
**PrivateNotes** | Pointer to **string** | The private notes of the payment | [optional] 
**IsManual** | Pointer to **bool** | Flags whether the payment was made manually or processed via a gateway | [optional] 
**IsDeleted** | Pointer to **bool** | Defines if the payment has been deleted | [optional] 
**Amount** | Pointer to **float32** | The amount of this payment | [optional] 
**Refunded** | Pointer to **float32** | The refunded amount of this payment | [optional] 
**UpdatedAt** | Pointer to **float32** | Timestamp | [optional] 
**ArchivedAt** | Pointer to **float32** | Timestamp | [optional] 
**CompanyGatewayId** | Pointer to **string** | The company gateway id | [optional] 
**Paymentables** | Pointer to [**Paymentable**](Paymentable.md) |  | [optional] 
**Invoices** | Pointer to [**[]InvoicePaymentable**](InvoicePaymentable.md) |  | [optional] 
**Credits** | Pointer to [**[]CreditPaymentable**](CreditPaymentable.md) |  | [optional] 
**Number** | Pointer to **string** | The payment number - is a unique alpha numeric number per payment per company | [optional] 

## Methods

### NewPayment

`func NewPayment() *Payment`

NewPayment instantiates a new Payment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentWithDefaults

`func NewPaymentWithDefaults() *Payment`

NewPaymentWithDefaults instantiates a new Payment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Payment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Payment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Payment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Payment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetClientId

`func (o *Payment) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Payment) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Payment) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Payment) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetInvitationId

`func (o *Payment) GetInvitationId() string`

GetInvitationId returns the InvitationId field if non-nil, zero value otherwise.

### GetInvitationIdOk

`func (o *Payment) GetInvitationIdOk() (*string, bool)`

GetInvitationIdOk returns a tuple with the InvitationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationId

`func (o *Payment) SetInvitationId(v string)`

SetInvitationId sets InvitationId field to given value.

### HasInvitationId

`func (o *Payment) HasInvitationId() bool`

HasInvitationId returns a boolean if a field has been set.

### GetClientContactId

`func (o *Payment) GetClientContactId() string`

GetClientContactId returns the ClientContactId field if non-nil, zero value otherwise.

### GetClientContactIdOk

`func (o *Payment) GetClientContactIdOk() (*string, bool)`

GetClientContactIdOk returns a tuple with the ClientContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientContactId

`func (o *Payment) SetClientContactId(v string)`

SetClientContactId sets ClientContactId field to given value.

### HasClientContactId

`func (o *Payment) HasClientContactId() bool`

HasClientContactId returns a boolean if a field has been set.

### GetUserId

`func (o *Payment) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Payment) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Payment) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Payment) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetTypeId

`func (o *Payment) GetTypeId() PaymentType`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *Payment) GetTypeIdOk() (*PaymentType, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *Payment) SetTypeId(v PaymentType)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *Payment) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetDate

`func (o *Payment) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Payment) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Payment) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *Payment) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetTransactionReference

`func (o *Payment) GetTransactionReference() string`

GetTransactionReference returns the TransactionReference field if non-nil, zero value otherwise.

### GetTransactionReferenceOk

`func (o *Payment) GetTransactionReferenceOk() (*string, bool)`

GetTransactionReferenceOk returns a tuple with the TransactionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionReference

`func (o *Payment) SetTransactionReference(v string)`

SetTransactionReference sets TransactionReference field to given value.

### HasTransactionReference

`func (o *Payment) HasTransactionReference() bool`

HasTransactionReference returns a boolean if a field has been set.

### GetAssignedUserId

`func (o *Payment) GetAssignedUserId() string`

GetAssignedUserId returns the AssignedUserId field if non-nil, zero value otherwise.

### GetAssignedUserIdOk

`func (o *Payment) GetAssignedUserIdOk() (*string, bool)`

GetAssignedUserIdOk returns a tuple with the AssignedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUserId

`func (o *Payment) SetAssignedUserId(v string)`

SetAssignedUserId sets AssignedUserId field to given value.

### HasAssignedUserId

`func (o *Payment) HasAssignedUserId() bool`

HasAssignedUserId returns a boolean if a field has been set.

### GetPrivateNotes

`func (o *Payment) GetPrivateNotes() string`

GetPrivateNotes returns the PrivateNotes field if non-nil, zero value otherwise.

### GetPrivateNotesOk

`func (o *Payment) GetPrivateNotesOk() (*string, bool)`

GetPrivateNotesOk returns a tuple with the PrivateNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNotes

`func (o *Payment) SetPrivateNotes(v string)`

SetPrivateNotes sets PrivateNotes field to given value.

### HasPrivateNotes

`func (o *Payment) HasPrivateNotes() bool`

HasPrivateNotes returns a boolean if a field has been set.

### GetIsManual

`func (o *Payment) GetIsManual() bool`

GetIsManual returns the IsManual field if non-nil, zero value otherwise.

### GetIsManualOk

`func (o *Payment) GetIsManualOk() (*bool, bool)`

GetIsManualOk returns a tuple with the IsManual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsManual

`func (o *Payment) SetIsManual(v bool)`

SetIsManual sets IsManual field to given value.

### HasIsManual

`func (o *Payment) HasIsManual() bool`

HasIsManual returns a boolean if a field has been set.

### GetIsDeleted

`func (o *Payment) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *Payment) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *Payment) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *Payment) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetAmount

`func (o *Payment) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Payment) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Payment) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Payment) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetRefunded

`func (o *Payment) GetRefunded() float32`

GetRefunded returns the Refunded field if non-nil, zero value otherwise.

### GetRefundedOk

`func (o *Payment) GetRefundedOk() (*float32, bool)`

GetRefundedOk returns a tuple with the Refunded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefunded

`func (o *Payment) SetRefunded(v float32)`

SetRefunded sets Refunded field to given value.

### HasRefunded

`func (o *Payment) HasRefunded() bool`

HasRefunded returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Payment) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Payment) GetUpdatedAtOk() (*float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Payment) SetUpdatedAt(v float32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Payment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetArchivedAt

`func (o *Payment) GetArchivedAt() float32`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *Payment) GetArchivedAtOk() (*float32, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *Payment) SetArchivedAt(v float32)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *Payment) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetCompanyGatewayId

`func (o *Payment) GetCompanyGatewayId() string`

GetCompanyGatewayId returns the CompanyGatewayId field if non-nil, zero value otherwise.

### GetCompanyGatewayIdOk

`func (o *Payment) GetCompanyGatewayIdOk() (*string, bool)`

GetCompanyGatewayIdOk returns a tuple with the CompanyGatewayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyGatewayId

`func (o *Payment) SetCompanyGatewayId(v string)`

SetCompanyGatewayId sets CompanyGatewayId field to given value.

### HasCompanyGatewayId

`func (o *Payment) HasCompanyGatewayId() bool`

HasCompanyGatewayId returns a boolean if a field has been set.

### GetPaymentables

`func (o *Payment) GetPaymentables() Paymentable`

GetPaymentables returns the Paymentables field if non-nil, zero value otherwise.

### GetPaymentablesOk

`func (o *Payment) GetPaymentablesOk() (*Paymentable, bool)`

GetPaymentablesOk returns a tuple with the Paymentables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentables

`func (o *Payment) SetPaymentables(v Paymentable)`

SetPaymentables sets Paymentables field to given value.

### HasPaymentables

`func (o *Payment) HasPaymentables() bool`

HasPaymentables returns a boolean if a field has been set.

### GetInvoices

`func (o *Payment) GetInvoices() []InvoicePaymentable`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *Payment) GetInvoicesOk() (*[]InvoicePaymentable, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *Payment) SetInvoices(v []InvoicePaymentable)`

SetInvoices sets Invoices field to given value.

### HasInvoices

`func (o *Payment) HasInvoices() bool`

HasInvoices returns a boolean if a field has been set.

### GetCredits

`func (o *Payment) GetCredits() []CreditPaymentable`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *Payment) GetCreditsOk() (*[]CreditPaymentable, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *Payment) SetCredits(v []CreditPaymentable)`

SetCredits sets Credits field to given value.

### HasCredits

`func (o *Payment) HasCredits() bool`

HasCredits returns a boolean if a field has been set.

### GetNumber

`func (o *Payment) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Payment) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Payment) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Payment) HasNumber() bool`

HasNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


