package openapi

import (
	"encoding/json"
)

// checks if the InvoicePaymentable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvoicePaymentable{}

// InvoicePaymentable struct for InvoicePaymentable
type InvoicePaymentable struct {
	// ______
	InvoiceId *string `json:"invoice_id,omitempty"`
	// ______
	Amount *string `json:"amount,omitempty"`
}

// NewInvoicePaymentable instantiates a new InvoicePaymentable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoicePaymentable() *InvoicePaymentable {
	this := InvoicePaymentable{}
	return &this
}

// NewInvoicePaymentableWithDefaults instantiates a new InvoicePaymentable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoicePaymentableWithDefaults() *InvoicePaymentable {
	this := InvoicePaymentable{}
	return &this
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise.
func (o *InvoicePaymentable) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoicePaymentable) GetInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceId) {
		return nil, false
	}
	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *InvoicePaymentable) HasInvoiceId() bool {
	if o != nil && !IsNil(o.InvoiceId) {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *InvoicePaymentable) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *InvoicePaymentable) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvoicePaymentable) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *InvoicePaymentable) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *InvoicePaymentable) SetAmount(v string) {
	o.Amount = &v
}

func (o InvoicePaymentable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvoicePaymentable) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.InvoiceId) {
		toSerialize["invoice_id"] = o.InvoiceId
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	return toSerialize, nil
}

type NullableInvoicePaymentable struct {
	value *InvoicePaymentable
	isSet bool
}

func (v NullableInvoicePaymentable) Get() *InvoicePaymentable {
	return v.value
}

func (v *NullableInvoicePaymentable) Set(val *InvoicePaymentable) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoicePaymentable) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoicePaymentable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoicePaymentable(val *InvoicePaymentable) *NullableInvoicePaymentable {
	return &NullableInvoicePaymentable{value: val, isSet: true}
}

func (v NullableInvoicePaymentable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoicePaymentable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
