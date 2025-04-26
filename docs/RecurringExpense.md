# RecurringExpense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The hashed id of the recurring expense | [optional] 
**UserId** | Pointer to **string** | The hashed id of the user who created the recurring expense | [optional] 
**AssignedUserId** | Pointer to **string** | The hashed id of the user assigned to this recurring expense | [optional] 
**ClientId** | Pointer to **string** | The hashed id of the client | [optional] 
**InvoiceId** | Pointer to **string** | The hashed id of the invoice | [optional] 
**BankId** | Pointer to **string** | The id of the bank associated with this recurring expense | [optional] 
**InvoiceCurrencyId** | Pointer to **string** | The currency id of the invoice associated with this recurring expense | [optional] 
**ExpenseCurrencyId** | Pointer to **string** | The currency id of the expense associated with this recurring expense | [optional] 
**InvoiceCategoryId** | Pointer to **string** | The category id of the invoice | [optional] 
**PaymentTypeId** | Pointer to **string** | The payment type id | [optional] 
**PrivateNotes** | Pointer to **string** | The recurring expense private notes | [optional] 
**PublicNotes** | Pointer to **string** | The recurring expense public notes | [optional] 
**TransactionReference** | Pointer to **string** | The recurring expense transaction reference | [optional] 
**TranscationId** | Pointer to **string** | The transaction id of the recurring expense | [optional] 
**CustomValue1** | Pointer to **string** | Custom value field | [optional] 
**CustomValue2** | Pointer to **string** | Custom value field | [optional] 
**CustomValue3** | Pointer to **string** | Custom value field | [optional] 
**CustomValue4** | Pointer to **string** | Custom value field | [optional] 
**TaxName1** | Pointer to **string** | The tax name | [optional] 
**TaxName2** | Pointer to **string** | The tax name | [optional] 
**TaxRate1** | Pointer to **float32** | The tax rate | [optional] 
**TaxRate2** | Pointer to **float32** | The tax rate | [optional] 
**TaxName3** | Pointer to **string** | The tax name | [optional] 
**TaxRate3** | Pointer to **float32** | The tax rate | [optional] 
**Amount** | Pointer to **float32** | The total amount of the recurring expense | [optional] 
**FrequencyId** | Pointer to **float32** | The frequency this recurring expense fires | [optional] 
**RemainingCycles** | Pointer to **float32** | The number of remaining cycles for this recurring expense | [optional] 
**ForeignAmount** | Pointer to **float32** | The foreign currency amount of the recurring expense | [optional] 
**ExchangeRate** | Pointer to **float32** | The exchange rate for the expernse | [optional] 
**Date** | Pointer to **string** | The date of the expense | [optional] 
**PaymentDate** | Pointer to **string** | The date the expense was paid | [optional] 
**ShouldBeInvoiced** | Pointer to **bool** | Boolean flag determining if the expense should be invoiced | [optional] 
**IsDeleted** | Pointer to **bool** | Boolean flag determining if the recurring expense is deleted | [optional] 
**LastSentDate** | Pointer to **string** | The Date it was sent last | [optional] 
**NextSendDate** | Pointer to **string** | The next send date | [optional] 
**InvoiceDocuments** | Pointer to **bool** | Boolean flag determining if the documents associated with this expense should be passed onto the invoice if it is converted to an invoice | [optional] 
**UpdatedAt** | Pointer to **float32** | Timestamp | [optional] 
**ArchivedAt** | Pointer to **float32** | Timestamp | [optional] 

## Methods

### NewRecurringExpense

`func NewRecurringExpense() *RecurringExpense`

NewRecurringExpense instantiates a new RecurringExpense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecurringExpenseWithDefaults

`func NewRecurringExpenseWithDefaults() *RecurringExpense`

NewRecurringExpenseWithDefaults instantiates a new RecurringExpense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RecurringExpense) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RecurringExpense) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RecurringExpense) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RecurringExpense) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *RecurringExpense) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RecurringExpense) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RecurringExpense) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *RecurringExpense) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAssignedUserId

