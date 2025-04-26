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
	"time"
)

// checks if the InvoiceItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoiceItem{}

// InvoiceItem struct for InvoiceItem
type InvoiceItem struct {
	// The quantity of the product offered for this line item
	Quantity *float32 `json:"quantity,omitempty"`
	// The cost of the product offered for this line item
	Cost *float32 `json:"cost,omitempty"`
	// The product key of the product offered for this line item (Referred to as Product in the product tab)
	ProductKey *string `json:"product_key,omitempty"`
	// The cost of the product offered for this line item (Referred to as Cost in the product tab)
	ProductCost *float32 `json:"product_cost,omitempty"`
	// The notes/description for the product offered for this line item
	Notes *string `json:"notes,omitempty"`
	// The discount applied to the product offered for this line item
	Discount *float32 `json:"discount,omitempty"`
	// Indicates whether the discount applied to the product offered for this line item is a fixed amount or a percentage
	IsAmountDiscount *bool `json:"is_amount_discount,omitempty"`
	// The name of the first tax applied to the product offered for this line item
	TaxName1 *string `json:"tax_name1,omitempty"`
	// The rate of the first tax applied to the product offered for this line item
	TaxRate1 *float32 `json:"tax_rate1,omitempty"`
	// The name of the second tax applied to the product offered for this line item
	TaxName2 *string `json:"tax_name2,omitempty"`
	// The rate of the second tax applied to the product offered for this line item
	TaxRate2 *float32 `json:"tax_rate2,omitempty"`
	// The name of the third tax applied to the product offered for this line item
	TaxName3 *string `json:"tax_name3,omitempty"`
	// The rate of the third tax applied to the product offered for this line item
	TaxRate3 *float32 `json:"tax_rate3,omitempty"`
	// Deprecated
	// Deprecated
	SortId *string `json:"sort_id,omitempty"`
	// The total amount of the product offered for this line item
	LineTotal *float32 `json:"line_total,omitempty"`
	// The total amount of the product offered for this line item before discounts
	GrossLineTotal *float32 `json:"gross_line_total,omitempty"`
	// The total amount of tax applied to the product offered for this line item
	TaxAmount *float32 `json:"tax_amount,omitempty"`
	// Deprecated
	// Deprecated
	Date *time.Time `json:"date,omitempty"`
	// The first custom value of the product offered for this line item
	CustomValue1 *string `json:"custom_value1,omitempty"`
	// The second custom value of the product offered for this line item
	CustomValue2 *string `json:"custom_value2,omitempty"`
	// The third custom value of the product offered for this line item
	CustomValue3 *string `json:"custom_value3,omitempty"`
	// The fourth custom value of the product offered for this line item
	CustomValue4 *string `json:"custom_value4,omitempty"`
	// 1 = product, 2 = service, 3 unpaid gateway fee, 4 paid gateway fee, 5 late fee, 6 expense
	TypeId *string `json:"type_id,omitempty"`
	// The tax ID of the product: 1 product, 2 service, 3 digital, 4 shipping, 5 exempt, 5 reduced tax, 7 override, 8 zero rate, 9 reverse tax
	TaxId *string `json:"tax_id,omitempty"`
}

// NewInvoiceItem instantiates a new InvoiceItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoiceItem() *InvoiceItem {
	this := InvoiceItem{}
	var typeId string = "1"
	this.TypeId = &typeId
	var taxId string = "1"
	this.TaxId = &taxId
	return &this
}

