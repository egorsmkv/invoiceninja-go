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

// checks if the ClientStatementRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientStatementRequest{}

// ClientStatementRequest struct for ClientStatementRequest
type ClientStatementRequest struct {
	// The start date of the statement period - format Y-m-d
	StartDate *string `json:"start_date,omitempty"`
	// The start date of the statement period - format Y-m-d
	EndDate *string `json:"end_date,omitempty"`
	// The hashed ID of the client
	ClientId *string `json:"client_id,omitempty"`
	// Flag which determines if the payments table is shown
	ShowPaymentsTable *bool `json:"show_payments_table,omitempty"`
	// Flag which determines if the credits table is shown
	ShowCreditsTable *bool `json:"show_credits_table,omitempty"`
	// Flag which determines if the aging table is shown
	ShowAgingTable *bool `json:"show_aging_table,omitempty"`
}

// NewClientStatementRequest instantiates a new ClientStatementRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientStatementRequest() *ClientStatementRequest {
	this := ClientStatementRequest{}
	return &this
}

// NewClientStatementRequestWithDefaults instantiates a new ClientStatementRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientStatementRequestWithDefaults() *ClientStatementRequest {
	this := ClientStatementRequest{}
	return &this
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ClientStatementRequest) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientStatementRequest) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *ClientStatementRequest) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *ClientStatementRequest) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ClientStatementRequest) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientStatementRequest) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ClientStatementRequest) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *ClientStatementRequest) SetEndDate(v string) {
	o.EndDate = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ClientStatementRequest) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientStatementRequest) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ClientStatementRequest) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ClientStatementRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetShowPaymentsTable returns the ShowPaymentsTable field value if set, zero value otherwise.
func (o *ClientStatementRequest) GetShowPaymentsTable() bool {
	if o == nil || IsNil(o.ShowPaymentsTable) {
		var ret bool
		return ret
	}
	return *o.ShowPaymentsTable
}

// GetShowPaymentsTableOk returns a tuple with the ShowPaymentsTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientStatementRequest) GetShowPaymentsTableOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowPaymentsTable) {
		return nil, false
	}
	return o.ShowPaymentsTable, true
}

// HasShowPaymentsTable returns a boolean if a field has been set.
func (o *ClientStatementRequest) HasShowPaymentsTable() bool {
	if o != nil && !IsNil(o.ShowPaymentsTable) {
		return true
	}

	return false
}

// SetShowPaymentsTable gets a reference to the given bool and assigns it to the ShowPaymentsTable field.
func (o *ClientStatementRequest) SetShowPaymentsTable(v bool) {
	o.ShowPaymentsTable = &v
}

// GetShowCreditsTable returns the ShowCreditsTable field value if set, zero value otherwise.
func (o *ClientStatementRequest) GetShowCreditsTable() bool {
	if o == nil || IsNil(o.ShowCreditsTable) {
		var ret bool
		return ret
	}
	return *o.ShowCreditsTable
}

// GetShowCreditsTableOk returns a tuple with the ShowCreditsTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientStatementRequest) GetShowCreditsTableOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowCreditsTable) {
		return nil, false
	}
	return o.ShowCreditsTable, true
}

// HasShowCreditsTable returns a boolean if a field has been set.
func (o *ClientStatementRequest) HasShowCreditsTable() bool {
	if o != nil && !IsNil(o.ShowCreditsTable) {
		return true
	}

	return false
}

// SetShowCreditsTable gets a reference to the given bool and assigns it to the ShowCreditsTable field.
func (o *ClientStatementRequest) SetShowCreditsTable(v bool) {
	o.ShowCreditsTable = &v
}

// GetShowAgingTable returns the ShowAgingTable field value if set, zero value otherwise.
func (o *ClientStatementRequest) GetShowAgingTable() bool {
	if o == nil || IsNil(o.ShowAgingTable) {
		var ret bool
		return ret
	}
	return *o.ShowAgingTable
}

// GetShowAgingTableOk returns a tuple with the ShowAgingTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientStatementRequest) GetShowAgingTableOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowAgingTable) {
		return nil, false
	}
	return o.ShowAgingTable, true
}

// HasShowAgingTable returns a boolean if a field has been set.
func (o *ClientStatementRequest) HasShowAgingTable() bool {
	if o != nil && !IsNil(o.ShowAgingTable) {
		return true
	}

	return false
}

// SetShowAgingTable gets a reference to the given bool and assigns it to the ShowAgingTable field.
func (o *ClientStatementRequest) SetShowAgingTable(v bool) {
	o.ShowAgingTable = &v
}

func (o ClientStatementRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientStatementRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.ShowPaymentsTable) {
		toSerialize["show_payments_table"] = o.ShowPaymentsTable
	}
	if !IsNil(o.ShowCreditsTable) {
		toSerialize["show_credits_table"] = o.ShowCreditsTable
	}
	if !IsNil(o.ShowAgingTable) {
		toSerialize["show_aging_table"] = o.ShowAgingTable
	}
	return toSerialize, nil
}

type NullableClientStatementRequest struct {
	value *ClientStatementRequest
	isSet bool
}

func (v NullableClientStatementRequest) Get() *ClientStatementRequest {
	return v.value
}

func (v *NullableClientStatementRequest) Set(val *ClientStatementRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableClientStatementRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableClientStatementRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientStatementRequest(val *ClientStatementRequest) *NullableClientStatementRequest {
	return &NullableClientStatementRequest{value: val, isSet: true}
}

func (v NullableClientStatementRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientStatementRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


