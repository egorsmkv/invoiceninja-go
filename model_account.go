package openapi

import (
	"encoding/json"
)

// checks if the Account type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Account{}

// Account struct for Account
type Account struct {
	// The account hashed id
	Id *string `json:"id,omitempty"`
	// Boolean flag if the account has been verified by sms
	AccountSmsVerified *string `json:"account_sms_verified,omitempty"`
}

// NewAccount instantiates a new Account object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccount() *Account {
	this := Account{}
	return &this
}

// NewAccountWithDefaults instantiates a new Account object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountWithDefaults() *Account {
	this := Account{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Account) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Account) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Account) SetId(v string) {
	o.Id = &v
}

// GetAccountSmsVerified returns the AccountSmsVerified field value if set, zero value otherwise.
func (o *Account) GetAccountSmsVerified() string {
	if o == nil || IsNil(o.AccountSmsVerified) {
		var ret string
		return ret
	}
	return *o.AccountSmsVerified
}

// GetAccountSmsVerifiedOk returns a tuple with the AccountSmsVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Account) GetAccountSmsVerifiedOk() (*string, bool) {
	if o == nil || IsNil(o.AccountSmsVerified) {
		return nil, false
	}
	return o.AccountSmsVerified, true
}

// HasAccountSmsVerified returns a boolean if a field has been set.
func (o *Account) HasAccountSmsVerified() bool {
	if o != nil && !IsNil(o.AccountSmsVerified) {
		return true
	}

	return false
}

// SetAccountSmsVerified gets a reference to the given string and assigns it to the AccountSmsVerified field.
func (o *Account) SetAccountSmsVerified(v string) {
	o.AccountSmsVerified = &v
}

func (o Account) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Account) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.AccountSmsVerified) {
		toSerialize["account_sms_verified"] = o.AccountSmsVerified
	}
	return toSerialize, nil
}

type NullableAccount struct {
	value *Account
	isSet bool
}

func (v NullableAccount) Get() *Account {
	return v.value
}

func (v *NullableAccount) Set(val *Account) {
	v.value = val
	v.isSet = true
}

func (v NullableAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccount(val *Account) *NullableAccount {
	return &NullableAccount{value: val, isSet: true}
}

func (v NullableAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
