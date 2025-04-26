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

// checks if the FeesAndLimits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeesAndLimits{}

// FeesAndLimits struct for FeesAndLimits
type FeesAndLimits struct {
	// The minimum amount accepted for this gateway
	MinLimit *string `json:"min_limit,omitempty"`
	// The maximum amount accepted for this gateway
	MaxLimit *string `json:"max_limit,omitempty"`
	// The gateway fee amount
	FeeAmount *float32 `json:"fee_amount,omitempty"`
	// The gateway fee percentage
	FeePercent *float32 `json:"fee_percent,omitempty"`
	// Fee tax name
	FeeTaxName1 *string `json:"fee_tax_name1,omitempty"`
	// Fee tax name
	FeeTaxName2 *string `json:"fee_tax_name2,omitempty"`
	// Fee tax name
	FeeTaxName3 *string `json:"fee_tax_name3,omitempty"`
	// The tax rate
	FeeTaxRate1 *float32 `json:"fee_tax_rate1,omitempty"`
	// The tax rate
	FeeTaxRate2 *float32 `json:"fee_tax_rate2,omitempty"`
	// The tax rate
	FeeTaxRate3 *float32 `json:"fee_tax_rate3,omitempty"`
	// If set the fee amount will be no higher than this amount
	FeeCap *float32 `json:"fee_cap,omitempty"`
	// Adjusts the fee to match the exact gateway fee.
	AdjustFeePercent *bool `json:"adjust_fee_percent,omitempty"`
}

// NewFeesAndLimits instantiates a new FeesAndLimits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeesAndLimits() *FeesAndLimits {
	this := FeesAndLimits{}
	return &this
}

// NewFeesAndLimitsWithDefaults instantiates a new FeesAndLimits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeesAndLimitsWithDefaults() *FeesAndLimits {
	this := FeesAndLimits{}
	return &this
}

// GetMinLimit returns the MinLimit field value if set, zero value otherwise.
func (o *FeesAndLimits) GetMinLimit() string {
	if o == nil || IsNil(o.MinLimit) {
		var ret string
		return ret
	}
	return *o.MinLimit
}

// GetMinLimitOk returns a tuple with the MinLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeesAndLimits) GetMinLimitOk() (*string, bool) {
	if o == nil || IsNil(o.MinLimit) {
		return nil, false
	}
	return o.MinLimit, true
}

// HasMinLimit returns a boolean if a field has been set.
func (o *FeesAndLimits) HasMinLimit() bool {
	if o != nil && !IsNil(o.MinLimit) {
		return true
	}

	return false
}

// SetMinLimit gets a reference to the given string and assigns it to the MinLimit field.
func (o *FeesAndLimits) SetMinLimit(v string) {
	o.MinLimit = &v
}

// GetMaxLimit returns the MaxLimit field value if set, zero value otherwise.
func (o *FeesAndLimits) GetMaxLimit() string {
	if o == nil || IsNil(o.MaxLimit) {
		var ret string
		return ret
	}
	return *o.MaxLimit
}

// GetMaxLimitOk returns a tuple with the MaxLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeesAndLimits) GetMaxLimitOk() (*string, bool) {
	if o == nil || IsNil(o.MaxLimit) {
		return nil, false
	}
	return o.MaxLimit, true
}

// HasMaxLimit returns a boolean if a field has been set.
func (o *FeesAndLimits) HasMaxLimit() bool {
	if o != nil && !IsNil(o.MaxLimit) {
		return true
	}

	return false
}

// SetMaxLimit gets a reference to the given string and assigns it to the MaxLimit field.
func (o *FeesAndLimits) SetMaxLimit(v string) {
	o.MaxLimit = &v
}

// GetFeeAmount returns the FeeAmount field value if set, zero value otherwise.
func (o *FeesAndLimits) GetFeeAmount() float32 {
	if o == nil || IsNil(o.FeeAmount) {
		var ret float32
		return ret
	}
	return *o.FeeAmount
}

