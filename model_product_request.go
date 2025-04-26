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

// checks if the ProductRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductRequest{}

// ProductRequest struct for ProductRequest
type ProductRequest struct {
	// The hashed ID of the user assigned to this product.
	AssignedUserId *string `json:"assigned_user_id,omitempty"`
	// The hashed ID of the project that this product is associated with.
	ProjectId *string `json:"project_id,omitempty"`
	// The hashed ID of the vendor that this product is associated with.
	VendorId *string `json:"vendor_id,omitempty"`
	// Custom value field 1.
	CustomValue1 *string `json:"custom_value1,omitempty"`
	// Custom value field 2.
	CustomValue2 *string `json:"custom_value2,omitempty"`
	// Custom value field 3.
	CustomValue3 *string `json:"custom_value3,omitempty"`
	// Custom value field 4.
	CustomValue4 *string `json:"custom_value4,omitempty"`
	// The product key.
	ProductKey *string `json:"product_key,omitempty"`
	// Notes about the product.
	Notes *string `json:"notes,omitempty"`
	// The cost of the product.
	Cost *float64 `json:"cost,omitempty"`
	// The price of the product.
	Price *float64 `json:"price,omitempty"`
	// The quantity of the product.
	Quantity *float64 `json:"quantity,omitempty"`
	// The default tax name associated with this product
	TaxName1 *string `json:"tax_name1,omitempty"`
	// The default tax rate for this product
	TaxRate1 *float64 `json:"tax_rate1,omitempty"`
	// The default tax name associated with this product
	TaxName2 *string `json:"tax_name2,omitempty"`
	// The default tax rate for this product
	TaxRate2 *float64 `json:"tax_rate2,omitempty"`
	// The default tax name associated with this product
	TaxName3 *string `json:"tax_name3,omitempty"`
	// The default tax rate for this product
	TaxRate3 *float64 `json:"tax_rate3,omitempty"`
	// The quantity of the product that is currently in stock.   **note** this field is not mutable without passing an extra query parameter which will allow modification of this value.  The query parameter ?update_in_stock_quantity=true **MUST** be passed if you wish to update this value manually. 
	InStockQuantity *int32 `json:"in_stock_quantity,omitempty"`
	// Indicates whether stock notifications are enabled for this product
	StockNotification *bool `json:"stock_notification,omitempty"`
	// The minimum quantity threshold for which stock notifications will be triggered
	StockNotificationThreshold *int32 `json:"stock_notification_threshold,omitempty"`
	// The maximum quantity that can be ordered for this product
	MaxQuantity *int32 `json:"max_quantity,omitempty"`
	// The URL of the product image
	ProductImage *string `json:"product_image,omitempty"`
	// The tax category id for this product.'  The following constants are available (default = '1')  ``` PRODUCT_TYPE_PHYSICAL = '1' PRODUCT_TYPE_SERVICE = '2' PRODUCT_TYPE_DIGITAL = '3' PRODUCT_TYPE_SHIPPING = '4' PRODUCT_TYPE_EXEMPT = '5' PRODUCT_TYPE_REDUCED_TAX = '6' PRODUCT_TYPE_OVERRIDE_TAX = '7' PRODUCT_TYPE_ZERO_RATED = '8' PRODUCT_TYPE_REVERSE_TAX = '9' ``` 
	TaxId *string `json:"tax_id,omitempty"`
}

// NewProductRequest instantiates a new ProductRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductRequest() *ProductRequest {
	this := ProductRequest{}
	var quantity float64 = 1
	this.Quantity = &quantity
	var inStockQuantity int32 = 0
	this.InStockQuantity = &inStockQuantity
	var stockNotification bool = true
	this.StockNotification = &stockNotification
	var stockNotificationThreshold int32 = 0
	this.StockNotificationThreshold = &stockNotificationThreshold
	var taxId string = "1"
	this.TaxId = &taxId
	return &this
}

