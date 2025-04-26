package openapi

import (
	"encoding/json"
)

// checks if the RecurringExpense type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecurringExpense{}

// RecurringExpense struct for RecurringExpense
type RecurringExpense struct {
	// The hashed id of the recurring expense
	Id *string `json:"id,omitempty"`
	// The hashed id of the user who created the recurring expense
	UserId *string `json:"user_id,omitempty"`
	// The hashed id of the user assigned to this recurring expense
	AssignedUserId *string `json:"assigned_user_id,omitempty"`
	// The hashed id of the client
	ClientId *string `json:"client_id,omitempty"`
	// The hashed id of the invoice
	InvoiceId *string `json:"invoice_id,omitempty"`
	// The id of the bank associated with this recurring expense
	BankId *string `json:"bank_id,omitempty"`
	// The currency id of the invoice associated with this recurring expense
	InvoiceCurrencyId *string `json:"invoice_currency_id,omitempty"`
	// The currency id of the expense associated with this recurring expense
	ExpenseCurrencyId *string `json:"expense_currency_id,omitempty"`
	// The category id of the invoice
	InvoiceCategoryId *string `json:"invoice_category_id,omitempty"`
	// The payment type id
	PaymentTypeId *string `json:"payment_type_id,omitempty"`
	// The recurring expense private notes
	PrivateNotes *string `json:"private_notes,omitempty"`
	// The recurring expense public notes
	PublicNotes *string `json:"public_notes,omitempty"`
	// The recurring expense transaction reference
	TransactionReference *string `json:"transaction_reference,omitempty"`
	// The transaction id of the recurring expense
	TranscationId *string `json:"transcation_id,omitempty"`
	// Custom value field
	CustomValue1 *string `json:"custom_value1,omitempty"`
	// Custom value field
	CustomValue2 *string `json:"custom_value2,omitempty"`
	// Custom value field
	CustomValue3 *string `json:"custom_value3,omitempty"`
	// Custom value field
	CustomValue4 *string `json:"custom_value4,omitempty"`
	// The tax name
	TaxName1 *string `json:"tax_name1,omitempty"`
	// The tax name
	TaxName2 *string `json:"tax_name2,omitempty"`
	// The tax rate
	TaxRate1 *float32 `json:"tax_rate1,omitempty"`
	// The tax rate
	TaxRate2 *float32 `json:"tax_rate2,omitempty"`
	// The tax name
	TaxName3 *string `json:"tax_name3,omitempty"`
	// The tax rate
	TaxRate3 *float32 `json:"tax_rate3,omitempty"`
	// The total amount of the recurring expense
	Amount *float32 `json:"amount,omitempty"`
	// The frequency this recurring expense fires
	FrequencyId *float32 `json:"frequency_id,omitempty"`
	// The number of remaining cycles for this recurring expense
	RemainingCycles *float32 `json:"remaining_cycles,omitempty"`
	// The foreign currency amount of the recurring expense
	ForeignAmount *float32 `json:"foreign_amount,omitempty"`
	// The exchange rate for the expernse
	ExchangeRate *float32 `json:"exchange_rate,omitempty"`
	// The date of the expense
	Date *string `json:"date,omitempty"`
	// The date the expense was paid
	PaymentDate *string `json:"payment_date,omitempty"`
	// Boolean flag determining if the expense should be invoiced
	ShouldBeInvoiced *bool `json:"should_be_invoiced,omitempty"`
	// Boolean flag determining if the recurring expense is deleted
	IsDeleted *bool `json:"is_deleted,omitempty"`
	// The Date it was sent last
	LastSentDate *string `json:"last_sent_date,omitempty"`
	// The next send date
	NextSendDate *string `json:"next_send_date,omitempty"`
	// Boolean flag determining if the documents associated with this expense should be passed onto the invoice if it is converted to an invoice
	InvoiceDocuments *bool `json:"invoice_documents,omitempty"`
	// Timestamp
	UpdatedAt *float32 `json:"updated_at,omitempty"`
	// Timestamp
	ArchivedAt *float32 `json:"archived_at,omitempty"`
}

// NewRecurringExpense instantiates a new RecurringExpense object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecurringExpense() *RecurringExpense {
	this := RecurringExpense{}
	return &this
}

// NewRecurringExpenseWithDefaults instantiates a new RecurringExpense object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecurringExpenseWithDefaults() *RecurringExpense {
	this := RecurringExpense{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RecurringExpense) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RecurringExpense) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RecurringExpense) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *RecurringExpense) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *RecurringExpense) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *RecurringExpense) SetUserId(v string) {
	o.UserId = &v
}

