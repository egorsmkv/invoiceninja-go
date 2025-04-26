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

// checks if the Credit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Credit{}

// Credit struct for Credit
type Credit struct {
	// The unique hashed ID of the credit
	Id *string `json:"id,omitempty"`
	// The unique hashed ID of the user associated with the credit
	UserId *string `json:"user_id,omitempty"`
	// The unique hashed ID of the assigned user responsible for the credit
	AssignedUserId *string `json:"assigned_user_id,omitempty"`
	// The unique hashed ID of the client associated with the credit
	ClientId *string `json:"client_id,omitempty"`
	// The ID representing the current status of the credit
	StatusId *string `json:"status_id,omitempty"`
	// The unique hashed ID of the linked invoice to which the credit is applied
	InvoiceId *string `json:"invoice_id,omitempty"`
	// The unique alphanumeric credit number per company
	Number *string `json:"number,omitempty"`
	// The purchase order number referred to by the credit
	PoNumber *string `json:"po_number,omitempty"`
	// The terms associated with the credit
	Terms *string `json:"terms,omitempty"`
	// Public notes for the credit
	PublicNotes *string `json:"public_notes,omitempty"`
	// Private notes for internal use, not visible to the client
	PrivateNotes *string `json:"private_notes,omitempty"`
	// The footer text for the credit
	Footer *string `json:"footer,omitempty"`
	// Custom value 1 for additional credit information
	CustomValue1 *string `json:"custom_value1,omitempty"`
	// Custom value 2 for additional credit information
	CustomValue2 *string `json:"custom_value2,omitempty"`
	// Custom value 3 for additional credit information
	CustomValue3 *string `json:"custom_value3,omitempty"`
	// Custom value 4 for additional credit information
	CustomValue4 *string `json:"custom_value4,omitempty"`
	// The name of the first tax applied to the credit
	TaxName1 *string `json:"tax_name1,omitempty"`
	// The name of the second tax applied to the credit
	TaxName2 *string `json:"tax_name2,omitempty"`
	// The rate of the first tax applied to the credit
	TaxRate1 *float32 `json:"tax_rate1,omitempty"`
	// The rate of the second tax applied to the credit
	TaxRate2 *float32 `json:"tax_rate2,omitempty"`
	// The name of the third tax applied to the credit
	TaxName3 *string `json:"tax_name3,omitempty"`
	// The rate of the third tax applied to the credit
	TaxRate3 *float32 `json:"tax_rate3,omitempty"`
	// The total amount of taxes for the credit
	TotalTaxes *float32 `json:"total_taxes,omitempty"`
	// An array of objects which define the line items of the credit
	LineItems []InvoiceItem `json:"line_items,omitempty"`
	// The total amount of the credit
	Amount *float32 `json:"amount,omitempty"`
	// The outstanding balance of the credit
	Balance *float32 `json:"balance,omitempty"`
	// The total amount paid to date for the credit
	PaidToDate *float32 `json:"paid_to_date,omitempty"`
	// The discount applied to the credit
	Discount *float32 `json:"discount,omitempty"`
	// The partial amount applied to the credit
	Partial *float32 `json:"partial,omitempty"`
	// Indicates whether the discount applied is a fixed amount or a percentage
	IsAmountDiscount *bool `json:"is_amount_discount,omitempty"`
	// Indicates whether the credit has been deleted
	IsDeleted *bool `json:"is_deleted,omitempty"`
	// Indicates whether the tax rates applied to the credit are inclusive or exclusive
	UsesInclusiveTaxes *bool `json:"uses_inclusive_taxes,omitempty"`
	// The date the credit was issued
	Date *string `json:"date,omitempty"`
	// The date the credit was last sent out
	LastSentDate *string `json:"last_sent_date,omitempty"`
	// The next scheduled date for sending a credit reminder
	NextSendDate *string `json:"next_send_date,omitempty"`
	// The due date for the partial amount of the credit
	PartialDueDate *string `json:"partial_due_date,omitempty"`
	// The due date for the total amount of the credit
	DueDate *string `json:"due_date,omitempty"`
	Settings *CompanySettings `json:"settings,omitempty"`
	// The timestamp of the last time the credit was viewed
	LastViewed *float32 `json:"last_viewed,omitempty"`
	// The timestamp of the last time the credit was updated
	UpdatedAt *float32 `json:"updated_at,omitempty"`
	// The timestamp of the last time the credit was archived
	ArchivedAt *float32 `json:"archived_at,omitempty"`
	// First custom surcharge amount
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
	// The client location id that this invoice relates to
	LocationId *string `json:"location_id,omitempty"`
}

// NewCredit instantiates a new Credit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredit() *Credit {
	this := Credit{}
	return &this
}

// NewCreditWithDefaults instantiates a new Credit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditWithDefaults() *Credit {
	this := Credit{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Credit) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Credit) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Credit) SetId(v string) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *Credit) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *Credit) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *Credit) SetUserId(v string) {
	o.UserId = &v
}

