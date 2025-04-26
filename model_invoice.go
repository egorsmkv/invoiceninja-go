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

// checks if the Invoice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Invoice{}

// Invoice struct for Invoice
type Invoice struct {
	// The invoice hashed id
	Id *string `json:"id,omitempty"`
	// The user hashed id
	UserId *string `json:"user_id,omitempty"`
	// The assigned user hashed id
	AssignedUserId *string `json:"assigned_user_id,omitempty"`
	// The client hashed id
	ClientId *string `json:"client_id,omitempty"`
	// The invoice status variable
	StatusId *string `json:"status_id,omitempty"`
	// The invoice number - is a unique alpha numeric number per invoice per company
	Number *string `json:"number,omitempty"`
	// The purchase order associated with this invoice
	PoNumber *string `json:"po_number,omitempty"`
	// The invoice terms
	Terms *string `json:"terms,omitempty"`
	// The public notes of the invoice
	PublicNotes *string `json:"public_notes,omitempty"`
	// The private notes of the invoice
	PrivateNotes *string `json:"private_notes,omitempty"`
	// The invoice footer notes
	Footer *string `json:"footer,omitempty"`
	// A custom field value
	CustomValue1 *string `json:"custom_value1,omitempty"`
	// A custom field value
	CustomValue2 *string `json:"custom_value2,omitempty"`
	// A custom field value
	CustomValue3 *string `json:"custom_value3,omitempty"`
	// A custom field value
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
	// The total taxes for the invoice
	TotalTaxes *float32 `json:"total_taxes,omitempty"`
	// An array of objects which define the line items of the invoice
	LineItems []InvoiceItem `json:"line_items,omitempty"`
	// An array of objects which define the invitations of the invoice
	Invitations []InvoiceInvitation `json:"invitations,omitempty"`
	// The invoice amount
	Amount *float32 `json:"amount,omitempty"`
	// The invoice balance
	Balance *float32 `json:"balance,omitempty"`
	// The amount paid on the invoice to date
	PaidToDate *float32 `json:"paid_to_date,omitempty"`
	// The invoice discount, can be an amount or a percentage
	Discount *float32 `json:"discount,omitempty"`
	// The deposit/partial amount
	Partial *float32 `json:"partial,omitempty"`
	// Flag determining if the discount is an amount or a percentage
	IsAmountDiscount *bool `json:"is_amount_discount,omitempty"`
	// Defines if the invoice has been deleted
	IsDeleted *bool `json:"is_deleted,omitempty"`
	// Defines the type of taxes used as either inclusive or exclusive
	UsesInclusiveTaxes *bool `json:"uses_inclusive_taxes,omitempty"`
	// The Invoice Date
	Date *string `json:"date,omitempty"`
	// The last date the invoice was sent out
	LastSentDate *string `json:"last_sent_date,omitempty"`
	// The Next date for a reminder to be sent
	NextSendDate *string `json:"next_send_date,omitempty"`
	// The due date for the deposit/partial amount
	PartialDueDate *string `json:"partial_due_date,omitempty"`
	// The due date of the invoice
	DueDate *string `json:"due_date,omitempty"`
	// Timestamp
	LastViewed *float32 `json:"last_viewed,omitempty"`
	// Timestamp
	UpdatedAt *float32 `json:"updated_at,omitempty"`
	// Timestamp
	ArchivedAt *float32 `json:"archived_at,omitempty"`
	// First Custom Surcharge
	CustomSurcharge1 *float32 `json:"custom_surcharge1,omitempty"`
	// Second Custom Surcharge
	CustomSurcharge2 *float32 `json:"custom_surcharge2,omitempty"`
	// Third Custom Surcharge
	CustomSurcharge3 *float32 `json:"custom_surcharge3,omitempty"`
	// Fourth Custom Surcharge
	CustomSurcharge4 *float32 `json:"custom_surcharge4,omitempty"`
	// Toggles charging taxes on custom surcharge amounts
	CustomSurchargeTax1 *bool `json:"custom_surcharge_tax1,omitempty"`
	// Toggles charging taxes on custom surcharge amounts
	CustomSurchargeTax2 *bool `json:"custom_surcharge_tax2,omitempty"`
	// Toggles charging taxes on custom surcharge amounts
	CustomSurchargeTax3 *bool `json:"custom_surcharge_tax3,omitempty"`
	// Toggles charging taxes on custom surcharge amounts
	CustomSurchargeTax4 *bool `json:"custom_surcharge_tax4,omitempty"`
	// The project associated with this invoice
	ProjectId *string `json:"project_id,omitempty"`
	// The number of times the invoice has attempted to be auto billed
	AutoBillTries *int32 `json:"auto_bill_tries,omitempty"`
	// Boolean flag determining if the invoice is set to auto bill
	AutoBillEnabled *bool `json:"auto_bill_enabled,omitempty"`
	// The subscription associated with this invoice
	SubscriptionId *string `json:"subscription_id,omitempty"`
	// The client location id that this invoice relates to
	LocationId *string `json:"location_id,omitempty"`
}

// NewInvoice instantiates a new Invoice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoice() *Invoice {
	this := Invoice{}
	return &this
}

// NewInvoiceWithDefaults instantiates a new Invoice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceWithDefaults() *Invoice {
	this := Invoice{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Invoice) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Invoice) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Invoice) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *Invoice) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *Invoice) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *Invoice) SetUserId(v string) {
	o.UserId = &v
}

// GetAssignedUserId returns the AssignedUserId field value if set, zero value otherwise.
func (o *Invoice) GetAssignedUserId() string {
	if o == nil || IsNil(o.AssignedUserId) {
		var ret string
		return ret
	}
	return *o.AssignedUserId
}

// GetAssignedUserIdOk returns a tuple with the AssignedUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetAssignedUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssignedUserId) {
		return nil, false
	}
	return o.AssignedUserId, true
}