// GetAssignedUserId returns the AssignedUserId field value if set, zero value otherwise.
func (o *RecurringExpense) GetAssignedUserId() string {
	if o == nil || IsNil(o.AssignedUserId) {
		var ret string
		return ret
	}
	return *o.AssignedUserId
}

// GetAssignedUserIdOk returns a tuple with the AssignedUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetAssignedUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssignedUserId) {
		return nil, false
	}
	return o.AssignedUserId, true
}

// HasAssignedUserId returns a boolean if a field has been set.
func (o *RecurringExpense) HasAssignedUserId() bool {
	if o != nil && !IsNil(o.AssignedUserId) {
		return true
	}

	return false
}

// SetAssignedUserId gets a reference to the given string and assigns it to the AssignedUserId field.
func (o *RecurringExpense) SetAssignedUserId(v string) {
	o.AssignedUserId = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *RecurringExpense) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *RecurringExpense) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *RecurringExpense) SetClientId(v string) {
	o.ClientId = &v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise.
func (o *RecurringExpense) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceId) {
		return nil, false
	}
	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *RecurringExpense) HasInvoiceId() bool {
	if o != nil && !IsNil(o.InvoiceId) {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *RecurringExpense) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

// GetBankId returns the BankId field value if set, zero value otherwise.
func (o *RecurringExpense) GetBankId() string {
	if o == nil || IsNil(o.BankId) {
		var ret string
		return ret
	}
	return *o.BankId
}

// GetBankIdOk returns a tuple with the BankId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetBankIdOk() (*string, bool) {
	if o == nil || IsNil(o.BankId) {
		return nil, false
	}
	return o.BankId, true
}

// HasBankId returns a boolean if a field has been set.
func (o *RecurringExpense) HasBankId() bool {
	if o != nil && !IsNil(o.BankId) {
		return true
	}

	return false
}

// SetBankId gets a reference to the given string and assigns it to the BankId field.
func (o *RecurringExpense) SetBankId(v string) {
	o.BankId = &v
}

// GetInvoiceCurrencyId returns the InvoiceCurrencyId field value if set, zero value otherwise.
func (o *RecurringExpense) GetInvoiceCurrencyId() string {
	if o == nil || IsNil(o.InvoiceCurrencyId) {
		var ret string
		return ret
	}
	return *o.InvoiceCurrencyId
}

// GetInvoiceCurrencyIdOk returns a tuple with the InvoiceCurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetInvoiceCurrencyIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceCurrencyId) {
		return nil, false
	}
	return o.InvoiceCurrencyId, true
}

// HasInvoiceCurrencyId returns a boolean if a field has been set.
func (o *RecurringExpense) HasInvoiceCurrencyId() bool {
	if o != nil && !IsNil(o.InvoiceCurrencyId) {
		return true
	}

	return false
}

// SetInvoiceCurrencyId gets a reference to the given string and assigns it to the InvoiceCurrencyId field.
func (o *RecurringExpense) SetInvoiceCurrencyId(v string) {
	o.InvoiceCurrencyId = &v
}

// GetExpenseCurrencyId returns the ExpenseCurrencyId field value if set, zero value otherwise.
func (o *RecurringExpense) GetExpenseCurrencyId() string {
	if o == nil || IsNil(o.ExpenseCurrencyId) {
		var ret string
		return ret
	}
	return *o.ExpenseCurrencyId
}

// GetExpenseCurrencyIdOk returns a tuple with the ExpenseCurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetExpenseCurrencyIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExpenseCurrencyId) {
		return nil, false
	}
	return o.ExpenseCurrencyId, true
}

// HasExpenseCurrencyId returns a boolean if a field has been set.
func (o *RecurringExpense) HasExpenseCurrencyId() bool {
	if o != nil && !IsNil(o.ExpenseCurrencyId) {
		return true
	}

	return false
}

// SetExpenseCurrencyId gets a reference to the given string and assigns it to the ExpenseCurrencyId field.
func (o *RecurringExpense) SetExpenseCurrencyId(v string) {
	o.ExpenseCurrencyId = &v
}

// GetInvoiceCategoryId returns the InvoiceCategoryId field value if set, zero value otherwise.
func (o *RecurringExpense) GetInvoiceCategoryId() string {
	if o == nil || IsNil(o.InvoiceCategoryId) {
		var ret string
		return ret
	}
	return *o.InvoiceCategoryId
}