// GetAssignedUserId returns the AssignedUserId field value if set, zero value otherwise.
func (o *Credit) GetAssignedUserId() string {
	if o == nil || IsNil(o.AssignedUserId) {
		var ret string
		return ret
	}
	return *o.AssignedUserId
}

// GetAssignedUserIdOk returns a tuple with the AssignedUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetAssignedUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssignedUserId) {
		return nil, false
	}
	return o.AssignedUserId, true
}

// HasAssignedUserId returns a boolean if a field has been set.
func (o *Credit) HasAssignedUserId() bool {
	if o != nil && !IsNil(o.AssignedUserId) {
		return true
	}

	return false
}

// SetAssignedUserId gets a reference to the given string and assigns it to the AssignedUserId field.
func (o *Credit) SetAssignedUserId(v string) {
	o.AssignedUserId = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *Credit) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *Credit) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *Credit) SetClientId(v string) {
	o.ClientId = &v
}

// GetStatusId returns the StatusId field value if set, zero value otherwise.
func (o *Credit) GetStatusId() string {
	if o == nil || IsNil(o.StatusId) {
		var ret string
		return ret
	}
	return *o.StatusId
}

// GetStatusIdOk returns a tuple with the StatusId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetStatusIdOk() (*string, bool) {
	if o == nil || IsNil(o.StatusId) {
		return nil, false
	}
	return o.StatusId, true
}

// HasStatusId returns a boolean if a field has been set.
func (o *Credit) HasStatusId() bool {
	if o != nil && !IsNil(o.StatusId) {
		return true
	}

	return false
}

// SetStatusId gets a reference to the given string and assigns it to the StatusId field.
func (o *Credit) SetStatusId(v string) {
	o.StatusId = &v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise.
func (o *Credit) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceId) {
		return nil, false
	}
	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *Credit) HasInvoiceId() bool {
	if o != nil && !IsNil(o.InvoiceId) {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *Credit) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *Credit) GetNumber() string {
	if o == nil || IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetNumberOk() (*string, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *Credit) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *Credit) SetNumber(v string) {
	o.Number = &v
}

// GetPoNumber returns the PoNumber field value if set, zero value otherwise.
func (o *Credit) GetPoNumber() string {
	if o == nil || IsNil(o.PoNumber) {
		var ret string
		return ret
	}
	return *o.PoNumber
}

// GetPoNumberOk returns a tuple with the PoNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetPoNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PoNumber) {
		return nil, false
	}
	return o.PoNumber, true
}

// HasPoNumber returns a boolean if a field has been set.
func (o *Credit) HasPoNumber() bool {
	if o != nil && !IsNil(o.PoNumber) {
		return true
	}

	return false
}

// SetPoNumber gets a reference to the given string and assigns it to the PoNumber field.
func (o *Credit) SetPoNumber(v string) {
	o.PoNumber = &v
}

// GetTerms returns the Terms field value if set, zero value otherwise.
func (o *Credit) GetTerms() string {
	if o == nil || IsNil(o.Terms) {
		var ret string
		return ret
	}
	return *o.Terms
}

// GetTermsOk returns a tuple with the Terms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetTermsOk() (*string, bool) {
	if o == nil || IsNil(o.Terms) {
		return nil, false
	}
	return o.Terms, true
}

// HasTerms returns a boolean if a field has been set.
func (o *Credit) HasTerms() bool {
	if o != nil && !IsNil(o.Terms) {
		return true
	}

	return false
}

// SetTerms gets a reference to the given string and assigns it to the Terms field.
func (o *Credit) SetTerms(v string) {
	o.Terms = &v
}

// GetPublicNotes returns the PublicNotes field value if set, zero value otherwise.
func (o *Credit) GetPublicNotes() string {
	if o == nil || IsNil(o.PublicNotes) {
		var ret string
		return ret
	}
	return *o.PublicNotes
}

// GetPublicNotesOk returns a tuple with the PublicNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetPublicNotesOk() (*string, bool) {
	if o == nil || IsNil(o.PublicNotes) {
		return nil, false
	}
	return o.PublicNotes, true
}

// HasPublicNotes returns a boolean if a field has been set.
func (o *Credit) HasPublicNotes() bool {
	if o != nil && !IsNil(o.PublicNotes) {
		return true
	}

	return false
}

// SetPublicNotes gets a reference to the given string and assigns it to the PublicNotes field.
func (o *Credit) SetPublicNotes(v string) {
	o.PublicNotes = &v
}

// GetPrivateNotes returns the PrivateNotes field value if set, zero value otherwise.
func (o *Credit) GetPrivateNotes() string {
	if o == nil || IsNil(o.PrivateNotes) {
		var ret string
		return ret
	}
	return *o.PrivateNotes
}

// GetPrivateNotesOk returns a tuple with the PrivateNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetPrivateNotesOk() (*string, bool) {
	if o == nil || IsNil(o.PrivateNotes) {
		return nil, false
	}
	return o.PrivateNotes, true
}

// HasPrivateNotes returns a boolean if a field has been set.
func (o *Credit) HasPrivateNotes() bool {
	if o != nil && !IsNil(o.PrivateNotes) {
		return true
	}

	return false
}