// HasAssignedUserId returns a boolean if a field has been set.
func (o *Invoice) HasAssignedUserId() bool {
	if o != nil && !IsNil(o.AssignedUserId) {
		return true
	}

	return false
}

// SetAssignedUserId gets a reference to the given string and assigns it to the AssignedUserId field.
func (o *Invoice) SetAssignedUserId(v string) {
	o.AssignedUserId = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *Invoice) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *Invoice) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *Invoice) SetClientId(v string) {
	o.ClientId = &v
}

// GetStatusId returns the StatusId field value if set, zero value otherwise.
func (o *Invoice) GetStatusId() string {
	if o == nil || IsNil(o.StatusId) {
		var ret string
		return ret
	}
	return *o.StatusId
}

// GetStatusIdOk returns a tuple with the StatusId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetStatusIdOk() (*string, bool) {
	if o == nil || IsNil(o.StatusId) {
		return nil, false
	}
	return o.StatusId, true
}

// HasStatusId returns a boolean if a field has been set.
func (o *Invoice) HasStatusId() bool {
	if o != nil && !IsNil(o.StatusId) {
		return true
	}

	return false
}

// SetStatusId gets a reference to the given string and assigns it to the StatusId field.
func (o *Invoice) SetStatusId(v string) {
	o.StatusId = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *Invoice) GetNumber() string {
	if o == nil || IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetNumberOk() (*string, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *Invoice) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *Invoice) SetNumber(v string) {
	o.Number = &v
}

// GetPoNumber returns the PoNumber field value if set, zero value otherwise.
func (o *Invoice) GetPoNumber() string {
	if o == nil || IsNil(o.PoNumber) {
		var ret string
		return ret
	}
	return *o.PoNumber
}

// GetPoNumberOk returns a tuple with the PoNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetPoNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PoNumber) {
		return nil, false
	}
	return o.PoNumber, true
}

// HasPoNumber returns a boolean if a field has been set.
func (o *Invoice) HasPoNumber() bool {
	if o != nil && !IsNil(o.PoNumber) {
		return true
	}

	return false
}

// SetPoNumber gets a reference to the given string and assigns it to the PoNumber field.
func (o *Invoice) SetPoNumber(v string) {
	o.PoNumber = &v
}

// GetTerms returns the Terms field value if set, zero value otherwise.
func (o *Invoice) GetTerms() string {
	if o == nil || IsNil(o.Terms) {
		var ret string
		return ret
	}
	return *o.Terms
}

// GetTermsOk returns a tuple with the Terms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetTermsOk() (*string, bool) {
	if o == nil || IsNil(o.Terms) {
		return nil, false
	}
	return o.Terms, true
}

// HasTerms returns a boolean if a field has been set.
func (o *Invoice) HasTerms() bool {
	if o != nil && !IsNil(o.Terms) {
		return true
	}

	return false
}

// SetTerms gets a reference to the given string and assigns it to the Terms field.
func (o *Invoice) SetTerms(v string) {
	o.Terms = &v
}

// GetPublicNotes returns the PublicNotes field value if set, zero value otherwise.
func (o *Invoice) GetPublicNotes() string {
	if o == nil || IsNil(o.PublicNotes) {
		var ret string
		return ret
	}
	return *o.PublicNotes
}

// GetPublicNotesOk returns a tuple with the PublicNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetPublicNotesOk() (*string, bool) {
	if o == nil || IsNil(o.PublicNotes) {
		return nil, false
	}
	return o.PublicNotes, true
}

// HasPublicNotes returns a boolean if a field has been set.
func (o *Invoice) HasPublicNotes() bool {
	if o != nil && !IsNil(o.PublicNotes) {
		return true
	}

	return false
}

// SetPublicNotes gets a reference to the given string and assigns it to the PublicNotes field.
func (o *Invoice) SetPublicNotes(v string) {
	o.PublicNotes = &v
}

// GetPrivateNotes returns the PrivateNotes field value if set, zero value otherwise.
func (o *Invoice) GetPrivateNotes() string {
	if o == nil || IsNil(o.PrivateNotes) {
		var ret string
		return ret
	}
	return *o.PrivateNotes
}

// GetPrivateNotesOk returns a tuple with the PrivateNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetPrivateNotesOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateNotes) {
		return nil, false
	}
	return o.PrivateNotes, true
}

// HasPrivateNotes returns a boolean if a field has been set.
func (o *Invoice) HasPrivateNotes() bool {
	if o != nil && !IsNil(o.PrivateNotes) {
		return true
	}

	return false
}

// SetPrivateNotes gets a reference to the given string and assigns it to the PrivateNotes field.
func (o *Invoice) SetPrivateNotes(v string) {
	o.PrivateNotes = &v
}

// GetFooter returns the Footer field value if set, zero value otherwise.
func (o *Invoice) GetFooter() string {
	if o == nil || IsNil(o.Footer) {
		var ret string
		return ret
	}
	return *o.Footer
}

// GetFooterOk returns a tuple with the Footer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetFooterOk() (*string, bool) {
	if o == nil || IsNil(o.Footer) {
		return nil, false
	}
	return o.Footer, true
}

// HasFooter returns a boolean if a field has been set.
func (o *Invoice) HasFooter() bool {
	if o != nil && !IsNil(o.Footer) {
		return true
	}

	return false
}

// SetFooter gets a reference to the given string and assigns it to the Footer field.
func (o *Invoice) SetFooter(v string) {
	o.Footer = &v
}

// GetCustomValue1 returns the CustomValue1 field value if set, zero value otherwise.
func (o *Invoice) GetCustomValue1() string {
	if o == nil || IsNil(o.CustomValue1) {
		var ret string
		return ret
	}
	return *o.CustomValue1
}

// GetCustomValue1Ok returns a tuple with the CustomValue1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetCustomValue1Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue1) {
		return nil, false
	}
	return o.CustomValue1, true
}

