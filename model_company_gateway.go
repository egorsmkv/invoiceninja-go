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

// checks if the CompanyGateway type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyGateway{}

// CompanyGateway struct for CompanyGateway
type CompanyGateway struct {
	// The hashed id of the company gateway
	Id *string `json:"id,omitempty"`
	// The gateway key (hash)
	GatewayKey *string `json:"gateway_key,omitempty"`
	// Bitmask representation of cards
	AcceptedCreditCards *int32 `json:"accepted_credit_cards,omitempty"`
	// Determines if the the billing address is required prior to payment.
	RequireBillingAddress *bool `json:"require_billing_address,omitempty"`
	// Determines if the the billing address is required prior to payment.
	RequireShippingAddress *bool `json:"require_shipping_address,omitempty"`
	// The configuration map for the gateway
	Config *string `json:"config,omitempty"`
	// Determines if the client details should be updated.
	UpdateDetails *bool `json:"update_details,omitempty"`
	// A mapped collection of the fees and limits for the configured gateway
	FeesAndLimits []FeesAndLimits `json:"fees_and_limits,omitempty"`
}

// NewCompanyGateway instantiates a new CompanyGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyGateway() *CompanyGateway {
	this := CompanyGateway{}
	return &this
}

// NewCompanyGatewayWithDefaults instantiates a new CompanyGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyGatewayWithDefaults() *CompanyGateway {
	this := CompanyGateway{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CompanyGateway) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyGateway) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CompanyGateway) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CompanyGateway) SetId(v string) {
	o.Id = &v
}

// GetGatewayKey returns the GatewayKey field value if set, zero value otherwise.
func (o *CompanyGateway) GetGatewayKey() string {
	if o == nil || IsNil(o.GatewayKey) {
		var ret string
		return ret
	}
	return *o.GatewayKey
}

// GetGatewayKeyOk returns a tuple with the GatewayKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyGateway) GetGatewayKeyOk() (*string, bool) {
	if o == nil || IsNil(o.GatewayKey) {
		return nil, false
	}
	return o.GatewayKey, true
}

// HasGatewayKey returns a boolean if a field has been set.
func (o *CompanyGateway) HasGatewayKey() bool {
	if o != nil && !IsNil(o.GatewayKey) {
		return true
	}

	return false
}

// SetGatewayKey gets a reference to the given string and assigns it to the GatewayKey field.
func (o *CompanyGateway) SetGatewayKey(v string) {
	o.GatewayKey = &v
}

// GetAcceptedCreditCards returns the AcceptedCreditCards field value if set, zero value otherwise.
func (o *CompanyGateway) GetAcceptedCreditCards() int32 {
	if o == nil || IsNil(o.AcceptedCreditCards) {
		var ret int32
		return ret
	}
	return *o.AcceptedCreditCards
}

// GetAcceptedCreditCardsOk returns a tuple with the AcceptedCreditCards field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyGateway) GetAcceptedCreditCardsOk() (*int32, bool) {
	if o == nil || IsNil(o.AcceptedCreditCards) {
		return nil, false
	}
	return o.AcceptedCreditCards, true
}

// HasAcceptedCreditCards returns a boolean if a field has been set.
func (o *CompanyGateway) HasAcceptedCreditCards() bool {
	if o != nil && !IsNil(o.AcceptedCreditCards) {
		return true
	}

	return false
}

// SetAcceptedCreditCards gets a reference to the given int32 and assigns it to the AcceptedCreditCards field.
func (o *CompanyGateway) SetAcceptedCreditCards(v int32) {
	o.AcceptedCreditCards = &v
}

// GetRequireBillingAddress returns the RequireBillingAddress field value if set, zero value otherwise.
func (o *CompanyGateway) GetRequireBillingAddress() bool {
	if o == nil || IsNil(o.RequireBillingAddress) {
		var ret bool
		return ret
	}
	return *o.RequireBillingAddress
}

// GetRequireBillingAddressOk returns a tuple with the RequireBillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyGateway) GetRequireBillingAddressOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireBillingAddress) {
		return nil, false
	}
	return o.RequireBillingAddress, true
}

// HasRequireBillingAddress returns a boolean if a field has been set.
func (o *CompanyGateway) HasRequireBillingAddress() bool {
	if o != nil && !IsNil(o.RequireBillingAddress) {
		return true
	}

	return false
}

// SetRequireBillingAddress gets a reference to the given bool and assigns it to the RequireBillingAddress field.
func (o *CompanyGateway) SetRequireBillingAddress(v bool) {
	o.RequireBillingAddress = &v
}