// SetPrivateNotes gets a reference to the given string and assigns it to the PrivateNotes field.
func (o *Credit) SetPrivateNotes(v string) {
	o.PrivateNotes = &v
}

// GetFooter returns the Footer field value if set, zero value otherwise.
func (o *Credit) GetFooter() string {
	if o == nil || IsNil(o.Footer) {
		var ret string
		return ret
	}
	return *o.Footer
}

// GetFooterOk returns a tuple with the Footer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetFooterOk() (*string, bool) {
	if o == nil || IsNil(o.Footer) {
		return nil, false
	}
	return o.Footer, true
}

// HasFooter returns a boolean if a field has been set.
func (o *Credit) HasFooter() bool {
	if o != nil && !IsNil(o.Footer) {
		return true
	}

	return false
}

// SetFooter gets a reference to the given string and assigns it to the Footer field.
func (o *Credit) SetFooter(v string) {
	o.Footer = &v
}

// GetCustomValue1 returns the CustomValue1 field value if set, zero value otherwise.
func (o *Credit) GetCustomValue1() string {
	if o == nil || IsNil(o.CustomValue1) {
		var ret string
		return ret
	}
	return *o.CustomValue1
}

// GetCustomValue1Ok returns a tuple with the CustomValue1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetCustomValue1Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue1) {
		return nil, false
	}
	return o.CustomValue1, true
}

// HasCustomValue1 returns a boolean if a field has been set.
func (o *Credit) HasCustomValue1() bool {
	if o != nil && !IsNil(o.CustomValue1) {
		return true
	}

	return false
}

// SetCustomValue1 gets a reference to the given string and assigns it to the CustomValue1 field.
func (o *Credit) SetCustomValue1(v string) {
	o.CustomValue1 = &v
}

// GetCustomValue2 returns the CustomValue2 field value if set, zero value otherwise.
func (o *Credit) GetCustomValue2() string {
	if o == nil || IsNil(o.CustomValue2) {
		var ret string
		return ret
	}
	return *o.CustomValue2
}

// GetCustomValue2Ok returns a tuple with the CustomValue2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetCustomValue2Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue2) {
		return nil, false
	}
	return o.CustomValue2, true
}

// HasCustomValue2 returns a boolean if a field has been set.
func (o *Credit) HasCustomValue2() bool {
	if o != nil && !IsNil(o.CustomValue2) {
		return true
	}

	return false
}

// SetCustomValue2 gets a reference to the given string and assigns it to the CustomValue2 field.
func (o *Credit) SetCustomValue2(v string) {
	o.CustomValue2 = &v
}

// GetCustomValue3 returns the CustomValue3 field value if set, zero value otherwise.
func (o *Credit) GetCustomValue3() string {
	if o == nil || IsNil(o.CustomValue3) {
		var ret string
		return ret
	}
	return *o.CustomValue3
}

// GetCustomValue3Ok returns a tuple with the CustomValue3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetCustomValue3Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue3) {
		return nil, false
	}
	return o.CustomValue3, true
}

// HasCustomValue3 returns a boolean if a field has been set.
func (o *Credit) HasCustomValue3() bool {
	if o != nil && !IsNil(o.CustomValue3) {
		return true
	}

	return false
}

// SetCustomValue3 gets a reference to the given string and assigns it to the CustomValue3 field.
func (o *Credit) SetCustomValue3(v string) {
	o.CustomValue3 = &v
}

// GetCustomValue4 returns the CustomValue4 field value if set, zero value otherwise.
func (o *Credit) GetCustomValue4() string {
	if o == nil || IsNil(o.CustomValue4) {
		var ret string
		return ret
	}
	return *o.CustomValue4
}

// GetCustomValue4Ok returns a tuple with the CustomValue4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetCustomValue4Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue4) {
		return nil, false
	}
	return o.CustomValue4, true
}

// HasCustomValue4 returns a boolean if a field has been set.
func (o *Credit) HasCustomValue4() bool {
	if o != nil && !IsNil(o.CustomValue4) {
		return true
	}

	return false
}

// SetCustomValue4 gets a reference to the given string and assigns it to the CustomValue4 field.
func (o *Credit) SetCustomValue4(v string) {
	o.CustomValue4 = &v
}

// GetTaxName1 returns the TaxName1 field value if set, zero value otherwise.
func (o *Credit) GetTaxName1() string {
	if o == nil || IsNil(o.TaxName1) {
		var ret string
		return ret
	}
	return *o.TaxName1
}

// GetTaxName1Ok returns a tuple with the TaxName1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetTaxName1Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName1) {
		return nil, false
	}
	return o.TaxName1, true
}

// HasTaxName1 returns a boolean if a field has been set.
func (o *Credit) HasTaxName1() bool {
	if o != nil && !IsNil(o.TaxName1) {
		return true
	}

	return false
}

// SetTaxName1 gets a reference to the given string and assigns it to the TaxName1 field.
func (o *Credit) SetTaxName1(v string) {
	o.TaxName1 = &v
}

