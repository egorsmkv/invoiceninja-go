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

// checks if the ClientGatewayToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientGatewayToken{}

// ClientGatewayToken struct for ClientGatewayToken
type ClientGatewayToken struct {
	// The hashed id of the client gateway token
	Id *string `json:"id,omitempty"`
	// The hashed_id of the client
	ClientId *string `json:"client_id,omitempty"`
	// The payment token
	Token *string `json:"token,omitempty"`
	// THe bank account routing number
	RoutingNumber *string `json:"routing_number,omitempty"`
	// The hashed id of the company gateway
	CompanyGatewayId *string `json:"company_gateway_id,omitempty"`
	// Flag determining if the token is the default payment method
	IsDefault *bool `json:"is_default,omitempty"`
}

// NewClientGatewayToken instantiates a new ClientGatewayToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientGatewayToken() *ClientGatewayToken {
	this := ClientGatewayToken{}
	return &this
}

// NewClientGatewayTokenWithDefaults instantiates a new ClientGatewayToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientGatewayTokenWithDefaults() *ClientGatewayToken {
	this := ClientGatewayToken{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ClientGatewayToken) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGatewayToken) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ClientGatewayToken) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ClientGatewayToken) SetId(v string) {
	o.Id = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ClientGatewayToken) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGatewayToken) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ClientGatewayToken) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ClientGatewayToken) SetClientId(v string) {
	o.ClientId = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ClientGatewayToken) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGatewayToken) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ClientGatewayToken) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ClientGatewayToken) SetToken(v string) {
	o.Token = &v
}

// GetRoutingNumber returns the RoutingNumber field value if set, zero value otherwise.
func (o *ClientGatewayToken) GetRoutingNumber() string {
	if o == nil || IsNil(o.RoutingNumber) {
		var ret string
		return ret
	}
	return *o.RoutingNumber
}

// GetRoutingNumberOk returns a tuple with the RoutingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGatewayToken) GetRoutingNumberOk() (*string, bool) {
	if o == nil || IsNil(o.RoutingNumber) {
		return nil, false
	}
	return o.RoutingNumber, true
}

// HasRoutingNumber returns a boolean if a field has been set.
func (o *ClientGatewayToken) HasRoutingNumber() bool {
	if o != nil && !IsNil(o.RoutingNumber) {
		return true
	}

	return false
}

// SetRoutingNumber gets a reference to the given string and assigns it to the RoutingNumber field.
func (o *ClientGatewayToken) SetRoutingNumber(v string) {
	o.RoutingNumber = &v
}

// GetCompanyGatewayId returns the CompanyGatewayId field value if set, zero value otherwise.
func (o *ClientGatewayToken) GetCompanyGatewayId() string {
	if o == nil || IsNil(o.CompanyGatewayId) {
		var ret string
		return ret
	}
	return *o.CompanyGatewayId
}

// GetCompanyGatewayIdOk returns a tuple with the CompanyGatewayId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGatewayToken) GetCompanyGatewayIdOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyGatewayId) {
		return nil, false
	}
	return o.CompanyGatewayId, true
}

// HasCompanyGatewayId returns a boolean if a field has been set.
func (o *ClientGatewayToken) HasCompanyGatewayId() bool {
	if o != nil && !IsNil(o.CompanyGatewayId) {
		return true
	}

	return false
}

// SetCompanyGatewayId gets a reference to the given string and assigns it to the CompanyGatewayId field.
func (o *ClientGatewayToken) SetCompanyGatewayId(v string) {
	o.CompanyGatewayId = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *ClientGatewayToken) GetIsDefault() bool {
	if o == nil || IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientGatewayToken) GetIsDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *ClientGatewayToken) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *ClientGatewayToken) SetIsDefault(v bool) {
	o.IsDefault = &v
}

func (o ClientGatewayToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientGatewayToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.RoutingNumber) {
		toSerialize["routing_number"] = o.RoutingNumber
	}
	if !IsNil(o.CompanyGatewayId) {
		toSerialize["company_gateway_id"] = o.CompanyGatewayId
	}
	if !IsNil(o.IsDefault) {
		toSerialize["is_default"] = o.IsDefault
	}
	return toSerialize, nil
}

type NullableClientGatewayToken struct {
	value *ClientGatewayToken
	isSet bool
}

func (v NullableClientGatewayToken) Get() *ClientGatewayToken {
	return v.value
}

func (v *NullableClientGatewayToken) Set(val *ClientGatewayToken) {
	v.value = val
	v.isSet = true
}

func (v NullableClientGatewayToken) IsSet() bool {
	return v.isSet
}

func (v *NullableClientGatewayToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientGatewayToken(val *ClientGatewayToken) *NullableClientGatewayToken {
	return &NullableClientGatewayToken{value: val, isSet: true}
}

func (v NullableClientGatewayToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientGatewayToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


