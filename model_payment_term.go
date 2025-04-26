package openapi

import (
	"encoding/json"
)

// checks if the PaymentTerm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentTerm{}

// PaymentTerm struct for PaymentTerm
type PaymentTerm struct {
	// The payment term length in days
	NumDays *int32 `json:"num_days,omitempty"`
	// The payment term length in string format
	Name *string `json:"name,omitempty"`
	// Timestamp
	CreatedAt *float32 `json:"created_at,omitempty"`
	// Timestamp
	UpdatedAt *float32 `json:"updated_at,omitempty"`
	// Timestamp
	ArchivedAt *float32 `json:"archived_at,omitempty"`
}

// NewPaymentTerm instantiates a new PaymentTerm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentTerm() *PaymentTerm {
	this := PaymentTerm{}
	return &this
}

// NewPaymentTermWithDefaults instantiates a new PaymentTerm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentTermWithDefaults() *PaymentTerm {
	this := PaymentTerm{}
	return &this
}

// GetNumDays returns the NumDays field value if set, zero value otherwise.
func (o *PaymentTerm) GetNumDays() int32 {
	if o == nil || IsNil(o.NumDays) {
		var ret int32
		return ret
	}
	return *o.NumDays
}

// GetNumDaysOk returns a tuple with the NumDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTerm) GetNumDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.NumDays) {
		return nil, false
	}
	return o.NumDays, true
}

// HasNumDays returns a boolean if a field has been set.
func (o *PaymentTerm) HasNumDays() bool {
	if o != nil && !IsNil(o.NumDays) {
		return true
	}

	return false
}

// SetNumDays gets a reference to the given int32 and assigns it to the NumDays field.
func (o *PaymentTerm) SetNumDays(v int32) {
	o.NumDays = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PaymentTerm) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTerm) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PaymentTerm) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PaymentTerm) SetName(v string) {
	o.Name = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PaymentTerm) GetCreatedAt() float32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret float32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTerm) GetCreatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PaymentTerm) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given float32 and assigns it to the CreatedAt field.
func (o *PaymentTerm) SetCreatedAt(v float32) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PaymentTerm) GetUpdatedAt() float32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret float32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTerm) GetUpdatedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PaymentTerm) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given float32 and assigns it to the UpdatedAt field.
func (o *PaymentTerm) SetUpdatedAt(v float32) {
	o.UpdatedAt = &v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *PaymentTerm) GetArchivedAt() float32 {
	if o == nil || IsNil(o.ArchivedAt) {
		var ret float32
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentTerm) GetArchivedAtOk() (*float32, bool) {
	if o == nil || IsNil(o.ArchivedAt) {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *PaymentTerm) HasArchivedAt() bool {
	if o != nil && !IsNil(o.ArchivedAt) {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given float32 and assigns it to the ArchivedAt field.
func (o *PaymentTerm) SetArchivedAt(v float32) {
	o.ArchivedAt = &v
}

func (o PaymentTerm) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentTerm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NumDays) {
		toSerialize["num_days"] = o.NumDays
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.ArchivedAt) {
		toSerialize["archived_at"] = o.ArchivedAt
	}
	return toSerialize, nil
}

type NullablePaymentTerm struct {
	value *PaymentTerm
	isSet bool
}

func (v NullablePaymentTerm) Get() *PaymentTerm {
	return v.value
}

func (v *NullablePaymentTerm) Set(val *PaymentTerm) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentTerm) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentTerm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentTerm(val *PaymentTerm) *NullablePaymentTerm {
	return &NullablePaymentTerm{value: val, isSet: true}
}

func (v NullablePaymentTerm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentTerm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