// GetInvoiceCategoryIdOk returns a tuple with the InvoiceCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetInvoiceCategoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceCategoryId) {
		return nil, false
	}
	return o.InvoiceCategoryId, true
}

// HasInvoiceCategoryId returns a boolean if a field has been set.
func (o *RecurringExpense) HasInvoiceCategoryId() bool {
	if o != nil && !IsNil(o.InvoiceCategoryId) {
		return true
	}

	return false
}

// SetInvoiceCategoryId gets a reference to the given string and assigns it to the InvoiceCategoryId field.
func (o *RecurringExpense) SetInvoiceCategoryId(v string) {
	o.InvoiceCategoryId = &v
}

// GetPaymentTypeId returns the PaymentTypeId field value if set, zero value otherwise.
func (o *RecurringExpense) GetPaymentTypeId() string {
	if o == nil || IsNil(o.PaymentTypeId) {
		var ret string
		return ret
	}
	return *o.PaymentTypeId
}

// GetPaymentTypeIdOk returns a tuple with the PaymentTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetPaymentTypeIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentTypeId) {
		return nil, false
	}
	return o.PaymentTypeId, true
}

// HasPaymentTypeId returns a boolean if a field has been set.
func (o *RecurringExpense) HasPaymentTypeId() bool {
	if o != nil && !IsNil(o.PaymentTypeId) {
		return true
	}

	return false
}

// SetPaymentTypeId gets a reference to the given string and assigns it to the PaymentTypeId field.
func (o *RecurringExpense) SetPaymentTypeId(v string) {
	o.PaymentTypeId = &v
}

// GetPrivateNotes returns the PrivateNotes field value if set, zero value otherwise.
func (o *RecurringExpense) GetPrivateNotes() string {
	if o == nil || IsNil(o.PrivateNotes) {
		var ret string
		return ret
	}
	return *o.PrivateNotes
}

// GetPrivateNotesOk returns a tuple with the PrivateNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetPrivateNotesOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateNotes) {
		return nil, false
	}
	return o.PrivateNotes, true
}

// HasPrivateNotes returns a boolean if a field has been set.
func (o *RecurringExpense) HasPrivateNotes() bool {
	if o != nil && !IsNil(o.PrivateNotes) {
		return true
	}

	return false
}

// SetPrivateNotes gets a reference to the given string and assigns it to the PrivateNotes field.
func (o *RecurringExpense) SetPrivateNotes(v string) {
	o.PrivateNotes = &v
}

// GetPublicNotes returns the PublicNotes field value if set, zero value otherwise.
func (o *RecurringExpense) GetPublicNotes() string {
	if o == nil || IsNil(o.PublicNotes) {
		var ret string
		return ret
	}
	return *o.PublicNotes
}

// GetPublicNotesOk returns a tuple with the PublicNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetPublicNotesOk() (*string, bool) {
	if o == nil || IsNil(o.PublicNotes) {
		return nil, false
	}
	return o.PublicNotes, true
}

// HasPublicNotes returns a boolean if a field has been set.
func (o *RecurringExpense) HasPublicNotes() bool {
	if o != nil && !IsNil(o.PublicNotes) {
		return true
	}

	return false
}

// SetPublicNotes gets a reference to the given string and assigns it to the PublicNotes field.
func (o *RecurringExpense) SetPublicNotes(v string) {
	o.PublicNotes = &v
}

// GetTransactionReference returns the TransactionReference field value if set, zero value otherwise.
func (o *RecurringExpense) GetTransactionReference() string {
	if o == nil || IsNil(o.TransactionReference) {
		var ret string
		return ret
	}
	return *o.TransactionReference
}

// GetTransactionReferenceOk returns a tuple with the TransactionReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetTransactionReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionReference) {
		return nil, false
	}
	return o.TransactionReference, true
}

// HasTransactionReference returns a boolean if a field has been set.
func (o *RecurringExpense) HasTransactionReference() bool {
	if o != nil && !IsNil(o.TransactionReference) {
		return true
	}

	return false
}

// SetTransactionReference gets a reference to the given string and assigns it to the TransactionReference field.
func (o *RecurringExpense) SetTransactionReference(v string) {
	o.TransactionReference = &v
}

// GetTranscationId returns the TranscationId field value if set, zero value otherwise.
func (o *RecurringExpense) GetTranscationId() string {
	if o == nil || IsNil(o.TranscationId) {
		var ret string
		return ret
	}
	return *o.TranscationId
}

// GetTranscationIdOk returns a tuple with the TranscationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetTranscationIdOk() (*string, bool) {
	if o == nil || IsNil(o.TranscationId) {
		return nil, false
	}
	return o.TranscationId, true
}