// NewProductRequestWithDefaults instantiates a new ProductRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductRequestWithDefaults() *ProductRequest {
	this := ProductRequest{}
	var quantity float64 = 1
	this.Quantity = &quantity
	var inStockQuantity int32 = 0
	this.InStockQuantity = &inStockQuantity
	var stockNotification bool = true
	this.StockNotification = &stockNotification
	var stockNotificationThreshold int32 = 0
	this.StockNotificationThreshold = &stockNotificationThreshold
	var taxId string = "1"
	this.TaxId = &taxId
	return &this
}

// GetAssignedUserId returns the AssignedUserId field value if set, zero value otherwise.
func (o *ProductRequest) GetAssignedUserId() string {
	if o == nil || IsNil(o.AssignedUserId) {
		var ret string
		return ret
	}
	return *o.AssignedUserId
}

// GetAssignedUserIdOk returns a tuple with the AssignedUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRequest) GetAssignedUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssignedUserId) {
		return nil, false
	}
	return o.AssignedUserId, true
}

// HasAssignedUserId returns a boolean if a field has been set.
func (o *ProductRequest) HasAssignedUserId() bool {
	if o != nil && !IsNil(o.AssignedUserId) {
		return true
	}

	return false
}

// SetAssignedUserId gets a reference to the given string and assigns it to the AssignedUserId field.
func (o *ProductRequest) SetAssignedUserId(v string) {
	o.AssignedUserId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ProductRequest) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRequest) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ProductRequest) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *ProductRequest) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetVendorId returns the VendorId field value if set, zero value otherwise.
func (o *ProductRequest) GetVendorId() string {
	if o == nil || IsNil(o.VendorId) {
		var ret string
		return ret
	}
	return *o.VendorId
}

// GetVendorIdOk returns a tuple with the VendorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRequest) GetVendorIdOk() (*string, bool) {
	if o == nil || IsNil(o.VendorId) {
		return nil, false
	}
	return o.VendorId, true
}

// HasVendorId returns a boolean if a field has been set.
func (o *ProductRequest) HasVendorId() bool {
	if o != nil && !IsNil(o.VendorId) {
		return true
	}

	return false
}

// SetVendorId gets a reference to the given string and assigns it to the VendorId field.
func (o *ProductRequest) SetVendorId(v string) {
	o.VendorId = &v
}

// GetCustomValue1 returns the CustomValue1 field value if set, zero value otherwise.
func (o *ProductRequest) GetCustomValue1() string {
	if o == nil || IsNil(o.CustomValue1) {
		var ret string
		return ret
	}
	return *o.CustomValue1
}

// GetCustomValue1Ok returns a tuple with the CustomValue1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRequest) GetCustomValue1Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue1) {
		return nil, false
	}
	return o.CustomValue1, true
}

// HasCustomValue1 returns a boolean if a field has been set.
func (o *ProductRequest) HasCustomValue1() bool {
	if o != nil && !IsNil(o.CustomValue1) {
		return true
	}

	return false
}

// SetCustomValue1 gets a reference to the given string and assigns it to the CustomValue1 field.
func (o *ProductRequest) SetCustomValue1(v string) {
	o.CustomValue1 = &v
}

// GetCustomValue2 returns the CustomValue2 field value if set, zero value otherwise.
func (o *ProductRequest) GetCustomValue2() string {
	if o == nil || IsNil(o.CustomValue2) {
		var ret string
		return ret
	}
	return *o.CustomValue2
}

// GetCustomValue2Ok returns a tuple with the CustomValue2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRequest) GetCustomValue2Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue2) {
		return nil, false
	}
	return o.CustomValue2, true
}

// HasCustomValue2 returns a boolean if a field has been set.
func (o *ProductRequest) HasCustomValue2() bool {
	if o != nil && !IsNil(o.CustomValue2) {
		return true
	}

	return false
}

