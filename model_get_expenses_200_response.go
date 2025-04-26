package openapi

import (
	"encoding/json"
)

// checks if the GetExpenses200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetExpenses200Response{}

// GetExpenses200Response struct for GetExpenses200Response
type GetExpenses200Response struct {
	Data []Expense `json:"data,omitempty"`
	Meta Meta      `json:"meta,omitempty"`
}

// NewGetExpenses200Response instantiates a new GetExpenses200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetExpenses200Response() *GetExpenses200Response {
	this := GetExpenses200Response{}
	return &this
}

// NewGetExpenses200ResponseWithDefaults instantiates a new GetExpenses200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetExpenses200ResponseWithDefaults() *GetExpenses200Response {
	this := GetExpenses200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetExpenses200Response) GetData() []Expense {
	if o == nil || IsNil(o.Data) {
		var ret []Expense
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExpenses200Response) GetDataOk() ([]Expense, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetExpenses200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Expense and assigns it to the Data field.
func (o *GetExpenses200Response) SetData(v []Expense) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetExpenses200Response) GetMeta() Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Meta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExpenses200Response) GetMetaOk() (Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return Meta{}, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetExpenses200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *GetExpenses200Response) SetMeta(v Meta) {
	o.Meta = v
}

func (o GetExpenses200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetExpenses200Response) ToMap() (map[string]any, error) {
	toSerialize := map[string]any{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetExpenses200Response struct {
	value *GetExpenses200Response
	isSet bool
}

func (v NullableGetExpenses200Response) Get() *GetExpenses200Response {
	return v.value
}

func (v *NullableGetExpenses200Response) Set(val *GetExpenses200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExpenses200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExpenses200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExpenses200Response(val *GetExpenses200Response) *NullableGetExpenses200Response {
	return &NullableGetExpenses200Response{value: val, isSet: true}
}

func (v NullableGetExpenses200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExpenses200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
