package openapi

import (
	"encoding/json"
)

// checks if the TaxRate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxRate{}

// TaxRate struct for TaxRate
type TaxRate struct {
	// Thie hashed id of the tax
	Id *string `json:"id,omitempty"`
	// The tax name
	Name *string `json:"name,omitempty"`
	// The tax rate
	Rate *float32 `json:"rate,omitempty"`
	// Boolean flag determining if the tax has been deleted
	IsDeleted *bool `json:"is_deleted,omitempty"`
}

// NewTaxRate instantiates a new TaxRate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxRate() *TaxRate {
	this := TaxRate{}
	return &this
}

// NewTaxRateWithDefaults instantiates a new TaxRate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxRateWithDefaults() *TaxRate {
	this := TaxRate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TaxRate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxRate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TaxRate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TaxRate) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TaxRate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxRate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TaxRate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TaxRate) SetName(v string) {
	o.Name = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *TaxRate) GetRate() float32 {
	if o == nil || IsNil(o.Rate) {
		var ret float32
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxRate) GetRateOk() (*float32, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *TaxRate) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given float32 and assigns it to the Rate field.
func (o *TaxRate) SetRate(v float32) {
	o.Rate = &v
}

// GetIsDeleted returns the IsDeleted field value if set, zero value otherwise.
func (o *TaxRate) GetIsDeleted() bool {
	if o == nil || IsNil(o.IsDeleted) {
		var ret bool
		return ret
	}
	return *o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxRate) GetIsDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDeleted) {
		return nil, false
	}
	return o.IsDeleted, true
}

// HasIsDeleted returns a boolean if a field has been set.
func (o *TaxRate) HasIsDeleted() bool {
	if o != nil && !IsNil(o.IsDeleted) {
		return true
	}

	return false
}

// SetIsDeleted gets a reference to the given bool and assigns it to the IsDeleted field.
func (o *TaxRate) SetIsDeleted(v bool) {
	o.IsDeleted = &v
}

func (o TaxRate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxRate) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	if !IsNil(o.IsDeleted) {
		toSerialize["is_deleted"] = o.IsDeleted
	}
	return toSerialize, nil
}

type NullableTaxRate struct {
	value *TaxRate
	isSet bool
}

func (v NullableTaxRate) Get() *TaxRate {
	return v.value
}

func (v *NullableTaxRate) Set(val *TaxRate) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxRate) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxRate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxRate(val *TaxRate) *NullableTaxRate {
	return &NullableTaxRate{value: val, isSet: true}
}

func (v NullableTaxRate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
