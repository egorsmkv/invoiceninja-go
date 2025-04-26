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

// checks if the Activity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Activity{}

// Activity struct for Activity
type Activity struct {
	// The id field of the activity
	Id *string `json:"id,omitempty"`
	// The activity type id
	ActivityTypeId *string `json:"activity_type_id,omitempty"`
	// The client hashed id
	ClientId *string `json:"client_id,omitempty"`
	// The company hashed id
	CompanyId *string `json:"company_id,omitempty"`
	// The user hashed id
	UserId *string `json:"user_id,omitempty"`
	// The invoice hashed id
	InvoiceId *string `json:"invoice_id,omitempty"`
	// The payment hashed id
	PaymentId *string `json:"payment_id,omitempty"`
	// The credit hashed id
	CreditId *string `json:"credit_id,omitempty"`
	// Unixtimestamp the last time the record was updated
	UpdatedAt *int32 `json:"updated_at,omitempty"`
	// The expense hashed id
	ExpenseId *string `json:"expense_id,omitempty"`
	// Defines is the activity was performed by the system
	IsSystem *bool `json:"is_system,omitempty"`
	// The contact hashed id
	ContactId *string `json:"contact_id,omitempty"`
	// The task hashed id
	TaskId *string `json:"task_id,omitempty"`
	// Activity Notes
	Notes *string `json:"notes,omitempty"`
	// The hashed ID of the token who performed the action
	TokenId *string `json:"token_id,omitempty"`
	// The IP Address of the user who performed the action
	Ip *string `json:"ip,omitempty"`
	User *User `json:"user,omitempty"`
	Client *Client `json:"client,omitempty"`
	Contact *ClientContact `json:"contact,omitempty"`
	RecurringInvoice *RecurringInvoice `json:"recurring_invoice,omitempty"`
	Invoice *Invoice `json:"invoice,omitempty"`
	Credit *Credit `json:"credit,omitempty"`
	Quote *Quote `json:"quote,omitempty"`
	Payment *Payment `json:"payment,omitempty"`
	Expense *Expense `json:"expense,omitempty"`
	Task *Task `json:"task,omitempty"`
	PurchaseOrder *PurchaseOrder `json:"purchase_order,omitempty"`
	Vendor *Vendor `json:"vendor,omitempty"`
	VendorContact *VendorContact `json:"vendor_contact,omitempty"`
}

// NewActivity instantiates a new Activity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivity() *Activity {
	this := Activity{}
	return &this
}

// NewActivityWithDefaults instantiates a new Activity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityWithDefaults() *Activity {
	this := Activity{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Activity) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Activity) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Activity) SetId(v string) {
	o.Id = &v
}

// GetActivityTypeId returns the ActivityTypeId field value if set, zero value otherwise.
func (o *Activity) GetActivityTypeId() string {
	if o == nil || IsNil(o.ActivityTypeId) {
		var ret string
		return ret
	}
	return *o.ActivityTypeId
}

// GetActivityTypeIdOk returns a tuple with the ActivityTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetActivityTypeIdOk() (*string, bool) {
	if o == nil || IsNil(o.ActivityTypeId) {
		return nil, false
	}
	return o.ActivityTypeId, true
}

// HasActivityTypeId returns a boolean if a field has been set.
func (o *Activity) HasActivityTypeId() bool {
	if o != nil && !IsNil(o.ActivityTypeId) {
		return true
	}

	return false
}

