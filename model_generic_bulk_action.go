package openapi

import (
	"encoding/json"
)

// checks if the GenericBulkAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenericBulkAction{}

// GenericBulkAction struct for GenericBulkAction
type GenericBulkAction struct {
	// The action to perform ie. archive / restore / delete
	Action *string  `json:"action,omitempty"`
	Ids    []string `json:"ids,omitempty"`
}

// NewGenericBulkAction instantiates a new GenericBulkAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericBulkAction() *GenericBulkAction {
	this := GenericBulkAction{}
	return &this
}

// NewGenericBulkActionWithDefaults instantiates a new GenericBulkAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericBulkActionWithDefaults() *GenericBulkAction {
	this := GenericBulkAction{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *GenericBulkAction) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericBulkAction) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *GenericBulkAction) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *GenericBulkAction) SetAction(v string) {
	o.Action = &v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *GenericBulkAction) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericBulkAction) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *GenericBulkAction) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *GenericBulkAction) SetIds(v []string) {
	o.Ids = v
}

func (o GenericBulkAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenericBulkAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	return toSerialize, nil
}

type NullableGenericBulkAction struct {
	value *GenericBulkAction
	isSet bool
}

func (v NullableGenericBulkAction) Get() *GenericBulkAction {
	return v.value
}

func (v *NullableGenericBulkAction) Set(val *GenericBulkAction) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericBulkAction) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericBulkAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericBulkAction(val *GenericBulkAction) *NullableGenericBulkAction {
	return &NullableGenericBulkAction{value: val, isSet: true}
}

func (v NullableGenericBulkAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericBulkAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
