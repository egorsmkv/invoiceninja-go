# Activity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id field of the activity | [optional] 
**ActivityTypeId** | Pointer to **string** | The activity type id | [optional] 
**ClientId** | Pointer to **string** | The client hashed id | [optional] 
**CompanyId** | Pointer to **string** | The company hashed id | [optional] 
**UserId** | Pointer to **string** | The user hashed id | [optional] 
**InvoiceId** | Pointer to **string** | The invoice hashed id | [optional] 
**PaymentId** | Pointer to **string** | The payment hashed id | [optional] 
**CreditId** | Pointer to **string** | The credit hashed id | [optional] 
**UpdatedAt** | Pointer to **int32** | Unixtimestamp the last time the record was updated | [optional] 
**ExpenseId** | Pointer to **string** | The expense hashed id | [optional] 
**IsSystem** | Pointer to **bool** | Defines is the activity was performed by the system | [optional] 
**ContactId** | Pointer to **string** | The contact hashed id | [optional] 
**TaskId** | Pointer to **string** | The task hashed id | [optional] 
**Notes** | Pointer to **string** | Activity Notes | [optional] 
**TokenId** | Pointer to **string** | The hashed ID of the token who performed the action | [optional] 
**Ip** | Pointer to **string** | The IP Address of the user who performed the action | [optional] 
**User** | Pointer to [**User**](User.md) |  | [optional] 
**Client** | Pointer to [**Client**](Client.md) |  | [optional] 
**Contact** | Pointer to [**ClientContact**](ClientContact.md) |  | [optional] 
**RecurringInvoice** | Pointer to [**RecurringInvoice**](RecurringInvoice.md) |  | [optional] 
**Invoice** | Pointer to [**Invoice**](Invoice.md) |  | [optional] 
**Credit** | Pointer to [**Credit**](Credit.md) |  | [optional] 
**Quote** | Pointer to [**Quote**](Quote.md) |  | [optional] 
**Payment** | Pointer to [**Payment**](Payment.md) |  | [optional] 
**Expense** | Pointer to [**Expense**](Expense.md) |  | [optional] 
**Task** | Pointer to [**Task**](Task.md) |  | [optional] 
**PurchaseOrder** | Pointer to [**PurchaseOrder**](PurchaseOrder.md) |  | [optional] 
**Vendor** | Pointer to [**Vendor**](Vendor.md) |  | [optional] 
**VendorContact** | Pointer to [**VendorContact**](VendorContact.md) |  | [optional] 

## Methods

### NewActivity

`func NewActivity() *Activity`

NewActivity instantiates a new Activity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityWithDefaults

`func NewActivityWithDefaults() *Activity`

NewActivityWithDefaults instantiates a new Activity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Activity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Activity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Activity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Activity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActivityTypeId

`func (o *Activity) GetActivityTypeId() string`

GetActivityTypeId returns the ActivityTypeId field if non-nil, zero value otherwise.

### GetActivityTypeIdOk

`func (o *Activity) GetActivityTypeIdOk() (*string, bool)`

GetActivityTypeIdOk returns a tuple with the ActivityTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityTypeId

`func (o *Activity) SetActivityTypeId(v string)`

SetActivityTypeId sets ActivityTypeId field to given value.

### HasActivityTypeId

`func (o *Activity) HasActivityTypeId() bool`

HasActivityTypeId returns a boolean if a field has been set.

### GetClientId

`func (o *Activity) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Activity) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Activity) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Activity) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetCompanyId

`func (o *Activity) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *Activity) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *Activity) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.

### HasCompanyId

`func (o *Activity) HasCompanyId() bool`

HasCompanyId returns a boolean if a field has been set.

### GetUserId

`func (o *Activity) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Activity) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Activity) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Activity) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *Activity) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *Activity) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *Activity) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *Activity) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetPaymentId

`func (o *Activity) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *Activity) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *Activity) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *Activity) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetCreditId

`func (o *Activity) GetCreditId() string`

GetCreditId returns the CreditId field if non-nil, zero value otherwise.

### GetCreditIdOk

`func (o *Activity) GetCreditIdOk() (*string, bool)`

GetCreditIdOk returns a tuple with the CreditId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditId

`func (o *Activity) SetCreditId(v string)`

SetCreditId sets CreditId field to given value.

### HasCreditId

`func (o *Activity) HasCreditId() bool`

HasCreditId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Activity) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Activity) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Activity) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Activity) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetExpenseId

`func (o *Activity) GetExpenseId() string`

GetExpenseId returns the ExpenseId field if non-nil, zero value otherwise.

### GetExpenseIdOk

`func (o *Activity) GetExpenseIdOk() (*string, bool)`

GetExpenseIdOk returns a tuple with the ExpenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseId

`func (o *Activity) SetExpenseId(v string)`

SetExpenseId sets ExpenseId field to given value.

### HasExpenseId

`func (o *Activity) HasExpenseId() bool`

