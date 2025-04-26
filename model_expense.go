/*
Invoice Ninja API Reference.

---   ![Invoice Ninja](https://invoicing.co/images/new_logo.png)   ## Introduction   Welcome to the Invoice Ninja API documentation, your comprehensive guide to integrating Invoice Ninja's powerful features into your applications. Whether you're building a custom client, automating workflows, or integrating with other systems, our API provides the tools you need to streamline your invoicing and billing processes.   ### What is Invoice Ninja?   Invoice Ninja is a robust source-available platform designed to simplify invoicing, billing, and payment management for freelancers, small businesses, and enterprises alike. With a user-friendly interface, customizable templates, and a suite of powerful features, Invoice Ninja empowers businesses to create professional invoices, track expenses, manage clients, and get paid faster.   ### Why use the Invoice Ninja API?   The Invoice Ninja API allows developers to extend the functionality of Invoice Ninja by programmatically accessing and manipulating data within their Invoice Ninja accounts. With the API, you can automate repetitive tasks, integrate with third-party services, and build custom solutions tailored to your specific business needs.   ### Getting Started   To get started with the Invoice Ninja API, you'll need an active Invoice Ninja account (or your own self hosted installation) and API credentials. If you haven't already done so, sign up for an account at Invoice Ninja and generate your API keys from the settings section.    Once you have your API credentials, you can start exploring the API endpoints, authentication methods, request and response formats, and more using the documentation provided here.   ### Explore the Documentation     This documentation is organized into sections to help you navigate and understand the various aspects of the Invoice Ninja API:    - Authentication: Learn how to authenticate your requests to the API using API tokens.   - Endpoints: Explore the available API endpoints for managing invoices, clients, payments, expenses, and more.   - Request and Response Formats: Understand the structure of API requests and responses, including parameters, headers, and payloads.   - Error Handling: Learn about error codes, status messages, and best practices for handling errors gracefully.   - Code Examples: Find code examples and tutorials to help you get started with integrating the Invoice Ninja API into your applications.         ### Need Help?         If you have any questions, encounter any issues, or need assistance with using the Invoice Ninja API, don't hesitate to reach out to our support team or join our community forums. We're here to help you succeed with Invoice Ninja and make the most of our API.        Let's start building together!   ### Endpoints      <div style=\"background-color: #2D394E; color: #fff padding: 20px; border-radius: 5px; border: 4px solid #212A3B; box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);\">       <p style=\"padding:10px; color: #DBDBDB;\"\">Production: https://invoicing.co</p>       <p style=\"padding:10px; color: #DBDBDB;\">Demo: https://demo.invoiceninja.com</p>   </div>    ### Client Libraries    PHP SDK can be found [here](https://github.com/invoiceninja/sdk-php)   ### Authentication:    Invoice Ninja uses API tokens to authenticate requests. You can view and manage your API keys in Settings > Account Management > Integrations > API tokens    API requests must be made over HTTPS. Calls made to HTTP will fail.   ### Errors:    Invoice Ninja uses standard HTTP response codes to indicate the success or failure of a request. below is a table of standard status codes and responses    | Status Code | Explanation                                                                 |   |-------------|-----------------------------------------------------------------------------|   | 200         | OK: The request has succeeded. The information returned with the response is dependent on the method used in the request. |   | 301         | Moved Permanently: This and all future requests should be directed to the given URI. |   | 303         | See Other: The response to the request can be found under another URI using the GET method. |   | 400         | Bad Request: The server cannot or will not process the request due to an apparent client error. |   | 401         | Unauthorized: Similar to 403 Forbidden, but specifically for use when authentication is required and has failed or has not yet been provided. |   | 403         | Forbidden: The request was valid, but the server is refusing action. |   | 404         | Not Found: The requested resource could not be found but may be available in the future. |   | 405         | Method Not Allowed: A request method is not supported for the requested resource. |   | 409         | Conflict: Indicates that the request could not be processed because of conflict in the request. |   | 422         | Unprocessable Entity: The request was well-formed but was unable to be followed due to semantic errors. |   | 429         | Too Many Requests: The user has sent too many requests in a given amount of time (\"rate limiting\"). |   | 500         | Internal Server Error: A generic error message, given when an unexpected condition was encountered and no more specific message is suitable. |   ### Pagination    When using index routes to retrieve lists of data, by default we limit the number of records returned to 20. You can using standard pagination to paginate results, ie: ?per_page=50 

API version: 5.11.48
Contact: contact@invoiceninja.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Expense type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Expense{}

// Expense struct for Expense
type Expense struct {
	// The expense hashed id
	Id *string `json:"id,omitempty"`
	// The user hashed id
	UserId *string `json:"user_id,omitempty"`
	// The assigned user hashed id
	AssignedUserId *string `json:"assigned_user_id,omitempty"`
	// The associated project_id
	ProjectId *string `json:"project_id,omitempty"`
	// The client hashed id
	ClientId *string `json:"client_id,omitempty"`
	// The related invoice hashed id
	InvoiceId *string `json:"invoice_id,omitempty"`
	// The bank id related to this expense
	BankId *string `json:"bank_id,omitempty"`
	// The currency id of the related invoice
	InvoiceCurrencyId *string `json:"invoice_currency_id,omitempty"`
	// The currency id of the expense
	CurrencyId *string `json:"currency_id,omitempty"`
	// The invoice category id
	InvoiceCategoryId *string `json:"invoice_category_id,omitempty"`
	// The payment type id
	PaymentTypeId *string `json:"payment_type_id,omitempty"`
	// The related recurring expense this expense was created from
	RecurringExpenseId *string `json:"recurring_expense_id,omitempty"`
	// The private notes of the expense
	PrivateNotes *string `json:"private_notes,omitempty"`
	// The public notes of the expense
	PublicNotes *string `json:"public_notes,omitempty"`
	// The transaction references of the expense
	TransactionReference *string `json:"transaction_reference,omitempty"`
	// The transaction id of the expense
	TranscationId *string `json:"transcation_id,omitempty"`
	// A custom value
	CustomValue1 *string `json:"custom_value1,omitempty"`
	// A custom value
	CustomValue2 *string `json:"custom_value2,omitempty"`
	// A custom value
	CustomValue3 *string `json:"custom_value3,omitempty"`
	// A custom value
	CustomValue4 *string `json:"custom_value4,omitempty"`
	// The tax amount
	TaxAmount *float32 `json:"tax_amount,omitempty"`
	// Tax Name 1
	TaxName1 *string `json:"tax_name1,omitempty"`
	// Tax Name 2
	TaxName2 *string `json:"tax_name2,omitempty"`
	// Tax Name 3
	TaxName3 *string `json:"tax_name3,omitempty"`
	// Tax rate 1
	TaxRate1 *float32 `json:"tax_rate1,omitempty"`
	// Tax rate 2
	TaxRate2 *float32 `json:"tax_rate2,omitempty"`
	// Tax rate 3
	TaxRate3 *float32 `json:"tax_rate3,omitempty"`
	// The total expense amont
	Amount *float32 `json:"amount,omitempty"`
	// The total foreign amount of the expense
	ForeignAmount *float32 `json:"foreign_amount,omitempty"`
	// The exchange rate at the time of the expense
	ExchangeRate *float32 `json:"exchange_rate,omitempty"`
	// The expense date format Y-m-d
	Date *string `json:"date,omitempty"`
	// The date of payment for the expense, format Y-m-d
	PaymentDate *string `json:"payment_date,omitempty"`
	// Flag whether the expense should be invoiced
	ShouldBeInvoiced *bool `json:"should_be_invoiced,omitempty"`
	// Boolean determining whether the expense has been deleted
	IsDeleted *bool `json:"is_deleted,omitempty"`
	// Passing the expense documents over to the invoice
	InvoiceDocuments *bool `json:"invoice_documents,omitempty"`
	// Timestamp
	UpdatedAt *float32 `json:"updated_at,omitempty"`
	// Timestamp
	ArchivedAt *float32 `json:"archived_at,omitempty"`
}

// NewExpense instantiates a new Expense object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpense() *Expense {
	this := Expense{}
	return &this
}

// NewExpenseWithDefaults instantiates a new Expense object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpenseWithDefaults() *Expense {
	this := Expense{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Expense) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Expense) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Expense) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *Expense) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *Expense) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *Expense) SetUserId(v string) {
	o.UserId = &v
}

// GetAssignedUserId returns the AssignedUserId field value if set, zero value otherwise.
func (o *Expense) GetAssignedUserId() string {
	if o == nil || IsNil(o.AssignedUserId) {
		var ret string
		return ret
	}
	return *o.AssignedUserId
}

// GetAssignedUserIdOk returns a tuple with the AssignedUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetAssignedUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssignedUserId) {
		return nil, false
	}
	return o.AssignedUserId, true
}

// HasAssignedUserId returns a boolean if a field has been set.
func (o *Expense) HasAssignedUserId() bool {
	if o != nil && !IsNil(o.AssignedUserId) {
		return true
	}

	return false
}

// SetAssignedUserId gets a reference to the given string and assigns it to the AssignedUserId field.
func (o *Expense) SetAssignedUserId(v string) {
	o.AssignedUserId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *Expense) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *Expense) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *Expense) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *Expense) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *Expense) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *Expense) SetClientId(v string) {
	o.ClientId = &v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise.
func (o *Expense) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceId) {
		return nil, false
	}
	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *Expense) HasInvoiceId() bool {
	if o != nil && !IsNil(o.InvoiceId) {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *Expense) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

// GetBankId returns the BankId field value if set, zero value otherwise.
func (o *Expense) GetBankId() string {
	if o == nil || IsNil(o.BankId) {
		var ret string
		return ret
	}
	return *o.BankId
}

// GetBankIdOk returns a tuple with the BankId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetBankIdOk() (*string, bool) {
	if o == nil || IsNil(o.BankId) {
		return nil, false
	}
	return o.BankId, true
}

// HasBankId returns a boolean if a field has been set.
func (o *Expense) HasBankId() bool {
	if o != nil && !IsNil(o.BankId) {
		return true
	}

	return false
}

// SetBankId gets a reference to the given string and assigns it to the BankId field.
func (o *Expense) SetBankId(v string) {
	o.BankId = &v
}

// GetInvoiceCurrencyId returns the InvoiceCurrencyId field value if set, zero value otherwise.
func (o *Expense) GetInvoiceCurrencyId() string {
	if o == nil || IsNil(o.InvoiceCurrencyId) {
		var ret string
		return ret
	}
	return *o.InvoiceCurrencyId
}

// GetInvoiceCurrencyIdOk returns a tuple with the InvoiceCurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetInvoiceCurrencyIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceCurrencyId) {
		return nil, false
	}
	return o.InvoiceCurrencyId, true
}

// HasInvoiceCurrencyId returns a boolean if a field has been set.
func (o *Expense) HasInvoiceCurrencyId() bool {
	if o != nil && !IsNil(o.InvoiceCurrencyId) {
		return true
	}

	return false
}

// SetInvoiceCurrencyId gets a reference to the given string and assigns it to the InvoiceCurrencyId field.
func (o *Expense) SetInvoiceCurrencyId(v string) {
	o.InvoiceCurrencyId = &v
}

// GetCurrencyId returns the CurrencyId field value if set, zero value otherwise.
func (o *Expense) GetCurrencyId() string {
	if o == nil || IsNil(o.CurrencyId) {
		var ret string
		return ret
	}
	return *o.CurrencyId
}

// GetCurrencyIdOk returns a tuple with the CurrencyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetCurrencyIdOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyId) {
		return nil, false
	}
	return o.CurrencyId, true
}

// HasCurrencyId returns a boolean if a field has been set.
func (o *Expense) HasCurrencyId() bool {
	if o != nil && !IsNil(o.CurrencyId) {
		return true
	}

	return false
}

// SetCurrencyId gets a reference to the given string and assigns it to the CurrencyId field.
func (o *Expense) SetCurrencyId(v string) {
	o.CurrencyId = &v
}

// GetInvoiceCategoryId returns the InvoiceCategoryId field value if set, zero value otherwise.
func (o *Expense) GetInvoiceCategoryId() string {
	if o == nil || IsNil(o.InvoiceCategoryId) {
		var ret string
		return ret
	}
	return *o.InvoiceCategoryId
}

// GetInvoiceCategoryIdOk returns a tuple with the InvoiceCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetInvoiceCategoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceCategoryId) {
		return nil, false
	}
	return o.InvoiceCategoryId, true
}

// HasInvoiceCategoryId returns a boolean if a field has been set.
func (o *Expense) HasInvoiceCategoryId() bool {
	if o != nil && !IsNil(o.InvoiceCategoryId) {
		return true
	}

	return false
}

// SetInvoiceCategoryId gets a reference to the given string and assigns it to the InvoiceCategoryId field.
func (o *Expense) SetInvoiceCategoryId(v string) {
	o.InvoiceCategoryId = &v
}

// GetPaymentTypeId returns the PaymentTypeId field value if set, zero value otherwise.
func (o *Expense) GetPaymentTypeId() string {
	if o == nil || IsNil(o.PaymentTypeId) {
		var ret string
		return ret
	}
	return *o.PaymentTypeId
}

// GetPaymentTypeIdOk returns a tuple with the PaymentTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetPaymentTypeIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentTypeId) {
		return nil, false
	}
	return o.PaymentTypeId, true
}

// HasPaymentTypeId returns a boolean if a field has been set.
func (o *Expense) HasPaymentTypeId() bool {
	if o != nil && !IsNil(o.PaymentTypeId) {
		return true
	}

	return false
}

// SetPaymentTypeId gets a reference to the given string and assigns it to the PaymentTypeId field.
func (o *Expense) SetPaymentTypeId(v string) {
	o.PaymentTypeId = &v
}

// GetRecurringExpenseId returns the RecurringExpenseId field value if set, zero value otherwise.
func (o *Expense) GetRecurringExpenseId() string {
	if o == nil || IsNil(o.RecurringExpenseId) {
		var ret string
		return ret
	}
	return *o.RecurringExpenseId
}

// GetRecurringExpenseIdOk returns a tuple with the RecurringExpenseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetRecurringExpenseIdOk() (*string, bool) {
	if o == nil || IsNil(o.RecurringExpenseId) {
		return nil, false
	}
	return o.RecurringExpenseId, true
}

// HasRecurringExpenseId returns a boolean if a field has been set.
func (o *Expense) HasRecurringExpenseId() bool {
	if o != nil && !IsNil(o.RecurringExpenseId) {
		return true
	}

	return false
}

// SetRecurringExpenseId gets a reference to the given string and assigns it to the RecurringExpenseId field.
func (o *Expense) SetRecurringExpenseId(v string) {
	o.RecurringExpenseId = &v
}

// GetPrivateNotes returns the PrivateNotes field value if set, zero value otherwise.
func (o *Expense) GetPrivateNotes() string {
	if o == nil || IsNil(o.PrivateNotes) {
		var ret string
		return ret
	}
	return *o.PrivateNotes
}

// GetPrivateNotesOk returns a tuple with the PrivateNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetPrivateNotesOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateNotes) {
		return nil, false
	}
	return o.PrivateNotes, true
}

// HasPrivateNotes returns a boolean if a field has been set.
func (o *Expense) HasPrivateNotes() bool {
	if o != nil && !IsNil(o.PrivateNotes) {
		return true
	}

	return false
}

// SetPrivateNotes gets a reference to the given string and assigns it to the PrivateNotes field.
func (o *Expense) SetPrivateNotes(v string) {
	o.PrivateNotes = &v
}

// GetPublicNotes returns the PublicNotes field value if set, zero value otherwise.
func (o *Expense) GetPublicNotes() string {
	if o == nil || IsNil(o.PublicNotes) {
		var ret string
		return ret
	}
	return *o.PublicNotes
}

// GetPublicNotesOk returns a tuple with the PublicNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetPublicNotesOk() (*string, bool) {
	if o == nil || IsNil(o.PublicNotes) {
		return nil, false
	}
	return o.PublicNotes, true
}

// HasPublicNotes returns a boolean if a field has been set.
func (o *Expense) HasPublicNotes() bool {
	if o != nil && !IsNil(o.PublicNotes) {
		return true
	}

	return false
}

// SetPublicNotes gets a reference to the given string and assigns it to the PublicNotes field.
func (o *Expense) SetPublicNotes(v string) {
	o.PublicNotes = &v
}

// GetTransactionReference returns the TransactionReference field value if set, zero value otherwise.
func (o *Expense) GetTransactionReference() string {
	if o == nil || IsNil(o.TransactionReference) {
		var ret string
		return ret
	}
	return *o.TransactionReference
}

// GetTransactionReferenceOk returns a tuple with the TransactionReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetTransactionReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionReference) {
		return nil, false
	}
	return o.TransactionReference, true
}

// HasTransactionReference returns a boolean if a field has been set.
func (o *Expense) HasTransactionReference() bool {
	if o != nil && !IsNil(o.TransactionReference) {
		return true
	}

	return false
}

// SetTransactionReference gets a reference to the given string and assigns it to the TransactionReference field.
func (o *Expense) SetTransactionReference(v string) {
	o.TransactionReference = &v
}

// GetTranscationId returns the TranscationId field value if set, zero value otherwise.
func (o *Expense) GetTranscationId() string {
	if o == nil || IsNil(o.TranscationId) {
		var ret string
		return ret
	}
	return *o.TranscationId
}

// GetTranscationIdOk returns a tuple with the TranscationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetTranscationIdOk() (*string, bool) {
	if o == nil || IsNil(o.TranscationId) {
		return nil, false
	}
	return o.TranscationId, true
}

// HasTranscationId returns a boolean if a field has been set.
func (o *Expense) HasTranscationId() bool {
	if o != nil && !IsNil(o.TranscationId) {
		return true
	}

	return false
}

// SetTranscationId gets a reference to the given string and assigns it to the TranscationId field.
func (o *Expense) SetTranscationId(v string) {
	o.TranscationId = &v
}

// GetCustomValue1 returns the CustomValue1 field value if set, zero value otherwise.
func (o *Expense) GetCustomValue1() string {
	if o == nil || IsNil(o.CustomValue1) {
		var ret string
		return ret
	}
	return *o.CustomValue1
}

// GetCustomValue1Ok returns a tuple with the CustomValue1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetCustomValue1Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue1) {
		return nil, false
	}
	return o.CustomValue1, true
}

// HasCustomValue1 returns a boolean if a field has been set.
func (o *Expense) HasCustomValue1() bool {
	if o != nil && !IsNil(o.CustomValue1) {
		return true
	}

	return false
}

// SetCustomValue1 gets a reference to the given string and assigns it to the CustomValue1 field.
func (o *Expense) SetCustomValue1(v string) {
	o.CustomValue1 = &v
}

// GetCustomValue2 returns the CustomValue2 field value if set, zero value otherwise.
func (o *Expense) GetCustomValue2() string {
	if o == nil || IsNil(o.CustomValue2) {
		var ret string
		return ret
	}
	return *o.CustomValue2
}

// GetCustomValue2Ok returns a tuple with the CustomValue2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetCustomValue2Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue2) {
		return nil, false
	}
	return o.CustomValue2, true
}

// HasCustomValue2 returns a boolean if a field has been set.
func (o *Expense) HasCustomValue2() bool {
	if o != nil && !IsNil(o.CustomValue2) {
		return true
	}

	return false
}

// SetCustomValue2 gets a reference to the given string and assigns it to the CustomValue2 field.
func (o *Expense) SetCustomValue2(v string) {
	o.CustomValue2 = &v
}

// GetCustomValue3 returns the CustomValue3 field value if set, zero value otherwise.
func (o *Expense) GetCustomValue3() string {
	if o == nil || IsNil(o.CustomValue3) {
		var ret string
		return ret
	}
	return *o.CustomValue3
}

// GetCustomValue3Ok returns a tuple with the CustomValue3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetCustomValue3Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue3) {
		return nil, false
	}
	return o.CustomValue3, true
}

// HasCustomValue3 returns a boolean if a field has been set.
func (o *Expense) HasCustomValue3() bool {
	if o != nil && !IsNil(o.CustomValue3) {
		return true
	}

	return false
}

// SetCustomValue3 gets a reference to the given string and assigns it to the CustomValue3 field.
func (o *Expense) SetCustomValue3(v string) {
	o.CustomValue3 = &v
}

// GetCustomValue4 returns the CustomValue4 field value if set, zero value otherwise.
func (o *Expense) GetCustomValue4() string {
	if o == nil || IsNil(o.CustomValue4) {
		var ret string
		return ret
	}
	return *o.CustomValue4
}

// GetCustomValue4Ok returns a tuple with the CustomValue4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetCustomValue4Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue4) {
		return nil, false
	}
	return o.CustomValue4, true
}

// HasCustomValue4 returns a boolean if a field has been set.
func (o *Expense) HasCustomValue4() bool {
	if o != nil && !IsNil(o.CustomValue4) {
		return true
	}

	return false
}

// SetCustomValue4 gets a reference to the given string and assigns it to the CustomValue4 field.
func (o *Expense) SetCustomValue4(v string) {
	o.CustomValue4 = &v
}

// GetTaxAmount returns the TaxAmount field value if set, zero value otherwise.
func (o *Expense) GetTaxAmount() float32 {
	if o == nil || IsNil(o.TaxAmount) {
		var ret float32
		return ret
	}
	return *o.TaxAmount
}

// GetTaxAmountOk returns a tuple with the TaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetTaxAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.TaxAmount) {
		return nil, false
	}
	return o.TaxAmount, true
}

// HasTaxAmount returns a boolean if a field has been set.
func (o *Expense) HasTaxAmount() bool {
	if o != nil && !IsNil(o.TaxAmount) {
		return true
	}

	return false
}

// SetTaxAmount gets a reference to the given float32 and assigns it to the TaxAmount field.
func (o *Expense) SetTaxAmount(v float32) {
	o.TaxAmount = &v
}

// GetTaxName1 returns the TaxName1 field value if set, zero value otherwise.
func (o *Expense) GetTaxName1() string {
	if o == nil || IsNil(o.TaxName1) {
		var ret string
		return ret
	}
	return *o.TaxName1
}

// GetTaxName1Ok returns a tuple with the TaxName1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetTaxName1Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName1) {
		return nil, false
	}
	return o.TaxName1, true
}

// HasTaxName1 returns a boolean if a field has been set.
func (o *Expense) HasTaxName1() bool {
	if o != nil && !IsNil(o.TaxName1) {
		return true
	}

	return false
}

// SetTaxName1 gets a reference to the given string and assigns it to the TaxName1 field.
func (o *Expense) SetTaxName1(v string) {
	o.TaxName1 = &v
}

// GetTaxName2 returns the TaxName2 field value if set, zero value otherwise.
func (o *Expense) GetTaxName2() string {
	if o == nil || IsNil(o.TaxName2) {
		var ret string
		return ret
	}
	return *o.TaxName2
}

// GetTaxName2Ok returns a tuple with the TaxName2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetTaxName2Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName2) {
		return nil, false
	}
	return o.TaxName2, true
}

// HasTaxName2 returns a boolean if a field has been set.
func (o *Expense) HasTaxName2() bool {
	if o != nil && !IsNil(o.TaxName2) {
		return true
	}

	return false
}

// SetTaxName2 gets a reference to the given string and assigns it to the TaxName2 field.
func (o *Expense) SetTaxName2(v string) {
	o.TaxName2 = &v
}

// GetTaxName3 returns the TaxName3 field value if set, zero value otherwise.
func (o *Expense) GetTaxName3() string {
	if o == nil || IsNil(o.TaxName3) {
		var ret string
		return ret
	}
	return *o.TaxName3
}

// GetTaxName3Ok returns a tuple with the TaxName3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetTaxName3Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName3) {
		return nil, false
	}
	return o.TaxName3, true
}

// HasTaxName3 returns a boolean if a field has been set.
func (o *Expense) HasTaxName3() bool {
	if o != nil && !IsNil(o.TaxName3) {
		return true
	}

	return false
}

// SetTaxName3 gets a reference to the given string and assigns it to the TaxName3 field.
func (o *Expense) SetTaxName3(v string) {
	o.TaxName3 = &v
}

// GetTaxRate1 returns the TaxRate1 field value if set, zero value otherwise.
func (o *Expense) GetTaxRate1() float32 {
	if o == nil || IsNil(o.TaxRate1) {
		var ret float32
		return ret
	}
	return *o.TaxRate1
}

// GetTaxRate1Ok returns a tuple with the TaxRate1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetTaxRate1Ok() (*float32, bool) {
	if o == nil || IsNil(o.TaxRate1) {
		return nil, false
	}
	return o.TaxRate1, true
}

// HasTaxRate1 returns a boolean if a field has been set.
func (o *Expense) HasTaxRate1() bool {
	if o != nil && !IsNil(o.TaxRate1) {
		return true
	}

	return false
}

// SetTaxRate1 gets a reference to the given float32 and assigns it to the TaxRate1 field.
func (o *Expense) SetTaxRate1(v float32) {
	o.TaxRate1 = &v
}

// GetTaxRate2 returns the TaxRate2 field value if set, zero value otherwise.
func (o *Expense) GetTaxRate2() float32 {
	if o == nil || IsNil(o.TaxRate2) {
		var ret float32
		return ret
	}
	return *o.TaxRate2
}

// GetTaxRate2Ok returns a tuple with the TaxRate2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetTaxRate2Ok() (*float32, bool) {
	if o == nil || IsNil(o.TaxRate2) {
		return nil, false
	}
	return o.TaxRate2, true
}

// HasTaxRate2 returns a boolean if a field has been set.
func (o *Expense) HasTaxRate2() bool {
	if o != nil && !IsNil(o.TaxRate2) {
		return true
	}

	return false
}

// SetTaxRate2 gets a reference to the given float32 and assigns it to the TaxRate2 field.
func (o *Expense) SetTaxRate2(v float32) {
	o.TaxRate2 = &v
}

// GetTaxRate3 returns the TaxRate3 field value if set, zero value otherwise.
func (o *Expense) GetTaxRate3() float32 {
	if o == nil || IsNil(o.TaxRate3) {
		var ret float32
		return ret
	}
	return *o.TaxRate3
}

// GetTaxRate3Ok returns a tuple with the TaxRate3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetTaxRate3Ok() (*float32, bool) {
	if o == nil || IsNil(o.TaxRate3) {
		return nil, false
	}
	return o.TaxRate3, true
}

// HasTaxRate3 returns a boolean if a field has been set.
func (o *Expense) HasTaxRate3() bool {
	if o != nil && !IsNil(o.TaxRate3) {
		return true
	}

	return false
}

// SetTaxRate3 gets a reference to the given float32 and assigns it to the TaxRate3 field.
func (o *Expense) SetTaxRate3(v float32) {
	o.TaxRate3 = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Expense) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Expense) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *Expense) SetAmount(v float32) {
	o.Amount = &v
}

// GetForeignAmount returns the ForeignAmount field value if set, zero value otherwise.
func (o *Expense) GetForeignAmount() float32 {
	if o == nil || IsNil(o.ForeignAmount) {
		var ret float32
		return ret
	}
	return *o.ForeignAmount
}

// GetForeignAmountOk returns a tuple with the ForeignAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetForeignAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.ForeignAmount) {
		return nil, false
	}
	return o.ForeignAmount, true
}

// HasForeignAmount returns a boolean if a field has been set.
func (o *Expense) HasForeignAmount() bool {
	if o != nil && !IsNil(o.ForeignAmount) {
		return true
	}

	return false
}

// SetForeignAmount gets a reference to the given float32 and assigns it to the ForeignAmount field.
func (o *Expense) SetForeignAmount(v float32) {
	o.ForeignAmount = &v
}

// GetExchangeRate returns the ExchangeRate field value if set, zero value otherwise.
func (o *Expense) GetExchangeRate() float32 {
	if o == nil || IsNil(o.ExchangeRate) {
		var ret float32
		return ret
	}
	return *o.ExchangeRate
}

// GetExchangeRateOk returns a tuple with the ExchangeRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetExchangeRateOk() (*float32, bool) {
	if o == nil || IsNil(o.ExchangeRate) {
		return nil, false
	}
	return o.ExchangeRate, true
}

// HasExchangeRate returns a boolean if a field has been set.
func (o *Expense) HasExchangeRate() bool {
	if o != nil && !IsNil(o.ExchangeRate) {
		return true
	}

	return false
}

// SetExchangeRate gets a reference to the given float32 and assigns it to the ExchangeRate field.
func (o *Expense) SetExchangeRate(v float32) {
	o.ExchangeRate = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *Expense) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *Expense) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *Expense) SetDate(v string) {
	o.Date = &v
}

// GetPaymentDate returns the PaymentDate field value if set, zero value otherwise.
func (o *Expense) GetPaymentDate() string {
	if o == nil || IsNil(o.PaymentDate) {
		var ret string
		return ret
	}
	return *o.PaymentDate
}

// GetPaymentDateOk returns a tuple with the PaymentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetPaymentDateOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentDate) {
		return nil, false
	}
	return o.PaymentDate, true
}

// HasPaymentDate returns a boolean if a field has been set.
func (o *Expense) HasPaymentDate() bool {
	if o != nil && !IsNil(o.PaymentDate) {
		return true
	}

	return false
}

// SetPaymentDate gets a reference to the given string and assigns it to the PaymentDate field.
func (o *Expense) SetPaymentDate(v string) {
	o.PaymentDate = &v
}

// GetShouldBeInvoiced returns the ShouldBeInvoiced field value if set, zero value otherwise.
func (o *Expense) GetShouldBeInvoiced() bool {
	if o == nil || IsNil(o.ShouldBeInvoiced) {
		var ret bool
		return ret
	}
	return *o.ShouldBeInvoiced
}

// GetShouldBeInvoicedOk returns a tuple with the ShouldBeInvoiced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetShouldBeInvoicedOk() (*bool, bool) {
	if o == nil || IsNil(o.ShouldBeInvoiced) {
		return nil, false
	}
	return o.ShouldBeInvoiced, true
}

// HasShouldBeInvoiced returns a boolean if a field has been set.
func (o *Expense) HasShouldBeInvoiced() bool {
	if o != nil && !IsNil(o.ShouldBeInvoiced) {
		return true
	}

	return false
}

// SetShouldBeInvoiced gets a reference to the given bool and assigns it to the ShouldBeInvoiced field.
func (o *Expense) SetShouldBeInvoiced(v bool) {
	o.ShouldBeInvoiced = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *Expense) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *Expense) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *Expense) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetInvoiceDocuments returns the InvoiceDocuments field value if set, zero value otherwise.
func (o *Expense) GetInvoiceDocuments() bool {
	if o == nil || IsNil(o.InvoiceDocuments) {
		var ret bool
		return ret
	}
	return *o.InvoiceDocuments
}

// GetInvoiceDocumentsOk returns a tuple with the InvoiceDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetInvoiceDocumentsOk() (*bool, bool) {
	if o == nil || IsNil(o.InvoiceDocuments) {
		return nil, false
	}
	return o.InvoiceDocuments, true
}

// HasInvoiceDocuments returns a boolean if a field has been set.
func (o *Expense) HasInvoiceDocuments() bool {
	if o != nil && !IsNil(o.InvoiceDocuments) {
		return true
	}

	return false
}

// SetInvoiceDocuments gets a reference to the given bool and assigns it to the InvoiceDocuments field.
func (o *Expense) SetInvoiceDocuments(v bool) {
	o.InvoiceDocuments = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Expense) GetUpdatedAt() float32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret float32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetUpdatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Expense) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given float32 and assigns it to the UpdatedAt field.
func (o *Expense) SetUpdatedAt(v float32) {
	o.UpdatedAt = &v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *Expense) GetArchivedAt() float32 {
	if o == nil || IsNil(o.ArchivedAt) {
		var ret float32
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Expense) GetArchivedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.ArchivedAt) {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *Expense) HasArchivedAt() bool {
	if o != nil && !IsNil(o.ArchivedAt) {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given float32 and assigns it to the ArchivedAt field.
func (o *Expense) SetArchivedAt(v float32) {
	o.ArchivedAt = &v
}

func (o Expense) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Expense) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.AssignedUserId) {
		toSerialize["assigned_user_id"] = o.AssignedUserId
	}
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
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
	if !IsNil(o.CurrencyId) {
		toSerialize["currency_id"] = o.CurrencyId
	}
	if !IsNil(o.InvoiceCategoryId) {
		toSerialize["invoice_category_id"] = o.InvoiceCategoryId
	}
	if !IsNil(o.PaymentTypeId) {
		toSerialize["payment_type_id"] = o.PaymentTypeId
	}
	if !IsNil(o.RecurringExpenseId) {
		toSerialize["recurring_expense_id"] = o.RecurringExpenseId
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
	if !IsNil(o.TaxAmount) {
		toSerialize["tax_amount"] = o.TaxAmount
	}
	if !IsNil(o.TaxName1) {
		toSerialize["tax_name1"] = o.TaxName1
	}
	if !IsNil(o.TaxName2) {
		toSerialize["tax_name2"] = o.TaxName2
	}
	if !IsNil(o.TaxName3) {
		toSerialize["tax_name3"] = o.TaxName3
	}
	if !IsNil(o.TaxRate1) {
		toSerialize["tax_rate1"] = o.TaxRate1
	}
	if !IsNil(o.TaxRate2) {
		toSerialize["tax_rate2"] = o.TaxRate2
	}
	if !IsNil(o.TaxRate3) {
		toSerialize["tax_rate3"] = o.TaxRate3
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
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

type NullableExpense struct {
	value *Expense
	isSet bool
}

func (v NullableExpense) Get() *Expense {
	return v.value
}

func (v *NullableExpense) Set(val *Expense) {
	v.value = val
	v.isSet = true
}

func (v NullableExpense) IsSet() bool {
	return v.isSet
}

func (v *NullableExpense) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpense(val *Expense) *NullableExpense {
	return &NullableExpense{value: val, isSet: true}
}

func (v NullableExpense) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpense) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


