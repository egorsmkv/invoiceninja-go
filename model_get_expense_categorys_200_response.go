package openapi

import (
	"encoding/json"
)

// checks if the GetExpenseCategorys200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetExpenseCategorys200Response{}

// GetExpenseCategorys200Response struct for GetExpenseCategorys200Response
type GetExpenseCategorys200Response struct {
	Data []ExpenseCategory `json:"data,omitempty"`
	Meta Meta              `json:"meta,omitempty"`
}

// NewGetExpenseCategorys200Response instantiates a new GetExpenseCategorys200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetExpenseCategorys200Response() *GetExpenseCategorys200Response {
	this := GetExpenseCategorys200Response{}
	return &this
}

// NewGetExpenseCategorys200ResponseWithDefaults instantiates a new GetExpenseCategorys200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetExpenseCategorys200ResponseWithDefaults() *GetExpenseCategorys200Response {
	this := GetExpenseCategorys200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetExpenseCategorys200Response) GetData() []ExpenseCategory {
	if o == nil || IsNil(o.Data) {
		var ret []ExpenseCategory
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExpenseCategorys200Response) GetDataOk() ([]ExpenseCategory, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetExpenseCategorys200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ExpenseCategory and assigns it to the Data field.
func (o *GetExpenseCategorys200Response) SetData(v []ExpenseCategory) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetExpenseCategorys200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExpenseCategorys200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetExpenseCategorys200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetExpenseCategorys200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetExpenseCategorys200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetExpenseCategorys200Response) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetExpenseCategorys200Response struct {
	value *GetExpenseCategorys200Response
	isSet bool
}

func (v NullableGetExpenseCategorys200Response) Get() *GetExpenseCategorys200Response {
	return v.value
}

func (v *NullableGetExpenseCategorys200Response) Set(val *GetExpenseCategorys200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExpenseCategorys200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExpenseCategorys200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExpenseCategorys200Response(val *GetExpenseCategorys200Response) *NullableGetExpenseCategorys200Response {
	return &NullableGetExpenseCategorys200Response{value: val, isSet: true}
}

func (v NullableGetExpenseCategorys200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExpenseCategorys200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
