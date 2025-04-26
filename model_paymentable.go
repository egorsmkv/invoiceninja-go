package openapi

import (
	"encoding/json"
)

// checks if the Paymentable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Paymentable{}

// Paymentable struct for Paymentable
type Paymentable struct {
	// The paymentable hashed id
	Id *string `json:"id,omitempty"`
	// The invoice hashed id
	InvoiceId *string `json:"invoice_id,omitempty"`
	// The credit hashed id
	CreditId *string `json:"credit_id,omitempty"`
	// The amount that has been refunded for this payment
	Refunded *float32 `json:"refunded,omitempty"`
	// The amount that has been applied to the payment
	Amount *float32 `json:"amount,omitempty"`
	// Timestamp
	UpdatedAt *float32 `json:"updated_at,omitempty"`
	// Timestamp
	CreatedAt *float32 `json:"created_at,omitempty"`
}

// NewPaymentable instantiates a new Paymentable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentable() *Paymentable {
	this := Paymentable{}
	return &this
}

// NewPaymentableWithDefaults instantiates a new Paymentable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentableWithDefaults() *Paymentable {
	this := Paymentable{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Paymentable) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Paymentable) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Paymentable) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Paymentable) SetId(v string) {
	o.Id = &v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise.
func (o *Paymentable) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Paymentable) GetInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceId) {
		return nil, false
	}
	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *Paymentable) HasInvoiceId() bool {
	if o != nil && !IsNil(o.InvoiceId) {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *Paymentable) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

// GetCreditId returns the CreditId field value if set, zero value otherwise.
func (o *Paymentable) GetCreditId() string {
	if o == nil || IsNil(o.CreditId) {
		var ret string
		return ret
	}
	return *o.CreditId
}

// GetCreditIdOk returns a tuple with the CreditId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Paymentable) GetCreditIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreditId) {
		return nil, false
	}
	return o.CreditId, true
}

// HasCreditId returns a boolean if a field has been set.
func (o *Paymentable) HasCreditId() bool {
	if o != nil && !IsNil(o.CreditId) {
		return true
	}

	return false
}

// SetCreditId gets a reference to the given string and assigns it to the CreditId field.
func (o *Paymentable) SetCreditId(v string) {
	o.CreditId = &v
}

// GetRefunded returns the Refunded field value if set, zero value otherwise.
func (o *Paymentable) GetRefunded() float32 {
	if o == nil || IsNil(o.Refunded) {
		var ret float32
		return ret
	}
	return *o.Refunded
}

// GetRefundedOk returns a tuple with the Refunded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Paymentable) GetRefundedOk() (*float32, bool) {
	if o == nil || IsNil(o.Refunded) {
		return nil, false
	}
	return o.Refunded, true
}

// HasRefunded returns a boolean if a field has been set.
func (o *Paymentable) HasRefunded() bool {
	if o != nil && !IsNil(o.Refunded) {
		return true
	}

	return false
}

// SetRefunded gets a reference to the given float32 and assigns it to the Refunded field.
func (o *Paymentable) SetRefunded(v float32) {
	o.Refunded = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Paymentable) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Paymentable) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Paymentable) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *Paymentable) SetAmount(v float32) {
	o.Amount = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Paymentable) GetUpdatedAt() float32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret float32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Paymentable) GetUpdatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Paymentable) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given float32 and assigns it to the UpdatedAt field.
func (o *Paymentable) SetUpdatedAt(v float32) {
	o.UpdatedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Paymentable) GetCreatedAt() float32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret float32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Paymentable) GetCreatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Paymentable) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given float32 and assigns it to the CreatedAt field.
func (o *Paymentable) SetCreatedAt(v float32) {
	o.CreatedAt = &v
}

func (o Paymentable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Paymentable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InvoiceId) {
		toSerialize["invoice_id"] = o.InvoiceId
	}
	if !IsNil(o.CreditId) {
		toSerialize["credit_id"] = o.CreditId
	}
	if !IsNil(o.Refunded) {
		toSerialize["refunded"] = o.Refunded
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullablePaymentable struct {
	value *Paymentable
	isSet bool
}

func (v NullablePaymentable) Get() *Paymentable {
	return v.value
}

func (v *NullablePaymentable) Set(val *Paymentable) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentable) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentable(val *Paymentable) *NullablePaymentable {
	return &NullablePaymentable{value: val, isSet: true}
}

func (v NullablePaymentable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