// GetTaxName2 returns the TaxName2 field value if set, zero value otherwise.
func (o *Credit) GetTaxName2() string {
	if o == nil || IsNil(o.TaxName2) {
		var ret string
		return ret
	}
	return *o.TaxName2
}

// GetTaxName2Ok returns a tuple with the TaxName2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetTaxName2Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName2) {
		return nil, false
	}
	return o.TaxName2, true
}

// HasTaxName2 returns a boolean if a field has been set.
func (o *Credit) HasTaxName2() bool {
	if o != nil && !IsNil(o.TaxName2) {
		return true
	}

	return false
}

// SetTaxName2 gets a reference to the given string and assigns it to the TaxName2 field.
func (o *Credit) SetTaxName2(v string) {
	o.TaxName2 = &v
}

// GetTaxRate1 returns the TaxRate1 field value if set, zero value otherwise.
func (o *Credit) GetTaxRate1() float32 {
	if o == nil || IsNil(o.TaxRate1) {
		var ret float32
		return ret
	}
	return *o.TaxRate1
}

// GetTaxRate1Ok returns a tuple with the TaxRate1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetTaxRate1Ok() (*float32, bool) {
	if o == nil || IsNil(o.TaxRate1) {
		return nil, false
	}
	return o.TaxRate1, true
}

// HasTaxRate1 returns a boolean if a field has been set.
func (o *Credit) HasTaxRate1() bool {
	if o != nil && !IsNil(o.TaxRate1) {
		return true
	}

	return false
}

// SetTaxRate1 gets a reference to the given float32 and assigns it to the TaxRate1 field.
func (o *Credit) SetTaxRate1(v float32) {
	o.TaxRate1 = &v
}

// GetTaxRate2 returns the TaxRate2 field value if set, zero value otherwise.
func (o *Credit) GetTaxRate2() float32 {
	if o == nil || IsNil(o.TaxRate2) {
		var ret float32
		return ret
	}
	return *o.TaxRate2
}

// GetTaxRate2Ok returns a tuple with the TaxRate2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetTaxRate2Ok() (*float32, bool) {
	if o == nil || IsNil(o.TaxRate2) {
		return nil, false
	}
	return o.TaxRate2, true
}

// HasTaxRate2 returns a boolean if a field has been set.
func (o *Credit) HasTaxRate2() bool {
	if o != nil && !IsNil(o.TaxRate2) {
		return true
	}

	return false
}

// SetTaxRate2 gets a reference to the given float32 and assigns it to the TaxRate2 field.
func (o *Credit) SetTaxRate2(v float32) {
	o.TaxRate2 = &v
}

// GetTaxName3 returns the TaxName3 field value if set, zero value otherwise.
func (o *Credit) GetTaxName3() string {
	if o == nil || IsNil(o.TaxName3) {
		var ret string
		return ret
	}
	return *o.TaxName3
}

// GetTaxName3Ok returns a tuple with the TaxName3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetTaxName3Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName3) {
		return nil, false
	}
	return o.TaxName3, true
}

// HasTaxName3 returns a boolean if a field has been set.
func (o *Credit) HasTaxName3() bool {
	if o != nil && !IsNil(o.TaxName3) {
		return true
	}

	return false
}

// SetTaxName3 gets a reference to the given string and assigns it to the TaxName3 field.
func (o *Credit) SetTaxName3(v string) {
	o.TaxName3 = &v
}

// GetTaxRate3 returns the TaxRate3 field value if set, zero value otherwise.
func (o *Credit) GetTaxRate3() float32 {
	if o == nil || IsNil(o.TaxRate3) {
		var ret float32
		return ret
	}
	return *o.TaxRate3
}

// GetTaxRate3Ok returns a tuple with the TaxRate3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetTaxRate3Ok() (*float32, bool) {
	if o == nil || IsNil(o.TaxRate3) {
		return nil, false
	}
	return o.TaxRate3, true
}

// HasTaxRate3 returns a boolean if a field has been set.
func (o *Credit) HasTaxRate3() bool {
	if o != nil && !IsNil(o.TaxRate3) {
		return true
	}

	return false
}

// SetTaxRate3 gets a reference to the given float32 and assigns it to the TaxRate3 field.
func (o *Credit) SetTaxRate3(v float32) {
	o.TaxRate3 = &v
}

// GetTotalTaxes returns the TotalTaxes field value if set, zero value otherwise.
func (o *Credit) GetTotalTaxes() float32 {
	if o == nil || IsNil(o.TotalTaxes) {
		var ret float32
		return ret
	}
	return *o.TotalTaxes
}

// GetTotalTaxesOk returns a tuple with the TotalTaxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetTotalTaxesOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalTaxes) {
		return nil, false
	}
	return o.TotalTaxes, true
}

// HasTotalTaxes returns a boolean if a field has been set.
func (o *Credit) HasTotalTaxes() bool {
	if o != nil && !IsNil(o.TotalTaxes) {
		return true
	}

	return false
}