`func (o *RecurringExpense) GetAssignedUserId() string`

GetAssignedUserId returns the AssignedUserId field if non-nil, zero value otherwise.

### GetAssignedUserIdOk

`func (o *RecurringExpense) GetAssignedUserIdOk() (*string, bool)`

GetAssignedUserIdOk returns a tuple with the AssignedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedUserId

`func (o *RecurringExpense) SetAssignedUserId(v string)`

SetAssignedUserId sets AssignedUserId field to given value.

### HasAssignedUserId

`func (o *RecurringExpense) HasAssignedUserId() bool`

HasAssignedUserId returns a boolean if a field has been set.

### GetClientId

`func (o *RecurringExpense) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *RecurringExpense) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *RecurringExpense) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *RecurringExpense) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetInvoiceId

`func (o *RecurringExpense) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### GetInvoiceIdOk

`func (o *RecurringExpense) GetInvoiceIdOk() (*string, bool)`

GetInvoiceIdOk returns a tuple with the InvoiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceId

`func (o *RecurringExpense) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### HasInvoiceId

`func (o *RecurringExpense) HasInvoiceId() bool`

HasInvoiceId returns a boolean if a field has been set.

### GetBankId

`func (o *RecurringExpense) GetBankId() string`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *RecurringExpense) GetBankIdOk() (*string, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *RecurringExpense) SetBankId(v string)`

SetBankId sets BankId field to given value.

### HasBankId

`func (o *RecurringExpense) HasBankId() bool`

HasBankId returns a boolean if a field has been set.

### GetInvoiceCurrencyId

`func (o *RecurringExpense) GetInvoiceCurrencyId() string`

GetInvoiceCurrencyId returns the InvoiceCurrencyId field if non-nil, zero value otherwise.

### GetInvoiceCurrencyIdOk

`func (o *RecurringExpense) GetInvoiceCurrencyIdOk() (*string, bool)`

GetInvoiceCurrencyIdOk returns a tuple with the InvoiceCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceCurrencyId

`func (o *RecurringExpense) SetInvoiceCurrencyId(v string)`

SetInvoiceCurrencyId sets InvoiceCurrencyId field to given value.

### HasInvoiceCurrencyId

`func (o *RecurringExpense) HasInvoiceCurrencyId() bool`

HasInvoiceCurrencyId returns a boolean if a field has been set.

### GetExpenseCurrencyId

`func (o *RecurringExpense) GetExpenseCurrencyId() string`

GetExpenseCurrencyId returns the ExpenseCurrencyId field if non-nil, zero value otherwise.

### GetExpenseCurrencyIdOk

`func (o *RecurringExpense) GetExpenseCurrencyIdOk() (*string, bool)`

GetExpenseCurrencyIdOk returns a tuple with the ExpenseCurrencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseCurrencyId

`func (o *RecurringExpense) SetExpenseCurrencyId(v string)`

SetExpenseCurrencyId sets ExpenseCurrencyId field to given value.

### HasExpenseCurrencyId

`func (o *RecurringExpense) HasExpenseCurrencyId() bool`

HasExpenseCurrencyId returns a boolean if a field has been set.

### GetInvoiceCategoryId

`func (o *RecurringExpense) GetInvoiceCategoryId() string`

GetInvoiceCategoryId returns the InvoiceCategoryId field if non-nil, zero value otherwise.

### GetInvoiceCategoryIdOk

`func (o *RecurringExpense) GetInvoiceCategoryIdOk() (*string, bool)`

GetInvoiceCategoryIdOk returns a tuple with the InvoiceCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceCategoryId

`func (o *RecurringExpense) SetInvoiceCategoryId(v string)`

SetInvoiceCategoryId sets InvoiceCategoryId field to given value.

### HasInvoiceCategoryId

`func (o *RecurringExpense) HasInvoiceCategoryId() bool`

HasInvoiceCategoryId returns a boolean if a field has been set.

### GetPaymentTypeId

`func (o *RecurringExpense) GetPaymentTypeId() string`