// GetFeeAmountOk returns a tuple with the FeeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeesAndLimits) GetFeeAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.FeeAmount) {
		return nil, false
	}
	return o.FeeAmount, true
}

// HasFeeAmount returns a boolean if a field has been set.
func (o *FeesAndLimits) HasFeeAmount() bool {
	if o != nil && !IsNil(o.FeeAmount) {
		return true
	}

	return false
}

// SetFeeAmount gets a reference to the given float32 and assigns it to the FeeAmount field.
func (o *FeesAndLimits) SetFeeAmount(v float32) {
	o.FeeAmount = &v
}

// GetFeePercent returns the FeePercent field value if set, zero value otherwise.
func (o *FeesAndLimits) GetFeePercent() float32 {
	if o == nil || IsNil(o.FeePercent) {
		var ret float32
		return ret
	}
	return *o.FeePercent
}

// GetFeePercentOk returns a tuple with the FeePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeesAndLimits) GetFeePercentOk() (*float32, bool) {
	if o == nil || IsNil(o.FeePercent) {
		return nil, false
	}
	return o.FeePercent, true
}

// HasFeePercent returns a boolean if a field has been set.
func (o *FeesAndLimits) HasFeePercent() bool {
	if o != nil && !IsNil(o.FeePercent) {
		return true
	}

	return false
}

// SetFeePercent gets a reference to the given float32 and assigns it to the FeePercent field.
func (o *FeesAndLimits) SetFeePercent(v float32) {
	o.FeePercent = &v
}

// GetFeeTaxName1 returns the FeeTaxName1 field value if set, zero value otherwise.
func (o *FeesAndLimits) GetFeeTaxName1() string {
	if o == nil || IsNil(o.FeeTaxName1) {
		var ret string
		return ret
	}
	return *o.FeeTaxName1
}

// GetFeeTaxName1Ok returns a tuple with the FeeTaxName1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeesAndLimits) GetFeeTaxName1Ok() (*string, bool) {
	if o == nil || IsNil(o.FeeTaxName1) {
		return nil, false
	}
	return o.FeeTaxName1, true
}

// HasFeeTaxName1 returns a boolean if a field has been set.
func (o *FeesAndLimits) HasFeeTaxName1() bool {
	if o != nil && !IsNil(o.FeeTaxName1) {
		return true
	}

	return false
}

// SetFeeTaxName1 gets a reference to the given string and assigns it to the FeeTaxName1 field.
func (o *FeesAndLimits) SetFeeTaxName1(v string) {
	o.FeeTaxName1 = &v
}

// GetFeeTaxName2 returns the FeeTaxName2 field value if set, zero value otherwise.
func (o *FeesAndLimits) GetFeeTaxName2() string {
	if o == nil || IsNil(o.FeeTaxName2) {
		var ret string
		return ret
	}
	return *o.FeeTaxName2
}

// GetFeeTaxName2Ok returns a tuple with the FeeTaxName2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeesAndLimits) GetFeeTaxName2Ok() (*string, bool) {
	if o == nil || IsNil(o.FeeTaxName2) {
		return nil, false
	}
	return o.FeeTaxName2, true
}

// HasFeeTaxName2 returns a boolean if a field has been set.
func (o *FeesAndLimits) HasFeeTaxName2() bool {
	if o != nil && !IsNil(o.FeeTaxName2) {
		return true
	}

	return false
}

// SetFeeTaxName2 gets a reference to the given string and assigns it to the FeeTaxName2 field.
func (o *FeesAndLimits) SetFeeTaxName2(v string) {
	o.FeeTaxName2 = &v
}

// GetFeeTaxName3 returns the FeeTaxName3 field value if set, zero value otherwise.
func (o *FeesAndLimits) GetFeeTaxName3() string {
	if o == nil || IsNil(o.FeeTaxName3) {
		var ret string
		return ret
	}
	return *o.FeeTaxName3
}

// GetFeeTaxName3Ok returns a tuple with the FeeTaxName3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeesAndLimits) GetFeeTaxName3Ok() (*string, bool) {
	if o == nil || IsNil(o.FeeTaxName3) {
		return nil, false
	}
	return o.FeeTaxName3, true
}

