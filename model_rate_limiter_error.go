package openapi

import (
	"encoding/json"
)

// checks if the RateLimiterError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RateLimiterError{}

// RateLimiterError struct for RateLimiterError
type RateLimiterError struct {
	// Rate limit exceeded.
	Message *string                `json:"message,omitempty"`
	Errors  *ValidationErrorErrors `json:"errors,omitempty"`
}

// NewRateLimiterError instantiates a new RateLimiterError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateLimiterError() *RateLimiterError {
	this := RateLimiterError{}
	return &this
}

// NewRateLimiterErrorWithDefaults instantiates a new RateLimiterError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateLimiterErrorWithDefaults() *RateLimiterError {
	this := RateLimiterError{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RateLimiterError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimiterError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RateLimiterError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RateLimiterError) SetMessage(v string) {
	o.Message = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *RateLimiterError) GetErrors() ValidationErrorErrors {
	if o == nil || IsNil(o.Errors) {
		var ret ValidationErrorErrors
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimiterError) GetErrorsOk() (*ValidationErrorErrors, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *RateLimiterError) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given ValidationErrorErrors and assigns it to the Errors field.
func (o *RateLimiterError) SetErrors(v ValidationErrorErrors) {
	o.Errors = &v
}

func (o RateLimiterError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RateLimiterError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableRateLimiterError struct {
	value *RateLimiterError
	isSet bool
}

func (v NullableRateLimiterError) Get() *RateLimiterError {
	return v.value
}

func (v *NullableRateLimiterError) Set(val *RateLimiterError) {
	v.value = val
	v.isSet = true
}

func (v NullableRateLimiterError) IsSet() bool {
	return v.isSet
}

func (v *NullableRateLimiterError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateLimiterError(val *RateLimiterError) *NullableRateLimiterError {
	return &NullableRateLimiterError{value: val, isSet: true}
}

func (v NullableRateLimiterError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateLimiterError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
