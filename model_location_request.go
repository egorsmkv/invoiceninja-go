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

// checks if the LocationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationRequest{}

// LocationRequest struct for LocationRequest
type LocationRequest struct {
	// The location name
	Name *string `json:"name,omitempty"`
	// The first line of the address
	Address1 *string `json:"address1,omitempty"`
	// The second line of the address
	Address2 *string `json:"address2,omitempty"`
	// The city name
	City *string `json:"city,omitempty"`
	// The state or region
	State *string `json:"state,omitempty"`
	// The postal or zip code
	PostalCode *string `json:"postal_code,omitempty"`
	// The ID of the associated country
	CountryId *string `json:"country_id,omitempty"`
	// Custom field value 1
	CustomValue1 NullableString `json:"custom_value1,omitempty"`
	// Custom field value 2
	CustomValue2 NullableString `json:"custom_value2,omitempty"`
	// Custom field value 3
	CustomValue3 NullableString `json:"custom_value3,omitempty"`
	// Custom field value 4
	CustomValue4 NullableString `json:"custom_value4,omitempty"`
	// Indicates if this is a shipping location
	IsShippingLocation *bool `json:"is_shipping_location,omitempty"`
	// The assigned user hashed id
	AssignedUserId *string `json:"assigned_user_id,omitempty"`
	// The client hashed id
	ClientId *string `json:"client_id,omitempty"`
	// The vendor hashed id
	VendorId *string `json:"vendor_id,omitempty"`
}

// NewLocationRequest instantiates a new LocationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationRequest() *LocationRequest {
	this := LocationRequest{}
	return &this
}

// NewLocationRequestWithDefaults instantiates a new LocationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationRequestWithDefaults() *LocationRequest {
	this := LocationRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LocationRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LocationRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LocationRequest) SetName(v string) {
	o.Name = &v
}

// GetAddress1 returns the Address1 field value if set, zero value otherwise.
func (o *LocationRequest) GetAddress1() string {
	if o == nil || IsNil(o.Address1) {
		var ret string
		return ret
	}
	return *o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationRequest) GetAddress1Ok() (*string, bool) {
	if o == nil || IsNil(o.Address1) {
		return nil, false
	}
	return o.Address1, true
}

// HasAddress1 returns a boolean if a field has been set.
func (o *LocationRequest) HasAddress1() bool {
	if o != nil && !IsNil(o.Address1) {
		return true
	}

	return false
}

// SetAddress1 gets a reference to the given string and assigns it to the Address1 field.
func (o *LocationRequest) SetAddress1(v string) {
	o.Address1 = &v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *LocationRequest) GetAddress2() string {
	if o == nil || IsNil(o.Address2) {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationRequest) GetAddress2Ok() (*string, bool) {
	if o == nil || IsNil(o.Address2) {
		return nil, false
	}
	return o.Address2, true
}

// HasAddress2 returns a boolean if a field has been set.
func (o *LocationRequest) HasAddress2() bool {
	if o != nil && !IsNil(o.Address2) {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *LocationRequest) SetAddress2(v string) {
	o.Address2 = &v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *LocationRequest) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationRequest) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *LocationRequest) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *LocationRequest) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *LocationRequest) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationRequest) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *LocationRequest) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *LocationRequest) SetState(v string) {
	o.State = &v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *LocationRequest) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationRequest) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *LocationRequest) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *LocationRequest) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetCountryId returns the CountryId field value if set, zero value otherwise.
func (o *LocationRequest) GetCountryId() string {
	if o == nil || IsNil(o.CountryId) {
		var ret string
		return ret
	}
	return *o.CountryId
}

// GetCountryIdOk returns a tuple with the CountryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationRequest) GetCountryIdOk() (*string, bool) {
	if o == nil || IsNil(o.CountryId) {
		return nil, false
	}
	return o.CountryId, true
}

// HasCountryId returns a boolean if a field has been set.
func (o *LocationRequest) HasCountryId() bool {
	if o != nil && !IsNil(o.CountryId) {
		return true
	}

	return false
}

// SetCountryId gets a reference to the given string and assigns it to the CountryId field.
func (o *LocationRequest) SetCountryId(v string) {
	o.CountryId = &v
}

// GetCustomValue1 returns the CustomValue1 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationRequest) GetCustomValue1() string {
	if o == nil || IsNil(o.CustomValue1.Get()) {
		var ret string
		return ret
	}
	return *o.CustomValue1.Get()
}