// SetActivityTypeId gets a reference to the given string and assigns it to the ActivityTypeId field.
func (o *Activity) SetActivityTypeId(v string) {
	o.ActivityTypeId = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *Activity) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *Activity) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *Activity) SetClientId(v string) {
	o.ClientId = &v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *Activity) GetCompanyId() string {
	if o == nil || IsNil(o.CompanyId) {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetCompanyIdOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyId) {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *Activity) HasCompanyId() bool {
	if o != nil && !IsNil(o.CompanyId) {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *Activity) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *Activity) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *Activity) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *Activity) SetUserId(v string) {
	o.UserId = &v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise.
func (o *Activity) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceId) {
		return nil, false
	}
	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *Activity) HasInvoiceId() bool {
	if o != nil && !IsNil(o.InvoiceId) {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *Activity) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *Activity) GetPaymentId() string {
	if o == nil || IsNil(o.PaymentId) {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentId) {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *Activity) HasPaymentId() bool {
	if o != nil && !IsNil(o.PaymentId) {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *Activity) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetCreditId returns the CreditId field value if set, zero value otherwise.
func (o *Activity) GetCreditId() string {
	if o == nil || IsNil(o.CreditId) {
		var ret string
		return ret
	}
	return *o.CreditId
}

// GetCreditIdOk returns a tuple with the CreditId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetCreditIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreditId) {
		return nil, false
	}
	return o.CreditId, true
}

// HasCreditId returns a boolean if a field has been set.
func (o *Activity) HasCreditId() bool {
	if o != nil && !IsNil(o.CreditId) {
		return true
	}

	return false
}

// SetCreditId gets a reference to the given string and assigns it to the CreditId field.
func (o *Activity) SetCreditId(v string) {
	o.CreditId = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Activity) GetUpdatedAt() int32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Activity) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *Activity) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

// GetExpenseId returns the ExpenseId field value if set, zero value otherwise.
func (o *Activity) GetExpenseId() string {
	if o == nil || IsNil(o.ExpenseId) {
		var ret string
		return ret
	}
	return *o.ExpenseId
}

// GetExpenseIdOk returns a tuple with the ExpenseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetExpenseIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExpenseId) {
		return nil, false
	}
	return o.ExpenseId, true
}

// HasExpenseId returns a boolean if a field has been set.
func (o *Activity) HasExpenseId() bool {
	if o != nil && !IsNil(o.ExpenseId) {
		return true
	}

	return false
}

// SetExpenseId gets a reference to the given string and assigns it to the ExpenseId field.
func (o *Activity) SetExpenseId(v string) {
	o.ExpenseId = &v
}

// GetIsSystem returns the IsSystem field value if set, zero value otherwise.
func (o *Activity) GetIsSystem() bool {
	if o == nil || IsNil(o.IsSystem) {
		var ret bool
		return ret
	}
	return *o.IsSystem
}

// GetIsSystemOk returns a tuple with the IsSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetIsSystemOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSystem) {
		return nil, false
	}
	return o.IsSystem, true
}

// HasIsSystem returns a boolean if a field has been set.
func (o *Activity) HasIsSystem() bool {
	if o != nil && !IsNil(o.IsSystem) {
		return true
	}

	return false
}

// SetIsSystem gets a reference to the given bool and assigns it to the IsSystem field.
func (o *Activity) SetIsSystem(v bool) {
	o.IsSystem = &v
}

// GetContactId returns the ContactId field value if set, zero value otherwise.
func (o *Activity) GetContactId() string {
	if o == nil || IsNil(o.ContactId) {
		var ret string
		return ret
	}
	return *o.ContactId
}

// GetContactIdOk returns a tuple with the ContactId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetContactIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContactId) {
		return nil, false
	}
	return o.ContactId, true
}

// HasContactId returns a boolean if a field has been set.
func (o *Activity) HasContactId() bool {
	if o != nil && !IsNil(o.ContactId) {
		return true
	}

	return false
}

