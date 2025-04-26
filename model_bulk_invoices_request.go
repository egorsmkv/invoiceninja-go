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

// checks if the BulkInvoicesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkInvoicesRequest{}

// BulkInvoicesRequest struct for BulkInvoicesRequest
type BulkInvoicesRequest struct {
	// The email type to be sent, when bulk emailing invoices
	EmailType *Enum `json:"email_type,omitempty"`
	// The action to be performed, options include:   - `bulk_download`     Bulk download an array of invoice PDFs (These are sent to the admin via email.)   - `download`     Download a single PDF. (Returns a single PDF object)   - `bulk_print`     Merges an array of Invoice PDFs for easy one click printing.   - `auto_bill`     Attempts to automatically bill the invoices with the payment method on file.   - `clone_to_invoice`     Returns a clone of the invoice.   - `clone_to_quote`     Returns a quote cloned using the properties of the given invoice.   - `mark_paid`     Marks an array of invoices as paid.   - `mark_sent`     Marks an array of invoices as sent.   - `restore`     Restores an array of invoices   - `delete`     Deletes an array of invoices   - `archive`     Archives an array of invoices   - `cancel`     Cancels an array of invoices   - `email`     Emails an array of invoices   - `send_email`     Emails an array of invoices. Requires additional properties to be sent. `email_type`   
	Action *string `json:"action,omitempty"`
	Ids []string `json:"ids,omitempty"`
}

// NewBulkInvoicesRequest instantiates a new BulkInvoicesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkInvoicesRequest() *BulkInvoicesRequest {
	this := BulkInvoicesRequest{}
	return &this
}

// NewBulkInvoicesRequestWithDefaults instantiates a new BulkInvoicesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkInvoicesRequestWithDefaults() *BulkInvoicesRequest {
	this := BulkInvoicesRequest{}
	return &this
}

// GetEmailType returns the EmailType field value if set, zero value otherwise.
func (o *BulkInvoicesRequest) GetEmailType() Enum {
	if o == nil || IsNil(o.EmailType) {
		var ret Enum
		return ret
	}
	return *o.EmailType
}

// GetEmailTypeOk returns a tuple with the EmailType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkInvoicesRequest) GetEmailTypeOk() (*Enum, bool) {
	if o == nil || IsNil(o.EmailType) {
		return nil, false
	}
	return o.EmailType, true
}

// HasEmailType returns a boolean if a field has been set.
func (o *BulkInvoicesRequest) HasEmailType() bool {
	if o != nil && !IsNil(o.EmailType) {
		return true
	}

	return false
}

// SetEmailType gets a reference to the given Enum and assigns it to the EmailType field.
func (o *BulkInvoicesRequest) SetEmailType(v Enum) {
	o.EmailType = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *BulkInvoicesRequest) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkInvoicesRequest) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *BulkInvoicesRequest) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *BulkInvoicesRequest) SetAction(v string) {
	o.Action = &v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *BulkInvoicesRequest) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkInvoicesRequest) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *BulkInvoicesRequest) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *BulkInvoicesRequest) SetIds(v []string) {
	o.Ids = v
}

func (o BulkInvoicesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkInvoicesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmailType) {
		toSerialize["email_type"] = o.EmailType
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	return toSerialize, nil
}

type NullableBulkInvoicesRequest struct {
	value *BulkInvoicesRequest
	isSet bool
}

func (v NullableBulkInvoicesRequest) Get() *BulkInvoicesRequest {
	return v.value
}

func (v *NullableBulkInvoicesRequest) Set(val *BulkInvoicesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkInvoicesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkInvoicesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkInvoicesRequest(val *BulkInvoicesRequest) *NullableBulkInvoicesRequest {
	return &NullableBulkInvoicesRequest{value: val, isSet: true}
}

func (v NullableBulkInvoicesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkInvoicesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