// HasCustomValue1 returns a boolean if a field has been set.
func (o *Invoice) HasCustomValue1() bool {
	if o != nil && !IsNil(o.CustomValue1) {
		return true
	}

	return false
}

// SetCustomValue1 gets a reference to the given string and assigns it to the CustomValue1 field.
func (o *Invoice) SetCustomValue1(v string) {
	o.CustomValue1 = &v
}

// GetCustomValue2 returns the CustomValue2 field value if set, zero value otherwise.
func (o *Invoice) GetCustomValue2() string {
	if o == nil || IsNil(o.CustomValue2) {
		var ret string
		return ret
	}
	return *o.CustomValue2
}

// GetCustomValue2Ok returns a tuple with the CustomValue2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetCustomValue2Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue2) {
		return nil, false
	}
	return o.CustomValue2, true
}

// HasCustomValue2 returns a boolean if a field has been set.
func (o *Invoice) HasCustomValue2() bool {
	if o != nil && !IsNil(o.CustomValue2) {
		return true
	}

	return false
}

// SetCustomValue2 gets a reference to the given string and assigns it to the CustomValue2 field.
func (o *Invoice) SetCustomValue2(v string) {
	o.CustomValue2 = &v
}

// GetCustomValue3 returns the CustomValue3 field value if set, zero value otherwise.
func (o *Invoice) GetCustomValue3() string {
	if o == nil || IsNil(o.CustomValue3) {
		var ret string
		return ret
	}
	return *o.CustomValue3
}

// GetCustomValue3Ok returns a tuple with the CustomValue3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetCustomValue3Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue3) {
		return nil, false
	}
	return o.CustomValue3, true
}

// HasCustomValue3 returns a boolean if a field has been set.
func (o *Invoice) HasCustomValue3() bool {
	if o != nil && !IsNil(o.CustomValue3) {
		return true
	}

	return false
}

// SetCustomValue3 gets a reference to the given string and assigns it to the CustomValue3 field.
func (o *Invoice) SetCustomValue3(v string) {
	o.CustomValue3 = &v
}

// GetCustomValue4 returns the CustomValue4 field value if set, zero value otherwise.
func (o *Invoice) GetCustomValue4() string {
	if o == nil || IsNil(o.CustomValue4) {
		var ret string
		return ret
	}
	return *o.CustomValue4
}

// GetCustomValue4Ok returns a tuple with the CustomValue4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetCustomValue4Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue4) {
		return nil, false
	}
	return o.CustomValue4, true
}

// HasCustomValue4 returns a boolean if a field has been set.
func (o *Invoice) HasCustomValue4() bool {
	if o != nil && !IsNil(o.CustomValue4) {
		return true
	}

	return false
}

// SetCustomValue4 gets a reference to the given string and assigns it to the CustomValue4 field.
func (o *Invoice) SetCustomValue4(v string) {
	o.CustomValue4 = &v
}

// GetTaxName1 returns the TaxName1 field value if set, zero value otherwise.
func (o *Invoice) GetTaxName1() string {
	if o == nil || IsNil(o.TaxName1) {
		var ret string
		return ret
	}
	return *o.TaxName1
}

// GetTaxName1Ok returns a tuple with the TaxName1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetTaxName1Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName1) {
		return nil, false
	}
	return o.TaxName1, true
}

// HasTaxName1 returns a boolean if a field has been set.
func (o *Invoice) HasTaxName1() bool {
	if o != nil && !IsNil(o.TaxName1) {
		return true
	}

	return false
}

// SetTaxName1 gets a reference to the given string and assigns it to the TaxName1 field.
func (o *Invoice) SetTaxName1(v string) {
	o.TaxName1 = &v
}

// GetTaxName2 returns the TaxName2 field value if set, zero value otherwise.
func (o *Invoice) GetTaxName2() string {
	if o == nil || IsNil(o.TaxName2) {
		var ret string
		return ret
	}
	return *o.TaxName2
}

// GetTaxName2Ok returns a tuple with the TaxName2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetTaxName2Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName2) {
		return nil, false
	}
	return o.TaxName2, true
}

// HasTaxName2 returns a boolean if a field has been set.
func (o *Invoice) HasTaxName2() bool {
	if o != nil && !IsNil(o.TaxName2) {
		return true
	}

	return false
}

// SetTaxName2 gets a reference to the given string and assigns it to the TaxName2 field.
func (o *Invoice) SetTaxName2(v string) {
	o.TaxName2 = &v
}

// GetTaxRate1 returns the TaxRate1 field value if set, zero value otherwise.
func (o *Invoice) GetTaxRate1() float32 {
	if o == nil || IsNil(o.TaxRate1) {
		var ret float32
		return ret
	}
	return *o.TaxRate1
}

// GetTaxRate1Ok returns a tuple with the TaxRate1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetTaxRate1Ok() (*float32, bool) {
	if o == nil || IsNil(o.TaxRate1) {
		return nil, false
	}
	return o.TaxRate1, true
}

// HasTaxRate1 returns a boolean if a field has been set.
func (o *Invoice) HasTaxRate1() bool {
	if o != nil && !IsNil(o.TaxRate1) {
		return true
	}

	return false
}

// SetTaxRate1 gets a reference to the given float32 and assigns it to the TaxRate1 field.
func (o *Invoice) SetTaxRate1(v float32) {
	o.TaxRate1 = &v
}

// GetTaxRate2 returns the TaxRate2 field value if set, zero value otherwise.
func (o *Invoice) GetTaxRate2() float32 {
	if o == nil || IsNil(o.TaxRate2) {
		var ret float32
		return ret
	}
	return *o.TaxRate2
}

// GetTaxRate2Ok returns a tuple with the TaxRate2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetTaxRate2Ok() (*float32, bool) {
	if o == nil || IsNil(o.TaxRate2) {
		return nil, false
	}
	return o.TaxRate2, true
}