// HasFeeTaxName3 returns a boolean if a field has been set.
func (o *FeesAndLimits) HasFeeTaxName3() bool {
	if o != nil && !IsNil(o.FeeTaxName3) {
		return true
	}

	return false
}

// SetFeeTaxName3 gets a reference to the given string and assigns it to the FeeTaxName3 field.
func (o *FeesAndLimits) SetFeeTaxName3(v string) {
	o.FeeTaxName3 = &v
}

// GetFeeTaxRate1 returns the FeeTaxRate1 field value if set, zero value otherwise.
func (o *FeesAndLimits) GetFeeTaxRate1() float32 {
	if o == nil || IsNil(o.FeeTaxRate1) {
		var ret float32
		return ret
	}
	return *o.FeeTaxRate1
}

// GetFeeTaxRate1Ok returns a tuple with the FeeTaxRate1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeesAndLimits) GetFeeTaxRate1Ok() (*float32, bool) {
	if o == nil || IsNil(o.FeeTaxRate1) {
		return nil, false
	}
	return o.FeeTaxRate1, true
}

// HasFeeTaxRate1 returns a boolean if a field has been set.
func (o *FeesAndLimits) HasFeeTaxRate1() bool {
	if o != nil && !IsNil(o.FeeTaxRate1) {
		return true
	}

	return false
}

// SetFeeTaxRate1 gets a reference to the given float32 and assigns it to the FeeTaxRate1 field.
func (o *FeesAndLimits) SetFeeTaxRate1(v float32) {
	o.FeeTaxRate1 = &v
}

// GetFeeTaxRate2 returns the FeeTaxRate2 field value if set, zero value otherwise.
func (o *FeesAndLimits) GetFeeTaxRate2() float32 {
	if o == nil || IsNil(o.FeeTaxRate2) {
		var ret float32
		return ret
	}
	return *o.FeeTaxRate2
}

// GetFeeTaxRate2Ok returns a tuple with the FeeTaxRate2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeesAndLimits) GetFeeTaxRate2Ok() (*float32, bool) {
	if o == nil || IsNil(o.FeeTaxRate2) {
		return nil, false
	}
	return o.FeeTaxRate2, true
}

// HasFeeTaxRate2 returns a boolean if a field has been set.
func (o *FeesAndLimits) HasFeeTaxRate2() bool {
	if o != nil && !IsNil(o.FeeTaxRate2) {
		return true
	}

	return false
}

// SetFeeTaxRate2 gets a reference to the given float32 and assigns it to the FeeTaxRate2 field.
func (o *FeesAndLimits) SetFeeTaxRate2(v float32) {
	o.FeeTaxRate2 = &v
}

// GetFeeTaxRate3 returns the FeeTaxRate3 field value if set, zero value otherwise.
func (o *FeesAndLimits) GetFeeTaxRate3() float32 {
	if o == nil || IsNil(o.FeeTaxRate3) {
		var ret float32
		return ret
	}
	return *o.FeeTaxRate3
}

// GetFeeTaxRate3Ok returns a tuple with the FeeTaxRate3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeesAndLimits) GetFeeTaxRate3Ok() (*float32, bool) {
	if o == nil || IsNil(o.FeeTaxRate3) {
		return nil, false
	}
	return o.FeeTaxRate3, true
}

// HasFeeTaxRate3 returns a boolean if a field has been set.
func (o *FeesAndLimits) HasFeeTaxRate3() bool {
	if o != nil && !IsNil(o.FeeTaxRate3) {
		return true
	}

	return false
}

// SetFeeTaxRate3 gets a reference to the given float32 and assigns it to the FeeTaxRate3 field.
func (o *FeesAndLimits) SetFeeTaxRate3(v float32) {
	o.FeeTaxRate3 = &v
}

