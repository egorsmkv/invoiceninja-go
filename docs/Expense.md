# Expense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The expense hashed id | [optional] 
**UserId** | Pointer to **string** | The user hashed id | [optional] 
**AssignedUserId** | Pointer to **string** | The assigned user hashed id | [optional] 
**ProjectId** | Pointer to **string** | The associated project_id | [optional] 
**ClientId** | Pointer to **string** | The client hashed id | [optional] 
**InvoiceId** | Pointer to **string** | The related invoice hashed id | [optional] 
**BankId** | Pointer to **string** | The bank id related to this expense | [optional] 
**InvoiceCurrencyId** | Pointer to **string** | The currency id of the related invoice | [optional] 
**CurrencyId** | Pointer to **string** | The currency id of the expense | [optional] 
**InvoiceCategoryId** | Pointer to **string** | The invoice category id | [optional] 
**PaymentTypeId** | Pointer to **string** | The payment type id | [optional] 
**RecurringExpenseId** | Pointer to **string** | The related recurring expense this expense was created from | [optional] 
**PrivateNotes** | Pointer to **string** | The private notes of the expense | [optional] 
**PublicNotes** | Pointer to **string** | The public notes of the expense | [optional] 
**TransactionReference** | Pointer to **string** | The transaction references of the expense | [optional] 
**TranscationId** | Pointer to **string** | The transaction id of the expense | [optional] 
**CustomValue1** | Pointer to **string** | A custom value | [optional] 
**CustomValue2** | Pointer to **string** | A custom value | [optional] 
**CustomValue3** | Pointer to **string** | A custom value | [optional] 
**CustomValue4** | Pointer to **string** | A custom value | [optional] 
**TaxAmount** | Pointer to **float32** | The tax amount | [optional] 
**TaxName1** | Pointer to **string** | Tax Name 1 | [optional] 
**TaxName2** | Pointer to **string** | Tax Name 2 | [optional] 
**TaxName3** | Pointer to **string** | Tax Name 3 | [optional] 
**TaxRate1** | Pointer to **float32** | Tax rate 1 | [optional] 
**TaxRate2** | Pointer to **float32** | Tax rate 2 | [optional] 
**TaxRate3** | Pointer to **float32** | Tax rate 3 | [optional] 
**Amount** | Pointer to **float32** | The total expense amont | [optional] 
**ForeignAmount** | Pointer to **float32** | The total foreign amount of the expense | [optional] 
**ExchangeRate** | Pointer to **float32** | The exchange rate at the time of the expense | [optional] 
**Date** | Pointer to **string** | The expense date format Y-m-d | [optional] 
**PaymentDate** | Pointer to **string** | The date of payment for the expense, format Y-m-d | [optional] 
**ShouldBeInvoiced** | Pointer to **bool** | Flag whether the expense should be invoiced | [optional] 
**IsDeleted** | Pointer to **bool** | Boolean determining whether the expense has been deleted | [optional] 
**InvoiceDocuments** | Pointer to **bool** | Passing the expense documents over to the invoice | [optional] 
**UpdatedAt** | Pointer to **float32** | Timestamp | [optional] 
**ArchivedAt** | Pointer to **float32** | Timestamp | [optional] 

## Methods

### NewExpense

`func NewExpense() *Expense`

NewExpense instantiates a new Expense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseWithDefaults

`func NewExpenseWithDefaults() *Expense`

NewExpenseWithDefaults instantiates a new Expense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Expense) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Expense) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Expense) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Expense) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *Expense) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Expense) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Expense) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Expense) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAssignedUserId

`func (o *Expense) GetAssignedUserId() string`

GetAssignedUserId returns the AssignedUserId field if non-nil, zero value otherwise.

### GetAssignedUserIdOk

`func (o *Expense) GetAssignedUserIdOk() (*string, bool)`

GetAssignedUserIdOk returns a tuple with the AssignedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUserId

`func (o *Expense) SetAssignedUserId(v string)`

SetAssignedUserId sets AssignedUserId field to given value.