// HasTaxRate2 returns a boolean if a field has been set.
func (o *Invoice) HasTaxRate2() bool {
	if o != nil && !IsNil(o.TaxRate2) {
		return true
	}

	return false
}

// SetTaxRate2 gets a reference to the given float32 and assigns it to the TaxRate2 field.
func (o *Invoice) SetTaxRate2(v float32) {
	o.TaxRate2 = &v
}

// GetTaxName3 returns the TaxName3 field value if set, zero value otherwise.
func (o *Invoice) GetTaxName3() string {
	if o == nil || IsNil(o.TaxName3) {
		var ret string
		return ret
	}
	return *o.TaxName3
}

// GetTaxName3Ok returns a tuple with the TaxName3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetTaxName3Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName3) {
		return nil, false
	}
	return o.TaxName3, true
}

// HasTaxName3 returns a boolean if a field has been set.
func (o *Invoice) HasTaxName3() bool {
	if o != nil && !IsNil(o.TaxName3) {
		return true
	}

	return false
}

// SetTaxName3 gets a reference to the given string and assigns it to the TaxName3 field.
func (o *Invoice) SetTaxName3(v string) {
	o.TaxName3 = &v
}

// GetTaxRate3 returns the TaxRate3 field value if set, zero value otherwise.
func (o *Invoice) GetTaxRate3() float32 {
	if o == nil || IsNil(o.TaxRate3) {
		var ret float32
		return ret
	}
	return *o.TaxRate3
}

// GetTaxRate3Ok returns a tuple with the TaxRate3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetTaxRate3Ok() (*float32, bool) {
	if o == nil || IsNil(o.TaxRate3) {
		return nil, false
	}
	return o.TaxRate3, true
}

// HasTaxRate3 returns a boolean if a field has been set.
func (o *Invoice) HasTaxRate3() bool {
	if o != nil && !IsNil(o.TaxRate3) {
		return true
	}

	return false
}

// SetTaxRate3 gets a reference to the given float32 and assigns it to the TaxRate3 field.
func (o *Invoice) SetTaxRate3(v float32) {
	o.TaxRate3 = &v
}

// GetTotalTaxes returns the TotalTaxes field value if set, zero value otherwise.
func (o *Invoice) GetTotalTaxes() float32 {
	if o == nil || IsNil(o.TotalTaxes) {
		var ret float32
		return ret
	}
	return *o.TotalTaxes
}

// GetTotalTaxesOk returns a tuple with the TotalTaxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetTotalTaxesOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalTaxes) {
		return nil, false
	}
	return o.TotalTaxes, true
}

// HasTotalTaxes returns a boolean if a field has been set.
func (o *Invoice) HasTotalTaxes() bool {
	if o != nil && !IsNil(o.TotalTaxes) {
		return true
	}

	return false
}

// SetTotalTaxes gets a reference to the given float32 and assigns it to the TotalTaxes field.
func (o *Invoice) SetTotalTaxes(v float32) {
	o.TotalTaxes = &v
}

// GetLineItems returns the LineItems field value if set, zero value otherwise.
func (o *Invoice) GetLineItems() []InvoiceItem {
	if o == nil || IsNil(o.LineItems) {
		var ret []InvoiceItem
		return ret
	}
	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetLineItemsOk() ([]InvoiceItem, bool) {
	if o == nil || IsNil(o.LineItems) {
		return nil, false
	}
	return o.LineItems, true
}

// HasLineItems returns a boolean if a field has been set.
func (o *Invoice) HasLineItems() bool {
	if o != nil && !IsNil(o.LineItems) {
		return true
	}

	return false
}

// SetLineItems gets a reference to the given []InvoiceItem and assigns it to the LineItems field.
func (o *Invoice) SetLineItems(v []InvoiceItem) {
	o.LineItems = v
}

// GetInvitations returns the Invitations field value if set, zero value otherwise.
func (o *Invoice) GetInvitations() []InvoiceInvitation {
	if o == nil || IsNil(o.Invitations) {
		var ret []InvoiceInvitation
		return ret
	}
	return o.Invitations
}

// GetInvitationsOk returns a tuple with the Invitations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetInvitationsOk() ([]InvoiceInvitation, bool) {
	if o == nil || IsNil(o.Invitations) {
		return nil, false
	}
	return o.Invitations, true
}

// HasInvitations returns a boolean if a field has been set.
func (o *Invoice) HasInvitations() bool {
	if o != nil && !IsNil(o.Invitations) {
		return true
	}

	return false
}