// SetCustomValue2 gets a reference to the given string and assigns it to the CustomValue2 field.
func (o *ProductRequest) SetCustomValue2(v string) {
	o.CustomValue2 = &v
}

// GetCustomValue3 returns the CustomValue3 field value if set, zero value otherwise.
func (o *ProductRequest) GetCustomValue3() string {
	if o == nil || IsNil(o.CustomValue3) {
		var ret string
		return ret
	}
	return *o.CustomValue3
}

// GetCustomValue3Ok returns a tuple with the CustomValue3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRequest) GetCustomValue3Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue3) {
		return nil, false
	}
	return o.CustomValue3, true
}

// HasCustomValue3 returns a boolean if a field has been set.
func (o *ProductRequest) HasCustomValue3() bool {
	if o != nil && !IsNil(o.CustomValue3) {
		return true
	}

	return false
}

// SetCustomValue3 gets a reference to the given string and assigns it to the CustomValue3 field.
func (o *ProductRequest) SetCustomValue3(v string) {
	o.CustomValue3 = &v
}

// GetCustomValue4 returns the CustomValue4 field value if set, zero value otherwise.
func (o *ProductRequest) GetCustomValue4() string {
	if o == nil || IsNil(o.CustomValue4) {
		var ret string
		return ret
	}
	return *o.CustomValue4
}

// GetCustomValue4Ok returns a tuple with the CustomValue4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRequest) GetCustomValue4Ok() (*string, bool) {
	if o == nil || IsNil(o.CustomValue4) {
		return nil, false
	}
	return o.CustomValue4, true
}

// HasCustomValue4 returns a boolean if a field has been set.
func (o *ProductRequest) HasCustomValue4() bool {
	if o != nil && !IsNil(o.CustomValue4) {
		return true
	}

	return false
}

// SetCustomValue4 gets a reference to the given string and assigns it to the CustomValue4 field.
func (o *ProductRequest) SetCustomValue4(v string) {
	o.CustomValue4 = &v
}

// GetProductKey returns the ProductKey field value if set, zero value otherwise.
func (o *ProductRequest) GetProductKey() string {
	if o == nil || IsNil(o.ProductKey) {
		var ret string
		return ret
	}
	return *o.ProductKey
}

// GetProductKeyOk returns a tuple with the ProductKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRequest) GetProductKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ProductKey) {
		return nil, false
	}
	return o.ProductKey, true
}

// HasProductKey returns a boolean if a field has been set.
func (o *ProductRequest) HasProductKey() bool {
	if o != nil && !IsNil(o.ProductKey) {
		return true
	}

	return false
}

// SetProductKey gets a reference to the given string and assigns it to the ProductKey field.
func (o *ProductRequest) SetProductKey(v string) {
	o.ProductKey = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *ProductRequest) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRequest) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *ProductRequest) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *ProductRequest) SetNotes(v string) {
	o.Notes = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *ProductRequest) GetCost() float64 {
	if o == nil || IsNil(o.Cost) {
		var ret float64
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRequest) GetCostOk() (*float64, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// HasCost returns a boolean if a field has been set.
func (o *ProductRequest) HasCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given float64 and assigns it to the Cost field.
func (o *ProductRequest) SetCost(v float64) {
	o.Cost = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ProductRequest) GetPrice() float64 {
	if o == nil || IsNil(o.Price) {
		var ret float64
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRequest) GetPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ProductRequest) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float64 and assigns it to the Price field.
func (o *ProductRequest) SetPrice(v float64) {
	o.Price = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *ProductRequest) GetQuantity() float64 {
	if o == nil || IsNil(o.Quantity) {
		var ret float64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRequest) GetQuantityOk() (*float64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *ProductRequest) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float64 and assigns it to the Quantity field.
func (o *ProductRequest) SetQuantity(v float64) {
	o.Quantity = &v
}

// GetTaxName1 returns the TaxName1 field value if set, zero value otherwise.
func (o *ProductRequest) GetTaxName1() string {
	if o == nil || IsNil(o.TaxName1) {
		var ret string
		return ret
	}
	return *o.TaxName1
}

// GetTaxName1Ok returns a tuple with the TaxName1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRequest) GetTaxName1Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName1) {
		return nil, false
	}
	return o.TaxName1, true
}

// HasTaxName1 returns a boolean if a field has been set.
func (o *ProductRequest) HasTaxName1() bool {
	if o != nil && !IsNil(o.TaxName1) {
		return true
	}

	return false
}

// SetTaxName1 gets a reference to the given string and assigns it to the TaxName1 field.
func (o *ProductRequest) SetTaxName1(v string) {
	o.TaxName1 = &v
}

// GetTaxRate1 returns the TaxRate1 field value if set, zero value otherwise.
func (o *ProductRequest) GetTaxRate1() float64 {
	if o == nil || IsNil(o.TaxRate1) {
		var ret float64
		return ret
	}
	return *o.TaxRate1
}

// GetTaxRate1Ok returns a tuple with the TaxRate1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRequest) GetTaxRate1Ok() (*float64, bool) {
	if o == nil || IsNil(o.TaxRate1) {
		return nil, false
	}
	return o.TaxRate1, true
}