// NewInvoiceItemWithDefaults instantiates a new InvoiceItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoiceItemWithDefaults() *InvoiceItem {
	this := InvoiceItem{}
	var typeId string = "1"
	this.TypeId = &typeId
	var taxId string = "1"
	this.TaxId = &taxId
	return &this
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *InvoiceItem) GetQuantity() float32 {
	if o == nil || IsNil(o.Quantity) {
		var ret float32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceItem) GetQuantityOk() (*float32, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *InvoiceItem) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float32 and assigns it to the Quantity field.
func (o *InvoiceItem) SetQuantity(v float32) {
	o.Quantity = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *InvoiceItem) GetCost() float32 {
	if o == nil || IsNil(o.Cost) {
		var ret float32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceItem) GetCostOk() (*float32, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *InvoiceItem) HasCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given float32 and assigns it to the Cost field.
func (o *InvoiceItem) SetCost(v float32) {
	o.Cost = &v
}

// GetProductKey returns the ProductKey field value if set, zero value otherwise.
func (o *InvoiceItem) GetProductKey() string {
	if o == nil || IsNil(o.ProductKey) {
		var ret string
		return ret
	}
	return *o.ProductKey
}

// GetProductKeyOk returns a tuple with the ProductKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceItem) GetProductKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ProductKey) {
		return nil, false
	}
	return o.ProductKey, true
}

// HasProductKey returns a boolean if a field has been set.
func (o *InvoiceItem) HasProductKey() bool {
	if o != nil && !IsNil(o.ProductKey) {
		return true
	}

	return false
}

// SetProductKey gets a reference to the given string and assigns it to the ProductKey field.
func (o *InvoiceItem) SetProductKey(v string) {
	o.ProductKey = &v
}

// GetProductCost returns the ProductCost field value if set, zero value otherwise.
func (o *InvoiceItem) GetProductCost() float32 {
	if o == nil || IsNil(o.ProductCost) {
		var ret float32
		return ret
	}
	return *o.ProductCost
}

// GetProductCostOk returns a tuple with the ProductCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceItem) GetProductCostOk() (*float32, bool) {
	if o == nil || IsNil(o.ProductCost) {
		return nil, false
	}
	return o.ProductCost, true
}

// HasProductCost returns a boolean if a field has been set.
func (o *InvoiceItem) HasProductCost() bool {
	if o != nil && !IsNil(o.ProductCost) {
		return true
	}

	return false
}

// SetProductCost gets a reference to the given float32 and assigns it to the ProductCost field.
func (o *InvoiceItem) SetProductCost(v float32) {
	o.ProductCost = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *InvoiceItem) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceItem) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *InvoiceItem) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *InvoiceItem) SetNotes(v string) {
	o.Notes = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *InvoiceItem) GetDiscount() float32 {
	if o == nil || IsNil(o.Discount) {
		var ret float32
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceItem) GetDiscountOk() (*float32, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *InvoiceItem) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given float32 and assigns it to the Discount field.
func (o *InvoiceItem) SetDiscount(v float32) {
	o.Discount = &v
}

// GetIsAmountDiscount returns the IsAmountDiscount field value if set, zero value otherwise.
func (o *InvoiceItem) GetIsAmountDiscount() bool {
	if o == nil || IsNil(o.IsAmountDiscount) {
		var ret bool
		return ret
	}
	return *o.IsAmountDiscount
}

// GetIsAmountDiscountOk returns a tuple with the IsAmountDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceItem) GetIsAmountDiscountOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAmountDiscount) {
		return nil, false
	}
	return o.IsAmountDiscount, true
}

// HasIsAmountDiscount returns a boolean if a field has been set.
func (o *InvoiceItem) HasIsAmountDiscount() bool {
	if o != nil && !IsNil(o.IsAmountDiscount) {
		return true
	}

	return false
}

// SetIsAmountDiscount gets a reference to the given bool and assigns it to the IsAmountDiscount field.
func (o *InvoiceItem) SetIsAmountDiscount(v bool) {
	o.IsAmountDiscount = &v
}

// GetTaxName1 returns the TaxName1 field value if set, zero value otherwise.
func (o *InvoiceItem) GetTaxName1() string {
	if o == nil || IsNil(o.TaxName1) {
		var ret string
		return ret
	}
	return *o.TaxName1
}