// SetTotalTaxes gets a reference to the given float32 and assigns it to the TotalTaxes field.
func (o *Credit) SetTotalTaxes(v float32) {
	o.TotalTaxes = &v
}

// GetLineItems returns the LineItems field value if set, zero value otherwise.
func (o *Credit) GetLineItems() []InvoiceItem {
	if o == nil || IsNil(o.LineItems) {
		var ret []InvoiceItem
		return ret
	}
	return o.LineItems
}

// GetLineItemsOk returns a tuple with the LineItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetLineItemsOk() ([]InvoiceItem, bool) {
	if o == nil || IsNil(o.LineItems) {
		return nil, false
	}
	return o.LineItems, true
}

// HasLineItems returns a boolean if a field has been set.
func (o *Credit) HasLineItems() bool {
	if o != nil && !IsNil(o.LineItems) {
		return true
	}

	return false
}

// SetLineItems gets a reference to the given []InvoiceItem and assigns it to the LineItems field.
func (o *Credit) SetLineItems(v []InvoiceItem) {
	o.LineItems = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Credit) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Credit) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *Credit) SetAmount(v float32) {
	o.Amount = &v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *Credit) GetBalance() float32 {
	if o == nil || IsNil(o.Balance) {
		var ret float32
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetBalanceOk() (*float32, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *Credit) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given float32 and assigns it to the Balance field.
func (o *Credit) SetBalance(v float32) {
	o.Balance = &v
}

// GetPaidToDate returns the PaidToDate field value if set, zero value otherwise.
func (o *Credit) GetPaidToDate() float32 {
	if o == nil || IsNil(o.PaidToDate) {
		var ret float32
		return ret
	}
	return *o.PaidToDate
}

// GetPaidToDateOk returns a tuple with the PaidToDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetPaidToDateOk() (*float32, bool) {
	if o == nil || IsNil(o.PaidToDate) {
		return nil, false
	}
	return o.PaidToDate, true
}

// HasPaidToDate returns a boolean if a field has been set.
func (o *Credit) HasPaidToDate() bool {
	if o != nil && !IsNil(o.PaidToDate) {
		return true
	}

	return false
}

// SetPaidToDate gets a reference to the given float32 and assigns it to the PaidToDate field.
func (o *Credit) SetPaidToDate(v float32) {
	o.PaidToDate = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *Credit) GetDiscount() float32 {
	if o == nil || IsNil(o.Discount) {
		var ret float32
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetDiscountOk() (*float32, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *Credit) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given float32 and assigns it to the Discount field.
func (o *Credit) SetDiscount(v float32) {
	o.Discount = &v
}

// GetPartial returns the Partial field value if set, zero value otherwise.
func (o *Credit) GetPartial() float32 {
	if o == nil || IsNil(o.Partial) {
		var ret float32
		return ret
	}
	return *o.Partial
}

// GetPartialOk returns a tuple with the Partial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetPartialOk() (*float32, bool) {
	if o == nil || IsNil(o.Partial) {
		return nil, false
	}
	return o.Partial, true
}

// HasPartial returns a boolean if a field has been set.
func (o *Credit) HasPartial() bool {
	if o != nil && !IsNil(o.Partial) {
		return true
	}

	return false
}

// SetPartial gets a reference to the given float32 and assigns it to the Partial field.
func (o *Credit) SetPartial(v float32) {
	o.Partial = &v
}

// GetIsAmountDiscount returns the IsAmountDiscount field value if set, zero value otherwise.
func (o *Credit) GetIsAmountDiscount() bool {
	if o == nil || IsNil(o.IsAmountDiscount) {
		var ret bool
		return ret
	}
	return *o.IsAmountDiscount
}

// GetIsAmountDiscountOk returns a tuple with the IsAmountDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetIsAmountDiscountOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAmountDiscount) {
		return nil, false
	}
	return o.IsAmountDiscount, true
}

// HasIsAmountDiscount returns a boolean if a field has been set.
func (o *Credit) HasIsAmountDiscount() bool {
	if o != nil && !IsNil(o.IsAmountDiscount) {
		return true
	}

	return false
}

// SetIsAmountDiscount gets a reference to the given bool and assigns it to the IsAmountDiscount field.
func (o *Credit) SetIsAmountDiscount(v bool) {
	o.IsAmountDiscount = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *Credit) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *Credit) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *Credit) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

// GetUsesInclusiveTaxes returns the UsesInclusiveTaxes field value if set, zero value otherwise.
func (o *Credit) GetUsesInclusiveTaxes() bool {
	if o == nil || IsNil(o.UsesInclusiveTaxes) {
		var ret bool
		return ret
	}
	return *o.UsesInclusiveTaxes
}

// GetUsesInclusiveTaxesOk returns a tuple with the UsesInclusiveTaxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetUsesInclusiveTaxesOk() (*bool, bool) {
	if o == nil || IsNil(o.UsesInclusiveTaxes) {
		return nil, false
	}
	return o.UsesInclusiveTaxes, true
}

