package openapi

import (
	"encoding/json"
)

// checks if the InvalidInputError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvalidInputError{}

// InvalidInputError struct for InvalidInputError
type InvalidInputError struct {
	// Invalid input
	Message *string                `json:"message,omitempty"`
	Errors  *ValidationErrorErrors `json:"errors,omitempty"`
}

// NewInvalidInputError instantiates a new InvalidInputError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvalidInputError() *InvalidInputError {
	this := InvalidInputError{}
	return &this
}

// NewInvalidInputErrorWithDefaults instantiates a new InvalidInputError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvalidInputErrorWithDefaults() *InvalidInputError {
	this := InvalidInputError{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InvalidInputError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvalidInputError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InvalidInputError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InvalidInputError) SetMessage(v string) {
	o.Message = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *InvalidInputError) GetErrors() ValidationErrorErrors {
	if o == nil || IsNil(o.Errors) {
		var ret ValidationErrorErrors
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvalidInputError) GetErrorsOk() (*ValidationErrorErrors, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *InvalidInputError) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given ValidationErrorErrors and assigns it to the Errors field.
func (o *InvalidInputError) SetErrors(v ValidationErrorErrors) {
	o.Errors = &v
}

func (o InvalidInputError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvalidInputError) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableInvalidInputError struct {
	value *InvalidInputError
	isSet bool
}

func (v NullableInvalidInputError) Get() *InvalidInputError {
	return v.value
}

func (v *NullableInvalidInputError) Set(val *InvalidInputError) {
	v.value = val
	v.isSet = true
}

func (v NullableInvalidInputError) IsSet() bool {
	return v.isSet
}

func (v *NullableInvalidInputError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvalidInputError(val *InvalidInputError) *NullableInvalidInputError {
	return &NullableInvalidInputError{value: val, isSet: true}
}

func (v NullableInvalidInputError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvalidInputError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