// GetFeeCap returns the FeeCap field value if set, zero value otherwise.
func (o *FeesAndLimits) GetFeeCap() float32 {
	if o == nil || IsNil(o.FeeCap) {
		var ret float32
		return ret
	}
	return *o.FeeCap
}

// GetFeeCapOk returns a tuple with the FeeCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeesAndLimits) GetFeeCapOk() (*float32, bool) {
	if o == nil || IsNil(o.FeeCap) {
		return nil, false
	}
	return o.FeeCap, true
}

// HasFeeCap returns a boolean if a field has been set.
func (o *FeesAndLimits) HasFeeCap() bool {
	if o != nil && !IsNil(o.FeeCap) {
		return true
	}

	return false
}

// SetFeeCap gets a reference to the given float32 and assigns it to the FeeCap field.
func (o *FeesAndLimits) SetFeeCap(v float32) {
	o.FeeCap = &v
}

// GetAdjustFeePercent returns the AdjustFeePercent field value if set, zero value otherwise.
func (o *FeesAndLimits) GetAdjustFeePercent() bool {
	if o == nil || IsNil(o.AdjustFeePercent) {
		var ret bool
		return ret
	}
	return *o.AdjustFeePercent
}

// GetAdjustFeePercentOk returns a tuple with the AdjustFeePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeesAndLimits) GetAdjustFeePercentOk() (*bool, bool) {
	if o == nil || IsNil(o.AdjustFeePercent) {
		return nil, false
	}
	return o.AdjustFeePercent, true
}

// HasAdjustFeePercent returns a boolean if a field has been set.
func (o *FeesAndLimits) HasAdjustFeePercent() bool {
	if o != nil && !IsNil(o.AdjustFeePercent) {
		return true
	}

	return false
}

// SetAdjustFeePercent gets a reference to the given bool and assigns it to the AdjustFeePercent field.
func (o *FeesAndLimits) SetAdjustFeePercent(v bool) {
	o.AdjustFeePercent = &v
}

func (o FeesAndLimits) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeesAndLimits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MinLimit) {
		toSerialize["min_limit"] = o.MinLimit
	}
	if !IsNil(o.MaxLimit) {
		toSerialize["max_limit"] = o.MaxLimit
	}
	if !IsNil(o.FeeAmount) {
		toSerialize["fee_amount"] = o.FeeAmount
	}
	if !IsNil(o.FeePercent) {
		toSerialize["fee_percent"] = o.FeePercent
	}
	if !IsNil(o.FeeTaxName1) {
		toSerialize["fee_tax_name1"] = o.FeeTaxName1
	}
	if !IsNil(o.FeeTaxName2) {
		toSerialize["fee_tax_name2"] = o.FeeTaxName2
	}
	if !IsNil(o.FeeTaxName3) {
		toSerialize["fee_tax_name3"] = o.FeeTaxName3
	}
	if !IsNil(o.FeeTaxRate1) {
		toSerialize["fee_tax_rate1"] = o.FeeTaxRate1
	}
	if !IsNil(o.FeeTaxRate2) {
		toSerialize["fee_tax_rate2"] = o.FeeTaxRate2
	}
	if !IsNil(o.FeeTaxRate3) {
		toSerialize["fee_tax_rate3"] = o.FeeTaxRate3
	}
	if !IsNil(o.FeeCap) {
		toSerialize["fee_cap"] = o.FeeCap
	}
	if !IsNil(o.AdjustFeePercent) {
		toSerialize["adjust_fee_percent"] = o.AdjustFeePercent
	}
	return toSerialize, nil
}

type NullableFeesAndLimits struct {
	value *FeesAndLimits
	isSet bool
}

func (v NullableFeesAndLimits) Get() *FeesAndLimits {
	return v.value
}

func (v *NullableFeesAndLimits) Set(val *FeesAndLimits) {
	v.value = val
	v.isSet = true
}

func (v NullableFeesAndLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableFeesAndLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeesAndLimits(val *FeesAndLimits) *NullableFeesAndLimits {
	return &NullableFeesAndLimits{value: val, isSet: true}
}

func (v NullableFeesAndLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeesAndLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


