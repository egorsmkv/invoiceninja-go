package openapi

import (
	"encoding/json"
)

// checks if the BTRules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BTRules{}

// BTRules struct for BTRules
type BTRules struct {
	// The key to search
	DataKey *string `json:"data_key,omitempty"`
	// The operator flag of the search
	Operator *string `json:"operator,omitempty"`
	// The value to search for
	Value *string `json:"value,omitempty"`
}

// NewBTRules instantiates a new BTRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTRules() *BTRules {
	this := BTRules{}
	return &this
}

// NewBTRulesWithDefaults instantiates a new BTRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTRulesWithDefaults() *BTRules {
	this := BTRules{}
	return &this
}

// GetDataKey returns the DataKey field value if set, zero value otherwise.
func (o *BTRules) GetDataKey() string {
	if o == nil || IsNil(o.DataKey) {
		var ret string
		return ret
	}
	return *o.DataKey
}

// GetDataKeyOk returns a tuple with the DataKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRules) GetDataKeyOk() (*string, bool) {
	if o == nil || IsNil(o.DataKey) {
		return nil, false
	}
	return o.DataKey, true
}

// HasDataKey returns a boolean if a field has been set.
func (o *BTRules) HasDataKey() bool {
	if o != nil && !IsNil(o.DataKey) {
		return true
	}

	return false
}

// SetDataKey gets a reference to the given string and assigns it to the DataKey field.
func (o *BTRules) SetDataKey(v string) {
	o.DataKey = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *BTRules) GetOperator() string {
	if o == nil || IsNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRules) GetOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *BTRules) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *BTRules) SetOperator(v string) {
	o.Operator = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTRules) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRules) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTRules) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *BTRules) SetValue(v string) {
	o.Value = &v
}

func (o BTRules) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BTRules) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.DataKey) {
		toSerialize["data_key"] = o.DataKey
	}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableBTRules struct {
	value *BTRules
	isSet bool
}

func (v NullableBTRules) Get() *BTRules {
	return v.value
}

func (v *NullableBTRules) Set(val *BTRules) {
	v.value = val
	v.isSet = true
}

func (v NullableBTRules) IsSet() bool {
	return v.isSet
}

func (v *NullableBTRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTRules(val *BTRules) *NullableBTRules {
	return &NullableBTRules{value: val, isSet: true}
}

func (v NullableBTRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