// SetInvitations gets a reference to the given []InvoiceInvitation and assigns it to the Invitations field.
func (o *Invoice) SetInvitations(v []InvoiceInvitation) {
	o.Invitations = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Invoice) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Invoice) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *Invoice) SetAmount(v float32) {
	o.Amount = &v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *Invoice) GetBalance() float32 {
	if o == nil || IsNil(o.Balance) {
		var ret float32
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetBalanceOk() (*float32, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *Invoice) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given float32 and assigns it to the Balance field.
func (o *Invoice) SetBalance(v float32) {
	o.Balance = &v
}

// GetPaidToDate returns the PaidToDate field value if set, zero value otherwise.
func (o *Invoice) GetPaidToDate() float32 {
	if o == nil || IsNil(o.PaidToDate) {
		var ret float32
		return ret
	}
	return *o.PaidToDate
}

// GetPaidToDateOk returns a tuple with the PaidToDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetPaidToDateOk() (*float32, bool) {
	if o == nil || IsNil(o.PaidToDate) {
		return nil, false
	}
	return o.PaidToDate, true
}

// HasPaidToDate returns a boolean if a field has been set.
func (o *Invoice) HasPaidToDate() bool {
	if o != nil && !IsNil(o.PaidToDate) {
		return true
	}

	return false
}

// SetPaidToDate gets a reference to the given float32 and assigns it to the PaidToDate field.
func (o *Invoice) SetPaidToDate(v float32) {
	o.PaidToDate = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *Invoice) GetDiscount() float32 {
	if o == nil || IsNil(o.Discount) {
		var ret float32
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetDiscountOk() (*float32, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *Invoice) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given float32 and assigns it to the Discount field.
func (o *Invoice) SetDiscount(v float32) {
	o.Discount = &v
}

// GetPartial returns the Partial field value if set, zero value otherwise.
func (o *Invoice) GetPartial() float32 {
	if o == nil || IsNil(o.Partial) {
		var ret float32
		return ret
	}
	return *o.Partial
}

// GetPartialOk returns a tuple with the Partial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetPartialOk() (*float32, bool) {
	if o == nil || IsNil(o.Partial) {
		return nil, false
	}
	return o.Partial, true
}

// HasPartial returns a boolean if a field has been set.
func (o *Invoice) HasPartial() bool {
	if o != nil && !IsNil(o.Partial) {
		return true
	}

	return false
}

// SetPartial gets a reference to the given float32 and assigns it to the Partial field.
func (o *Invoice) SetPartial(v float32) {
	o.Partial = &v
}

// GetIsAmountDiscount returns the IsAmountDiscount field value if set, zero value otherwise.
func (o *Invoice) GetIsAmountDiscount() bool {
	if o == nil || IsNil(o.IsAmountDiscount) {
		var ret bool
		return ret
	}
	return *o.IsAmountDiscount
}

// GetIsAmountDiscountOk returns a tuple with the IsAmountDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetIsAmountDiscountOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAmountDiscount) {
		return nil, false
	}
	return o.IsAmountDiscount, true
}

// HasIsAmountDiscount returns a boolean if a field has been set.
func (o *Invoice) HasIsAmountDiscount() bool {
	if o != nil && !IsNil(o.IsAmountDiscount) {
		return true
	}

	return false
}

// SetIsAmountDiscount gets a reference to the given bool and assigns it to the IsAmountDiscount field.
func (o *Invoice) SetIsAmountDiscount(v bool) {
	o.IsAmountDiscount = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *Invoice) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *Invoice) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *Invoice) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetUsesInclusiveTaxes returns the UsesInclusiveTaxes field value if set, zero value otherwise.
func (o *Invoice) GetUsesInclusiveTaxes() bool {
	if o == nil || IsNil(o.UsesInclusiveTaxes) {
		var ret bool
		return ret
	}
	return *o.UsesInclusiveTaxes
}

// GetUsesInclusiveTaxesOk returns a tuple with the UsesInclusiveTaxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetUsesInclusiveTaxesOk() (*bool, bool) {
	if o == nil || IsNil(o.UsesInclusiveTaxes) {
		return nil, false
	}
	return o.UsesInclusiveTaxes, true
}

// HasUsesInclusiveTaxes returns a boolean if a field has been set.
func (o *Invoice) HasUsesInclusiveTaxes() bool {
	if o != nil && !IsNil(o.UsesInclusiveTaxes) {
		return true
	}

	return false
}

// SetUsesInclusiveTaxes gets a reference to the given bool and assigns it to the UsesInclusiveTaxes field.
func (o *Invoice) SetUsesInclusiveTaxes(v bool) {
	o.UsesInclusiveTaxes = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *Invoice) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *Invoice) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *Invoice) SetDate(v string) {
	o.Date = &v
}

// GetLastSentDate returns the LastSentDate field value if set, zero value otherwise.
func (o *Invoice) GetLastSentDate() string {
	if o == nil || IsNil(o.LastSentDate) {
		var ret string
		return ret
	}
	return *o.LastSentDate
}

// GetLastSentDateOk returns a tuple with the LastSentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetLastSentDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastSentDate) {
		return nil, false
	}
	return o.LastSentDate, true
}

// HasLastSentDate returns a boolean if a field has been set.
func (o *Invoice) HasLastSentDate() bool {
	if o != nil && !IsNil(o.LastSentDate) {
		return true
	}

	return false
}

// SetLastSentDate gets a reference to the given string and assigns it to the LastSentDate field.
func (o *Invoice) SetLastSentDate(v string) {
	o.LastSentDate = &v
}

// GetNextSendDate returns the NextSendDate field value if set, zero value otherwise.
func (o *Invoice) GetNextSendDate() string {
	if o == nil || IsNil(o.NextSendDate) {
		var ret string
		return ret
	}
	return *o.NextSendDate
}

// GetNextSendDateOk returns a tuple with the NextSendDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetNextSendDateOk() (*string, bool) {
	if o == nil || IsNil(o.NextSendDate) {
		return nil, false
	}
	return o.NextSendDate, true
}

// HasNextSendDate returns a boolean if a field has been set.
func (o *Invoice) HasNextSendDate() bool {
	if o != nil && !IsNil(o.NextSendDate) {
		return true
	}

	return false
}

// SetNextSendDate gets a reference to the given string and assigns it to the NextSendDate field.
func (o *Invoice) SetNextSendDate(v string) {
	o.NextSendDate = &v
}

// GetPartialDueDate returns the PartialDueDate field value if set, zero value otherwise.
func (o *Invoice) GetPartialDueDate() string {
	if o == nil || IsNil(o.PartialDueDate) {
		var ret string
		return ret
	}
	return *o.PartialDueDate
}

// GetPartialDueDateOk returns a tuple with the PartialDueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetPartialDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.PartialDueDate) {
		return nil, false
	}
	return o.PartialDueDate, true
}

// HasPartialDueDate returns a boolean if a field has been set.
func (o *Invoice) HasPartialDueDate() bool {
	if o != nil && !IsNil(o.PartialDueDate) {
		return true
	}

	return false
}