// HasUsesInclusiveTaxes returns a boolean if a field has been set.
func (o *Credit) HasUsesInclusiveTaxes() bool {
	if o != nil && !IsNil(o.UsesInclusiveTaxes) {
		return true
	}

	return false
}

// SetUsesInclusiveTaxes gets a reference to the given bool and assigns it to the UsesInclusiveTaxes field.
func (o *Credit) SetUsesInclusiveTaxes(v bool) {
	o.UsesInclusiveTaxes = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *Credit) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *Credit) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *Credit) SetDate(v string) {
	o.Date = &v
}

// GetLastSentDate returns the LastSentDate field value if set, zero value otherwise.
func (o *Credit) GetLastSentDate() string {
	if o == nil || IsNil(o.LastSentDate) {
		var ret string
		return ret
	}
	return *o.LastSentDate
}

// GetLastSentDateOk returns a tuple with the LastSentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetLastSentDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastSentDate) {
		return nil, false
	}
	return o.LastSentDate, true
}

// HasLastSentDate returns a boolean if a field has been set.
func (o *Credit) HasLastSentDate() bool {
	if o != nil && !IsNil(o.LastSentDate) {
		return true
	}

	return false
}

// SetLastSentDate gets a reference to the given string and assigns it to the LastSentDate field.
func (o *Credit) SetLastSentDate(v string) {
	o.LastSentDate = &v
}

// GetNextSendDate returns the NextSendDate field value if set, zero value otherwise.
func (o *Credit) GetNextSendDate() string {
	if o == nil || IsNil(o.NextSendDate) {
		var ret string
		return ret
	}
	return *o.NextSendDate
}

// GetNextSendDateOk returns a tuple with the NextSendDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetNextSendDateOk() (*string, bool) {
	if o == nil || IsNil(o.NextSendDate) {
		return nil, false
	}
	return o.NextSendDate, true
}

// HasNextSendDate returns a boolean if a field has been set.
func (o *Credit) HasNextSendDate() bool {
	if o != nil && !IsNil(o.NextSendDate) {
		return true
	}

	return false
}

// SetNextSendDate gets a reference to the given string and assigns it to the NextSendDate field.
func (o *Credit) SetNextSendDate(v string) {
	o.NextSendDate = &v
}

// GetPartialDueDate returns the PartialDueDate field value if set, zero value otherwise.
func (o *Credit) GetPartialDueDate() string {
	if o == nil || IsNil(o.PartialDueDate) {
		var ret string
		return ret
	}
	return *o.PartialDueDate
}

// GetPartialDueDateOk returns a tuple with the PartialDueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetPartialDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.PartialDueDate) {
		return nil, false
	}
	return o.PartialDueDate, true
}

// HasPartialDueDate returns a boolean if a field has been set.
func (o *Credit) HasPartialDueDate() bool {
	if o != nil && !IsNil(o.PartialDueDate) {
		return true
	}

	return false
}

// SetPartialDueDate gets a reference to the given string and assigns it to the PartialDueDate field.
func (o *Credit) SetPartialDueDate(v string) {
	o.PartialDueDate = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *Credit) GetDueDate() string {
	if o == nil || IsNil(o.DueDate) {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *Credit) HasDueDate() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *Credit) SetDueDate(v string) {
	o.DueDate = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *Credit) GetSettings() CompanySettings {
	if o == nil || IsNil(o.Settings) {
		var ret CompanySettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetSettingsOk() (*CompanySettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *Credit) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given CompanySettings and assigns it to the Settings field.
func (o *Credit) SetSettings(v CompanySettings) {
	o.Settings = &v
}

// GetLastViewed returns the LastViewed field value if set, zero value otherwise.
func (o *Credit) GetLastViewed() float32 {
	if o == nil || IsNil(o.LastViewed) {
		var ret float32
		return ret
	}
	return *o.LastViewed
}

// GetLastViewedOk returns a tuple with the LastViewed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetLastViewedOk() (*float32, bool) {
	if o == nil || IsNil(o.LastViewed) {
		return nil, false
	}
	return o.LastViewed, true
}

// HasLastViewed returns a boolean if a field has been set.
func (o *Credit) HasLastViewed() bool {
	if o != nil && !IsNil(o.LastViewed) {
		return true
	}

	return false
}

// SetLastViewed gets a reference to the given float32 and assigns it to the LastViewed field.
func (o *Credit) SetLastViewed(v float32) {
	o.LastViewed = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Credit) GetUpdatedAt() float32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret float32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetUpdatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Credit) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given float32 and assigns it to the UpdatedAt field.