// HasTranscationId returns a boolean if a field has been set.
func (o *RecurringExpense) HasTranscationId() bool {
	if o != nil && !IsNil(o.TranscationId) {
		return true
	}

	return false
}

// SetTranscationId gets a reference to the given string and assigns it to the TranscationId field.
func (o *RecurringExpense) SetTranscationId(v string) {
	o.TranscationId = &v
}

// GetCustomValue1 returns the CustomValue1 field value if set, zero value otherwise.
func (o *RecurringExpense) GetCustomValue1() string {
	if o == nil || IsNil(o.CustomValue1) {
		var ret string
		return ret
	}
	return *o.CustomValue1
}

// GetCustomValue1Ok returns a tuple with the CustomValue1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetCustomValue1Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue1) {
		return nil, false
	}
	return o.CustomValue1, true
}

// HasCustomValue1 returns a boolean if a field has been set.
func (o *RecurringExpense) HasCustomValue1() bool {
	if o != nil && !IsNil(o.CustomValue1) {
		return true
	}

	return false
}

// SetCustomValue1 gets a reference to the given string and assigns it to the CustomValue1 field.
func (o *RecurringExpense) SetCustomValue1(v string) {
	o.CustomValue1 = &v
}

// GetCustomValue2 returns the CustomValue2 field value if set, zero value otherwise.
func (o *RecurringExpense) GetCustomValue2() string {
	if o == nil || IsNil(o.CustomValue2) {
		var ret string
		return ret
	}
	return *o.CustomValue2
}

// GetCustomValue2Ok returns a tuple with the CustomValue2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetCustomValue2Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue2) {
		return nil, false
	}
	return o.CustomValue2, true
}

// HasCustomValue2 returns a boolean if a field has been set.
func (o *RecurringExpense) HasCustomValue2() bool {
	if o != nil && !IsNil(o.CustomValue2) {
		return true
	}

	return false
}

// SetCustomValue2 gets a reference to the given string and assigns it to the CustomValue2 field.
func (o *RecurringExpense) SetCustomValue2(v string) {
	o.CustomValue2 = &v
}

// GetCustomValue3 returns the CustomValue3 field value if set, zero value otherwise.
func (o *RecurringExpense) GetCustomValue3() string {
	if o == nil || IsNil(o.CustomValue3) {
		var ret string
		return ret
	}
	return *o.CustomValue3
}

// GetCustomValue3Ok returns a tuple with the CustomValue3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetCustomValue3Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue3) {
		return nil, false
	}
	return o.CustomValue3, true
}

// HasCustomValue3 returns a boolean if a field has been set.
func (o *RecurringExpense) HasCustomValue3() bool {
	if o != nil && !IsNil(o.CustomValue3) {
		return true
	}

	return false
}

// SetCustomValue3 gets a reference to the given string and assigns it to the CustomValue3 field.
func (o *RecurringExpense) SetCustomValue3(v string) {
	o.CustomValue3 = &v
}

// GetCustomValue4 returns the CustomValue4 field value if set, zero value otherwise.
func (o *RecurringExpense) GetCustomValue4() string {
	if o == nil || IsNil(o.CustomValue4) {
		var ret string
		return ret
	}
	return *o.CustomValue4
}

// GetCustomValue4Ok returns a tuple with the CustomValue4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetCustomValue4Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue4) {
		return nil, false
	}
	return o.CustomValue4, true
}

// HasCustomValue4 returns a boolean if a field has been set.
func (o *RecurringExpense) HasCustomValue4() bool {
	if o != nil && !IsNil(o.CustomValue4) {
		return true
	}

	return false
}

// SetCustomValue4 gets a reference to the given string and assigns it to the CustomValue4 field.
func (o *RecurringExpense) SetCustomValue4(v string) {
	o.CustomValue4 = &v
}

// GetTaxName1 returns the TaxName1 field value if set, zero value otherwise.
func (o *RecurringExpense) GetTaxName1() string {
	if o == nil || IsNil(o.TaxName1) {
		var ret string
		return ret
	}
	return *o.TaxName1
}

// GetTaxName1Ok returns a tuple with the TaxName1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetTaxName1Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName1) {
		return nil, false
	}
	return o.TaxName1, true
}

// HasTaxName1 returns a boolean if a field has been set.
func (o *RecurringExpense) HasTaxName1() bool {
	if o != nil && !IsNil(o.TaxName1) {
		return true
	}

	return false
}