// GetRequireShippingAddress returns the RequireShippingAddress field value if set, zero value otherwise.
func (o *CompanyGateway) GetRequireShippingAddress() bool {
	if o == nil || IsNil(o.RequireShippingAddress) {
		var ret bool
		return ret
	}
	return *o.RequireShippingAddress
}

// GetRequireShippingAddressOk returns a tuple with the RequireShippingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyGateway) GetRequireShippingAddressOk() (*bool, bool) {
	if o == nil || IsNil(o.RequireShippingAddress) {
		return nil, false
	}
	return o.RequireShippingAddress, true
}

// HasRequireShippingAddress returns a boolean if a field has been set.
func (o *CompanyGateway) HasRequireShippingAddress() bool {
	if o != nil && !IsNil(o.RequireShippingAddress) {
		return true
	}

	return false
}

// SetRequireShippingAddress gets a reference to the given bool and assigns it to the RequireShippingAddress field.
func (o *CompanyGateway) SetRequireShippingAddress(v bool) {
	o.RequireShippingAddress = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *CompanyGateway) GetConfig() string {
	if o == nil || IsNil(o.Config) {
		var ret string
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyGateway) GetConfigOk() (*string, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *CompanyGateway) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given string and assigns it to the Config field.
func (o *CompanyGateway) SetConfig(v string) {
	o.Config = &v
}

// GetUpdateDetails returns the UpdateDetails field value if set, zero value otherwise.
func (o *CompanyGateway) GetUpdateDetails() bool {
	if o == nil || IsNil(o.UpdateDetails) {
		var ret bool
		return ret
	}
	return *o.UpdateDetails
}

// GetUpdateDetailsOk returns a tuple with the UpdateDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyGateway) GetUpdateDetailsOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateDetails) {
		return nil, false
	}
	return o.UpdateDetails, true
}

// HasUpdateDetails returns a boolean if a field has been set.
func (o *CompanyGateway) HasUpdateDetails() bool {
	if o != nil && !IsNil(o.UpdateDetails) {
		return true
	}

	return false
}

// SetUpdateDetails gets a reference to the given bool and assigns it to the UpdateDetails field.
func (o *CompanyGateway) SetUpdateDetails(v bool) {
	o.UpdateDetails = &v
}

// GetFeesAndLimits returns the FeesAndLimits field value if set, zero value otherwise.
func (o *CompanyGateway) GetFeesAndLimits() []FeesAndLimits {
	if o == nil || IsNil(o.FeesAndLimits) {
		var ret []FeesAndLimits
		return ret
	}
	return o.FeesAndLimits
}

// GetFeesAndLimitsOk returns a tuple with the FeesAndLimits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyGateway) GetFeesAndLimitsOk() ([]FeesAndLimits, bool) {
	if o == nil || IsNil(o.FeesAndLimits) {
		return nil, false
	}
	return o.FeesAndLimits, true
}

// HasFeesAndLimits returns a boolean if a field has been set.
func (o *CompanyGateway) HasFeesAndLimits() bool {
	if o != nil && !IsNil(o.FeesAndLimits) {
		return true
	}

	return false
}

// SetFeesAndLimits gets a reference to the given []FeesAndLimits and assigns it to the FeesAndLimits field.
func (o *CompanyGateway) SetFeesAndLimits(v []FeesAndLimits) {
	o.FeesAndLimits = v
}

func (o CompanyGateway) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyGateway) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.GatewayKey) {
		toSerialize["gateway_key"] = o.GatewayKey
	}
	if !IsNil(o.AcceptedCreditCards) {
		toSerialize["accepted_credit_cards"] = o.AcceptedCreditCards
	}
	if !IsNil(o.RequireBillingAddress) {
		toSerialize["require_billing_address"] = o.RequireBillingAddress
	}
	if !IsNil(o.RequireShippingAddress) {
		toSerialize["require_shipping_address"] = o.RequireShippingAddress
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.UpdateDetails) {
		toSerialize["update_details"] = o.UpdateDetails
	}
	if !IsNil(o.FeesAndLimits) {
		toSerialize["fees_and_limits"] = o.FeesAndLimits
	}
	return toSerialize, nil
}

type NullableCompanyGateway struct {
	value *CompanyGateway
	isSet bool
}

func (v NullableCompanyGateway) Get() *CompanyGateway {
	return v.value
}

func (v *NullableCompanyGateway) Set(val *CompanyGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyGateway(val *CompanyGateway) *NullableCompanyGateway {
	return &NullableCompanyGateway{value: val, isSet: true}
}

func (v NullableCompanyGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