// GetTaxName1Ok returns a tuple with the TaxName1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceItem) GetTaxName1Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName1) {
		return nil, false
	}
	return o.TaxName1, true
}

// HasTaxName1 returns a boolean if a field has been set.
func (o *InvoiceItem) HasTaxName1() bool {
	if o != nil && !IsNil(o.TaxName1) {
		return true
	}

	return false
}

// SetTaxName1 gets a reference to the given string and assigns it to the TaxName1 field.
func (o *InvoiceItem) SetTaxName1(v string) {
	o.TaxName1 = &v
}

// GetTaxRate1 returns the TaxRate1 field value if set, zero value otherwise.
func (o *InvoiceItem) GetTaxRate1() float32 {
	if o == nil || IsNil(o.TaxRate1) {
		var ret float32
		return ret
	}
	return *o.TaxRate1
}

// GetTaxRate1Ok returns a tuple with the TaxRate1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceItem) GetTaxRate1Ok() (*float32, bool) {
	if o == nil || IsNil(o.TaxRate1) {
		return nil, false
	}
	return o.TaxRate1, true
}

// HasTaxRate1 returns a boolean if a field has been set.
func (o *InvoiceItem) HasTaxRate1() bool {
	if o != nil && !IsNil(o.TaxRate1) {
		return true
	}

	return false
}

// SetTaxRate1 gets a reference to the given float32 and assigns it to the TaxRate1 field.
func (o *InvoiceItem) SetTaxRate1(v float32) {
	o.TaxRate1 = &v
}

// GetTaxName2 returns the TaxName2 field value if set, zero value otherwise.
func (o *InvoiceItem) GetTaxName2() string {
	if o == nil || IsNil(o.TaxName2) {
		var ret string
		return ret
	}
	return *o.TaxName2
}

// GetTaxName2Ok returns a tuple with the TaxName2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceItem) GetTaxName2Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName2) {
		return nil, false
	}
	return o.TaxName2, true
}

// HasTaxName2 returns a boolean if a field has been set.
func (o *InvoiceItem) HasTaxName2() bool {
	if o != nil && !IsNil(o.TaxName2) {
		return true
	}

	return false
}

// SetTaxName2 gets a reference to the given string and assigns it to the TaxName2 field.
func (o *InvoiceItem) SetTaxName2(v string) {
	o.TaxName2 = &v
}

// GetTaxRate2 returns the TaxRate2 field value if set, zero value otherwise.
func (o *InvoiceItem) GetTaxRate2() float32 {
	if o == nil || IsNil(o.TaxRate2) {
		var ret float32
		return ret
	}
	return *o.TaxRate2
}

// GetTaxRate2Ok returns a tuple with the TaxRate2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceItem) GetTaxRate2Ok() (*float32, bool) {
	if o == nil || IsNil(o.TaxRate2) {
		return nil, false
	}
	return o.TaxRate2, true
}

// HasTaxRate2 returns a boolean if a field has been set.
func (o *InvoiceItem) HasTaxRate2() bool {
	if o != nil && !IsNil(o.TaxRate2) {
		return true
	}

	return false
}

// SetTaxRate2 gets a reference to the given float32 and assigns it to the TaxRate2 field.
func (o *InvoiceItem) SetTaxRate2(v float32) {
	o.TaxRate2 = &v
}

// GetTaxName3 returns the TaxName3 field value if set, zero value otherwise.
func (o *InvoiceItem) GetTaxName3() string {
	if o == nil || IsNil(o.TaxName3) {
		var ret string
		return ret
	}
	return *o.TaxName3
}

// GetTaxName3Ok returns a tuple with the TaxName3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceItem) GetTaxName3Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName3) {
		return nil, false
	}
	return o.TaxName3, true
}

// HasTaxName3 returns a boolean if a field has been set.
func (o *InvoiceItem) HasTaxName3() bool {
	if o != nil && !IsNil(o.TaxName3) {
		return true
	}

	return false
}

