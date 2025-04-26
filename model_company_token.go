package openapi

import (
	"encoding/json"
)

// checks if the CompanyToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyToken{}

// CompanyToken struct for CompanyToken
type CompanyToken struct {
	// The token name
	Name *string `json:"name,omitempty"`
	// The token value
	Token *string `json:"token,omitempty"`
	// Determines whether the token is created by the system rather than a user
	IsSystem *bool `json:"is_system,omitempty"`
}

// NewCompanyToken instantiates a new CompanyToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyToken() *CompanyToken {
	this := CompanyToken{}
	return &this
}

// NewCompanyTokenWithDefaults instantiates a new CompanyToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyTokenWithDefaults() *CompanyToken {
	this := CompanyToken{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CompanyToken) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyToken) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CompanyToken) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CompanyToken) SetName(v string) {
	o.Name = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CompanyToken) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyToken) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CompanyToken) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CompanyToken) SetToken(v string) {
	o.Token = &v
}

// GetIsSystem returns the IsSystem field value if set, zero value otherwise.
func (o *CompanyToken) GetIsSystem() bool {
	if o == nil || IsNil(o.IsSystem) {
		var ret bool
		return ret
	}
	return *o.IsSystem
}

// GetIsSystemOk returns a tuple with the IsSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyToken) GetIsSystemOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSystem) {
		return nil, false
	}
	return o.IsSystem, true
}

// HasIsSystem returns a boolean if a field has been set.
func (o *CompanyToken) HasIsSystem() bool {
	if o != nil && !IsNil(o.IsSystem) {
		return true
	}

	return false
}

// SetIsSystem gets a reference to the given bool and assigns it to the IsSystem field.
func (o *CompanyToken) SetIsSystem(v bool) {
	o.IsSystem = &v
}

func (o CompanyToken) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.IsSystem) {
		toSerialize["is_system"] = o.IsSystem
	}
	return toSerialize, nil
}

type NullableCompanyToken struct {
	value *CompanyToken
	isSet bool
}

func (v NullableCompanyToken) Get() *CompanyToken {
	return v.value
}

func (v *NullableCompanyToken) Set(val *CompanyToken) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyToken) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyToken(val *CompanyToken) *NullableCompanyToken {
	return &NullableCompanyToken{value: val, isSet: true}
}

func (v NullableCompanyToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