// SetTaxName1 gets a reference to the given string and assigns it to the TaxName1 field.
func (o *RecurringExpense) SetTaxName1(v string) {
	o.TaxName1 = &v
}

// GetTaxName2 returns the TaxName2 field value if set, zero value otherwise.
func (o *RecurringExpense) GetTaxName2() string {
	if o == nil || IsNil(o.TaxName2) {
		var ret string
		return ret
	}
	return *o.TaxName2
}

// GetTaxName2Ok returns a tuple with the TaxName2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetTaxName2Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName2) {
		return nil, false
	}
	return o.TaxName2, true
}

// HasTaxName2 returns a boolean if a field has been set.
func (o *RecurringExpense) HasTaxName2() bool {
	if o != nil && !IsNil(o.TaxName2) {
		return true
	}

	return false
}

// SetTaxName2 gets a reference to the given string and assigns it to the TaxName2 field.
func (o *RecurringExpense) SetTaxName2(v string) {
	o.TaxName2 = &v
}

// GetTaxRate1 returns the TaxRate1 field value if set, zero value otherwise.
func (o *RecurringExpense) GetTaxRate1() float32 {
	if o == nil || IsNil(o.TaxRate1) {
		var ret float32
		return ret
	}
	return *o.TaxRate1
}

// GetTaxRate1Ok returns a tuple with the TaxRate1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetTaxRate1Ok() (*float32, bool) {
	if o == nil || IsNil(o.TaxRate1) {
		return nil, false
	}
	return o.TaxRate1, true
}

// HasTaxRate1 returns a boolean if a field has been set.
func (o *RecurringExpense) HasTaxRate1() bool {
	if o != nil && !IsNil(o.TaxRate1) {
		return true
	}

	return false
}

// SetTaxRate1 gets a reference to the given float32 and assigns it to the TaxRate1 field.
func (o *RecurringExpense) SetTaxRate1(v float32) {
	o.TaxRate1 = &v
}

// GetTaxRate2 returns the TaxRate2 field value if set, zero value otherwise.
func (o *RecurringExpense) GetTaxRate2() float32 {
	if o == nil || IsNil(o.TaxRate2) {
		var ret float32
		return ret
	}
	return *o.TaxRate2
}

// GetTaxRate2Ok returns a tuple with the TaxRate2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetTaxRate2Ok() (*float32, bool) {
	if o == nil || IsNil(o.TaxRate2) {
		return nil, false
	}
	return o.TaxRate2, true
}

// HasTaxRate2 returns a boolean if a field has been set.
func (o *RecurringExpense) HasTaxRate2() bool {
	if o != nil && !IsNil(o.TaxRate2) {
		return true
	}

	return false
}

// SetTaxRate2 gets a reference to the given float32 and assigns it to the TaxRate2 field.
func (o *RecurringExpense) SetTaxRate2(v float32) {
	o.TaxRate2 = &v
}

// GetTaxName3 returns the TaxName3 field value if set, zero value otherwise.
func (o *RecurringExpense) GetTaxName3() string {
	if o == nil || IsNil(o.TaxName3) {
		var ret string
		return ret
	}
	return *o.TaxName3
}

// GetTaxName3Ok returns a tuple with the TaxName3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetTaxName3Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName3) {
		return nil, false
	}
	return o.TaxName3, true
}

// HasTaxName3 returns a boolean if a field has been set.
func (o *RecurringExpense) HasTaxName3() bool {
	if o != nil && !IsNil(o.TaxName3) {
		return true
	}

	return false
}

// SetTaxName3 gets a reference to the given string and assigns it to the TaxName3 field.
func (o *RecurringExpense) SetTaxName3(v string) {
	o.TaxName3 = &v
}

// GetTaxRate3 returns the TaxRate3 field value if set, zero value otherwise.
func (o *RecurringExpense) GetTaxRate3() float32 {
	if o == nil || IsNil(o.TaxRate3) {
		var ret float32
		return ret
	}
	return *o.TaxRate3
}

// GetTaxRate3Ok returns a tuple with the TaxRate3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetTaxRate3Ok() (*float32, bool) {
	if o == nil || IsNil(o.TaxRate3) {
		return nil, false
	}
	return o.TaxRate3, true
}

// HasTaxRate3 returns a boolean if a field has been set.
func (o *RecurringExpense) HasTaxRate3() bool {
	if o != nil && !IsNil(o.TaxRate3) {
		return true
	}

	return false
}