// SetContactId gets a reference to the given string and assigns it to the ContactId field.
func (o *Activity) SetContactId(v string) {
	o.ContactId = &v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *Activity) GetTaskId() string {
	if o == nil || IsNil(o.TaskId) {
		var ret string
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetTaskIdOk() (*string, bool) {
	if o == nil || IsNil(o.TaskId) {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *Activity) HasTaskId() bool {
	if o != nil && !IsNil(o.TaskId) {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given string and assigns it to the TaskId field.
func (o *Activity) SetTaskId(v string) {
	o.TaskId = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *Activity) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *Activity) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *Activity) SetNotes(v string) {
	o.Notes = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *Activity) GetTokenId() string {
	if o == nil || IsNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetTokenIdOk() (*string, bool) {
	if o == nil || IsNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *Activity) HasTokenId() bool {
	if o != nil && !IsNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *Activity) SetTokenId(v string) {
	o.TokenId = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *Activity) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *Activity) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *Activity) SetIp(v string) {
	o.Ip = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Activity) GetUser() User {
	if o == nil || IsNil(o.User) {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetUserOk() (*User, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Activity) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *Activity) SetUser(v User) {
	o.User = &v
}

// GetClient returns the Client field value if set, zero value otherwise.
func (o *Activity) GetClient() Client {
	if o == nil || IsNil(o.Client) {
		var ret Client
		return ret
	}
	return *o.Client
}

// GetClientOk returns a tuple with the Client field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetClientOk() (*Client, bool) {
	if o == nil || IsNil(o.Client) {
		return nil, false
	}
	return o.Client, true
}

// HasClient returns a boolean if a field has been set.
func (o *Activity) HasClient() bool {
	if o != nil && !IsNil(o.Client) {
		return true
	}

	return false
}

// SetClient gets a reference to the given Client and assigns it to the Client field.
func (o *Activity) SetClient(v Client) {
	o.Client = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *Activity) GetContact() ClientContact {
	if o == nil || IsNil(o.Contact) {
		var ret ClientContact
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetContactOk() (*ClientContact, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *Activity) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given ClientContact and assigns it to the Contact field.
func (o *Activity) SetContact(v ClientContact) {
	o.Contact = &v
}

// GetRecurringInvoice returns the RecurringInvoice field value if set, zero value otherwise.
func (o *Activity) GetRecurringInvoice() RecurringInvoice {
	if o == nil || IsNil(o.RecurringInvoice) {
		var ret RecurringInvoice
		return ret
	}
	return *o.RecurringInvoice
}

// GetRecurringInvoiceOk returns a tuple with the RecurringInvoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetRecurringInvoiceOk() (*RecurringInvoice, bool) {
	if o == nil || IsNil(o.RecurringInvoice) {
		return nil, false
	}
	return o.RecurringInvoice, true
}

// HasRecurringInvoice returns a boolean if a field has been set.
func (o *Activity) HasRecurringInvoice() bool {
	if o != nil && !IsNil(o.RecurringInvoice) {
		return true
	}

	return false
}

// SetRecurringInvoice gets a reference to the given RecurringInvoice and assigns it to the RecurringInvoice field.
func (o *Activity) SetRecurringInvoice(v RecurringInvoice) {
	o.RecurringInvoice = &v
}

// GetInvoice returns the Invoice field value if set, zero value otherwise.
func (o *Activity) GetInvoice() Invoice {
	if o == nil || IsNil(o.Invoice) {
		var ret Invoice
		return ret
	}
	return *o.Invoice
}

// GetInvoiceOk returns a tuple with the Invoice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetInvoiceOk() (*Invoice, bool) {
	if o == nil || IsNil(o.Invoice) {
		return nil, false
	}
	return o.Invoice, true
}

// HasInvoice returns a boolean if a field has been set.
func (o *Activity) HasInvoice() bool {
	if o != nil && !IsNil(o.Invoice) {
		return true
	}

	return false
}

// SetInvoice gets a reference to the given Invoice and assigns it to the Invoice field.
func (o *Activity) SetInvoice(v Invoice) {
	o.Invoice = &v
}

// GetCredit returns the Credit field value if set, zero value otherwise.
func (o *Activity) GetCredit() Credit {
	if o == nil || IsNil(o.Credit) {
		var ret Credit
		return ret
	}
	return *o.Credit
}

// GetCreditOk returns a tuple with the Credit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetCreditOk() (*Credit, bool) {
	if o == nil || IsNil(o.Credit) {
		return nil, false
	}
	return o.Credit, true
}

// HasCredit returns a boolean if a field has been set.
func (o *Activity) HasCredit() bool {
	if o != nil && !IsNil(o.Credit) {
		return true
	}

	return false
}

// SetCredit gets a reference to the given Credit and assigns it to the Credit field.
func (o *Activity) SetCredit(v Credit) {
	o.Credit = &v
}

// GetQuote returns the Quote field value if set, zero value otherwise.
func (o *Activity) GetQuote() Quote {
	if o == nil || IsNil(o.Quote) {
		var ret Quote
		return ret
	}
	return *o.Quote
}

// GetQuoteOk returns a tuple with the Quote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetQuoteOk() (*Quote, bool) {
	if o == nil || IsNil(o.Quote) {
		return nil, false
	}
	return o.Quote, true
}