GetPaymentTypeId returns the PaymentTypeId field if non-nil, zero value otherwise.

### GetPaymentTypeIdOk

`func (o *RecurringExpense) GetPaymentTypeIdOk() (*string, bool)`

GetPaymentTypeIdOk returns a tuple with the PaymentTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTypeId

`func (o *RecurringExpense) SetPaymentTypeId(v string)`

SetPaymentTypeId sets PaymentTypeId field to given value.

### HasPaymentTypeId

`func (o *RecurringExpense) HasPaymentTypeId() bool`

HasPaymentTypeId returns a boolean if a field has been set.

### GetPrivateNotes

`func (o *RecurringExpense) GetPrivateNotes() string`

GetPrivateNotes returns the PrivateNotes field if non-nil, zero value otherwise.

### GetPrivateNotesOk

`func (o *RecurringExpense) GetPrivateNotesOk() (*string, bool)`

GetPrivateNotesOk returns a tuple with the PrivateNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateNotes

`func (o *RecurringExpense) SetPrivateNotes(v string)`

SetPrivateNotes sets PrivateNotes field to given value.

### HasPrivateNotes

`func (o *RecurringExpense) HasPrivateNotes() bool`

HasPrivateNotes returns a boolean if a field has been set.

### GetPublicNotes

`func (o *RecurringExpense) GetPublicNotes() string`

GetPublicNotes returns the PublicNotes field if non-nil, zero value otherwise.

### GetPublicNotesOk

`func (o *RecurringExpense) GetPublicNotesOk() (*string, bool)`

GetPublicNotesOk returns a tuple with the PublicNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicNotes

`func (o *RecurringExpense) SetPublicNotes(v string)`

SetPublicNotes sets PublicNotes field to given value.

### HasPublicNotes

`func (o *RecurringExpense) HasPublicNotes() bool`

HasPublicNotes returns a boolean if a field has been set.

### GetTransactionReference

`func (o *RecurringExpense) GetTransactionReference() string`

GetTransactionReference returns the TransactionReference field if non-nil, zero value otherwise.

### GetTransactionReferenceOk

`func (o *RecurringExpense) GetTransactionReferenceOk() (*string, bool)`

GetTransactionReferenceOk returns a tuple with the TransactionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionReference

`func (o *RecurringExpense) SetTransactionReference(v string)`

SetTransactionReference sets TransactionReference field to given value.

### HasTransactionReference

`func (o *RecurringExpense) HasTransactionReference() bool`

HasTransactionReference returns a boolean if a field has been set.

### GetTranscationId

`func (o *RecurringExpense) GetTranscationId() string`

GetTranscationId returns the TranscationId field if non-nil, zero value otherwise.

### GetTranscationIdOk

`func (o *RecurringExpense) GetTranscationIdOk() (*string, bool)`

GetTranscationIdOk returns a tuple with the TranscationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscationId

`func (o *RecurringExpense) SetTranscationId(v string)`

SetTranscationId sets TranscationId field to given value.

### HasTranscationId

`func (o *RecurringExpense) HasTranscationId() bool`

HasTranscationId returns a boolean if a field has been set.

### GetCustomValue1

`func (o *RecurringExpense) GetCustomValue1() string`

GetCustomValue1 returns the CustomValue1 field if non-nil, zero value otherwise.

### GetCustomValue1Ok

`func (o *RecurringExpense) GetCustomValue1Ok() (*string, bool)`

GetCustomValue1Ok returns a tuple with the CustomValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue1

`func (o *RecurringExpense) SetCustomValue1(v string)`

SetCustomValue1 sets CustomValue1 field to given value.

### HasCustomValue1

`func (o *RecurringExpense) HasCustomValue1() bool`

HasCustomValue1 returns a boolean if a field has been set.

### GetCustomValue2

`func (o *RecurringExpense) GetCustomValue2() string`

GetCustomValue2 returns the CustomValue2 field if non-nil, zero value otherwise.

### GetCustomValue2Ok