// HasTaxRate1 returns a boolean if a field has been set.
func (o *ProductRequest) HasTaxRate1() bool {
	if o != nil && !IsNil(o.TaxRate1) {
		return true
	}

	return false
}

// SetTaxRate1 gets a reference to the given float64 and assigns it to the TaxRate1 field.
func (o *ProductRequest) SetTaxRate1(v float64) {
	o.TaxRate1 = &v
}

// GetTaxName2 returns the TaxName2 field value if set, zero value otherwise.
func (o *ProductRequest) GetTaxName2() string {
	if o == nil || IsNil(o.TaxName2) {
		var ret string
		return ret
	}
	return *o.TaxName2
}

// GetTaxName2Ok returns a tuple with the TaxName2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRequest) GetTaxName2Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName2) {
		return nil, false
	}
	return o.TaxName2, true
}

// HasTaxName2 returns a boolean if a field has been set.
func (o *ProductRequest) HasTaxName2() bool {
	if o != nil && !IsNil(o.TaxName2) {
		return true
	}

	return false
}

// SetTaxName2 gets a reference to the given string and assigns it to the TaxName2 field.
func (o *ProductRequest) SetTaxName2(v string) {
	o.TaxName2 = &v
}

// GetTaxRate2 returns the TaxRate2 field value if set, zero value otherwise.
func (o *ProductRequest) GetTaxRate2() float64 {
	if o == nil || IsNil(o.TaxRate2) {
		var ret float64
		return ret
	}
	return *o.TaxRate2
}

// GetTaxRate2Ok returns a tuple with the TaxRate2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRequest) GetTaxRate2Ok() (*float64, bool) {
	if o == nil || IsNil(o.TaxRate2) {
		return nil, false
	}
	return o.TaxRate2, true
}

// HasTaxRate2 returns a boolean if a field has been set.
func (o *ProductRequest) HasTaxRate2() bool {
	if o != nil && !IsNil(o.TaxRate2) {
		return true
	}

	return false
}

// SetTaxRate2 gets a reference to the given float64 and assigns it to the TaxRate2 field.
func (o *ProductRequest) SetTaxRate2(v float64) {
	o.TaxRate2 = &v
}

// GetTaxName3 returns the TaxName3 field value if set, zero value otherwise.
func (o *ProductRequest) GetTaxName3() string {
	if o == nil || IsNil(o.TaxName3) {
		var ret string
		return ret
	}
	return *o.TaxName3
}

// GetTaxName3Ok returns a tuple with the TaxName3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRequest) GetTaxName3Ok() (*string, bool) {
	if o == nil || IsNil(o.TaxName3) {
		return nil, false
	}
	return o.TaxName3, true
}