// SetPartialDueDate gets a reference to the given string and assigns it to the PartialDueDate field.
func (o *Invoice) SetPartialDueDate(v string) {
	o.PartialDueDate = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *Invoice) GetDueDate() string {
	if o == nil || IsNil(o.DueDate) {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *Invoice) HasDueDate() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *Invoice) SetDueDate(v string) {
	o.DueDate = &v
}

// GetLastViewed returns the LastViewed field value if set, zero value otherwise.
func (o *Invoice) GetLastViewed() float32 {
	if o == nil || IsNil(o.LastViewed) {
		var ret float32
		return ret
	}
	return *o.LastViewed
}

// GetLastViewedOk returns a tuple with the LastViewed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetLastViewedOk() (*float32, bool) {
	if o == nil || IsNil(o.LastViewed) {
		return nil, false
	}
	return o.LastViewed, true
}

// HasLastViewed returns a boolean if a field has been set.
func (o *Invoice) HasLastViewed() bool {
	if o != nil && !IsNil(o.LastViewed) {
		return true
	}

	return false
}

// SetLastViewed gets a reference to the given float32 and assigns it to the LastViewed field.
func (o *Invoice) SetLastViewed(v float32) {
	o.LastViewed = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Invoice) GetUpdatedAt() float32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret float32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetUpdatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Invoice) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given float32 and assigns it to the UpdatedAt field.
func (o *Invoice) SetUpdatedAt(v float32) {
	o.UpdatedAt = &v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *Invoice) GetArchivedAt() float32 {
	if o == nil || IsNil(o.ArchivedAt) {
		var ret float32
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetArchivedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.ArchivedAt) {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *Invoice) HasArchivedAt() bool {
	if o != nil && !IsNil(o.ArchivedAt) {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given float32 and assigns it to the ArchivedAt field.
func (o *Invoice) SetArchivedAt(v float32) {
	o.ArchivedAt = &v
}

// GetCustomSurcharge1 returns the CustomSurcharge1 field value if set, zero value otherwise.
func (o *Invoice) GetCustomSurcharge1() float32 {
	if o == nil || IsNil(o.CustomSurcharge1) {
		var ret float32
		return ret
	}
	return *o.CustomSurcharge1
}

// GetCustomSurcharge1Ok returns a tuple with the CustomSurcharge1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetCustomSurcharge1Ok() (*float32, bool) {
	if o == nil || IsNil(o.CustomSurcharge1) {
		return nil, false
	}
	return o.CustomSurcharge1, true
}

// HasCustomSurcharge1 returns a boolean if a field has been set.
func (o *Invoice) HasCustomSurcharge1() bool {
	if o != nil && !IsNil(o.CustomSurcharge1) {
		return true
	}

	return false
}

// SetCustomSurcharge1 gets a reference to the given float32 and assigns it to the CustomSurcharge1 field.
func (o *Invoice) SetCustomSurcharge1(v float32) {
	o.CustomSurcharge1 = &v
}

// GetCustomSurcharge2 returns the CustomSurcharge2 field value if set, zero value otherwise.
func (o *Invoice) GetCustomSurcharge2() float32 {
	if o == nil || IsNil(o.CustomSurcharge2) {
		var ret float32
		return ret
	}
	return *o.CustomSurcharge2
}

// GetCustomSurcharge2Ok returns a tuple with the CustomSurcharge2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetCustomSurcharge2Ok() (*float32, bool) {
	if o == nil || IsNil(o.CustomSurcharge2) {
		return nil, false
	}
	return o.CustomSurcharge2, true
}

// HasCustomSurcharge2 returns a boolean if a field has been set.
func (o *Invoice) HasCustomSurcharge2() bool {
	if o != nil && !IsNil(o.CustomSurcharge2) {
		return true
	}

	return false
}

// SetCustomSurcharge2 gets a reference to the given float32 and assigns it to the CustomSurcharge2 field.
func (o *Invoice) SetCustomSurcharge2(v float32) {
	o.CustomSurcharge2 = &v
}

// GetCustomSurcharge3 returns the CustomSurcharge3 field value if set, zero value otherwise.
func (o *Invoice) GetCustomSurcharge3() float32 {
	if o == nil || IsNil(o.CustomSurcharge3) {
		var ret float32
		return ret
	}
	return *o.CustomSurcharge3
}

// GetCustomSurcharge3Ok returns a tuple with the CustomSurcharge3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetCustomSurcharge3Ok() (*float32, bool) {
	if o == nil || IsNil(o.CustomSurcharge3) {
		return nil, false
	}
	return o.CustomSurcharge3, true
}

// HasCustomSurcharge3 returns a boolean if a field has been set.
func (o *Invoice) HasCustomSurcharge3() bool {
	if o != nil && !IsNil(o.CustomSurcharge3) {
		return true
	}

	return false
}

// SetCustomSurcharge3 gets a reference to the given float32 and assigns it to the CustomSurcharge3 field.
func (o *Invoice) SetCustomSurcharge3(v float32) {
	o.CustomSurcharge3 = &v
}

// GetCustomSurcharge4 returns the CustomSurcharge4 field value if set, zero value otherwise.
func (o *Invoice) GetCustomSurcharge4() float32 {
	if o == nil || IsNil(o.CustomSurcharge4) {
		var ret float32
		return ret
	}
	return *o.CustomSurcharge4
}

// GetCustomSurcharge4Ok returns a tuple with the CustomSurcharge4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetCustomSurcharge4Ok() (*float32, bool) {
	if o == nil || IsNil(o.CustomSurcharge4) {
		return nil, false
	}
	return o.CustomSurcharge4, true
}

// HasCustomSurcharge4 returns a boolean if a field has been set.
func (o *Invoice) HasCustomSurcharge4() bool {
	if o != nil && !IsNil(o.CustomSurcharge4) {
		return true
	}

	return false
}