// SetTaxRate3 gets a reference to the given float32 and assigns it to the TaxRate3 field.
func (o *RecurringExpense) SetTaxRate3(v float32) {
	o.TaxRate3 = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *RecurringExpense) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *RecurringExpense) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *RecurringExpense) SetAmount(v float32) {
	o.Amount = &v
}

// GetFrequencyId returns the FrequencyId field value if set, zero value otherwise.
func (o *RecurringExpense) GetFrequencyId() float32 {
	if o == nil || IsNil(o.FrequencyId) {
		var ret float32
		return ret
	}
	return *o.FrequencyId
}

// GetFrequencyIdOk returns a tuple with the FrequencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetFrequencyIdOk() (*float32, bool) {
	if o == nil || IsNil(o.FrequencyId) {
		return nil, false
	}
	return o.FrequencyId, true
}

// HasFrequencyId returns a boolean if a field has been set.
func (o *RecurringExpense) HasFrequencyId() bool {
	if o != nil && !IsNil(o.FrequencyId) {
		return true
	}

	return false
}

// SetFrequencyId gets a reference to the given float32 and assigns it to the FrequencyId field.
func (o *RecurringExpense) SetFrequencyId(v float32) {
	o.FrequencyId = &v
}

// GetRemainingCycles returns the RemainingCycles field value if set, zero value otherwise.
func (o *RecurringExpense) GetRemainingCycles() float32 {
	if o == nil || IsNil(o.RemainingCycles) {
		var ret float32
		return ret
	}
	return *o.RemainingCycles
}

// GetRemainingCyclesOk returns a tuple with the RemainingCycles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetRemainingCyclesOk() (*float32, bool) {
	if o == nil || IsNil(o.RemainingCycles) {
		return nil, false
	}
	return o.RemainingCycles, true
}

// HasRemainingCycles returns a boolean if a field has been set.
func (o *RecurringExpense) HasRemainingCycles() bool {
	if o != nil && !IsNil(o.RemainingCycles) {
		return true
	}

	return false
}

// SetRemainingCycles gets a reference to the given float32 and assigns it to the RemainingCycles field.
func (o *RecurringExpense) SetRemainingCycles(v float32) {
	o.RemainingCycles = &v
}

// GetForeignAmount returns the ForeignAmount field value if set, zero value otherwise.
func (o *RecurringExpense) GetForeignAmount() float32 {
	if o == nil || IsNil(o.ForeignAmount) {
		var ret float32
		return ret
	}
	return *o.ForeignAmount
}

// GetForeignAmountOk returns a tuple with the ForeignAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetForeignAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.ForeignAmount) {
		return nil, false
	}
	return o.ForeignAmount, true
}

// HasForeignAmount returns a boolean if a field has been set.
func (o *RecurringExpense) HasForeignAmount() bool {
	if o != nil && !IsNil(o.ForeignAmount) {
		return true
	}

	return false
}

// SetForeignAmount gets a reference to the given float32 and assigns it to the ForeignAmount field.
func (o *RecurringExpense) SetForeignAmount(v float32) {
	o.ForeignAmount = &v
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise.
func (o *RecurringExpense) GetExchangeRate() float32 {
	if o == nil || IsNil(o.ExchangeRate) {
		var ret float32
		return ret
	}
	return *o.ExchangeRate
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetExchangeRateOk() (*float32, bool) {
	if o == nil || IsNil(o.ExchangeRate) {
		return nil, false
	}
	return o.ExchangeRate, true
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *RecurringExpense) HasExchangeRate() bool {
	if o != nil && !IsNil(o.ExchangeRate) {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given float32 and assigns it to the ExchangeRate field.
func (o *RecurringExpense) SetExchangeRate(v float32) {
	o.ExchangeRate = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *RecurringExpense) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *RecurringExpense) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *RecurringExpense) SetDate(v string) {
	o.Date = &v
}

// GetPaymentDate returns the PaymentDate field value if set, zero value otherwise.
func (o *RecurringExpense) GetPaymentDate() string {
	if o == nil || IsNil(o.PaymentDate) {
		var ret string
		return ret
	}
	return *o.PaymentDate
}

// GetPaymentDateOk returns a tuple with the PaymentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetPaymentDateOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentDate) {
		return nil, false
	}
	return o.PaymentDate, true
}

// HasPaymentDate returns a boolean if a field has been set.
func (o *RecurringExpense) HasPaymentDate() bool {
	if o != nil && !IsNil(o.PaymentDate) {
		return true
	}

	return false
}

// SetPaymentDate gets a reference to the given string and assigns it to the PaymentDate field.
func (o *RecurringExpense) SetPaymentDate(v string) {
	o.PaymentDate = &v
}