// HasTaxName3 returns a boolean if a field has been set.
func (o *ProductRequest) HasTaxName3() bool {
	if o != nil && !IsNil(o.TaxName3) {
		return true
	}

	return false
}

// SetTaxName3 gets a reference to the given string and assigns it to the TaxName3 field.
func (o *ProductRequest) SetTaxName3(v string) {
	o.TaxName3 = &v
}

// GetTaxRate3 returns the TaxRate3 field value if set, zero value otherwise.
func (o *ProductRequest) GetTaxRate3() float64 {
	if o == nil || IsNil(o.TaxRate3) {
		var ret float64
		return ret
	}
	return *o.TaxRate3
}

// GetTaxRate3Ok returns a tuple with the TaxRate3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRequest) GetTaxRate3Ok() (*float64, bool) {
	if o == nil || IsNil(o.TaxRate3) {
		return nil, false
	}
	return o.TaxRate3, true
}

// HasTaxRate3 returns a boolean if a field has been set.
func (o *ProductRequest) HasTaxRate3() bool {
	if o != nil && !IsNil(o.TaxRate3) {
		return true
	}

	return false
}

// SetTaxRate3 gets a reference to the given float64 and assigns it to the TaxRate3 field.
func (o *ProductRequest) SetTaxRate3(v float64) {
	o.TaxRate3 = &v
}

// GetInStockQuantity returns the InStockQuantity field value if set, zero value otherwise.
func (o *ProductRequest) GetInStockQuantity() int32 {
	if o == nil || IsNil(o.InStockQuantity) {
		var ret int32
		return ret
	}
	return *o.InStockQuantity
}

// GetInStockQuantityOk returns a tuple with the InStockQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRequest) GetInStockQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.InStockQuantity) {
		return nil, false
	}
	return o.InStockQuantity, true
}

// HasInStockQuantity returns a boolean if a field has been set.
func (o *ProductRequest) HasInStockQuantity() bool {
	if o != nil && !IsNil(o.InStockQuantity) {
		return true
	}

	return false
}

// SetInStockQuantity gets a reference to the given int32 and assigns it to the InStockQuantity field.
func (o *ProductRequest) SetInStockQuantity(v int32) {
	o.InStockQuantity = &v
}

// GetStockNotification returns the StockNotification field value if set, zero value otherwise.
func (o *ProductRequest) GetStockNotification() bool {
	if o == nil || IsNil(o.StockNotification) {
		var ret bool
		return ret
	}
	return *o.StockNotification
}

// GetStockNotificationOk returns a tuple with the StockNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRequest) GetStockNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.StockNotification) {
		return nil, false
	}
	return o.StockNotification, true
}

// HasStockNotification returns a boolean if a field has been set.
func (o *ProductRequest) HasStockNotification() bool {
	if o != nil && !IsNil(o.StockNotification) {
		return true
	}

	return false
}

// SetStockNotification gets a reference to the given bool and assigns it to the StockNotification field.
func (o *ProductRequest) SetStockNotification(v bool) {
	o.StockNotification = &v
}

// GetStockNotificationThreshold returns the StockNotificationThreshold field value if set, zero value otherwise.
func (o *ProductRequest) GetStockNotificationThreshold() int32 {
	if o == nil || IsNil(o.StockNotificationThreshold) {
		var ret int32
		return ret
	}
	return *o.StockNotificationThreshold
}

// GetStockNotificationThresholdOk returns a tuple with the StockNotificationThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRequest) GetStockNotificationThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.StockNotificationThreshold) {
		return nil, false
	}
	return o.StockNotificationThreshold, true
}

// HasStockNotificationThreshold returns a boolean if a field has been set.
func (o *ProductRequest) HasStockNotificationThreshold() bool {
	if o != nil && !IsNil(o.StockNotificationThreshold) {
		return true
	}

	return false
}

