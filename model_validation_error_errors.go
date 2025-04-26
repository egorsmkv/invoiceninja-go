package openapi

import (
	"encoding/json"
)

// checks if the ValidationErrorErrors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidationErrorErrors{}

// ValidationErrorErrors struct for ValidationErrorErrors
type ValidationErrorErrors struct {
	Value []string `json:"value,omitempty"`
}

// NewValidationErrorErrors instantiates a new ValidationErrorErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidationErrorErrors() *ValidationErrorErrors {
	this := ValidationErrorErrors{}
	return &this
}

// NewValidationErrorErrorsWithDefaults instantiates a new ValidationErrorErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationErrorErrorsWithDefaults() *ValidationErrorErrors {
	this := ValidationErrorErrors{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ValidationErrorErrors) GetValue() []string {
	if o == nil || IsNil(o.Value) {
		var ret []string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationErrorErrors) GetValueOk() ([]string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ValidationErrorErrors) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given []string and assigns it to the Value field.
func (o *ValidationErrorErrors) SetValue(v []string) {
	o.Value = v
}

func (o ValidationErrorErrors) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidationErrorErrors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableValidationErrorErrors struct {
	value *ValidationErrorErrors
	isSet bool
}

func (v NullableValidationErrorErrors) Get() *ValidationErrorErrors {
	return v.value
}

func (v *NullableValidationErrorErrors) Set(val *ValidationErrorErrors) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationErrorErrors) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationErrorErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationErrorErrors(val *ValidationErrorErrors) *NullableValidationErrorErrors {
	return &NullableValidationErrorErrors{value: val, isSet: true}
}

func (v NullableValidationErrorErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationErrorErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