// GetShouldBeInvoiced returns the ShouldBeInvoiced field value if set, zero value otherwise.
func (o *RecurringExpense) GetShouldBeInvoiced() bool {
	if o == nil || IsNil(o.ShouldBeInvoiced) {
		var ret bool
		return ret
	}
	return *o.ShouldBeInvoiced
}

// GetShouldBeInvoicedOk returns a tuple with the ShouldBeInvoiced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetShouldBeInvoicedOk() (*bool, bool) {
	if o == nil || IsNil(o.ShouldBeInvoiced) {
		return nil, false
	}
	return o.ShouldBeInvoiced, true
}

// HasShouldBeInvoiced returns a boolean if a field has been set.
func (o *RecurringExpense) HasShouldBeInvoiced() bool {
	if o != nil && !IsNil(o.ShouldBeInvoiced) {
		return true
	}

	return false
}

// SetShouldBeInvoiced gets a reference to the given bool and assigns it to the ShouldBeInvoiced field.
func (o *RecurringExpense) SetShouldBeInvoiced(v bool) {
	o.ShouldBeInvoiced = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *RecurringExpense) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *RecurringExpense) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *RecurringExpense) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetLastSentDate returns the LastSentDate field value if set, zero value otherwise.
func (o *RecurringExpense) GetLastSentDate() string {
	if o == nil || IsNil(o.LastSentDate) {
		var ret string
		return ret
	}
	return *o.LastSentDate
}

// GetLastSentDateOk returns a tuple with the LastSentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetLastSentDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastSentDate) {
		return nil, false
	}
	return o.LastSentDate, true
}

// HasLastSentDate returns a boolean if a field has been set.
func (o *RecurringExpense) HasLastSentDate() bool {
	if o != nil && !IsNil(o.LastSentDate) {
		return true
	}

	return false
}

// SetLastSentDate gets a reference to the given string and assigns it to the LastSentDate field.
func (o *RecurringExpense) SetLastSentDate(v string) {
	o.LastSentDate = &v
}

// GetNextSendDate returns the NextSendDate field value if set, zero value otherwise.
func (o *RecurringExpense) GetNextSendDate() string {
	if o == nil || IsNil(o.NextSendDate) {
		var ret string
		return ret
	}
	return *o.NextSendDate
}

// GetNextSendDateOk returns a tuple with the NextSendDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetNextSendDateOk() (*string, bool) {
	if o == nil || IsNil(o.NextSendDate) {
		return nil, false
	}
	return o.NextSendDate, true
}

// HasNextSendDate returns a boolean if a field has been set.
func (o *RecurringExpense) HasNextSendDate() bool {
	if o != nil && !IsNil(o.NextSendDate) {
		return true
	}

	return false
}

// SetNextSendDate gets a reference to the given string and assigns it to the NextSendDate field.
func (o *RecurringExpense) SetNextSendDate(v string) {
	o.NextSendDate = &v
}

// GetInvoiceDocuments returns the InvoiceDocuments field value if set, zero value otherwise.
func (o *RecurringExpense) GetInvoiceDocuments() bool {
	if o == nil || IsNil(o.InvoiceDocuments) {
		var ret bool
		return ret
	}
	return *o.InvoiceDocuments
}

// GetInvoiceDocumentsOk returns a tuple with the InvoiceDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetInvoiceDocumentsOk() (*bool, bool) {
	if o == nil || IsNil(o.InvoiceDocuments) {
		return nil, false
	}
	return o.InvoiceDocuments, true
}

// HasInvoiceDocuments returns a boolean if a field has been set.
func (o *RecurringExpense) HasInvoiceDocuments() bool {
	if o != nil && !IsNil(o.InvoiceDocuments) {
		return true
	}

	return false
}

// SetInvoiceDocuments gets a reference to the given bool and assigns it to the InvoiceDocuments field.
func (o *RecurringExpense) SetInvoiceDocuments(v bool) {
	o.InvoiceDocuments = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RecurringExpense) GetUpdatedAt() float32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret float32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetUpdatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RecurringExpense) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given float32 and assigns it to the UpdatedAt field.