// SetCustomSurcharge4 gets a reference to the given float32 and assigns it to the CustomSurcharge4 field.
func (o *Invoice) SetCustomSurcharge4(v float32) {
	o.CustomSurcharge4 = &v
}

// GetCustomSurchargeTax1 returns the CustomSurchargeTax1 field value if set, zero value otherwise.
func (o *Invoice) GetCustomSurchargeTax1() bool {
	if o == nil || IsNil(o.CustomSurchargeTax1) {
		var ret bool
		return ret
	}
	return *o.CustomSurchargeTax1
}

// GetCustomSurchargeTax1Ok returns a tuple with the CustomSurchargeTax1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetCustomSurchargeTax1Ok() (*bool, bool) {
	if o == nil || IsNil(o.CustomSurchargeTax1) {
		return nil, false
	}
	return o.CustomSurchargeTax1, true
}

// HasCustomSurchargeTax1 returns a boolean if a field has been set.
func (o *Invoice) HasCustomSurchargeTax1() bool {
	if o != nil && !IsNil(o.CustomSurchargeTax1) {
		return true
	}

	return false
}

// SetCustomSurchargeTax1 gets a reference to the given bool and assigns it to the CustomSurchargeTax1 field.
func (o *Invoice) SetCustomSurchargeTax1(v bool) {
	o.CustomSurchargeTax1 = &v
}

// GetCustomSurchargeTax2 returns the CustomSurchargeTax2 field value if set, zero value otherwise.
func (o *Invoice) GetCustomSurchargeTax2() bool {
	if o == nil || IsNil(o.CustomSurchargeTax2) {
		var ret bool
		return ret
	}
	return *o.CustomSurchargeTax2
}

// GetCustomSurchargeTax2Ok returns a tuple with the CustomSurchargeTax2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetCustomSurchargeTax2Ok() (*bool, bool) {
	if o == nil || IsNil(o.CustomSurchargeTax2) {
		return nil, false
	}
	return o.CustomSurchargeTax2, true
}

// HasCustomSurchargeTax2 returns a boolean if a field has been set.
func (o *Invoice) HasCustomSurchargeTax2() bool {
	if o != nil && !IsNil(o.CustomSurchargeTax2) {
		return true
	}

	return false
}

// SetCustomSurchargeTax2 gets a reference to the given bool and assigns it to the CustomSurchargeTax2 field.
func (o *Invoice) SetCustomSurchargeTax2(v bool) {
	o.CustomSurchargeTax2 = &v
}

// GetCustomSurchargeTax3 returns the CustomSurchargeTax3 field value if set, zero value otherwise.
func (o *Invoice) GetCustomSurchargeTax3() bool {
	if o == nil || IsNil(o.CustomSurchargeTax3) {
		var ret bool
		return ret
	}
	return *o.CustomSurchargeTax3
}

// GetCustomSurchargeTax3Ok returns a tuple with the CustomSurchargeTax3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetCustomSurchargeTax3Ok() (*bool, bool) {
	if o == nil || IsNil(o.CustomSurchargeTax3) {
		return nil, false
	}
	return o.CustomSurchargeTax3, true
}

// HasCustomSurchargeTax3 returns a boolean if a field has been set.
func (o *Invoice) HasCustomSurchargeTax3() bool {
	if o != nil && !IsNil(o.CustomSurchargeTax3) {
		return true
	}

	return false
}

// SetCustomSurchargeTax3 gets a reference to the given bool and assigns it to the CustomSurchargeTax3 field.
func (o *Invoice) SetCustomSurchargeTax3(v bool) {
	o.CustomSurchargeTax3 = &v
}

// GetCustomSurchargeTax4 returns the CustomSurchargeTax4 field value if set, zero value otherwise.
func (o *Invoice) GetCustomSurchargeTax4() bool {
	if o == nil || IsNil(o.CustomSurchargeTax4) {
		var ret bool
		return ret
	}
	return *o.CustomSurchargeTax4
}

// GetCustomSurchargeTax4Ok returns a tuple with the CustomSurchargeTax4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetCustomSurchargeTax4Ok() (*bool, bool) {
	if o == nil || IsNil(o.CustomSurchargeTax4) {
		return nil, false
	}
	return o.CustomSurchargeTax4, true
}

// HasCustomSurchargeTax4 returns a boolean if a field has been set.
func (o *Invoice) HasCustomSurchargeTax4() bool {
	if o != nil && !IsNil(o.CustomSurchargeTax4) {
		return true
	}

	return false
}

// SetCustomSurchargeTax4 gets a reference to the given bool and assigns it to the CustomSurchargeTax4 field.
func (o *Invoice) SetCustomSurchargeTax4(v bool) {
	o.CustomSurchargeTax4 = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *Invoice) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *Invoice) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *Invoice) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetAutoBillTries returns the AutoBillTries field value if set, zero value otherwise.
func (o *Invoice) GetAutoBillTries() int32 {
	if o == nil || IsNil(o.AutoBillTries) {
		var ret int32
		return ret
	}
	return *o.AutoBillTries
}

// GetAutoBillTriesOk returns a tuple with the AutoBillTries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetAutoBillTriesOk() (*int32, bool) {
	if o == nil || IsNil(o.AutoBillTries) {
		return nil, false
	}
	return o.AutoBillTries, true
}

// HasAutoBillTries returns a boolean if a field has been set.
func (o *Invoice) HasAutoBillTries() bool {
	if o != nil && !IsNil(o.AutoBillTries) {
		return true
	}

	return false
}

// SetAutoBillTries gets a reference to the given int32 and assigns it to the AutoBillTries field.
func (o *Invoice) SetAutoBillTries(v int32) {
	o.AutoBillTries = &v
}

// GetAutoBillEnabled returns the AutoBillEnabled field value if set, zero value otherwise.
func (o *Invoice) GetAutoBillEnabled() bool {
	if o == nil || IsNil(o.AutoBillEnabled) {
		var ret bool
		return ret
	}
	return *o.AutoBillEnabled
}