// SetTaxName3 gets a reference to the given string and assigns it to the TaxName3 field.
func (o *InvoiceItem) SetTaxName3(v string) {
	o.TaxName3 = &v
}

// GetTaxRate3 returns the TaxRate3 field value if set, zero value otherwise.
func (o *InvoiceItem) GetTaxRate3() float32 {
	if o == nil || IsNil(o.TaxRate3) {
		var ret float32
		return ret
	}
	return *o.TaxRate3
}

// GetTaxRate3Ok returns a tuple with the TaxRate3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceItem) GetTaxRate3Ok() (*float32, bool) {
	if o == nil || IsNil(o.TaxRate3) {
		return nil, false
	}
	return o.TaxRate3, true
}

// HasTaxRate3 returns a boolean if a field has been set.
func (o *InvoiceItem) HasTaxRate3() bool {
	if o != nil && !IsNil(o.TaxRate3) {
		return true
	}

	return false
}

// SetTaxRate3 gets a reference to the given float32 and assigns it to the TaxRate3 field.
func (o *InvoiceItem) SetTaxRate3(v float32) {
	o.TaxRate3 = &v
}

// GetSortId returns the SortId field value if set, zero value otherwise.
// Deprecated
func (o *InvoiceItem) GetSortId() string {
	if o == nil || IsNil(o.SortId) {
		var ret string
		return ret
	}
	return *o.SortId
}

// GetSortIdOk returns a tuple with the SortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *InvoiceItem) GetSortIdOk() (*string, bool) {
	if o == nil || IsNil(o.SortId) {
		return nil, false
	}
	return o.SortId, true
}

// HasSortId returns a boolean if a field has been set.
func (o *InvoiceItem) HasSortId() bool {
	if o != nil && !IsNil(o.SortId) {
		return true
	}

	return false
}

// SetSortId gets a reference to the given string and assigns it to the SortId field.
// Deprecated
func (o *InvoiceItem) SetSortId(v string) {
	o.SortId = &v
}

// GetLineTotal returns the LineTotal field value if set, zero value otherwise.
func (o *InvoiceItem) GetLineTotal() float32 {
	if o == nil || IsNil(o.LineTotal) {
		var ret float32
		return ret
	}
	return *o.LineTotal
}

// GetLineTotalOk returns a tuple with the LineTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceItem) GetLineTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.LineTotal) {
		return nil, false
	}
	return o.LineTotal, true
}

// HasLineTotal returns a boolean if a field has been set.
func (o *InvoiceItem) HasLineTotal() bool {
	if o != nil && !IsNil(o.LineTotal) {
		return true
	}

	return false
}

// SetLineTotal gets a reference to the given float32 and assigns it to the LineTotal field.
func (o *InvoiceItem) SetLineTotal(v float32) {
	o.LineTotal = &v
}

// GetGrossLineTotal returns the GrossLineTotal field value if set, zero value otherwise.
func (o *InvoiceItem) GetGrossLineTotal() float32 {
	if o == nil || IsNil(o.GrossLineTotal) {
		var ret float32
		return ret
	}
	return *o.GrossLineTotal
}

// GetGrossLineTotalOk returns a tuple with the GrossLineTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceItem) GetGrossLineTotalOk() (*float32, bool) {
	if o == nil || IsNil(o.GrossLineTotal) {
		return nil, false
	}
	return o.GrossLineTotal, true
}

// HasGrossLineTotal returns a boolean if a field has been set.
func (o *InvoiceItem) HasGrossLineTotal() bool {
	if o != nil && !IsNil(o.GrossLineTotal) {
		return true
	}

	return false
}

// SetGrossLineTotal gets a reference to the given float32 and assigns it to the GrossLineTotal field.
func (o *InvoiceItem) SetGrossLineTotal(v float32) {
	o.GrossLineTotal = &v
}