func (o *RecurringExpense) SetUpdatedAt(v float32) {
	o.UpdatedAt = &v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *RecurringExpense) GetArchivedAt() float32 {
	if o == nil || IsNil(o.ArchivedAt) {
		var ret float32
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecurringExpense) GetArchivedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.ArchivedAt) {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *RecurringExpense) HasArchivedAt() bool {
	if o != nil && !IsNil(o.ArchivedAt) {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given float32 and assigns it to the ArchivedAt field.
func (o *RecurringExpense) SetArchivedAt(v float32) {
	o.ArchivedAt = &v
}

func (o RecurringExpense) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecurringExpense) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.AssignedUserId) {
		toSerialize["assigned_user_id"] = o.AssignedUserId
	}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.InvoiceId) {
		toSerialize["invoice_id"] = o.InvoiceId
	}
	if !IsNil(o.BankId) {
		toSerialize["bank_id"] = o.BankId
	}
	if !IsNil(o.InvoiceCurrencyId) {
		toSerialize["invoice_currency_id"] = o.InvoiceCurrencyId
	}
	if !IsNil(o.ExpenseCurrencyId) {
		toSerialize["expense_currency_id"] = o.ExpenseCurrencyId
	}
	if !IsNil(o.InvoiceCategoryId) {
		toSerialize["invoice_category_id"] = o.InvoiceCategoryId
	}
	if !IsNil(o.PaymentTypeId) {
		toSerialize["payment_type_id"] = o.PaymentTypeId
	}
	if !IsNil(o.PrivateNotes) {
		toSerialize["private_notes"] = o.PrivateNotes
	}
	if !IsNil(o.PublicNotes) {
		toSerialize["public_notes"] = o.PublicNotes
	}
	if !IsNil(o.TransactionReference) {
		toSerialize["transaction_reference"] = o.TransactionReference
	}
	if !IsNil(o.TranscationId) {
		toSerialize["transcation_id"] = o.TranscationId
	}
	if !IsNil(o.CustomValue1) {
		toSerialize["custom_value1"] = o.CustomValue1
	}
	if !IsNil(o.CustomValue2) {
		toSerialize["custom_value2"] = o.CustomValue2
	}
	if !IsNil(o.CustomValue3) {
		toSerialize["custom_value3"] = o.CustomValue3
	}
	if !IsNil(o.CustomValue4) {
		toSerialize["custom_value4"] = o.CustomValue4
	}
	if !IsNil(o.TaxName1) {
		toSerialize["tax_name1"] = o.TaxName1
	}
	if !IsNil(o.TaxName2) {
		toSerialize["tax_name2"] = o.TaxName2
	}
	if !IsNil(o.TaxRate1) {
		toSerialize["tax_rate1"] = o.TaxRate1
	}
	if !IsNil(o.TaxRate2) {
		toSerialize["tax_rate2"] = o.TaxRate2
	}
	if !IsNil(o.TaxName3) {
		toSerialize["tax_name3"] = o.TaxName3
	}
	if !IsNil(o.TaxRate3) {
		toSerialize["tax_rate3"] = o.TaxRate3
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.FrequencyId) {
		toSerialize["frequency_id"] = o.FrequencyId
	}
	if !IsNil(o.RemainingCycles) {
		toSerialize["remaining_cycles"] = o.RemainingCycles
	}
	if !IsNil(o.ForeignAmount) {
		toSerialize["foreign_amount"] = o.ForeignAmount
	}
	if !IsNil(o.ExchangeRate) {
		toSerialize["exchange_rate"] = o.ExchangeRate
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.PaymentDate) {
		toSerialize["payment_date"] = o.PaymentDate
	}
	if !IsNil(o.ShouldBeInvoiced) {
		toSerialize["should_be_invoiced"] = o.ShouldBeInvoiced
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["is_deleted"] = o.IsDeleted
	}
	if !IsNil(o.LastSentDate) {
		toSerialize["last_sent_date"] = o.LastSentDate
	}
	if !IsNil(o.NextSendDate) {
		toSerialize["next_send_date"] = o.NextSendDate
	}
	if !IsNil(o.InvoiceDocuments) {
		toSerialize["invoice_documents"] = o.InvoiceDocuments
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.ArchivedAt) {
		toSerialize["archived_at"] = o.ArchivedAt
	}
	return toSerialize, nil
}

type NullableRecurringExpense struct {
	value *RecurringExpense
	isSet bool
}

func (v NullableRecurringExpense) Get() *RecurringExpense {
	return v.value
}

func (v *NullableRecurringExpense) Set(val *RecurringExpense) {
	v.value = val
	v.isSet = true
}

func (v NullableRecurringExpense) IsSet() bool {
	return v.isSet
}

func (v *NullableRecurringExpense) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecurringExpense(val *RecurringExpense) *NullableRecurringExpense {
	return &NullableRecurringExpense{value: val, isSet: true}
}

func (v NullableRecurringExpense) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecurringExpense) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