`func (o *RecurringExpense) GetCustomValue2Ok() (*string, bool)`

GetCustomValue2Ok returns a tuple with the CustomValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue2

`func (o *RecurringExpense) SetCustomValue2(v string)`

SetCustomValue2 sets CustomValue2 field to given value.

### HasCustomValue2

`func (o *RecurringExpense) HasCustomValue2() bool`

HasCustomValue2 returns a boolean if a field has been set.

### GetCustomValue3

`func (o *RecurringExpense) GetCustomValue3() string`

GetCustomValue3 returns the CustomValue3 field if non-nil, zero value otherwise.

### GetCustomValue3Ok

`func (o *RecurringExpense) GetCustomValue3Ok() (*string, bool)`

GetCustomValue3Ok returns a tuple with the CustomValue3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue3

`func (o *RecurringExpense) SetCustomValue3(v string)`

SetCustomValue3 sets CustomValue3 field to given value.

### HasCustomValue3

`func (o *RecurringExpense) HasCustomValue3() bool`

HasCustomValue3 returns a boolean if a field has been set.

### GetCustomValue4

`func (o *RecurringExpense) GetCustomValue4() string`

GetCustomValue4 returns the CustomValue4 field if non-nil, zero value otherwise.

### GetCustomValue4Ok

`func (o *RecurringExpense) GetCustomValue4Ok() (*string, bool)`

GetCustomValue4Ok returns a tuple with the CustomValue4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValue4

`func (o *RecurringExpense) SetCustomValue4(v string)`

SetCustomValue4 sets CustomValue4 field to given value.

### HasCustomValue4

`func (o *RecurringExpense) HasCustomValue4() bool`

HasCustomValue4 returns a boolean if a field has been set.

### GetTaxName1

`func (o *RecurringExpense) GetTaxName1() string`

GetTaxName1 returns the TaxName1 field if non-nil, zero value otherwise.

### GetTaxName1Ok

`func (o *RecurringExpense) GetTaxName1Ok() (*string, bool)`

GetTaxName1Ok returns a tuple with the TaxName1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName1

`func (o *RecurringExpense) SetTaxName1(v string)`

SetTaxName1 sets TaxName1 field to given value.

### HasTaxName1

`func (o *RecurringExpense) HasTaxName1() bool`

HasTaxName1 returns a boolean if a field has been set.

### GetTaxName2

`func (o *RecurringExpense) GetTaxName2() string`

GetTaxName2 returns the TaxName2 field if non-nil, zero value otherwise.

### GetTaxName2Ok

`func (o *RecurringExpense) GetTaxName2Ok() (*string, bool)`

GetTaxName2Ok returns a tuple with the TaxName2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName2

`func (o *RecurringExpense) SetTaxName2(v string)`

SetTaxName2 sets TaxName2 field to given value.

### HasTaxName2

`func (o *RecurringExpense) HasTaxName2() bool`

HasTaxName2 returns a boolean if a field has been set.

### GetTaxRate1

`func (o *RecurringExpense) GetTaxRate1() float32`

GetTaxRate1 returns the TaxRate1 field if non-nil, zero value otherwise.

### GetTaxRate1Ok

`func (o *RecurringExpense) GetTaxRate1Ok() (*float32, bool)`

GetTaxRate1Ok returns a tuple with the TaxRate1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate1

`func (o *RecurringExpense) SetTaxRate1(v float32)`

SetTaxRate1 sets TaxRate1 field to given value.

### HasTaxRate1

`func (o *RecurringExpense) HasTaxRate1() bool`

HasTaxRate1 returns a boolean if a field has been set.

### GetTaxRate2

`func (o *RecurringExpense) GetTaxRate2() float32`

GetTaxRate2 returns the TaxRate2 field if non-nil, zero value otherwise.

### GetTaxRate2Ok

`func (o *RecurringExpense) GetTaxRate2Ok() (*float32, bool)`

GetTaxRate2Ok returns a tuple with the TaxRate2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate2

`func (o *RecurringExpense) SetTaxRate2(v float32)`