// GetTaxAmount returns the TaxAmount field value if set, zero value otherwise.
func (o *InvoiceItem) GetTaxAmount() float32 {
	if o == nil || IsNil(o.TaxAmount) {
		var ret float32
		return ret
	}
	return *o.TaxAmount
}

// GetTaxAmountOk returns a tuple with the TaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceItem) GetTaxAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.TaxAmount) {
		return nil, false
	}
	return o.TaxAmount, true
}

// HasTaxAmount returns a boolean if a field has been set.
func (o *InvoiceItem) HasTaxAmount() bool {
	if o != nil && !IsNil(o.TaxAmount) {
		return true
	}

	return false
}

// SetTaxAmount gets a reference to the given float32 and assigns it to the TaxAmount field.
func (o *InvoiceItem) SetTaxAmount(v float32) {
	o.TaxAmount = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
// Deprecated
func (o *InvoiceItem) GetDate() time.Time {
	if o == nil || IsNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *InvoiceItem) GetDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *InvoiceItem) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
// Deprecated
func (o *InvoiceItem) SetDate(v time.Time) {
	o.Date = &v
}

// GetCustomValue1 returns the CustomValue1 field value if set, zero value otherwise.
func (o *InvoiceItem) GetCustomValue1() string {
	if o == nil || IsNil(o.CustomValue1) {
		var ret string
		return ret
	}
	return *o.CustomValue1
}

// GetCustomValue1Ok returns a tuple with the CustomValue1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceItem) GetCustomValue1Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue1) {
		return nil, false
	}
	return o.CustomValue1, true
}

// HasCustomValue1 returns a boolean if a field has been set.
func (o *InvoiceItem) HasCustomValue1() bool {
	if o != nil && !IsNil(o.CustomValue1) {
		return true
	}

	return false
}

// SetCustomValue1 gets a reference to the given string and assigns it to the CustomValue1 field.
func (o *InvoiceItem) SetCustomValue1(v string) {
	o.CustomValue1 = &v
}

// GetCustomValue2 returns the CustomValue2 field value if set, zero value otherwise.
func (o *InvoiceItem) GetCustomValue2() string {
	if o == nil || IsNil(o.CustomValue2) {
		var ret string
		return ret
	}
	return *o.CustomValue2
}

// GetCustomValue2Ok returns a tuple with the CustomValue2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceItem) GetCustomValue2Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue2) {
		return nil, false
	}
	return o.CustomValue2, true
}

// HasCustomValue2 returns a boolean if a field has been set.
func (o *InvoiceItem) HasCustomValue2() bool {
	if o != nil && !IsNil(o.CustomValue2) {
		return true
	}

	return false
}

// SetCustomValue2 gets a reference to the given string and assigns it to the CustomValue2 field.
func (o *InvoiceItem) SetCustomValue2(v string) {
	o.CustomValue2 = &v
}

// GetCustomValue3 returns the CustomValue3 field value if set, zero value otherwise.
func (o *InvoiceItem) GetCustomValue3() string {
	if o == nil || IsNil(o.CustomValue3) {
		var ret string
		return ret
	}
	return *o.CustomValue3
}

// GetCustomValue3Ok returns a tuple with the CustomValue3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceItem) GetCustomValue3Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue3) {
		return nil, false
	}
	return o.CustomValue3, true
}

// HasCustomValue3 returns a boolean if a field has been set.
func (o *InvoiceItem) HasCustomValue3() bool {
	if o != nil && !IsNil(o.CustomValue3) {
		return true
	}

	return false
}

// SetCustomValue3 gets a reference to the given string and assigns it to the CustomValue3 field.
func (o *InvoiceItem) SetCustomValue3(v string) {
	o.CustomValue3 = &v
}

// GetCustomValue4 returns the CustomValue4 field value if set, zero value otherwise.
func (o *InvoiceItem) GetCustomValue4() string {
	if o == nil || IsNil(o.CustomValue4) {
		var ret string
		return ret
	}
	return *o.CustomValue4
}