### HasAssignedUserId

`func (o *Expense) HasAssignedUserId() bool`

HasAssignedUserId returns a boolean if a field has been set.

### GetProjectId

`func (o *Expense) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Expense) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Expense) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Expense) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetClientId

`func (o *Expense) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Expense) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Expense) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Expense) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *Expense) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *Expense) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *Expense) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *Expense) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetBankId

`func (o *Expense) GetBankId() string`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *Expense) GetBankIdOk() (*string, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *Expense) SetBankId(v string)`

SetBankId sets BankId field to given value.

### HasBankId

`func (o *Expense) HasBankId() bool`

HasBankId returns a boolean if a field has been set.

### GetInvoiceCurrencyId

`func (o *Expense) GetInvoiceCurrencyId() string`

GetInvoiceCurrencyId returns the InvoiceCurrencyId field if non-nil, zero value otherwise.

### GetInvoiceCurrencyIdOk

`func (o *Expense) GetInvoiceCurrencyIdOk() (*string, bool)`

GetInvoiceCurrencyIdOk returns a tuple with the InvoiceCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceCurrencyId

`func (o *Expense) SetInvoiceCurrencyId(v string)`

SetInvoiceCurrencyId sets InvoiceCurrencyId field to given value.

### HasInvoiceCurrencyId

`func (o *Expense) HasInvoiceCurrencyId() bool`

HasInvoiceCurrencyId returns a boolean if a field has been set.

### GetCurrencyId

`func (o *Expense) GetCurrencyId() string`

GetCurrencyId returns the CurrencyId field if non-nil, zero value otherwise.

### GetCurrencyIdOk

`func (o *Expense) GetCurrencyIdOk() (*string, bool)`

GetCurrencyIdOk returns a tuple with the CurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyId

`func (o *Expense) SetCurrencyId(v string)`

SetCurrencyId sets CurrencyId field to given value.

### HasCurrencyId

`func (o *Expense) HasCurrencyId() bool`

HasCurrencyId returns a boolean if a field has been set.

### GetInvoiceCategoryId

`func (o *Expense) GetInvoiceCategoryId() string`

GetInvoiceCategoryId returns the InvoiceCategoryId field if non-nil, zero value otherwise.

### GetInvoiceCategoryIdOk

`func (o *Expense) GetInvoiceCategoryIdOk() (*string, bool)`

GetInvoiceCategoryIdOk returns a tuple with the InvoiceCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceCategoryId

`func (o *Expense) SetInvoiceCategoryId(v string)`

SetInvoiceCategoryId sets InvoiceCategoryId field to given value.

### HasInvoiceCategoryId

`func (o *Expense) HasInvoiceCategoryId() bool`

HasInvoiceCategoryId returns a boolean if a field has been set.

### GetPaymentTypeId

`func (o *Expense) GetPaymentTypeId() string`

GetPaymentTypeId returns the PaymentTypeId field if non-nil, zero value otherwise.

### GetPaymentTypeIdOk

`func (o *Expense) GetPaymentTypeIdOk() (*string, bool)`

GetPaymentTypeIdOk returns a tuple with the PaymentTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTypeId

`func (o *Expense) SetPaymentTypeId(v string)`

SetPaymentTypeId sets PaymentTypeId field to given value.

### HasPaymentTypeId

`func (o *Expense) HasPaymentTypeId() bool`

HasPaymentTypeId returns a boolean if a field has been set.

### GetRecurringExpenseId

`func (o *Expense) GetRecurringExpenseId() string`

GetRecurringExpenseId returns the RecurringExpenseId field if non-nil, zero value otherwise.

### GetRecurringExpenseIdOk

`func (o *Expense) GetRecurringExpenseIdOk() (*string, bool)`

GetRecurringExpenseIdOk returns a tuple with the RecurringExpenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringExpenseId

`func (o *Expense) SetRecurringExpenseId(v string)`

SetRecurringExpenseId sets RecurringExpenseId field to given value.

### HasRecurringExpenseId

`func (o *Expense) HasRecurringExpenseId() bool`

HasRecurringExpenseId returns a boolean if a field has been set.

### GetPrivateNotes

`func (o *Expense) GetPrivateNotes() string`

GetPrivateNotes returns the PrivateNotes field if non-nil, zero value otherwise.

### GetPrivateNotesOk

`func (o *Expense) GetPrivateNotesOk() (*string, bool)`

GetPrivateNotesOk returns a tuple with the PrivateNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNotes

`func (o *Expense) SetPrivateNotes(v string)`

SetPrivateNotes sets PrivateNotes field to given value.

### HasPrivateNotes

`func (o *Expense) HasPrivateNotes() bool`

HasPrivateNotes returns a boolean if a field has been set.

### GetPublicNotes

`func (o *Expense) GetPublicNotes() string`

GetPublicNotes returns the PublicNotes field if non-nil, zero value otherwise.

### GetPublicNotesOk

`func (o *Expense) GetPublicNotesOk() (*string, bool)`

GetPublicNotesOk returns a tuple with the PublicNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNotes

`func (o *Expense) SetPublicNotes(v string)`

SetPublicNotes sets PublicNotes field to given value.

### HasPublicNotes

`func (o *Expense) HasPublicNotes() bool`

HasPublicNotes returns a boolean if a field has been set.

### GetTransactionReference

`func (o *Expense) GetTransactionReference() string`

GetTransactionReference returns the TransactionReference field if non-nil, zero value otherwise.

### GetTransactionReferenceOk

`func (o *Expense) GetTransactionReferenceOk() (*string, bool)`

GetTransactionReferenceOk returns a tuple with the TransactionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionReference

`func (o *Expense) SetTransactionReference(v string)`

SetTransactionReference sets TransactionReference field to given value.

### HasTransactionReference

`func (o *Expense) HasTransactionReference() bool`

HasTransactionReference returns a boolean if a field has been set.

### GetTranscationId

`func (o *Expense) GetTranscationId() string`

GetTranscationId returns the TranscationId field if non-nil, zero value otherwise.

### GetTranscationIdOk

`func (o *Expense) GetTranscationIdOk() (*string, bool)`

GetTranscationIdOk returns a tuple with the TranscationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscationId

`func (o *Expense) SetTranscationId(v string)`

SetTranscationId sets TranscationId field to given value.

### HasTranscationId

`func (o *Expense) HasTranscationId() bool`

HasTranscationId returns a boolean if a field has been set.

### GetCustomValue1

`func (o *Expense) GetCustomValue1() string`

GetCustomValue1 returns the CustomValue1 field if non-nil, zero value otherwise.

### GetCustomValue1Ok

`func (o *Expense) GetCustomValue1Ok() (*string, bool)`

GetCustomValue1Ok returns a tuple with the CustomValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue1

`func (o *Expense) SetCustomValue1(v string)`

SetCustomValue1 sets CustomValue1 field to given value.

### HasCustomValue1

`func (o *Expense) HasCustomValue1() bool`

HasCustomValue1 returns a boolean if a field has been set.

### GetCustomValue2

`func (o *Expense) GetCustomValue2() string`

GetCustomValue2 returns the CustomValue2 field if non-nil, zero value otherwise.

### GetCustomValue2Ok

`func (o *Expense) GetCustomValue2Ok() (*string, bool)`

GetCustomValue2Ok returns a tuple with the CustomValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue2

`func (o *Expense) SetCustomValue2(v string)`

SetCustomValue2 sets CustomValue2 field to given value.

### HasCustomValue2

`func (o *Expense) HasCustomValue2() bool`

HasCustomValue2 returns a boolean if a field has been set.

### GetCustomValue3

`func (o *Expense) GetCustomValue3() string`

GetCustomValue3 returns the CustomValue3 field if non-nil, zero value otherwise.

### GetCustomValue3Ok

`func (o *Expense) GetCustomValue3Ok() (*string, bool)`

GetCustomValue3Ok returns a tuple with the CustomValue3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue3

`func (o *Expense) SetCustomValue3(v string)`

SetCustomValue3 sets CustomValue3 field to given value.

### HasCustomValue3

`func (o *Expense) HasCustomValue3() bool`

HasCustomValue3 returns a boolean if a field has been set.

### GetCustomValue4

`func (o *Expense) GetCustomValue4() string`

GetCustomValue4 returns the CustomValue4 field if non-nil, zero value otherwise.

### GetCustomValue4Ok

`func (o *Expense) GetCustomValue4Ok() (*string, bool)`

GetCustomValue4Ok returns a tuple with the CustomValue4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue4

`func (o *Expense) SetCustomValue4(v string)`

SetCustomValue4 sets CustomValue4 field to given value.

### HasCustomValue4

`func (o *Expense) HasCustomValue4() bool`

HasCustomValue4 returns a boolean if a field has been set.

### GetTaxAmount

`func (o *Expense) GetTaxAmount() float32`

GetTaxAmount returns the TaxAmount field if non-nil, zero value otherwise.

### GetTaxAmountOk

`func (o *Expense) GetTaxAmountOk() (*float32, bool)`

GetTaxAmountOk returns a tuple with the TaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxAmount

`func (o *Expense) SetTaxAmount(v float32)`

SetTaxAmount sets TaxAmount field to given value.

### HasTaxAmount

`func (o *Expense) HasTaxAmount() bool`

HasTaxAmount returns a boolean if a field has been set.

### GetTaxName1

`func (o *Expense) GetTaxName1() string`

GetTaxName1 returns the TaxName1 field if non-nil, zero value otherwise.

### GetTaxName1Ok

`func (o *Expense) GetTaxName1Ok() (*string, bool)`

GetTaxName1Ok returns a tuple with the TaxName1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName1

`func (o *Expense) SetTaxName1(v string)`

SetTaxName1 sets TaxName1 field to given value.

### HasTaxName1

`func (o *Expense) HasTaxName1() bool`

HasTaxName1 returns a boolean if a field has been set.

### GetTaxName2

`func (o *Expense) GetTaxName2() string`

GetTaxName2 returns the TaxName2 field if non-nil, zero value otherwise.

### GetTaxName2Ok

`func (o *Expense) GetTaxName2Ok() (*string, bool)`

GetTaxName2Ok returns a tuple with the TaxName2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName2

`func (o *Expense) SetTaxName2(v string)`

SetTaxName2 sets TaxName2 field to given value.

### HasTaxName2

`func (o *Expense) HasTaxName2() bool`

HasTaxName2 returns a boolean if a field has been set.

### GetTaxName3

`func (o *Expense) GetTaxName3() string`

GetTaxName3 returns the TaxName3 field if non-nil, zero value otherwise.

### GetTaxName3Ok

`func (o *Expense) GetTaxName3Ok() (*string, bool)`

GetTaxName3Ok returns a tuple with the TaxName3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName3

`func (o *Expense) SetTaxName3(v string)`

SetTaxName3 sets TaxName3 field to given value.

### HasTaxName3

`func (o *Expense) HasTaxName3() bool`

HasTaxName3 returns a boolean if a field has been set.

### GetTaxRate1

`func (o *Expense) GetTaxRate1() float32`

GetTaxRate1 returns the TaxRate1 field if non-nil, zero value otherwise.

### GetTaxRate1Ok

`func (o *Expense) GetTaxRate1Ok() (*float32, bool)`

GetTaxRate1Ok returns a tuple with the TaxRate1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate1

`func (o *Expense) SetTaxRate1(v float32)`

SetTaxRate1 sets TaxRate1 field to given value.

### HasTaxRate1

`func (o *Expense) HasTaxRate1() bool`

HasTaxRate1 returns a boolean if a field has been set.

### GetTaxRate2

`func (o *Expense) GetTaxRate2() float32`

GetTaxRate2 returns the TaxRate2 field if non-nil, zero value otherwise.

### GetTaxRate2Ok

`func (o *Expense) GetTaxRate2Ok() (*float32, bool)`

GetTaxRate2Ok returns a tuple with the TaxRate2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate2

`func (o *Expense) SetTaxRate2(v float32)`

SetTaxRate2 sets TaxRate2 field to given value.

### HasTaxRate2

`func (o *Expense) HasTaxRate2() bool`

HasTaxRate2 returns a boolean if a field has been set.

### GetTaxRate3

`func (o *Expense) GetTaxRate3() float32`

GetTaxRate3 returns the TaxRate3 field if non-nil, zero value otherwise.

### GetTaxRate3Ok

`func (o *Expense) GetTaxRate3Ok() (*float32, bool)`

GetTaxRate3Ok returns a tuple with the TaxRate3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate3

`func (o *Expense) SetTaxRate3(v float32)`

SetTaxRate3 sets TaxRate3 field to given value.

### HasTaxRate3

`func (o *Expense) HasTaxRate3() bool`

HasTaxRate3 returns a boolean if a field has been set.

### GetAmount

`func (o *Expense) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Expense) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Expense) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Expense) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetForeignAmount

`func (o *Expense) GetForeignAmount() float32`

GetForeignAmount returns the ForeignAmount field if non-nil, zero value otherwise.

### GetForeignAmountOk

`func (o *Expense) GetForeignAmountOk() (*float32, bool)`

GetForeignAmountOk returns a tuple with the ForeignAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignAmount

`func (o *Expense) SetForeignAmount(v float32)`

SetForeignAmount sets ForeignAmount field to given value.

### HasForeignAmount

`func (o *Expense) HasForeignAmount() bool`

HasForeignAmount returns a boolean if a field has been set.

### GetExchangeRate

`func (o *Expense) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *Expense) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *Expense) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *Expense) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetDate

`func (o *Expense) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Expense) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Expense) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *Expense) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetPaymentDate

`func (o *Expense) GetPaymentDate() string`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *Expense) GetPaymentDateOk() (*string, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *Expense) SetPaymentDate(v string)`

SetPaymentDate sets PaymentDate field to given value.

### HasPaymentDate

`func (o *Expense) HasPaymentDate() bool`

HasPaymentDate returns a boolean if a field has been set.

### GetShouldBeInvoiced

`func (o *Expense) GetShouldBeInvoiced() bool`

GetShouldBeInvoiced returns the ShouldBeInvoiced field if non-nil, zero value otherwise.

### GetShouldBeInvoicedOk

`func (o *Expense) GetShouldBeInvoicedOk() (*bool, bool)`

GetShouldBeInvoicedOk returns a tuple with the ShouldBeInvoiced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldBeInvoiced

`func (o *Expense) SetShouldBeInvoiced(v bool)`

SetShouldBeInvoiced sets ShouldBeInvoiced field to given value.

### HasShouldBeInvoiced

`func (o *Expense) HasShouldBeInvoiced() bool`

HasShouldBeInvoiced returns a boolean if a field has been set.

### GetIsDeleted

`func (o *Expense) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *Expense) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *Expense) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *Expense) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetInvoiceDocuments

`func (o *Expense) GetInvoiceDocuments() bool`

GetInvoiceDocuments returns the InvoiceDocuments field if non-nil, zero value otherwise.

### GetInvoiceDocumentsOk

`func (o *Expense) GetInvoiceDocumentsOk() (*bool, bool)`

GetInvoiceDocumentsOk returns a tuple with the InvoiceDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDocuments

`func (o *Expense) SetInvoiceDocuments(v bool)`

SetInvoiceDocuments sets InvoiceDocuments field to given value.

### HasInvoiceDocuments

`func (o *Expense) HasInvoiceDocuments() bool`

HasInvoiceDocuments returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Expense) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Expense) GetUpdatedAtOk() (*float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Expense) SetUpdatedAt(v float32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Expense) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetArchivedAt

`func (o *Expense) GetArchivedAt() float32`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *Expense) GetArchivedAtOk() (*float32, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *Expense) SetArchivedAt(v float32)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *Expense) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