func (o *Credit) SetUpdatedAt(v float32) {
	o.UpdatedAt = &v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *Credit) GetArchivedAt() float32 {
	if o == nil || IsNil(o.ArchivedAt) {
		var ret float32
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetArchivedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.ArchivedAt) {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *Credit) HasArchivedAt() bool {
	if o != nil && !IsNil(o.ArchivedAt) {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given float32 and assigns it to the ArchivedAt field.
func (o *Credit) SetArchivedAt(v float32) {
	o.ArchivedAt = &v
}

// GetCustomSurcharge1 returns the CustomSurcharge1 field value if set, zero value otherwise.
func (o *Credit) GetCustomSurcharge1() float32 {
	if o == nil || IsNil(o.CustomSurcharge1) {
		var ret float32
		return ret
	}
	return *o.CustomSurcharge1
}

// GetCustomSurcharge1Ok returns a tuple with the CustomSurcharge1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetCustomSurcharge1Ok() (*float32, bool) {
	if o == nil || IsNil(o.CustomSurcharge1) {
		return nil, false
	}
	return o.CustomSurcharge1, true
}

// HasCustomSurcharge1 returns a boolean if a field has been set.
func (o *Credit) HasCustomSurcharge1() bool {
	if o != nil && !IsNil(o.CustomSurcharge1) {
		return true
	}

	return false
}

// SetCustomSurcharge1 gets a reference to the given float32 and assigns it to the CustomSurcharge1 field.
func (o *Credit) SetCustomSurcharge1(v float32) {
	o.CustomSurcharge1 = &v
}

// GetCustomSurcharge2 returns the CustomSurcharge2 field value if set, zero value otherwise.
func (o *Credit) GetCustomSurcharge2() float32 {
	if o == nil || IsNil(o.CustomSurcharge2) {
		var ret float32
		return ret
	}
	return *o.CustomSurcharge2
}

// GetCustomSurcharge2Ok returns a tuple with the CustomSurcharge2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetCustomSurcharge2Ok() (*float32, bool) {
	if o == nil || IsNil(o.CustomSurcharge2) {
		return nil, false
	}
	return o.CustomSurcharge2, true
}

// HasCustomSurcharge2 returns a boolean if a field has been set.
func (o *Credit) HasCustomSurcharge2() bool {
	if o != nil && !IsNil(o.CustomSurcharge2) {
		return true
	}

	return false
}

// SetCustomSurcharge2 gets a reference to the given float32 and assigns it to the CustomSurcharge2 field.
func (o *Credit) SetCustomSurcharge2(v float32) {
	o.CustomSurcharge2 = &v
}

// GetCustomSurcharge3 returns the CustomSurcharge3 field value if set, zero value otherwise.
func (o *Credit) GetCustomSurcharge3() float32 {
	if o == nil || IsNil(o.CustomSurcharge3) {
		var ret float32
		return ret
	}
	return *o.CustomSurcharge3
}

// GetCustomSurcharge3Ok returns a tuple with the CustomSurcharge3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetCustomSurcharge3Ok() (*float32, bool) {
	if o == nil || IsNil(o.CustomSurcharge3) {
		return nil, false
	}
	return o.CustomSurcharge3, true
}

// HasCustomSurcharge3 returns a boolean if a field has been set.
func (o *Credit) HasCustomSurcharge3() bool {
	if o != nil && !IsNil(o.CustomSurcharge3) {
		return true
	}

	return false
}

// SetCustomSurcharge3 gets a reference to the given float32 and assigns it to the CustomSurcharge3 field.
func (o *Credit) SetCustomSurcharge3(v float32) {
	o.CustomSurcharge3 = &v
}

// GetCustomSurcharge4 returns the CustomSurcharge4 field value if set, zero value otherwise.
func (o *Credit) GetCustomSurcharge4() float32 {
	if o == nil || IsNil(o.CustomSurcharge4) {
		var ret float32
		return ret
	}
	return *o.CustomSurcharge4
}

// GetCustomSurcharge4Ok returns a tuple with the CustomSurcharge4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetCustomSurcharge4Ok() (*float32, bool) {
	if o == nil || IsNil(o.CustomSurcharge4) {
		return nil, false
	}
	return o.CustomSurcharge4, true
}

// HasCustomSurcharge4 returns a boolean if a field has been set.
func (o *Credit) HasCustomSurcharge4() bool {
	if o != nil && !IsNil(o.CustomSurcharge4) {
		return true
	}

	return false
}

// SetCustomSurcharge4 gets a reference to the given float32 and assigns it to the CustomSurcharge4 field.
func (o *Credit) SetCustomSurcharge4(v float32) {
	o.CustomSurcharge4 = &v
}

// GetCustomSurchargeTax1 returns the CustomSurchargeTax1 field value if set, zero value otherwise.
func (o *Credit) GetCustomSurchargeTax1() bool {
	if o == nil || IsNil(o.CustomSurchargeTax1) {
		var ret bool
		return ret
	}
	return *o.CustomSurchargeTax1
}

// GetCustomSurchargeTax1Ok returns a tuple with the CustomSurchargeTax1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetCustomSurchargeTax1Ok() (*bool, bool) {
	if o == nil || IsNil(o.CustomSurchargeTax1) {
		return nil, false
	}
	return o.CustomSurchargeTax1, true
}

// HasCustomSurchargeTax1 returns a boolean if a field has been set.
func (o *Credit) HasCustomSurchargeTax1() bool {
	if o != nil && !IsNil(o.CustomSurchargeTax1) {
		return true
	}

	return false
}

