package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PostLoginRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostLoginRequest{}

// PostLoginRequest struct for PostLoginRequest
type PostLoginRequest struct {
	// The users email address.
	Email string `json:"email"`
	// The user password. Must meet minimum criteria ~ > 6 characters
	Password string `json:"password"`
	// The one time password if 2FA is enabled
	OneTimePassword *string `json:"one_time_password,omitempty"`
}

type _PostLoginRequest PostLoginRequest

// NewPostLoginRequest instantiates a new PostLoginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostLoginRequest(email string, password string) *PostLoginRequest {
	this := PostLoginRequest{}
	this.Email = email
	this.Password = password
	return &this
}

// NewPostLoginRequestWithDefaults instantiates a new PostLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostLoginRequestWithDefaults() *PostLoginRequest {
	this := PostLoginRequest{}
	return &this
}

// GetEmail returns the Email field value
func (o *PostLoginRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *PostLoginRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *PostLoginRequest) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *PostLoginRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *PostLoginRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *PostLoginRequest) SetPassword(v string) {
	o.Password = v
}

// GetOneTimePassword returns the OneTimePassword field value if set, zero value otherwise.
func (o *PostLoginRequest) GetOneTimePassword() string {
	if o == nil || IsNil(o.OneTimePassword) {
		var ret string
		return ret
	}
	return *o.OneTimePassword
}

// GetOneTimePasswordOk returns a tuple with the OneTimePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLoginRequest) GetOneTimePasswordOk() (*string, bool) {
	if o == nil || IsNil(o.OneTimePassword) {
		return nil, false
	}
	return o.OneTimePassword, true
}

// HasOneTimePassword returns a boolean if a field has been set.
func (o *PostLoginRequest) HasOneTimePassword() bool {
	if o != nil && !IsNil(o.OneTimePassword) {
		return true
	}

	return false
}

// SetOneTimePassword gets a reference to the given string and assigns it to the OneTimePassword field.
func (o *PostLoginRequest) SetOneTimePassword(v string) {
	o.OneTimePassword = &v
}

func (o PostLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostLoginRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	toSerialize["password"] = o.Password
	if !IsNil(o.OneTimePassword) {
		toSerialize["one_time_password"] = o.OneTimePassword
	}
	return toSerialize, nil
}

func (o *PostLoginRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
		"password",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPostLoginRequest := _PostLoginRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostLoginRequest)

	if err != nil {
		return err
	}

	*o = PostLoginRequest(varPostLoginRequest)

	return err
}

type NullablePostLoginRequest struct {
	value *PostLoginRequest
	isSet bool
}

func (v NullablePostLoginRequest) Get() *PostLoginRequest {
	return v.value
}

func (v *NullablePostLoginRequest) Set(val *PostLoginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostLoginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostLoginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostLoginRequest(val *PostLoginRequest) *NullablePostLoginRequest {
	return &NullablePostLoginRequest{value: val, isSet: true}
}

func (v NullablePostLoginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostLoginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