SetTaxRate2 sets TaxRate2 field to given value.

### HasTaxRate2

`func (o *RecurringExpense) HasTaxRate2() bool`

HasTaxRate2 returns a boolean if a field has been set.

### GetTaxName3

`func (o *RecurringExpense) GetTaxName3() string`

GetTaxName3 returns the TaxName3 field if non-nil, zero value otherwise.

### GetTaxName3Ok

`func (o *RecurringExpense) GetTaxName3Ok() (*string, bool)`

GetTaxName3Ok returns a tuple with the TaxName3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxName3

`func (o *RecurringExpense) SetTaxName3(v string)`

SetTaxName3 sets TaxName3 field to given value.

### HasTaxName3

`func (o *RecurringExpense) HasTaxName3() bool`

HasTaxName3 returns a boolean if a field has been set.

### GetTaxRate3

`func (o *RecurringExpense) GetTaxRate3() float32`

GetTaxRate3 returns the TaxRate3 field if non-nil, zero value otherwise.

### GetTaxRate3Ok

`func (o *RecurringExpense) GetTaxRate3Ok() (*float32, bool)`

GetTaxRate3Ok returns a tuple with the TaxRate3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate3

`func (o *RecurringExpense) SetTaxRate3(v float32)`

SetTaxRate3 sets TaxRate3 field to given value.

### HasTaxRate3

`func (o *RecurringExpense) HasTaxRate3() bool`

HasTaxRate3 returns a boolean if a field has been set.

### GetAmount

`func (o *RecurringExpense) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RecurringExpense) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RecurringExpense) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *RecurringExpense) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetFrequencyId

`func (o *RecurringExpense) GetFrequencyId() float32`

GetFrequencyId returns the FrequencyId field if non-nil, zero value otherwise.

### GetFrequencyIdOk

`func (o *RecurringExpense) GetFrequencyIdOk() (*float32, bool)`

GetFrequencyIdOk returns a tuple with the FrequencyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyId

`func (o *RecurringExpense) SetFrequencyId(v float32)`

SetFrequencyId sets FrequencyId field to given value.

### HasFrequencyId

`func (o *RecurringExpense) HasFrequencyId() bool`

HasFrequencyId returns a boolean if a field has been set.

### GetRemainingCycles

`func (o *RecurringExpense) GetRemainingCycles() float32`

GetRemainingCycles returns the RemainingCycles field if non-nil, zero value otherwise.

### GetRemainingCyclesOk

`func (o *RecurringExpense) GetRemainingCyclesOk() (*float32, bool)`

GetRemainingCyclesOk returns a tuple with the RemainingCycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingCycles

`func (o *RecurringExpense) SetRemainingCycles(v float32)`

SetRemainingCycles sets RemainingCycles field to given value.

### HasRemainingCycles

`func (o *RecurringExpense) HasRemainingCycles() bool`

HasRemainingCycles returns a boolean if a field has been set.

### GetForeignAmount

`func (o *RecurringExpense) GetForeignAmount() float32`

GetForeignAmount returns the ForeignAmount field if non-nil, zero value otherwise.

### GetForeignAmountOk

`func (o *RecurringExpense) GetForeignAmountOk() (*float32, bool)`

GetForeignAmountOk returns a tuple with the ForeignAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignAmount

`func (o *RecurringExpense) SetForeignAmount(v float32)`

SetForeignAmount sets ForeignAmount field to given value.

### HasForeignAmount

`func (o *RecurringExpense) HasForeignAmount() bool`

HasForeignAmount returns a boolean if a field has been set.

### GetExchangeRate

`func (o *RecurringExpense) GetExchangeRate() float32`

GetExchangeRate returns the ExchangeRate field if non-nil, zero value otherwise.

### GetExchangeRateOk

`func (o *RecurringExpense) GetExchangeRateOk() (*float32, bool)`

GetExchangeRateOk returns a tuple with the ExchangeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRate

`func (o *RecurringExpense) SetExchangeRate(v float32)`

SetExchangeRate sets ExchangeRate field to given value.

### HasExchangeRate

`func (o *RecurringExpense) HasExchangeRate() bool`

HasExchangeRate returns a boolean if a field has been set.

### GetDate

`func (o *RecurringExpense) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *RecurringExpense) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *RecurringExpense) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *RecurringExpense) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetPaymentDate

`func (o *RecurringExpense) GetPaymentDate() string`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *RecurringExpense) GetPaymentDateOk() (*string, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *RecurringExpense) SetPaymentDate(v string)`

SetPaymentDate sets PaymentDate field to given value.

### HasPaymentDate

`func (o *RecurringExpense) HasPaymentDate() bool`

HasPaymentDate returns a boolean if a field has been set.

### GetShouldBeInvoiced

`func (o *RecurringExpense) GetShouldBeInvoiced() bool`

GetShouldBeInvoiced returns the ShouldBeInvoiced field if non-nil, zero value otherwise.

### GetShouldBeInvoicedOk

`func (o *RecurringExpense) GetShouldBeInvoicedOk() (*bool, bool)`

GetShouldBeInvoicedOk returns a tuple with the ShouldBeInvoiced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldBeInvoiced

`func (o *RecurringExpense) SetShouldBeInvoiced(v bool)`

SetShouldBeInvoiced sets ShouldBeInvoiced field to given value.

### HasShouldBeInvoiced

`func (o *RecurringExpense) HasShouldBeInvoiced() bool`

HasShouldBeInvoiced returns a boolean if a field has been set.

### GetIsDeleted

`func (o *RecurringExpense) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *RecurringExpense) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *RecurringExpense) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *RecurringExpense) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetLastSentDate

`func (o *RecurringExpense) GetLastSentDate() string`

GetLastSentDate returns the LastSentDate field if non-nil, zero value otherwise.

### GetLastSentDateOk

`func (o *RecurringExpense) GetLastSentDateOk() (*string, bool)`

GetLastSentDateOk returns a tuple with the LastSentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSentDate

`func (o *RecurringExpense) SetLastSentDate(v string)`

SetLastSentDate sets LastSentDate field to given value.

### HasLastSentDate

`func (o *RecurringExpense) HasLastSentDate() bool`

HasLastSentDate returns a boolean if a field has been set.

### GetNextSendDate

`func (o *RecurringExpense) GetNextSendDate() string`

GetNextSendDate returns the NextSendDate field if non-nil, zero value otherwise.

### GetNextSendDateOk

`func (o *RecurringExpense) GetNextSendDateOk() (*string, bool)`

GetNextSendDateOk returns a tuple with the NextSendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSendDate

`func (o *RecurringExpense) SetNextSendDate(v string)`

SetNextSendDate sets NextSendDate field to given value.

### HasNextSendDate

`func (o *RecurringExpense) HasNextSendDate() bool`

HasNextSendDate returns a boolean if a field has been set.

### GetInvoiceDocuments

`func (o *RecurringExpense) GetInvoiceDocuments() bool`

GetInvoiceDocuments returns the InvoiceDocuments field if non-nil, zero value otherwise.

### GetInvoiceDocumentsOk

`func (o *RecurringExpense) GetInvoiceDocumentsOk() (*bool, bool)`

GetInvoiceDocumentsOk returns a tuple with the InvoiceDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDocuments

`func (o *RecurringExpense) SetInvoiceDocuments(v bool)`

SetInvoiceDocuments sets InvoiceDocuments field to given value.

### HasInvoiceDocuments

`func (o *RecurringExpense) HasInvoiceDocuments() bool`

HasInvoiceDocuments returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RecurringExpense) GetUpdatedAt() float32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RecurringExpense) GetUpdatedAtOk() (*float32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RecurringExpense) SetUpdatedAt(v float32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RecurringExpense) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetArchivedAt

`func (o *RecurringExpense) GetArchivedAt() float32`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *RecurringExpense) GetArchivedAtOk() (*float32, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *RecurringExpense) SetArchivedAt(v float32)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *RecurringExpense) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