// GetCustomValue4Ok returns a tuple with the CustomValue4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceItem) GetCustomValue4Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue4) {
		return nil, false
	}
	return o.CustomValue4, true
}

// HasCustomValue4 returns a boolean if a field has been set.
func (o *InvoiceItem) HasCustomValue4() bool {
	if o != nil && !IsNil(o.CustomValue4) {
		return true
	}

	return false
}

// SetCustomValue4 gets a reference to the given string and assigns it to the CustomValue4 field.
func (o *InvoiceItem) SetCustomValue4(v string) {
	o.CustomValue4 = &v
}

// GetTypeId returns the TypeId field value if set, zero value otherwise.
func (o *InvoiceItem) GetTypeId() string {
	if o == nil || IsNil(o.TypeId) {
		var ret string
		return ret
	}
	return *o.TypeId
}

// GetTypeIdOk returns a tuple with the TypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceItem) GetTypeIdOk() (*string, bool) {
	if o == nil || IsNil(o.TypeId) {
		return nil, false
	}
	return o.TypeId, true
}

// HasTypeId returns a boolean if a field has been set.
func (o *InvoiceItem) HasTypeId() bool {
	if o != nil && !IsNil(o.TypeId) {
		return true
	}

	return false
}

// SetTypeId gets a reference to the given string and assigns it to the TypeId field.
func (o *InvoiceItem) SetTypeId(v string) {
	o.TypeId = &v
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *InvoiceItem) GetTaxId() string {
	if o == nil || IsNil(o.TaxId) {
		var ret string
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoiceItem) GetTaxIdOk() (*string, bool) {
	if o == nil || IsNil(o.TaxId) {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *InvoiceItem) HasTaxId() bool {
	if o != nil && !IsNil(o.TaxId) {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given string and assigns it to the TaxId field.
func (o *InvoiceItem) SetTaxId(v string) {
	o.TaxId = &v
}

func (o InvoiceItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoiceItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	if !IsNil(o.ProductKey) {
		toSerialize["product_key"] = o.ProductKey
	}
	if !IsNil(o.ProductCost) {
		toSerialize["product_cost"] = o.ProductCost
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}
	if !IsNil(o.IsAmountDiscount) {
		toSerialize["is_amount_discount"] = o.IsAmountDiscount
	}
	if !IsNil(o.TaxName1) {
		toSerialize["tax_name1"] = o.TaxName1
	}
	if !IsNil(o.TaxRate1) {
		toSerialize["tax_rate1"] = o.TaxRate1
	}
	if !IsNil(o.TaxName2) {
		toSerialize["tax_name2"] = o.TaxName2
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
	if !IsNil(o.SortId) {
		toSerialize["sort_id"] = o.SortId
	}
	if !IsNil(o.LineTotal) {
		toSerialize["line_total"] = o.LineTotal
	}
	if !IsNil(o.GrossLineTotal) {
		toSerialize["gross_line_total"] = o.GrossLineTotal
	}
	if !IsNil(o.TaxAmount) {
		toSerialize["tax_amount"] = o.TaxAmount
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
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
	if !IsNil(o.TypeId) {
		toSerialize["type_id"] = o.TypeId
	}
	if !IsNil(o.TaxId) {
		toSerialize["tax_id"] = o.TaxId
	}
	return toSerialize, nil
}

type NullableInvoiceItem struct {
	value *InvoiceItem
	isSet bool
}

func (v NullableInvoiceItem) Get() *InvoiceItem {
	return v.value
}

func (v *NullableInvoiceItem) Set(val *InvoiceItem) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoiceItem) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoiceItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoiceItem(val *InvoiceItem) *NullableInvoiceItem {
	return &NullableInvoiceItem{value: val, isSet: true}
}

func (v NullableInvoiceItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoiceItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


