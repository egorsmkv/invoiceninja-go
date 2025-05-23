package openapi

import (
	"encoding/json"
)

// checks if the AuthorizationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizationError{}

// AuthorizationError struct for AuthorizationError
type AuthorizationError struct {
	// Insufficient permissions for this resource.
	Message *string                `json:"message,omitempty"`
	Errors  *ValidationErrorErrors `json:"errors,omitempty"`
}

// NewAuthorizationError instantiates a new AuthorizationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationError() *AuthorizationError {
	this := AuthorizationError{}
	return &this
}

// NewAuthorizationErrorWithDefaults instantiates a new AuthorizationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationErrorWithDefaults() *AuthorizationError {
	this := AuthorizationError{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AuthorizationError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AuthorizationError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AuthorizationError) SetMessage(v string) {
	o.Message = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *AuthorizationError) GetErrors() ValidationErrorErrors {
	if o == nil || IsNil(o.Errors) {
		var ret ValidationErrorErrors
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationError) GetErrorsOk() (*ValidationErrorErrors, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *AuthorizationError) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given ValidationErrorErrors and assigns it to the Errors field.
func (o *AuthorizationError) SetErrors(v ValidationErrorErrors) {
	o.Errors = &v
}

func (o AuthorizationError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizationError) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableAuthorizationError struct {
	value *AuthorizationError
	isSet bool
}

func (v NullableAuthorizationError) Get() *AuthorizationError {
	return v.value
}

func (v *NullableAuthorizationError) Set(val *AuthorizationError) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationError) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationError(val *AuthorizationError) *NullableAuthorizationError {
	return &NullableAuthorizationError{value: val, isSet: true}
}

func (v NullableAuthorizationError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