// HasQuote returns a boolean if a field has been set.
func (o *Activity) HasQuote() bool {
	if o != nil && !IsNil(o.Quote) {
		return true
	}

	return false
}

// SetQuote gets a reference to the given Quote and assigns it to the Quote field.
func (o *Activity) SetQuote(v Quote) {
	o.Quote = &v
}

// GetPayment returns the Payment field value if set, zero value otherwise.
func (o *Activity) GetPayment() Payment {
	if o == nil || IsNil(o.Payment) {
		var ret Payment
		return ret
	}
	return *o.Payment
}

// GetPaymentOk returns a tuple with the Payment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetPaymentOk() (*Payment, bool) {
	if o == nil || IsNil(o.Payment) {
		return nil, false
	}
	return o.Payment, true
}

// HasPayment returns a boolean if a field has been set.
func (o *Activity) HasPayment() bool {
	if o != nil && !IsNil(o.Payment) {
		return true
	}

	return false
}

// SetPayment gets a reference to the given Payment and assigns it to the Payment field.
func (o *Activity) SetPayment(v Payment) {
	o.Payment = &v
}

// GetExpense returns the Expense field value if set, zero value otherwise.
func (o *Activity) GetExpense() Expense {
	if o == nil || IsNil(o.Expense) {
		var ret Expense
		return ret
	}
	return *o.Expense
}

// GetExpenseOk returns a tuple with the Expense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetExpenseOk() (*Expense, bool) {
	if o == nil || IsNil(o.Expense) {
		return nil, false
	}
	return o.Expense, true
}

// HasExpense returns a boolean if a field has been set.
func (o *Activity) HasExpense() bool {
	if o != nil && !IsNil(o.Expense) {
		return true
	}

	return false
}

// SetExpense gets a reference to the given Expense and assigns it to the Expense field.
func (o *Activity) SetExpense(v Expense) {
	o.Expense = &v
}

// GetTask returns the Task field value if set, zero value otherwise.
func (o *Activity) GetTask() Task {
	if o == nil || IsNil(o.Task) {
		var ret Task
		return ret
	}
	return *o.Task
}

// GetTaskOk returns a tuple with the Task field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetTaskOk() (*Task, bool) {
	if o == nil || IsNil(o.Task) {
		return nil, false
	}
	return o.Task, true
}

// HasTask returns a boolean if a field has been set.
func (o *Activity) HasTask() bool {
	if o != nil && !IsNil(o.Task) {
		return true
	}

	return false
}

// SetTask gets a reference to the given Task and assigns it to the Task field.
func (o *Activity) SetTask(v Task) {
	o.Task = &v
}

// GetPurchaseOrder returns the PurchaseOrder field value if set, zero value otherwise.
func (o *Activity) GetPurchaseOrder() PurchaseOrder {
	if o == nil || IsNil(o.PurchaseOrder) {
		var ret PurchaseOrder
		return ret
	}
	return *o.PurchaseOrder
}

// GetPurchaseOrderOk returns a tuple with the PurchaseOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetPurchaseOrderOk() (*PurchaseOrder, bool) {
	if o == nil || IsNil(o.PurchaseOrder) {
		return nil, false
	}
	return o.PurchaseOrder, true
}

// HasPurchaseOrder returns a boolean if a field has been set.
func (o *Activity) HasPurchaseOrder() bool {
	if o != nil && !IsNil(o.PurchaseOrder) {
		return true
	}

	return false
}