// SetCustomSurchargeTax1 gets a reference to the given bool and assigns it to the CustomSurchargeTax1 field.
func (o *Credit) SetCustomSurchargeTax1(v bool) {
	o.CustomSurchargeTax1 = &v
}

// GetCustomSurchargeTax2 returns the CustomSurchargeTax2 field value if set, zero value otherwise.
func (o *Credit) GetCustomSurchargeTax2() bool {
	if o == nil || IsNil(o.CustomSurchargeTax2) {
		var ret bool
		return ret
	}
	return *o.CustomSurchargeTax2
}

// GetCustomSurchargeTax2Ok returns a tuple with the CustomSurchargeTax2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetCustomSurchargeTax2Ok() (*bool, bool) {
	if o == nil || IsNil(o.CustomSurchargeTax2) {
		return nil, false
	}
	return o.CustomSurchargeTax2, true
}

// HasCustomSurchargeTax2 returns a boolean if a field has been set.
func (o *Credit) HasCustomSurchargeTax2() bool {
	if o != nil && !IsNil(o.CustomSurchargeTax2) {
		return true
	}

	return false
}

// SetCustomSurchargeTax2 gets a reference to the given bool and assigns it to the CustomSurchargeTax2 field.
func (o *Credit) SetCustomSurchargeTax2(v bool) {
	o.CustomSurchargeTax2 = &v
}

// GetCustomSurchargeTax3 returns the CustomSurchargeTax3 field value if set, zero value otherwise.
func (o *Credit) GetCustomSurchargeTax3() bool {
	if o == nil || IsNil(o.CustomSurchargeTax3) {
		var ret bool
		return ret
	}
	return *o.CustomSurchargeTax3
}

// GetCustomSurchargeTax3Ok returns a tuple with the CustomSurchargeTax3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetCustomSurchargeTax3Ok() (*bool, bool) {
	if o == nil || IsNil(o.CustomSurchargeTax3) {
		return nil, false
	}
	return o.CustomSurchargeTax3, true
}

// HasCustomSurchargeTax3 returns a boolean if a field has been set.
func (o *Credit) HasCustomSurchargeTax3() bool {
	if o != nil && !IsNil(o.CustomSurchargeTax3) {
		return true
	}

	return false
}

// SetCustomSurchargeTax3 gets a reference to the given bool and assigns it to the CustomSurchargeTax3 field.
func (o *Credit) SetCustomSurchargeTax3(v bool) {
	o.CustomSurchargeTax3 = &v
}

// GetCustomSurchargeTax4 returns the CustomSurchargeTax4 field value if set, zero value otherwise.
func (o *Credit) GetCustomSurchargeTax4() bool {
	if o == nil || IsNil(o.CustomSurchargeTax4) {
		var ret bool
		return ret
	}
	return *o.CustomSurchargeTax4
}

// GetCustomSurchargeTax4Ok returns a tuple with the CustomSurchargeTax4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetCustomSurchargeTax4Ok() (*bool, bool) {
	if o == nil || IsNil(o.CustomSurchargeTax4) {
		return nil, false
	}
	return o.CustomSurchargeTax4, true
}

// HasCustomSurchargeTax4 returns a boolean if a field has been set.
func (o *Credit) HasCustomSurchargeTax4() bool {
	if o != nil && !IsNil(o.CustomSurchargeTax4) {
		return true
	}

	return false
}

// SetCustomSurchargeTax4 gets a reference to the given bool and assigns it to the CustomSurchargeTax4 field.
func (o *Credit) SetCustomSurchargeTax4(v bool) {
	o.CustomSurchargeTax4 = &v
}

// GetLocationId returns the LocationId field value if set, zero value otherwise.
func (o *Credit) GetLocationId() string {
	if o == nil || IsNil(o.LocationId) {
		var ret string
		return ret
	}
	return *o.LocationId
}

// GetLocationIdOk returns a tuple with the LocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Credit) GetLocationIdOk() (*string, bool) {
	if o == nil || IsNil(o.LocationId) {
		return nil, false
	}
	return o.LocationId, true
}

// HasLocationId returns a boolean if a field has been set.
func (o *Credit) HasLocationId() bool {
	if o != nil && !IsNil(o.LocationId) {
		return true
	}

	return false
}

// SetLocationId gets a reference to the given string and assigns it to the LocationId field.
func (o *Credit) SetLocationId(v string) {
	o.LocationId = &v
}

func (o Credit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Credit) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.InvoiceId) {
		toSerialize["invoice_id"] = o.InvoiceId
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
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
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
	if !IsNil(o.LocationId) {
		toSerialize["location_id"] = o.LocationId
	}
	return toSerialize, nil
}

type NullableCredit struct {
	value *Credit
	isSet bool
}

func (v NullableCredit) Get() *Credit {
	return v.value
}

func (v *NullableCredit) Set(val *Credit) {
	v.value = val
	v.isSet = true
}

func (v NullableCredit) IsSet() bool {
	return v.isSet
}

func (v *NullableCredit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredit(val *Credit) *NullableCredit {
	return &NullableCredit{value: val, isSet: true}
}

func (v NullableCredit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