// GetCustomValue1Ok returns a tuple with the CustomValue1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationRequest) GetCustomValue1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomValue1.Get(), o.CustomValue1.IsSet()
}

// HasCustomValue1 returns a boolean if a field has been set.
func (o *LocationRequest) HasCustomValue1() bool {
	if o != nil && o.CustomValue1.IsSet() {
		return true
	}

	return false
}

// SetCustomValue1 gets a reference to the given NullableString and assigns it to the CustomValue1 field.
func (o *LocationRequest) SetCustomValue1(v string) {
	o.CustomValue1.Set(&v)
}
// SetCustomValue1Nil sets the value for CustomValue1 to be an explicit nil
func (o *LocationRequest) SetCustomValue1Nil() {
	o.CustomValue1.Set(nil)
}

// UnsetCustomValue1 ensures that no value is present for CustomValue1, not even an explicit nil
func (o *LocationRequest) UnsetCustomValue1() {
	o.CustomValue1.Unset()
}

// GetCustomValue2 returns the CustomValue2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationRequest) GetCustomValue2() string {
	if o == nil || IsNil(o.CustomValue2.Get()) {
		var ret string
		return ret
	}
	return *o.CustomValue2.Get()
}

// GetCustomValue2Ok returns a tuple with the CustomValue2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationRequest) GetCustomValue2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomValue2.Get(), o.CustomValue2.IsSet()
}

// HasCustomValue2 returns a boolean if a field has been set.
func (o *LocationRequest) HasCustomValue2() bool {
	if o != nil && o.CustomValue2.IsSet() {
		return true
	}

	return false
}

// SetCustomValue2 gets a reference to the given NullableString and assigns it to the CustomValue2 field.
func (o *LocationRequest) SetCustomValue2(v string) {
	o.CustomValue2.Set(&v)
}
// SetCustomValue2Nil sets the value for CustomValue2 to be an explicit nil
func (o *LocationRequest) SetCustomValue2Nil() {
	o.CustomValue2.Set(nil)
}

// UnsetCustomValue2 ensures that no value is present for CustomValue2, not even an explicit nil
func (o *LocationRequest) UnsetCustomValue2() {
	o.CustomValue2.Unset()
}

// GetCustomValue3 returns the CustomValue3 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationRequest) GetCustomValue3() string {
	if o == nil || IsNil(o.CustomValue3.Get()) {
		var ret string
		return ret
	}
	return *o.CustomValue3.Get()
}

// GetCustomValue3Ok returns a tuple with the CustomValue3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationRequest) GetCustomValue3Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomValue3.Get(), o.CustomValue3.IsSet()
}

// HasCustomValue3 returns a boolean if a field has been set.
func (o *LocationRequest) HasCustomValue3() bool {
	if o != nil && o.CustomValue3.IsSet() {
		return true
	}

	return false
}

// SetCustomValue3 gets a reference to the given NullableString and assigns it to the CustomValue3 field.
func (o *LocationRequest) SetCustomValue3(v string) {
	o.CustomValue3.Set(&v)
}
// SetCustomValue3Nil sets the value for CustomValue3 to be an explicit nil
func (o *LocationRequest) SetCustomValue3Nil() {
	o.CustomValue3.Set(nil)
}

// UnsetCustomValue3 ensures that no value is present for CustomValue3, not even an explicit nil
func (o *LocationRequest) UnsetCustomValue3() {
	o.CustomValue3.Unset()
}

// GetCustomValue4 returns the CustomValue4 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocationRequest) GetCustomValue4() string {
	if o == nil || IsNil(o.CustomValue4.Get()) {
		var ret string
		return ret
	}
	return *o.CustomValue4.Get()
}

// GetCustomValue4Ok returns a tuple with the CustomValue4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocationRequest) GetCustomValue4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomValue4.Get(), o.CustomValue4.IsSet()
}

// HasCustomValue4 returns a boolean if a field has been set.
func (o *LocationRequest) HasCustomValue4() bool {
	if o != nil && o.CustomValue4.IsSet() {
		return true
	}

	return false
}

// SetCustomValue4 gets a reference to the given NullableString and assigns it to the CustomValue4 field.
func (o *LocationRequest) SetCustomValue4(v string) {
	o.CustomValue4.Set(&v)
}
// SetCustomValue4Nil sets the value for CustomValue4 to be an explicit nil
func (o *LocationRequest) SetCustomValue4Nil() {
	o.CustomValue4.Set(nil)
}

