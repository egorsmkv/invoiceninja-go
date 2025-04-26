package openapi

import (
	"encoding/json"
)

// checks if the CreditPaymentable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditPaymentable{}

// CreditPaymentable struct for CreditPaymentable
type CreditPaymentable struct {
	// The credit hashed id
	CreditId *string `json:"credit_id,omitempty"`
	// The credit amount
	Amount *string `json:"amount,omitempty"`
}

// NewCreditPaymentable instantiates a new CreditPaymentable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditPaymentable() *CreditPaymentable {
	this := CreditPaymentable{}
	return &this
}

// NewCreditPaymentableWithDefaults instantiates a new CreditPaymentable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditPaymentableWithDefaults() *CreditPaymentable {
	this := CreditPaymentable{}
	return &this
}

// GetCreditId returns the CreditId field value if set, zero value otherwise.
func (o *CreditPaymentable) GetCreditId() string {
	if o == nil || IsNil(o.CreditId) {
		var ret string
		return ret
	}
	return *o.CreditId
}

// GetCreditIdOk returns a tuple with the CreditId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditPaymentable) GetCreditIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreditId) {
		return nil, false
	}
	return o.CreditId, true
}

// HasCreditId returns a boolean if a field has been set.
func (o *CreditPaymentable) HasCreditId() bool {
	if o != nil && !IsNil(o.CreditId) {
		return true
	}

	return false
}

// SetCreditId gets a reference to the given string and assigns it to the CreditId field.
func (o *CreditPaymentable) SetCreditId(v string) {
	o.CreditId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *CreditPaymentable) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditPaymentable) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *CreditPaymentable) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *CreditPaymentable) SetAmount(v string) {
	o.Amount = &v
}

func (o CreditPaymentable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditPaymentable) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.CreditId) {
		toSerialize["credit_id"] = o.CreditId
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	return toSerialize, nil
}

type NullableCreditPaymentable struct {
	value *CreditPaymentable
	isSet bool
}

func (v NullableCreditPaymentable) Get() *CreditPaymentable {
	return v.value
}

func (v *NullableCreditPaymentable) Set(val *CreditPaymentable) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditPaymentable) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditPaymentable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditPaymentable(val *CreditPaymentable) *NullableCreditPaymentable {
	return &NullableCreditPaymentable{value: val, isSet: true}
}

func (v NullableCreditPaymentable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditPaymentable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
