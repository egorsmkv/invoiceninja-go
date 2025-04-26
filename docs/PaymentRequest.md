# PaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | The client hashed id | [optional] 
**ClientContactId** | Pointer to **string** | The client contact hashed id | [optional] 
**UserId** | Pointer to **string** | The user hashed id | [optional] 
**TypeId** | Pointer to [**PaymentType**](PaymentType.md) |  | [optional] 
**Date** | Pointer to **string** | The Payment date | [optional] 
**TransactionReference** | Pointer to **string** | The transaction reference as defined by the payment gateway | [optional] 
**AssignedUserId** | Pointer to **string** | The assigned user hashed id | [optional] 
**PrivateNotes** | Pointer to **string** | The private notes of the payment | [optional] 
**Amount** | Pointer to **float32** | The amount of this payment | [optional] 
**Invoices** | Pointer to [**[]InvoicePaymentable**](InvoicePaymentable.md) |  | [optional] 
**Credits** | Pointer to [**[]CreditPaymentable**](CreditPaymentable.md) |  | [optional] 
**Number** | Pointer to **string** | The payment number - is a unique alpha numeric number per payment per company | [optional] 

## Methods

### NewPaymentRequest

`func NewPaymentRequest() *PaymentRequest`

NewPaymentRequest instantiates a new PaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRequestWithDefaults

`func NewPaymentRequestWithDefaults() *PaymentRequest`

NewPaymentRequestWithDefaults instantiates a new PaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *PaymentRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PaymentRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PaymentRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PaymentRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientContactId

`func (o *PaymentRequest) GetClientContactId() string`

GetClientContactId returns the ClientContactId field if non-nil, zero value otherwise.

### GetClientContactIdOk

`func (o *PaymentRequest) GetClientContactIdOk() (*string, bool)`

GetClientContactIdOk returns a tuple with the ClientContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientContactId

`func (o *PaymentRequest) SetClientContactId(v string)`

SetClientContactId sets ClientContactId field to given value.

### HasClientContactId

`func (o *PaymentRequest) HasClientContactId() bool`

HasClientContactId returns a boolean if a field has been set.

### GetUserId

`func (o *PaymentRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PaymentRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PaymentRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PaymentRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetTypeId

`func (o *PaymentRequest) GetTypeId() PaymentType`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *PaymentRequest) GetTypeIdOk() (*PaymentType, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *PaymentRequest) SetTypeId(v PaymentType)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *PaymentRequest) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetDate

`func (o *PaymentRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PaymentRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PaymentRequest) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *PaymentRequest) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetTransactionReference

`func (o *PaymentRequest) GetTransactionReference() string`

GetTransactionReference returns the TransactionReference field if non-nil, zero value otherwise.

### GetTransactionReferenceOk

`func (o *PaymentRequest) GetTransactionReferenceOk() (*string, bool)`

GetTransactionReferenceOk returns a tuple with the TransactionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionReference

`func (o *PaymentRequest) SetTransactionReference(v string)`

SetTransactionReference sets TransactionReference field to given value.

### HasTransactionReference

`func (o *PaymentRequest) HasTransactionReference() bool`

HasTransactionReference returns a boolean if a field has been set.

### GetAssignedUserId

`func (o *PaymentRequest) GetAssignedUserId() string`

GetAssignedUserId returns the AssignedUserId field if non-nil, zero value otherwise.

### GetAssignedUserIdOk

`func (o *PaymentRequest) GetAssignedUserIdOk() (*string, bool)`

GetAssignedUserIdOk returns a tuple with the AssignedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUserId

`func (o *PaymentRequest) SetAssignedUserId(v string)`

SetAssignedUserId sets AssignedUserId field to given value.

### HasAssignedUserId

`func (o *PaymentRequest) HasAssignedUserId() bool`

HasAssignedUserId returns a boolean if a field has been set.

### GetPrivateNotes

`func (o *PaymentRequest) GetPrivateNotes() string`

GetPrivateNotes returns the PrivateNotes field if non-nil, zero value otherwise.

### GetPrivateNotesOk

`func (o *PaymentRequest) GetPrivateNotesOk() (*string, bool)`

GetPrivateNotesOk returns a tuple with the PrivateNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNotes

`func (o *PaymentRequest) SetPrivateNotes(v string)`

SetPrivateNotes sets PrivateNotes field to given value.

### HasPrivateNotes

`func (o *PaymentRequest) HasPrivateNotes() bool`

HasPrivateNotes returns a boolean if a field has been set.

### GetAmount

`func (o *PaymentRequest) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentRequest) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentRequest) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PaymentRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetInvoices

`func (o *PaymentRequest) GetInvoices() []InvoicePaymentable`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *PaymentRequest) GetInvoicesOk() (*[]InvoicePaymentable, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *PaymentRequest) SetInvoices(v []InvoicePaymentable)`

SetInvoices sets Invoices field to given value.

### HasInvoices

`func (o *PaymentRequest) HasInvoices() bool`

HasInvoices returns a boolean if a field has been set.

### GetCredits

`func (o *PaymentRequest) GetCredits() []CreditPaymentable`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *PaymentRequest) GetCreditsOk() (*[]CreditPaymentable, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *PaymentRequest) SetCredits(v []CreditPaymentable)`

SetCredits sets Credits field to given value.

### HasCredits

`func (o *PaymentRequest) HasCredits() bool`

HasCredits returns a boolean if a field has been set.

### GetNumber

`func (o *PaymentRequest) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PaymentRequest) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PaymentRequest) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *PaymentRequest) HasNumber() bool`

HasNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