// SetPurchaseOrder gets a reference to the given PurchaseOrder and assigns it to the PurchaseOrder field.
func (o *Activity) SetPurchaseOrder(v PurchaseOrder) {
	o.PurchaseOrder = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *Activity) GetVendor() Vendor {
	if o == nil || IsNil(o.Vendor) {
		var ret Vendor
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetVendorOk() (*Vendor, bool) {
	if o == nil || IsNil(o.Vendor) {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *Activity) HasVendor() bool {
	if o != nil && !IsNil(o.Vendor) {
		return true
	}

	return false
}

// SetVendor gets a reference to the given Vendor and assigns it to the Vendor field.
func (o *Activity) SetVendor(v Vendor) {
	o.Vendor = &v
}

// GetVendorContact returns the VendorContact field value if set, zero value otherwise.
func (o *Activity) GetVendorContact() VendorContact {
	if o == nil || IsNil(o.VendorContact) {
		var ret VendorContact
		return ret
	}
	return *o.VendorContact
}

// GetVendorContactOk returns a tuple with the VendorContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Activity) GetVendorContactOk() (*VendorContact, bool) {
	if o == nil || IsNil(o.VendorContact) {
		return nil, false
	}
	return o.VendorContact, true
}

// HasVendorContact returns a boolean if a field has been set.
func (o *Activity) HasVendorContact() bool {
	if o != nil && !IsNil(o.VendorContact) {
		return true
	}

	return false
}

// SetVendorContact gets a reference to the given VendorContact and assigns it to the VendorContact field.
func (o *Activity) SetVendorContact(v VendorContact) {
	o.VendorContact = &v
}

func (o Activity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Activity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ActivityTypeId) {
		toSerialize["activity_type_id"] = o.ActivityTypeId
	}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.CompanyId) {
		toSerialize["company_id"] = o.CompanyId
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.InvoiceId) {
		toSerialize["invoice_id"] = o.InvoiceId
	}
	if !IsNil(o.PaymentId) {
		toSerialize["payment_id"] = o.PaymentId
	}
	if !IsNil(o.CreditId) {
		toSerialize["credit_id"] = o.CreditId
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.ExpenseId) {
		toSerialize["expense_id"] = o.ExpenseId
	}
	if !IsNil(o.IsSystem) {
		toSerialize["is_system"] = o.IsSystem
	}
	if !IsNil(o.ContactId) {
		toSerialize["contact_id"] = o.ContactId
	}
	if !IsNil(o.TaskId) {
		toSerialize["task_id"] = o.TaskId
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.TokenId) {
		toSerialize["token_id"] = o.TokenId
	}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Client) {
		toSerialize["client"] = o.Client
	}
	if !IsNil(o.Contact) {
		toSerialize["contact"] = o.Contact
	}
	if !IsNil(o.RecurringInvoice) {
		toSerialize["recurring_invoice"] = o.RecurringInvoice
	}
	if !IsNil(o.Invoice) {
		toSerialize["invoice"] = o.Invoice
	}
	if !IsNil(o.Credit) {
		toSerialize["credit"] = o.Credit
	}
	if !IsNil(o.Quote) {
		toSerialize["quote"] = o.Quote
	}
	if !IsNil(o.Payment) {
		toSerialize["payment"] = o.Payment
	}
	if !IsNil(o.Expense) {
		toSerialize["expense"] = o.Expense
	}
	if !IsNil(o.Task) {
		toSerialize["task"] = o.Task
	}
	if !IsNil(o.PurchaseOrder) {
		toSerialize["purchase_order"] = o.PurchaseOrder
	}
	if !IsNil(o.Vendor) {
		toSerialize["vendor"] = o.Vendor
	}
	if !IsNil(o.VendorContact) {
		toSerialize["vendor_contact"] = o.VendorContact
	}
	return toSerialize, nil
}

type NullableActivity struct {
	value *Activity
	isSet bool
}

func (v NullableActivity) Get() *Activity {
	return v.value
}

func (v *NullableActivity) Set(val *Activity) {
	v.value = val
	v.isSet = true
}

func (v NullableActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivity(val *Activity) *NullableActivity {
	return &NullableActivity{value: val, isSet: true}
}

func (v NullableActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


