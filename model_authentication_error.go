package openapi

import (
	"encoding/json"
)

// checks if the AuthenticationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticationError{}

// AuthenticationError struct for AuthenticationError
type AuthenticationError struct {
	// These credentials do not match our records / Invalid Token
	Message *string `json:"message,omitempty"`
}

// NewAuthenticationError instantiates a new AuthenticationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationError() *AuthenticationError {
	this := AuthenticationError{}
	return &this
}

// NewAuthenticationErrorWithDefaults instantiates a new AuthenticationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationErrorWithDefaults() *AuthenticationError {
	this := AuthenticationError{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *AuthenticationError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *AuthenticationError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *AuthenticationError) SetMessage(v string) {
	o.Message = &v
}

func (o AuthenticationError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableAuthenticationError struct {
	value *AuthenticationError
	isSet bool
}

func (v NullableAuthenticationError) Get() *AuthenticationError {
	return v.value
}

func (v *NullableAuthenticationError) Set(val *AuthenticationError) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationError) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationError(val *AuthenticationError) *NullableAuthenticationError {
	return &NullableAuthenticationError{value: val, isSet: true}
}

func (v NullableAuthenticationError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