// SetStockNotificationThreshold gets a reference to the given int32 and assigns it to the StockNotificationThreshold field.
func (o *ProductRequest) SetStockNotificationThreshold(v int32) {
	o.StockNotificationThreshold = &v
}

// GetMaxQuantity returns the MaxQuantity field value if set, zero value otherwise.
func (o *ProductRequest) GetMaxQuantity() int32 {
	if o == nil || IsNil(o.MaxQuantity) {
		var ret int32
		return ret
	}
	return *o.MaxQuantity
}

// GetMaxQuantityOk returns a tuple with the MaxQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRequest) GetMaxQuantityOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxQuantity) {
		return nil, false
	}
	return o.MaxQuantity, true
}

// HasMaxQuantity returns a boolean if a field has been set.
func (o *ProductRequest) HasMaxQuantity() bool {
	if o != nil && !IsNil(o.MaxQuantity) {
		return true
	}

	return false
}

// SetMaxQuantity gets a reference to the given int32 and assigns it to the MaxQuantity field.
func (o *ProductRequest) SetMaxQuantity(v int32) {
	o.MaxQuantity = &v
}

// GetProductImage returns the ProductImage field value if set, zero value otherwise.
func (o *ProductRequest) GetProductImage() string {
	if o == nil || IsNil(o.ProductImage) {
		var ret string
		return ret
	}
	return *o.ProductImage
}

// GetProductImageOk returns a tuple with the ProductImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRequest) GetProductImageOk() (*string, bool) {
	if o == nil || IsNil(o.ProductImage) {
		return nil, false
	}
	return o.ProductImage, true
}

// HasProductImage returns a boolean if a field has been set.
func (o *ProductRequest) HasProductImage() bool {
	if o != nil && !IsNil(o.ProductImage) {
		return true
	}

	return false
}

// SetProductImage gets a reference to the given string and assigns it to the ProductImage field.
func (o *ProductRequest) SetProductImage(v string) {
	o.ProductImage = &v
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *ProductRequest) GetTaxId() string {
	if o == nil || IsNil(o.TaxId) {
		var ret string
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductRequest) GetTaxIdOk() (*string, bool) {
	if o == nil || IsNil(o.TaxId) {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *ProductRequest) HasTaxId() bool {
	if o != nil && !IsNil(o.TaxId) {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given string and assigns it to the TaxId field.
func (o *ProductRequest) SetTaxId(v string) {
	o.TaxId = &v
}

func (o ProductRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssignedUserId) {
		toSerialize["assigned_user_id"] = o.AssignedUserId
	}
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	if !IsNil(o.VendorId) {
		toSerialize["vendor_id"] = o.VendorId
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
	if !IsNil(o.ProductKey) {
		toSerialize["product_key"] = o.ProductKey
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
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
	if !IsNil(o.InStockQuantity) {
		toSerialize["in_stock_quantity"] = o.InStockQuantity
	}
	if !IsNil(o.StockNotification) {
		toSerialize["stock_notification"] = o.StockNotification
	}
	if !IsNil(o.StockNotificationThreshold) {
		toSerialize["stock_notification_threshold"] = o.StockNotificationThreshold
	}
	if !IsNil(o.MaxQuantity) {
		toSerialize["max_quantity"] = o.MaxQuantity
	}
	if !IsNil(o.ProductImage) {
		toSerialize["product_image"] = o.ProductImage
	}
	if !IsNil(o.TaxId) {
		toSerialize["tax_id"] = o.TaxId
	}
	return toSerialize, nil
}

type NullableProductRequest struct {
	value *ProductRequest
	isSet bool
}

func (v NullableProductRequest) Get() *ProductRequest {
	return v.value
}

func (v *NullableProductRequest) Set(val *ProductRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductRequest(val *ProductRequest) *NullableProductRequest {
	return &NullableProductRequest{value: val, isSet: true}
}

func (v NullableProductRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