HasExpenseId returns a boolean if a field has been set.

### GetIsSystem

`func (o *Activity) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *Activity) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *Activity) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.

### HasIsSystem

`func (o *Activity) HasIsSystem() bool`

HasIsSystem returns a boolean if a field has been set.

### GetContactId

`func (o *Activity) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *Activity) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *Activity) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *Activity) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### GetTaskId

`func (o *Activity) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *Activity) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *Activity) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *Activity) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetNotes

`func (o *Activity) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Activity) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Activity) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Activity) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetTokenId

`func (o *Activity) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *Activity) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *Activity) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *Activity) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetIp

`func (o *Activity) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *Activity) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *Activity) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *Activity) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetUser

`func (o *Activity) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Activity) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Activity) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *Activity) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetClient

`func (o *Activity) GetClient() Client`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *Activity) GetClientOk() (*Client, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *Activity) SetClient(v Client)`

SetClient sets Client field to given value.

### HasClient

`func (o *Activity) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetContact

`func (o *Activity) GetContact() ClientContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Activity) GetContactOk() (*ClientContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Activity) SetContact(v ClientContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *Activity) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetRecurringInvoice

`func (o *Activity) GetRecurringInvoice() RecurringInvoice`

GetRecurringInvoice returns the RecurringInvoice field if non-nil, zero value otherwise.

### GetRecurringInvoiceOk

`func (o *Activity) GetRecurringInvoiceOk() (*RecurringInvoice, bool)`

GetRecurringInvoiceOk returns a tuple with the RecurringInvoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringInvoice

`func (o *Activity) SetRecurringInvoice(v RecurringInvoice)`

SetRecurringInvoice sets RecurringInvoice field to given value.

### HasRecurringInvoice

`func (o *Activity) HasRecurringInvoice() bool`

HasRecurringInvoice returns a boolean if a field has been set.

### GetInvoice

`func (o *Activity) GetInvoice() Invoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *Activity) GetInvoiceOk() (*Invoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *Activity) SetInvoice(v Invoice)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *Activity) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetCredit

`func (o *Activity) GetCredit() Credit`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *Activity) GetCreditOk() (*Credit, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *Activity) SetCredit(v Credit)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *Activity) HasCredit() bool`

HasCredit returns a boolean if a field has been set.

### GetQuote

`func (o *Activity) GetQuote() Quote`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *Activity) GetQuoteOk() (*Quote, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *Activity) SetQuote(v Quote)`

SetQuote sets Quote field to given value.

### HasQuote

`func (o *Activity) HasQuote() bool`

HasQuote returns a boolean if a field has been set.

### GetPayment

`func (o *Activity) GetPayment() Payment`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *Activity) GetPaymentOk() (*Payment, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *Activity) SetPayment(v Payment)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *Activity) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetExpense

`func (o *Activity) GetExpense() Expense`

GetExpense returns the Expense field if non-nil, zero value otherwise.

### GetExpenseOk

`func (o *Activity) GetExpenseOk() (*Expense, bool)`

GetExpenseOk returns a tuple with the Expense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpense

`func (o *Activity) SetExpense(v Expense)`

SetExpense sets Expense field to given value.

### HasExpense

`func (o *Activity) HasExpense() bool`

HasExpense returns a boolean if a field has been set.

### GetTask

`func (o *Activity) GetTask() Task`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *Activity) GetTaskOk() (*Task, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *Activity) SetTask(v Task)`

SetTask sets Task field to given value.

### HasTask

`func (o *Activity) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetPurchaseOrder

`func (o *Activity) GetPurchaseOrder() PurchaseOrder`

GetPurchaseOrder returns the PurchaseOrder field if non-nil, zero value otherwise.

### GetPurchaseOrderOk

`func (o *Activity) GetPurchaseOrderOk() (*PurchaseOrder, bool)`

GetPurchaseOrderOk returns a tuple with the PurchaseOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrder

`func (o *Activity) SetPurchaseOrder(v PurchaseOrder)`

SetPurchaseOrder sets PurchaseOrder field to given value.

### HasPurchaseOrder

`func (o *Activity) HasPurchaseOrder() bool`

HasPurchaseOrder returns a boolean if a field has been set.

### GetVendor

`func (o *Activity) GetVendor() Vendor`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *Activity) GetVendorOk() (*Vendor, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *Activity) SetVendor(v Vendor)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *Activity) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVendorContact

`func (o *Activity) GetVendorContact() VendorContact`

GetVendorContact returns the VendorContact field if non-nil, zero value otherwise.

### GetVendorContactOk

`func (o *Activity) GetVendorContactOk() (*VendorContact, bool)`

GetVendorContactOk returns a tuple with the VendorContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorContact

`func (o *Activity) SetVendorContact(v VendorContact)`

SetVendorContact sets VendorContact field to given value.

### HasVendorContact

`func (o *Activity) HasVendorContact() bool`

HasVendorContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