// UnsetCustomValue4 ensures that no value is present for CustomValue4, not even an explicit nil
func (o *LocationRequest) UnsetCustomValue4() {
	o.CustomValue4.Unset()
}

// GetIsShippingLocation returns the IsShippingLocation field value if set, zero value otherwise.
func (o *LocationRequest) GetIsShippingLocation() bool {
	if o == nil || IsNil(o.IsShippingLocation) {
		var ret bool
		return ret
	}
	return *o.IsShippingLocation
}

// GetIsShippingLocationOk returns a tuple with the IsShippingLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationRequest) GetIsShippingLocationOk() (*bool, bool) {
	if o == nil || IsNil(o.IsShippingLocation) {
		return nil, false
	}
	return o.IsShippingLocation, true
}

// HasIsShippingLocation returns a boolean if a field has been set.
func (o *LocationRequest) HasIsShippingLocation() bool {
	if o != nil && !IsNil(o.IsShippingLocation) {
		return true
	}

	return false
}

// SetIsShippingLocation gets a reference to the given bool and assigns it to the IsShippingLocation field.
func (o *LocationRequest) SetIsShippingLocation(v bool) {
	o.IsShippingLocation = &v
}

// GetAssignedUserId returns the AssignedUserId field value if set, zero value otherwise.
func (o *LocationRequest) GetAssignedUserId() string {
	if o == nil || IsNil(o.AssignedUserId) {
		var ret string
		return ret
	}
	return *o.AssignedUserId
}

// GetAssignedUserIdOk returns a tuple with the AssignedUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationRequest) GetAssignedUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssignedUserId) {
		return nil, false
	}
	return o.AssignedUserId, true
}

// HasAssignedUserId returns a boolean if a field has been set.
func (o *LocationRequest) HasAssignedUserId() bool {
	if o != nil && !IsNil(o.AssignedUserId) {
		return true
	}

	return false
}

// SetAssignedUserId gets a reference to the given string and assigns it to the AssignedUserId field.
func (o *LocationRequest) SetAssignedUserId(v string) {
	o.AssignedUserId = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *LocationRequest) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationRequest) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *LocationRequest) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *LocationRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetVendorId returns the VendorId field value if set, zero value otherwise.
func (o *LocationRequest) GetVendorId() string {
	if o == nil || IsNil(o.VendorId) {
		var ret string
		return ret
	}
	return *o.VendorId
}

// GetVendorIdOk returns a tuple with the VendorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationRequest) GetVendorIdOk() (*string, bool) {
	if o == nil || IsNil(o.VendorId) {
		return nil, false
	}
	return o.VendorId, true
}

// HasVendorId returns a boolean if a field has been set.
func (o *LocationRequest) HasVendorId() bool {
	if o != nil && !IsNil(o.VendorId) {
		return true
	}

	return false
}

// SetVendorId gets a reference to the given string and assigns it to the VendorId field.
func (o *LocationRequest) SetVendorId(v string) {
	o.VendorId = &v
}

func (o LocationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Address1) {
		toSerialize["address1"] = o.Address1
	}
	if !IsNil(o.Address2) {
		toSerialize["address2"] = o.Address2
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.PostalCode) {
		toSerialize["postal_code"] = o.PostalCode
	}
	if !IsNil(o.CountryId) {
		toSerialize["country_id"] = o.CountryId
	}
	if o.CustomValue1.IsSet() {
		toSerialize["custom_value1"] = o.CustomValue1.Get()
	}
	if o.CustomValue2.IsSet() {
		toSerialize["custom_value2"] = o.CustomValue2.Get()
	}
	if o.CustomValue3.IsSet() {
		toSerialize["custom_value3"] = o.CustomValue3.Get()
	}
	if o.CustomValue4.IsSet() {
		toSerialize["custom_value4"] = o.CustomValue4.Get()
	}
	if !IsNil(o.IsShippingLocation) {
		toSerialize["is_shipping_location"] = o.IsShippingLocation
	}
	if !IsNil(o.AssignedUserId) {
		toSerialize["assigned_user_id"] = o.AssignedUserId
	}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.VendorId) {
		toSerialize["vendor_id"] = o.VendorId
	}
	return toSerialize, nil
}

type NullableLocationRequest struct {
	value *LocationRequest
	isSet bool
}

func (v NullableLocationRequest) Get() *LocationRequest {
	return v.value
}

func (v *NullableLocationRequest) Set(val *LocationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationRequest(val *LocationRequest) *NullableLocationRequest {
	return &NullableLocationRequest{value: val, isSet: true}
}

func (v NullableLocationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