// GetAutoBillEnabledOk returns a tuple with the AutoBillEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetAutoBillEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoBillEnabled) {
		return nil, false
	}
	return o.AutoBillEnabled, true
}

// HasAutoBillEnabled returns a boolean if a field has been set.
func (o *Invoice) HasAutoBillEnabled() bool {
	if o != nil && !IsNil(o.AutoBillEnabled) {
		return true
	}

	return false
}

// SetAutoBillEnabled gets a reference to the given bool and assigns it to the AutoBillEnabled field.
func (o *Invoice) SetAutoBillEnabled(v bool) {
	o.AutoBillEnabled = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *Invoice) GetSubscriptionId() string {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret string
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetSubscriptionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *Invoice) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given string and assigns it to the SubscriptionId field.
func (o *Invoice) SetSubscriptionId(v string) {
	o.SubscriptionId = &v
}

// GetLocationId returns the LocationId field value if set, zero value otherwise.
func (o *Invoice) GetLocationId() string {
	if o == nil || IsNil(o.LocationId) {
		var ret string
		return ret
	}
	return *o.LocationId
}

// GetLocationIdOk returns a tuple with the LocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoice) GetLocationIdOk() (*string, bool) {
	if o == nil || IsNil(o.LocationId) {
		return nil, false
	}
	return o.LocationId, true
}

// HasLocationId returns a boolean if a field has been set.
func (o *Invoice) HasLocationId() bool {
	if o != nil && !IsNil(o.LocationId) {
		return true
	}

	return false
}

// SetLocationId gets a reference to the given string and assigns it to the LocationId field.
func (o *Invoice) SetLocationId(v string) {
	o.LocationId = &v
}

func (o Invoice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Invoice) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.StatusId) {
		toSerialize["status_id"] = o.StatusId
	}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if !IsNil(o.PoNumber) {
		toSerialize["po_number"] = o.PoNumber
	}
	if !IsNil(o.Terms) {
		toSerialize["terms"] = o.Terms
	}
	if !IsNil(o.PublicNotes) {
		toSerialize["public_notes"] = o.PublicNotes
	}
	if !IsNil(o.PrivateNotes) {
		toSerialize["private_notes"] = o.PrivateNotes
	}
	if !IsNil(o.Footer) {
		toSerialize["footer"] = o.Footer
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
	if !IsNil(o.TotalTaxes) {
		toSerialize["total_taxes"] = o.TotalTaxes
	}
	if !IsNil(o.LineItems) {
		toSerialize["line_items"] = o.LineItems
	}
	if !IsNil(o.Invitations) {
		toSerialize["invitations"] = o.Invitations
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	if !IsNil(o.PaidToDate) {
		toSerialize["paid_to_date"] = o.PaidToDate
	}
	if !IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}
	if !IsNil(o.Partial) {
		toSerialize["partial"] = o.Partial
	}
	if !IsNil(o.IsAmountDiscount) {
		toSerialize["is_amount_discount"] = o.IsAmountDiscount
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["is_deleted"] = o.IsDeleted
	}
	if !IsNil(o.UsesInclusiveTaxes) {
		toSerialize["uses_inclusive_taxes"] = o.UsesInclusiveTaxes
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.LastSentDate) {
		toSerialize["last_sent_date"] = o.LastSentDate
	}
	if !IsNil(o.NextSendDate) {
		toSerialize["next_send_date"] = o.NextSendDate
	}
	if !IsNil(o.PartialDueDate) {
		toSerialize["partial_due_date"] = o.PartialDueDate
	}
	if !IsNil(o.DueDate) {
		toSerialize["due_date"] = o.DueDate
	}
	if !IsNil(o.LastViewed) {
		toSerialize["last_viewed"] = o.LastViewed
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.ArchivedAt) {
		toSerialize["archived_at"] = o.ArchivedAt
	}
	if !IsNil(o.CustomSurcharge1) {
		toSerialize["custom_surcharge1"] = o.CustomSurcharge1
	}
	if !IsNil(o.CustomSurcharge2) {
		toSerialize["custom_surcharge2"] = o.CustomSurcharge2
	}
	if !IsNil(o.CustomSurcharge3) {
		toSerialize["custom_surcharge3"] = o.CustomSurcharge3
	}
	if !IsNil(o.CustomSurcharge4) {
		toSerialize["custom_surcharge4"] = o.CustomSurcharge4
	}
	if !IsNil(o.CustomSurchargeTax1) {
		toSerialize["custom_surcharge_tax1"] = o.CustomSurchargeTax1
	}
	if !IsNil(o.CustomSurchargeTax2) {
		toSerialize["custom_surcharge_tax2"] = o.CustomSurchargeTax2
	}
	if !IsNil(o.CustomSurchargeTax3) {
		toSerialize["custom_surcharge_tax3"] = o.CustomSurchargeTax3
	}
	if !IsNil(o.CustomSurchargeTax4) {
		toSerialize["custom_surcharge_tax4"] = o.CustomSurchargeTax4
	}
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	if !IsNil(o.AutoBillTries) {
		toSerialize["auto_bill_tries"] = o.AutoBillTries
	}
	if !IsNil(o.AutoBillEnabled) {
		toSerialize["auto_bill_enabled"] = o.AutoBillEnabled
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscription_id"] = o.SubscriptionId
	}
	if !IsNil(o.LocationId) {
		toSerialize["location_id"] = o.LocationId
	}
	return toSerialize, nil
}

type NullableInvoice struct {
	value *Invoice
	isSet bool
}

func (v NullableInvoice) Get() *Invoice {
	return v.value
}

func (v *NullableInvoice) Set(val *Invoice) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoice) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoice(val *Invoice) *NullableInvoice {
	return &NullableInvoice{value: val, isSet: true}
}

func (v NullableInvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


